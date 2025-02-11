/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package agent

import (
	"bytes"
	"context"
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/google/renameio/v2"
	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/lib/defaults"
)

// Base paths for constructing namespaced directories.
const (
	teleportOptDir     = "/opt/teleport"
	systemdAdminDir    = "/etc/systemd/system"
	systemdPIDFile     = "/run/teleport.pid"
	needrestartConfDir = "/etc/needrestart/conf.d"
	versionsDirName    = "versions"
	lockFileName       = "update.lock"
	defaultNamespace   = "default"
	systemNamespace    = "system"
)

const (
	// deprecatedTimerName is the timer for the deprecated upgrader should be disabled on setup.
	deprecatedTimerName = "teleport-upgrade.timer"
)

const (
	updateServiceTemplate = `# teleport-update
# DO NOT EDIT THIS FILE
[Unit]
Description=Teleport auto-update service

[Service]
Type=oneshot
ExecStart={{.UpdaterCommand}}
`
	updateTimerTemplate = `# teleport-update
# DO NOT EDIT THIS FILE
[Unit]
Description=Teleport auto-update timer unit

[Timer]
OnActiveSec=1m
OnUnitActiveSec=5m
RandomizedDelaySec=1m

[Install]
WantedBy={{.TeleportService}}
`
	teleportDropInTemplate = `# teleport-update
# DO NOT EDIT THIS FILE
[Service]
Environment=TELEPORT_UPDATE_CONFIG_FILE={{.UpdaterConfigFile}}
`
	// This configuration sets the default value for needrestart-trigger automatic restarts for teleport.service to disabled.
	// Users may still choose to enable needrestart for teleport.service when installing packaging interactively (or via dpkg config),
	// but doing so will result in a hard restart that disconnects the agent whenever any dependent libraries are updated.
	// Other network services, like openvpn, follow this pattern.
	// It is possible to configure needrestart to trigger a soft restart (via restart.d script), but given that Teleport subprocesses
	// can use a wide variety of installed binaries (when executed by the user), this could trigger many unexpected reloads.
	needrestartConfTemplate = `$nrconf{override_rc}{qr(^{{replace .TeleportService "." "\\."}})} = 0;
`
)

type confParams struct {
	TeleportService   string
	UpdaterCommand    string
	UpdaterConfigFile string
}

// Namespace represents a namespace within various system paths for a isolated installation of Teleport.
type Namespace struct {
	log *slog.Logger
	// name of namespace
	name string
	// dataDir for Teleport
	dataDir string
	// linkDir for Teleport binaries (ns: /opt/teleport/myns/bin)
	linkDir string
	// versionsDir for Teleport versions (ns: /opt/teleport/myns/versions)
	versionsDir string
	// serviceFile for the Teleport systemd service (ns: /etc/systemd/system/teleport_myns.service)
	serviceFile string
	// configFile for Teleport config (ns: /etc/teleport_myns.yaml)
	configFile string
	// pidFile for Teleport (ns: /run/teleport_myns.pid)
	pidFile string
	// updaterLockFile for locking the updater (ns: /opt/teleport/myns/update.lock)
	updaterLockFile string
	// updaterConfigFile for configuring updates (ns: /opt/teleport/myns/update.yaml)
	updaterConfigFile string
	// updaterBinFile for the updater when linked (linkDir + name)
	updaterBinFile string
	// updaterServiceFile is the systemd service path for the updater
	updaterServiceFile string
	// updaterTimerFile is the systemd timer path for the updater
	updaterTimerFile string
	// dropInFile is the Teleport systemd drop-in path extending Teleport
	dropInFile string
	// needrestartConfFile is the path to needrestart configuration for Teleport
	needrestartConfFile string
}

var alphanum = regexp.MustCompile("^[a-zA-Z0-9-]*$")

// NewNamespace validates and returns a Namespace.
// Namespaces must be alphanumeric + `-`.
func NewNamespace(log *slog.Logger, name, dataDir, linkDir string) (*Namespace, error) {
	if name == defaultNamespace ||
		name == systemNamespace {
		return nil, trace.Errorf("namespace %s is reserved", name)
	}
	if !alphanum.MatchString(name) {
		return nil, trace.Errorf("invalid namespace name %s, must be alphanumeric", name)
	}
	if name == "" {
		if dataDir == "" {
			dataDir = defaults.DataDir
		}
		if linkDir == "" {
			linkDir = DefaultLinkDir
		}
		return &Namespace{
			log:                 log,
			name:                name,
			dataDir:             dataDir,
			linkDir:             linkDir,
			versionsDir:         filepath.Join(namespaceDir(name), versionsDirName),
			serviceFile:         filepath.Join("/", serviceDir, serviceName),
			configFile:          defaults.ConfigFilePath,
			pidFile:             systemdPIDFile,
			updaterLockFile:     filepath.Join(namespaceDir(name), lockFileName),
			updaterConfigFile:   filepath.Join(namespaceDir(name), updateConfigName),
			updaterBinFile:      filepath.Join(linkDir, BinaryName),
			updaterServiceFile:  filepath.Join(systemdAdminDir, BinaryName+".service"),
			updaterTimerFile:    filepath.Join(systemdAdminDir, BinaryName+".timer"),
			dropInFile:          filepath.Join(systemdAdminDir, "teleport.service.d", BinaryName+".conf"),
			needrestartConfFile: filepath.Join(needrestartConfDir, BinaryName+".conf"),
		}, nil
	}

	prefix := "teleport_" + name
	if dataDir == "" {
		dataDir = filepath.Join(filepath.Dir(defaults.DataDir), prefix)
	}
	if linkDir == "" {
		linkDir = filepath.Join(namespaceDir(name), "bin")
	}
	return &Namespace{
		log:                 log,
		name:                name,
		dataDir:             dataDir,
		linkDir:             linkDir,
		versionsDir:         filepath.Join(namespaceDir(name), versionsDirName),
		serviceFile:         filepath.Join(systemdAdminDir, prefix+".service"),
		configFile:          filepath.Join(filepath.Dir(defaults.ConfigFilePath), prefix+".yaml"),
		pidFile:             filepath.Join(filepath.Dir(systemdPIDFile), prefix+".pid"),
		updaterLockFile:     filepath.Join(namespaceDir(name), lockFileName),
		updaterConfigFile:   filepath.Join(namespaceDir(name), updateConfigName),
		updaterBinFile:      filepath.Join(linkDir, BinaryName),
		updaterServiceFile:  filepath.Join(systemdAdminDir, BinaryName+"_"+name+".service"),
		updaterTimerFile:    filepath.Join(systemdAdminDir, BinaryName+"_"+name+".timer"),
		dropInFile:          filepath.Join(systemdAdminDir, "teleport.service.d", BinaryName+"_"+name+".conf"),
		needrestartConfFile: filepath.Join(needrestartConfDir, BinaryName+"_"+name+".conf"),
	}, nil
}

func namespaceDir(name string) string {
	if name == "" {
		name = defaultNamespace
	}
	return filepath.Join(teleportOptDir, name)
}

// Init create the initial directory structure and returns the lockfile for a Namespace.
func (ns *Namespace) Init() (lockFile string, err error) {
	if err := os.MkdirAll(ns.versionsDir, systemDirMode); err != nil {
		return "", trace.Wrap(err)
	}
	return ns.updaterLockFile, nil
}

// Setup installs service and timer files for the teleport-update binary.
// Afterwords, Setup reloads systemd and enables the timer with --now.
func (ns *Namespace) Setup(ctx context.Context) error {
	if ok, err := hasSystemD(); err == nil && !ok {
		ns.log.WarnContext(ctx, "Systemd is not running, skipping updater installation.")
		return nil
	}

	err := ns.writeConfigFiles(ctx)
	if err != nil {
		return trace.Wrap(err, "failed to write teleport-update systemd config files")
	}
	timer := &SystemdService{
		ServiceName: filepath.Base(ns.updaterTimerFile),
		Log:         ns.log,
	}
	if err := timer.Sync(ctx); err != nil {
		return trace.Wrap(err, "failed to sync systemd config")
	}
	if err := timer.Enable(ctx, true); err != nil {
		return trace.Wrap(err, "failed to enable teleport-update systemd timer")
	}
	if ns.name == "" {
		oldTimer := &SystemdService{
			ServiceName: deprecatedTimerName,
			Log:         ns.log,
		}
		// If the old teleport-upgrade script is detected, disable it to ensure they do not interfere.
		// Note that the schedule is also set to nop by the Teleport agent -- this just prevents restarts.
		enabled, err := isActiveOrEnabled(ctx, oldTimer)
		if err != nil {
			return trace.Wrap(err, "failed to determine if deprecated teleport-upgrade systemd timer is enabled")
		}
		if enabled {
			if err := oldTimer.Disable(ctx, true); err != nil {
				ns.log.ErrorContext(ctx, "The deprecated teleport-ent-updater package is installed on this server, and it cannot be disabled due to an error. You must remove the teleport-ent-updater package after verifying that teleport-update is working.", errorKey, err)
			} else {
				ns.log.WarnContext(ctx, "The deprecated teleport-ent-updater package is installed on this server. This package has been disabled to prevent conflicts. Please remove the teleport-ent-updater package after verifying that teleport-update is working.")
			}
		}
	}
	return nil
}

// Teardown removes all traces of the auto-updater, including its configuration.
func (ns *Namespace) Teardown(ctx context.Context) error {
	if ok, err := hasSystemD(); err == nil && !ok {
		ns.log.WarnContext(ctx, "Systemd is not running, skipping updater removal.")
		if err := os.RemoveAll(namespaceDir(ns.name)); err != nil {
			return trace.Wrap(err, "failed to remove versions directory")
		}
		return nil
	}

	svc := &SystemdService{
		ServiceName: filepath.Base(ns.updaterTimerFile),
		Log:         ns.log,
	}
	if err := svc.Disable(ctx, true); err != nil {
		return trace.Wrap(err, "failed to disable teleport-update systemd timer")
	}
	for _, p := range []string{
		ns.updaterServiceFile,
		ns.updaterTimerFile,
		ns.dropInFile,
		ns.needrestartConfFile,
	} {
		if err := os.Remove(p); err != nil && !errors.Is(err, fs.ErrNotExist) {
			return trace.Wrap(err, "failed to remove %s", filepath.Base(p))
		}
	}
	if err := svc.Sync(ctx); err != nil {
		return trace.Wrap(err, "failed to sync systemd config")
	}
	if err := os.RemoveAll(namespaceDir(ns.name)); err != nil {
		return trace.Wrap(err, "failed to remove versions directory")
	}
	if ns.name == "" {
		oldTimer := &SystemdService{
			ServiceName: deprecatedTimerName,
			Log:         ns.log,
		}
		// If the old upgrader exists, attempt to re-enable it automatically
		present, err := oldTimer.IsPresent(ctx)
		if err != nil {
			return trace.Wrap(err, "failed to determine if deprecated teleport-upgrade systemd timer is present")
		}
		if present {
			if err := oldTimer.Enable(ctx, true); err != nil {
				ns.log.ErrorContext(ctx, "The deprecated teleport-ent-updater package is installed on this server, and it cannot be re-enabled due to an error. Please fix the teleport-ent-updater package if you intend to use the deprecated updater.", errorKey, err)
			} else {
				ns.log.WarnContext(ctx, "The deprecated teleport-ent-updater package is installed on this server. This package has been re-enabled to ensure continued updates. To disable automatic updates entirely, please remove the teleport-ent-updater package.")
			}
		}
	}
	return nil
}

func (ns *Namespace) writeConfigFiles(ctx context.Context) error {
	var args string
	if ns.name != "" {
		args = " --install-suffix=" + ns.name
	}
	teleportService := filepath.Base(ns.serviceFile)
	params := confParams{
		TeleportService:   teleportService,
		UpdaterCommand:    ns.updaterBinFile + args + " update",
		UpdaterConfigFile: ns.updaterConfigFile,
	}
	err := writeSystemTemplate(ns.updaterServiceFile, updateServiceTemplate, params)
	if err != nil {
		return trace.Wrap(err)
	}
	err = writeSystemTemplate(ns.updaterTimerFile, updateTimerTemplate, params)
	if err != nil {
		return trace.Wrap(err)
	}
	err = writeSystemTemplate(ns.dropInFile, teleportDropInTemplate, params)
	if err != nil {
		return trace.Wrap(err)
	}
	// Needrestart config is non-critical for updater functionality.
	_, err = os.Stat(filepath.Dir(ns.needrestartConfFile))
	if os.IsNotExist(err) {
		return nil // needrestart is not present
	}
	if err != nil {
		ns.log.ErrorContext(ctx, "Unable to disable needrestart.", errorKey, err)
		return nil
	}
	ns.log.InfoContext(ctx, "Disabling needrestart.", unitKey, teleportService)
	err = writeSystemTemplate(ns.needrestartConfFile, needrestartConfTemplate, params)
	if err != nil {
		ns.log.ErrorContext(ctx, "Unable to disable needrestart.", errorKey, err)
		return nil
	}
	return nil
}

// writeSystemTemplate atomically writes a template to a system file, creating any needed directories.
// Temporarily files are stored in the target path to ensure the file has needed SELinux contexts.
func writeSystemTemplate(path, t string, values any) error {
	dir, file := filepath.Split(path)
	if err := os.MkdirAll(dir, systemDirMode); err != nil {
		return trace.Wrap(err)
	}
	opts := []renameio.Option{
		renameio.WithPermissions(configFileMode),
		renameio.WithExistingPermissions(),
		renameio.WithTempDir(dir),
	}
	f, err := renameio.NewPendingFile(path, opts...)
	if err != nil {
		return trace.Wrap(err)
	}
	defer f.Cleanup()

	tmpl, err := template.New(file).Funcs(template.FuncMap{
		"replace": func(s, old, new string) string {
			return strings.ReplaceAll(s, old, new)
		},
	}).Parse(t)
	if err != nil {
		return trace.Wrap(err)
	}
	err = tmpl.Execute(f, values)
	if err != nil {
		return trace.Wrap(err)
	}
	return trace.Wrap(f.CloseAtomicallyReplace())
}

// replaceTeleportService replaces the default paths in the Teleport service config with namespaced paths.
func (ns *Namespace) replaceTeleportService(cfg []byte) []byte {
	for _, rep := range []struct {
		old, new string
	}{
		{
			old: "/usr/local/bin/",
			new: ns.linkDir + "/",
		},
		{
			old: "/etc/teleport.yaml",
			new: ns.configFile,
		},
		{
			old: "/run/teleport.pid",
			new: ns.pidFile,
		},
	} {
		cfg = bytes.ReplaceAll(cfg, []byte(rep.old), []byte(rep.new))
	}
	return cfg
}

func (ns *Namespace) LogWarning(ctx context.Context) {
	ns.log.WarnContext(ctx, "Custom install suffix specified. Teleport data_dir must be configured in the config file.",
		"data_dir", ns.dataDir,
		"path", ns.linkDir,
		"config", ns.configFile,
		"service", filepath.Base(ns.serviceFile),
		"pid", ns.pidFile,
	)
}

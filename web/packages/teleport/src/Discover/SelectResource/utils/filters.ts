/**
 * Teleport
 * Copyright (C) 2025  Gravitational, Inc.
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

import { Platform } from 'design/platform';

import { AuthType } from 'teleport/services/user';

import { SelectResourceSpec } from '../resources';

export function filterBySupportedPlatformsAndAuthTypes(
  platform: Platform,
  authType: AuthType,
  resources: SelectResourceSpec[]
) {
  return resources.filter(resource => {
    const resourceSupportsPlatform =
      !resource.supportedPlatforms?.length ||
      resource.supportedPlatforms.includes(platform);

    const resourceSupportsAuthType =
      !resource.supportedAuthTypes?.length ||
      resource.supportedAuthTypes.includes(authType);

    return resourceSupportsPlatform && resourceSupportsAuthType;
  });
}

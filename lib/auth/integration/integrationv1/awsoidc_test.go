/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
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

package integrationv1

import (
	"testing"

	"github.com/gravitational/trace"
	"github.com/stretchr/testify/require"

	integrationv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/integration/v1"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/defaults"
	"github.com/gravitational/teleport/lib/integrations/awsoidc"
	"github.com/gravitational/teleport/lib/jwt"
	"github.com/gravitational/teleport/lib/utils"
)

func TestGenerateAWSOIDCToken(t *testing.T) {
	t.Parallel()
	clusterName := "test-cluster"
	integrationNameWithoutIssuer := "my-integration-without-issuer"
	integrationNameWithIssuer := "my-integration-with-issuer"

	publicURL := "https://example.com"
	s3BucketURI := "s3://mybucket/my-idp"
	s3IssuerURL := "https://mybucket.s3.amazonaws.com/my-idp"

	ca := newCertAuthority(t, types.HostCA, clusterName)
	ctx, localClient, resourceSvc := initSvc(t, ca, clusterName, publicURL)

	ig, err := types.NewIntegrationAWSOIDC(
		types.Metadata{Name: integrationNameWithoutIssuer},
		&types.AWSOIDCIntegrationSpecV1{
			RoleARN: "arn:aws:iam::123456789012:role/OpsTeam",
		},
	)
	require.NoError(t, err)
	_, err = localClient.CreateIntegration(ctx, ig)
	require.NoError(t, err)

	ig, err = types.NewIntegrationAWSOIDC(
		types.Metadata{Name: integrationNameWithIssuer},
		&types.AWSOIDCIntegrationSpecV1{
			RoleARN:     "arn:aws:iam::123456789012:role/OpsTeam",
			IssuerS3URI: s3BucketURI,
		},
	)
	require.NoError(t, err)
	_, err = localClient.CreateIntegration(ctx, ig)
	require.NoError(t, err)

	ctx = authorizerForDummyUser(t, ctx, types.RoleSpecV6{
		Allow: types.RoleConditions{Rules: []types.Rule{
			{Resources: []string{types.KindIntegration}, Verbs: []string{types.VerbUse}},
		}},
	}, localClient)

	// Get Public Key
	require.NotEmpty(t, ca.GetActiveKeys().JWT)
	jwtPubKey := ca.GetActiveKeys().JWT[0].PublicKey

	publicKey, err := utils.ParsePublicKey(jwtPubKey)
	require.NoError(t, err)

	// Validate JWT against public key
	key, err := jwt.New(&jwt.Config{
		Algorithm:   defaults.ApplicationTokenAlgorithm,
		ClusterName: clusterName,
		Clock:       resourceSvc.clock,
		PublicKey:   publicKey,
	})
	require.NoError(t, err)

	t.Run("without integration (old clients)", func(t *testing.T) {
		resp, err := resourceSvc.GenerateAWSOIDCToken(ctx, &integrationv1.GenerateAWSOIDCTokenRequest{})
		require.NoError(t, err)

		_, err = key.VerifyAWSOIDC(jwt.AWSOIDCVerifyParams{
			RawToken: resp.GetToken(),
			Issuer:   publicURL,
		})
		require.NoError(t, err)
		// Fails if the issuer is different
		_, err = key.VerifyAWSOIDC(jwt.AWSOIDCVerifyParams{
			RawToken: resp.GetToken(),
			Issuer:   publicURL + "3",
		})
		require.Error(t, err)
	})
	t.Run("with integration in rpc call but no issuer defined", func(t *testing.T) {
		resp, err := resourceSvc.GenerateAWSOIDCToken(ctx, &integrationv1.GenerateAWSOIDCTokenRequest{
			Integration: integrationNameWithoutIssuer,
		})
		require.NoError(t, err)

		_, err = key.VerifyAWSOIDC(jwt.AWSOIDCVerifyParams{
			RawToken: resp.GetToken(),
			Issuer:   publicURL,
		})
		require.NoError(t, err)
		// Fails if the issuer is different
		_, err = key.VerifyAWSOIDC(jwt.AWSOIDCVerifyParams{
			RawToken: resp.GetToken(),
			Issuer:   publicURL + "3",
		})
		require.Error(t, err)
	})
	t.Run("with integration in rpc call and issuer defined", func(t *testing.T) {
		resp, err := resourceSvc.GenerateAWSOIDCToken(ctx, &integrationv1.GenerateAWSOIDCTokenRequest{
			Integration: integrationNameWithIssuer,
		})
		require.NoError(t, err)

		_, err = key.VerifyAWSOIDC(jwt.AWSOIDCVerifyParams{
			RawToken: resp.GetToken(),
			Issuer:   s3IssuerURL,
		})
		require.NoError(t, err)
		// Fails if the issuer is different
		_, err = key.VerifyAWSOIDC(jwt.AWSOIDCVerifyParams{
			RawToken: resp.GetToken(),
			Issuer:   publicURL,
		})
		require.Error(t, err)
	})
}

func TestConvertSecurityGroupRulesToProto(t *testing.T) {
	for _, tt := range []struct {
		name     string
		in       []awsoidc.SecurityGroupRule
		expected []*integrationv1.SecurityGroupRule
	}{
		{
			name: "valid",
			in: []awsoidc.SecurityGroupRule{{
				IPProtocol: "tcp",
				FromPort:   8080,
				ToPort:     8081,
				CIDRs: []awsoidc.CIDR{{
					CIDR:        "10.10.10.0/24",
					Description: "cidr x",
				}},
			}},
			expected: []*integrationv1.SecurityGroupRule{{
				IpProtocol: "tcp",
				FromPort:   8080,
				ToPort:     8081,
				Cidrs: []*integrationv1.SecurityGroupRuleCIDR{{
					Cidr:        "10.10.10.0/24",
					Description: "cidr x",
				}},
			}},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			out := convertSecurityGroupRulesToProto(tt.in)
			require.Equal(t, tt.expected, out)
		})
	}
}

func TestListEICE(t *testing.T) {
	t.Parallel()

	clusterName := "test-cluster"
	proxyPublicAddr := "127.0.0.1.nip.io"
	integrationName := "my-awsoidc-integration"
	ig, err := types.NewIntegrationAWSOIDC(
		types.Metadata{Name: integrationName},
		&types.AWSOIDCIntegrationSpecV1{
			RoleARN: "arn:aws:iam::123456789012:role/OpsTeam",
		},
	)
	require.NoError(t, err)

	ca := newCertAuthority(t, types.HostCA, clusterName)
	ctx, localClient, resourceSvc := initSvc(t, ca, clusterName, proxyPublicAddr)

	_, err = localClient.CreateIntegration(ctx, ig)
	require.NoError(t, err)

	awsoidService, err := NewAWSOIDCService(&AWSOIDCServiceConfig{
		IntegrationService:    resourceSvc,
		Authorizer:            resourceSvc.authorizer,
		ProxyPublicAddrGetter: func() string { return "128.0.0.1" },
		Cache:                 &mockCache{},
	})
	require.NoError(t, err)

	t.Run("fails when user doesn't have access to integration.use", func(t *testing.T) {
		role := types.RoleSpecV6{
			Allow: types.RoleConditions{Rules: []types.Rule{{
				Resources: []string{types.KindIntegration},
				Verbs:     []string{types.VerbRead},
			}}},
		}

		userCtx := authorizerForDummyUser(t, ctx, role, localClient)

		_, err = awsoidService.ListEICE(userCtx, &integrationv1.ListEICERequest{
			Integration: integrationName,
			Region:      "my-region",
			VpcIds:      []string{"vpc-123"},
			NextToken:   "",
		})
		require.True(t, trace.IsAccessDenied(err), "expected AccessDenied error, but got %T", err)
	})
	t.Run("calls awsoidc package when user has access to integration.use/read", func(t *testing.T) {
		role := types.RoleSpecV6{
			Allow: types.RoleConditions{Rules: []types.Rule{{
				Resources: []string{types.KindIntegration},
				Verbs:     []string{types.VerbRead, types.VerbUse},
			}}},
		}

		userCtx := authorizerForDummyUser(t, ctx, role, localClient)

		_, err = awsoidService.ListEICE(userCtx, &integrationv1.ListEICERequest{
			Integration: integrationName,
			Region:      "",
			VpcIds:      []string{"vpc-123"},
			NextToken:   "",
		})
		require.True(t, trace.IsBadParameter(err), "expected BadParameter error, but got %T", err)
	})
}

func TestEnrollEKSClusters(t *testing.T) {
	t.Parallel()

	clusterName := "test-cluster"
	proxyPublicAddr := "127.0.0.1"
	integrationName := "my-awsoidc-integration"
	ig, err := types.NewIntegrationAWSOIDC(
		types.Metadata{Name: integrationName},
		&types.AWSOIDCIntegrationSpecV1{
			RoleARN: "arn:aws:iam::123456789012:role/OpsTeam",
		},
	)
	require.NoError(t, err)

	ca := newCertAuthority(t, types.HostCA, clusterName)
	ctx, localClient, resourceSvc := initSvc(t, ca, clusterName, proxyPublicAddr)

	_, err = localClient.CreateIntegration(ctx, ig)
	require.NoError(t, err)

	awsoidService, err := NewAWSOIDCService(&AWSOIDCServiceConfig{
		IntegrationService:    resourceSvc,
		Authorizer:            resourceSvc.authorizer,
		ProxyPublicAddrGetter: func() string { return "128.0.0.1" },
		Cache:                 &mockCache{},
	})
	require.NoError(t, err)

	t.Run("fails when user doesn't have access to integration.use", func(t *testing.T) {
		role := types.RoleSpecV6{
			Allow: types.RoleConditions{Rules: []types.Rule{{
				Resources: []string{types.KindIntegration},
				Verbs:     []string{types.VerbRead},
			}}},
		}

		userCtx := authorizerForDummyUser(t, ctx, role, localClient)

		_, err = awsoidService.EnrollEKSClusters(userCtx, &integrationv1.EnrollEKSClustersRequest{
			Integration:     integrationName,
			Region:          "my-region",
			EksClusterNames: []string{"EKS1"},
			AgentVersion:    "10.0.0",
		})
		require.True(t, trace.IsAccessDenied(err), "expected AccessDenied error, but got %T", err)
	})
	t.Run("calls awsoidc package when user has access to integration.use/read", func(t *testing.T) {
		role := types.RoleSpecV6{
			Allow: types.RoleConditions{Rules: []types.Rule{{
				Resources: []string{types.KindIntegration},
				Verbs:     []string{types.VerbRead, types.VerbUse},
			}}},
		}

		userCtx := authorizerForDummyUser(t, ctx, role, localClient)

		_, err := awsoidService.EnrollEKSClusters(userCtx, &integrationv1.EnrollEKSClustersRequest{
			Integration:     integrationName,
			Region:          "my-region",
			EksClusterNames: []string{"EKS1"},
		})
		require.True(t, trace.IsBadParameter(err), "expected BadParameter error, but got %T", err)
	})
}

func TestDeployService(t *testing.T) {
	t.Parallel()

	clusterName := "test-cluster"
	proxyPublicAddr := "127.0.0.1.nip.io"
	integrationName := "my-awsoidc-integration"
	ig, err := types.NewIntegrationAWSOIDC(
		types.Metadata{Name: integrationName},
		&types.AWSOIDCIntegrationSpecV1{
			RoleARN: "arn:aws:iam::123456789012:role/OpsTeam",
		},
	)
	require.NoError(t, err)

	ca := newCertAuthority(t, types.HostCA, clusterName)
	ctx, localClient, resourceSvc := initSvc(t, ca, clusterName, proxyPublicAddr)

	_, err = localClient.CreateIntegration(ctx, ig)
	require.NoError(t, err)

	awsoidService, err := NewAWSOIDCService(&AWSOIDCServiceConfig{
		IntegrationService:    resourceSvc,
		Authorizer:            resourceSvc.authorizer,
		ProxyPublicAddrGetter: func() string { return "128.0.0.1" },
		Cache:                 &mockCache{},
	})
	require.NoError(t, err)

	t.Run("fails when user doesn't have access to integration.use", func(t *testing.T) {
		role := types.RoleSpecV6{
			Allow: types.RoleConditions{Rules: []types.Rule{{
				Resources: []string{types.KindIntegration},
				Verbs:     []string{types.VerbRead},
			}}},
		}

		userCtx := authorizerForDummyUser(t, ctx, role, localClient)

		_, err = awsoidService.DeployService(userCtx, &integrationv1.DeployServiceRequest{
			Integration: integrationName,
			Region:      "my-region",
		})
		require.True(t, trace.IsAccessDenied(err), "expected AccessDenied error, but got %T", err)
	})
	t.Run("calls awsoidc package when user has access to integration.use/read", func(t *testing.T) {
		role := types.RoleSpecV6{
			Allow: types.RoleConditions{Rules: []types.Rule{{
				Resources: []string{types.KindIntegration},
				Verbs:     []string{types.VerbRead, types.VerbUse},
			}}},
		}

		userCtx := authorizerForDummyUser(t, ctx, role, localClient)

		_, err = awsoidService.DeployService(userCtx, &integrationv1.DeployServiceRequest{
			Integration: integrationName,
			Region:      "my-region",
		})
		require.True(t, trace.IsBadParameter(err), "expected BadParameter error, but got %T", err)
	})
}

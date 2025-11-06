// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/knocklabs/knock-go"
	"github.com/knocklabs/knock-go/internal/testutil"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/shared"
)

func TestTenantListWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Tenants.List(context.TODO(), knock.TenantListParams{
		After:    knock.F("after"),
		Before:   knock.F("before"),
		Name:     knock.F("name"),
		PageSize: knock.F(int64(0)),
		TenantID: knock.F("tenant_id"),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantDelete(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Tenants.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantGet(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Tenants.Get(context.TODO(), "id")
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantSetWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Tenants.Set(
		context.TODO(),
		"id",
		knock.TenantSetParams{
			ChannelData: knock.F(knock.InlineChannelDataRequestParam{
				"97c5837d-c65c-4d54-aa39-080eeb81c69d": knock.PushChannelDataTokensOnlyParam{
					Tokens: knock.F([]string{"push_token_xxx"}),
				},
			}),
			Name: knock.F("Jurassic Park"),
			Settings: knock.F(knock.TenantSetParamsSettings{
				Branding: knock.F(knock.TenantSetParamsSettingsBranding{
					IconURL:              knock.F("https://example.com/trex_silhouette_icon.png"),
					LogoURL:              knock.F("https://example.com/amber_fossil_logo.png"),
					PrimaryColor:         knock.F("#DF1A22"),
					PrimaryColorContrast: knock.F("#FFDE00"),
				}),
				PreferenceSet: knock.F(knock.PreferenceSetRequestParam{
					PersistenceStrategy: knock.F(knock.PreferenceSetRequest_PersistenceStrategyMerge),
					Categories: knock.F(map[string]knock.PreferenceSetRequestCategoriesUnionParam{
						"marketing": shared.UnionBool(false),
						"transactional": knock.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
							ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
								Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
								Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
								HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
								Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
								SMS: knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](knock.PreferenceSetChannelTypeSettingParam{
									Conditions: knock.F([]shared.ConditionParam{{
										Argument: knock.F("US"),
										Operator: knock.F(shared.ConditionOperatorEqualTo),
										Variable: knock.F("recipient.country_code"),
									}}),
								}),
							}),
							Channels: knock.F(map[string]knock.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectChannelsUnionParam{
								"aef6e715-df82-4ab6-b61e-b743e249f7b6": shared.UnionBool(true),
							}),
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("frog_genome"),
								Operator: knock.F(shared.ConditionOperatorContains),
								Variable: knock.F("specimen.dna_sequence"),
							}}),
						},
					}),
					ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
						Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
						Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
						HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
						Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
						SMS: knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](knock.PreferenceSetChannelTypeSettingParam{
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("US"),
								Operator: knock.F(shared.ConditionOperatorEqualTo),
								Variable: knock.F("recipient.country_code"),
							}}),
						}),
					}),
					Channels: knock.F(map[string]knock.PreferenceSetRequestChannelsUnionParam{
						"2f641633-95d3-4555-9222-9f1eb7888a80": knock.PreferenceSetChannelSettingParam{
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("US"),
								Operator: knock.F(shared.ConditionOperatorEqualTo),
								Variable: knock.F("recipient.country_code"),
							}}),
						},
						"aef6e715-df82-4ab6-b61e-b743e249f7b6": shared.UnionBool(true),
					}),
					CommercialSubscribed: knock.F(true),
					Workflows: knock.F(map[string]knock.PreferenceSetRequestWorkflowsUnionParam{
						"dinosaurs-loose": knock.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
							ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
								Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
								Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
								HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
								Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
								SMS: knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](knock.PreferenceSetChannelTypeSettingParam{
									Conditions: knock.F([]shared.ConditionParam{{
										Argument: knock.F("US"),
										Operator: knock.F(shared.ConditionOperatorEqualTo),
										Variable: knock.F("recipient.country_code"),
									}}),
								}),
							}),
							Channels: knock.F(map[string]knock.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelsUnionParam{
								"aef6e715-df82-4ab6-b61e-b743e249f7b6": shared.UnionBool(true),
							}),
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("frog_genome"),
								Operator: knock.F(shared.ConditionOperatorContains),
								Variable: knock.F("specimen.dna_sequence"),
							}}),
						},
					}),
				}),
			}),
		},
	)
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

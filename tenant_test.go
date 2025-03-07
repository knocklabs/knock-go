// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/knock-go"
	"github.com/stainless-sdks/knock-go/internal/testutil"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
)

func TestTenantListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Tenants.List(context.TODO(), knock.TenantListParams{
		After:    knock.F("after"),
		Before:   knock.F("before"),
		PageSize: knock.F(int64(0)),
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
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Tenants.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTenantGet(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
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
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Tenants.Set(
		context.TODO(),
		"id",
		knock.TenantSetParams{
			ChannelData: knock.F(map[string]knock.TenantSetParamsChannelData{
				"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
					Data: knock.F[knock.TenantSetParamsChannelDataDataUnion](knock.TenantSetParamsChannelDataDataPushChannelData{
						Tokens: knock.F([]string{"push_token_xxx"}),
					}),
				},
			}),
			Preferences: knock.F(map[string]knock.TenantSetParamsPreferences{
				"default": {
					Categories: knock.F(map[string]knock.TenantSetParamsPreferencesCategoriesUnion{
						"transactional": knock.TenantSetParamsPreferencesCategoriesObject{
							ChannelTypes: knock.F(knock.TenantSetParamsPreferencesCategoriesObjectChannelTypes{
								Chat:      knock.F[knock.TenantSetParamsPreferencesCategoriesObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.TenantSetParamsPreferencesCategoriesObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.TenantSetParamsPreferencesCategoriesObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.TenantSetParamsPreferencesCategoriesObjectConditionsOperatorEqualTo),
								Variable: knock.F("recipient.property"),
							}}),
						},
					}),
					ChannelTypes: knock.F(knock.TenantSetParamsPreferencesChannelTypes{
						Chat:      knock.F[knock.TenantSetParamsPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.TenantSetParamsPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
						HTTP:      knock.F[knock.TenantSetParamsPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.TenantSetParamsPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.TenantSetParamsPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.TenantSetParamsPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Workflows: knock.F(map[string]knock.TenantSetParamsPreferencesWorkflowsUnion{
						"dinosaurs-loose": knock.TenantSetParamsPreferencesWorkflowsObject{
							ChannelTypes: knock.F(knock.TenantSetParamsPreferencesWorkflowsObjectChannelTypes{
								Chat:      knock.F[knock.TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailUnion](shared.UnionBool(true)),
								HTTP:      knock.F[knock.TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.TenantSetParamsPreferencesWorkflowsObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorEqualTo),
								Variable: knock.F("recipient.property"),
							}}),
						},
					}),
				},
			}),
			Settings: knock.F(knock.TenantSetParamsSettings{
				Branding: knock.F(knock.TenantSetParamsSettingsBranding{
					IconURL:              knock.F("https://example.com/icon.png"),
					LogoURL:              knock.F("https://example.com/logo.png"),
					PrimaryColor:         knock.F("#000000"),
					PrimaryColorContrast: knock.F("#FFFFFF"),
				}),
				PreferenceSet: knock.F(knock.TenantSetParamsSettingsPreferenceSet{
					Categories: knock.F(map[string]knock.TenantSetParamsSettingsPreferenceSetCategoriesUnion{
						"marketing": shared.UnionBool(false),
						"transactional": knock.TenantSetParamsSettingsPreferenceSetCategoriesObject{
							ChannelTypes: knock.F(knock.TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypes{
								Chat:      knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.TenantSetParamsSettingsPreferenceSetCategoriesObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorEqualTo),
								Variable: knock.F("recipient.property"),
							}}),
						},
					}),
					ChannelTypes: knock.F(knock.TenantSetParamsSettingsPreferenceSetChannelTypes{
						Chat:      knock.F[knock.TenantSetParamsSettingsPreferenceSetChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.TenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion](shared.UnionBool(true)),
						HTTP:      knock.F[knock.TenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.TenantSetParamsSettingsPreferenceSetChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.TenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Workflows: knock.F(map[string]knock.TenantSetParamsSettingsPreferenceSetWorkflowsUnion{
						"dinosaurs-loose": knock.TenantSetParamsSettingsPreferenceSetWorkflowsObject{
							ChannelTypes: knock.F(knock.TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypes{
								Chat:      knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.TenantSetParamsSettingsPreferenceSetWorkflowsObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorEqualTo),
								Variable: knock.F("recipient.property"),
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

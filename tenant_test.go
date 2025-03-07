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
						"transactional": knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject{
							ChannelTypes: knock.F(knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
								Chat:      knock.F[knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
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
						"dinosaurs-loose": knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject{
							ChannelTypes: knock.F(knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
								Chat:      knock.F[knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(true)),
								HTTP:      knock.F[knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
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
						"transactional": knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject{
							ChannelTypes: knock.F(knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
								Chat:      knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
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
						"dinosaurs-loose": knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject{
							ChannelTypes: knock.F(knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
								Chat:      knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
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

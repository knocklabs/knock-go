// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/knock-go"
	"github.com/stainless-sdks/knock-go/internal/testutil"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
)

func TestUserBulkDelete(t *testing.T) {
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
	_, err := client.Users.Bulk.Delete(context.TODO(), knock.UserBulkDeleteParams{
		QueryUserIDs: knock.F([]string{"string"}),
		BodyUserIDs:  knock.F([]string{"user_1", "user_2"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserBulkIdentify(t *testing.T) {
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
	_, err := client.Users.Bulk.Identify(context.TODO(), knock.UserBulkIdentifyParams{
		Users: knock.F([]knock.UserBulkIdentifyParamsUser{{
			ID: knock.F("user_1"),
			ChannelData: knock.F(map[string]knock.UserBulkIdentifyParamsUsersChannelData{
				"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
					Data: knock.F[knock.UserBulkIdentifyParamsUsersChannelDataDataUnion](knock.UserBulkIdentifyParamsUsersChannelDataDataPushChannelData{
						Tokens: knock.F([]string{"push_token_xxx"}),
					}),
				},
			}),
			CreatedAt: knock.F(time.Now()),
			Preferences: knock.F(map[string]knock.UserBulkIdentifyParamsUsersPreferences{
				"default": {
					Categories: knock.F(map[string]knock.UserBulkIdentifyParamsUsersPreferencesCategoriesUnion{
						"transactional": knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject{
							ChannelTypes: knock.F(knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
								Chat:      knock.F[knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
								Variable: knock.F("recipient.property"),
							}}),
						},
					}),
					ChannelTypes: knock.F(knock.UserBulkIdentifyParamsUsersPreferencesChannelTypes{
						Chat:      knock.F[knock.UserBulkIdentifyParamsUsersPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
						HTTP:      knock.F[knock.UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.UserBulkIdentifyParamsUsersPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Workflows: knock.F(map[string]knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsUnion{
						"dinosaurs-loose": knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject{
							ChannelTypes: knock.F(knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
								Chat:      knock.F[knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(true)),
								HTTP:      knock.F[knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
								Variable: knock.F("recipient.property"),
							}}),
						},
					}),
				},
			}),
		}}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserBulkSetPreferencesWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Bulk.SetPreferences(context.TODO(), knock.UserBulkSetPreferencesParams{
		Preferences: knock.F(knock.UserBulkSetPreferencesParamsPreferences{
			Categories: knock.F(map[string]knock.UserBulkSetPreferencesParamsPreferencesCategoriesUnion{
				"marketing": shared.UnionBool(false),
				"transactional": knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject{
					ChannelTypes: knock.F(knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
						Chat:      knock.F[knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
						HTTP:      knock.F[knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Conditions: knock.F([]knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
						Argument: knock.F("some_property"),
						Operator: knock.F(knock.UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
						Variable: knock.F("recipient.property"),
					}}),
				},
			}),
			ChannelTypes: knock.F(knock.UserBulkSetPreferencesParamsPreferencesChannelTypes{
				Chat:      knock.F[knock.UserBulkSetPreferencesParamsPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
				Email:     knock.F[knock.UserBulkSetPreferencesParamsPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
				HTTP:      knock.F[knock.UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
				InAppFeed: knock.F[knock.UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
				Push:      knock.F[knock.UserBulkSetPreferencesParamsPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
				SMS:       knock.F[knock.UserBulkSetPreferencesParamsPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
			}),
			Workflows: knock.F(map[string]knock.UserBulkSetPreferencesParamsPreferencesWorkflowsUnion{
				"dinosaurs-loose": knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject{
					ChannelTypes: knock.F(knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
						Chat:      knock.F[knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
						HTTP:      knock.F[knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Conditions: knock.F([]knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
						Argument: knock.F("some_property"),
						Operator: knock.F(knock.UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
						Variable: knock.F("recipient.property"),
					}}),
				},
			}),
		}),
		UserIDs: knock.F([]string{"user_1", "user_2"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

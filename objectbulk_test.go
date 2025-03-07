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

func TestObjectBulkDelete(t *testing.T) {
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
	_, err := client.Objects.Bulk.Delete(
		context.TODO(),
		"collection",
		knock.ObjectBulkDeleteParams{
			ObjectIDs: knock.F([]string{"string"}),
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

func TestObjectBulkAddSubscriptions(t *testing.T) {
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
	_, err := client.Objects.Bulk.AddSubscriptions(
		context.TODO(),
		"projects",
		knock.ObjectBulkAddSubscriptionsParams{
			Subscriptions: knock.F([]knock.ObjectBulkAddSubscriptionsParamsSubscription{{
				ID: knock.F("project-1"),
				Recipients: knock.F([]knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion{knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequest{
					ID: knock.F("user_1"),
					ChannelData: knock.F(map[string]knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelData{
						"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
							Data: knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion](knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataPushChannelData{
								Tokens: knock.F([]string{"push_token_xxx"}),
							}),
						},
					}),
					CreatedAt: knock.F(time.Now()),
					Preferences: knock.F(map[string]knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferences{
						"default": {
							Categories: knock.F(map[string]knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion{
								"transactional": knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject{
									ChannelTypes: knock.F(knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
										Chat:      knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
										Email:     knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
										HTTP:      knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
										InAppFeed: knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
										Push:      knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
										SMS:       knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
										Argument: knock.F("some_property"),
										Operator: knock.F(knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
										Variable: knock.F("recipient.property"),
									}}),
								},
							}),
							ChannelTypes: knock.F(knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypes{
								Chat:      knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
								HTTP:      knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Workflows: knock.F(map[string]knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion{
								"dinosaurs-loose": knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject{
									ChannelTypes: knock.F(knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
										Chat:      knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
										Email:     knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(true)),
										HTTP:      knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
										InAppFeed: knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
										Push:      knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
										SMS:       knock.F[knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
										Argument: knock.F("some_property"),
										Operator: knock.F(knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
										Variable: knock.F("recipient.property"),
									}}),
								},
							}),
						},
					}),
				}}),
				Properties: knock.F(map[string]interface{}{
					"foo": "bar",
				}),
			}}),
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

func TestObjectBulkSet(t *testing.T) {
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
	_, err := client.Objects.Bulk.Set(
		context.TODO(),
		"collection",
		knock.ObjectBulkSetParams{
			Objects: knock.F([]knock.ObjectBulkSetParamsObject{{
				ID:         knock.F("project_1"),
				Collection: knock.F("projects"),
				ChannelData: knock.F(map[string]knock.ObjectBulkSetParamsObjectsChannelData{
					"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
						Data: knock.F[knock.ObjectBulkSetParamsObjectsChannelDataDataUnion](knock.ObjectBulkSetParamsObjectsChannelDataDataPushChannelData{
							Tokens: knock.F([]string{"push_token_xxx"}),
						}),
					},
				}),
				CreatedAt: knock.F(time.Now()),
				Preferences: knock.F(map[string]knock.ObjectBulkSetParamsObjectsPreferences{
					"default": {
						Categories: knock.F(map[string]knock.ObjectBulkSetParamsObjectsPreferencesCategoriesUnion{
							"transactional": knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject{
								ChannelTypes: knock.F(knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
									Chat:      knock.F[knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
									Email:     knock.F[knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
									HTTP:      knock.F[knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
									InAppFeed: knock.F[knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
									Push:      knock.F[knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
									SMS:       knock.F[knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
									Argument: knock.F("some_property"),
									Operator: knock.F(knock.ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
									Variable: knock.F("recipient.property"),
								}}),
							},
						}),
						ChannelTypes: knock.F(knock.ObjectBulkSetParamsObjectsPreferencesChannelTypes{
							Chat:      knock.F[knock.ObjectBulkSetParamsObjectsPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
							Email:     knock.F[knock.ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
							HTTP:      knock.F[knock.ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
							InAppFeed: knock.F[knock.ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
							Push:      knock.F[knock.ObjectBulkSetParamsObjectsPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
							SMS:       knock.F[knock.ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
						}),
						Workflows: knock.F(map[string]knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsUnion{
							"dinosaurs-loose": knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject{
								ChannelTypes: knock.F(knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
									Chat:      knock.F[knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
									Email:     knock.F[knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(true)),
									HTTP:      knock.F[knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
									InAppFeed: knock.F[knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
									Push:      knock.F[knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
									SMS:       knock.F[knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
									Argument: knock.F("some_property"),
									Operator: knock.F(knock.ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
									Variable: knock.F("recipient.property"),
								}}),
							},
						}),
					},
				}),
			}}),
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

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

func TestObjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Objects.List(
		context.TODO(),
		"collection",
		knock.ObjectListParams{
			After:    knock.F("after"),
			Before:   knock.F("before"),
			PageSize: knock.F(int64(0)),
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

func TestObjectDelete(t *testing.T) {
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
	_, err := client.Objects.Delete(
		context.TODO(),
		"collection",
		"id",
	)
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectAddSubscriptionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Objects.AddSubscriptions(
		context.TODO(),
		"collection",
		"object_id",
		knock.ObjectAddSubscriptionsParams{
			Recipients: knock.F([]knock.ObjectAddSubscriptionsParamsRecipientUnion{shared.UnionString("user_1"), shared.UnionString("user_2")}),
			Properties: knock.F(map[string]interface{}{
				"key": "bar",
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

func TestObjectDeleteSubscriptions(t *testing.T) {
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
	_, err := client.Objects.DeleteSubscriptions(
		context.TODO(),
		"collection",
		"object_id",
		knock.ObjectDeleteSubscriptionsParams{
			Recipients: knock.F([]knock.ObjectDeleteSubscriptionsParamsRecipientUnion{knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequest{
				ID: knock.F("user_1"),
				ChannelData: knock.F(map[string]knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestChannelData{
					"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
						Data: knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestChannelDataDataUnion](knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestChannelDataDataPushChannelData{
							Tokens: knock.F([]string{"push_token_xxx"}),
						}),
					},
				}),
				CreatedAt: knock.F(time.Now()),
				Preferences: knock.F(map[string]knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferences{
					"default": {
						Categories: knock.F(map[string]knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion{
							"transactional": knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject{
								ChannelTypes: knock.F(knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
									Chat:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
									Email:     knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
									HTTP:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
									InAppFeed: knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
									Push:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
									SMS:       knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
									Argument: knock.F("some_property"),
									Operator: knock.F(knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
									Variable: knock.F("recipient.property"),
								}}),
							},
						}),
						ChannelTypes: knock.F(knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypes{
							Chat:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
							Email:     knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
							HTTP:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
							InAppFeed: knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
							Push:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
							SMS:       knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
						}),
						Workflows: knock.F(map[string]knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion{
							"dinosaurs-loose": knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject{
								ChannelTypes: knock.F(knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
									Chat:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
									Email:     knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(true)),
									HTTP:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
									InAppFeed: knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
									Push:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
									SMS:       knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
									Argument: knock.F("some_property"),
									Operator: knock.F(knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
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

func TestObjectGet(t *testing.T) {
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
	_, err := client.Objects.Get(
		context.TODO(),
		"collection",
		"id",
	)
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectGetChannelData(t *testing.T) {
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
	_, err := client.Objects.GetChannelData(
		context.TODO(),
		"collection",
		"object_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectGetPreferencesWithOptionalParams(t *testing.T) {
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
	_, err := client.Objects.GetPreferences(
		context.TODO(),
		"collection",
		"object_id",
		"id",
		knock.ObjectGetPreferencesParams{
			Tenant: knock.F("tenant"),
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

func TestObjectListMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Objects.ListMessages(
		context.TODO(),
		"projects",
		"project-123",
		knock.ObjectListMessagesParams{
			After:                  knock.F("after"),
			Before:                 knock.F("before"),
			ChannelID:              knock.F("channel_id"),
			EngagementStatus:       knock.F([]knock.ObjectListMessagesParamsEngagementStatus{knock.ObjectListMessagesParamsEngagementStatusSeen}),
			MessageIDs:             knock.F([]string{"string"}),
			PageSize:               knock.F(int64(0)),
			Source:                 knock.F("source"),
			Status:                 knock.F([]knock.ObjectListMessagesParamsStatus{knock.ObjectListMessagesParamsStatusQueued}),
			Tenant:                 knock.F("tenant"),
			TriggerData:            knock.F("trigger_data"),
			WorkflowCategories:     knock.F([]string{"string"}),
			WorkflowRecipientRunID: knock.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			WorkflowRunID:          knock.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestObjectListPreferences(t *testing.T) {
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
	_, err := client.Objects.ListPreferences(
		context.TODO(),
		"collection",
		"object_id",
	)
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectListSchedulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Objects.ListSchedules(
		context.TODO(),
		"collection",
		"id",
		knock.ObjectListSchedulesParams{
			After:    knock.F("after"),
			Before:   knock.F("before"),
			PageSize: knock.F(int64(0)),
			Tenant:   knock.F("tenant"),
			Workflow: knock.F("workflow"),
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

func TestObjectListSubscriptionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Objects.ListSubscriptions(
		context.TODO(),
		"collection",
		"object_id",
		knock.ObjectListSubscriptionsParams{
			After:      knock.F("after"),
			Before:     knock.F("before"),
			Mode:       knock.F(knock.ObjectListSubscriptionsParamsModeRecipient),
			PageSize:   knock.F(int64(0)),
			Recipients: knock.F([]knock.ObjectListSubscriptionsParamsRecipientUnion{shared.UnionString("user_123")}),
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

func TestObjectSetWithOptionalParams(t *testing.T) {
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
	_, err := client.Objects.Set(
		context.TODO(),
		"collection",
		"id",
		knock.ObjectSetParams{
			ChannelData: knock.F(map[string]knock.ObjectSetParamsChannelData{
				"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
					Data: knock.F[knock.ObjectSetParamsChannelDataDataUnion](knock.ObjectSetParamsChannelDataDataPushChannelData{
						Tokens: knock.F([]string{"push_token_xxx"}),
					}),
				},
			}),
			Preferences: knock.F(map[string]knock.ObjectSetParamsPreferences{
				"default": {
					Categories: knock.F(map[string]knock.ObjectSetParamsPreferencesCategoriesUnion{
						"transactional": knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject{
							ChannelTypes: knock.F(knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
								Chat:      knock.F[knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.ObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
								Variable: knock.F("recipient.property"),
							}}),
						},
					}),
					ChannelTypes: knock.F(knock.ObjectSetParamsPreferencesChannelTypes{
						Chat:      knock.F[knock.ObjectSetParamsPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.ObjectSetParamsPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
						HTTP:      knock.F[knock.ObjectSetParamsPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.ObjectSetParamsPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.ObjectSetParamsPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.ObjectSetParamsPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Workflows: knock.F(map[string]knock.ObjectSetParamsPreferencesWorkflowsUnion{
						"dinosaurs-loose": knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject{
							ChannelTypes: knock.F(knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
								Chat:      knock.F[knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(true)),
								HTTP:      knock.F[knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.ObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
								Variable: knock.F("recipient.property"),
							}}),
						},
					}),
				},
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

func TestObjectSetChannelDataWithOptionalParams(t *testing.T) {
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
	_, err := client.Objects.SetChannelData(
		context.TODO(),
		"collection",
		"object_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		knock.ObjectSetChannelDataParams{
			Data: knock.F[knock.ObjectSetChannelDataParamsDataUnion](knock.ObjectSetChannelDataParamsDataPushChannelData{
				Tokens: knock.F([]string{"push_token_1"}),
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

func TestObjectSetPreferencesWithOptionalParams(t *testing.T) {
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
	_, err := client.Objects.SetPreferences(
		context.TODO(),
		"collection",
		"object_id",
		"id",
		knock.ObjectSetPreferencesParams{
			Categories: knock.F(map[string]knock.ObjectSetPreferencesParamsCategoriesUnion{
				"marketing": shared.UnionBool(false),
				"transactional": knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObject{
					ChannelTypes: knock.F(knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
						Chat:      knock.F[knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
						HTTP:      knock.F[knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Conditions: knock.F([]knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
						Argument: knock.F("some_property"),
						Operator: knock.F(knock.ObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
						Variable: knock.F("recipient.property"),
					}}),
				},
			}),
			ChannelTypes: knock.F(knock.ObjectSetPreferencesParamsChannelTypes{
				Chat:      knock.F[knock.ObjectSetPreferencesParamsChannelTypesChatUnion](shared.UnionBool(true)),
				Email:     knock.F[knock.ObjectSetPreferencesParamsChannelTypesEmailUnion](shared.UnionBool(true)),
				HTTP:      knock.F[knock.ObjectSetPreferencesParamsChannelTypesHTTPUnion](shared.UnionBool(true)),
				InAppFeed: knock.F[knock.ObjectSetPreferencesParamsChannelTypesInAppFeedUnion](shared.UnionBool(true)),
				Push:      knock.F[knock.ObjectSetPreferencesParamsChannelTypesPushUnion](shared.UnionBool(true)),
				SMS:       knock.F[knock.ObjectSetPreferencesParamsChannelTypesSMSUnion](shared.UnionBool(true)),
			}),
			Workflows: knock.F(map[string]knock.ObjectSetPreferencesParamsWorkflowsUnion{
				"dinosaurs-loose": knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObject{
					ChannelTypes: knock.F(knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
						Chat:      knock.F[knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
						HTTP:      knock.F[knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Conditions: knock.F([]knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
						Argument: knock.F("some_property"),
						Operator: knock.F(knock.ObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
						Variable: knock.F("recipient.property"),
					}}),
				},
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

func TestObjectUnsetChannelData(t *testing.T) {
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
	_, err := client.Objects.UnsetChannelData(
		context.TODO(),
		"collection",
		"object_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

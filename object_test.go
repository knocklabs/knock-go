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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
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
							"transactional": knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObject{
								ChannelTypes: knock.F(knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObjectChannelTypes{
									Chat:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObjectChannelTypesChatUnion](shared.UnionBool(true)),
									Email:     knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObjectChannelTypesEmailUnion](shared.UnionBool(false)),
									HTTP:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
									InAppFeed: knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
									Push:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObjectChannelTypesPushUnion](shared.UnionBool(true)),
									SMS:       knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObjectChannelTypesSMSUnion](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObjectCondition{{
									Argument: knock.F("some_property"),
									Operator: knock.F(knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesObjectConditionsOperatorEqualTo),
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
							"dinosaurs-loose": knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObject{
								ChannelTypes: knock.F(knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObjectChannelTypes{
									Chat:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObjectChannelTypesChatUnion](shared.UnionBool(true)),
									Email:     knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObjectChannelTypesEmailUnion](shared.UnionBool(true)),
									HTTP:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
									InAppFeed: knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
									Push:      knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObjectChannelTypesPushUnion](shared.UnionBool(true)),
									SMS:       knock.F[knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObjectChannelTypesSMSUnion](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObjectCondition{{
									Argument: knock.F("some_property"),
									Operator: knock.F(knock.ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsObjectConditionsOperatorEqualTo),
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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
	)
	_, err := client.Objects.ListSubscriptions(
		context.TODO(),
		"collection",
		"object_id",
		knock.ObjectListSubscriptionsParams{
			After:      knock.F("after"),
			Before:     knock.F("before"),
			Mode:       knock.F(knock.ObjectListSubscriptionsParamsModeRecipient),
			Objects:    knock.F([]knock.ObjectListSubscriptionsParamsObjectUnion{shared.UnionString("user_123")}),
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
		option.WithToken("My Token"),
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
						"transactional": knock.ObjectSetParamsPreferencesCategoriesObject{
							ChannelTypes: knock.F(knock.ObjectSetParamsPreferencesCategoriesObjectChannelTypes{
								Chat:      knock.F[knock.ObjectSetParamsPreferencesCategoriesObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.ObjectSetParamsPreferencesCategoriesObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.ObjectSetParamsPreferencesCategoriesObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.ObjectSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.ObjectSetParamsPreferencesCategoriesObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.ObjectSetParamsPreferencesCategoriesObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.ObjectSetParamsPreferencesCategoriesObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.ObjectSetParamsPreferencesCategoriesObjectConditionsOperatorEqualTo),
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
						"dinosaurs-loose": knock.ObjectSetParamsPreferencesWorkflowsObject{
							ChannelTypes: knock.F(knock.ObjectSetParamsPreferencesWorkflowsObjectChannelTypes{
								Chat:      knock.F[knock.ObjectSetParamsPreferencesWorkflowsObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.ObjectSetParamsPreferencesWorkflowsObjectChannelTypesEmailUnion](shared.UnionBool(true)),
								HTTP:      knock.F[knock.ObjectSetParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.ObjectSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.ObjectSetParamsPreferencesWorkflowsObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.ObjectSetParamsPreferencesWorkflowsObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.ObjectSetParamsPreferencesWorkflowsObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.ObjectSetParamsPreferencesWorkflowsObjectConditionsOperatorEqualTo),
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

func TestObjectSetChannelData(t *testing.T) {
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
		option.WithToken("My Token"),
	)
	_, err := client.Objects.SetChannelData(
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
		option.WithToken("My Token"),
	)
	_, err := client.Objects.SetPreferences(
		context.TODO(),
		"collection",
		"object_id",
		"id",
		knock.ObjectSetPreferencesParams{
			Categories: knock.F(map[string]knock.ObjectSetPreferencesParamsCategoriesUnion{
				"marketing": shared.UnionBool(false),
				"transactional": knock.ObjectSetPreferencesParamsCategoriesObject{
					ChannelTypes: knock.F(knock.ObjectSetPreferencesParamsCategoriesObjectChannelTypes{
						Chat:      knock.F[knock.ObjectSetPreferencesParamsCategoriesObjectChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.ObjectSetPreferencesParamsCategoriesObjectChannelTypesEmailUnion](shared.UnionBool(false)),
						HTTP:      knock.F[knock.ObjectSetPreferencesParamsCategoriesObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.ObjectSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.ObjectSetPreferencesParamsCategoriesObjectChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.ObjectSetPreferencesParamsCategoriesObjectChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Conditions: knock.F([]knock.ObjectSetPreferencesParamsCategoriesObjectCondition{{
						Argument: knock.F("some_property"),
						Operator: knock.F(knock.ObjectSetPreferencesParamsCategoriesObjectConditionsOperatorEqualTo),
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
				"dinosaurs-loose": knock.ObjectSetPreferencesParamsWorkflowsObject{
					ChannelTypes: knock.F(knock.ObjectSetPreferencesParamsWorkflowsObjectChannelTypes{
						Chat:      knock.F[knock.ObjectSetPreferencesParamsWorkflowsObjectChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.ObjectSetPreferencesParamsWorkflowsObjectChannelTypesEmailUnion](shared.UnionBool(false)),
						HTTP:      knock.F[knock.ObjectSetPreferencesParamsWorkflowsObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.ObjectSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.ObjectSetPreferencesParamsWorkflowsObjectChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.ObjectSetPreferencesParamsWorkflowsObjectChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Conditions: knock.F([]knock.ObjectSetPreferencesParamsWorkflowsObjectCondition{{
						Argument: knock.F("some_property"),
						Operator: knock.F(knock.ObjectSetPreferencesParamsWorkflowsObjectConditionsOperatorEqualTo),
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
		option.WithToken("My Token"),
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

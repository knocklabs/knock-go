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
			Recipients: knock.F([]shared.RecipientRequestUnionParam{shared.UnionString("user_1"), shared.UnionString("user_2")}),
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
			Recipients: knock.F([]shared.RecipientRequestUnionParam{shared.InlineIdentifyUserRequestParam{
				ID: knock.F("user_1"),
				ChannelData: knock.F(shared.InlineChannelDataRequestParam{
					"97c5837d-c65c-4d54-aa39-080eeb81c69d": shared.ChannelDataRequestParam{
						Data: knock.F[shared.ChannelDataRequestDataUnionParam](shared.PushChannelDataParam{
							Tokens: knock.F([]string{"push_token_xxx"}),
						}),
					},
				}),
				CreatedAt: knock.F(time.Now()),
				Preferences: knock.F(shared.InlinePreferenceSetRequestParam{
					"default": shared.PreferenceSetRequestParam{
						Categories: knock.F(map[string]shared.PreferenceSetRequestCategoriesUnionParam{
							"transactional": shared.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
								ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
									Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
									Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
									HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
									InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
									Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
									SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]shared.ConditionParam{{
									Argument: knock.F("some_property"),
									Operator: knock.F(shared.ConditionOperatorEqualTo),
									Variable: knock.F("recipient.property"),
								}}),
							},
						}),
						ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
							Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
							Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
							HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
							InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
							Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
							SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
						}),
						Workflows: knock.F(map[string]shared.PreferenceSetRequestWorkflowsUnionParam{
							"dinosaurs-loose": shared.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
								ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
									Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
									Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
									HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
									InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
									Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
									SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]shared.ConditionParam{{
									Argument: knock.F("some_property"),
									Operator: knock.F(shared.ConditionOperatorEqualTo),
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
			ChannelData: knock.F(shared.InlineChannelDataRequestParam{
				"97c5837d-c65c-4d54-aa39-080eeb81c69d": shared.ChannelDataRequestParam{
					Data: knock.F[shared.ChannelDataRequestDataUnionParam](shared.PushChannelDataParam{
						Tokens: knock.F([]string{"push_token_xxx"}),
					}),
				},
			}),
			Preferences: knock.F(shared.InlinePreferenceSetRequestParam{
				"default": shared.PreferenceSetRequestParam{
					Categories: knock.F(map[string]shared.PreferenceSetRequestCategoriesUnionParam{
						"transactional": shared.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
							ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
								Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
								Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
								HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
								InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
								Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
								SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("some_property"),
								Operator: knock.F(shared.ConditionOperatorEqualTo),
								Variable: knock.F("recipient.property"),
							}}),
						},
					}),
					ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
						Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
						Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
						HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
						InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
						Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
						SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
					}),
					Workflows: knock.F(map[string]shared.PreferenceSetRequestWorkflowsUnionParam{
						"dinosaurs-loose": shared.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
							ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
								Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
								Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
								HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
								InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
								Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
								SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("some_property"),
								Operator: knock.F(shared.ConditionOperatorEqualTo),
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
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Objects.SetChannelData(
		context.TODO(),
		"collection",
		"object_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		knock.ObjectSetChannelDataParams{
			ChannelDataRequest: shared.ChannelDataRequestParam{
				Data: knock.F[shared.ChannelDataRequestDataUnionParam](shared.PushChannelDataParam{
					Tokens: knock.F([]string{"push_token_1"}),
				}),
			},
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
			PreferenceSetRequest: shared.PreferenceSetRequestParam{
				Categories: knock.F(map[string]shared.PreferenceSetRequestCategoriesUnionParam{
					"marketing": shared.UnionBool(false),
					"transactional": shared.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
						ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
							Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
							Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
							HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
							InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
							Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
							SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
						}),
						Conditions: knock.F([]shared.ConditionParam{{
							Argument: knock.F("some_property"),
							Operator: knock.F(shared.ConditionOperatorEqualTo),
							Variable: knock.F("recipient.property"),
						}}),
					},
				}),
				ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
					Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
					Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
					HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
					InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
					Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
					SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
				}),
				Workflows: knock.F(map[string]shared.PreferenceSetRequestWorkflowsUnionParam{
					"dinosaurs-loose": shared.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
						ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
							Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
							Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
							HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
							InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
							Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
							SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
						}),
						Conditions: knock.F([]shared.ConditionParam{{
							Argument: knock.F("some_property"),
							Operator: knock.F(shared.ConditionOperatorEqualTo),
							Variable: knock.F("recipient.property"),
						}}),
					},
				}),
			},
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

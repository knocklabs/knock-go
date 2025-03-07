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

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Update(
		context.TODO(),
		"user_id",
		knock.UserUpdateParams{
			ChannelData: knock.F(map[string]knock.UserUpdateParamsChannelData{
				"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
					Data: knock.F[knock.UserUpdateParamsChannelDataDataUnion](knock.UserUpdateParamsChannelDataDataPushChannelData{
						Tokens: knock.F([]string{"push_token_xxx"}),
					}),
				},
			}),
			CreatedAt: knock.F(time.Now()),
			Preferences: knock.F(map[string]knock.UserUpdateParamsPreferences{
				"default": {
					Categories: knock.F(map[string]knock.UserUpdateParamsPreferencesCategoriesUnion{
						"transactional": knock.UserUpdateParamsPreferencesCategoriesObject{
							ChannelTypes: knock.F(knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypes{
								Chat:      knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion](shared.UnionBool(false)),
								HTTP:      knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.UserUpdateParamsPreferencesCategoriesObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEqualTo),
								Variable: knock.F("recipient.property"),
							}}),
						},
					}),
					ChannelTypes: knock.F(knock.UserUpdateParamsPreferencesChannelTypes{
						Chat:      knock.F[knock.UserUpdateParamsPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.UserUpdateParamsPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
						HTTP:      knock.F[knock.UserUpdateParamsPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.UserUpdateParamsPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.UserUpdateParamsPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.UserUpdateParamsPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Workflows: knock.F(map[string]knock.UserUpdateParamsPreferencesWorkflowsUnion{
						"dinosaurs-loose": knock.UserUpdateParamsPreferencesWorkflowsObject{
							ChannelTypes: knock.F(knock.UserUpdateParamsPreferencesWorkflowsObjectChannelTypes{
								Chat:      knock.F[knock.UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailUnion](shared.UnionBool(true)),
								HTTP:      knock.F[knock.UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Conditions: knock.F([]knock.UserUpdateParamsPreferencesWorkflowsObjectCondition{{
								Argument: knock.F("some_property"),
								Operator: knock.F(knock.UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorEqualTo),
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

func TestUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.List(context.TODO(), knock.UserListParams{
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

func TestUserDelete(t *testing.T) {
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
	_, err := client.Users.Delete(context.TODO(), "user_id")
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGet(t *testing.T) {
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
	_, err := client.Users.Get(context.TODO(), "user_id")
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGetChannelData(t *testing.T) {
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
	_, err := client.Users.GetChannelData(
		context.TODO(),
		"user_id",
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

func TestUserGetPreferencesWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.GetPreferences(
		context.TODO(),
		"user_id",
		"id",
		knock.UserGetPreferencesParams{
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

func TestUserListMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.ListMessages(
		context.TODO(),
		"user-123",
		knock.UserListMessagesParams{
			After:                  knock.F("after"),
			Before:                 knock.F("before"),
			ChannelID:              knock.F("channel_id"),
			EngagementStatus:       knock.F([]knock.UserListMessagesParamsEngagementStatus{knock.UserListMessagesParamsEngagementStatusSeen}),
			MessageIDs:             knock.F([]string{"string"}),
			PageSize:               knock.F(int64(0)),
			Source:                 knock.F("source"),
			Status:                 knock.F([]knock.UserListMessagesParamsStatus{knock.UserListMessagesParamsStatusQueued}),
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

func TestUserListPreferences(t *testing.T) {
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
	_, err := client.Users.ListPreferences(context.TODO(), "user_id")
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserListSchedulesWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.ListSchedules(
		context.TODO(),
		"user_id",
		knock.UserListSchedulesParams{
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

func TestUserListSubscriptionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.ListSubscriptions(
		context.TODO(),
		"user_id",
		knock.UserListSubscriptionsParams{
			After:    knock.F("after"),
			Before:   knock.F("before"),
			Objects:  knock.F([]knock.UserListSubscriptionsParamsObjectUnion{shared.UnionString("user_123")}),
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

func TestUserMergeWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Merge(
		context.TODO(),
		"user_id",
		knock.UserMergeParams{
			FromUserID: knock.F("user_1"),
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

func TestUserSetChannelData(t *testing.T) {
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
	_, err := client.Users.SetChannelData(
		context.TODO(),
		"user_id",
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

func TestUserSetPreferencesWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.SetPreferences(
		context.TODO(),
		"user_id",
		"id",
		knock.UserSetPreferencesParams{
			Categories: knock.F(map[string]knock.UserSetPreferencesParamsCategoriesUnion{
				"marketing": shared.UnionBool(false),
				"transactional": knock.UserSetPreferencesParamsCategoriesObject{
					ChannelTypes: knock.F(knock.UserSetPreferencesParamsCategoriesObjectChannelTypes{
						Chat:      knock.F[knock.UserSetPreferencesParamsCategoriesObjectChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.UserSetPreferencesParamsCategoriesObjectChannelTypesEmailUnion](shared.UnionBool(false)),
						HTTP:      knock.F[knock.UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.UserSetPreferencesParamsCategoriesObjectChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.UserSetPreferencesParamsCategoriesObjectChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Conditions: knock.F([]knock.UserSetPreferencesParamsCategoriesObjectCondition{{
						Argument: knock.F("some_property"),
						Operator: knock.F(knock.UserSetPreferencesParamsCategoriesObjectConditionsOperatorEqualTo),
						Variable: knock.F("recipient.property"),
					}}),
				},
			}),
			ChannelTypes: knock.F(knock.UserSetPreferencesParamsChannelTypes{
				Chat:      knock.F[knock.UserSetPreferencesParamsChannelTypesChatUnion](shared.UnionBool(true)),
				Email:     knock.F[knock.UserSetPreferencesParamsChannelTypesEmailUnion](shared.UnionBool(true)),
				HTTP:      knock.F[knock.UserSetPreferencesParamsChannelTypesHTTPUnion](shared.UnionBool(true)),
				InAppFeed: knock.F[knock.UserSetPreferencesParamsChannelTypesInAppFeedUnion](shared.UnionBool(true)),
				Push:      knock.F[knock.UserSetPreferencesParamsChannelTypesPushUnion](shared.UnionBool(true)),
				SMS:       knock.F[knock.UserSetPreferencesParamsChannelTypesSMSUnion](shared.UnionBool(true)),
			}),
			Workflows: knock.F(map[string]knock.UserSetPreferencesParamsWorkflowsUnion{
				"dinosaurs-loose": knock.UserSetPreferencesParamsWorkflowsObject{
					ChannelTypes: knock.F(knock.UserSetPreferencesParamsWorkflowsObjectChannelTypes{
						Chat:      knock.F[knock.UserSetPreferencesParamsWorkflowsObjectChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailUnion](shared.UnionBool(false)),
						HTTP:      knock.F[knock.UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.UserSetPreferencesParamsWorkflowsObjectChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Conditions: knock.F([]knock.UserSetPreferencesParamsWorkflowsObjectCondition{{
						Argument: knock.F("some_property"),
						Operator: knock.F(knock.UserSetPreferencesParamsWorkflowsObjectConditionsOperatorEqualTo),
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

func TestUserUnsetChannelData(t *testing.T) {
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
	_, err := client.Users.UnsetChannelData(
		context.TODO(),
		"user_id",
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

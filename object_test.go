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
			Recipients: knock.F([]knock.ObjectDeleteSubscriptionsParamsRecipientUnion{shared.UnionString("string")}),
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
						Tokens: knock.F([]string{"string"}),
					}),
				},
			}),
			Preferences: knock.F(map[string]knock.ObjectSetParamsPreferences{
				"default": {
					Categories: knock.F(map[string]knock.ObjectSetParamsPreferencesCategoriesUnion{
						"transactional": shared.UnionBool(true),
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
						"dinosaurs-loose": shared.UnionBool(true),
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
				"marketing":     shared.UnionBool(false),
				"transactional": shared.UnionBool(true),
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
				"dinosaurs-loose": shared.UnionBool(true),
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

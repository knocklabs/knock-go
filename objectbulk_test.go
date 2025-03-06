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
		option.WithToken("My Token"),
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
		option.WithToken("My Token"),
	)
	_, err := client.Objects.Bulk.AddSubscriptions(
		context.TODO(),
		"projects",
		knock.ObjectBulkAddSubscriptionsParams{
			Subscriptions: knock.F([]knock.ObjectBulkAddSubscriptionsParamsSubscription{{
				ID:         knock.F("project-1"),
				Recipients: knock.F([]knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion{shared.UnionString("string")}),
				Properties: knock.F[any](nil),
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
		option.WithToken("My Token"),
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
							Tokens: knock.F([]string{"string"}),
						}),
					},
				}),
				CreatedAt: knock.F(time.Now()),
				Preferences: knock.F(map[string]knock.ObjectBulkSetParamsObjectsPreferences{
					"default": {
						Categories: knock.F(map[string]knock.ObjectBulkSetParamsObjectsPreferencesCategoriesUnion{
							"transactional": shared.UnionBool(true),
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
							"dinosaurs-loose": shared.UnionBool(true),
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

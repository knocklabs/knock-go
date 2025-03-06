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
		option.WithToken("My Token"),
	)
	_, err := client.Users.Bulk.Delete(context.TODO(), knock.UserBulkDeleteParams{
		UserIDs1: knock.F([]string{"string"}),
		UserIDs2: knock.F([]string{"user_1", "user_2"}),
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
		option.WithToken("My Token"),
	)
	_, err := client.Users.Bulk.Identify(context.TODO(), knock.UserBulkIdentifyParams{
		Users: knock.F([]knock.UserBulkIdentifyParamsUser{{
			ID: knock.F("user_1"),
			ChannelData: knock.F(map[string]knock.UserBulkIdentifyParamsUsersChannelData{
				"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
					Data: knock.F[knock.UserBulkIdentifyParamsUsersChannelDataDataUnion](knock.UserBulkIdentifyParamsUsersChannelDataDataPushChannelData{
						Tokens: knock.F([]string{"string"}),
					}),
				},
			}),
			CreatedAt: knock.F(time.Now()),
			Preferences: knock.F(map[string]knock.UserBulkIdentifyParamsUsersPreferences{
				"default": {
					Categories: knock.F(map[string]knock.UserBulkIdentifyParamsUsersPreferencesCategoriesUnion{
						"transactional": shared.UnionBool(true),
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
						"dinosaurs-loose": shared.UnionBool(true),
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
		option.WithToken("My Token"),
	)
	_, err := client.Users.Bulk.SetPreferences(context.TODO(), knock.UserBulkSetPreferencesParams{
		Preferences: knock.F(knock.UserBulkSetPreferencesParamsPreferences{
			Categories: knock.F(map[string]knock.UserBulkSetPreferencesParamsPreferencesCategoriesUnion{
				"marketing":     shared.UnionBool(false),
				"transactional": shared.UnionBool(true),
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
				"dinosaurs-loose": shared.UnionBool(true),
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

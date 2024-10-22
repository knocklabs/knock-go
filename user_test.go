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
					Data: knock.F(knock.UserUpdateParamsChannelDataData{
						Tokens: knock.F([]string{"push_token_xxx"}),
					}),
				},
			}),
			CreatedAt: knock.F(time.Now()),
			Preferences: knock.F(map[string]knock.UserUpdateParamsPreferences{
				"default": {
					Categories: knock.F[knock.UserUpdateParamsPreferencesCategoriesUnion](knock.UserUpdateParamsPreferencesCategoriesObject{
						ChannelTypes: knock.F(knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypes{
							Chat:      knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion](shared.UnionBool(true)),
							Email:     knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion](shared.UnionBool(true)),
							HTTP:      knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
							InAppFeed: knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
							Push:      knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion](shared.UnionBool(true)),
							SMS:       knock.F[knock.UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion](shared.UnionBool(true)),
						}),
						Conditions: knock.F([]knock.UserUpdateParamsPreferencesCategoriesObjectCondition{{
							Argument: knock.F("some_property"),
							Operator: knock.F(knock.UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEqualTo),
							Variable: knock.F("recipient.property"),
						}, {
							Argument: knock.F("some_property"),
							Operator: knock.F(knock.UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEqualTo),
							Variable: knock.F("recipient.property"),
						}, {
							Argument: knock.F("some_property"),
							Operator: knock.F(knock.UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEqualTo),
							Variable: knock.F("recipient.property"),
						}}),
					}),
					ChannelTypes: knock.F(knock.UserUpdateParamsPreferencesChannelTypes{
						Chat:      knock.F[knock.UserUpdateParamsPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
						Email:     knock.F[knock.UserUpdateParamsPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
						HTTP:      knock.F[knock.UserUpdateParamsPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.UserUpdateParamsPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
						Push:      knock.F[knock.UserUpdateParamsPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
						SMS:       knock.F[knock.UserUpdateParamsPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
					}),
					Workflows: knock.F[knock.UserUpdateParamsPreferencesWorkflowsUnion](shared.UnionBool(false)),
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
			FromUserID: knock.F("from_user_id"),
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

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
			ObjectIDs: knock.F([]string{"obj_123", "obj_456", "obj_789"}),
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
				Recipients: knock.F([]knock.RecipientRequestUnionParam{knock.InlineIdentifyUserRequestParam{
					ID: knock.F("user_1"),
					ChannelData: knock.F(knock.InlineChannelDataRequestParam{knock.InlineChannelDataRequestItemParam{
						ChannelID: knock.F("97c5837d-c65c-4d54-aa39-080eeb81c69d"),
						Data: knock.F[knock.InlineChannelDataRequestItemDataUnionParam](knock.PushChannelDataParam{
							Tokens:   knock.F([]string{"push_token_xxx"}),
							Type:     knock.F(knock.PushChannelDataTypePushFcm),
							Typename: knock.F(knock.PushChannelData_TypenamePushChannelData),
						}),
						Provider: knock.F("push_fcm"),
					}}),
					CreatedAt:   knock.F(time.Now()),
					Preferences: knock.F[any](map[string]interface{}{}),
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
			Objects: knock.F([]knock.InlineObjectRequestParam{{
				ID:         knock.F("project_1"),
				Collection: knock.F("projects"),
				ChannelData: knock.F(knock.InlineChannelDataRequestParam{knock.InlineChannelDataRequestItemParam{
					ChannelID: knock.F("97c5837d-c65c-4d54-aa39-080eeb81c69d"),
					Data: knock.F[knock.InlineChannelDataRequestItemDataUnionParam](knock.PushChannelDataParam{
						Tokens:   knock.F([]string{"push_token_xxx"}),
						Type:     knock.F(knock.PushChannelDataTypePushFcm),
						Typename: knock.F(knock.PushChannelData_TypenamePushChannelData),
					}),
					Provider: knock.F("push_fcm"),
				}}),
				CreatedAt:   knock.F(time.Now()),
				Preferences: knock.F[any](map[string]interface{}{}),
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

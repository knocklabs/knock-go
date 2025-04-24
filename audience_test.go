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

func TestAudienceAddMembers(t *testing.T) {
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
	_, err := client.Audiences.AddMembers(
		context.TODO(),
		"key",
		knock.AudienceAddMembersParams{
			Members: knock.F([]knock.AudienceAddMembersParamsMember{{
				User: knock.F(knock.InlineIdentifyUserRequestParam{
					ID: knock.F("dr_sattler"),
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
				}),
				Tenant: knock.F("ingen_isla_nublar"),
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

func TestAudienceListMembers(t *testing.T) {
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
	_, err := client.Audiences.ListMembers(context.TODO(), "key")
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAudienceRemoveMembers(t *testing.T) {
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
	_, err := client.Audiences.RemoveMembers(
		context.TODO(),
		"key",
		knock.AudienceRemoveMembersParams{
			Members: knock.F([]knock.AudienceRemoveMembersParamsMember{{
				User: knock.F(knock.InlineIdentifyUserRequestParam{
					ID: knock.F("dr_sattler"),
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
				}),
				Tenant: knock.F("ingen_isla_nublar"),
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

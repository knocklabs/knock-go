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
)

func TestUserGuideGetChannelWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Guides.GetChannel(
		context.TODO(),
		"user_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		knock.UserGuideGetChannelParams{
			Data:   knock.F("data"),
			Tenant: knock.F("tenant"),
			Type:   knock.F("type"),
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

func TestUserGuideMarkMessageAsArchivedWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Guides.MarkMessageAsArchived(
		context.TODO(),
		"user_id",
		"message_id",
		knock.UserGuideMarkMessageAsArchivedParams{
			ChannelID:    knock.F("123e4567-e89b-12d3-a456-426614174000"),
			GuideID:      knock.F("323e4567-e89b-12d3-a456-426614174000"),
			GuideKey:     knock.F("guide_12345"),
			GuideStepRef: knock.F("step_12345"),
			Content: knock.F[any](map[string]interface{}{
				"body":  "Guide content body",
				"title": "Guide Title",
			}),
			Data: knock.F[any](map[string]interface{}{
				"product_id": "product_123",
			}),
			IsFinal: knock.F(true),
			Metadata: knock.F(map[string]interface{}{
				"source": "bar",
			}),
			Tenant: knock.F("tenant_12345"),
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

func TestUserGuideMarkMessageAsInteractedWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Guides.MarkMessageAsInteracted(
		context.TODO(),
		"user_id",
		"message_id",
		knock.UserGuideMarkMessageAsInteractedParams{
			ChannelID:    knock.F("123e4567-e89b-12d3-a456-426614174000"),
			GuideID:      knock.F("323e4567-e89b-12d3-a456-426614174000"),
			GuideKey:     knock.F("guide_12345"),
			GuideStepRef: knock.F("step_12345"),
			Content: knock.F[any](map[string]interface{}{
				"body":  "Guide content body",
				"title": "Guide Title",
			}),
			Data: knock.F[any](map[string]interface{}{
				"product_id": "product_123",
			}),
			IsFinal: knock.F(true),
			Metadata: knock.F(map[string]interface{}{
				"source": "bar",
			}),
			Tenant: knock.F("tenant_12345"),
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

func TestUserGuideMarkMessageAsSeenWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Guides.MarkMessageAsSeen(
		context.TODO(),
		"user_id",
		"message_id",
		knock.UserGuideMarkMessageAsSeenParams{
			ChannelID:    knock.F("123e4567-e89b-12d3-a456-426614174000"),
			GuideID:      knock.F("323e4567-e89b-12d3-a456-426614174000"),
			GuideKey:     knock.F("guide_12345"),
			GuideStepRef: knock.F("step_12345"),
			Content: knock.F[any](map[string]interface{}{
				"body":  "Guide content body",
				"title": "Guide Title",
			}),
			Data: knock.F[any](map[string]interface{}{
				"product_id": "product_123",
			}),
			IsFinal: knock.F(true),
			Metadata: knock.F(map[string]interface{}{
				"source": "bar",
			}),
			Tenant: knock.F("tenant_12345"),
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

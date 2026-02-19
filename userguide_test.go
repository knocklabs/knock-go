// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/knocklabs/knock-go"
	"github.com/knocklabs/knock-go/internal/testutil"
	"github.com/knocklabs/knock-go/option"
)

func TestUserGuideGetChannelWithOptionalParams(t *testing.T) {
	t.Skip("Mock server doesn't support callbacks yet")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
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
	t.Skip("Mock server doesn't support callbacks yet")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Users.Guides.MarkMessageAsArchived(
		context.TODO(),
		"user_id",
		"message_id",
		knock.UserGuideMarkMessageAsArchivedParams{
			ChannelID:    knock.F("123e4567-e89b-12d3-a456-426614174000"),
			GuideID:      knock.F("7e9dc78c-b3b1-4127-a54e-71f1899b831a"),
			GuideKey:     knock.F("tour_notification"),
			GuideStepRef: knock.F("lab_tours"),
			Content: knock.F(map[string]interface{}{
				"body":  "bar",
				"title": "bar",
			}),
			Data: knock.F(map[string]interface{}{
				"next_time":  "bar",
				"spots_left": "bar",
				"tour_id":    "bar",
			}),
			IsFinal: knock.F(false),
			Metadata: knock.F(map[string]interface{}{
				"cta":   "bar",
				"theme": "bar",
				"type":  "bar",
			}),
			Tenant: knock.F("ingen_isla_nublar"),
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
	t.Skip("Mock server doesn't support callbacks yet")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Users.Guides.MarkMessageAsInteracted(
		context.TODO(),
		"user_id",
		"message_id",
		knock.UserGuideMarkMessageAsInteractedParams{
			ChannelID:    knock.F("123e4567-e89b-12d3-a456-426614174000"),
			GuideID:      knock.F("7e9dc78c-b3b1-4127-a54e-71f1899b831a"),
			GuideKey:     knock.F("tour_notification"),
			GuideStepRef: knock.F("lab_tours"),
			Content: knock.F(map[string]interface{}{
				"body":  "bar",
				"title": "bar",
			}),
			Data: knock.F(map[string]interface{}{
				"next_time":  "bar",
				"spots_left": "bar",
				"tour_id":    "bar",
			}),
			IsFinal: knock.F(false),
			Metadata: knock.F(map[string]interface{}{
				"cta":   "bar",
				"theme": "bar",
				"type":  "bar",
			}),
			Tenant: knock.F("ingen_isla_nublar"),
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
	t.Skip("Mock server doesn't support callbacks yet")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Users.Guides.MarkMessageAsSeen(
		context.TODO(),
		"user_id",
		"message_id",
		knock.UserGuideMarkMessageAsSeenParams{
			ChannelID:    knock.F("123e4567-e89b-12d3-a456-426614174000"),
			GuideID:      knock.F("7e9dc78c-b3b1-4127-a54e-71f1899b831a"),
			GuideKey:     knock.F("tour_notification"),
			GuideStepRef: knock.F("lab_tours"),
			Content: knock.F(map[string]interface{}{
				"body":  "bar",
				"title": "bar",
			}),
			Data: knock.F(map[string]interface{}{
				"next_time":  "bar",
				"spots_left": "bar",
				"tour_id":    "bar",
			}),
			IsFinal: knock.F(false),
			Metadata: knock.F(map[string]interface{}{
				"cta":   "bar",
				"theme": "bar",
				"type":  "bar",
			}),
			Tenant: knock.F("ingen_isla_nublar"),
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

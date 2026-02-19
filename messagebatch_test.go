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

func TestMessageBatchArchive(t *testing.T) {
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
	_, err := client.Messages.Batch.Archive(context.TODO(), knock.MessageBatchArchiveParams{
		MessageIDs: knock.F([]string{"2w3YUpTTOxuDvZFji8OMsKrG176", "2w3YVRbPXMIh8Zq6oBFcVDA5xes"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchGetContent(t *testing.T) {
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
	_, err := client.Messages.Batch.GetContent(context.TODO(), knock.MessageBatchGetContentParams{
		MessageIDs: knock.F([]string{"string"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchMarkAsInteractedWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Batch.MarkAsInteracted(context.TODO(), knock.MessageBatchMarkAsInteractedParams{
		MessageIDs: knock.F([]string{"1jNaXzB2RZX3LY8wVQnfCKyPnv7"}),
		Metadata: knock.F(map[string]interface{}{
			"key": "bar",
		}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchMarkAsRead(t *testing.T) {
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
	_, err := client.Messages.Batch.MarkAsRead(context.TODO(), knock.MessageBatchMarkAsReadParams{
		MessageIDs: knock.F([]string{"2w3YUpTTOxuDvZFji8OMsKrG176", "2w3YVRbPXMIh8Zq6oBFcVDA5xes"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchMarkAsSeen(t *testing.T) {
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
	_, err := client.Messages.Batch.MarkAsSeen(context.TODO(), knock.MessageBatchMarkAsSeenParams{
		MessageIDs: knock.F([]string{"2w3YUpTTOxuDvZFji8OMsKrG176", "2w3YVRbPXMIh8Zq6oBFcVDA5xes"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchMarkAsUnread(t *testing.T) {
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
	_, err := client.Messages.Batch.MarkAsUnread(context.TODO(), knock.MessageBatchMarkAsUnreadParams{
		MessageIDs: knock.F([]string{"2w3YUpTTOxuDvZFji8OMsKrG176", "2w3YVRbPXMIh8Zq6oBFcVDA5xes"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchMarkAsUnseen(t *testing.T) {
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
	_, err := client.Messages.Batch.MarkAsUnseen(context.TODO(), knock.MessageBatchMarkAsUnseenParams{
		MessageIDs: knock.F([]string{"2w3YUpTTOxuDvZFji8OMsKrG176", "2w3YVRbPXMIh8Zq6oBFcVDA5xes"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchUnarchive(t *testing.T) {
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
	_, err := client.Messages.Batch.Unarchive(context.TODO(), knock.MessageBatchUnarchiveParams{
		MessageIDs: knock.F([]string{"2w3YUpTTOxuDvZFji8OMsKrG176", "2w3YVRbPXMIh8Zq6oBFcVDA5xes"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/knocklabs/knock-go"
	"github.com/knocklabs/knock-go/internal/testutil"
	"github.com/knocklabs/knock-go/option"
)

func TestChannelBulkUpdateMessageStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.Bulk.UpdateMessageStatus(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		knock.ChannelBulkUpdateMessageStatusParamsActionSeen,
		knock.ChannelBulkUpdateMessageStatusParams{
			Archived:         knock.F(knock.ChannelBulkUpdateMessageStatusParamsArchivedInclude),
			DeliveryStatus:   knock.F(knock.ChannelBulkUpdateMessageStatusParamsDeliveryStatusDelivered),
			EngagementStatus: knock.F(knock.ChannelBulkUpdateMessageStatusParamsEngagementStatusSeen),
			HasTenant:        knock.F(true),
			NewerThan:        knock.F(time.Now()),
			OlderThan:        knock.F(time.Now()),
			RecipientIDs:     knock.F([]string{"recipient1", "recipient2"}),
			Tenants:          knock.F([]string{"tenant1", "tenant2"}),
			TriggerData:      knock.F(`{"key":"value"}`),
			Workflows:        knock.F([]string{"workflow1", "workflow2"}),
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

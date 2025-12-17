// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock_test

import (
	"context"
	"os"
	"testing"

	"github.com/knocklabs/knock-go"
	"github.com/knocklabs/knock-go/internal/testutil"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/shared"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism doesn't support callbacks yet")
	response, err := client.Workflows.Trigger(
		context.TODO(),
		"dinosaurs-loose",
		knock.WorkflowTriggerParams{
			Recipients: knock.F([]knock.RecipientRequestUnionParam{shared.UnionString("dnedry")}),
			Data: knock.F(map[string]interface{}{
				"dinosaur": "triceratops",
			}),
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", response.WorkflowRunID)
}

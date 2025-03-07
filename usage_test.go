// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/knock-go"
	"github.com/stainless-sdks/knock-go/internal/testutil"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
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
		option.WithBearerToken("My Bearer Token"),
	)
	response, err := client.Workflows.Trigger(
		context.TODO(),
		"dinosaurs-loose",
		knock.WorkflowTriggerParams{
			Data: knock.F(map[string]interface{}{
				"dinosaur": "triceratops",
			}),
			Recipients: knock.F([]knock.WorkflowTriggerParamsRecipientUnion{shared.UnionString("dnedry")}),
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response.WorkflowRunID)
}

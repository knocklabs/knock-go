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
	"github.com/stainless-sdks/knock-go/shared"
)

func TestWorkflowCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Cancel(
		context.TODO(),
		"key",
		knock.WorkflowCancelParams{
			CancellationKey: knock.F("cancel-workflow-123"),
			Recipients:      knock.F([]string{"jhammond"}),
			Tenant:          knock.F("prk_1"),
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

func TestWorkflowTriggerWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Trigger(
		context.TODO(),
		"key",
		knock.WorkflowTriggerParams{
			Actor:           knock.F[knock.WorkflowTriggerParamsActorUnion](shared.UnionString("string")),
			CancellationKey: knock.Null[string](),
			Data: knock.F(map[string]interface{}{
				"dinosaur_names":  "bar",
				"is_alert":        "bar",
				"park_id":         "bar",
				"severity":        "bar",
				"welcome_message": "bar",
			}),
			Recipients: knock.F([]knock.WorkflowTriggerParamsRecipientUnion{shared.UnionString("jhammond")}),
			Tenant:     knock.F[knock.WorkflowTriggerParamsTenantUnion](shared.UnionString("acme_corp")),
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

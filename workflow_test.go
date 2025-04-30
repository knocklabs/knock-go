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
	"github.com/knocklabs/knock-go/shared"
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
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Workflows.Cancel(
		context.TODO(),
		"key",
		knock.WorkflowCancelParams{
			CancellationKey: knock.F("cancel-workflow-123"),
			Recipients:      knock.F([]knock.RecipientReferenceUnionParam{shared.UnionString("jhammond")}),
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
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Workflows.Trigger(
		context.TODO(),
		"key",
		knock.WorkflowTriggerParams{
			Recipients:      knock.F([]knock.RecipientRequestUnionParam{shared.UnionString("dr_grant"), shared.UnionString("dr_sattler"), shared.UnionString("dr_malcolm")}),
			Actor:           knock.F[knock.RecipientRequestUnionParam](shared.UnionString("mr_dna")),
			CancellationKey: knock.F("isla_nublar_incident_1993"),
			Data: knock.F(map[string]interface{}{
				"affected_areas":      "bar",
				"attraction_id":       "bar",
				"evacuation_protocol": "bar",
				"message":             "bar",
				"severity":            "bar",
				"system_status":       "bar",
			}),
			Tenant: knock.F[knock.InlineTenantRequestUnionParam](shared.UnionString("ingen_isla_nublar")),
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

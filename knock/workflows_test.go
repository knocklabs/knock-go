package knock

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestWorkflows_Trigger(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"workflow_run_id":"e2898d04-cb0c-5a1b-93e0-6c3f6bad82ef"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))

	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	request := &TriggerWorkflowRequest{
		Workflow: "test",
		Data: map[string]interface{}{
			"life":      "uh, finds a way",
			"dinosaurs": "loose",
		},
	}
	request.AddRecipientByID("tim")
	request.AddRecipientByID("lex")
	request.AddRecipientByEntity(map[string]interface{}{
		"id":         "projects-2",
		"collection": "projects",
	})

	request.AddActorByEntity(map[string]interface{}{
		"id":         "projects-1",
		"collection": "projects",
	})

	fmt.Printf("%v\n", request)

	workflowRunId, err := client.Workflows.Trigger(ctx, request)

	want := "e2898d04-cb0c-5a1b-93e0-6c3f6bad82ef"

	c.Assert(err, qt.IsNil)
	c.Assert(workflowRunId, qt.DeepEquals, want)
}

func TestWorkflows_Cancel(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	request := &CancelWorkflowRequest{
		Workflow:        "test",
		CancellationKey: "cancellation-key-123",
	}

	request.AddRecipientByID("tom")
	request.AddRecipientByEntity(map[string]interface{}{
		"id":         "projects-2",
		"collection": "projects",
	})

	err = client.Workflows.Cancel(ctx, request)

	c.Assert(err, qt.IsNil)
}

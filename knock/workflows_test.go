package knock

import (
	"context"
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

	workflowRunId, err := client.Workflows.Trigger(ctx, request, nil)

	want := "e2898d04-cb0c-5a1b-93e0-6c3f6bad82ef"

	c.Assert(err, qt.IsNil)
	c.Assert(workflowRunId, qt.DeepEquals, want)
}

func TestWorkflows_Trigger_Idempotence(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"workflow_run_id":"e2898d04-cb0c-5a1b-93e0-6c3f6bad82ef"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)

		// Verify the idempotency key was set
		c.Assert(r.Header.Get("Idempotency-Key"), qt.Not(qt.Equals), "")
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

	options := &MethodOptions{
		IdempotencyKey: "idempotency-key-123",
	}
	workflowRunId, err := client.Workflows.Trigger(ctx, request, options)

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

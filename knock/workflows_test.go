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

	workflow, err := client.Workflows.Trigger(ctx, &TriggerWorkflowRequest{
		Workflow:   "test",
		Recipients: []string{"tim", "lex"},
		Data: map[string]interface{}{
			"life":      "uh, finds a way",
			"dinosaurs": "loose",
		},
	})

	want := &TriggerWorkflowResponse{
		WorkflowRunId: "e2898d04-cb0c-5a1b-93e0-6c3f6bad82ef",
	}

	c.Assert(err, qt.IsNil)
	c.Assert(workflow, qt.DeepEquals, want)
}

func TestWorkflows_Cancel(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	err = client.Workflows.Cancel(ctx, &CancelWorkflowRequest{
		Workflow:        "test",
		CancellationKey: "user-123",
	})

	c.Assert(err, qt.IsNil)
}

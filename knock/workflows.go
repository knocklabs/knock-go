package knock

// TriggerWorkflowRequest encapsulates the request for triggering a workflow.
type TriggerWorkflowRequest struct {
	Workflow  string
	Recipient string
	Actor     string `json:"actor,omitempty"`
}

type TriggerWorkflowResponse struct {
	WorkflowRunId string `json:"workflow_run_id"`
}

package knock

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// WorkflowsService is an interface for communicating with the Knock
// Workflows API endpoints.
type WorkflowsService interface {
	Trigger(context.Context, *TriggerWorkflowRequest) (*TriggerWorkflowResponse, error)
	Cancel(context.Context, *CancelWorkflowRequest) error
}

type workflowsService struct {
	client *Client
}

var _ WorkflowsService = &workflowsService{}

func NewWorkflowsService(client *Client) *workflowsService {
	return &workflowsService{
		client: client,
	}
}

// TriggerWorkflowRequest encapsulates the request for triggering a workflow.
type TriggerWorkflowRequest struct {
	Workflow        string                 `json:"workflow"`
	Recipients      []string               `json:"recipients"`
	Actor           string                 `json:"actor,omitempty"`
	CancellationKey string                 `json:"cancellation_key,omitempty"`
	Tenant          string                 `json:"tenant,omitempty"`
	Data            map[string]interface{} `json:"data,omitempty"`
}

type TriggerWorkflowResponse struct {
	WorkflowRunId string `json:"workflow_run_id"`
}

// CancelWorkflowRequest encapsulates the request for cancelling a workflow.
type CancelWorkflowRequest struct {
	Workflow        string   `json:"workflow"`
	Recipients      []string `json:"recipients,omitempty"`
	CancellationKey string   `json:"cancellation_key"`
}

func workflowsAPIPath(workflowId string) string {
	return fmt.Sprintf("v1/workflows/%s", workflowId)
}

func (ws *workflowsService) Trigger(ctx context.Context, triggerReq *TriggerWorkflowRequest) (*TriggerWorkflowResponse, error) {
	path := fmt.Sprintf("%s/trigger", workflowsAPIPath(triggerReq.Workflow))
	req, err := ws.client.newRequest(http.MethodPost, path, triggerReq)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for trigger workflow")
	}
	workflowResponse := &TriggerWorkflowResponse{}
	_, err = ws.client.do(ctx, req, workflowResponse)

	if err != nil {
		return nil, errors.Wrap(err, "error making request for trigger workflow")
	}

	return workflowResponse, nil
}

func (ws *workflowsService) Cancel(ctx context.Context, cancelReq *CancelWorkflowRequest) error {
	path := fmt.Sprintf("%s/cancel", workflowsAPIPath(cancelReq.Workflow))
	req, err := ws.client.newRequest(http.MethodPost, path, cancelReq)
	if err != nil {
		return errors.Wrap(err, "error creating request to cancel workflow")
	}

	_, err = ws.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request to cancel workflow")
	}

	return err
}

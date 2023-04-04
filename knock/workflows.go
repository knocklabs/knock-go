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
	Trigger(context.Context, *TriggerWorkflowRequest, *MethodOptions) (string, error)
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

// Client structs
type TriggerWorkflowRequest struct {
	Workflow        string                 `json:"workflow"`
	Recipients      []interface{}          `json:"recipients"`
	Actor           interface{}            `json:"actor,omitempty"`
	CancellationKey string                 `json:"cancellation_key,omitempty"`
	Tenant          string                 `json:"tenant,omitempty"`
	Data            map[string]interface{} `json:"data,omitempty"`
}

type TriggerWorkflowResponse struct {
	WorkflowRunId string `json:"workflow_run_id"`
}

type CancelWorkflowRequest struct {
	Workflow        string        `json:"workflow"`
	Recipients      []interface{} `json:"recipients"`
	CancellationKey string        `json:"cancellation_key"`
}

func workflowsAPIPath(workflowId string) string {
	return fmt.Sprintf("v1/workflows/%s", workflowId)
}

func (tr *TriggerWorkflowRequest) AddRecipientByID(recipientID string) TriggerWorkflowRequest {
	tr.Recipients = append(tr.Recipients, recipientID)
	return *tr
}

func (tr *TriggerWorkflowRequest) AddRecipientByEntity(entity map[string]interface{}) TriggerWorkflowRequest {
	tr.Recipients = append(tr.Recipients, entity)
	return *tr
}

func (tr *TriggerWorkflowRequest) AddActorByID(actorID string) TriggerWorkflowRequest {
	tr.Actor = actorID
	return *tr
}

func (tr *TriggerWorkflowRequest) AddActorByEntity(entity map[string]interface{}) TriggerWorkflowRequest {
	tr.Actor = entity
	return *tr
}

func (ws *workflowsService) Trigger(ctx context.Context, triggerReq *TriggerWorkflowRequest, methodOptions *MethodOptions) (string, error) {
	path := fmt.Sprintf("%s/trigger", workflowsAPIPath(triggerReq.Workflow))
	req, err := ws.client.newRequest(http.MethodPost, path, triggerReq, methodOptions)

	if err != nil {
		return "", errors.Wrap(err, "error creating request for trigger workflow")
	}
	workflowResponse := &TriggerWorkflowResponse{}
	_, err = ws.client.do(ctx, req, workflowResponse)

	if err != nil {
		return "", errors.Wrap(err, "error making request for trigger workflow")
	}

	return workflowResponse.WorkflowRunId, nil
}

func (tr *CancelWorkflowRequest) AddRecipientByID(recipientID string) CancelWorkflowRequest {
	tr.Recipients = append(tr.Recipients, recipientID)
	return *tr
}

func (tr *CancelWorkflowRequest) AddRecipientByEntity(entity map[string]interface{}) CancelWorkflowRequest {
	tr.Recipients = append(tr.Recipients, entity)
	return *tr
}

func (ws *workflowsService) Cancel(ctx context.Context, cancelReq *CancelWorkflowRequest) error {
	path := fmt.Sprintf("%s/cancel", workflowsAPIPath(cancelReq.Workflow))
	req, err := ws.client.newRequest(http.MethodPost, path, cancelReq, nil)
	if err != nil {
		return errors.Wrap(err, "error creating request to cancel workflow")
	}

	_, err = ws.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request to cancel workflow")
	}

	return err
}

package knock

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// WorkflowsService is an interface for communicating with the Knock
// Workflows API endpoints.
type WorkflowsService interface {
	Trigger(context.Context, *TriggerWorkflowRequest, *MethodOptions) (string, error)
	Cancel(context.Context, *CancelWorkflowRequest) error
	CreateSchedules(context.Context, *CreateSchedulesRequest) ([]*Schedule, error)
	UpdateSchedules(context.Context, *UpdateSchedulesRequest) ([]*Schedule, error)
	DeleteSchedules(context.Context, *DeleteSchedulesRequest) ([]*Schedule, error)
	ListSchedules(context.Context, *ListSchedulesRequest) ([]*Schedule, *PageInfo, error)
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

type ScheduleRepeat struct {
	Frequency  string   `json:"frequency"`
	Interval   int      `json:"interval,omitempty"`
	Days       []string `json:"days"`
	DayOfMonth int      `json:"day_of_month,omitempty"`
	Hours      int      `json:"hours,omitempty"`
	Minutes    int      `json:"minutes,omitempty"`
}

type Schedule struct {
	Cursor           string                 `json:"__cursor"`
	ID               string                 `json:"id"`
	Recipient        interface{}            `json:"recipient"`
	Repeats          []*ScheduleRepeat      `json:"repeats"`
	Workflow         string                 `json:"workflow"`
	Tenant           string                 `json:"tenant"`
	NextOccurrenceAt time.Time              `json:"next_occurrence_at"`
	InsertedAt       time.Time              `json:"inserted_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
	Data             map[string]interface{} `json:"data"`
}

type CreateSchedulesRequest struct {
	Workflow    string                 `json:"workflow"`
	Recipients  []interface{}          `json:"recipients"`
	Repeats     []*ScheduleRepeat      `json:"repeats"`
	Actor       interface{}            `json:"actor,omitempty"`
	Tenant      string                 `json:"tenant,omitempty"`
	ScheduledAt time.Time              `json:"scheduled_at,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}

type UpdateSchedulesRequest struct {
	ScheduleIDs []string               `json:"schedule_ids"`
	Repeats     []*ScheduleRepeat      `json:"repeats"`
	Actor       interface{}            `json:"actor,omitempty"`
	Tenant      string                 `json:"tenant,omitempty"`
	ScheduledAt time.Time              `json:"scheduled_at,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}

type DeleteSchedulesRequest struct {
	ScheduleIDs []string `json:"schedule_ids"`
}

type ListSchedulesRequest struct {
	PageSize   int           `url:"page_size,omitempty"`
	Cursor     string        `url:"cursor,omitempty"`
	Before     string        `url:"before,omitempty"`
	After      string        `url:"after,omitempty"`
	Workflow   string        `url:"workflow"`
	Tenant     string        `url:"tenant,omitempty"`
	Recipients []interface{} `url:"recipients,omitempty"`
}

type ListSchedulesResponse struct {
	Entries  []*Schedule `json:"entries"`
	PageInfo *PageInfo   `json:"page_info"`
}

func workflowsAPIPath(workflowId string) string {
	return fmt.Sprintf("v1/workflows/%s", workflowId)
}

func schedulesAPIPath() string {
	return "v1/schedules"
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

func (csr *CreateSchedulesRequest) AddRecipientByID(recipientID string) CreateSchedulesRequest {
	csr.Recipients = append(csr.Recipients, recipientID)
	return *csr
}

func (csr *CreateSchedulesRequest) AddRecipientByEntity(entity map[string]interface{}) CreateSchedulesRequest {
	csr.Recipients = append(csr.Recipients, entity)
	return *csr
}

func (csr *CreateSchedulesRequest) AddActorByID(actorID string) CreateSchedulesRequest {
	csr.Actor = actorID
	return *csr
}

func (csr *CreateSchedulesRequest) AddActorByEntity(entity map[string]interface{}) CreateSchedulesRequest {
	csr.Actor = entity
	return *csr
}

func (csr *CreateSchedulesRequest) AddRepeat(repeat *ScheduleRepeat) CreateSchedulesRequest {
	csr.Repeats = append(csr.Repeats, repeat)
	return *csr
}

func (ws *workflowsService) CreateSchedules(ctx context.Context, createSchedulesReq *CreateSchedulesRequest) ([]*Schedule, error) {
	path := schedulesAPIPath()
	req, err := ws.client.newRequest(http.MethodPost, path, createSchedulesReq, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for creating schedules")
	}
	createSchedulesResponse := []*Schedule{}
	_, err = ws.client.do(ctx, req, createSchedulesResponse)

	if err != nil {
		return nil, errors.Wrap(err, "error making request for creating schedules")
	}

	return createSchedulesResponse, nil
}

func (usr *UpdateSchedulesRequest) AddActorByID(actorID string) UpdateSchedulesRequest {
	usr.Actor = actorID
	return *usr
}

func (usr *UpdateSchedulesRequest) AddActorByEntity(entity map[string]interface{}) UpdateSchedulesRequest {
	usr.Actor = entity
	return *usr
}

func (usr *UpdateSchedulesRequest) AddRepeat(repeat *ScheduleRepeat) UpdateSchedulesRequest {
	usr.Repeats = append(usr.Repeats, repeat)
	return *usr
}

func (ws *workflowsService) UpdateSchedules(ctx context.Context, updateSchedulesReq *UpdateSchedulesRequest) ([]*Schedule, error) {
	path := schedulesAPIPath()
	req, err := ws.client.newRequest(http.MethodPut, path, updateSchedulesReq, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for updating schedules")
	}
	updateSchedulesResponse := []*Schedule{}
	_, err = ws.client.do(ctx, req, updateSchedulesResponse)

	if err != nil {
		return nil, errors.Wrap(err, "error making request for updating schedules")
	}

	return updateSchedulesResponse, nil
}

func (ws *workflowsService) DeleteSchedules(ctx context.Context, deleteSchedulesReq *DeleteSchedulesRequest) ([]*Schedule, error) {
	path := schedulesAPIPath()
	req, err := ws.client.newRequest(http.MethodDelete, path, deleteSchedulesReq, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for deleting schedules")
	}
	deleteSchedulesResponse := []*Schedule{}
	_, err = ws.client.do(ctx, req, deleteSchedulesResponse)

	if err != nil {
		return nil, errors.Wrap(err, "error making request for deleting schedules")
	}

	return deleteSchedulesResponse, nil
}

func (ws *workflowsService) ListSchedules(ctx context.Context, listReq *ListSchedulesRequest) ([]*Schedule, *PageInfo, error) {
	queryString, err := query.Values(listReq)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error parsing query params to list schedules")
	}

	listRes := &ListSchedulesResponse{}

	path := fmt.Sprintf("%s?%s", schedulesAPIPath(), queryString.Encode())

	req, err := ws.client.newRequest(http.MethodGet, path, listReq, nil)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request to list schedules")
	}

	_, err = ws.client.do(ctx, req, listRes)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request to list messages")
	}

	return listRes.Entries, listRes.PageInfo, nil
}

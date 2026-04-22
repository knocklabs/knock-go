// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/apiquery"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/packages/pagination"
)

// A workflow run represents an individual execution of a workflow for a specific
// recipient.
//
// WorkflowRecipientRunService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowRecipientRunService] method instead.
type WorkflowRecipientRunService struct {
	Options []option.RequestOption
}

// NewWorkflowRecipientRunService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowRecipientRunService(opts ...option.RequestOption) (r *WorkflowRecipientRunService) {
	r = &WorkflowRecipientRunService{}
	r.Options = opts
	return
}

// Returns a paginated list of workflow recipient runs for the current environment.
func (r *WorkflowRecipientRunService) List(ctx context.Context, query WorkflowRecipientRunListParams, opts ...option.RequestOption) (res *pagination.ItemsCursor[WorkflowRecipientRun], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/workflow_recipient_runs"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a paginated list of workflow recipient runs for the current environment.
func (r *WorkflowRecipientRunService) ListAutoPaging(ctx context.Context, query WorkflowRecipientRunListParams, opts ...option.RequestOption) *pagination.ItemsCursorAutoPager[WorkflowRecipientRun] {
	return pagination.NewItemsCursorAutoPager(r.List(ctx, query, opts...))
}

// Returns a single workflow recipient run with its associated events.
func (r *WorkflowRecipientRunService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkflowRecipientRunDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workflow_recipient_runs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A workflow recipient run represents an individual execution of a workflow for a
// specific recipient.
type WorkflowRecipientRun struct {
	// The unique identifier for the workflow recipient run (per-recipient).
	ID string `json:"id" api:"required" format:"uuid"`
	// The typename of the schema.
	Typename string `json:"__typename" api:"required"`
	// Timestamp when the resource was created.
	InsertedAt time.Time `json:"inserted_at" api:"required" format:"date-time"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (ID, collection).
	Recipient RecipientReferenceUnion `json:"recipient" api:"required"`
	// The current status of the workflow recipient run. One of `queued`, `processing`,
	// `paused`, `completed`, or `cancelled`.
	Status WorkflowRecipientRunStatus `json:"status" api:"required"`
	// Describes how the workflow was triggered.
	TriggerSource WorkflowRecipientRunTriggerSource `json:"trigger_source" api:"required"`
	// The timestamp when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The key of the workflow that was executed.
	Workflow string `json:"workflow" api:"required"`
	// The identifier for the top-level workflow run shared across all recipients in a
	// single trigger.
	WorkflowRunID string `json:"workflow_run_id" api:"required" format:"uuid"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (ID, collection).
	Actor RecipientReferenceUnion `json:"actor" api:"nullable"`
	// The number of errors encountered during the workflow recipient run.
	ErrorCount int64 `json:"error_count"`
	// The tenant associated with the workflow recipient run.
	Tenant string                   `json:"tenant" api:"nullable"`
	JSON   workflowRecipientRunJSON `json:"-"`
}

// workflowRecipientRunJSON contains the JSON metadata for the struct
// [WorkflowRecipientRun]
type workflowRecipientRunJSON struct {
	ID            apijson.Field
	Typename      apijson.Field
	InsertedAt    apijson.Field
	Recipient     apijson.Field
	Status        apijson.Field
	TriggerSource apijson.Field
	UpdatedAt     apijson.Field
	Workflow      apijson.Field
	WorkflowRunID apijson.Field
	Actor         apijson.Field
	ErrorCount    apijson.Field
	Tenant        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkflowRecipientRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowRecipientRunJSON) RawJSON() string {
	return r.raw
}

// The current status of the workflow recipient run. One of `queued`, `processing`,
// `paused`, `completed`, or `cancelled`.
type WorkflowRecipientRunStatus string

const (
	WorkflowRecipientRunStatusQueued     WorkflowRecipientRunStatus = "queued"
	WorkflowRecipientRunStatusProcessing WorkflowRecipientRunStatus = "processing"
	WorkflowRecipientRunStatusPaused     WorkflowRecipientRunStatus = "paused"
	WorkflowRecipientRunStatusCompleted  WorkflowRecipientRunStatus = "completed"
	WorkflowRecipientRunStatusCancelled  WorkflowRecipientRunStatus = "cancelled"
)

func (r WorkflowRecipientRunStatus) IsKnown() bool {
	switch r {
	case WorkflowRecipientRunStatusQueued, WorkflowRecipientRunStatusProcessing, WorkflowRecipientRunStatusPaused, WorkflowRecipientRunStatusCompleted, WorkflowRecipientRunStatusCancelled:
		return true
	}
	return false
}

// Describes how the workflow was triggered.
type WorkflowRecipientRunTriggerSource struct {
	// The type of trigger source. One of `api`, `audience`, `schedule`, `broadcast`,
	// `workflow_step`, `integration`, or `rehearsal`.
	Type WorkflowRecipientRunTriggerSourceType `json:"type" api:"required"`
	// The key of the audience that triggered the workflow.
	AudienceKey string `json:"audience_key" api:"nullable"`
	// The cancellation key provided when the workflow was triggered via the API.
	CancellationKey string `json:"cancellation_key" api:"nullable"`
	// The ID of the schedule that triggered the workflow.
	ScheduleID string                                `json:"schedule_id" api:"nullable"`
	JSON       workflowRecipientRunTriggerSourceJSON `json:"-"`
}

// workflowRecipientRunTriggerSourceJSON contains the JSON metadata for the struct
// [WorkflowRecipientRunTriggerSource]
type workflowRecipientRunTriggerSourceJSON struct {
	Type            apijson.Field
	AudienceKey     apijson.Field
	CancellationKey apijson.Field
	ScheduleID      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkflowRecipientRunTriggerSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowRecipientRunTriggerSourceJSON) RawJSON() string {
	return r.raw
}

// The type of trigger source. One of `api`, `audience`, `schedule`, `broadcast`,
// `workflow_step`, `integration`, or `rehearsal`.
type WorkflowRecipientRunTriggerSourceType string

const (
	WorkflowRecipientRunTriggerSourceTypeAPI          WorkflowRecipientRunTriggerSourceType = "api"
	WorkflowRecipientRunTriggerSourceTypeAudience     WorkflowRecipientRunTriggerSourceType = "audience"
	WorkflowRecipientRunTriggerSourceTypeSchedule     WorkflowRecipientRunTriggerSourceType = "schedule"
	WorkflowRecipientRunTriggerSourceTypeBroadcast    WorkflowRecipientRunTriggerSourceType = "broadcast"
	WorkflowRecipientRunTriggerSourceTypeWorkflowStep WorkflowRecipientRunTriggerSourceType = "workflow_step"
	WorkflowRecipientRunTriggerSourceTypeIntegration  WorkflowRecipientRunTriggerSourceType = "integration"
	WorkflowRecipientRunTriggerSourceTypeRehearsal    WorkflowRecipientRunTriggerSourceType = "rehearsal"
)

func (r WorkflowRecipientRunTriggerSourceType) IsKnown() bool {
	switch r {
	case WorkflowRecipientRunTriggerSourceTypeAPI, WorkflowRecipientRunTriggerSourceTypeAudience, WorkflowRecipientRunTriggerSourceTypeSchedule, WorkflowRecipientRunTriggerSourceTypeBroadcast, WorkflowRecipientRunTriggerSourceTypeWorkflowStep, WorkflowRecipientRunTriggerSourceTypeIntegration, WorkflowRecipientRunTriggerSourceTypeRehearsal:
		return true
	}
	return false
}

// A single workflow recipient run with its events.
type WorkflowRecipientRunDetail struct {
	// A list of events that occurred during the workflow recipient run.
	Events []WorkflowRecipientRunEvent    `json:"events" api:"required"`
	JSON   workflowRecipientRunDetailJSON `json:"-"`
	WorkflowRecipientRun
}

// workflowRecipientRunDetailJSON contains the JSON metadata for the struct
// [WorkflowRecipientRunDetail]
type workflowRecipientRunDetailJSON struct {
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowRecipientRunDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowRecipientRunDetailJSON) RawJSON() string {
	return r.raw
}

// An event that occurred during a workflow recipient run.
type WorkflowRecipientRunEvent struct {
	// The unique identifier for the event.
	ID string `json:"id" api:"required"`
	// The typename of the schema.
	Typename string `json:"__typename" api:"required"`
	// The type of event that occurred.
	Event string `json:"event" api:"required"`
	// Timestamp when the resource was created.
	InsertedAt time.Time `json:"inserted_at" api:"required" format:"date-time"`
	// Whether the event represents a successful or error state.
	Status WorkflowRecipientRunEventStatus `json:"status" api:"required"`
	// The attempt number of the workflow recipient run event. Increments for each
	// retry.
	Attempt int64 `json:"attempt"`
	// Event-specific data associated with the event.
	Data map[string]interface{} `json:"data" api:"nullable"`
	// The reference of the workflow step associated with this event.
	StepRef string `json:"step_ref" api:"nullable"`
	// The type of workflow step associated with this event.
	StepType string                        `json:"step_type" api:"nullable"`
	JSON     workflowRecipientRunEventJSON `json:"-"`
}

// workflowRecipientRunEventJSON contains the JSON metadata for the struct
// [WorkflowRecipientRunEvent]
type workflowRecipientRunEventJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Event       apijson.Field
	InsertedAt  apijson.Field
	Status      apijson.Field
	Attempt     apijson.Field
	Data        apijson.Field
	StepRef     apijson.Field
	StepType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowRecipientRunEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowRecipientRunEventJSON) RawJSON() string {
	return r.raw
}

// Whether the event represents a successful or error state.
type WorkflowRecipientRunEventStatus string

const (
	WorkflowRecipientRunEventStatusOk    WorkflowRecipientRunEventStatus = "ok"
	WorkflowRecipientRunEventStatusError WorkflowRecipientRunEventStatus = "error"
)

func (r WorkflowRecipientRunEventStatus) IsKnown() bool {
	switch r {
	case WorkflowRecipientRunEventStatusOk, WorkflowRecipientRunEventStatusError:
		return true
	}
	return false
}

type WorkflowRecipientRunListParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// Limits the results to workflow recipient runs started before the given date.
	EndingAt param.Field[time.Time] `query:"ending_at" format:"date-time"`
	// Limits the results to workflow recipient runs that have errors.
	HasErrors param.Field[bool] `query:"has_errors"`
	// The number of items per page (defaults to 50).
	PageSize param.Field[int64] `query:"page_size"`
	// Limits the results to workflow recipient runs for the given recipient. Accepts a
	// user ID string or an object reference with `id` and `collection`.
	Recipient param.Field[RecipientReferenceUnionParam] `query:"recipient"`
	// Limits the results to workflow recipient runs started after the given date.
	StartingAt param.Field[time.Time] `query:"starting_at" format:"date-time"`
	// Limits the results to workflow recipient runs with the given status.
	Status param.Field[[]WorkflowRecipientRunListParamsStatus] `query:"status"`
	// Limits the results to workflow recipient runs for the given tenant.
	Tenant param.Field[string] `query:"tenant"`
	// Limits the results to workflow recipient runs for the given workflow key.
	Workflow param.Field[string] `query:"workflow"`
}

// URLQuery serializes [WorkflowRecipientRunListParams]'s query parameters as
// `url.Values`.
func (r WorkflowRecipientRunListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowRecipientRunListParamsStatus string

const (
	WorkflowRecipientRunListParamsStatusQueued     WorkflowRecipientRunListParamsStatus = "queued"
	WorkflowRecipientRunListParamsStatusProcessing WorkflowRecipientRunListParamsStatus = "processing"
	WorkflowRecipientRunListParamsStatusPaused     WorkflowRecipientRunListParamsStatus = "paused"
	WorkflowRecipientRunListParamsStatusCompleted  WorkflowRecipientRunListParamsStatus = "completed"
	WorkflowRecipientRunListParamsStatusCancelled  WorkflowRecipientRunListParamsStatus = "cancelled"
)

func (r WorkflowRecipientRunListParamsStatus) IsKnown() bool {
	switch r {
	case WorkflowRecipientRunListParamsStatusQueued, WorkflowRecipientRunListParamsStatusProcessing, WorkflowRecipientRunListParamsStatusPaused, WorkflowRecipientRunListParamsStatusCompleted, WorkflowRecipientRunListParamsStatusCancelled:
		return true
	}
	return false
}

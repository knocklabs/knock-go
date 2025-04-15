// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
)

// WorkflowService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options []option.RequestOption
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r *WorkflowService) {
	r = &WorkflowService{}
	r.Options = opts
	return
}

// Issues a cancellation request to inflight workflow runs
func (r *WorkflowService) Cancel(ctx context.Context, key string, body WorkflowCancelParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/workflows/%s/cancel", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Triggers a workflow
func (r *WorkflowService) Trigger(ctx context.Context, key string, body WorkflowTriggerParams, opts ...option.RequestOption) (res *WorkflowTriggerResponse, err error) {
	opts = append(r.Options[:], opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/workflows/%s/trigger", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The response from triggering a workflow
type WorkflowTriggerResponse struct {
	// The ID of the workflow trigger. This value allows you to track individual
	// workflow runs associated with this trigger request.
	WorkflowRunID string                      `json:"workflow_run_id,required" format:"uuid"`
	JSON          workflowTriggerResponseJSON `json:"-"`
}

// workflowTriggerResponseJSON contains the JSON metadata for the struct
// [WorkflowTriggerResponse]
type workflowTriggerResponseJSON struct {
	WorkflowRunID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkflowTriggerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowTriggerResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowCancelParams struct {
	// The cancellation key supplied to the workflow trigger endpoint to use for
	// cancelling one or more workflow runs.
	CancellationKey param.Field[string] `json:"cancellation_key,required"`
	// An optional list of recipients to cancel the workflow for using the cancellation
	// key.
	Recipients param.Field[[]string] `json:"recipients"`
	Tenant     param.Field[string]   `json:"tenant"`
}

func (r WorkflowCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkflowTriggerParams struct {
	// Specifies a recipient in a request. This can either be a user identifier
	// (string), an inline user request (object), or an inline object request, which is
	// determined by the presence of a `collection` property.
	Actor param.Field[RecipientRequestUnionParam] `json:"actor"`
	// An optional key that is used in the workflow cancellation endpoint to target a
	// cancellation of any workflow runs associated with this trigger.
	CancellationKey param.Field[string] `json:"cancellation_key"`
	// An optional map of data to be used in the workflow. This data will be available
	// to the workflow as a map in the `data` field.
	Data param.Field[map[string]string] `json:"data"`
	// The recipients to trigger the workflow for.
	Recipients param.Field[[]RecipientRequestUnionParam] `json:"recipients"`
	// An inline tenant request
	Tenant param.Field[InlineTenantRequestUnionParam] `json:"tenant"`
}

func (r WorkflowTriggerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

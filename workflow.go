// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
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

// When invoked for a workflow using a specific workflow key and cancellation key,
// will cancel any queued workflow runs associated with that key/cancellation key
// pair. Can optionally be provided one or more recipients to scope the request to.
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

// Trigger a workflow (specified by the key) to run for the given recipients, using
// the parameters provided. Returns an identifier for the workflow run request. All
// workflow runs are executed asynchronously. This endpoint also handles
// [inline identifications](/managing-recipients/identifying-recipients#inline-identifying-recipients)
// for the `actor`, `recipient`, and `tenant` fields.
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

// The response from triggering a workflow.
type WorkflowTriggerResponse struct {
	// This value allows you to track individual messages associated with this trigger
	// request.
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
	// An optional key that is used to reference a specific workflow trigger request
	// when issuing a [workflow cancellation](/send-notifications/canceling-workflows)
	// request. Must be provided while triggering a workflow in order to enable
	// subsequent cancellation. Should be unique across trigger requests to avoid
	// unintentional cancellations.
	CancellationKey param.Field[string] `json:"cancellation_key,required"`
	// A list of recipients to cancel the notification for. If omitted, cancels for all
	// recipients associated with the cancellation key.
	Recipients param.Field[[]RecipientReferenceUnionParam] `json:"recipients"`
}

func (r WorkflowCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkflowTriggerParams struct {
	// The recipients to trigger the workflow for. Can inline identify users, objects,
	// or use a list of user IDs. Limited to 1,000 recipients.
	Recipients param.Field[[]RecipientRequestUnionParam] `json:"recipients,required"`
	// Specifies a recipient in a request. This can either be a user identifier
	// (string), an inline user request (object), or an inline object request, which is
	// determined by the presence of a `collection` property.
	Actor param.Field[RecipientRequestUnionParam] `json:"actor"`
	// An optional key that is used to reference a specific workflow trigger request
	// when issuing a [workflow cancellation](/send-notifications/canceling-workflows)
	// request. Must be provided while triggering a workflow in order to enable
	// subsequent cancellation. Should be unique across trigger requests to avoid
	// unintentional cancellations.
	CancellationKey param.Field[string] `json:"cancellation_key"`
	// An optional map of data to pass into the workflow execution. There is a 10MB
	// limit on the size of the full `data` payload. Any individual string value
	// greater than 1024 bytes in length will be
	// [truncated](/developer-tools/api-logs#log-truncation) in your logs.
	Data param.Field[map[string]interface{}] `json:"data"`
	// An request to set a tenant inline.
	Tenant param.Field[InlineTenantRequestUnionParam] `json:"tenant"`
}

func (r WorkflowTriggerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
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

// Trigger a workflow specified by the key to run for the given recipients, using
// the parameters provided. Returns an identifier for the workflow run request. All
// workflow runs are executed asynchronously.
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
	Actor param.Field[WorkflowTriggerParamsActorUnion] `json:"actor"`
	// An optional key that is used in the workflow cancellation endpoint to target a
	// cancellation of any workflow runs associated with this trigger.
	CancellationKey param.Field[string] `json:"cancellation_key"`
	// An optional map of data to be used in the workflow. This data will be available
	// to the workflow as a map in the `data` field.
	Data param.Field[map[string]interface{}] `json:"data"`
	// The recipients to trigger the workflow for. Cannot exceed 1000 recipients in a
	// single trigger.
	Recipients param.Field[[]WorkflowTriggerParamsRecipientUnion] `json:"recipients"`
	// An inline tenant request
	Tenant param.Field[WorkflowTriggerParamsTenantUnion] `json:"tenant"`
}

func (r WorkflowTriggerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type WorkflowTriggerParamsActor struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID          param.Field[string]      `json:"id,required"`
	ChannelData param.Field[interface{}] `json:"channel_data"`
	Collection  param.Field[string]      `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt   param.Field[time.Time]   `json:"created_at" format:"date-time"`
	Preferences param.Field[interface{}] `json:"preferences"`
}

func (r WorkflowTriggerParamsActor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkflowTriggerParamsActor) ImplementsWorkflowTriggerParamsActorUnion() {}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// Satisfied by [shared.UnionString],
// [WorkflowTriggerParamsActorInlineIdentifyUserRequest],
// [WorkflowTriggerParamsActorInlineIdentifyObjectRequest],
// [WorkflowTriggerParamsActor].
type WorkflowTriggerParamsActorUnion interface {
	ImplementsWorkflowTriggerParamsActorUnion()
}

// A set of parameters to inline-identify a user with. Inline identifying the user
// will ensure that the user is available before the request is executed in Knock.
// It will perform an upsert against the user you're supplying, replacing any
// properties specified.
type WorkflowTriggerParamsActorInlineIdentifyUserRequest struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]shared.ChannelDataRequestParam] `json:"channel_data"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]shared.PreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                                   `json:"-,extras"`
}

func (r WorkflowTriggerParamsActorInlineIdentifyUserRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkflowTriggerParamsActorInlineIdentifyUserRequest) ImplementsWorkflowTriggerParamsActorUnion() {
}

// Inline identifies a custom object belonging to a collection
type WorkflowTriggerParamsActorInlineIdentifyObjectRequest struct {
	ID         param.Field[string] `json:"id,required"`
	Collection param.Field[string] `json:"collection,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]shared.ChannelDataRequestParam] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                                 `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]shared.PreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                                   `json:"-,extras"`
}

func (r WorkflowTriggerParamsActorInlineIdentifyObjectRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkflowTriggerParamsActorInlineIdentifyObjectRequest) ImplementsWorkflowTriggerParamsActorUnion() {
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type WorkflowTriggerParamsRecipient struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID          param.Field[string]      `json:"id,required"`
	ChannelData param.Field[interface{}] `json:"channel_data"`
	Collection  param.Field[string]      `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt   param.Field[time.Time]   `json:"created_at" format:"date-time"`
	Preferences param.Field[interface{}] `json:"preferences"`
}

func (r WorkflowTriggerParamsRecipient) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkflowTriggerParamsRecipient) ImplementsWorkflowTriggerParamsRecipientUnion() {}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// Satisfied by [shared.UnionString],
// [WorkflowTriggerParamsRecipientsInlineIdentifyUserRequest],
// [WorkflowTriggerParamsRecipientsInlineIdentifyObjectRequest],
// [WorkflowTriggerParamsRecipient].
type WorkflowTriggerParamsRecipientUnion interface {
	ImplementsWorkflowTriggerParamsRecipientUnion()
}

// A set of parameters to inline-identify a user with. Inline identifying the user
// will ensure that the user is available before the request is executed in Knock.
// It will perform an upsert against the user you're supplying, replacing any
// properties specified.
type WorkflowTriggerParamsRecipientsInlineIdentifyUserRequest struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]shared.ChannelDataRequestParam] `json:"channel_data"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]shared.PreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                                   `json:"-,extras"`
}

func (r WorkflowTriggerParamsRecipientsInlineIdentifyUserRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkflowTriggerParamsRecipientsInlineIdentifyUserRequest) ImplementsWorkflowTriggerParamsRecipientUnion() {
}

// Inline identifies a custom object belonging to a collection
type WorkflowTriggerParamsRecipientsInlineIdentifyObjectRequest struct {
	ID         param.Field[string] `json:"id,required"`
	Collection param.Field[string] `json:"collection,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]shared.ChannelDataRequestParam] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                                 `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]shared.PreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                                   `json:"-,extras"`
}

func (r WorkflowTriggerParamsRecipientsInlineIdentifyObjectRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkflowTriggerParamsRecipientsInlineIdentifyObjectRequest) ImplementsWorkflowTriggerParamsRecipientUnion() {
}

// An inline tenant request
//
// Satisfied by [shared.UnionString], [shared.TenantRequestParam].
type WorkflowTriggerParamsTenantUnion interface {
	ImplementsWorkflowTriggerParamsTenantUnion()
}

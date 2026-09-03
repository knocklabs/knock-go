// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/apiquery"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
)

// A message sent to a single recipient on a channel.
//
// MessageBatchService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageBatchService] method instead.
type MessageBatchService struct {
	Options []option.RequestOption
}

// NewMessageBatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessageBatchService(opts ...option.RequestOption) (r *MessageBatchService) {
	r = &MessageBatchService{}
	r.Options = opts
	return
}

// Marks the given messages as archived. Archived messages are hidden from the
// default message list in the feed but can still be accessed and unarchived later.
func (r *MessageBatchService) Archive(ctx context.Context, body MessageBatchArchiveParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batch/archived"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get the contents of multiple messages in a single request.
func (r *MessageBatchService) GetContent(ctx context.Context, query MessageBatchGetContentParams, opts ...option.RequestOption) (res *[]MessageContents, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batch/content"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Marks the given messages as interacted with by the user. This can include any
// user action on the message, with optional metadata about the specific
// interaction. Cannot include more than 5 key-value pairs, must not contain nested
// data. Read more about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsInteracted(ctx context.Context, body MessageBatchMarkAsInteractedParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batch/interacted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks the given messages as `read`. Read more about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsRead(ctx context.Context, body MessageBatchMarkAsReadParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks the given messages as `seen`. This indicates that the user has viewed the
// message in their feed or inbox. Read more about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsSeen(ctx context.Context, body MessageBatchMarkAsSeenParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batch/seen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks the given messages as `unread`. This reverses the `read` state. Read more
// about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsUnread(ctx context.Context, body MessageBatchMarkAsUnreadParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batch/unread"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks the given messages as `unseen`. This reverses the `seen` state. Read more
// about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsUnseen(ctx context.Context, body MessageBatchMarkAsUnseenParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batch/unseen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks the given messages as unarchived. This reverses the `archived` state.
// Archived messages are hidden from the default message list in the feed but can
// still be accessed and unarchived later.
func (r *MessageBatchService) Unarchive(ctx context.Context, body MessageBatchUnarchiveParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batch/unarchived"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request to update the status of multiple messages in batch.
type BatchMessagesStatusRequestParam struct {
	// The message IDs to update the status of.
	MessageIDs param.Field[[]string] `json:"message_ids" api:"required"`
}

func (r BatchMessagesStatusRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchArchiveParams struct {
	// Request to update the status of multiple messages in batch.
	BatchMessagesStatusRequest BatchMessagesStatusRequestParam `json:"batch_messages_status_request" api:"required"`
}

func (r MessageBatchArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BatchMessagesStatusRequest)
}

type MessageBatchGetContentParams struct {
	// The IDs of the messages to fetch contents of.
	MessageIDs param.Field[[]string] `query:"message_ids" api:"required"`
}

// URLQuery serializes [MessageBatchGetContentParams]'s query parameters as
// `url.Values`.
func (r MessageBatchGetContentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageBatchMarkAsInteractedParams struct {
	// The message IDs to batch mark as interacted with.
	MessageIDs param.Field[[]string] `json:"message_ids" api:"required"`
	// Metadata about the interaction.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r MessageBatchMarkAsInteractedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchMarkAsReadParams struct {
	// Request to update the status of multiple messages in batch.
	BatchMessagesStatusRequest BatchMessagesStatusRequestParam `json:"batch_messages_status_request" api:"required"`
}

func (r MessageBatchMarkAsReadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BatchMessagesStatusRequest)
}

type MessageBatchMarkAsSeenParams struct {
	// Request to update the status of multiple messages in batch.
	BatchMessagesStatusRequest BatchMessagesStatusRequestParam `json:"batch_messages_status_request" api:"required"`
}

func (r MessageBatchMarkAsSeenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BatchMessagesStatusRequest)
}

type MessageBatchMarkAsUnreadParams struct {
	// Request to update the status of multiple messages in batch.
	BatchMessagesStatusRequest BatchMessagesStatusRequestParam `json:"batch_messages_status_request" api:"required"`
}

func (r MessageBatchMarkAsUnreadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BatchMessagesStatusRequest)
}

type MessageBatchMarkAsUnseenParams struct {
	// Request to update the status of multiple messages in batch.
	BatchMessagesStatusRequest BatchMessagesStatusRequestParam `json:"batch_messages_status_request" api:"required"`
}

func (r MessageBatchMarkAsUnseenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BatchMessagesStatusRequest)
}

type MessageBatchUnarchiveParams struct {
	// Request to update the status of multiple messages in batch.
	BatchMessagesStatusRequest BatchMessagesStatusRequestParam `json:"batch_messages_status_request" api:"required"`
}

func (r MessageBatchUnarchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BatchMessagesStatusRequest)
}

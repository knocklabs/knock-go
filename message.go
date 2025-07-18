// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/apiquery"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/packages/pagination"
	"github.com/knocklabs/knock-go/shared"
	"github.com/tidwall/gjson"
)

// MessageService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
	Batch   *MessageBatchService
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r *MessageService) {
	r = &MessageService{}
	r.Options = opts
	r.Batch = NewMessageBatchService(opts...)
	return
}

// Returns a paginated list of messages for the current environment.
func (r *MessageService) List(ctx context.Context, query MessageListParams, opts ...option.RequestOption) (res *pagination.ItemsCursor[Message], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/messages"
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

// Returns a paginated list of messages for the current environment.
func (r *MessageService) ListAutoPaging(ctx context.Context, query MessageListParams, opts ...option.RequestOption) *pagination.ItemsCursorAutoPager[Message] {
	return pagination.NewItemsCursorAutoPager(r.List(ctx, query, opts...))
}

// Archives a message for the user. Archived messages are hidden from the default
// message list in the feed but can still be accessed and unarchived later.
func (r *MessageService) Archive(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/archived", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Retrieves a specific message by its ID.
func (r *MessageService) Get(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the fully rendered contents of a message, where the response depends on
// which channel the message was sent through.
func (r *MessageService) GetContent(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageGetContentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/content", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a paginated list of activities for the specified message.
func (r *MessageService) ListActivities(ctx context.Context, messageID string, query MessageListActivitiesParams, opts ...option.RequestOption) (res *pagination.ItemsCursor[Activity], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/activities", messageID)
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

// Returns a paginated list of activities for the specified message.
func (r *MessageService) ListActivitiesAutoPaging(ctx context.Context, messageID string, query MessageListActivitiesParams, opts ...option.RequestOption) *pagination.ItemsCursorAutoPager[Activity] {
	return pagination.NewItemsCursorAutoPager(r.ListActivities(ctx, messageID, query, opts...))
}

// Returns a paginated list of delivery logs for the specified message.
func (r *MessageService) ListDeliveryLogs(ctx context.Context, messageID string, query MessageListDeliveryLogsParams, opts ...option.RequestOption) (res *pagination.ItemsCursor[MessageDeliveryLog], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/delivery_logs", messageID)
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

// Returns a paginated list of delivery logs for the specified message.
func (r *MessageService) ListDeliveryLogsAutoPaging(ctx context.Context, messageID string, query MessageListDeliveryLogsParams, opts ...option.RequestOption) *pagination.ItemsCursorAutoPager[MessageDeliveryLog] {
	return pagination.NewItemsCursorAutoPager(r.ListDeliveryLogs(ctx, messageID, query, opts...))
}

// Returns a paginated list of events for the specified message.
func (r *MessageService) ListEvents(ctx context.Context, messageID string, query MessageListEventsParams, opts ...option.RequestOption) (res *pagination.ItemsCursor[MessageEvent], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/events", messageID)
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

// Returns a paginated list of events for the specified message.
func (r *MessageService) ListEventsAutoPaging(ctx context.Context, messageID string, query MessageListEventsParams, opts ...option.RequestOption) *pagination.ItemsCursorAutoPager[MessageEvent] {
	return pagination.NewItemsCursorAutoPager(r.ListEvents(ctx, messageID, query, opts...))
}

// Marks a message as `interacted` with by the user. This can include any user
// action on the message, with optional metadata about the specific interaction.
// Cannot include more than 5 key-value pairs, must not contain nested data. Read
// more about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageService) MarkAsInteracted(ctx context.Context, messageID string, body MessageMarkAsInteractedParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/interacted", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Marks a message as `read`. This indicates that the user has read the message
// content. Read more about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageService) MarkAsRead(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/read", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Marks a message as `seen`. This indicates that the user has viewed the message
// in their feed or inbox. Read more about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageService) MarkAsSeen(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/seen", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Marks a message as `unread`. This reverses the `read` state. Read more about
// message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageService) MarkAsUnread(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/read", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Marks a message as `unseen`. This reverses the `seen` state. Read more about
// message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageService) MarkAsUnseen(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/seen", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Removes a message from the archived state, making it visible in the default
// message list in the feed again.
func (r *MessageService) Unarchive(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/archived", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// An activity associated with a workflow trigger request. Messages produced after
// a [batch step](/designing-workflows/batch-function) can be associated with one
// or more activities. Non-batched messages will always be associated with a single
// activity.
type Activity struct {
	// Unique identifier for the activity.
	ID string `json:"id"`
	// The typename of the schema.
	Typename string `json:"__typename"`
	// A recipient of a notification, which is either a user or an object.
	Actor Recipient `json:"actor,nullable"`
	// The workflow trigger `data` payload associated with the activity.
	Data map[string]interface{} `json:"data,nullable"`
	// Timestamp when the activity was created.
	InsertedAt time.Time `json:"inserted_at" format:"date-time"`
	// A recipient of a notification, which is either a user or an object.
	Recipient Recipient `json:"recipient"`
	// Timestamp when the activity was last updated.
	UpdatedAt time.Time    `json:"updated_at" format:"date-time"`
	JSON      activityJSON `json:"-"`
}

// activityJSON contains the JSON metadata for the struct [Activity]
type activityJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Actor       apijson.Field
	Data        apijson.Field
	InsertedAt  apijson.Field
	Recipient   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Activity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activityJSON) RawJSON() string {
	return r.raw
}

// Represents a single message that was generated by a workflow for a given
// channel.
type Message struct {
	// The unique identifier for the message.
	ID string `json:"id"`
	// The typename of the schema.
	Typename string `json:"__typename"`
	// One or more actors that are associated with this message. Note: this is a list
	// that can contain up to 10 actors if the message is produced from a
	// [batch](/designing-workflows/batch-function).
	Actors []RecipientReferenceUnion `json:"actors"`
	// Timestamp when the message was archived.
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// The ID for the channel the message was sent through.
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when the message was clicked.
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Data associated with the message’s workflow run. Includes the workflow trigger
	// request’s `data` payload merged with any additional data returned by a
	// [fetch function](/designing-workflows/fetch-function). For messages produced
	// after a [batch step](/designing-workflows/batch-function), includes the payload
	// `data` from the most-recent trigger request (the final `activity` in the batch).
	Data map[string]interface{} `json:"data,nullable"`
	// A list of engagement statuses.
	EngagementStatuses []MessageEngagementStatus `json:"engagement_statuses"`
	// Timestamp when the resource was created.
	InsertedAt time.Time `json:"inserted_at" format:"date-time"`
	// Timestamp when the message was interacted with.
	InteractedAt time.Time `json:"interacted_at,nullable" format:"date-time"`
	// Timestamp when a link in the message was clicked.
	LinkClickedAt time.Time `json:"link_clicked_at,nullable" format:"date-time"`
	// The metadata associated with the message.
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Timestamp when the message was read.
	ReadAt time.Time `json:"read_at,nullable" format:"date-time"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (ID, collection).
	Recipient RecipientReferenceUnion `json:"recipient"`
	// Timestamp when the message was scheduled to be sent.
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when the message was seen.
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// The workflow that triggered the message.
	Source MessageSource `json:"source"`
	// The message delivery status.
	Status MessageStatus `json:"status"`
	// The ID of the `tenant` associated with the message. Only present when a `tenant`
	// is provided on a workflow trigger request.
	Tenant string `json:"tenant,nullable"`
	// The timestamp when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The key of the workflow that generated the message.
	//
	// Deprecated: deprecated
	Workflow string      `json:"workflow,nullable"`
	JSON     messageJSON `json:"-"`
}

// messageJSON contains the JSON metadata for the struct [Message]
type messageJSON struct {
	ID                 apijson.Field
	Typename           apijson.Field
	Actors             apijson.Field
	ArchivedAt         apijson.Field
	ChannelID          apijson.Field
	ClickedAt          apijson.Field
	Data               apijson.Field
	EngagementStatuses apijson.Field
	InsertedAt         apijson.Field
	InteractedAt       apijson.Field
	LinkClickedAt      apijson.Field
	Metadata           apijson.Field
	ReadAt             apijson.Field
	Recipient          apijson.Field
	ScheduledAt        apijson.Field
	SeenAt             apijson.Field
	Source             apijson.Field
	Status             apijson.Field
	Tenant             apijson.Field
	UpdatedAt          apijson.Field
	Workflow           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageJSON) RawJSON() string {
	return r.raw
}

// An engagement status for a message. Can be one of: read, seen, interacted,
// link_clicked, archived.
type MessageEngagementStatus string

const (
	MessageEngagementStatusSeen        MessageEngagementStatus = "seen"
	MessageEngagementStatusRead        MessageEngagementStatus = "read"
	MessageEngagementStatusInteracted  MessageEngagementStatus = "interacted"
	MessageEngagementStatusLinkClicked MessageEngagementStatus = "link_clicked"
	MessageEngagementStatusArchived    MessageEngagementStatus = "archived"
)

func (r MessageEngagementStatus) IsKnown() bool {
	switch r {
	case MessageEngagementStatusSeen, MessageEngagementStatusRead, MessageEngagementStatusInteracted, MessageEngagementStatusLinkClicked, MessageEngagementStatusArchived:
		return true
	}
	return false
}

// The workflow that triggered the message.
type MessageSource struct {
	Typename string `json:"__typename,required"`
	// The categories associated with the message.
	Categories []string `json:"categories,required"`
	// The key of the workflow that triggered the message.
	Key string `json:"key,required"`
	// The ID of the version of the workflow that triggered the message.
	VersionID string `json:"version_id,required" format:"uuid"`
	// The step reference for the step in the workflow that generated the message
	StepRef string            `json:"step_ref,nullable"`
	JSON    messageSourceJSON `json:"-"`
}

// messageSourceJSON contains the JSON metadata for the struct [MessageSource]
type messageSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	StepRef     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageSourceJSON) RawJSON() string {
	return r.raw
}

// The message delivery status.
type MessageStatus string

const (
	MessageStatusQueued            MessageStatus = "queued"
	MessageStatusSent              MessageStatus = "sent"
	MessageStatusDelivered         MessageStatus = "delivered"
	MessageStatusDeliveryAttempted MessageStatus = "delivery_attempted"
	MessageStatusUndelivered       MessageStatus = "undelivered"
	MessageStatusNotSent           MessageStatus = "not_sent"
	MessageStatusBounced           MessageStatus = "bounced"
)

func (r MessageStatus) IsKnown() bool {
	switch r {
	case MessageStatusQueued, MessageStatusSent, MessageStatusDelivered, MessageStatusDeliveryAttempted, MessageStatusUndelivered, MessageStatusNotSent, MessageStatusBounced:
		return true
	}
	return false
}

// A message delivery log contains a `request` from Knock to a downstream provider
// and the `response` that was returned.
type MessageDeliveryLog struct {
	// The unique identifier for the message delivery log.
	ID string `json:"id,required"`
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The ID of the environment in which the message delivery occurred.
	EnvironmentID string `json:"environment_id,required" format:"uuid"`
	// Timestamp when the message delivery log was created.
	InsertedAt string `json:"inserted_at,required"`
	// A message delivery log request.
	Request MessageDeliveryLogRequest `json:"request,required"`
	// A message delivery log response.
	Response MessageDeliveryLogResponse `json:"response,required"`
	// The name of the service that processed the delivery.
	ServiceName string                 `json:"service_name,required"`
	JSON        messageDeliveryLogJSON `json:"-"`
}

// messageDeliveryLogJSON contains the JSON metadata for the struct
// [MessageDeliveryLog]
type messageDeliveryLogJSON struct {
	ID            apijson.Field
	Typename      apijson.Field
	EnvironmentID apijson.Field
	InsertedAt    apijson.Field
	Request       apijson.Field
	Response      apijson.Field
	ServiceName   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MessageDeliveryLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageDeliveryLogJSON) RawJSON() string {
	return r.raw
}

// A message delivery log request.
type MessageDeliveryLogRequest struct {
	// The body content that was sent with the request.
	Body MessageDeliveryLogRequestBodyUnion `json:"body"`
	// The headers that were sent with the request.
	Headers map[string]interface{} `json:"headers,nullable"`
	// The host to which the request was sent.
	Host string `json:"host"`
	// The HTTP method used for the request.
	Method MessageDeliveryLogRequestMethod `json:"method"`
	// The path of the URL that was requested.
	Path string `json:"path"`
	// The query string of the URL that was requested.
	Query string                        `json:"query,nullable"`
	JSON  messageDeliveryLogRequestJSON `json:"-"`
}

// messageDeliveryLogRequestJSON contains the JSON metadata for the struct
// [MessageDeliveryLogRequest]
type messageDeliveryLogRequestJSON struct {
	Body        apijson.Field
	Headers     apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageDeliveryLogRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageDeliveryLogRequestJSON) RawJSON() string {
	return r.raw
}

// The body content that was sent with the request.
//
// Union satisfied by [shared.UnionString] or [MessageDeliveryLogRequestBodyMap].
type MessageDeliveryLogRequestBodyUnion interface {
	ImplementsMessageDeliveryLogRequestBodyUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageDeliveryLogRequestBodyUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageDeliveryLogRequestBodyMap{}),
		},
	)
}

type MessageDeliveryLogRequestBodyMap map[string]interface{}

func (r MessageDeliveryLogRequestBodyMap) ImplementsMessageDeliveryLogRequestBodyUnion() {}

// The HTTP method used for the request.
type MessageDeliveryLogRequestMethod string

const (
	MessageDeliveryLogRequestMethodGet    MessageDeliveryLogRequestMethod = "GET"
	MessageDeliveryLogRequestMethodPost   MessageDeliveryLogRequestMethod = "POST"
	MessageDeliveryLogRequestMethodPut    MessageDeliveryLogRequestMethod = "PUT"
	MessageDeliveryLogRequestMethodDelete MessageDeliveryLogRequestMethod = "DELETE"
	MessageDeliveryLogRequestMethodPatch  MessageDeliveryLogRequestMethod = "PATCH"
)

func (r MessageDeliveryLogRequestMethod) IsKnown() bool {
	switch r {
	case MessageDeliveryLogRequestMethodGet, MessageDeliveryLogRequestMethodPost, MessageDeliveryLogRequestMethodPut, MessageDeliveryLogRequestMethodDelete, MessageDeliveryLogRequestMethodPatch:
		return true
	}
	return false
}

// A message delivery log response.
type MessageDeliveryLogResponse struct {
	// The body content that was received with the response.
	Body MessageDeliveryLogResponseBodyUnion `json:"body"`
	// The headers that were received with the response.
	Headers map[string]interface{} `json:"headers,nullable"`
	// The HTTP status code of the response.
	Status int64                          `json:"status"`
	JSON   messageDeliveryLogResponseJSON `json:"-"`
}

// messageDeliveryLogResponseJSON contains the JSON metadata for the struct
// [MessageDeliveryLogResponse]
type messageDeliveryLogResponseJSON struct {
	Body        apijson.Field
	Headers     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageDeliveryLogResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageDeliveryLogResponseJSON) RawJSON() string {
	return r.raw
}

// The body content that was received with the response.
//
// Union satisfied by [shared.UnionString] or [MessageDeliveryLogResponseBodyMap].
type MessageDeliveryLogResponseBodyUnion interface {
	ImplementsMessageDeliveryLogResponseBodyUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageDeliveryLogResponseBodyUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageDeliveryLogResponseBodyMap{}),
		},
	)
}

type MessageDeliveryLogResponseBodyMap map[string]interface{}

func (r MessageDeliveryLogResponseBodyMap) ImplementsMessageDeliveryLogResponseBodyUnion() {}

// A message event. Occurs when a message
// [delivery or engagement status](/send-notifications/message-statuses) changes.
type MessageEvent struct {
	// The unique identifier for the message event.
	ID string `json:"id,required"`
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// Timestamp when the event was created.
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (ID, collection).
	Recipient RecipientReferenceUnion `json:"recipient,required"`
	// The type of event that occurred.
	Type MessageEventType `json:"type,required"`
	// The data associated with the message event. Only present for some event types.
	Data map[string]interface{} `json:"data,nullable"`
	JSON messageEventJSON       `json:"-"`
}

// messageEventJSON contains the JSON metadata for the struct [MessageEvent]
type messageEventJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	InsertedAt  apijson.Field
	Recipient   apijson.Field
	Type        apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageEventJSON) RawJSON() string {
	return r.raw
}

// The type of event that occurred.
type MessageEventType string

const (
	MessageEventTypeMessageArchived          MessageEventType = "message.archived"
	MessageEventTypeMessageBounced           MessageEventType = "message.bounced"
	MessageEventTypeMessageDelivered         MessageEventType = "message.delivered"
	MessageEventTypeMessageDeliveryAttempted MessageEventType = "message.delivery_attempted"
	MessageEventTypeMessageInteracted        MessageEventType = "message.interacted"
	MessageEventTypeMessageLinkClicked       MessageEventType = "message.link_clicked"
	MessageEventTypeMessageNotSent           MessageEventType = "message.not_sent"
	MessageEventTypeMessageQueued            MessageEventType = "message.queued"
	MessageEventTypeMessageRead              MessageEventType = "message.read"
	MessageEventTypeMessageSeen              MessageEventType = "message.seen"
	MessageEventTypeMessageSent              MessageEventType = "message.sent"
	MessageEventTypeMessageUnarchived        MessageEventType = "message.unarchived"
	MessageEventTypeMessageUndelivered       MessageEventType = "message.undelivered"
	MessageEventTypeMessageUnread            MessageEventType = "message.unread"
	MessageEventTypeMessageUnseen            MessageEventType = "message.unseen"
)

func (r MessageEventType) IsKnown() bool {
	switch r {
	case MessageEventTypeMessageArchived, MessageEventTypeMessageBounced, MessageEventTypeMessageDelivered, MessageEventTypeMessageDeliveryAttempted, MessageEventTypeMessageInteracted, MessageEventTypeMessageLinkClicked, MessageEventTypeMessageNotSent, MessageEventTypeMessageQueued, MessageEventTypeMessageRead, MessageEventTypeMessageSeen, MessageEventTypeMessageSent, MessageEventTypeMessageUnarchived, MessageEventTypeMessageUndelivered, MessageEventTypeMessageUnread, MessageEventTypeMessageUnseen:
		return true
	}
	return false
}

// The content of a message.
type MessageGetContentResponse struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// Content data specific to the channel type.
	Data MessageGetContentResponseData `json:"data,required"`
	// Timestamp when the message content was created.
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// The unique identifier for the message content.
	MessageID string                        `json:"message_id,required"`
	JSON      messageGetContentResponseJSON `json:"-"`
}

// messageGetContentResponseJSON contains the JSON metadata for the struct
// [MessageGetContentResponse]
type messageGetContentResponseJSON struct {
	Typename    apijson.Field
	Data        apijson.Field
	InsertedAt  apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseJSON) RawJSON() string {
	return r.raw
}

// Content data specific to the channel type.
type MessageGetContentResponseData struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The device token to send the push notification to.
	Token string `json:"token"`
	// The BCC email addresses.
	Bcc string `json:"bcc,nullable"`
	// This field can have the runtime type of
	// [[]MessageGetContentResponseDataMessageInAppFeedContentBlock].
	Blocks interface{} `json:"blocks"`
	// The content body of the SMS message.
	Body string `json:"body"`
	// The CC email addresses.
	Cc string `json:"cc,nullable"`
	// This field can have the runtime type of [map[string]interface{}].
	Connection interface{} `json:"connection"`
	// This field can have the runtime type of [map[string]interface{}].
	Data interface{} `json:"data"`
	// The sender's email address.
	From string `json:"from"`
	// The HTML body of the email message.
	HTMLBody string `json:"html_body"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{} `json:"metadata"`
	// The reply-to email address.
	ReplyTo string `json:"reply_to,nullable"`
	// The subject line of the email message.
	SubjectLine string `json:"subject_line"`
	// This field can have the runtime type of
	// [MessageGetContentResponseDataMessageChatContentTemplate].
	Template interface{} `json:"template"`
	// The text body of the email message.
	TextBody string `json:"text_body"`
	// The title of the push notification.
	Title string `json:"title"`
	// The recipient's email address.
	To    string                            `json:"to"`
	JSON  messageGetContentResponseDataJSON `json:"-"`
	union MessageGetContentResponseDataUnion
}

// messageGetContentResponseDataJSON contains the JSON metadata for the struct
// [MessageGetContentResponseData]
type messageGetContentResponseDataJSON struct {
	Typename    apijson.Field
	Token       apijson.Field
	Bcc         apijson.Field
	Blocks      apijson.Field
	Body        apijson.Field
	Cc          apijson.Field
	Connection  apijson.Field
	Data        apijson.Field
	From        apijson.Field
	HTMLBody    apijson.Field
	Metadata    apijson.Field
	ReplyTo     apijson.Field
	SubjectLine apijson.Field
	Template    apijson.Field
	TextBody    apijson.Field
	Title       apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageGetContentResponseDataJSON) RawJSON() string {
	return r.raw
}

func (r *MessageGetContentResponseData) UnmarshalJSON(data []byte) (err error) {
	*r = MessageGetContentResponseData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageGetContentResponseDataUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageGetContentResponseDataMessageEmailContent],
// [MessageGetContentResponseDataMessageSMSContent],
// [MessageGetContentResponseDataMessagePushContent],
// [MessageGetContentResponseDataMessageChatContent],
// [MessageGetContentResponseDataMessageInAppFeedContent].
func (r MessageGetContentResponseData) AsUnion() MessageGetContentResponseDataUnion {
	return r.union
}

// Content data specific to the channel type.
//
// Union satisfied by [MessageGetContentResponseDataMessageEmailContent],
// [MessageGetContentResponseDataMessageSMSContent],
// [MessageGetContentResponseDataMessagePushContent],
// [MessageGetContentResponseDataMessageChatContent] or
// [MessageGetContentResponseDataMessageInAppFeedContent].
type MessageGetContentResponseDataUnion interface {
	implementsMessageGetContentResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageGetContentResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetContentResponseDataMessageEmailContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetContentResponseDataMessageSMSContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetContentResponseDataMessagePushContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetContentResponseDataMessageChatContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetContentResponseDataMessageInAppFeedContent{}),
		},
	)
}

// The content of an email message.
type MessageGetContentResponseDataMessageEmailContent struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The sender's email address.
	From string `json:"from,required"`
	// The HTML body of the email message.
	HTMLBody string `json:"html_body,required"`
	// The subject line of the email message.
	SubjectLine string `json:"subject_line,required"`
	// The text body of the email message.
	TextBody string `json:"text_body,required"`
	// The recipient's email address.
	To string `json:"to,required"`
	// The BCC email addresses.
	Bcc string `json:"bcc,nullable"`
	// The CC email addresses.
	Cc string `json:"cc,nullable"`
	// The reply-to email address.
	ReplyTo string                                               `json:"reply_to,nullable"`
	JSON    messageGetContentResponseDataMessageEmailContentJSON `json:"-"`
}

// messageGetContentResponseDataMessageEmailContentJSON contains the JSON metadata
// for the struct [MessageGetContentResponseDataMessageEmailContent]
type messageGetContentResponseDataMessageEmailContentJSON struct {
	Typename    apijson.Field
	From        apijson.Field
	HTMLBody    apijson.Field
	SubjectLine apijson.Field
	TextBody    apijson.Field
	To          apijson.Field
	Bcc         apijson.Field
	Cc          apijson.Field
	ReplyTo     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageEmailContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageEmailContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetContentResponseDataMessageEmailContent) implementsMessageGetContentResponseData() {}

// The content of an SMS message.
type MessageGetContentResponseDataMessageSMSContent struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The content body of the SMS message.
	Body string `json:"body,required"`
	// The phone number the SMS was sent to.
	To   string                                             `json:"to,required"`
	JSON messageGetContentResponseDataMessageSMSContentJSON `json:"-"`
}

// messageGetContentResponseDataMessageSMSContentJSON contains the JSON metadata
// for the struct [MessageGetContentResponseDataMessageSMSContent]
type messageGetContentResponseDataMessageSMSContentJSON struct {
	Typename    apijson.Field
	Body        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageSMSContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageSMSContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetContentResponseDataMessageSMSContent) implementsMessageGetContentResponseData() {}

// Push channel data.
type MessageGetContentResponseDataMessagePushContent struct {
	// The device token to send the push notification to.
	Token string `json:"token,required"`
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The content body of the push notification.
	Body string `json:"body,required"`
	// The title of the push notification.
	Title string `json:"title,required"`
	// Additional data payload for the push notification.
	Data map[string]interface{}                              `json:"data,nullable"`
	JSON messageGetContentResponseDataMessagePushContentJSON `json:"-"`
}

// messageGetContentResponseDataMessagePushContentJSON contains the JSON metadata
// for the struct [MessageGetContentResponseDataMessagePushContent]
type messageGetContentResponseDataMessagePushContentJSON struct {
	Token       apijson.Field
	Typename    apijson.Field
	Body        apijson.Field
	Title       apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessagePushContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessagePushContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetContentResponseDataMessagePushContent) implementsMessageGetContentResponseData() {}

// The content of a chat message.
type MessageGetContentResponseDataMessageChatContent struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The channel data connection from the recipient to the underlying provider.
	Connection map[string]interface{} `json:"connection,required"`
	// The template structure for the chat message.
	Template MessageGetContentResponseDataMessageChatContentTemplate `json:"template,required"`
	// Additional metadata associated with the chat message.
	Metadata map[string]interface{}                              `json:"metadata,nullable"`
	JSON     messageGetContentResponseDataMessageChatContentJSON `json:"-"`
}

// messageGetContentResponseDataMessageChatContentJSON contains the JSON metadata
// for the struct [MessageGetContentResponseDataMessageChatContent]
type messageGetContentResponseDataMessageChatContentJSON struct {
	Typename    apijson.Field
	Connection  apijson.Field
	Template    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageChatContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageChatContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetContentResponseDataMessageChatContent) implementsMessageGetContentResponseData() {}

// The template structure for the chat message.
type MessageGetContentResponseDataMessageChatContentTemplate struct {
	// The blocks of the message in a chat.
	Blocks []MessageGetContentResponseDataMessageChatContentTemplateBlock `json:"blocks,nullable"`
	// The JSON content of the message.
	JsonContent map[string]interface{} `json:"json_content,nullable"`
	// The summary of the chat message.
	Summary string                                                      `json:"summary,nullable"`
	JSON    messageGetContentResponseDataMessageChatContentTemplateJSON `json:"-"`
}

// messageGetContentResponseDataMessageChatContentTemplateJSON contains the JSON
// metadata for the struct
// [MessageGetContentResponseDataMessageChatContentTemplate]
type messageGetContentResponseDataMessageChatContentTemplateJSON struct {
	Blocks      apijson.Field
	JsonContent apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageChatContentTemplate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageChatContentTemplateJSON) RawJSON() string {
	return r.raw
}

// A block in a message in a chat.
type MessageGetContentResponseDataMessageChatContentTemplateBlock struct {
	// The actual content of the block.
	Content string `json:"content,required"`
	// The name of the block for identification.
	Name string `json:"name,required"`
	// The type of block in a message in a chat (text or markdown).
	Type MessageGetContentResponseDataMessageChatContentTemplateBlocksType `json:"type,required"`
	JSON messageGetContentResponseDataMessageChatContentTemplateBlockJSON  `json:"-"`
}

// messageGetContentResponseDataMessageChatContentTemplateBlockJSON contains the
// JSON metadata for the struct
// [MessageGetContentResponseDataMessageChatContentTemplateBlock]
type messageGetContentResponseDataMessageChatContentTemplateBlockJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageChatContentTemplateBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageChatContentTemplateBlockJSON) RawJSON() string {
	return r.raw
}

// The type of block in a message in a chat (text or markdown).
type MessageGetContentResponseDataMessageChatContentTemplateBlocksType string

const (
	MessageGetContentResponseDataMessageChatContentTemplateBlocksTypeText     MessageGetContentResponseDataMessageChatContentTemplateBlocksType = "text"
	MessageGetContentResponseDataMessageChatContentTemplateBlocksTypeMarkdown MessageGetContentResponseDataMessageChatContentTemplateBlocksType = "markdown"
)

func (r MessageGetContentResponseDataMessageChatContentTemplateBlocksType) IsKnown() bool {
	switch r {
	case MessageGetContentResponseDataMessageChatContentTemplateBlocksTypeText, MessageGetContentResponseDataMessageChatContentTemplateBlocksTypeMarkdown:
		return true
	}
	return false
}

// The content of an in-app feed message.
type MessageGetContentResponseDataMessageInAppFeedContent struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The blocks of the message in an app feed.
	Blocks []MessageGetContentResponseDataMessageInAppFeedContentBlock `json:"blocks,required"`
	JSON   messageGetContentResponseDataMessageInAppFeedContentJSON    `json:"-"`
}

// messageGetContentResponseDataMessageInAppFeedContentJSON contains the JSON
// metadata for the struct [MessageGetContentResponseDataMessageInAppFeedContent]
type messageGetContentResponseDataMessageInAppFeedContentJSON struct {
	Typename    apijson.Field
	Blocks      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageInAppFeedContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageInAppFeedContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetContentResponseDataMessageInAppFeedContent) implementsMessageGetContentResponseData() {
}

// A block in a message in an app feed.
type MessageGetContentResponseDataMessageInAppFeedContentBlock struct {
	// The name of the block in a message in an app feed.
	Name string `json:"name,required"`
	// The type of block in a message in an app feed.
	Type MessageGetContentResponseDataMessageInAppFeedContentBlocksType `json:"type,required"`
	// This field can have the runtime type of
	// [[]MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton].
	Buttons interface{} `json:"buttons"`
	// The content of the block in a message in an app feed.
	Content string `json:"content"`
	// The rendered HTML version of the content.
	Rendered string                                                        `json:"rendered"`
	JSON     messageGetContentResponseDataMessageInAppFeedContentBlockJSON `json:"-"`
	union    MessageGetContentResponseDataMessageInAppFeedContentBlocksUnion
}

// messageGetContentResponseDataMessageInAppFeedContentBlockJSON contains the JSON
// metadata for the struct
// [MessageGetContentResponseDataMessageInAppFeedContentBlock]
type messageGetContentResponseDataMessageInAppFeedContentBlockJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Buttons     apijson.Field
	Content     apijson.Field
	Rendered    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageGetContentResponseDataMessageInAppFeedContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r *MessageGetContentResponseDataMessageInAppFeedContentBlock) UnmarshalJSON(data []byte) (err error) {
	*r = MessageGetContentResponseDataMessageInAppFeedContentBlock{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock],
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock].
func (r MessageGetContentResponseDataMessageInAppFeedContentBlock) AsUnion() MessageGetContentResponseDataMessageInAppFeedContentBlocksUnion {
	return r.union
}

// A block in a message in an app feed.
//
// Union satisfied by
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock]
// or
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock].
type MessageGetContentResponseDataMessageInAppFeedContentBlocksUnion interface {
	implementsMessageGetContentResponseDataMessageInAppFeedContentBlock()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageGetContentResponseDataMessageInAppFeedContentBlocksUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock{}),
		},
	)
}

// A block in a message in an app feed.
type MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock struct {
	// The content of the block in a message in an app feed.
	Content string `json:"content,required"`
	// The name of the block in a message in an app feed.
	Name string `json:"name,required"`
	// The rendered HTML version of the content.
	Rendered string `json:"rendered,required"`
	// The type of block in a message in an app feed.
	Type MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType `json:"type,required"`
	JSON messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockJSON `json:"-"`
}

// messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockJSON
// contains the JSON metadata for the struct
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock]
type messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Rendered    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock) implementsMessageGetContentResponseDataMessageInAppFeedContentBlock() {
}

// The type of block in a message in an app feed.
type MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType string

const (
	MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockTypeMarkdown MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType = "markdown"
	MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockTypeText     MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType = "text"
)

func (r MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType) IsKnown() bool {
	switch r {
	case MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockTypeMarkdown, MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockTypeText:
		return true
	}
	return false
}

// A button set block in a message in an app feed.
type MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock struct {
	// A list of buttons in an in app feed message.
	Buttons []MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton `json:"buttons,required"`
	// The name of the button set in a message in an app feed.
	Name string `json:"name,required"`
	// The type of block in a message in an app feed.
	Type MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockType `json:"type,required"`
	JSON messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockJSON `json:"-"`
}

// messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockJSON
// contains the JSON metadata for the struct
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock]
type messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockJSON struct {
	Buttons     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock) implementsMessageGetContentResponseDataMessageInAppFeedContentBlock() {
}

// A button in an in app feed message.
type MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton struct {
	// The action to take when the button is clicked.
	Action string `json:"action,required"`
	// The label of the button.
	Label string `json:"label,required"`
	// The name of the button.
	Name string                                                                                             `json:"name,required"`
	JSON messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButtonJSON `json:"-"`
}

// messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButtonJSON
// contains the JSON metadata for the struct
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton]
type messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButtonJSON struct {
	Action      apijson.Field
	Label       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButtonJSON) RawJSON() string {
	return r.raw
}

// The type of block in a message in an app feed.
type MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockType string

const (
	MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockTypeButtonSet MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockType = "button_set"
)

func (r MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockType) IsKnown() bool {
	switch r {
	case MessageGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockTypeButtonSet:
		return true
	}
	return false
}

// The type of block in a message in an app feed.
type MessageGetContentResponseDataMessageInAppFeedContentBlocksType string

const (
	MessageGetContentResponseDataMessageInAppFeedContentBlocksTypeMarkdown  MessageGetContentResponseDataMessageInAppFeedContentBlocksType = "markdown"
	MessageGetContentResponseDataMessageInAppFeedContentBlocksTypeText      MessageGetContentResponseDataMessageInAppFeedContentBlocksType = "text"
	MessageGetContentResponseDataMessageInAppFeedContentBlocksTypeButtonSet MessageGetContentResponseDataMessageInAppFeedContentBlocksType = "button_set"
)

func (r MessageGetContentResponseDataMessageInAppFeedContentBlocksType) IsKnown() bool {
	switch r {
	case MessageGetContentResponseDataMessageInAppFeedContentBlocksTypeMarkdown, MessageGetContentResponseDataMessageInAppFeedContentBlocksTypeText, MessageGetContentResponseDataMessageInAppFeedContentBlocksTypeButtonSet:
		return true
	}
	return false
}

type MessageListParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// Limits the results to items with the corresponding channel ID.
	ChannelID param.Field[string] `query:"channel_id"`
	// Limits the results to messages with the given engagement status.
	EngagementStatus param.Field[[]MessageListParamsEngagementStatus] `query:"engagement_status"`
	InsertedAt       param.Field[MessageListParamsInsertedAt]         `query:"inserted_at"`
	// Limits the results to only the message IDs given (max 50). Note: when using this
	// option, the results will be subject to any other filters applied to the query.
	MessageIDs param.Field[[]string] `query:"message_ids"`
	// The number of items per page (defaults to 50).
	PageSize param.Field[int64] `query:"page_size"`
	// Limits the results to messages triggered by the given workflow key.
	Source param.Field[string] `query:"source"`
	// Limits the results to messages with the given delivery status.
	Status param.Field[[]MessageListParamsStatus] `query:"status"`
	// Limits the results to items with the corresponding tenant.
	Tenant param.Field[string] `query:"tenant"`
	// Limits the results to only messages that were generated with the given data. See
	// [trigger data filtering](/api-reference/overview/trigger-data-filtering) for
	// more information.
	TriggerData param.Field[string] `query:"trigger_data"`
	// Limits the results to messages related to any of the provided categories.
	WorkflowCategories param.Field[[]string] `query:"workflow_categories"`
	// Limits the results to messages for a specific recipient's workflow run.
	WorkflowRecipientRunID param.Field[string] `query:"workflow_recipient_run_id" format:"uuid"`
	// Limits the results to messages associated with the top-level workflow run ID
	// returned by the workflow trigger request.
	WorkflowRunID param.Field[string] `query:"workflow_run_id" format:"uuid"`
}

// URLQuery serializes [MessageListParams]'s query parameters as `url.Values`.
func (r MessageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageListParamsEngagementStatus string

const (
	MessageListParamsEngagementStatusSeen        MessageListParamsEngagementStatus = "seen"
	MessageListParamsEngagementStatusUnseen      MessageListParamsEngagementStatus = "unseen"
	MessageListParamsEngagementStatusRead        MessageListParamsEngagementStatus = "read"
	MessageListParamsEngagementStatusUnread      MessageListParamsEngagementStatus = "unread"
	MessageListParamsEngagementStatusArchived    MessageListParamsEngagementStatus = "archived"
	MessageListParamsEngagementStatusUnarchived  MessageListParamsEngagementStatus = "unarchived"
	MessageListParamsEngagementStatusLinkClicked MessageListParamsEngagementStatus = "link_clicked"
	MessageListParamsEngagementStatusInteracted  MessageListParamsEngagementStatus = "interacted"
)

func (r MessageListParamsEngagementStatus) IsKnown() bool {
	switch r {
	case MessageListParamsEngagementStatusSeen, MessageListParamsEngagementStatusUnseen, MessageListParamsEngagementStatusRead, MessageListParamsEngagementStatusUnread, MessageListParamsEngagementStatusArchived, MessageListParamsEngagementStatusUnarchived, MessageListParamsEngagementStatusLinkClicked, MessageListParamsEngagementStatusInteracted:
		return true
	}
	return false
}

type MessageListParamsInsertedAt struct {
	// Limits the results to messages inserted after the given date.
	Gt param.Field[string] `query:"gt"`
	// Limits the results to messages inserted after or on the given date.
	Gte param.Field[string] `query:"gte"`
	// Limits the results to messages inserted before the given date.
	Lt param.Field[string] `query:"lt"`
	// Limits the results to messages inserted before or on the given date.
	Lte param.Field[string] `query:"lte"`
}

// URLQuery serializes [MessageListParamsInsertedAt]'s query parameters as
// `url.Values`.
func (r MessageListParamsInsertedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageListParamsStatus string

const (
	MessageListParamsStatusQueued            MessageListParamsStatus = "queued"
	MessageListParamsStatusSent              MessageListParamsStatus = "sent"
	MessageListParamsStatusDelivered         MessageListParamsStatus = "delivered"
	MessageListParamsStatusDeliveryAttempted MessageListParamsStatus = "delivery_attempted"
	MessageListParamsStatusUndelivered       MessageListParamsStatus = "undelivered"
	MessageListParamsStatusNotSent           MessageListParamsStatus = "not_sent"
	MessageListParamsStatusBounced           MessageListParamsStatus = "bounced"
)

func (r MessageListParamsStatus) IsKnown() bool {
	switch r {
	case MessageListParamsStatusQueued, MessageListParamsStatusSent, MessageListParamsStatusDelivered, MessageListParamsStatusDeliveryAttempted, MessageListParamsStatusUndelivered, MessageListParamsStatusNotSent, MessageListParamsStatusBounced:
		return true
	}
	return false
}

type MessageListActivitiesParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// The number of items per page (defaults to 50).
	PageSize param.Field[int64] `query:"page_size"`
	// The trigger data to filter activities by.
	TriggerData param.Field[string] `query:"trigger_data"`
}

// URLQuery serializes [MessageListActivitiesParams]'s query parameters as
// `url.Values`.
func (r MessageListActivitiesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageListDeliveryLogsParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// The number of items per page (defaults to 50).
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [MessageListDeliveryLogsParams]'s query parameters as
// `url.Values`.
func (r MessageListDeliveryLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageListEventsParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// The number of items per page (defaults to 50).
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [MessageListEventsParams]'s query parameters as
// `url.Values`.
func (r MessageListEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageMarkAsInteractedParams struct {
	// Metadata about the interaction.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r MessageMarkAsInteractedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

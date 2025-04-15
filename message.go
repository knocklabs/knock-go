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

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/packages/pagination"
	"github.com/stainless-sdks/knock-go/shared"
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

// Returns a paginated list of messages
func (r *MessageService) List(ctx context.Context, query MessageListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Message], err error) {
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

// Returns a paginated list of messages
func (r *MessageService) ListAutoPaging(ctx context.Context, query MessageListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Message] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Archives a message
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

// Retrieves a single message
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

// Get the contents of a message
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

// Get activities for a message
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

// Get activities for a message
func (r *MessageService) ListActivitiesAutoPaging(ctx context.Context, messageID string, query MessageListActivitiesParams, opts ...option.RequestOption) *pagination.ItemsCursorAutoPager[Activity] {
	return pagination.NewItemsCursorAutoPager(r.ListActivities(ctx, messageID, query, opts...))
}

// Get delivery logs for a message
func (r *MessageService) ListDeliveryLogs(ctx context.Context, messageID string, query MessageListDeliveryLogsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[MessageDeliveryLog], err error) {
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

// Get delivery logs for a message
func (r *MessageService) ListDeliveryLogsAutoPaging(ctx context.Context, messageID string, query MessageListDeliveryLogsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[MessageDeliveryLog] {
	return pagination.NewEntriesCursorAutoPager(r.ListDeliveryLogs(ctx, messageID, query, opts...))
}

// Get events for a message
func (r *MessageService) ListEvents(ctx context.Context, messageID string, query MessageListEventsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[MessageEvent], err error) {
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

// Get events for a message
func (r *MessageService) ListEventsAutoPaging(ctx context.Context, messageID string, query MessageListEventsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[MessageEvent] {
	return pagination.NewEntriesCursorAutoPager(r.ListEvents(ctx, messageID, query, opts...))
}

// Marks a message as interacted with
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

// Marks a message as read
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

// Marks a message as seen
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

// Marks a message as unread
func (r *MessageService) MarkAsUnread(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/unread", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Marks a message as unseen
func (r *MessageService) MarkAsUnseen(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/unseen", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Unarchives a message
func (r *MessageService) Unarchive(ctx context.Context, messageID string, opts ...option.RequestOption) (res *Message, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/unarchived", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// An activity associated with a workflow run
type Activity struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor Recipient `json:"actor,nullable"`
	// The data associated with the activity
	Data       interface{} `json:"data,nullable"`
	InsertedAt time.Time   `json:"inserted_at" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient Recipient    `json:"recipient"`
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
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageEngagementStatus `json:"engagement_statuses"`
	// Timestamp of creation
	InsertedAt time.Time `json:"inserted_at" format:"date-time"`
	// Timestamp when message was interacted with
	InteractedAt time.Time `json:"interacted_at,nullable" format:"date-time"`
	// Timestamp when a link in the message was clicked
	LinkClickedAt time.Time `json:"link_clicked_at,nullable" format:"date-time"`
	// Message metadata
	Metadata interface{} `json:"metadata,nullable"`
	// Timestamp when message was read
	ReadAt time.Time `json:"read_at,nullable" format:"date-time"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (id, collection).
	Recipient MessageRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageSource `json:"source"`
	// Message delivery status
	Status MessageStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
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

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or [MessageActorsObject].
type MessageActorsUnion interface {
	ImplementsMessageActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                  `json:"collection,required"`
	JSON       messageActorsObjectJSON `json:"-"`
}

// messageActorsObjectJSON contains the JSON metadata for the struct
// [MessageActorsObject]
type messageActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageActorsObject) ImplementsMessageActorsUnion() {}

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

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or [MessageRecipientObject].
type MessageRecipientUnion interface {
	ImplementsMessageRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                     `json:"collection,required"`
	JSON       messageRecipientObjectJSON `json:"-"`
}

// messageRecipientObjectJSON contains the JSON metadata for the struct
// [MessageRecipientObject]
type messageRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageRecipientObject) ImplementsMessageRecipientUnion() {}

// Source information
type MessageSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string            `json:"version_id,required" format:"uuid"`
	JSON      messageSourceJSON `json:"-"`
}

// messageSourceJSON contains the JSON metadata for the struct [MessageSource]
type messageSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
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

// A message delivery log
type MessageDeliveryLog struct {
	ID            string `json:"id,required"`
	Typename      string `json:"__typename,required"`
	EnvironmentID string `json:"environment_id,required" format:"uuid"`
	InsertedAt    string `json:"inserted_at,required"`
	// A message delivery log request
	Request MessageDeliveryLogRequest `json:"request,required"`
	// A message delivery log response
	Response    MessageDeliveryLogResponse `json:"response,required"`
	ServiceName string                     `json:"service_name,required"`
	JSON        messageDeliveryLogJSON     `json:"-"`
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

// A message delivery log request
type MessageDeliveryLogRequest struct {
	Body    interface{}                     `json:"body"`
	Headers interface{}                     `json:"headers,nullable"`
	Host    string                          `json:"host"`
	Method  MessageDeliveryLogRequestMethod `json:"method"`
	Path    string                          `json:"path"`
	Query   string                          `json:"query,nullable"`
	JSON    messageDeliveryLogRequestJSON   `json:"-"`
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

// A message delivery log response
type MessageDeliveryLogResponse struct {
	Body    interface{}                    `json:"body"`
	Headers interface{}                    `json:"headers,nullable"`
	Status  int64                          `json:"status"`
	JSON    messageDeliveryLogResponseJSON `json:"-"`
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

// A single event that occurred for a message
type MessageEvent struct {
	ID         string    `json:"id,required"`
	Typename   string    `json:"__typename,required"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (id, collection).
	Recipient MessageEventRecipientUnion `json:"recipient,required"`
	Type      MessageEventType           `json:"type,required"`
	// The data associated with the event. Only present for some event types
	Data interface{}      `json:"data,nullable"`
	JSON messageEventJSON `json:"-"`
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

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or [MessageEventRecipientObject].
type MessageEventRecipientUnion interface {
	ImplementsMessageEventRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageEventRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageEventRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageEventRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                          `json:"collection,required"`
	JSON       messageEventRecipientObjectJSON `json:"-"`
}

// messageEventRecipientObjectJSON contains the JSON metadata for the struct
// [MessageEventRecipientObject]
type messageEventRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageEventRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageEventRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageEventRecipientObject) ImplementsMessageEventRecipientUnion() {}

type MessageEventType string

const (
	MessageEventTypeMessageQueued      MessageEventType = "message.queued"
	MessageEventTypeMessageSent        MessageEventType = "message.sent"
	MessageEventTypeMessageDelivered   MessageEventType = "message.delivered"
	MessageEventTypeMessageUndelivered MessageEventType = "message.undelivered"
	MessageEventTypeMessageBounced     MessageEventType = "message.bounced"
	MessageEventTypeMessageRead        MessageEventType = "message.read"
	MessageEventTypeMessageUnread      MessageEventType = "message.unread"
	MessageEventTypeMessageLinkClicked MessageEventType = "message.link_clicked"
	MessageEventTypeMessageInteracted  MessageEventType = "message.interacted"
	MessageEventTypeMessageSeen        MessageEventType = "message.seen"
	MessageEventTypeMessageUnseen      MessageEventType = "message.unseen"
	MessageEventTypeMessageArchived    MessageEventType = "message.archived"
	MessageEventTypeMessageUnarchived  MessageEventType = "message.unarchived"
)

func (r MessageEventType) IsKnown() bool {
	switch r {
	case MessageEventTypeMessageQueued, MessageEventTypeMessageSent, MessageEventTypeMessageDelivered, MessageEventTypeMessageUndelivered, MessageEventTypeMessageBounced, MessageEventTypeMessageRead, MessageEventTypeMessageUnread, MessageEventTypeMessageLinkClicked, MessageEventTypeMessageInteracted, MessageEventTypeMessageSeen, MessageEventTypeMessageUnseen, MessageEventTypeMessageArchived, MessageEventTypeMessageUnarchived:
		return true
	}
	return false
}

// The contents of a message
type MessageGetContentResponse struct {
	Typename string `json:"__typename,required"`
	// The contents of an email message
	Data       MessageGetContentResponseData `json:"data,required"`
	InsertedAt time.Time                     `json:"inserted_at,required" format:"date-time"`
	MessageID  string                        `json:"message_id,required"`
	JSON       messageGetContentResponseJSON `json:"-"`
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

// The contents of an email message
type MessageGetContentResponseData struct {
	Typename string `json:"__typename,required"`
	Token    string `json:"token"`
	Bcc      string `json:"bcc,nullable"`
	// This field can have the runtime type of
	// [[]MessageGetContentResponseDataMessageInAppFeedContentBlock].
	Blocks interface{} `json:"blocks"`
	Body   string      `json:"body"`
	Cc     string      `json:"cc,nullable"`
	// This field can have the runtime type of [interface{}].
	Connection interface{} `json:"connection"`
	// This field can have the runtime type of [interface{}].
	Data     interface{} `json:"data"`
	From     string      `json:"from"`
	HTMLBody string      `json:"html_body"`
	// This field can have the runtime type of [interface{}].
	Metadata    interface{} `json:"metadata"`
	ReplyTo     string      `json:"reply_to,nullable"`
	SubjectLine string      `json:"subject_line"`
	// This field can have the runtime type of
	// [MessageGetContentResponseDataMessageChatContentTemplate].
	Template interface{}                       `json:"template"`
	TextBody string                            `json:"text_body"`
	Title    string                            `json:"title"`
	To       string                            `json:"to"`
	JSON     messageGetContentResponseDataJSON `json:"-"`
	union    MessageGetContentResponseDataUnion
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

// The contents of an email message
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

// The contents of an email message
type MessageGetContentResponseDataMessageEmailContent struct {
	Typename    string                                               `json:"__typename,required"`
	From        string                                               `json:"from,required"`
	HTMLBody    string                                               `json:"html_body,required"`
	SubjectLine string                                               `json:"subject_line,required"`
	TextBody    string                                               `json:"text_body,required"`
	To          string                                               `json:"to,required"`
	Bcc         string                                               `json:"bcc,nullable"`
	Cc          string                                               `json:"cc,nullable"`
	ReplyTo     string                                               `json:"reply_to,nullable"`
	JSON        messageGetContentResponseDataMessageEmailContentJSON `json:"-"`
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

// The contents of an SMS message
type MessageGetContentResponseDataMessageSMSContent struct {
	Typename string                                             `json:"__typename,required"`
	Body     string                                             `json:"body,required"`
	To       string                                             `json:"to,required"`
	JSON     messageGetContentResponseDataMessageSMSContentJSON `json:"-"`
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

// The contents of a push message
type MessageGetContentResponseDataMessagePushContent struct {
	Token    string                                              `json:"token,required"`
	Typename string                                              `json:"__typename,required"`
	Body     string                                              `json:"body,required"`
	Title    string                                              `json:"title,required"`
	Data     interface{}                                         `json:"data,nullable"`
	JSON     messageGetContentResponseDataMessagePushContentJSON `json:"-"`
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

// The contents of a chat message
type MessageGetContentResponseDataMessageChatContent struct {
	Typename string `json:"__typename,required"`
	// The channel data connection from the recipient to the underlying provider
	Connection interface{}                                             `json:"connection,required"`
	Template   MessageGetContentResponseDataMessageChatContentTemplate `json:"template,required"`
	Metadata   interface{}                                             `json:"metadata,nullable"`
	JSON       messageGetContentResponseDataMessageChatContentJSON     `json:"-"`
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

type MessageGetContentResponseDataMessageChatContentTemplate struct {
	// The structured blocks of the message
	Blocks []MessageGetContentResponseDataMessageChatContentTemplateBlock `json:"blocks,nullable"`
	// The JSON content of the message
	JsonContent interface{}                                                 `json:"json_content,nullable"`
	Summary     string                                                      `json:"summary,nullable"`
	JSON        messageGetContentResponseDataMessageChatContentTemplateJSON `json:"-"`
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

// A block in a chat message
type MessageGetContentResponseDataMessageChatContentTemplateBlock struct {
	Content string                                                            `json:"content,required"`
	Name    string                                                            `json:"name,required"`
	Type    MessageGetContentResponseDataMessageChatContentTemplateBlocksType `json:"type,required"`
	JSON    messageGetContentResponseDataMessageChatContentTemplateBlockJSON  `json:"-"`
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

// The contents of a message in an app feed
type MessageGetContentResponseDataMessageInAppFeedContent struct {
	Typename string `json:"__typename,required"`
	// The blocks of the message
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

// A content (text or markdown) block in a message in an app feed
type MessageGetContentResponseDataMessageInAppFeedContentBlock struct {
	Name string                                                         `json:"name,required"`
	Type MessageGetContentResponseDataMessageInAppFeedContentBlocksType `json:"type,required"`
	// This field can have the runtime type of
	// [[]MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton].
	Buttons  interface{}                                                   `json:"buttons"`
	Content  string                                                        `json:"content"`
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
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlock],
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock].
func (r MessageGetContentResponseDataMessageInAppFeedContentBlock) AsUnion() MessageGetContentResponseDataMessageInAppFeedContentBlocksUnion {
	return r.union
}

// A content (text or markdown) block in a message in an app feed
//
// Union satisfied by
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlock] or
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock].
type MessageGetContentResponseDataMessageInAppFeedContentBlocksUnion interface {
	implementsMessageGetContentResponseDataMessageInAppFeedContentBlock()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageGetContentResponseDataMessageInAppFeedContentBlocksUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlock{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock{}),
		},
	)
}

// A content (text or markdown) block in a message in an app feed
type MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlock struct {
	Content  string                                                                     `json:"content,required"`
	Name     string                                                                     `json:"name,required"`
	Rendered string                                                                     `json:"rendered,required"`
	Type     MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType `json:"type,required"`
	JSON     messageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockJSON `json:"-"`
}

// messageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockJSON
// contains the JSON metadata for the struct
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlock]
type messageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Rendered    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlock) implementsMessageGetContentResponseDataMessageInAppFeedContentBlock() {
}

type MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType string

const (
	MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockTypeMarkdown MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType = "markdown"
	MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockTypeText     MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType = "text"
)

func (r MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType) IsKnown() bool {
	switch r {
	case MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockTypeMarkdown, MessageGetContentResponseDataMessageInAppFeedContentBlocksContentBlockTypeText:
		return true
	}
	return false
}

// A set of buttons in a message in an app feed
type MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock struct {
	Buttons []MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton `json:"buttons,required"`
	Name    string                                                                           `json:"name,required"`
	Type    MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockType     `json:"type,required"`
	JSON    messageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockJSON     `json:"-"`
}

// messageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockJSON
// contains the JSON metadata for the struct
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock]
type messageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockJSON struct {
	Buttons     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock) implementsMessageGetContentResponseDataMessageInAppFeedContentBlock() {
}

// A button in a set of buttons
type MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton struct {
	Action string                                                                             `json:"action,required"`
	Label  string                                                                             `json:"label,required"`
	Name   string                                                                             `json:"name,required"`
	JSON   messageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButtonJSON `json:"-"`
}

// messageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButtonJSON
// contains the JSON metadata for the struct
// [MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton]
type messageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButtonJSON struct {
	Action      apijson.Field
	Label       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButtonJSON) RawJSON() string {
	return r.raw
}

type MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockType string

const (
	MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockTypeButtonSet MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockType = "button_set"
)

func (r MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockType) IsKnown() bool {
	switch r {
	case MessageGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockTypeButtonSet:
		return true
	}
	return false
}

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
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The channel ID
	ChannelID param.Field[string] `query:"channel_id"`
	// The engagement status of the message
	EngagementStatus param.Field[[]MessageListParamsEngagementStatus] `query:"engagement_status"`
	// The message IDs to filter messages by
	MessageIDs param.Field[[]string] `query:"message_ids"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// The source of the message (workflow key)
	Source param.Field[string] `query:"source"`
	// The status of the message
	Status param.Field[[]MessageListParamsStatus] `query:"status"`
	// The tenant ID
	Tenant param.Field[string] `query:"tenant"`
	// The trigger data to filter messages by. Must be a valid JSON object.
	TriggerData param.Field[string] `query:"trigger_data"`
	// The workflow categories to filter messages by
	WorkflowCategories param.Field[[]string] `query:"workflow_categories"`
	// The workflow recipient run ID to filter messages by
	WorkflowRecipientRunID param.Field[string] `query:"workflow_recipient_run_id" format:"uuid"`
	// The workflow run ID to filter messages by
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
	MessageListParamsEngagementStatusRead        MessageListParamsEngagementStatus = "read"
	MessageListParamsEngagementStatusInteracted  MessageListParamsEngagementStatus = "interacted"
	MessageListParamsEngagementStatusLinkClicked MessageListParamsEngagementStatus = "link_clicked"
	MessageListParamsEngagementStatusArchived    MessageListParamsEngagementStatus = "archived"
)

func (r MessageListParamsEngagementStatus) IsKnown() bool {
	switch r {
	case MessageListParamsEngagementStatusSeen, MessageListParamsEngagementStatusRead, MessageListParamsEngagementStatusInteracted, MessageListParamsEngagementStatusLinkClicked, MessageListParamsEngagementStatusArchived:
		return true
	}
	return false
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
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// The trigger data to filter activities by
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
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
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
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
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
	// Metadata about the interaction
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r MessageMarkAsInteractedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

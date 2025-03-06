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
func (r *MessageService) List(ctx context.Context, query MessageListParams, opts ...option.RequestOption) (res *MessageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Archives a message
func (r *MessageService) Archive(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageArchiveResponse, err error) {
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
func (r *MessageService) Get(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageGetResponse, err error) {
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
func (r *MessageService) ListActivities(ctx context.Context, messageID string, query MessageListActivitiesParams, opts ...option.RequestOption) (res *MessageListActivitiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/activities", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get delivery logs for a message
func (r *MessageService) ListDeliveryLogs(ctx context.Context, messageID string, query MessageListDeliveryLogsParams, opts ...option.RequestOption) (res *MessageListDeliveryLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/delivery_logs", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get events for a message
func (r *MessageService) ListEvents(ctx context.Context, messageID string, query MessageListEventsParams, opts ...option.RequestOption) (res *MessageListEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/events", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Marks a message as interacted with
func (r *MessageService) MarkAsInteracted(ctx context.Context, messageID string, body MessageMarkAsInteractedParams, opts ...option.RequestOption) (res *MessageMarkAsInteractedResponse, err error) {
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
func (r *MessageService) MarkAsRead(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageMarkAsReadResponse, err error) {
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
func (r *MessageService) MarkAsSeen(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageMarkAsSeenResponse, err error) {
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
func (r *MessageService) MarkAsUnread(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageMarkAsUnreadResponse, err error) {
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
func (r *MessageService) MarkAsUnseen(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageMarkAsUnseenResponse, err error) {
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
func (r *MessageService) Unarchive(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageUnarchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/unarchived", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A paginated list of messages.
type MessageListResponse struct {
	// The list of messages
	Entries []MessageListResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo MessageListResponsePageInfo `json:"page_info,required"`
	JSON     messageListResponseJSON     `json:"-"`
}

// messageListResponseJSON contains the JSON metadata for the struct
// [MessageListResponse]
type messageListResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single message that was generated by a workflow for a given
// channel.
type MessageListResponseEntry struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageListResponseEntriesActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageListResponseEntriesEngagementStatus `json:"engagement_statuses"`
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
	Recipient MessageListResponseEntriesRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageListResponseEntriesSource `json:"source"`
	// Message delivery status
	Status MessageListResponseEntriesStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                       `json:"workflow,nullable"`
	JSON     messageListResponseEntryJSON `json:"-"`
}

// messageListResponseEntryJSON contains the JSON metadata for the struct
// [MessageListResponseEntry]
type messageListResponseEntryJSON struct {
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

func (r *MessageListResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageListResponseEntriesActorsObject].
type MessageListResponseEntriesActorsUnion interface {
	ImplementsMessageListResponseEntriesActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageListResponseEntriesActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageListResponseEntriesActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageListResponseEntriesActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                     `json:"collection,required"`
	JSON       messageListResponseEntriesActorsObjectJSON `json:"-"`
}

// messageListResponseEntriesActorsObjectJSON contains the JSON metadata for the
// struct [MessageListResponseEntriesActorsObject]
type messageListResponseEntriesActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListResponseEntriesActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListResponseEntriesActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageListResponseEntriesActorsObject) ImplementsMessageListResponseEntriesActorsUnion() {}

type MessageListResponseEntriesEngagementStatus string

const (
	MessageListResponseEntriesEngagementStatusSeen        MessageListResponseEntriesEngagementStatus = "seen"
	MessageListResponseEntriesEngagementStatusRead        MessageListResponseEntriesEngagementStatus = "read"
	MessageListResponseEntriesEngagementStatusInteracted  MessageListResponseEntriesEngagementStatus = "interacted"
	MessageListResponseEntriesEngagementStatusLinkClicked MessageListResponseEntriesEngagementStatus = "link_clicked"
	MessageListResponseEntriesEngagementStatusArchived    MessageListResponseEntriesEngagementStatus = "archived"
)

func (r MessageListResponseEntriesEngagementStatus) IsKnown() bool {
	switch r {
	case MessageListResponseEntriesEngagementStatusSeen, MessageListResponseEntriesEngagementStatusRead, MessageListResponseEntriesEngagementStatusInteracted, MessageListResponseEntriesEngagementStatusLinkClicked, MessageListResponseEntriesEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageListResponseEntriesRecipientObject].
type MessageListResponseEntriesRecipientUnion interface {
	ImplementsMessageListResponseEntriesRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageListResponseEntriesRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageListResponseEntriesRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageListResponseEntriesRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                        `json:"collection,required"`
	JSON       messageListResponseEntriesRecipientObjectJSON `json:"-"`
}

// messageListResponseEntriesRecipientObjectJSON contains the JSON metadata for the
// struct [MessageListResponseEntriesRecipientObject]
type messageListResponseEntriesRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListResponseEntriesRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListResponseEntriesRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageListResponseEntriesRecipientObject) ImplementsMessageListResponseEntriesRecipientUnion() {
}

// Source information
type MessageListResponseEntriesSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                               `json:"version_id,required" format:"uuid"`
	JSON      messageListResponseEntriesSourceJSON `json:"-"`
}

// messageListResponseEntriesSourceJSON contains the JSON metadata for the struct
// [MessageListResponseEntriesSource]
type messageListResponseEntriesSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListResponseEntriesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListResponseEntriesSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type MessageListResponseEntriesStatus string

const (
	MessageListResponseEntriesStatusQueued            MessageListResponseEntriesStatus = "queued"
	MessageListResponseEntriesStatusSent              MessageListResponseEntriesStatus = "sent"
	MessageListResponseEntriesStatusDelivered         MessageListResponseEntriesStatus = "delivered"
	MessageListResponseEntriesStatusDeliveryAttempted MessageListResponseEntriesStatus = "delivery_attempted"
	MessageListResponseEntriesStatusUndelivered       MessageListResponseEntriesStatus = "undelivered"
	MessageListResponseEntriesStatusNotSent           MessageListResponseEntriesStatus = "not_sent"
	MessageListResponseEntriesStatusBounced           MessageListResponseEntriesStatus = "bounced"
)

func (r MessageListResponseEntriesStatus) IsKnown() bool {
	switch r {
	case MessageListResponseEntriesStatusQueued, MessageListResponseEntriesStatusSent, MessageListResponseEntriesStatusDelivered, MessageListResponseEntriesStatusDeliveryAttempted, MessageListResponseEntriesStatusUndelivered, MessageListResponseEntriesStatusNotSent, MessageListResponseEntriesStatusBounced:
		return true
	}
	return false
}

// The information about a paginated result
type MessageListResponsePageInfo struct {
	Typename string                          `json:"__typename,required"`
	PageSize int64                           `json:"page_size,required"`
	After    string                          `json:"after,nullable"`
	Before   string                          `json:"before,nullable"`
	JSON     messageListResponsePageInfoJSON `json:"-"`
}

// messageListResponsePageInfoJSON contains the JSON metadata for the struct
// [MessageListResponsePageInfo]
type messageListResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

// Represents a single message that was generated by a workflow for a given
// channel.
type MessageArchiveResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageArchiveResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageArchiveResponseEngagementStatus `json:"engagement_statuses"`
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
	Recipient MessageArchiveResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageArchiveResponseSource `json:"source"`
	// Message delivery status
	Status MessageArchiveResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                     `json:"workflow,nullable"`
	JSON     messageArchiveResponseJSON `json:"-"`
}

// messageArchiveResponseJSON contains the JSON metadata for the struct
// [MessageArchiveResponse]
type messageArchiveResponseJSON struct {
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

func (r *MessageArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageArchiveResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or [MessageArchiveResponseActorsObject].
type MessageArchiveResponseActorsUnion interface {
	ImplementsMessageArchiveResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageArchiveResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageArchiveResponseActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageArchiveResponseActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                 `json:"collection,required"`
	JSON       messageArchiveResponseActorsObjectJSON `json:"-"`
}

// messageArchiveResponseActorsObjectJSON contains the JSON metadata for the struct
// [MessageArchiveResponseActorsObject]
type messageArchiveResponseActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageArchiveResponseActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageArchiveResponseActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageArchiveResponseActorsObject) ImplementsMessageArchiveResponseActorsUnion() {}

type MessageArchiveResponseEngagementStatus string

const (
	MessageArchiveResponseEngagementStatusSeen        MessageArchiveResponseEngagementStatus = "seen"
	MessageArchiveResponseEngagementStatusRead        MessageArchiveResponseEngagementStatus = "read"
	MessageArchiveResponseEngagementStatusInteracted  MessageArchiveResponseEngagementStatus = "interacted"
	MessageArchiveResponseEngagementStatusLinkClicked MessageArchiveResponseEngagementStatus = "link_clicked"
	MessageArchiveResponseEngagementStatusArchived    MessageArchiveResponseEngagementStatus = "archived"
)

func (r MessageArchiveResponseEngagementStatus) IsKnown() bool {
	switch r {
	case MessageArchiveResponseEngagementStatusSeen, MessageArchiveResponseEngagementStatusRead, MessageArchiveResponseEngagementStatusInteracted, MessageArchiveResponseEngagementStatusLinkClicked, MessageArchiveResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageArchiveResponseRecipientObject].
type MessageArchiveResponseRecipientUnion interface {
	ImplementsMessageArchiveResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageArchiveResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageArchiveResponseRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageArchiveResponseRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                    `json:"collection,required"`
	JSON       messageArchiveResponseRecipientObjectJSON `json:"-"`
}

// messageArchiveResponseRecipientObjectJSON contains the JSON metadata for the
// struct [MessageArchiveResponseRecipientObject]
type messageArchiveResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageArchiveResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageArchiveResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageArchiveResponseRecipientObject) ImplementsMessageArchiveResponseRecipientUnion() {}

// Source information
type MessageArchiveResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                           `json:"version_id,required" format:"uuid"`
	JSON      messageArchiveResponseSourceJSON `json:"-"`
}

// messageArchiveResponseSourceJSON contains the JSON metadata for the struct
// [MessageArchiveResponseSource]
type messageArchiveResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageArchiveResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageArchiveResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type MessageArchiveResponseStatus string

const (
	MessageArchiveResponseStatusQueued            MessageArchiveResponseStatus = "queued"
	MessageArchiveResponseStatusSent              MessageArchiveResponseStatus = "sent"
	MessageArchiveResponseStatusDelivered         MessageArchiveResponseStatus = "delivered"
	MessageArchiveResponseStatusDeliveryAttempted MessageArchiveResponseStatus = "delivery_attempted"
	MessageArchiveResponseStatusUndelivered       MessageArchiveResponseStatus = "undelivered"
	MessageArchiveResponseStatusNotSent           MessageArchiveResponseStatus = "not_sent"
	MessageArchiveResponseStatusBounced           MessageArchiveResponseStatus = "bounced"
)

func (r MessageArchiveResponseStatus) IsKnown() bool {
	switch r {
	case MessageArchiveResponseStatusQueued, MessageArchiveResponseStatusSent, MessageArchiveResponseStatusDelivered, MessageArchiveResponseStatusDeliveryAttempted, MessageArchiveResponseStatusUndelivered, MessageArchiveResponseStatusNotSent, MessageArchiveResponseStatusBounced:
		return true
	}
	return false
}

// Represents a single message that was generated by a workflow for a given
// channel.
type MessageGetResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageGetResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageGetResponseEngagementStatus `json:"engagement_statuses"`
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
	Recipient MessageGetResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageGetResponseSource `json:"source"`
	// Message delivery status
	Status MessageGetResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                 `json:"workflow,nullable"`
	JSON     messageGetResponseJSON `json:"-"`
}

// messageGetResponseJSON contains the JSON metadata for the struct
// [MessageGetResponse]
type messageGetResponseJSON struct {
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

func (r *MessageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or [MessageGetResponseActorsObject].
type MessageGetResponseActorsUnion interface {
	ImplementsMessageGetResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageGetResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetResponseActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageGetResponseActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                             `json:"collection,required"`
	JSON       messageGetResponseActorsObjectJSON `json:"-"`
}

// messageGetResponseActorsObjectJSON contains the JSON metadata for the struct
// [MessageGetResponseActorsObject]
type messageGetResponseActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetResponseActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetResponseActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetResponseActorsObject) ImplementsMessageGetResponseActorsUnion() {}

type MessageGetResponseEngagementStatus string

const (
	MessageGetResponseEngagementStatusSeen        MessageGetResponseEngagementStatus = "seen"
	MessageGetResponseEngagementStatusRead        MessageGetResponseEngagementStatus = "read"
	MessageGetResponseEngagementStatusInteracted  MessageGetResponseEngagementStatus = "interacted"
	MessageGetResponseEngagementStatusLinkClicked MessageGetResponseEngagementStatus = "link_clicked"
	MessageGetResponseEngagementStatusArchived    MessageGetResponseEngagementStatus = "archived"
)

func (r MessageGetResponseEngagementStatus) IsKnown() bool {
	switch r {
	case MessageGetResponseEngagementStatusSeen, MessageGetResponseEngagementStatusRead, MessageGetResponseEngagementStatusInteracted, MessageGetResponseEngagementStatusLinkClicked, MessageGetResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or [MessageGetResponseRecipientObject].
type MessageGetResponseRecipientUnion interface {
	ImplementsMessageGetResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageGetResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageGetResponseRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageGetResponseRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                `json:"collection,required"`
	JSON       messageGetResponseRecipientObjectJSON `json:"-"`
}

// messageGetResponseRecipientObjectJSON contains the JSON metadata for the struct
// [MessageGetResponseRecipientObject]
type messageGetResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageGetResponseRecipientObject) ImplementsMessageGetResponseRecipientUnion() {}

// Source information
type MessageGetResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                       `json:"version_id,required" format:"uuid"`
	JSON      messageGetResponseSourceJSON `json:"-"`
}

// messageGetResponseSourceJSON contains the JSON metadata for the struct
// [MessageGetResponseSource]
type messageGetResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageGetResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageGetResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type MessageGetResponseStatus string

const (
	MessageGetResponseStatusQueued            MessageGetResponseStatus = "queued"
	MessageGetResponseStatusSent              MessageGetResponseStatus = "sent"
	MessageGetResponseStatusDelivered         MessageGetResponseStatus = "delivered"
	MessageGetResponseStatusDeliveryAttempted MessageGetResponseStatus = "delivery_attempted"
	MessageGetResponseStatusUndelivered       MessageGetResponseStatus = "undelivered"
	MessageGetResponseStatusNotSent           MessageGetResponseStatus = "not_sent"
	MessageGetResponseStatusBounced           MessageGetResponseStatus = "bounced"
)

func (r MessageGetResponseStatus) IsKnown() bool {
	switch r {
	case MessageGetResponseStatusQueued, MessageGetResponseStatusSent, MessageGetResponseStatusDelivered, MessageGetResponseStatusDeliveryAttempted, MessageGetResponseStatusUndelivered, MessageGetResponseStatusNotSent, MessageGetResponseStatusBounced:
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

// A paginated list of activities
type MessageListActivitiesResponse struct {
	Items []MessageListActivitiesResponseItem `json:"items,required"`
	// The information about a paginated result
	PageInfo MessageListActivitiesResponsePageInfo `json:"page_info,required"`
	JSON     messageListActivitiesResponseJSON     `json:"-"`
}

// messageListActivitiesResponseJSON contains the JSON metadata for the struct
// [MessageListActivitiesResponse]
type messageListActivitiesResponseJSON struct {
	Items       apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListActivitiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListActivitiesResponseJSON) RawJSON() string {
	return r.raw
}

// An activity associated with a workflow run
type MessageListActivitiesResponseItem struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor MessageListActivitiesResponseItemsActor `json:"actor,nullable"`
	// The data associated with the activity
	Data       interface{} `json:"data,nullable"`
	InsertedAt time.Time   `json:"inserted_at" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient MessageListActivitiesResponseItemsRecipient `json:"recipient"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      messageListActivitiesResponseItemJSON       `json:"-"`
}

// messageListActivitiesResponseItemJSON contains the JSON metadata for the struct
// [MessageListActivitiesResponseItem]
type messageListActivitiesResponseItemJSON struct {
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

func (r *MessageListActivitiesResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListActivitiesResponseItemJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type MessageListActivitiesResponseItemsActor struct {
	ID          string                                      `json:"id,required"`
	Typename    string                                      `json:"__typename,required"`
	UpdatedAt   time.Time                                   `json:"updated_at,required" format:"date-time"`
	Avatar      string                                      `json:"avatar,nullable"`
	Collection  string                                      `json:"collection"`
	CreatedAt   time.Time                                   `json:"created_at,nullable" format:"date-time"`
	Email       string                                      `json:"email,nullable" format:"email"`
	Name        string                                      `json:"name,nullable"`
	PhoneNumber string                                      `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                                      `json:"timezone,nullable"`
	JSON        messageListActivitiesResponseItemsActorJSON `json:"-"`
	union       MessageListActivitiesResponseItemsActorUnion
}

// messageListActivitiesResponseItemsActorJSON contains the JSON metadata for the
// struct [MessageListActivitiesResponseItemsActor]
type messageListActivitiesResponseItemsActorJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageListActivitiesResponseItemsActorJSON) RawJSON() string {
	return r.raw
}

func (r *MessageListActivitiesResponseItemsActor) UnmarshalJSON(data []byte) (err error) {
	*r = MessageListActivitiesResponseItemsActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageListActivitiesResponseItemsActorUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [MessageListActivitiesResponseItemsActorObject].
func (r MessageListActivitiesResponseItemsActor) AsUnion() MessageListActivitiesResponseItemsActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [MessageListActivitiesResponseItemsActorObject].
type MessageListActivitiesResponseItemsActorUnion interface {
	implementsMessageListActivitiesResponseItemsActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageListActivitiesResponseItemsActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageListActivitiesResponseItemsActorObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type MessageListActivitiesResponseItemsActorObject struct {
	ID          string                                            `json:"id,required"`
	Typename    string                                            `json:"__typename,required"`
	Collection  string                                            `json:"collection,required"`
	UpdatedAt   time.Time                                         `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                         `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                            `json:"-,extras"`
	JSON        messageListActivitiesResponseItemsActorObjectJSON `json:"-"`
}

// messageListActivitiesResponseItemsActorObjectJSON contains the JSON metadata for
// the struct [MessageListActivitiesResponseItemsActorObject]
type messageListActivitiesResponseItemsActorObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListActivitiesResponseItemsActorObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListActivitiesResponseItemsActorObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageListActivitiesResponseItemsActorObject) implementsMessageListActivitiesResponseItemsActor() {
}

// A recipient, which is either a user or an object
type MessageListActivitiesResponseItemsRecipient struct {
	ID          string                                          `json:"id,required"`
	Typename    string                                          `json:"__typename,required"`
	UpdatedAt   time.Time                                       `json:"updated_at,required" format:"date-time"`
	Avatar      string                                          `json:"avatar,nullable"`
	Collection  string                                          `json:"collection"`
	CreatedAt   time.Time                                       `json:"created_at,nullable" format:"date-time"`
	Email       string                                          `json:"email,nullable" format:"email"`
	Name        string                                          `json:"name,nullable"`
	PhoneNumber string                                          `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                                          `json:"timezone,nullable"`
	JSON        messageListActivitiesResponseItemsRecipientJSON `json:"-"`
	union       MessageListActivitiesResponseItemsRecipientUnion
}

// messageListActivitiesResponseItemsRecipientJSON contains the JSON metadata for
// the struct [MessageListActivitiesResponseItemsRecipient]
type messageListActivitiesResponseItemsRecipientJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageListActivitiesResponseItemsRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *MessageListActivitiesResponseItemsRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = MessageListActivitiesResponseItemsRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageListActivitiesResponseItemsRecipientUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [MessageListActivitiesResponseItemsRecipientObject].
func (r MessageListActivitiesResponseItemsRecipient) AsUnion() MessageListActivitiesResponseItemsRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or
// [MessageListActivitiesResponseItemsRecipientObject].
type MessageListActivitiesResponseItemsRecipientUnion interface {
	implementsMessageListActivitiesResponseItemsRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageListActivitiesResponseItemsRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageListActivitiesResponseItemsRecipientObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type MessageListActivitiesResponseItemsRecipientObject struct {
	ID          string                                                `json:"id,required"`
	Typename    string                                                `json:"__typename,required"`
	Collection  string                                                `json:"collection,required"`
	UpdatedAt   time.Time                                             `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                             `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                                `json:"-,extras"`
	JSON        messageListActivitiesResponseItemsRecipientObjectJSON `json:"-"`
}

// messageListActivitiesResponseItemsRecipientObjectJSON contains the JSON metadata
// for the struct [MessageListActivitiesResponseItemsRecipientObject]
type messageListActivitiesResponseItemsRecipientObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListActivitiesResponseItemsRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListActivitiesResponseItemsRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageListActivitiesResponseItemsRecipientObject) implementsMessageListActivitiesResponseItemsRecipient() {
}

// The information about a paginated result
type MessageListActivitiesResponsePageInfo struct {
	Typename string                                    `json:"__typename,required"`
	PageSize int64                                     `json:"page_size,required"`
	After    string                                    `json:"after,nullable"`
	Before   string                                    `json:"before,nullable"`
	JSON     messageListActivitiesResponsePageInfoJSON `json:"-"`
}

// messageListActivitiesResponsePageInfoJSON contains the JSON metadata for the
// struct [MessageListActivitiesResponsePageInfo]
type messageListActivitiesResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListActivitiesResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListActivitiesResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

// A paginated list of message delivery logs
type MessageListDeliveryLogsResponse struct {
	Entries []MessageListDeliveryLogsResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo MessageListDeliveryLogsResponsePageInfo `json:"page_info,required"`
	JSON     messageListDeliveryLogsResponseJSON     `json:"-"`
}

// messageListDeliveryLogsResponseJSON contains the JSON metadata for the struct
// [MessageListDeliveryLogsResponse]
type messageListDeliveryLogsResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListDeliveryLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListDeliveryLogsResponseJSON) RawJSON() string {
	return r.raw
}

// A message delivery log
type MessageListDeliveryLogsResponseEntry struct {
	ID            string `json:"id,required"`
	Typename      string `json:"__typename,required"`
	EnvironmentID string `json:"environment_id,required" format:"uuid"`
	InsertedAt    string `json:"inserted_at,required"`
	// A message delivery log request
	Request MessageListDeliveryLogsResponseEntriesRequest `json:"request,required"`
	// A message delivery log response
	Response    MessageListDeliveryLogsResponseEntriesResponse `json:"response,required"`
	ServiceName string                                         `json:"service_name,required"`
	JSON        messageListDeliveryLogsResponseEntryJSON       `json:"-"`
}

// messageListDeliveryLogsResponseEntryJSON contains the JSON metadata for the
// struct [MessageListDeliveryLogsResponseEntry]
type messageListDeliveryLogsResponseEntryJSON struct {
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

func (r *MessageListDeliveryLogsResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListDeliveryLogsResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A message delivery log request
type MessageListDeliveryLogsResponseEntriesRequest struct {
	Body    interface{}                                         `json:"body"`
	Headers interface{}                                         `json:"headers,nullable"`
	Host    string                                              `json:"host"`
	Method  MessageListDeliveryLogsResponseEntriesRequestMethod `json:"method"`
	Path    string                                              `json:"path"`
	Query   string                                              `json:"query,nullable"`
	JSON    messageListDeliveryLogsResponseEntriesRequestJSON   `json:"-"`
}

// messageListDeliveryLogsResponseEntriesRequestJSON contains the JSON metadata for
// the struct [MessageListDeliveryLogsResponseEntriesRequest]
type messageListDeliveryLogsResponseEntriesRequestJSON struct {
	Body        apijson.Field
	Headers     apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListDeliveryLogsResponseEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListDeliveryLogsResponseEntriesRequestJSON) RawJSON() string {
	return r.raw
}

type MessageListDeliveryLogsResponseEntriesRequestMethod string

const (
	MessageListDeliveryLogsResponseEntriesRequestMethodGet    MessageListDeliveryLogsResponseEntriesRequestMethod = "GET"
	MessageListDeliveryLogsResponseEntriesRequestMethodPost   MessageListDeliveryLogsResponseEntriesRequestMethod = "POST"
	MessageListDeliveryLogsResponseEntriesRequestMethodPut    MessageListDeliveryLogsResponseEntriesRequestMethod = "PUT"
	MessageListDeliveryLogsResponseEntriesRequestMethodDelete MessageListDeliveryLogsResponseEntriesRequestMethod = "DELETE"
	MessageListDeliveryLogsResponseEntriesRequestMethodPatch  MessageListDeliveryLogsResponseEntriesRequestMethod = "PATCH"
)

func (r MessageListDeliveryLogsResponseEntriesRequestMethod) IsKnown() bool {
	switch r {
	case MessageListDeliveryLogsResponseEntriesRequestMethodGet, MessageListDeliveryLogsResponseEntriesRequestMethodPost, MessageListDeliveryLogsResponseEntriesRequestMethodPut, MessageListDeliveryLogsResponseEntriesRequestMethodDelete, MessageListDeliveryLogsResponseEntriesRequestMethodPatch:
		return true
	}
	return false
}

// A message delivery log response
type MessageListDeliveryLogsResponseEntriesResponse struct {
	Body    interface{}                                        `json:"body"`
	Headers interface{}                                        `json:"headers,nullable"`
	Status  int64                                              `json:"status"`
	JSON    messageListDeliveryLogsResponseEntriesResponseJSON `json:"-"`
}

// messageListDeliveryLogsResponseEntriesResponseJSON contains the JSON metadata
// for the struct [MessageListDeliveryLogsResponseEntriesResponse]
type messageListDeliveryLogsResponseEntriesResponseJSON struct {
	Body        apijson.Field
	Headers     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListDeliveryLogsResponseEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListDeliveryLogsResponseEntriesResponseJSON) RawJSON() string {
	return r.raw
}

// The information about a paginated result
type MessageListDeliveryLogsResponsePageInfo struct {
	Typename string                                      `json:"__typename,required"`
	PageSize int64                                       `json:"page_size,required"`
	After    string                                      `json:"after,nullable"`
	Before   string                                      `json:"before,nullable"`
	JSON     messageListDeliveryLogsResponsePageInfoJSON `json:"-"`
}

// messageListDeliveryLogsResponsePageInfoJSON contains the JSON metadata for the
// struct [MessageListDeliveryLogsResponsePageInfo]
type messageListDeliveryLogsResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListDeliveryLogsResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListDeliveryLogsResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

// A paginated list of message events.
type MessageListEventsResponse struct {
	// The list of message events
	Entries []MessageListEventsResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo MessageListEventsResponsePageInfo `json:"page_info,required"`
	JSON     messageListEventsResponseJSON     `json:"-"`
}

// messageListEventsResponseJSON contains the JSON metadata for the struct
// [MessageListEventsResponse]
type messageListEventsResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListEventsResponseJSON) RawJSON() string {
	return r.raw
}

// A single event that occurred for a message
type MessageListEventsResponseEntry struct {
	ID         string    `json:"id,required"`
	Typename   string    `json:"__typename,required"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (id, collection).
	Recipient MessageListEventsResponseEntriesRecipientUnion `json:"recipient,required"`
	Type      MessageListEventsResponseEntriesType           `json:"type,required"`
	// The data associated with the event. Only present for some event types
	Data interface{}                        `json:"data,nullable"`
	JSON messageListEventsResponseEntryJSON `json:"-"`
}

// messageListEventsResponseEntryJSON contains the JSON metadata for the struct
// [MessageListEventsResponseEntry]
type messageListEventsResponseEntryJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	InsertedAt  apijson.Field
	Recipient   apijson.Field
	Type        apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListEventsResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListEventsResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageListEventsResponseEntriesRecipientObject].
type MessageListEventsResponseEntriesRecipientUnion interface {
	ImplementsMessageListEventsResponseEntriesRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageListEventsResponseEntriesRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageListEventsResponseEntriesRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageListEventsResponseEntriesRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                              `json:"collection,required"`
	JSON       messageListEventsResponseEntriesRecipientObjectJSON `json:"-"`
}

// messageListEventsResponseEntriesRecipientObjectJSON contains the JSON metadata
// for the struct [MessageListEventsResponseEntriesRecipientObject]
type messageListEventsResponseEntriesRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListEventsResponseEntriesRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListEventsResponseEntriesRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageListEventsResponseEntriesRecipientObject) ImplementsMessageListEventsResponseEntriesRecipientUnion() {
}

type MessageListEventsResponseEntriesType string

const (
	MessageListEventsResponseEntriesTypeMessageQueued      MessageListEventsResponseEntriesType = "message.queued"
	MessageListEventsResponseEntriesTypeMessageSent        MessageListEventsResponseEntriesType = "message.sent"
	MessageListEventsResponseEntriesTypeMessageDelivered   MessageListEventsResponseEntriesType = "message.delivered"
	MessageListEventsResponseEntriesTypeMessageUndelivered MessageListEventsResponseEntriesType = "message.undelivered"
	MessageListEventsResponseEntriesTypeMessageBounced     MessageListEventsResponseEntriesType = "message.bounced"
	MessageListEventsResponseEntriesTypeMessageRead        MessageListEventsResponseEntriesType = "message.read"
	MessageListEventsResponseEntriesTypeMessageUnread      MessageListEventsResponseEntriesType = "message.unread"
	MessageListEventsResponseEntriesTypeMessageLinkClicked MessageListEventsResponseEntriesType = "message.link_clicked"
	MessageListEventsResponseEntriesTypeMessageInteracted  MessageListEventsResponseEntriesType = "message.interacted"
	MessageListEventsResponseEntriesTypeMessageSeen        MessageListEventsResponseEntriesType = "message.seen"
	MessageListEventsResponseEntriesTypeMessageUnseen      MessageListEventsResponseEntriesType = "message.unseen"
	MessageListEventsResponseEntriesTypeMessageArchived    MessageListEventsResponseEntriesType = "message.archived"
	MessageListEventsResponseEntriesTypeMessageUnarchived  MessageListEventsResponseEntriesType = "message.unarchived"
)

func (r MessageListEventsResponseEntriesType) IsKnown() bool {
	switch r {
	case MessageListEventsResponseEntriesTypeMessageQueued, MessageListEventsResponseEntriesTypeMessageSent, MessageListEventsResponseEntriesTypeMessageDelivered, MessageListEventsResponseEntriesTypeMessageUndelivered, MessageListEventsResponseEntriesTypeMessageBounced, MessageListEventsResponseEntriesTypeMessageRead, MessageListEventsResponseEntriesTypeMessageUnread, MessageListEventsResponseEntriesTypeMessageLinkClicked, MessageListEventsResponseEntriesTypeMessageInteracted, MessageListEventsResponseEntriesTypeMessageSeen, MessageListEventsResponseEntriesTypeMessageUnseen, MessageListEventsResponseEntriesTypeMessageArchived, MessageListEventsResponseEntriesTypeMessageUnarchived:
		return true
	}
	return false
}

// The information about a paginated result
type MessageListEventsResponsePageInfo struct {
	Typename string                                `json:"__typename,required"`
	PageSize int64                                 `json:"page_size,required"`
	After    string                                `json:"after,nullable"`
	Before   string                                `json:"before,nullable"`
	JSON     messageListEventsResponsePageInfoJSON `json:"-"`
}

// messageListEventsResponsePageInfoJSON contains the JSON metadata for the struct
// [MessageListEventsResponsePageInfo]
type messageListEventsResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageListEventsResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageListEventsResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

// Represents a single message that was generated by a workflow for a given
// channel.
type MessageMarkAsInteractedResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageMarkAsInteractedResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageMarkAsInteractedResponseEngagementStatus `json:"engagement_statuses"`
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
	Recipient MessageMarkAsInteractedResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageMarkAsInteractedResponseSource `json:"source"`
	// Message delivery status
	Status MessageMarkAsInteractedResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                              `json:"workflow,nullable"`
	JSON     messageMarkAsInteractedResponseJSON `json:"-"`
}

// messageMarkAsInteractedResponseJSON contains the JSON metadata for the struct
// [MessageMarkAsInteractedResponse]
type messageMarkAsInteractedResponseJSON struct {
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

func (r *MessageMarkAsInteractedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsInteractedResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsInteractedResponseActorsObject].
type MessageMarkAsInteractedResponseActorsUnion interface {
	ImplementsMessageMarkAsInteractedResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsInteractedResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsInteractedResponseActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsInteractedResponseActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                          `json:"collection,required"`
	JSON       messageMarkAsInteractedResponseActorsObjectJSON `json:"-"`
}

// messageMarkAsInteractedResponseActorsObjectJSON contains the JSON metadata for
// the struct [MessageMarkAsInteractedResponseActorsObject]
type messageMarkAsInteractedResponseActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsInteractedResponseActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsInteractedResponseActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsInteractedResponseActorsObject) ImplementsMessageMarkAsInteractedResponseActorsUnion() {
}

type MessageMarkAsInteractedResponseEngagementStatus string

const (
	MessageMarkAsInteractedResponseEngagementStatusSeen        MessageMarkAsInteractedResponseEngagementStatus = "seen"
	MessageMarkAsInteractedResponseEngagementStatusRead        MessageMarkAsInteractedResponseEngagementStatus = "read"
	MessageMarkAsInteractedResponseEngagementStatusInteracted  MessageMarkAsInteractedResponseEngagementStatus = "interacted"
	MessageMarkAsInteractedResponseEngagementStatusLinkClicked MessageMarkAsInteractedResponseEngagementStatus = "link_clicked"
	MessageMarkAsInteractedResponseEngagementStatusArchived    MessageMarkAsInteractedResponseEngagementStatus = "archived"
)

func (r MessageMarkAsInteractedResponseEngagementStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsInteractedResponseEngagementStatusSeen, MessageMarkAsInteractedResponseEngagementStatusRead, MessageMarkAsInteractedResponseEngagementStatusInteracted, MessageMarkAsInteractedResponseEngagementStatusLinkClicked, MessageMarkAsInteractedResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsInteractedResponseRecipientObject].
type MessageMarkAsInteractedResponseRecipientUnion interface {
	ImplementsMessageMarkAsInteractedResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsInteractedResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsInteractedResponseRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsInteractedResponseRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                             `json:"collection,required"`
	JSON       messageMarkAsInteractedResponseRecipientObjectJSON `json:"-"`
}

// messageMarkAsInteractedResponseRecipientObjectJSON contains the JSON metadata
// for the struct [MessageMarkAsInteractedResponseRecipientObject]
type messageMarkAsInteractedResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsInteractedResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsInteractedResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsInteractedResponseRecipientObject) ImplementsMessageMarkAsInteractedResponseRecipientUnion() {
}

// Source information
type MessageMarkAsInteractedResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                                    `json:"version_id,required" format:"uuid"`
	JSON      messageMarkAsInteractedResponseSourceJSON `json:"-"`
}

// messageMarkAsInteractedResponseSourceJSON contains the JSON metadata for the
// struct [MessageMarkAsInteractedResponseSource]
type messageMarkAsInteractedResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsInteractedResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsInteractedResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type MessageMarkAsInteractedResponseStatus string

const (
	MessageMarkAsInteractedResponseStatusQueued            MessageMarkAsInteractedResponseStatus = "queued"
	MessageMarkAsInteractedResponseStatusSent              MessageMarkAsInteractedResponseStatus = "sent"
	MessageMarkAsInteractedResponseStatusDelivered         MessageMarkAsInteractedResponseStatus = "delivered"
	MessageMarkAsInteractedResponseStatusDeliveryAttempted MessageMarkAsInteractedResponseStatus = "delivery_attempted"
	MessageMarkAsInteractedResponseStatusUndelivered       MessageMarkAsInteractedResponseStatus = "undelivered"
	MessageMarkAsInteractedResponseStatusNotSent           MessageMarkAsInteractedResponseStatus = "not_sent"
	MessageMarkAsInteractedResponseStatusBounced           MessageMarkAsInteractedResponseStatus = "bounced"
)

func (r MessageMarkAsInteractedResponseStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsInteractedResponseStatusQueued, MessageMarkAsInteractedResponseStatusSent, MessageMarkAsInteractedResponseStatusDelivered, MessageMarkAsInteractedResponseStatusDeliveryAttempted, MessageMarkAsInteractedResponseStatusUndelivered, MessageMarkAsInteractedResponseStatusNotSent, MessageMarkAsInteractedResponseStatusBounced:
		return true
	}
	return false
}

// Represents a single message that was generated by a workflow for a given
// channel.
type MessageMarkAsReadResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageMarkAsReadResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageMarkAsReadResponseEngagementStatus `json:"engagement_statuses"`
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
	Recipient MessageMarkAsReadResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageMarkAsReadResponseSource `json:"source"`
	// Message delivery status
	Status MessageMarkAsReadResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                        `json:"workflow,nullable"`
	JSON     messageMarkAsReadResponseJSON `json:"-"`
}

// messageMarkAsReadResponseJSON contains the JSON metadata for the struct
// [MessageMarkAsReadResponse]
type messageMarkAsReadResponseJSON struct {
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

func (r *MessageMarkAsReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsReadResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsReadResponseActorsObject].
type MessageMarkAsReadResponseActorsUnion interface {
	ImplementsMessageMarkAsReadResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsReadResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsReadResponseActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsReadResponseActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                    `json:"collection,required"`
	JSON       messageMarkAsReadResponseActorsObjectJSON `json:"-"`
}

// messageMarkAsReadResponseActorsObjectJSON contains the JSON metadata for the
// struct [MessageMarkAsReadResponseActorsObject]
type messageMarkAsReadResponseActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsReadResponseActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsReadResponseActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsReadResponseActorsObject) ImplementsMessageMarkAsReadResponseActorsUnion() {}

type MessageMarkAsReadResponseEngagementStatus string

const (
	MessageMarkAsReadResponseEngagementStatusSeen        MessageMarkAsReadResponseEngagementStatus = "seen"
	MessageMarkAsReadResponseEngagementStatusRead        MessageMarkAsReadResponseEngagementStatus = "read"
	MessageMarkAsReadResponseEngagementStatusInteracted  MessageMarkAsReadResponseEngagementStatus = "interacted"
	MessageMarkAsReadResponseEngagementStatusLinkClicked MessageMarkAsReadResponseEngagementStatus = "link_clicked"
	MessageMarkAsReadResponseEngagementStatusArchived    MessageMarkAsReadResponseEngagementStatus = "archived"
)

func (r MessageMarkAsReadResponseEngagementStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsReadResponseEngagementStatusSeen, MessageMarkAsReadResponseEngagementStatusRead, MessageMarkAsReadResponseEngagementStatusInteracted, MessageMarkAsReadResponseEngagementStatusLinkClicked, MessageMarkAsReadResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsReadResponseRecipientObject].
type MessageMarkAsReadResponseRecipientUnion interface {
	ImplementsMessageMarkAsReadResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsReadResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsReadResponseRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsReadResponseRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                       `json:"collection,required"`
	JSON       messageMarkAsReadResponseRecipientObjectJSON `json:"-"`
}

// messageMarkAsReadResponseRecipientObjectJSON contains the JSON metadata for the
// struct [MessageMarkAsReadResponseRecipientObject]
type messageMarkAsReadResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsReadResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsReadResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsReadResponseRecipientObject) ImplementsMessageMarkAsReadResponseRecipientUnion() {
}

// Source information
type MessageMarkAsReadResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                              `json:"version_id,required" format:"uuid"`
	JSON      messageMarkAsReadResponseSourceJSON `json:"-"`
}

// messageMarkAsReadResponseSourceJSON contains the JSON metadata for the struct
// [MessageMarkAsReadResponseSource]
type messageMarkAsReadResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsReadResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsReadResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type MessageMarkAsReadResponseStatus string

const (
	MessageMarkAsReadResponseStatusQueued            MessageMarkAsReadResponseStatus = "queued"
	MessageMarkAsReadResponseStatusSent              MessageMarkAsReadResponseStatus = "sent"
	MessageMarkAsReadResponseStatusDelivered         MessageMarkAsReadResponseStatus = "delivered"
	MessageMarkAsReadResponseStatusDeliveryAttempted MessageMarkAsReadResponseStatus = "delivery_attempted"
	MessageMarkAsReadResponseStatusUndelivered       MessageMarkAsReadResponseStatus = "undelivered"
	MessageMarkAsReadResponseStatusNotSent           MessageMarkAsReadResponseStatus = "not_sent"
	MessageMarkAsReadResponseStatusBounced           MessageMarkAsReadResponseStatus = "bounced"
)

func (r MessageMarkAsReadResponseStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsReadResponseStatusQueued, MessageMarkAsReadResponseStatusSent, MessageMarkAsReadResponseStatusDelivered, MessageMarkAsReadResponseStatusDeliveryAttempted, MessageMarkAsReadResponseStatusUndelivered, MessageMarkAsReadResponseStatusNotSent, MessageMarkAsReadResponseStatusBounced:
		return true
	}
	return false
}

// Represents a single message that was generated by a workflow for a given
// channel.
type MessageMarkAsSeenResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageMarkAsSeenResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageMarkAsSeenResponseEngagementStatus `json:"engagement_statuses"`
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
	Recipient MessageMarkAsSeenResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageMarkAsSeenResponseSource `json:"source"`
	// Message delivery status
	Status MessageMarkAsSeenResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                        `json:"workflow,nullable"`
	JSON     messageMarkAsSeenResponseJSON `json:"-"`
}

// messageMarkAsSeenResponseJSON contains the JSON metadata for the struct
// [MessageMarkAsSeenResponse]
type messageMarkAsSeenResponseJSON struct {
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

func (r *MessageMarkAsSeenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsSeenResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsSeenResponseActorsObject].
type MessageMarkAsSeenResponseActorsUnion interface {
	ImplementsMessageMarkAsSeenResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsSeenResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsSeenResponseActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsSeenResponseActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                    `json:"collection,required"`
	JSON       messageMarkAsSeenResponseActorsObjectJSON `json:"-"`
}

// messageMarkAsSeenResponseActorsObjectJSON contains the JSON metadata for the
// struct [MessageMarkAsSeenResponseActorsObject]
type messageMarkAsSeenResponseActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsSeenResponseActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsSeenResponseActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsSeenResponseActorsObject) ImplementsMessageMarkAsSeenResponseActorsUnion() {}

type MessageMarkAsSeenResponseEngagementStatus string

const (
	MessageMarkAsSeenResponseEngagementStatusSeen        MessageMarkAsSeenResponseEngagementStatus = "seen"
	MessageMarkAsSeenResponseEngagementStatusRead        MessageMarkAsSeenResponseEngagementStatus = "read"
	MessageMarkAsSeenResponseEngagementStatusInteracted  MessageMarkAsSeenResponseEngagementStatus = "interacted"
	MessageMarkAsSeenResponseEngagementStatusLinkClicked MessageMarkAsSeenResponseEngagementStatus = "link_clicked"
	MessageMarkAsSeenResponseEngagementStatusArchived    MessageMarkAsSeenResponseEngagementStatus = "archived"
)

func (r MessageMarkAsSeenResponseEngagementStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsSeenResponseEngagementStatusSeen, MessageMarkAsSeenResponseEngagementStatusRead, MessageMarkAsSeenResponseEngagementStatusInteracted, MessageMarkAsSeenResponseEngagementStatusLinkClicked, MessageMarkAsSeenResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsSeenResponseRecipientObject].
type MessageMarkAsSeenResponseRecipientUnion interface {
	ImplementsMessageMarkAsSeenResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsSeenResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsSeenResponseRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsSeenResponseRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                       `json:"collection,required"`
	JSON       messageMarkAsSeenResponseRecipientObjectJSON `json:"-"`
}

// messageMarkAsSeenResponseRecipientObjectJSON contains the JSON metadata for the
// struct [MessageMarkAsSeenResponseRecipientObject]
type messageMarkAsSeenResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsSeenResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsSeenResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsSeenResponseRecipientObject) ImplementsMessageMarkAsSeenResponseRecipientUnion() {
}

// Source information
type MessageMarkAsSeenResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                              `json:"version_id,required" format:"uuid"`
	JSON      messageMarkAsSeenResponseSourceJSON `json:"-"`
}

// messageMarkAsSeenResponseSourceJSON contains the JSON metadata for the struct
// [MessageMarkAsSeenResponseSource]
type messageMarkAsSeenResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsSeenResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsSeenResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type MessageMarkAsSeenResponseStatus string

const (
	MessageMarkAsSeenResponseStatusQueued            MessageMarkAsSeenResponseStatus = "queued"
	MessageMarkAsSeenResponseStatusSent              MessageMarkAsSeenResponseStatus = "sent"
	MessageMarkAsSeenResponseStatusDelivered         MessageMarkAsSeenResponseStatus = "delivered"
	MessageMarkAsSeenResponseStatusDeliveryAttempted MessageMarkAsSeenResponseStatus = "delivery_attempted"
	MessageMarkAsSeenResponseStatusUndelivered       MessageMarkAsSeenResponseStatus = "undelivered"
	MessageMarkAsSeenResponseStatusNotSent           MessageMarkAsSeenResponseStatus = "not_sent"
	MessageMarkAsSeenResponseStatusBounced           MessageMarkAsSeenResponseStatus = "bounced"
)

func (r MessageMarkAsSeenResponseStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsSeenResponseStatusQueued, MessageMarkAsSeenResponseStatusSent, MessageMarkAsSeenResponseStatusDelivered, MessageMarkAsSeenResponseStatusDeliveryAttempted, MessageMarkAsSeenResponseStatusUndelivered, MessageMarkAsSeenResponseStatusNotSent, MessageMarkAsSeenResponseStatusBounced:
		return true
	}
	return false
}

// Represents a single message that was generated by a workflow for a given
// channel.
type MessageMarkAsUnreadResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageMarkAsUnreadResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageMarkAsUnreadResponseEngagementStatus `json:"engagement_statuses"`
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
	Recipient MessageMarkAsUnreadResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageMarkAsUnreadResponseSource `json:"source"`
	// Message delivery status
	Status MessageMarkAsUnreadResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                          `json:"workflow,nullable"`
	JSON     messageMarkAsUnreadResponseJSON `json:"-"`
}

// messageMarkAsUnreadResponseJSON contains the JSON metadata for the struct
// [MessageMarkAsUnreadResponse]
type messageMarkAsUnreadResponseJSON struct {
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

func (r *MessageMarkAsUnreadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsUnreadResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsUnreadResponseActorsObject].
type MessageMarkAsUnreadResponseActorsUnion interface {
	ImplementsMessageMarkAsUnreadResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsUnreadResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsUnreadResponseActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsUnreadResponseActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                      `json:"collection,required"`
	JSON       messageMarkAsUnreadResponseActorsObjectJSON `json:"-"`
}

// messageMarkAsUnreadResponseActorsObjectJSON contains the JSON metadata for the
// struct [MessageMarkAsUnreadResponseActorsObject]
type messageMarkAsUnreadResponseActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsUnreadResponseActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsUnreadResponseActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsUnreadResponseActorsObject) ImplementsMessageMarkAsUnreadResponseActorsUnion() {}

type MessageMarkAsUnreadResponseEngagementStatus string

const (
	MessageMarkAsUnreadResponseEngagementStatusSeen        MessageMarkAsUnreadResponseEngagementStatus = "seen"
	MessageMarkAsUnreadResponseEngagementStatusRead        MessageMarkAsUnreadResponseEngagementStatus = "read"
	MessageMarkAsUnreadResponseEngagementStatusInteracted  MessageMarkAsUnreadResponseEngagementStatus = "interacted"
	MessageMarkAsUnreadResponseEngagementStatusLinkClicked MessageMarkAsUnreadResponseEngagementStatus = "link_clicked"
	MessageMarkAsUnreadResponseEngagementStatusArchived    MessageMarkAsUnreadResponseEngagementStatus = "archived"
)

func (r MessageMarkAsUnreadResponseEngagementStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsUnreadResponseEngagementStatusSeen, MessageMarkAsUnreadResponseEngagementStatusRead, MessageMarkAsUnreadResponseEngagementStatusInteracted, MessageMarkAsUnreadResponseEngagementStatusLinkClicked, MessageMarkAsUnreadResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsUnreadResponseRecipientObject].
type MessageMarkAsUnreadResponseRecipientUnion interface {
	ImplementsMessageMarkAsUnreadResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsUnreadResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsUnreadResponseRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsUnreadResponseRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                         `json:"collection,required"`
	JSON       messageMarkAsUnreadResponseRecipientObjectJSON `json:"-"`
}

// messageMarkAsUnreadResponseRecipientObjectJSON contains the JSON metadata for
// the struct [MessageMarkAsUnreadResponseRecipientObject]
type messageMarkAsUnreadResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsUnreadResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsUnreadResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsUnreadResponseRecipientObject) ImplementsMessageMarkAsUnreadResponseRecipientUnion() {
}

// Source information
type MessageMarkAsUnreadResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                                `json:"version_id,required" format:"uuid"`
	JSON      messageMarkAsUnreadResponseSourceJSON `json:"-"`
}

// messageMarkAsUnreadResponseSourceJSON contains the JSON metadata for the struct
// [MessageMarkAsUnreadResponseSource]
type messageMarkAsUnreadResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsUnreadResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsUnreadResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type MessageMarkAsUnreadResponseStatus string

const (
	MessageMarkAsUnreadResponseStatusQueued            MessageMarkAsUnreadResponseStatus = "queued"
	MessageMarkAsUnreadResponseStatusSent              MessageMarkAsUnreadResponseStatus = "sent"
	MessageMarkAsUnreadResponseStatusDelivered         MessageMarkAsUnreadResponseStatus = "delivered"
	MessageMarkAsUnreadResponseStatusDeliveryAttempted MessageMarkAsUnreadResponseStatus = "delivery_attempted"
	MessageMarkAsUnreadResponseStatusUndelivered       MessageMarkAsUnreadResponseStatus = "undelivered"
	MessageMarkAsUnreadResponseStatusNotSent           MessageMarkAsUnreadResponseStatus = "not_sent"
	MessageMarkAsUnreadResponseStatusBounced           MessageMarkAsUnreadResponseStatus = "bounced"
)

func (r MessageMarkAsUnreadResponseStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsUnreadResponseStatusQueued, MessageMarkAsUnreadResponseStatusSent, MessageMarkAsUnreadResponseStatusDelivered, MessageMarkAsUnreadResponseStatusDeliveryAttempted, MessageMarkAsUnreadResponseStatusUndelivered, MessageMarkAsUnreadResponseStatusNotSent, MessageMarkAsUnreadResponseStatusBounced:
		return true
	}
	return false
}

// Represents a single message that was generated by a workflow for a given
// channel.
type MessageMarkAsUnseenResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageMarkAsUnseenResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageMarkAsUnseenResponseEngagementStatus `json:"engagement_statuses"`
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
	Recipient MessageMarkAsUnseenResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageMarkAsUnseenResponseSource `json:"source"`
	// Message delivery status
	Status MessageMarkAsUnseenResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                          `json:"workflow,nullable"`
	JSON     messageMarkAsUnseenResponseJSON `json:"-"`
}

// messageMarkAsUnseenResponseJSON contains the JSON metadata for the struct
// [MessageMarkAsUnseenResponse]
type messageMarkAsUnseenResponseJSON struct {
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

func (r *MessageMarkAsUnseenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsUnseenResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsUnseenResponseActorsObject].
type MessageMarkAsUnseenResponseActorsUnion interface {
	ImplementsMessageMarkAsUnseenResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsUnseenResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsUnseenResponseActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsUnseenResponseActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                      `json:"collection,required"`
	JSON       messageMarkAsUnseenResponseActorsObjectJSON `json:"-"`
}

// messageMarkAsUnseenResponseActorsObjectJSON contains the JSON metadata for the
// struct [MessageMarkAsUnseenResponseActorsObject]
type messageMarkAsUnseenResponseActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsUnseenResponseActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsUnseenResponseActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsUnseenResponseActorsObject) ImplementsMessageMarkAsUnseenResponseActorsUnion() {}

type MessageMarkAsUnseenResponseEngagementStatus string

const (
	MessageMarkAsUnseenResponseEngagementStatusSeen        MessageMarkAsUnseenResponseEngagementStatus = "seen"
	MessageMarkAsUnseenResponseEngagementStatusRead        MessageMarkAsUnseenResponseEngagementStatus = "read"
	MessageMarkAsUnseenResponseEngagementStatusInteracted  MessageMarkAsUnseenResponseEngagementStatus = "interacted"
	MessageMarkAsUnseenResponseEngagementStatusLinkClicked MessageMarkAsUnseenResponseEngagementStatus = "link_clicked"
	MessageMarkAsUnseenResponseEngagementStatusArchived    MessageMarkAsUnseenResponseEngagementStatus = "archived"
)

func (r MessageMarkAsUnseenResponseEngagementStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsUnseenResponseEngagementStatusSeen, MessageMarkAsUnseenResponseEngagementStatusRead, MessageMarkAsUnseenResponseEngagementStatusInteracted, MessageMarkAsUnseenResponseEngagementStatusLinkClicked, MessageMarkAsUnseenResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageMarkAsUnseenResponseRecipientObject].
type MessageMarkAsUnseenResponseRecipientUnion interface {
	ImplementsMessageMarkAsUnseenResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageMarkAsUnseenResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageMarkAsUnseenResponseRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageMarkAsUnseenResponseRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                         `json:"collection,required"`
	JSON       messageMarkAsUnseenResponseRecipientObjectJSON `json:"-"`
}

// messageMarkAsUnseenResponseRecipientObjectJSON contains the JSON metadata for
// the struct [MessageMarkAsUnseenResponseRecipientObject]
type messageMarkAsUnseenResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsUnseenResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsUnseenResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageMarkAsUnseenResponseRecipientObject) ImplementsMessageMarkAsUnseenResponseRecipientUnion() {
}

// Source information
type MessageMarkAsUnseenResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                                `json:"version_id,required" format:"uuid"`
	JSON      messageMarkAsUnseenResponseSourceJSON `json:"-"`
}

// messageMarkAsUnseenResponseSourceJSON contains the JSON metadata for the struct
// [MessageMarkAsUnseenResponseSource]
type messageMarkAsUnseenResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageMarkAsUnseenResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageMarkAsUnseenResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type MessageMarkAsUnseenResponseStatus string

const (
	MessageMarkAsUnseenResponseStatusQueued            MessageMarkAsUnseenResponseStatus = "queued"
	MessageMarkAsUnseenResponseStatusSent              MessageMarkAsUnseenResponseStatus = "sent"
	MessageMarkAsUnseenResponseStatusDelivered         MessageMarkAsUnseenResponseStatus = "delivered"
	MessageMarkAsUnseenResponseStatusDeliveryAttempted MessageMarkAsUnseenResponseStatus = "delivery_attempted"
	MessageMarkAsUnseenResponseStatusUndelivered       MessageMarkAsUnseenResponseStatus = "undelivered"
	MessageMarkAsUnseenResponseStatusNotSent           MessageMarkAsUnseenResponseStatus = "not_sent"
	MessageMarkAsUnseenResponseStatusBounced           MessageMarkAsUnseenResponseStatus = "bounced"
)

func (r MessageMarkAsUnseenResponseStatus) IsKnown() bool {
	switch r {
	case MessageMarkAsUnseenResponseStatusQueued, MessageMarkAsUnseenResponseStatusSent, MessageMarkAsUnseenResponseStatusDelivered, MessageMarkAsUnseenResponseStatusDeliveryAttempted, MessageMarkAsUnseenResponseStatusUndelivered, MessageMarkAsUnseenResponseStatusNotSent, MessageMarkAsUnseenResponseStatusBounced:
		return true
	}
	return false
}

// Represents a single message that was generated by a workflow for a given
// channel.
type MessageUnarchiveResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []MessageUnarchiveResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []MessageUnarchiveResponseEngagementStatus `json:"engagement_statuses"`
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
	Recipient MessageUnarchiveResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source MessageUnarchiveResponseSource `json:"source"`
	// Message delivery status
	Status MessageUnarchiveResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                       `json:"workflow,nullable"`
	JSON     messageUnarchiveResponseJSON `json:"-"`
}

// messageUnarchiveResponseJSON contains the JSON metadata for the struct
// [MessageUnarchiveResponse]
type messageUnarchiveResponseJSON struct {
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

func (r *MessageUnarchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageUnarchiveResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageUnarchiveResponseActorsObject].
type MessageUnarchiveResponseActorsUnion interface {
	ImplementsMessageUnarchiveResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageUnarchiveResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageUnarchiveResponseActorsObject{}),
		},
	)
}

// An object reference to a recipient
type MessageUnarchiveResponseActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                   `json:"collection,required"`
	JSON       messageUnarchiveResponseActorsObjectJSON `json:"-"`
}

// messageUnarchiveResponseActorsObjectJSON contains the JSON metadata for the
// struct [MessageUnarchiveResponseActorsObject]
type messageUnarchiveResponseActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageUnarchiveResponseActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageUnarchiveResponseActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageUnarchiveResponseActorsObject) ImplementsMessageUnarchiveResponseActorsUnion() {}

type MessageUnarchiveResponseEngagementStatus string

const (
	MessageUnarchiveResponseEngagementStatusSeen        MessageUnarchiveResponseEngagementStatus = "seen"
	MessageUnarchiveResponseEngagementStatusRead        MessageUnarchiveResponseEngagementStatus = "read"
	MessageUnarchiveResponseEngagementStatusInteracted  MessageUnarchiveResponseEngagementStatus = "interacted"
	MessageUnarchiveResponseEngagementStatusLinkClicked MessageUnarchiveResponseEngagementStatus = "link_clicked"
	MessageUnarchiveResponseEngagementStatusArchived    MessageUnarchiveResponseEngagementStatus = "archived"
)

func (r MessageUnarchiveResponseEngagementStatus) IsKnown() bool {
	switch r {
	case MessageUnarchiveResponseEngagementStatusSeen, MessageUnarchiveResponseEngagementStatusRead, MessageUnarchiveResponseEngagementStatusInteracted, MessageUnarchiveResponseEngagementStatusLinkClicked, MessageUnarchiveResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [MessageUnarchiveResponseRecipientObject].
type MessageUnarchiveResponseRecipientUnion interface {
	ImplementsMessageUnarchiveResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageUnarchiveResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageUnarchiveResponseRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type MessageUnarchiveResponseRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                      `json:"collection,required"`
	JSON       messageUnarchiveResponseRecipientObjectJSON `json:"-"`
}

// messageUnarchiveResponseRecipientObjectJSON contains the JSON metadata for the
// struct [MessageUnarchiveResponseRecipientObject]
type messageUnarchiveResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageUnarchiveResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageUnarchiveResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r MessageUnarchiveResponseRecipientObject) ImplementsMessageUnarchiveResponseRecipientUnion() {}

// Source information
type MessageUnarchiveResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                             `json:"version_id,required" format:"uuid"`
	JSON      messageUnarchiveResponseSourceJSON `json:"-"`
}

// messageUnarchiveResponseSourceJSON contains the JSON metadata for the struct
// [MessageUnarchiveResponseSource]
type messageUnarchiveResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageUnarchiveResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageUnarchiveResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type MessageUnarchiveResponseStatus string

const (
	MessageUnarchiveResponseStatusQueued            MessageUnarchiveResponseStatus = "queued"
	MessageUnarchiveResponseStatusSent              MessageUnarchiveResponseStatus = "sent"
	MessageUnarchiveResponseStatusDelivered         MessageUnarchiveResponseStatus = "delivered"
	MessageUnarchiveResponseStatusDeliveryAttempted MessageUnarchiveResponseStatus = "delivery_attempted"
	MessageUnarchiveResponseStatusUndelivered       MessageUnarchiveResponseStatus = "undelivered"
	MessageUnarchiveResponseStatusNotSent           MessageUnarchiveResponseStatus = "not_sent"
	MessageUnarchiveResponseStatusBounced           MessageUnarchiveResponseStatus = "bounced"
)

func (r MessageUnarchiveResponseStatus) IsKnown() bool {
	switch r {
	case MessageUnarchiveResponseStatusQueued, MessageUnarchiveResponseStatusSent, MessageUnarchiveResponseStatusDelivered, MessageUnarchiveResponseStatusDeliveryAttempted, MessageUnarchiveResponseStatusUndelivered, MessageUnarchiveResponseStatusNotSent, MessageUnarchiveResponseStatusBounced:
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/apiquery"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
	"github.com/tidwall/gjson"
)

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
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/archived"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the contents of multiple messages in a single request.
func (r *MessageBatchService) GetContent(ctx context.Context, query MessageBatchGetContentParams, opts ...option.RequestOption) (res *[]MessageBatchGetContentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/content"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Marks the given messages as interacted with by the user. This can include any
// user action on the message, with optional metadata about the specific
// interaction. Cannot include more than 5 key-value pairs, must not contain nested
// data. Read more about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsInteracted(ctx context.Context, body MessageBatchMarkAsInteractedParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/interacted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Marks the given messages as `read`. Read more about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsRead(ctx context.Context, body MessageBatchMarkAsReadParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Marks the given messages as `seen`. This indicates that the user has viewed the
// message in their feed or inbox. Read more about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsSeen(ctx context.Context, body MessageBatchMarkAsSeenParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/seen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Marks the given messages as `unread`. This reverses the `read` state. Read more
// about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsUnread(ctx context.Context, body MessageBatchMarkAsUnreadParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/unread"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Marks the given messages as `unseen`. This reverses the `seen` state. Read more
// about message engagement statuses
// [here](/send-notifications/message-statuses#engagement-status).
func (r *MessageBatchService) MarkAsUnseen(ctx context.Context, body MessageBatchMarkAsUnseenParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/unseen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Marks the given messages as unarchived. This reverses the `archived` state.
// Archived messages are hidden from the default message list in the feed but can
// still be accessed and unarchived later.
func (r *MessageBatchService) Unarchive(ctx context.Context, body MessageBatchUnarchiveParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/unarchived"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The content of a message.
type MessageBatchGetContentResponse struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// Content data specific to the channel type.
	Data MessageBatchGetContentResponseData `json:"data,required"`
	// Timestamp when the message content was created.
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// The unique identifier for the message content.
	MessageID string                             `json:"message_id,required"`
	JSON      messageBatchGetContentResponseJSON `json:"-"`
}

// messageBatchGetContentResponseJSON contains the JSON metadata for the struct
// [MessageBatchGetContentResponse]
type messageBatchGetContentResponseJSON struct {
	Typename    apijson.Field
	Data        apijson.Field
	InsertedAt  apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseJSON) RawJSON() string {
	return r.raw
}

// Content data specific to the channel type.
type MessageBatchGetContentResponseData struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The device token to send the push notification to.
	Token string `json:"token"`
	// The BCC email addresses.
	Bcc string `json:"bcc,nullable"`
	// This field can have the runtime type of
	// [[]MessageBatchGetContentResponseDataMessageInAppFeedContentBlock].
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
	// [MessageBatchGetContentResponseDataMessageChatContentTemplate].
	Template interface{} `json:"template"`
	// The text body of the email message.
	TextBody string `json:"text_body"`
	// The title of the push notification.
	Title string `json:"title"`
	// The recipient's email address.
	To    string                                 `json:"to"`
	JSON  messageBatchGetContentResponseDataJSON `json:"-"`
	union MessageBatchGetContentResponseDataUnion
}

// messageBatchGetContentResponseDataJSON contains the JSON metadata for the struct
// [MessageBatchGetContentResponseData]
type messageBatchGetContentResponseDataJSON struct {
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

func (r messageBatchGetContentResponseDataJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchGetContentResponseData) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchGetContentResponseData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageBatchGetContentResponseDataUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchGetContentResponseDataMessageEmailContent],
// [MessageBatchGetContentResponseDataMessageSMSContent],
// [MessageBatchGetContentResponseDataMessagePushContent],
// [MessageBatchGetContentResponseDataMessageChatContent],
// [MessageBatchGetContentResponseDataMessageInAppFeedContent].
func (r MessageBatchGetContentResponseData) AsUnion() MessageBatchGetContentResponseDataUnion {
	return r.union
}

// Content data specific to the channel type.
//
// Union satisfied by [MessageBatchGetContentResponseDataMessageEmailContent],
// [MessageBatchGetContentResponseDataMessageSMSContent],
// [MessageBatchGetContentResponseDataMessagePushContent],
// [MessageBatchGetContentResponseDataMessageChatContent] or
// [MessageBatchGetContentResponseDataMessageInAppFeedContent].
type MessageBatchGetContentResponseDataUnion interface {
	implementsMessageBatchGetContentResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchGetContentResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageBatchGetContentResponseDataMessageEmailContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageBatchGetContentResponseDataMessageSMSContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageBatchGetContentResponseDataMessagePushContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageBatchGetContentResponseDataMessageChatContent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageBatchGetContentResponseDataMessageInAppFeedContent{}),
		},
	)
}

// The content of an email message.
type MessageBatchGetContentResponseDataMessageEmailContent struct {
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
	ReplyTo string                                                    `json:"reply_to,nullable"`
	JSON    messageBatchGetContentResponseDataMessageEmailContentJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessageEmailContentJSON contains the JSON
// metadata for the struct [MessageBatchGetContentResponseDataMessageEmailContent]
type messageBatchGetContentResponseDataMessageEmailContentJSON struct {
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

func (r *MessageBatchGetContentResponseDataMessageEmailContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageEmailContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchGetContentResponseDataMessageEmailContent) implementsMessageBatchGetContentResponseData() {
}

// The content of an SMS message.
type MessageBatchGetContentResponseDataMessageSMSContent struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The content body of the SMS message.
	Body string `json:"body,required"`
	// The phone number the SMS was sent to.
	To   string                                                  `json:"to,required"`
	JSON messageBatchGetContentResponseDataMessageSMSContentJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessageSMSContentJSON contains the JSON
// metadata for the struct [MessageBatchGetContentResponseDataMessageSMSContent]
type messageBatchGetContentResponseDataMessageSMSContentJSON struct {
	Typename    apijson.Field
	Body        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageSMSContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageSMSContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchGetContentResponseDataMessageSMSContent) implementsMessageBatchGetContentResponseData() {
}

// Push channel data.
type MessageBatchGetContentResponseDataMessagePushContent struct {
	// The device token to send the push notification to.
	Token string `json:"token,required"`
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The content body of the push notification.
	Body string `json:"body,required"`
	// The title of the push notification.
	Title string `json:"title,required"`
	// Additional data payload for the push notification.
	Data map[string]interface{}                                   `json:"data,nullable"`
	JSON messageBatchGetContentResponseDataMessagePushContentJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessagePushContentJSON contains the JSON
// metadata for the struct [MessageBatchGetContentResponseDataMessagePushContent]
type messageBatchGetContentResponseDataMessagePushContentJSON struct {
	Token       apijson.Field
	Typename    apijson.Field
	Body        apijson.Field
	Title       apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessagePushContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessagePushContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchGetContentResponseDataMessagePushContent) implementsMessageBatchGetContentResponseData() {
}

// The content of a chat message.
type MessageBatchGetContentResponseDataMessageChatContent struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The channel data connection from the recipient to the underlying provider.
	Connection map[string]interface{} `json:"connection,required"`
	// The template structure for the chat message.
	Template MessageBatchGetContentResponseDataMessageChatContentTemplate `json:"template,required"`
	// Additional metadata associated with the chat message.
	Metadata map[string]interface{}                                   `json:"metadata,nullable"`
	JSON     messageBatchGetContentResponseDataMessageChatContentJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessageChatContentJSON contains the JSON
// metadata for the struct [MessageBatchGetContentResponseDataMessageChatContent]
type messageBatchGetContentResponseDataMessageChatContentJSON struct {
	Typename    apijson.Field
	Connection  apijson.Field
	Template    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageChatContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageChatContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchGetContentResponseDataMessageChatContent) implementsMessageBatchGetContentResponseData() {
}

// The template structure for the chat message.
type MessageBatchGetContentResponseDataMessageChatContentTemplate struct {
	// The blocks of the message in a chat.
	Blocks []MessageBatchGetContentResponseDataMessageChatContentTemplateBlock `json:"blocks,nullable"`
	// The JSON content of the message.
	JsonContent map[string]interface{} `json:"json_content,nullable"`
	// The summary of the chat message.
	Summary string                                                           `json:"summary,nullable"`
	JSON    messageBatchGetContentResponseDataMessageChatContentTemplateJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessageChatContentTemplateJSON contains the
// JSON metadata for the struct
// [MessageBatchGetContentResponseDataMessageChatContentTemplate]
type messageBatchGetContentResponseDataMessageChatContentTemplateJSON struct {
	Blocks      apijson.Field
	JsonContent apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageChatContentTemplate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageChatContentTemplateJSON) RawJSON() string {
	return r.raw
}

// A block in a message in a chat.
type MessageBatchGetContentResponseDataMessageChatContentTemplateBlock struct {
	// The actual content of the block.
	Content string `json:"content,required"`
	// The name of the block for identification.
	Name string `json:"name,required"`
	// The type of block in a message in a chat (text or markdown).
	Type MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksType `json:"type,required"`
	JSON messageBatchGetContentResponseDataMessageChatContentTemplateBlockJSON  `json:"-"`
}

// messageBatchGetContentResponseDataMessageChatContentTemplateBlockJSON contains
// the JSON metadata for the struct
// [MessageBatchGetContentResponseDataMessageChatContentTemplateBlock]
type messageBatchGetContentResponseDataMessageChatContentTemplateBlockJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageChatContentTemplateBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageChatContentTemplateBlockJSON) RawJSON() string {
	return r.raw
}

// The type of block in a message in a chat (text or markdown).
type MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksType string

const (
	MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksTypeText     MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksType = "text"
	MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksTypeMarkdown MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksType = "markdown"
)

func (r MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksType) IsKnown() bool {
	switch r {
	case MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksTypeText, MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksTypeMarkdown:
		return true
	}
	return false
}

// The content of an in-app feed message.
type MessageBatchGetContentResponseDataMessageInAppFeedContent struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The blocks of the message in an app feed.
	Blocks []MessageBatchGetContentResponseDataMessageInAppFeedContentBlock `json:"blocks,required"`
	JSON   messageBatchGetContentResponseDataMessageInAppFeedContentJSON    `json:"-"`
}

// messageBatchGetContentResponseDataMessageInAppFeedContentJSON contains the JSON
// metadata for the struct
// [MessageBatchGetContentResponseDataMessageInAppFeedContent]
type messageBatchGetContentResponseDataMessageInAppFeedContentJSON struct {
	Typename    apijson.Field
	Blocks      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageInAppFeedContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageInAppFeedContentJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchGetContentResponseDataMessageInAppFeedContent) implementsMessageBatchGetContentResponseData() {
}

// A block in a message in an app feed.
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlock struct {
	// The name of the block in a message in an app feed.
	Name string `json:"name,required"`
	// The type of block in a message in an app feed.
	Type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksType `json:"type,required"`
	// This field can have the runtime type of
	// [[]MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton].
	Buttons interface{} `json:"buttons"`
	// The content of the block in a message in an app feed.
	Content string `json:"content"`
	// The rendered HTML version of the content.
	Rendered string                                                             `json:"rendered"`
	JSON     messageBatchGetContentResponseDataMessageInAppFeedContentBlockJSON `json:"-"`
	union    MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksUnion
}

// messageBatchGetContentResponseDataMessageInAppFeedContentBlockJSON contains the
// JSON metadata for the struct
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlock]
type messageBatchGetContentResponseDataMessageInAppFeedContentBlockJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Buttons     apijson.Field
	Content     apijson.Field
	Rendered    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageBatchGetContentResponseDataMessageInAppFeedContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchGetContentResponseDataMessageInAppFeedContentBlock) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchGetContentResponseDataMessageInAppFeedContentBlock{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock],
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock].
func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlock) AsUnion() MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksUnion {
	return r.union
}

// A block in a message in an app feed.
//
// Union satisfied by
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock]
// or
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock].
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksUnion interface {
	implementsMessageBatchGetContentResponseDataMessageInAppFeedContentBlock()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock{}),
		},
	)
}

// A block in a message in an app feed.
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock struct {
	// The content of the block in a message in an app feed.
	Content string `json:"content,required"`
	// The name of the block in a message in an app feed.
	Name string `json:"name,required"`
	// The rendered HTML version of the content.
	Rendered string `json:"rendered,required"`
	// The type of block in a message in an app feed.
	Type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType `json:"type,required"`
	JSON messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock]
type messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Rendered    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlock) implementsMessageBatchGetContentResponseDataMessageInAppFeedContentBlock() {
}

// The type of block in a message in an app feed.
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType string

const (
	MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockTypeMarkdown MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType = "markdown"
	MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockTypeText     MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType = "text"
)

func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockType) IsKnown() bool {
	switch r {
	case MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockTypeMarkdown, MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedContentBlockTypeText:
		return true
	}
	return false
}

// A button set block in a message in an app feed.
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock struct {
	// A list of buttons in an in app feed message.
	Buttons []MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton `json:"buttons,required"`
	// The name of the button set in a message in an app feed.
	Name string `json:"name,required"`
	// The type of block in a message in an app feed.
	Type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockType `json:"type,required"`
	JSON messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock]
type messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockJSON struct {
	Buttons     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlock) implementsMessageBatchGetContentResponseDataMessageInAppFeedContentBlock() {
}

// A button in an in app feed message.
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton struct {
	// The action to take when the button is clicked.
	Action string `json:"action,required"`
	// The label of the button.
	Label string `json:"label,required"`
	// The name of the button.
	Name string                                                                                                  `json:"name,required"`
	JSON messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButtonJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButtonJSON
// contains the JSON metadata for the struct
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton]
type messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButtonJSON struct {
	Action      apijson.Field
	Label       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButton) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockButtonJSON) RawJSON() string {
	return r.raw
}

// The type of block in a message in an app feed.
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockType string

const (
	MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockTypeButtonSet MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockType = "button_set"
)

func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockType) IsKnown() bool {
	switch r {
	case MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksMessageInAppFeedButtonSetBlockTypeButtonSet:
		return true
	}
	return false
}

// The type of block in a message in an app feed.
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksType string

const (
	MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksTypeMarkdown  MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksType = "markdown"
	MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksTypeText      MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksType = "text"
	MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksTypeButtonSet MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksType = "button_set"
)

func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksType) IsKnown() bool {
	switch r {
	case MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksTypeMarkdown, MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksTypeText, MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksTypeButtonSet:
		return true
	}
	return false
}

type MessageBatchArchiveParams struct {
	// The message IDs to update the status of.
	MessageIDs param.Field[[]string] `json:"message_ids,required"`
}

func (r MessageBatchArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchGetContentParams struct {
	// The IDs of the messages to fetch contents of.
	MessageIDs param.Field[[]string] `query:"message_ids,required"`
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
	MessageIDs param.Field[[]string] `json:"message_ids,required"`
	// Metadata about the interaction.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r MessageBatchMarkAsInteractedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchMarkAsReadParams struct {
	// The message IDs to update the status of.
	MessageIDs param.Field[[]string] `json:"message_ids,required"`
}

func (r MessageBatchMarkAsReadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchMarkAsSeenParams struct {
	// The message IDs to update the status of.
	MessageIDs param.Field[[]string] `json:"message_ids,required"`
}

func (r MessageBatchMarkAsSeenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchMarkAsUnreadParams struct {
	// The message IDs to update the status of.
	MessageIDs param.Field[[]string] `json:"message_ids,required"`
}

func (r MessageBatchMarkAsUnreadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchMarkAsUnseenParams struct {
	// The message IDs to update the status of.
	MessageIDs param.Field[[]string] `json:"message_ids,required"`
}

func (r MessageBatchMarkAsUnseenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchUnarchiveParams struct {
	// The message IDs to update the status of.
	MessageIDs param.Field[[]string] `json:"message_ids,required"`
}

func (r MessageBatchUnarchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

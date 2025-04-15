// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
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

// Marks one or more messages as archived
func (r *MessageBatchService) Archive(ctx context.Context, body MessageBatchArchiveParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/archived"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the contents of multiple messages
func (r *MessageBatchService) GetContent(ctx context.Context, query MessageBatchGetContentParams, opts ...option.RequestOption) (res *[]MessageBatchGetContentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/content"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Marks one or more messages as interacted
func (r *MessageBatchService) MarkAsInteracted(ctx context.Context, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/interacted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Marks one or more messages as read
func (r *MessageBatchService) MarkAsRead(ctx context.Context, body MessageBatchMarkAsReadParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Marks one or more messages as seen
func (r *MessageBatchService) MarkAsSeen(ctx context.Context, body MessageBatchMarkAsSeenParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/seen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Marks one or more messages as unread
func (r *MessageBatchService) MarkAsUnread(ctx context.Context, body MessageBatchMarkAsUnreadParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/unread"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Marks one or more messages as unseen
func (r *MessageBatchService) MarkAsUnseen(ctx context.Context, body MessageBatchMarkAsUnseenParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/unseen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Marks one or more messages as unarchived
func (r *MessageBatchService) Unarchive(ctx context.Context, body MessageBatchUnarchiveParams, opts ...option.RequestOption) (res *[]Message, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/batch/unarchived"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The contents of a message
type MessageBatchGetContentResponse struct {
	Typename string `json:"__typename,required"`
	// The contents of an email message
	Data       MessageBatchGetContentResponseData `json:"data,required"`
	InsertedAt time.Time                          `json:"inserted_at,required" format:"date-time"`
	MessageID  string                             `json:"message_id,required"`
	JSON       messageBatchGetContentResponseJSON `json:"-"`
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

// The contents of an email message
type MessageBatchGetContentResponseData struct {
	Typename string `json:"__typename,required"`
	Token    string `json:"token"`
	Bcc      string `json:"bcc,nullable"`
	// This field can have the runtime type of
	// [[]MessageBatchGetContentResponseDataMessageInAppFeedContentBlock].
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
	// [MessageBatchGetContentResponseDataMessageChatContentTemplate].
	Template interface{}                            `json:"template"`
	TextBody string                                 `json:"text_body"`
	Title    string                                 `json:"title"`
	To       string                                 `json:"to"`
	JSON     messageBatchGetContentResponseDataJSON `json:"-"`
	union    MessageBatchGetContentResponseDataUnion
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

// The contents of an email message
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

// The contents of an email message
type MessageBatchGetContentResponseDataMessageEmailContent struct {
	Typename    string                                                    `json:"__typename,required"`
	From        string                                                    `json:"from,required"`
	HTMLBody    string                                                    `json:"html_body,required"`
	SubjectLine string                                                    `json:"subject_line,required"`
	TextBody    string                                                    `json:"text_body,required"`
	To          string                                                    `json:"to,required"`
	Bcc         string                                                    `json:"bcc,nullable"`
	Cc          string                                                    `json:"cc,nullable"`
	ReplyTo     string                                                    `json:"reply_to,nullable"`
	JSON        messageBatchGetContentResponseDataMessageEmailContentJSON `json:"-"`
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

// The contents of an SMS message
type MessageBatchGetContentResponseDataMessageSMSContent struct {
	Typename string                                                  `json:"__typename,required"`
	Body     string                                                  `json:"body,required"`
	To       string                                                  `json:"to,required"`
	JSON     messageBatchGetContentResponseDataMessageSMSContentJSON `json:"-"`
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

// The contents of a push message
type MessageBatchGetContentResponseDataMessagePushContent struct {
	Token    string                                                   `json:"token,required"`
	Typename string                                                   `json:"__typename,required"`
	Body     string                                                   `json:"body,required"`
	Title    string                                                   `json:"title,required"`
	Data     interface{}                                              `json:"data,nullable"`
	JSON     messageBatchGetContentResponseDataMessagePushContentJSON `json:"-"`
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

// The contents of a chat message
type MessageBatchGetContentResponseDataMessageChatContent struct {
	Typename string `json:"__typename,required"`
	// The channel data connection from the recipient to the underlying provider
	Connection interface{}                                                  `json:"connection,required"`
	Template   MessageBatchGetContentResponseDataMessageChatContentTemplate `json:"template,required"`
	Metadata   interface{}                                                  `json:"metadata,nullable"`
	JSON       messageBatchGetContentResponseDataMessageChatContentJSON     `json:"-"`
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

type MessageBatchGetContentResponseDataMessageChatContentTemplate struct {
	// The structured blocks of the message
	Blocks []MessageBatchGetContentResponseDataMessageChatContentTemplateBlock `json:"blocks,nullable"`
	// The JSON content of the message
	JsonContent interface{}                                                      `json:"json_content,nullable"`
	Summary     string                                                           `json:"summary,nullable"`
	JSON        messageBatchGetContentResponseDataMessageChatContentTemplateJSON `json:"-"`
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

// A block in a chat message
type MessageBatchGetContentResponseDataMessageChatContentTemplateBlock struct {
	Content string                                                                 `json:"content,required"`
	Name    string                                                                 `json:"name,required"`
	Type    MessageBatchGetContentResponseDataMessageChatContentTemplateBlocksType `json:"type,required"`
	JSON    messageBatchGetContentResponseDataMessageChatContentTemplateBlockJSON  `json:"-"`
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

// The contents of a message in an app feed
type MessageBatchGetContentResponseDataMessageInAppFeedContent struct {
	Typename string `json:"__typename,required"`
	// The blocks of the message
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

// A content (text or markdown) block in a message in an app feed
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlock struct {
	Name string                                                              `json:"name,required"`
	Type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksType `json:"type,required"`
	// This field can have the runtime type of
	// [[]MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton].
	Buttons  interface{}                                                        `json:"buttons"`
	Content  string                                                             `json:"content"`
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
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlock],
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock].
func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlock) AsUnion() MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksUnion {
	return r.union
}

// A content (text or markdown) block in a message in an app feed
//
// Union satisfied by
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlock] or
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock].
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksUnion interface {
	implementsMessageBatchGetContentResponseDataMessageInAppFeedContentBlock()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlock{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock{}),
		},
	)
}

// A content (text or markdown) block in a message in an app feed
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlock struct {
	Content  string                                                                          `json:"content,required"`
	Name     string                                                                          `json:"name,required"`
	Rendered string                                                                          `json:"rendered,required"`
	Type     MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType `json:"type,required"`
	JSON     messageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlock]
type messageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Rendered    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlock) implementsMessageBatchGetContentResponseDataMessageInAppFeedContentBlock() {
}

type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType string

const (
	MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockTypeMarkdown MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType = "markdown"
	MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockTypeText     MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType = "text"
)

func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockType) IsKnown() bool {
	switch r {
	case MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockTypeMarkdown, MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksContentBlockTypeText:
		return true
	}
	return false
}

// A set of buttons in a message in an app feed
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock struct {
	Buttons []MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton `json:"buttons,required"`
	Name    string                                                                                `json:"name,required"`
	Type    MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockType     `json:"type,required"`
	JSON    messageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockJSON     `json:"-"`
}

// messageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock]
type messageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockJSON struct {
	Buttons     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlock) implementsMessageBatchGetContentResponseDataMessageInAppFeedContentBlock() {
}

// A button in a set of buttons
type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton struct {
	Action string                                                                                  `json:"action,required"`
	Label  string                                                                                  `json:"label,required"`
	Name   string                                                                                  `json:"name,required"`
	JSON   messageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButtonJSON `json:"-"`
}

// messageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButtonJSON
// contains the JSON metadata for the struct
// [MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton]
type messageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButtonJSON struct {
	Action      apijson.Field
	Label       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButton) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockButtonJSON) RawJSON() string {
	return r.raw
}

type MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockType string

const (
	MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockTypeButtonSet MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockType = "button_set"
)

func (r MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockType) IsKnown() bool {
	switch r {
	case MessageBatchGetContentResponseDataMessageInAppFeedContentBlocksButtonSetBlockTypeButtonSet:
		return true
	}
	return false
}

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
	MessageIDs param.Field[[]string] `json:"message_ids,required" format:"uuid"`
}

type MessageBatchGetContentParams struct {
	// The IDs of the messages to fetch contents of
	MessageIDs param.Field[[]interface{}] `query:"message_ids,required"`
}

// URLQuery serializes [MessageBatchGetContentParams]'s query parameters as
// `url.Values`.
func (r MessageBatchGetContentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageBatchMarkAsReadParams struct {
	MessageIDs param.Field[[]string] `json:"message_ids,required" format:"uuid"`
}

type MessageBatchMarkAsSeenParams struct {
	MessageIDs param.Field[[]string] `json:"message_ids,required" format:"uuid"`
}

type MessageBatchMarkAsUnreadParams struct {
	MessageIDs param.Field[[]string] `json:"message_ids,required" format:"uuid"`
}

type MessageBatchMarkAsUnseenParams struct {
	MessageIDs param.Field[[]string] `json:"message_ids,required" format:"uuid"`
}

type MessageBatchUnarchiveParams struct {
	MessageIDs param.Field[[]string] `json:"message_ids,required" format:"uuid"`
}

package knock

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// MessagesService is an interface for communicating with the Knock
// Messages API endpoints.
type MessagesService interface {
	List(context.Context, *ListMessagesRequest) ([]*Message, *PageInfo, error)

	Get(context.Context, *GetMessageRequest) (*Message, error)
	GetEvents(context.Context, *GetMessageEventsRequest) ([]*MessageEvent, *PageInfo, error)
	GetActivities(context.Context, *GetMessageActivitiesRequest) ([]*MessageActivity, *PageInfo, error)
	GetContent(context.Context, *GetMessageContentRequest) (*MessageContent, error)

	SetStatus(context.Context, *SetStatusRequest) (*Message, error)
	DeleteStatus(context.Context, *DeleteStatusRequest) (*Message, error)

	BatchSetStatus(context.Context, *BatchSetStatusRequest) ([]*Message, error)
	BulkChangeChannelStatus(context.Context, *BulkChangeChannelStatusRequest) (*BulkOperation, error)
}

type messagesService struct {
	client *Client
}

var _ MessagesService = &messagesService{}

func NewMessagesService(client *Client) *messagesService {
	return &messagesService{
		client: client,
	}
}

type EngagementStatus string

// User message query statuses
const (
	Queued      EngagementStatus = "queued"
	Sent        EngagementStatus = "sent"
	Delivered   EngagementStatus = "delivered"
	Undelivered EngagementStatus = "undelivered"
)

type MessageStatus string

const (
	NotSent    MessageStatus = "seen"
	Read       MessageStatus = "read"
	Archived   MessageStatus = "archived"
	Unarchived MessageStatus = "unarchived"
	Unread     MessageStatus = "unread"
)

type Message struct {
	Cursor     string                 `json:"__cursor"`
	ID         string                 `json:"id"`
	ChannelID  string                 `json:"channel_id"`
	Recipient  string                 `json:"recipient"`
	Workflow   string                 `json:"workflow"`
	Tenant     string                 `json:"tenant"`
	Status     EngagementStatus       `json:"status"`
	ReadAt     time.Time              `json:"read_at"`
	SeenAt     time.Time              `json:"seen_at"`
	ArchivedAt time.Time              `json:"archived_at"`
	InsertedAt time.Time              `json:"inserted_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	Source     *NotificationSource    `json:"source"`
	Data       map[string]interface{} `json:"data"`
}

type NotificationSource struct {
	Key       string `json:"key"`
	VersionID string `json:"version_id"`
}

type MessageEvent struct {
	Cursor        string                 `json:"__cursor"`
	ID            string                 `json:"id"`
	EnvironmentID string                 `json:"environment_id"`
	Recipient     string                 `json:"recipient"`
	Data          map[string]interface{} `json:"data"`
}

type MessageActivity struct {
	Cursor    string                 `json:"__cursor"`
	ID        string                 `json:"id"`
	Data      map[string]interface{} `json:"data"`
	Actor     *User                  `json:"actor"`
	Recipient *User                  `json:"recipient"`
}

type MessageContent struct {
	ID         string                 `json:"id"`
	Data       map[string]interface{} `json:"data"`
	InsertedAt time.Time              `json:"inserted_at"`
}

// requests and responses
type ListMessagesRequest struct {
	PageSize  int                `url:"page_size,omitempty"`
	Cursor    string             `url:"page_size,omitempty"`
	Before    string             `url:"before,omitempty"`
	After     string             `url:"after,omitempty"`
	Source    string             `url:"source,omitempty"`
	Tenant    string             `url:"tenant,omitempty"`
	Status    []EngagementStatus `url:"status,omitempty"`
	ChannelID string             `url:"channel_id,omitempty"`
}
type ListMessagesResponse struct {
	Items    []*Message `json:"items"`
	PageInfo *PageInfo  `json:"page_info"`
}

type GetMessageRequest struct {
	ID string
}
type GetMessageResponse struct {
	Message *Message
}

type GetMessageEventsRequest = GetMessageRequest
type GetMessageEventsResponse struct {
	Items    []*MessageEvent `json:"items"`
	PageInfo *PageInfo       `json:"page_info"`
}

type GetMessageActivitiesRequest = GetMessageRequest
type GetMessageActivitiesResponse struct {
	Activities []*MessageActivity `json:"items"`
	PageInfo   *PageInfo          `json:"page_info"`
}

type GetMessageContentRequest = GetMessageRequest
type GetMessageContentResponse struct {
	Content *MessageContent
}

type SetStatusRequest struct {
	ID     string
	Status MessageStatus
}
type SetStatusResponse = GetMessageResponse

type DeleteStatusRequest = SetStatusRequest
type DeleteStatusResponse = GetMessageResponse

type BatchSetStatusRequest struct {
	Status     MessageStatus `json:"-"`
	MessageIDs []string      `json:"message_ids"`
}
type BatchSetStatusResponse struct {
	Messages []*Message
}

type BulkChangeChannelStatusRequest struct {
	ChannelID string        `json:"-"`
	Status    MessageStatus `json:"-"`

	UserIDs          []string         `json:"user_ids,omitempty"`
	OlderThan        *time.Time       `json:"older_than,omitempty"`
	NewerThan        *time.Time       `json:"newer_than,omitempty"`
	DeliveryStatus   EngagementStatus `json:"delivery_status,omitempty"`
	EngagementStatus string           `json:"engagement_status,omitempty"`
}
type BulkChangeChannelStatusResponse struct {
	BulkOperation *BulkOperation
}

const messagesAPIBasePath = "/v1/messages"

func messagesAPIPath(messageID string) string {
	return fmt.Sprintf("%s/%s", messagesAPIBasePath, messageID)
}

func (ms *messagesService) List(ctx context.Context, listReq *ListMessagesRequest) ([]*Message, *PageInfo, error) {

	queryString, err := query.Values(listReq)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error parsing request to list messages")
	}
	path := fmt.Sprintf("%s?%s", messagesAPIBasePath, queryString.Encode())

	req, err := ms.client.newRequest(http.MethodGet, path, listReq)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request to list messages")
	}
	listRes := &ListMessagesResponse{}

	_, err = ms.client.do(ctx, req, listRes)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request to list messages")
	}

	return listRes.Items, listRes.PageInfo, nil
}

func (ms *messagesService) Get(ctx context.Context, getReq *GetMessageRequest) (*Message, error) {

	path := messagesAPIPath(getReq.ID)

	req, err := ms.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to list messages")
	}
	getRes := &GetMessageResponse{Message: &Message{}}
	_, err = ms.client.do(ctx, req, getRes.Message)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to list messages")
	}

	return getRes.Message, nil
}

func (ms *messagesService) GetEvents(ctx context.Context, getEventsReq *GetMessageEventsRequest) ([]*MessageEvent, *PageInfo, error) {

	path := fmt.Sprintf("%s/events", messagesAPIPath(getEventsReq.ID))

	req, err := ms.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request to list message events")
	}
	getMessageEventsResponse := &GetMessageEventsResponse{}
	_, err = ms.client.do(ctx, req, getMessageEventsResponse)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request to list message events")
	}

	return getMessageEventsResponse.Items, getMessageEventsResponse.PageInfo, nil
}

func (ms *messagesService) GetActivities(ctx context.Context, getActivitiesReq *GetMessageActivitiesRequest) ([]*MessageActivity, *PageInfo, error) {

	path := fmt.Sprintf("%s/activities", messagesAPIPath(getActivitiesReq.ID))

	req, err := ms.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request to list message activities")
	}
	getMessageActivitiesResponse := &GetMessageActivitiesResponse{}
	_, err = ms.client.do(ctx, req, getMessageActivitiesResponse)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request to list message activities")
	}

	return getMessageActivitiesResponse.Activities, getMessageActivitiesResponse.PageInfo, nil
}

func (ms *messagesService) GetContent(ctx context.Context, getContentReq *GetMessageContentRequest) (*MessageContent, error) {

	path := fmt.Sprintf("%s/activities", messagesAPIPath(getContentReq.ID))

	req, err := ms.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to list message activities")
	}
	getMessageContentResponse := &GetMessageContentResponse{Content: &MessageContent{}}
	_, err = ms.client.do(ctx, req, getMessageContentResponse.Content)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to list message activities")
	}

	return getMessageContentResponse.Content, nil
}

func (ms *messagesService) SetStatus(ctx context.Context, setStatusReq *SetStatusRequest) (*Message, error) {

	path := fmt.Sprintf("%s/%s", messagesAPIPath(setStatusReq.ID), setStatusReq.Status)

	req, err := ms.client.newRequest(http.MethodPut, path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set message status")
	}
	setMessageStatusResponse := &SetStatusResponse{Message: &Message{}}
	_, err = ms.client.do(ctx, req, setMessageStatusResponse.Message)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to set message status")
	}

	return setMessageStatusResponse.Message, nil
}

func (ms *messagesService) DeleteStatus(ctx context.Context, deleteStatusReq *DeleteStatusRequest) (*Message, error) {

	path := fmt.Sprintf("%s/%s", messagesAPIPath(deleteStatusReq.ID), deleteStatusReq.Status)

	req, err := ms.client.newRequest(http.MethodDelete, path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to delete message status")
	}
	deleteMessageStatusResponse := &SetStatusResponse{Message: &Message{}}
	_, err = ms.client.do(ctx, req, deleteMessageStatusResponse.Message)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to delete message status")
	}

	return deleteMessageStatusResponse.Message, nil
}

func (ms *messagesService) BatchSetStatus(ctx context.Context, batchSetStatus *BatchSetStatusRequest) ([]*Message, error) {

	path := fmt.Sprintf("%s/%s", messagesAPIPath("batch"), batchSetStatus.Status)

	req, err := ms.client.newRequest(http.MethodPost, path, batchSetStatus)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set message status")
	}

	batchSetStatusRes := BatchSetStatusResponse{Messages: []*Message{}}

	_, err = ms.client.do(ctx, req, &batchSetStatusRes.Messages)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to set message status")
	}

	return batchSetStatusRes.Messages, nil
}

func (ms *messagesService) BulkChangeChannelStatus(ctx context.Context, bulkStatusReq *BulkChangeChannelStatusRequest) (*BulkOperation, error) {

	path := fmt.Sprintf("v1/channels/%s/messages/bulk/%s", bulkStatusReq.ChannelID, bulkStatusReq.Status)

	req, err := ms.client.newRequest(http.MethodPost, path, bulkStatusReq)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set message status")
	}

	batchSetStatusRes := BulkChangeChannelStatusResponse{BulkOperation: &BulkOperation{}}

	_, err = ms.client.do(ctx, req, &batchSetStatusRes.BulkOperation)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to set message status")
	}

	return batchSetStatusRes.BulkOperation, nil
}

package knock

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

type MessagesService interface {
	List(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error)
	Get(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
	GetEvents(context.Context, *GetMessageEventsRequest) (*GetMessageEventsResponse, error)
	GetActivities(context.Context, *GetMessageActivitiesRequest) (*GetMessageActivitiesResponse, error)
	GetContent(context.Context, *GetMessageContentRequest) (*GetMessageContentResponse, error)
	SetStatus(context.Context, *SetStatusRequest) (*SetStatusResponse, error)
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

const messagesAPIBasePath = "/v1/messages"

func messagesAPIPath(messageID string) string {
	return fmt.Sprintf("%s/%s", messagesAPIBasePath, messageID)
}

type UserMessageStatus string

// User message query statuses
const (
	Queued      UserMessageStatus = "queued"
	Sent        UserMessageStatus = "sent"
	Delivered   UserMessageStatus = "delivered"
	Undelivered UserMessageStatus = "undelivered"
)

type MessageStatus string

const (
	NotSent  MessageStatus = "not_sent"
	Read     MessageStatus = "read"
	Archived MessageStatus = "archived"
)

type NotificationSource struct {
	Key       string `json:"key"`
	VersionID string `json:"version_id"`
}

type Message struct {
	Cursor     string                 `json:"__cursor"`
	ID         string                 `json:"id"`
	ChannelID  string                 `json:"channel_id"`
	Recipient  string                 `json:"recipient"`
	Workflow   string                 `json:"workflow"`
	Tenant     string                 `json:"tenant"`
	Status     UserMessageStatus      `json:"status"`
	ReadAt     time.Time              `json:"read_at"`
	SeenAt     time.Time              `json:"seen_at"`
	ArchivedAt time.Time              `json:"archived_at"`
	InsertedAt time.Time              `json:"inserted_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	Source     *NotificationSource    `json:"source"`
	Data       map[string]interface{} `json:"data"`
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

type ListMessagesRequest struct {
	PageSize  int                 `url:"page_size,omitempty"`
	Cursor    int                 `url:"page_size,omitempty"`
	Before    string              `url:"before,omitempty"`
	After     string              `url:"after,omitempty"`
	Source    string              `url:"source,omitempty"`
	Tenant    string              `url:"tenant,omitempty"`
	Status    []UserMessageStatus `url:"status,omitempty"`
	ChannelID string              `url:"channel_id,omitempty"`
}
type ListMessagesResponse struct {
	Items    []*Message `json:"items"`
	PageInfo PageInfo   `json:"page_info"`
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
	PageInfo PageInfo        `json:"page_info"`
}

type GetMessageActivitiesRequest = GetMessageRequest
type GetMessageActivitiesResponse struct {
	Activities []*MessageActivity `json:"items"`
	PageInfo   PageInfo           `json:"page_info"`
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

func (ms *messagesService) List(ctx context.Context, listReq *ListMessagesRequest) (*ListMessagesResponse, error) {

	queryString, err := query.Values(listReq)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to list messages")
	}
	path := fmt.Sprintf("%s?%s", messagesAPIBasePath, queryString.Encode())

	req, err := ms.client.newRequest(http.MethodGet, path, listReq)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to list messages")
	}
	listMessagesResponse := &ListMessagesResponse{}

	_, err = ms.client.do(ctx, req, listMessagesResponse)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to list messages")
	}

	return listMessagesResponse, nil
}

func (ms *messagesService) Get(ctx context.Context, getReq *GetMessageRequest) (*GetMessageResponse, error) {

	path := messagesAPIPath(getReq.ID)

	req, err := ms.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to list messages")
	}
	getMessageResponse := &GetMessageResponse{Message: &Message{}}
	_, err = ms.client.do(ctx, req, getMessageResponse.Message)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to list messages")
	}

	return getMessageResponse, nil
}

func (ms *messagesService) GetEvents(ctx context.Context, getEventsReq *GetMessageEventsRequest) (*GetMessageEventsResponse, error) {

	path := fmt.Sprintf("%s/events", messagesAPIPath(getEventsReq.ID))

	req, err := ms.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to list message events")
	}
	getMessageEventsResponse := &GetMessageEventsResponse{}
	_, err = ms.client.do(ctx, req, getMessageEventsResponse)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to list message events")
	}

	return getMessageEventsResponse, nil
}

func (ms *messagesService) GetActivities(ctx context.Context, getActivitiesReq *GetMessageActivitiesRequest) (*GetMessageActivitiesResponse, error) {

	path := fmt.Sprintf("%s/activities", messagesAPIPath(getActivitiesReq.ID))

	req, err := ms.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to list message activities")
	}
	getMessageActivitiesResponse := &GetMessageActivitiesResponse{}
	_, err = ms.client.do(ctx, req, getMessageActivitiesResponse)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to list message activities")
	}

	return getMessageActivitiesResponse, nil
}

func (ms *messagesService) GetContent(ctx context.Context, getContentReq *GetMessageContentRequest) (*GetMessageContentResponse, error) {

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

	return getMessageContentResponse, nil
}

func (ms *messagesService) SetStatus(ctx context.Context, setStatusReq *SetStatusRequest) (*SetStatusResponse, error) {

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

	return setMessageStatusResponse, nil
}

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

type MessageStatus string

const (
	Queued      MessageStatus = "queued"
	Sent        MessageStatus = "sent"
	Delivered   MessageStatus = "delivered"
	Undelivered MessageStatus = "undelivered"
	NotSent     MessageStatus = "not_sent"
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
	Status     MessageStatus          `json:"status"`
	ReadAt     time.Time              `json:"read_at"`
	SeenAt     time.Time              `json:"seen_at"`
	ArchivedAt time.Time              `json:"archived_at"`
	InsertedAt time.Time              `json:"inserted_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	Source     *NotificationSource    `json:"source"`
	Data       map[string]interface{} `json:"data"`
}

type ListMessagesRequest struct {
	PageSize  int             `url:"page_size,omitempty"`
	Cursor    int             `url:"page_size,omitempty"`
	Before    string          `url:"before,omitempty"`
	After     string          `url:"after,omitempty"`
	Source    string          `url:"source,omitempty"`
	Tenant    string          `url:"tenant,omitempty"`
	Status    []MessageStatus `url:"status,omitempty"`
	ChannelID string          `url:"channel_id,omitempty"`
}

type PageInfoMessages struct {
	PageSize int    `json:"page_size,omitempty"`
	Before   string `json:"before,omitempty"`
	After    string `json:"after,omitempty"`
}

type ListMessagesResponse struct {
	Items    []*Message       `json:"items"`
	PageInfo PageInfoMessages `json:"page_info"`
}

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

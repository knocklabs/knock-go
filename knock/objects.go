package knock

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

type ObjectsService interface {
	Get(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
	GetMessages(context.Context, *GetObjectMessagesRequest) (*GetObjectMessagesResponse, error)
	Set(context.Context, *SetObjectRequest) (*SetObjectResponse, error)
	Delete(context.Context, *DeleteObjectRequest) error
}

type objectsService struct {
	client *Client
}

var _ ObjectsService = &objectsService{}

func NewObjectService(client *Client) *objectsService {
	return &objectsService{
		client: client,
	}
}

type Object struct {
	CollectionID string                 `json:"collection"`
	ObjectID     string                 `json:"id"`
	Properties   map[string]interface{} `json:"properties"`
	UpdatedAt    time.Time              `json:"updated_at"`
	CreatedAt    time.Time              `json:"created_at"`
}

type ObjectMessage struct {
	Cursor     string                 `json:"__cursor"`
	ID         string                 `json:"id"`
	ChannelID  string                 `json:"channel_id"`
	Recipient  map[string]interface{} `json:"recipient"`
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

type GetObjectRequest struct {
	ID           string
	CollectionID string
}
type GetObjectResponse struct {
	Object *Object
}

type GetObjectMessagesRequest struct {
	ObjectID     string `url:"-"`
	CollectionID string `url:"-"`

	PageSize  int                 `url:"page_size,omitempty"`
	Before    string              `url:"before,omitempty"`
	After     string              `url:"after,omitempty"`
	Source    string              `url:"source,omitempty"`
	Tenant    string              `url:"tenant,omitempty"`
	Status    []UserMessageStatus `url:"status,omitempty"`
	ChannelID string              `url:"channel_id,omitempty"`
}
type GetObjectMessagesResponse struct {
	Items    []*ObjectMessage `json:"entries"`
	PageInfo PageInfo         `json:"page_info"`
}

type SetObjectRequest struct {
	ID           string                 `json:"-"`
	CollectionID string                 `json:"-"`
	Properties   map[string]interface{} `json:""`
}
type SetObjectResponse = GetObjectResponse

type DeleteObjectRequest = GetObjectRequest

func objectAPIPath(collectionID string, objectID string) string {
	return fmt.Sprintf("v1/objects/%s/%s", collectionID, objectID)
}

func (os *objectsService) Get(ctx context.Context, getObjectRequest *GetObjectRequest) (*GetObjectResponse, error) {
	path := objectAPIPath(getObjectRequest.CollectionID, getObjectRequest.ID)

	req, err := os.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get object")
	}

	getObjectResponse := &GetObjectResponse{Object: &Object{}}
	_, err = os.client.do(ctx, req, getObjectResponse.Object)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get object")
	}

	return getObjectResponse, nil
}

func (os *objectsService) GetMessages(ctx context.Context, getObjectMessagesRequest *GetObjectMessagesRequest) (*GetObjectMessagesResponse, error) {
	queryString, err := query.Values(getObjectMessagesRequest)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing query parameters to list object messages")
	}
	path := fmt.Sprintf("%s/messages?%s", objectAPIPath(getObjectMessagesRequest.CollectionID, getObjectMessagesRequest.ObjectID), queryString.Encode())

	req, err := os.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to list object messages")
	}

	getObjectMessagesResponse := &GetObjectMessagesResponse{}
	_, err = os.client.do(ctx, req, getObjectMessagesResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to list object messages")
	}

	return getObjectMessagesResponse, nil
}

func (os *objectsService) Set(ctx context.Context, setObjectRequest *SetObjectRequest) (*SetObjectResponse, error) {
	path := objectAPIPath(setObjectRequest.CollectionID, setObjectRequest.ID)

	if len(setObjectRequest.Properties) == 0 {
		return nil, &Error{msg: "Must set at least one property"}
	}

	req, err := os.client.newRequest(http.MethodPut, path, setObjectRequest.Properties)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for set object")
	}

	setObjectResponse := &SetObjectResponse{Object: &Object{}}
	_, err = os.client.do(ctx, req, setObjectResponse.Object)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for set object")
	}

	return setObjectResponse, nil
}

// TODO object messages
// TODO

func (os *objectsService) Delete(ctx context.Context, deleteObjectRequest *DeleteObjectRequest) error {
	path := objectAPIPath(deleteObjectRequest.CollectionID, deleteObjectRequest.ID)

	req, err := os.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return errors.Wrap(err, "error creating request for delete object")
	}

	_, err = os.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request for delete object")
	}

	return nil
}

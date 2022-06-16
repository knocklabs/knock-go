package knock

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// ObjectsService is an interface for communicating with the Knock
// Objects API endpoints.
type ObjectsService interface {
	Get(context.Context, *GetObjectRequest) (*Object, error)
	Set(context.Context, *SetObjectRequest) (*Object, error)
	Delete(context.Context, *DeleteObjectRequest) error

	GetMessages(context.Context, *GetObjectMessagesRequest) ([]*ObjectMessage, *PageInfo, error)

	GetChannelData(context.Context, *GetObjectChannelDataRequest) (map[string]interface{}, error)
	SetChannelData(context.Context, *SetObjectChannelDataRequest) (map[string]interface{}, error)
	DeleteChannelData(context.Context, *DeleteObjectChannelDataRequest) error

	GetPreferences(context.Context, *GetObjectPreferencesRequest) (*PreferenceSet, error)
	SetPreferences(context.Context, *SetObjectPreferencesRequest) (*PreferenceSet, error)
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

// Context structs
type Object struct {
	Collection string                 `json:"collection"`
	ID         string                 `json:"id"`
	Properties map[string]interface{} `json:"properties"`
	UpdatedAt  time.Time              `json:"updated_at"`
	CreatedAt  time.Time              `json:"created_at"`
}

type ObjectMessage struct {
	Cursor     string                 `json:"__cursor"`
	ID         string                 `json:"id"`
	ChannelID  string                 `json:"channel_id"`
	Recipient  map[string]interface{} `json:"recipient"`
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

// Client structs
type GetObjectRequest struct {
	ID         string
	Collection string
}
type GetObjectResponse struct {
	Object *Object
}

type SetObjectRequest struct {
	ID         string                 `json:"-"`
	Collection string                 `json:"-"`
	Properties map[string]interface{} `json:""`
}
type SetObjectResponse = GetObjectResponse

type DeleteObjectRequest = GetObjectRequest

type GetObjectMessagesRequest struct {
	ObjectID   string `url:"-"`
	Collection string `url:"-"`

	PageSize  int                `url:"page_size,omitempty"`
	Before    string             `url:"before,omitempty"`
	After     string             `url:"after,omitempty"`
	Source    string             `url:"source,omitempty"`
	Tenant    string             `url:"tenant,omitempty"`
	Status    []EngagementStatus `url:"status,omitempty"`
	ChannelID string             `url:"channel_id,omitempty"`
}
type GetObjectMessagesResponse struct {
	Items    []*ObjectMessage `json:"entries"`
	PageInfo *PageInfo        `json:"page_info"`
}

type GetObjectChannelDataRequest struct {
	Collection string
	ChannelID  string
	ObjectID   string
}
type GetObjectChannelDataResponse = GetUserChannelDataResponse

type SetObjectChannelDataRequest struct {
	Collection string                 `json:"-"`
	ChannelID  string                 `json:"-"`
	ObjectID   string                 `json:"-"`
	Data       map[string]interface{} `json:"data"`
}

type SetObjectChannelDataResponse = GetUserChannelDataResponse

type DeleteObjectChannelDataRequest = GetObjectChannelDataRequest

type GetObjectPreferencesRequest struct {
	PreferenceID string `json:"-"`
	ObjectID     string `json:"-"`
	Collection   string `json:"-"`
}
type GetObjectPreferencesResponse struct {
	Preferences *PreferenceSet
}
type SetObjectPreferencesRequest struct {
	PreferenceID string `json:"-"`
	ObjectID     string `json:"-"`
	Collection   string `json:"-"`

	Workflows    map[string]interface{} `json:"workflows,omitempty"`
	ChannelTypes map[string]interface{} `json:"channel_types,omitempty"`
	Categories   map[string]interface{} `json:"categories,omitempty"`
}
type SetObjectPreferencesResponse = GetObjectPreferencesResponse

func objectAPIPath(collection string, objectID string) string {
	return fmt.Sprintf("v1/objects/%s/%s", collection, objectID)
}

func objectChannelDataAPIPath(collection string, objectID string, channelID string) string {
	return fmt.Sprintf("v1/objects/%s/%s/channel_data/%s", collection, objectID, channelID)
}

func (os *objectsService) Get(ctx context.Context, getObjectRequest *GetObjectRequest) (*Object, error) {
	path := objectAPIPath(getObjectRequest.Collection, getObjectRequest.ID)

	req, err := os.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get object")
	}

	getObjectResponse := &GetObjectResponse{Object: &Object{}}
	_, err = os.client.do(ctx, req, getObjectResponse.Object)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get object")
	}

	return getObjectResponse.Object, nil
}

func (os *objectsService) GetMessages(ctx context.Context, getObjectMessagesRequest *GetObjectMessagesRequest) ([]*ObjectMessage, *PageInfo, error) {
	queryString, err := query.Values(getObjectMessagesRequest)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error parsing query parameters to list object messages")
	}
	path := fmt.Sprintf("%s/messages?%s", objectAPIPath(getObjectMessagesRequest.Collection, getObjectMessagesRequest.ObjectID), queryString.Encode())

	req, err := os.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request to list object messages")
	}

	getObjectMessagesResponse := &GetObjectMessagesResponse{}
	_, err = os.client.do(ctx, req, getObjectMessagesResponse)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request to list object messages")
	}

	return getObjectMessagesResponse.Items, getObjectMessagesResponse.PageInfo, nil
}

func (os *objectsService) Set(ctx context.Context, setObjectRequest *SetObjectRequest) (*Object, error) {
	path := objectAPIPath(setObjectRequest.Collection, setObjectRequest.ID)

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

	return setObjectResponse.Object, nil
}

func (os *objectsService) Delete(ctx context.Context, deleteObjectRequest *DeleteObjectRequest) error {
	path := objectAPIPath(deleteObjectRequest.Collection, deleteObjectRequest.ID)

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

func (os *objectsService) GetChannelData(ctx context.Context, getChannelDataReq *GetObjectChannelDataRequest) (map[string]interface{}, error) {
	path := objectChannelDataAPIPath(getChannelDataReq.Collection, getChannelDataReq.ObjectID, getChannelDataReq.ChannelID)

	req, err := os.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to get object channel data")
	}

	channelDataResponse := &GetUserChannelDataResponse{}
	_, err = os.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to get object channel data")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to get object channel data")
	}

	return channelDataResponse.ChannelData, nil
}

func (os *objectsService) SetChannelData(ctx context.Context, getChannelDataReq *SetObjectChannelDataRequest) (map[string]interface{}, error) {
	path := objectChannelDataAPIPath(getChannelDataReq.Collection, getChannelDataReq.ObjectID, getChannelDataReq.ChannelID)

	req, err := os.client.newRequest(http.MethodPut, path, getChannelDataReq)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set object channel data")
	}

	channelDataResponse := &SetObjectChannelDataResponse{}
	_, err = os.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to set object channel data")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to set object channel data")
	}

	return channelDataResponse.ChannelData, nil
}

func (os *objectsService) DeleteChannelData(ctx context.Context, deleteObjectChannelDataReq *DeleteObjectChannelDataRequest) error {
	path := objectChannelDataAPIPath(deleteObjectChannelDataReq.Collection, deleteObjectChannelDataReq.ObjectID, deleteObjectChannelDataReq.ChannelID)

	req, err := os.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return errors.Wrap(err, "error creating request to delete object channel data")
	}

	_, err = os.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request to delete object channel data")
	}

	if err != nil {
		return errors.Wrap(err, "error parsing request to delete object channel data")
	}

	return nil
}

func (os *objectsService) GetPreferences(ctx context.Context, getObjectPreferencesReq *GetObjectPreferencesRequest) (*PreferenceSet, error) {
	if getObjectPreferencesReq.PreferenceID == "" {
		getObjectPreferencesReq.PreferenceID = "default"
	}
	path := fmt.Sprintf("%s/preferences/%s", objectAPIPath(getObjectPreferencesReq.Collection, getObjectPreferencesReq.ObjectID), getObjectPreferencesReq.PreferenceID)

	req, err := os.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to get object preferences")
	}

	getPreferenceResponse := GetObjectPreferencesResponse{Preferences: &PreferenceSet{}}

	_, err = os.client.do(ctx, req, &getPreferenceResponse.Preferences)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to get object preferences")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to get object preferences")
	}

	return getPreferenceResponse.Preferences, nil

}

func (sr *SetObjectPreferencesRequest) AddChannelTypesPreference(channelType map[string]interface{}) SetObjectPreferencesRequest {
	sr.ChannelTypes = PreferencesMapAppend(sr.ChannelTypes, channelType)
	return *sr
}

func (sr *SetObjectPreferencesRequest) AddWorkflowsPreference(workflows map[string]interface{}) SetObjectPreferencesRequest {
	sr.Workflows = PreferencesMapAppend(sr.Workflows, workflows)
	return *sr
}

func (sr *SetObjectPreferencesRequest) AddCategoriesPreference(categories map[string]interface{}) SetObjectPreferencesRequest {
	sr.Categories = PreferencesMapAppend(sr.Categories, categories)
	return *sr
}

func (os *objectsService) SetPreferences(ctx context.Context, setObjectPreferencesReq *SetObjectPreferencesRequest) (*PreferenceSet, error) {

	if setObjectPreferencesReq.PreferenceID == "" {
		setObjectPreferencesReq.PreferenceID = "default"
	}

	path := fmt.Sprintf("%s/preferences/%s", objectAPIPath(setObjectPreferencesReq.Collection, setObjectPreferencesReq.ObjectID), setObjectPreferencesReq.PreferenceID)

	req, err := os.client.newRequest(http.MethodPut, path, setObjectPreferencesReq)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set object preferences")
	}

	setPreferenceResponse := SetObjectPreferencesResponse{Preferences: &PreferenceSet{}}

	_, err = os.client.do(ctx, req, &setPreferenceResponse.Preferences)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to set object preferences")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to set object preferences")
	}

	return setPreferenceResponse.Preferences, nil
}

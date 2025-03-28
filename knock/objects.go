package knock

import (
	"context"
	"encoding/json"
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
	BulkSet(context.Context, *BulkSetObjectsRequest) (*BulkOperation, error)
	Delete(context.Context, *DeleteObjectRequest) error
	BulkDelete(context.Context, *BulkDeleteObjectsRequest) (*BulkOperation, error)
	List(context.Context, *ListObjectsRequest) ([]*Object, *PageInfo, error)

	GetMessages(context.Context, *GetObjectMessagesRequest) ([]*ObjectMessage, *PageInfo, error)
	GetSchedules(context.Context, *GetObjectSchedulesRequest) ([]*Schedule, *PageInfo, error)
	GetSubscriptions(context.Context, *GetObjectSubscriptionsRequest) ([]*ObjectSubscription, *PageInfo, error)

	GetChannelData(context.Context, *GetObjectChannelDataRequest) (map[string]interface{}, error)
	SetChannelData(context.Context, *SetObjectChannelDataRequest) (map[string]interface{}, error)
	DeleteChannelData(context.Context, *DeleteObjectChannelDataRequest) error

	GetPreferences(context.Context, *GetObjectPreferencesRequest) (*PreferenceSet, error)
	SetPreferences(context.Context, *SetObjectPreferencesRequest) (*PreferenceSet, error)

	AddSubscriptions(context.Context, *AddSubscriptionsRequest) ([]*ObjectSubscription, error)
	BulkAddSubscriptions(context.Context, *BulkAddSubscriptionsRequest) (*BulkOperation, error)
	DeleteSubscriptions(context.Context, *DeleteSubscriptionsRequest) ([]*ObjectSubscription, error)
	ListSubscriptions(context.Context, *ListSubscriptionsRequest) (*ListSubscriptionsResponse, error)
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

type BulkSetObject struct {
	ID         string `json:"id"`
	Properties map[string]interface{}
}

type BulkSetObjectsRequest struct {
	Collection string           `json:"-"`
	Objects    []*BulkSetObject `json:"objects"`
}

type ListObjectsRequest struct {
	Collection string `url:"-"`
	ObjectId   string `url:"object_id,omitempty"`
	Name       string `url:"name,omitempty"`
	PageSize   int    `url:"page_size,omitempty"`
	Before     string `url:"before,omitempty"`
	After      string `url:"after,omitempty"`
}

type SetObjectResponse = GetObjectResponse

type DeleteObjectRequest = GetObjectRequest

type BulkDeleteObjectsRequest struct {
	Collection string   `json:"-"`
	ObjectIDs  []string `json:"object_ids"`
}

type ListObjectsResponse struct {
	Entries  []*Object `json:"entries"`
	PageInfo *PageInfo `json:"page_info"`
}

type GetObjectMessagesRequest struct {
	ObjectID   string `url:"-"`
	Collection string `url:"-"`

	PageSize               int                    `url:"page_size,omitempty"`
	Before                 string                 `url:"before,omitempty"`
	After                  string                 `url:"after,omitempty"`
	Source                 string                 `url:"source,omitempty"`
	Tenant                 string                 `url:"tenant,omitempty"`
	Status                 []EngagementStatus     `url:"status,omitempty"`
	ChannelID              string                 `url:"channel_id,omitempty"`
	TriggerData            map[string]interface{} `url:"-"`
	WorkflowRunID          string                 `url:"workflow_run_id,omitempty"`
	WorkflowRecipientRunID string                 `url:"workflow_recipient_run_id,omitempty"`
}

type GetObjectMessagesResponse struct {
	Items    []*ObjectMessage `json:"items"`
	PageInfo *PageInfo        `json:"page_info"`
}

type GetObjectSchedulesRequest struct {
	ObjectID   string `url:"-"`
	Collection string `url:"-"`

	PageSize int    `url:"page_size,omitempty"`
	Before   string `url:"before,omitempty"`
	After    string `url:"after,omitempty"`
	Workflow string `url:"workflow,omitempty"`
	Tenant   string `url:"tenant,omitempty"`
}

type GetObjectSchedulesResponse struct {
	Items    []*Schedule `json:"entries"`
	PageInfo *PageInfo   `json:"page_info"`
}

type GetObjectSubscriptionsRequest struct {
	Collection string `url:"-"`
	ObjectID   string `url:"-"`
	PageSize   int    `url:"page_size,omitempty"`
	Before     string `url:"before,omitempty"`
	After      string `url:"after,omitempty"`
}

type GetObjectSubscriptionsResponse struct {
	Entries  []*ObjectSubscription `json:"entries"`
	PageInfo *PageInfo             `json:"page_info"`
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

type ObjectSubscription struct {
	Recipient  interface{}            `json:"recipient"`
	Object     Object                 `json:"object"`
	Properties map[string]interface{} `json:"properties"`
	InsertedAt time.Time              `json:"inserted_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}

type AddSubscriptionsRequest struct {
	ObjectID   string                 `json:"-"`
	Collection string                 `json:"-"`
	Recipients []interface{}          `json:"recipients"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

type AddSubscriptionsResponse struct {
	Subscriptions []*ObjectSubscription
}

type BulkAddObjectSubscription struct {
	ID         string                 `json:"id"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Recipients []interface{}          `json:"recipients"`
}

type BulkAddSubscriptionsRequest struct {
	Collection    string                       `json:"-"`
	Subscriptions []*BulkAddObjectSubscription `json:"subscriptions"`
}

type DeleteSubscriptionsRequest struct {
	ObjectID   string        `json:"-"`
	Collection string        `json:"-"`
	Recipients []interface{} `json:"recipients"`
}

type DeleteSubscriptionsResponse struct {
	Subscriptions []*ObjectSubscription
}

type ListSubscriptionsRequest struct {
	ObjectID   string `url:"-"`
	Collection string `url:"-"`
	PageSize   int    `url:"page_size,omitempty"`
	Before     string `url:"before,omitempty"`
	After      string `url:"after,omitempty"`
}

type ListSubscriptionsResponse struct {
	Entries  []*ObjectSubscription `json:"entries"`
	PageInfo *PageInfo             `json:"page_info"`
}

type BulkRequestResponse struct {
	BulkOperation *BulkOperation
}

func collectionAPIPath(collection string) string {
	return fmt.Sprintf("v1/objects/%s", collection)
}

func objectAPIPath(collection string, objectID string) string {
	return fmt.Sprintf("v1/objects/%s/%s", collection, objectID)
}

func objectChannelDataAPIPath(collection string, objectID string, channelID string) string {
	return fmt.Sprintf("v1/objects/%s/%s/channel_data/%s", collection, objectID, channelID)
}

func (os *objectsService) Get(ctx context.Context, getObjectRequest *GetObjectRequest) (*Object, error) {
	path := objectAPIPath(getObjectRequest.Collection, getObjectRequest.ID)

	req, err := os.client.newRequest(http.MethodGet, path, nil, nil)
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

	if getObjectMessagesRequest.TriggerData != nil {
		triggerDataJson, err := json.Marshal(getObjectMessagesRequest.TriggerData)

		if err != nil {
			return nil, nil, errors.Wrap(err, "error converting TriggerData into JSON")
		}

		queryString.Add("trigger_data", string(triggerDataJson))
	}

	path := fmt.Sprintf("%s/messages?%s", objectAPIPath(getObjectMessagesRequest.Collection, getObjectMessagesRequest.ObjectID), queryString.Encode())

	req, err := os.client.newRequest(http.MethodGet, path, nil, nil)
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

func (os *objectsService) GetSchedules(ctx context.Context, getObjectSchedulesRequest *GetObjectSchedulesRequest) ([]*Schedule, *PageInfo, error) {
	queryString, err := query.Values(getObjectSchedulesRequest)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error parsing query parameters to list object schedules")
	}

	path := fmt.Sprintf("%s/schedules?%s", objectAPIPath(getObjectSchedulesRequest.Collection, getObjectSchedulesRequest.ObjectID), queryString.Encode())

	req, err := os.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request to list object schedules")
	}

	getObjectSchedulesResponse := &GetObjectSchedulesResponse{}
	_, err = os.client.do(ctx, req, getObjectSchedulesResponse)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request to list object schedules")
	}

	return getObjectSchedulesResponse.Items, getObjectSchedulesResponse.PageInfo, nil
}

func (us *objectsService) GetSubscriptions(ctx context.Context, getObjectSubscriptionsReq *GetObjectSubscriptionsRequest) ([]*ObjectSubscription, *PageInfo, error) {
	queryString, _ := query.Values(getObjectSubscriptionsReq)
	path := fmt.Sprintf("%s/subscriptions?mode=recipient&%s", objectAPIPath(getObjectSubscriptionsReq.Collection, getObjectSubscriptionsReq.ObjectID), queryString.Encode())

	req, err := us.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request to get object subscriptions")
	}
	getObjectSubscriptionsResponse := &GetObjectSubscriptionsResponse{}

	_, err = us.client.do(ctx, req, getObjectSubscriptionsResponse)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request to get object subscriptions")
	}

	return getObjectSubscriptionsResponse.Entries, getObjectSubscriptionsResponse.PageInfo, nil
}

func (os *objectsService) Set(ctx context.Context, setObjectRequest *SetObjectRequest) (*Object, error) {
	path := objectAPIPath(setObjectRequest.Collection, setObjectRequest.ID)

	if len(setObjectRequest.Properties) == 0 {
		return nil, &Error{msg: "Must set at least one property"}
	}

	req, err := os.client.newRequest(http.MethodPut, path, setObjectRequest.Properties, nil)
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

func (os *objectsService) BulkSet(ctx context.Context, bulkSetObjectsRequest *BulkSetObjectsRequest) (*BulkOperation, error) {
	path := fmt.Sprintf("%s/bulk/set", collectionAPIPath(bulkSetObjectsRequest.Collection))

	req, err := os.client.newRequest(http.MethodPost, path, bulkSetObjectsRequest, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to bulk set objects")
	}

	bulkSetObjectsRes := BulkRequestResponse{BulkOperation: &BulkOperation{}}

	_, err = os.client.do(ctx, req, bulkSetObjectsRes.BulkOperation)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to bulk set objects")
	}

	return bulkSetObjectsRes.BulkOperation, nil
}

func (os *objectsService) Delete(ctx context.Context, deleteObjectRequest *DeleteObjectRequest) error {
	path := objectAPIPath(deleteObjectRequest.Collection, deleteObjectRequest.ID)

	req, err := os.client.newRequest(http.MethodDelete, path, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error creating request for delete object")
	}

	_, err = os.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request for delete object")
	}

	return nil
}

func (os *objectsService) BulkDelete(ctx context.Context, bulkDeleteObjectsRequest *BulkDeleteObjectsRequest) (*BulkOperation, error) {
	path := fmt.Sprintf("%s/bulk/delete", collectionAPIPath(bulkDeleteObjectsRequest.Collection))

	req, err := os.client.newRequest(http.MethodPost, path, bulkDeleteObjectsRequest, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to bulk delete objects")
	}

	bulkDeleteObjectsRes := BulkRequestResponse{BulkOperation: &BulkOperation{}}

	_, err = os.client.do(ctx, req, bulkDeleteObjectsRes.BulkOperation)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to bulk delete objects")
	}

	return bulkDeleteObjectsRes.BulkOperation, nil
}

func (os *objectsService) List(ctx context.Context, listObjectsRequest *ListObjectsRequest) ([]*Object, *PageInfo, error) {
	queryString, err := query.Values(listObjectsRequest)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error parsing query params to list objects")
	}

	collectionPath := collectionAPIPath(listObjectsRequest.Collection)
	path := fmt.Sprintf("%s?%s", collectionPath, queryString.Encode())

	req, err := os.client.newRequest(http.MethodGet, path, listObjectsRequest, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request for listing objects")
	}

	listObjectsResponse := &ListObjectsResponse{}
	_, err = os.client.do(ctx, req, listObjectsResponse)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request for listing objects")
	}

	return listObjectsResponse.Entries, listObjectsResponse.PageInfo, nil
}

func (os *objectsService) GetChannelData(ctx context.Context, getChannelDataReq *GetObjectChannelDataRequest) (map[string]interface{}, error) {
	path := objectChannelDataAPIPath(getChannelDataReq.Collection, getChannelDataReq.ObjectID, getChannelDataReq.ChannelID)

	req, err := os.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to get object channel data")
	}

	channelDataResponse := &GetUserChannelDataResponse{}
	_, err = os.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to get object channel data")
	}

	return channelDataResponse.ChannelData, nil
}

func (os *objectsService) SetChannelData(ctx context.Context, getChannelDataReq *SetObjectChannelDataRequest) (map[string]interface{}, error) {
	path := objectChannelDataAPIPath(getChannelDataReq.Collection, getChannelDataReq.ObjectID, getChannelDataReq.ChannelID)

	req, err := os.client.newRequest(http.MethodPut, path, getChannelDataReq, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set object channel data")
	}

	channelDataResponse := &SetObjectChannelDataResponse{}
	_, err = os.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to set object channel data")
	}

	return channelDataResponse.ChannelData, nil
}

func (os *objectsService) DeleteChannelData(ctx context.Context, deleteObjectChannelDataReq *DeleteObjectChannelDataRequest) error {
	path := objectChannelDataAPIPath(deleteObjectChannelDataReq.Collection, deleteObjectChannelDataReq.ObjectID, deleteObjectChannelDataReq.ChannelID)

	req, err := os.client.newRequest(http.MethodDelete, path, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error creating request to delete object channel data")
	}

	_, err = os.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request to delete object channel data")
	}

	return nil
}

func (os *objectsService) GetPreferences(ctx context.Context, getObjectPreferencesReq *GetObjectPreferencesRequest) (*PreferenceSet, error) {
	if getObjectPreferencesReq.PreferenceID == "" {
		getObjectPreferencesReq.PreferenceID = "default"
	}
	path := fmt.Sprintf("%s/preferences/%s", objectAPIPath(getObjectPreferencesReq.Collection, getObjectPreferencesReq.ObjectID), getObjectPreferencesReq.PreferenceID)

	req, err := os.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to get object preferences")
	}

	getPreferenceResponse := GetObjectPreferencesResponse{Preferences: &PreferenceSet{}}

	_, err = os.client.do(ctx, req, &getPreferenceResponse.Preferences)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to get object preferences")
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

	req, err := os.client.newRequest(http.MethodPut, path, setObjectPreferencesReq, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set object preferences")
	}

	setPreferenceResponse := SetObjectPreferencesResponse{Preferences: &PreferenceSet{}}

	_, err = os.client.do(ctx, req, &setPreferenceResponse.Preferences)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to set object preferences")
	}

	return setPreferenceResponse.Preferences, nil
}

func (os *objectsService) AddSubscriptions(ctx context.Context, addSubscriptionsReq *AddSubscriptionsRequest) ([]*ObjectSubscription, error) {
	path := fmt.Sprintf("%s/subscriptions", objectAPIPath(addSubscriptionsReq.Collection, addSubscriptionsReq.ObjectID))

	req, err := os.client.newRequest(http.MethodPost, path, addSubscriptionsReq, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to add subscriptions")
	}

	addSubscriptionsResponse := AddSubscriptionsResponse{Subscriptions: []*ObjectSubscription{}}

	_, err = os.client.do(ctx, req, &addSubscriptionsResponse.Subscriptions)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to add subscriptions")
	}

	return addSubscriptionsResponse.Subscriptions, nil
}

func (os *objectsService) BulkAddSubscriptions(ctx context.Context, bulkAddSubscriptionsRequest *BulkAddSubscriptionsRequest) (*BulkOperation, error) {
	path := fmt.Sprintf("%s/bulk/subscriptions/add", collectionAPIPath(bulkAddSubscriptionsRequest.Collection))

	req, err := os.client.newRequest(http.MethodPost, path, bulkAddSubscriptionsRequest, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to bulk add subscriptions")
	}

	bulkAddSubscriptionsRes := BulkRequestResponse{BulkOperation: &BulkOperation{}}

	_, err = os.client.do(ctx, req, bulkAddSubscriptionsRes.BulkOperation)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to bulk add subscriptions")
	}

	return bulkAddSubscriptionsRes.BulkOperation, nil
}

func (os *objectsService) DeleteSubscriptions(ctx context.Context, deleteSubscriptionsReq *DeleteSubscriptionsRequest) ([]*ObjectSubscription, error) {
	path := fmt.Sprintf("%s/subscriptions", objectAPIPath(deleteSubscriptionsReq.Collection, deleteSubscriptionsReq.ObjectID))

	req, err := os.client.newRequest(http.MethodDelete, path, deleteSubscriptionsReq, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to delete subscriptions")
	}

	deleteSubscriptionsResponse := DeleteSubscriptionsResponse{Subscriptions: []*ObjectSubscription{}}

	_, err = os.client.do(ctx, req, &deleteSubscriptionsResponse.Subscriptions)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to delete subscriptions")
	}

	return deleteSubscriptionsResponse.Subscriptions, nil
}

func (os *objectsService) ListSubscriptions(ctx context.Context, listSubscriptionsReq *ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
	queryString, err := query.Values(listSubscriptionsReq)

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to list subscriptions")
	}

	path := fmt.Sprintf("%s/subscriptions?%s", objectAPIPath(listSubscriptionsReq.Collection, listSubscriptionsReq.ObjectID), queryString.Encode())
	req, err := os.client.newRequest(http.MethodGet, path, listSubscriptionsReq, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request to list subscriptions")
	}

	listRes := &ListSubscriptionsResponse{}
	_, err = os.client.do(ctx, req, listRes)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to list subscriptions")
	}

	return listRes, nil
}

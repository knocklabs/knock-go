package knock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// UsersService is an interface for communicating with the Knock
// Users API endpoints.
type UsersService interface {
	Get(context.Context, *GetUserRequest) (*User, error)
	Identify(context.Context, *IdentifyUserRequest) (*User, error)
	Merge(context.Context, *MergeUserRequest) (*User, error)
	Delete(context.Context, *DeleteUserRequest) error

	GetMessages(context.Context, *GetUserMessagesRequest) ([]*Message, error)
	GetSchedules(context.Context, *GetUserSchedulesRequest) ([]*Schedule, *PageInfo, error)

	BulkIdentify(context.Context, *BulkIdentifyUserRequest) (*BulkOperation, error)
	BulkDelete(context.Context, *BulkDeleteUserRequest) (*BulkOperation, error)

	GetFeed(context.Context, *GetFeedRequest) (*Feed, error)

	GetChannelData(context.Context, *GetUserChannelDataRequest) (map[string]interface{}, error)
	SetChannelData(context.Context, *SetUserChannelDataRequest) (map[string]interface{}, error)
	DeleteChannelData(context.Context, *DeleteUserChannelDataRequest) error

	GetAllPreferences(context.Context, *GetAllPreferencesRequest) ([]*PreferenceSet, error)
	GetPreferences(context.Context, *GetUserPreferencesRequest) (*PreferenceSet, error)
	SetPreferences(context.Context, *SetUserPreferencesRequest) (*PreferenceSet, error)
	BulkSetPreferences(context.Context, *BulkSetUserPreferencesRequest) (*BulkOperation, error)
}

type usersService struct {
	client *Client
}

var _ UsersService = &usersService{}

func NewUsersService(client *Client) *usersService {
	return &usersService{
		client: client,
	}
}

type User struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
	Avatar           string    `json:"avatar,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CustomProperties map[string]interface{}
}

// Context structs
type Feed struct {
	FeedItems    []*FeedItem            `json:"entries"`
	PageInfo     *PageInfo              `json:"page_info"`
	Vars         map[string]interface{} `json:"vars"`
	FeedMetadata *FeedMetadata          `json:"meta"`
}

type FeedItem struct {
	Activities      []*MessageActivity     `json:"activities"`
	Actors          []*User                `json:"actors"`
	TotalActivities int                    `json:"total_activities"`
	TotalActors     int                    `json:"total_actors"`
	Blocks          []*FeedBlock           `json:"blocks"`
	Data            map[string]interface{} `json:"vars"`
	InsertedAt      time.Time              `json:"inserted_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	ReadAt          time.Time              `json:"read_at"`
	SeenAt          time.Time              `json:"seen_at"`
	ArchivedAt      time.Time              `json:"archived_at"`
	Source          NotificationSource     `json:"source"`
	Tenant          string                 `json:"tenant"`
}

type FeedMetadata struct {
	UnreadCount int `json:"unread_count"`
	UnseenCount int `json:"unseen_count"`
}

type FeedBlock struct {
	Content  string `json:"content"`
	Name     string `json:"name"`
	Rendered string `json:"rendered"`
	Type     string `json:"type"`
}

// Client structs
type GetUserRequest struct {
	ID string
}

type GetUserResponse struct {
	User *User
}
type IdentifyUserRequest = User
type IdentifyUserResponse = GetUserResponse

type MergeUserRequest struct {
	ID         string
	FromUserID string `json:"from_user_id"`
}
type MergeUserResponse = GetUserResponse

type DeleteUserRequest struct {
	ID string
}

type GetUserMessagesRequest struct {
	ID          string                 `url:"-"`
	PageSize    int                    `url:"page_size,omitempty"`
	After       string                 `url:"after,omitempty"`
	Before      string                 `url:"before,omitempty"`
	Source      string                 `url:"source,omitempty"`
	Tenant      string                 `url:"tenant,omitempty"`
	Status      []EngagementStatus     `url:"status,omitempty"`
	ChannelID   string                 `url:"channel_id,omitempty"`
	TriggerData map[string]interface{} `url:"-"`
}

type GetUserMessagesResponse struct {
	Messages []*Message `json:"items"`
}

type GetUserSchedulesRequest struct {
	ID string `url:"-"`

	PageSize int    `url:"page_size,omitempty"`
	Before   string `url:"before,omitempty"`
	After    string `url:"after,omitempty"`
	Workflow string `url:"workflow,omitempty"`
	Tenant   string `url:"tenant,omitempty"`
}
type GetUserSchedulesResponse struct {
	Items    []*Schedule `json:"entries"`
	PageInfo *PageInfo   `json:"page_info"`
}

type BulkIdentifyUserRequest struct {
	Users []*IdentifyUserRequest
}

type BulkIdentifyUserResponse struct {
	BulkOperation *BulkOperation
}

type BulkDeleteUserRequest struct {
	UserIDs []string `json:"user_ids"`
}
type BulkDeleteUserResponse = BulkIdentifyUserResponse

type GetFeedRequest struct {
	UserID      string                 `url:"-"`
	FeedID      string                 `url:"-"`
	PageSize    int                    `url:"page_size,omitempty"`
	After       string                 `url:"after,omitempty"`
	Before      string                 `url:"before,omitempty"`
	Status      []string               `url:"status,omitempty"`
	Source      string                 `url:"source,omitempty"`
	Tenant      string                 `url:"tenant,omitempty"`
	HasTenant   bool                   `url:"has_tenant,omitempty"`
	Archived    string                 `url:"archived,omitempty"`
	TriggerData map[string]interface{} `url:"-"`
}
type GetFeedResponse struct {
	Feed *Feed
}

type GetUserChannelDataRequest struct {
	UserID    string
	ChannelID string
}
type GetUserChannelDataResponse struct {
	ChannelData map[string]interface{} `json:"data"`
}

type SetUserChannelDataRequest struct {
	UserID    string                 `json:"-"`
	ChannelID string                 `json:"-"`
	Data      map[string]interface{} `json:"data"`
}
type SetUserChannelDataResponse = GetUserChannelDataResponse

type DeleteUserChannelDataRequest = GetUserChannelDataRequest

type GetAllPreferencesRequest struct {
	UserID string `json:"-"`
}

type GetAllPreferencesResponse struct {
	Preferences []*PreferenceSet
}

type GetUserPreferencesRequest struct {
	UserID       string `json:"-"`
	PreferenceID string `json:"-"`
}
type GetUserPreferencesResponse struct {
	Preferences *PreferenceSet
}

type SetUserPreferencesRequest struct {
	UserId       string `json:"-"`
	PreferenceID string `json:"-"`

	Workflows    map[string]interface{} `json:"workflows,omitempty"`
	ChannelTypes map[string]interface{} `json:"channel_types,omitempty"`
	Categories   map[string]interface{} `json:"categories,omitempty"`
}
type SetUserPreferencesResponse = GetUserPreferencesResponse

type BulkSetUserPreferencesRequest struct {
	UserIDs     []string      `json:"user_ids"`
	Preferences PreferenceSet `json:"preferences"`
}

type BulkSetUserPreferencesResponse = BulkIdentifyUserResponse

func UsersAPIPath(userId string) string {
	return fmt.Sprintf("v1/users/%s", userId)
}

func usersChannelDataAPIPath(userID string, channelID string) string {
	return fmt.Sprintf("%s/channel_data/%s", UsersAPIPath(userID), channelID)
}

func feedsAPIPath(userID string, FeedID string) string {
	return fmt.Sprintf("%s/feeds/%s", UsersAPIPath(userID), FeedID)
}

func (us *usersService) Identify(ctx context.Context, identifyReq *IdentifyUserRequest) (*User, error) {
	path := UsersAPIPath(identifyReq.ID)

	identifyBody, err := identifyReq.toMapWithCustomProperties()
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for identify user")
	}

	req, err := us.client.newRequest(http.MethodPut, path, identifyBody, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for identify user")
	}

	body, err := us.client.do(ctx, req, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for identify user")
	}

	user, err := parseRawUserResponseCustomProperties(body)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for identify user")
	}

	return user, nil
}

func (us *usersService) Get(ctx context.Context, getReq *GetUserRequest) (*User, error) {
	path := UsersAPIPath(getReq.ID)

	req, err := us.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get user")
	}

	body, err := us.client.do(ctx, req, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get user")
	}

	user, err := parseRawUserResponseCustomProperties(body)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for get user")
	}

	return user, nil
}

func (us *usersService) Delete(ctx context.Context, deleteReq *DeleteUserRequest) error {
	path := UsersAPIPath(deleteReq.ID)
	req, err := us.client.newRequest(http.MethodDelete, path, nil, nil)

	if err != nil {
		return errors.Wrap(err, "error creating request for delete user")
	}

	_, err = us.client.do(ctx, req, nil)
	return err
}

func (us *usersService) Merge(ctx context.Context, mergeReq *MergeUserRequest) (*User, error) {
	path := fmt.Sprintf("%s/merge", UsersAPIPath(mergeReq.ID))

	req, err := us.client.newRequest(http.MethodPost, path, mergeReq, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for merge user")
	}

	body, err := us.client.do(ctx, req, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for merge user")
	}
	user, err := parseRawUserResponseCustomProperties(body)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for merge user")
	}

	return user, nil
}

func (us *usersService) GetMessages(ctx context.Context, getUserMessagesReq *GetUserMessagesRequest) ([]*Message, error) {
	queryString, _ := query.Values(getUserMessagesReq)

	if getUserMessagesReq.TriggerData != nil {
		triggerDataJson, err := json.Marshal(getUserMessagesReq.TriggerData)

		if err != nil {
			return nil, errors.Wrap(err, "error converting TriggerData into JSON")
		}

		queryString.Add("trigger_data", string(triggerDataJson))
	}

	path := fmt.Sprintf("%s/messages?%s", UsersAPIPath(getUserMessagesReq.ID), queryString.Encode())

	req, err := us.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get user messages")
	}
	getUserMessagesResponse := &GetUserMessagesResponse{}

	_, err = us.client.do(ctx, req, getUserMessagesResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to list user messages")
	}

	return getUserMessagesResponse.Messages, nil
}

func (us *usersService) GetSchedules(ctx context.Context, getUserSchedulesReq *GetUserSchedulesRequest) ([]*Schedule, *PageInfo, error) {
	queryString, _ := query.Values(getUserSchedulesReq)
	path := fmt.Sprintf("%s/schedules?%s", UsersAPIPath(getUserSchedulesReq.ID), queryString.Encode())

	req, err := us.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request for get user schedules")
	}
	getUserSchedulesResponse := &GetUserSchedulesResponse{}

	_, err = us.client.do(ctx, req, getUserSchedulesResponse)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request for get user schedules")
	}

	return getUserSchedulesResponse.Items, getUserSchedulesResponse.PageInfo, nil
}

func (us *usersService) BulkIdentify(ctx context.Context, bulkIdentifyReq *BulkIdentifyUserRequest) (*BulkOperation, error) {
	path := UsersAPIPath("bulk/identify")

	var bulkUserFlatMap []map[string]interface{}

	for _, rawUser := range bulkIdentifyReq.Users {
		userFlatMap, err := rawUser.toMapWithCustomProperties()
		if err != nil {
			return nil, errors.Wrap(err, "error parsing user custom properties in bulk identify users")
		}
		bulkUserFlatMap = append(bulkUserFlatMap, userFlatMap)
	}

	// Define a struct used only to wrap this data in a `users` json key.
	type BulkUserList struct {
		BulkUserFlatMap []map[string]interface{} `json:"users"`
	}
	req, err := us.client.newRequest(http.MethodPost, path, BulkUserList{BulkUserFlatMap: bulkUserFlatMap}, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for bulk identify user")
	}

	bulkIdentifyRes := &BulkIdentifyUserResponse{BulkOperation: &BulkOperation{}}

	_, err = us.client.do(ctx, req, bulkIdentifyRes.BulkOperation)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for bulk identify user")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for bulk identify user")
	}

	return bulkIdentifyRes.BulkOperation, nil
}

func (us *usersService) BulkDelete(ctx context.Context, bulkDeleteReq *BulkDeleteUserRequest) (*BulkOperation, error) {
	path := UsersAPIPath("bulk/delete")

	req, err := us.client.newRequest(http.MethodPost, path, bulkDeleteReq, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for bulk delete user")
	}

	bulkDeleteRes := &BulkDeleteUserResponse{BulkOperation: &BulkOperation{}}

	_, err = us.client.do(ctx, req, bulkDeleteRes.BulkOperation)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for bulk delete user")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for bulk delete user")
	}

	return bulkDeleteRes.BulkOperation, nil
}

func (us *usersService) GetFeed(ctx context.Context, getFeedReq *GetFeedRequest) (*Feed, error) {
	queryString, _ := query.Values(getFeedReq)

	if getFeedReq.TriggerData != nil {
		triggerDataJson, err := json.Marshal(getFeedReq.TriggerData)

		if err != nil {
			return nil, errors.Wrap(err, "error converting TriggerData into JSON")
		}

		queryString.Add("trigger_data", string(triggerDataJson))
	}

	path := fmt.Sprintf("%s?%s", feedsAPIPath(getFeedReq.UserID, getFeedReq.FeedID), queryString.Encode())

	req, err := us.client.newRequest(http.MethodGet, path, getFeedReq, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get feed")
	}

	getFeedRes := &GetFeedResponse{Feed: &Feed{}}
	_, err = us.client.do(ctx, req, getFeedRes.Feed)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get feed")
	}

	return getFeedRes.Feed, nil
}

func (us *usersService) GetChannelData(ctx context.Context, getChannelDataReq *GetUserChannelDataRequest) (map[string]interface{}, error) {
	path := usersChannelDataAPIPath(getChannelDataReq.UserID, getChannelDataReq.ChannelID)

	req, err := us.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for channel data for user")
	}

	channelDataResponse := &GetUserChannelDataResponse{}
	_, err = us.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for channel data for user")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for channel data for user")
	}

	return channelDataResponse.ChannelData, nil
}

func (us *usersService) SetChannelData(ctx context.Context, getChannelDataReq *SetUserChannelDataRequest) (map[string]interface{}, error) {
	path := usersChannelDataAPIPath(getChannelDataReq.UserID, getChannelDataReq.ChannelID)

	req, err := us.client.newRequest(http.MethodPut, path, getChannelDataReq, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set channel data for user")
	}

	channelDataResponse := &SetUserChannelDataResponse{}
	_, err = us.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to set channel data for user")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to set channel data for user")
	}

	return channelDataResponse.ChannelData, nil
}

func (us *usersService) DeleteChannelData(ctx context.Context, deleteUserChannelDataReq *DeleteUserChannelDataRequest) error {
	path := usersChannelDataAPIPath(deleteUserChannelDataReq.UserID, deleteUserChannelDataReq.ChannelID)

	req, err := us.client.newRequest(http.MethodDelete, path, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error creating request to delete channel data for a user")
	}

	_, err = us.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request to delete channel data for a user")
	}

	if err != nil {
		return errors.Wrap(err, "error parsing request to delete channel data for a user")
	}

	return nil
}

func (us *usersService) GetAllPreferences(ctx context.Context, allPreferencesReq *GetAllPreferencesRequest) ([]*PreferenceSet, error) {

	path := fmt.Sprintf("%s/preferences", UsersAPIPath(allPreferencesReq.UserID))

	req, err := us.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to get all preferences")
	}

	allPreferencesResponse := GetAllPreferencesResponse{Preferences: []*PreferenceSet{}}

	_, err = us.client.do(ctx, req, &allPreferencesResponse.Preferences)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to get all preferences")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to get all preferences")
	}

	return allPreferencesResponse.Preferences, nil
}

func (us *usersService) GetPreferences(ctx context.Context, getPreferencesReq *GetUserPreferencesRequest) (*PreferenceSet, error) {
	if getPreferencesReq.PreferenceID == "" {
		getPreferencesReq.PreferenceID = "default"
	}

	path := fmt.Sprintf("%s/preferences/%s", UsersAPIPath(getPreferencesReq.UserID), getPreferencesReq.PreferenceID)

	req, err := us.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to get all preferences")
	}

	getPreferenceResponse := GetUserPreferencesResponse{Preferences: &PreferenceSet{}}

	_, err = us.client.do(ctx, req, &getPreferenceResponse.Preferences)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to get all preferences")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to get all preferences")
	}

	return getPreferenceResponse.Preferences, nil
}

func (sr *SetUserPreferencesRequest) AddChannelTypesPreference(channelType map[string]interface{}) SetUserPreferencesRequest {
	sr.ChannelTypes = PreferencesMapAppend(sr.ChannelTypes, channelType)
	return *sr
}

func (sr *SetUserPreferencesRequest) AddWorkflowsPreference(workflows map[string]interface{}) SetUserPreferencesRequest {
	sr.Workflows = PreferencesMapAppend(sr.Workflows, workflows)
	return *sr
}

func (sr *SetUserPreferencesRequest) AddCategoriesPreference(categories map[string]interface{}) SetUserPreferencesRequest {
	sr.Categories = PreferencesMapAppend(sr.Categories, categories)
	return *sr
}

func (us *usersService) SetPreferences(ctx context.Context, setPreferencesReq *SetUserPreferencesRequest) (*PreferenceSet, error) {
	if setPreferencesReq.PreferenceID == "" {
		setPreferencesReq.PreferenceID = "default"
	}
	path := fmt.Sprintf("%s/preferences/%s", UsersAPIPath(setPreferencesReq.UserId), setPreferencesReq.PreferenceID)

	req, err := us.client.newRequest(http.MethodPut, path, setPreferencesReq, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set all preferences")
	}

	setPreferenceResponse := SetUserPreferencesResponse{Preferences: &PreferenceSet{}}

	_, err = us.client.do(ctx, req, &setPreferenceResponse.Preferences)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to set all preferences")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to set all preferences")
	}

	return setPreferenceResponse.Preferences, nil
}

func (br *BulkSetUserPreferencesRequest) AddChannelTypesPreference(channelType map[string]interface{}) BulkSetUserPreferencesRequest {
	br.Preferences.ChannelTypes = PreferencesMapAppend(br.Preferences.ChannelTypes, channelType)
	return *br
}

func (br *BulkSetUserPreferencesRequest) AddWorkflowsPreference(workflow map[string]interface{}) BulkSetUserPreferencesRequest {
	br.Preferences.Workflows = PreferencesMapAppend(br.Preferences.Workflows, workflow)
	return *br
}

func (br *BulkSetUserPreferencesRequest) AddCategoriesPreference(category map[string]interface{}) BulkSetUserPreferencesRequest {
	br.Preferences.Categories = PreferencesMapAppend(br.Preferences.Categories, category)
	return *br
}

func (us *usersService) BulkSetPreferences(ctx context.Context, setPreferencesReq *BulkSetUserPreferencesRequest) (*BulkOperation, error) {

	req, err := us.client.newRequest(http.MethodPost, "v1/users/bulk/preferences", setPreferencesReq, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to bulk set user preferences")
	}

	bulkSetPreferencesResponse := BulkSetUserPreferencesResponse{}

	_, err = us.client.do(ctx, req, &bulkSetPreferencesResponse.BulkOperation)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to bulk set user preferences")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to bulk set user preferences")
	}

	return bulkSetPreferencesResponse.BulkOperation, nil
}

// IdentifyUserRequests can contain arbitrary customer data that must be stored, and must be mapped
// appropriately to one big flat list of key value pairs.
func (user *User) toMapWithCustomProperties() (map[string]interface{}, error) {
	flatMap := make(map[string]interface{})

	// Note that marshalling and then immediately un-marshalling transforms struct keys
	// into the parameter names accepted by the API (i.e. PhoneNumber -> phone_number)
	data, _ := json.Marshal(user)
	err := json.Unmarshal(data, &flatMap)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing user custom properties")
	}

	if flatMap["CustomProperties"] != nil {
		// Move all keys from a nested key in the map to the top level of the map
		for k, v := range flatMap["CustomProperties"].(map[string]interface{}) {
			flatMap[k] = v
		}
		delete(flatMap, "CustomProperties")
	}
	return flatMap, nil
}

func parseRawUserResponseCustomProperties(rawResponse []byte) (*User, error) {
	user := User{}
	err := json.Unmarshal(rawResponse, &user)
	if err != nil {
		return nil, err
	}

	// Create a map of the full API response, removing keys explicitly defined in the struct
	// Any remaining keys are custom properties, which will be added to the struct's CustomProperties field
	var customProperties map[string]interface{}
	err = json.Unmarshal(rawResponse, &customProperties)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing user custom properties")
	}

	val := reflect.ValueOf(user)
	for i := 0; i < val.Type().NumField(); i++ {
		delete(customProperties, val.Type().Field(i).Tag.Get("json"))
	}
	// This is returned in API responses but is not used
	delete(customProperties, "__typename")

	user.CustomProperties = customProperties

	return &user, nil
}

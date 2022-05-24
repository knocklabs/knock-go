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
	Identify(context.Context, *IdentifyUserRequest) (*IdentifyUserResponse, error)
	Get(context.Context, *GetUserRequest) (*GetUserResponse, error)
	Delete(context.Context, *DeleteUserRequest) error
	Merge(context.Context, *MergeUserRequest) (*MergeUserResponse, error)
	GetMessages(context.Context, *GetUserMessagesRequest) (*GetUserMessagesResponse, error)
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
	Avatar           string    `json:"avatar"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CustomProperties map[string]interface{}
}

type GetUserRequest struct {
	// User unique identifier
	ID string
}

type DeleteUserRequest struct {
	// User unique identifier
	ID string
}

type IdentifyUserRequest struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar,omitempty"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	CustomProperties map[string]interface{}
}

type MergeUserRequest struct {
	// User unique identifier
	ID         string
	FromUserID string `json:"from_user_id"`
}

type GetUserMessagesRequest struct {
	// User unique identifier
	ID        string              `url:"-"`
	PageSize  int                 `url:"page_size,omitempty"`
	After     string              `url:"after,omitempty"`
	Before    string              `url:"before,omitempty"`
	Source    string              `url:"source,omitempty"`
	Tenant    string              `url:"tenant,omitempty"`
	Status    []UserMessageStatus `url:"status,omitempty"`
	ChannelID string              `url:"channel_id,omitempty"`
}

type GetUserMessagesResponse struct {
	Messages []*Message `json:"items"`
}

type GetUserResponse struct {
	User *User
}
type MergeUserResponse = GetUserResponse
type IdentifyUserResponse = GetUserResponse

func (us *usersService) Identify(ctx context.Context, identifyReq *IdentifyUserRequest) (*IdentifyUserResponse, error) {
	path := UsersAPIPath(identifyReq.ID)

	identifyBody := identifyReq.toMapWithCustomProperties()
	req, err := us.client.newRequest(http.MethodPut, path, identifyBody)

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

	return &IdentifyUserResponse{
		User: user,
	}, nil
}

func (us *usersService) Get(ctx context.Context, getReq *GetUserRequest) (*GetUserResponse, error) {
	path := UsersAPIPath(getReq.ID)

	req, err := us.client.newRequest(http.MethodGet, path, nil)
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

	return &GetUserResponse{
		User: user,
	}, nil
}

func (us *usersService) Delete(ctx context.Context, deleteReq *DeleteUserRequest) error {
	path := UsersAPIPath(deleteReq.ID)
	req, err := us.client.newRequest(http.MethodDelete, path, nil)

	if err != nil {
		return errors.Wrap(err, "error creating request for delete user")
	}

	_, err = us.client.do(ctx, req, nil)
	return err
}

func UsersAPIPath(userId string) string {
	return fmt.Sprintf("v1/users/%s", userId)
}

func (us *usersService) Merge(ctx context.Context, mergeReq *MergeUserRequest) (*MergeUserResponse, error) {
	path := fmt.Sprintf("%s/merge", UsersAPIPath(mergeReq.ID))

	req, err := us.client.newRequest(http.MethodPost, path, mergeReq)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for merge user")
	}

	body, err := us.client.do(ctx, req, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for merge user")
	}
	user, err := parseRawUserResponseCustomProperties(body)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for get user")
	}

	return &MergeUserResponse{
		User: user,
	}, nil
}

func (us *usersService) GetMessages(ctx context.Context, getUserMessagesReq *GetUserMessagesRequest) (*GetUserMessagesResponse, error) {

	queryString, _ := query.Values(getUserMessagesReq)
	path := fmt.Sprintf("%s/messages?%s", UsersAPIPath(getUserMessagesReq.ID), queryString.Encode())

	req, err := us.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get user messages")
	}
	getUserMessagesResponse := &GetUserMessagesResponse{}

	_, err = us.client.do(ctx, req, getUserMessagesResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get user messages")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for get user messages")
	}

	return getUserMessagesResponse, nil
}

// IdentifyUserRequests can contain arbitrary customer data that must be stored, and must be mapped
// appropriately to one big flat list of key value pairs.
func (identifyReq *IdentifyUserRequest) toMapWithCustomProperties() map[string]interface{} {
	flatMap := make(map[string]interface{})

	// Note that marshalling and then immediately un-marshalling transforms struct keys
	// into the parameter names accepted by the API (i.e. PhoneNumber -> phone_number)
	data, _ := json.Marshal(identifyReq)
	json.Unmarshal(data, &flatMap)

	if flatMap["CustomProperties"] != nil {
		// Move all keys from a nested key in the map to the top level of the map
		for k, v := range flatMap["CustomProperties"].(map[string]interface{}) {
			flatMap[k] = v
		}
		delete(flatMap, "CustomProperties")
	}
	return flatMap
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
	json.Unmarshal(rawResponse, &customProperties)
	val := reflect.ValueOf(user)
	for i := 0; i < val.Type().NumField(); i++ {
		delete(customProperties, val.Type().Field(i).Tag.Get("json"))
	}
	// This is returned in API responses but is not used
	delete(customProperties, "__typename")

	user.CustomProperties = customProperties

	return &user, nil
}

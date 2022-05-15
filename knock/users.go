package knock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// UsersService is an interface for communicating with the PlanetScale
// Users API endpoints.
type UsersService interface {
	Get(context.Context, *GetUserRequest) (*User, error)
	Delete(context.Context, *DeleteUserRequest) error
	Identify(context.Context, *IdentifyUserRequest) (*User, error)
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
	ID               string                 `mapstructure:"id"`
	Name             string                 `mapstructure:"name"`
	Email            string                 `mapstructure:"email"`
	PhoneNumber      string                 `mapstructure:"phone_number"`
	Avatar           string                 `mapstructure:"avatar"`
	CreatedAt        time.Time              `mapstructure:"created_at"`
	UpdatedAt        time.Time              `mapstructure:"updated_at"`
	CustomProperties map[string]interface{} `mapstructure:",remain"`
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

func (us *usersService) Delete(ctx context.Context, deleteReq *DeleteUserRequest) error {
	path := usersAPIPath(deleteReq.ID)
	req, err := us.client.newRequest(http.MethodDelete, path, nil)

	if err != nil {
		return errors.Wrap(err, "error creating request for delete user")
	}

	err = us.client.do(ctx, req, nil)
	return err
}

func (us *usersService) Identify(ctx context.Context, identifyReq *IdentifyUserRequest) (*User, error) {
	path := usersAPIPath(identifyReq.ID)

	identifyBody := identifyReq.toMapWithCustomProperties()
	req, err := us.client.newRequest(http.MethodPut, path, identifyBody)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for identify user")
	}

	var identifyResponse map[string]interface{}
	err = us.client.do(ctx, req, identifyResponse)
	if err != nil {
		return nil, err
	}

	user, err := toUserWithCustomProperties(identifyResponse)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func usersAPIPath(userId string) string {
	return fmt.Sprintf("v1/users/%s", userId)
}

// toUserWithCustomProperties takes a map of keys returned from an API
// response and returns a User struct, correctly handling the remapping
// of custom properties that may be present in user data.
func toUserWithCustomProperties(input map[string]interface{}) (*User, error) {
	delete(input, "__typename")
	user := User{}

	stringToDateTimeHook := func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t == reflect.TypeOf(time.Time{}) && f == reflect.TypeOf("") {
			return time.Parse(time.RFC3339, data.(string))
		}

		return data, nil
	}

	config := mapstructure.DecoderConfig{
		DecodeHook: stringToDateTimeHook,
		Result:     &user,
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(input)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// IdentifyUserRequests can contain arbitrary customer data that must be stored, and must be mapped
// appropriately to one big flat list of key value pairs.
func (identifyReq *IdentifyUserRequest) toMapWithCustomProperties() map[string]interface{} {
	flatMap := make(map[string]interface{})

	// Note that marshalling and then immediately un-marshalling transforms struct keys
	// into the parameter names accepted by the API (i.e. PhoneNumber -> phone_number)
	data, _ := json.Marshal(identifyReq)
	json.Unmarshal(data, &flatMap)

	// Move all keys from a nested key in the map to the top level of the map
	for k, v := range flatMap["CustomProperties"].(map[string]interface{}) {
		flatMap[k] = v
	}
	delete(flatMap, "CustomProperties")

	return flatMap
}

func (us *usersService) Get(ctx context.Context, getReq *GetUserRequest) (*User, error) {
	path := usersAPIPath(getReq.ID)

	req, err := us.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get user")
	}

	var getResponse map[string]interface{}
	err = us.client.do(ctx, req, &getResponse)
	if err != nil {
		return nil, err
	}

	user, err := toUserWithCustomProperties(getResponse)
	if err != nil {
		return nil, err
	}

	return user, nil
}

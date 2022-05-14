package knock

import (
	"context"
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
	Phone_Number     string                 `mapstructure:"phone_number"`
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
	ID         string
	Properties map[string]interface{}
	CreatedAt  string
}

func BuildUserStruct(input map[string]interface{}) (*User, error) {
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

func (us *usersService) Get(ctx context.Context, getReq *GetUserRequest) (*User, error) {
	path := usersAPIPath(getReq.ID)

	req, err := us.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get user")
	}

	var userResultMap map[string]interface{}
	err = us.client.do(ctx, req, &userResultMap)
	fmt.Printf("result: %+v", userResultMap)

	if err != nil {
		return nil, err
	}

	user, err := BuildUserStruct(userResultMap)

	if err != nil {
		return nil, err
	}

	return user, nil
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

func usersAPIPath(userId string) string {
	return fmt.Sprintf("v1/users/%s", userId)
}

package knock

import (
	"context"
	"fmt"
	"net/http"
	"time"

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

type UserProperties map[string]interface{}

type User struct {
	ID         string `json:"id"`
	Properties UserProperties
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
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

func (us *usersService) Get(ctx context.Context, getReq *GetUserRequest) (*User, error) {
	path := usersAPIPath(getReq.ID)
	req, err := us.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get user")
	}

	user := &User{}
	err = us.client.do(ctx, req, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *usersService) Delete(ctx context.Context, deleteReq *DeleteUserRequest) error {
	path := usersAPIPath(deleteReq.ID)
	req, err := us.client.newRequest(http.MethodDelete, path, nil)

	if err != nil {
		return errors.Wrap(err, "error creating request for delete database")
	}

	err = us.client.do(ctx, req, nil)
	return err
}

func usersAPIPath(userId string) string {
	return fmt.Sprintf("v1/users/%s", userId)
}

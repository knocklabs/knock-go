package knock

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ChannelDataService interface {
	GetForUser(context.Context, *GetUserChannelDataRequest) (*GetUserChannelDataResponse, error)
	SetForUser(context.Context, *SetUserChannelDataRequest) (*SetUserChannelDataResponse, error)
	DeleteForUser(context.Context, *DeleteUserChannelDataRequest) error

	// GetForObject(context.Context, *GetObjectChannelDataRequest) (*GetObjectChannelDataResponse, error)
	// SetForObject(context.Context, *SetObjectChannelDataRequest) (*SetObjectChannelDataResponse, error)
	// DeleteForObject(context.Context, *DeleteObjectChannelDataRequest) (*DeleteObjectChannelDataResponse, error)
}

type channelDataService struct {
	client *Client
}

var _ ChannelDataService = &channelDataService{}

func NewChannelDataService(client *Client) *channelDataService {
	return &channelDataService{
		client: client,
	}
}

type ChannelData struct {
	ChannelID string                 `json:"channel_id"`
	Data      map[string]interface{} `json:"data"`
}

type GetUserChannelDataRequest struct {
	UserID    string
	ChannelID string
}
type DeleteUserChannelDataRequest = GetUserChannelDataRequest

type SetUserChannelDataRequest struct {
	UserID    string                 `json:"-"`
	ChannelID string                 `json:"-"`
	Data      map[string]interface{} `json:"data"`
}

type GetUserChannelDataResponse struct {
	ChannelData map[string]interface{} `json:"data"`
}
type SetUserChannelDataResponse = GetUserChannelDataResponse
type DeleteUserChannelDataResponse = GetUserChannelDataResponse

func usersChannelDataAPIPath(userID string, channelID string) string {
	return fmt.Sprintf("%s/channel_data/%s", UsersAPIPath(userID), channelID)
}

// func objectChannelDataAPIPath(collectionID string, objectID string, channelID string) string {
// 	return fmt.Sprintf("v1/object/%s/%s/channel_data/%s", collectionID, objectID, channelID)
// }

func (cds *channelDataService) GetForUser(ctx context.Context, getChannelDataReq *GetUserChannelDataRequest) (*GetUserChannelDataResponse, error) {
	path := usersChannelDataAPIPath(getChannelDataReq.UserID, getChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get user")
	}

	channelDataResponse := &GetUserChannelDataResponse{}
	_, err = cds.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get user")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for get user")
	}

	return channelDataResponse, nil
}

func (cds *channelDataService) SetForUser(ctx context.Context, getChannelDataReq *SetUserChannelDataRequest) (*SetUserChannelDataResponse, error) {
	path := usersChannelDataAPIPath(getChannelDataReq.UserID, getChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodPut, path, getChannelDataReq)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get user")
	}

	channelDataResponse := &SetUserChannelDataResponse{}
	_, err = cds.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get user")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for get user")
	}

	return channelDataResponse, nil
}

func (cds *channelDataService) DeleteForUser(ctx context.Context, deleteUserChannelDataReq *DeleteUserChannelDataRequest) error {
	path := usersChannelDataAPIPath(deleteUserChannelDataReq.UserID, deleteUserChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		errors.Wrap(err, "error creating request for get user")
	}

	_, err = cds.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request for get user")
	}

	if err != nil {
		return errors.Wrap(err, "error parsing request for get user")
	}

	return nil
}

func (cds *channelDataService) GetForObject(ctx context.Context, getChannelDataReq *GetUserChannelDataRequest) (*GetUserChannelDataResponse, error) {
	path := usersChannelDataAPIPath(getChannelDataReq.UserID, getChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get user")
	}

	channelDataResponse := &GetUserChannelDataResponse{}
	_, err = cds.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get user")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for  get user")
	}

	return channelDataResponse, nil
}

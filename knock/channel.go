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

	GetForObject(context.Context, *GetObjectChannelDataRequest) (*GetObjectChannelDataResponse, error)
	SetForObject(context.Context, *SetObjectChannelDataRequest) (*SetObjectChannelDataResponse, error)
	DeleteForObject(context.Context, *DeleteObjectChannelDataRequest) error
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

type GetObjectChannelDataRequest struct {
	Collection string
	ChannelID  string
	ObjectID   string
}
type GetObjectChannelDataResponse = GetUserChannelDataResponse

type SetObjectChannelDataRequest = GetObjectChannelDataRequest
type SetObjectChannelDataResponse = GetUserChannelDataResponse

type DeleteObjectChannelDataRequest = GetObjectChannelDataRequest

func usersChannelDataAPIPath(userID string, channelID string) string {
	return fmt.Sprintf("%s/channel_data/%s", UsersAPIPath(userID), channelID)
}

func objectChannelDataAPIPath(collection string, objectID string, channelID string) string {
	return fmt.Sprintf("v1/objects/%s/%s/channel_data/%s", collection, objectID, channelID)
}

func (cds *channelDataService) GetForUser(ctx context.Context, getChannelDataReq *GetUserChannelDataRequest) (*GetUserChannelDataResponse, error) {
	path := usersChannelDataAPIPath(getChannelDataReq.UserID, getChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for channel data for user")
	}

	channelDataResponse := &GetUserChannelDataResponse{}
	_, err = cds.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for channel data for user")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request for channel data for user")
	}

	return channelDataResponse, nil
}

func (cds *channelDataService) SetForUser(ctx context.Context, getChannelDataReq *SetUserChannelDataRequest) (*SetUserChannelDataResponse, error) {
	path := usersChannelDataAPIPath(getChannelDataReq.UserID, getChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodPut, path, getChannelDataReq)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set channel data for user")
	}

	channelDataResponse := &SetUserChannelDataResponse{}
	_, err = cds.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to set channel data for user")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to set channel data for user")
	}

	return channelDataResponse, nil
}

func (cds *channelDataService) DeleteForUser(ctx context.Context, deleteUserChannelDataReq *DeleteUserChannelDataRequest) error {
	path := usersChannelDataAPIPath(deleteUserChannelDataReq.UserID, deleteUserChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		errors.Wrap(err, "error creating request to delete channel data for a user")
	}

	_, err = cds.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request to delete channel data for a user")
	}

	if err != nil {
		return errors.Wrap(err, "error parsing request to delete channel data for a user")
	}

	return nil
}

func (cds *channelDataService) GetForObject(ctx context.Context, getChannelDataReq *GetObjectChannelDataRequest) (*GetObjectChannelDataResponse, error) {
	path := objectChannelDataAPIPath(getChannelDataReq.Collection, getChannelDataReq.ObjectID, getChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to get object channel data")
	}

	channelDataResponse := &GetUserChannelDataResponse{}
	_, err = cds.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to get object channel data")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to get object channel data")
	}

	return channelDataResponse, nil
}

func (cds *channelDataService) SetForObject(ctx context.Context, getChannelDataReq *SetObjectChannelDataRequest) (*SetObjectChannelDataResponse, error) {
	path := objectChannelDataAPIPath(getChannelDataReq.Collection, getChannelDataReq.ObjectID, getChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodPut, path, getChannelDataReq)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to set object channel data")
	}

	channelDataResponse := &SetObjectChannelDataResponse{}
	_, err = cds.client.do(ctx, req, channelDataResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request to set object channel data")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error parsing request to set object channel data")
	}

	return channelDataResponse, nil
}

func (cds *channelDataService) DeleteForObject(ctx context.Context, deleteObjectChannelDataReq *DeleteObjectChannelDataRequest) error {
	path := objectChannelDataAPIPath(deleteObjectChannelDataReq.Collection, deleteObjectChannelDataReq.ObjectID, deleteObjectChannelDataReq.ChannelID)

	req, err := cds.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		errors.Wrap(err, "error creating request to delete object channel data")
	}

	_, err = cds.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request to delete object channel data")
	}

	if err != nil {
		return errors.Wrap(err, "error parsing request to delete object channel data")
	}

	return nil
}

package knock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

// ProvidersService is an interface for communicating with the Knock
// Providers API endpoints.
type ProvidersService interface {
	AuthCheck(context.Context, *ProviderAuthCheckRequest) (*ProviderAuthCheckResponse, error)
	ListChannels(context.Context, *ProviderListChannelsRequest) (*ProviderListChannelsResponse, error)
	RevokeAccess(context.Context, *ProviderRevokeAccessRequest) (bool, error)
}

type providersService struct {
	client *Client
}

var _ ProvidersService = &providersService{}

func NewProvidersService(client *Client) *providersService {
	return &providersService{
		client: client,
	}
}

// Client structs
type ProviderContext struct {
	// ProviderName is included as a path parameter
	ProviderName string `json:"-"`
	// ChannelId is included as a path parameter
	ChannelId string `json:"-"`
}

type ProviderAccessTokenObject struct {
	ObjectId   string `json:"object_id"`
	Collection string `json:"collection"`
}

type ProviderAuthCheckRequest struct {
	ProviderContext   `json:"-"`
	AccessTokenObject ProviderAccessTokenObject `json:"access_token_object"`
}

type ProviderAuthCheckResponse struct {
	Ok     bool   `json:"ok"`
	Url    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamId string `json:"team_id"`
	UserId string `json:"user_id"`
	Error  string `json:"error,omitempty"`
}

type ProviderListChannelsRequest struct {
	ProviderContext   `json:"-"`
	AccessTokenObject ProviderAccessTokenObject `json:"access_token_object"`
	SlackQueryOptions *SlackQueryOptions        `json:"query_options,omitempty"`
}

type ProviderListChannelsResponse struct {
	SlackChannels []SlackChannel `json:"slack_channels"`
	NextCursor    string         `json:"next_cursor"`
}

type SlackQueryOptions struct {
	Cursor          string `json:"cursor,omitempty"`
	ExcludeArchived bool   `json:"exclude_archived,omitempty"`
	Limit           int    `json:"limit,omitempty"`
	TeamId          string `json:"team_id,omitempty"`
	Types           string `json:"types,omitempty"`
}

type SlackChannel struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	IsPrivate     bool   `json:"is_private"`
	IsIM          bool   `json:"is_im"`
	ContextTeamId string `json:"context_team_id"`
}

type ProviderRevokeAccessRequest struct {
	ProviderContext   `json:"-"`
	AccessTokenObject ProviderAccessTokenObject `json:"access_token_object"`
}

func providersAPIPath(providerName, channelId string) string {
	return fmt.Sprintf("v1/providers/%s/%s", providerName, channelId)
}

func (ps providersService) AuthCheck(ctx context.Context, request *ProviderAuthCheckRequest) (*ProviderAuthCheckResponse, error) {
	path := providersAPIPath(request.ProviderName, request.ChannelId)
	path = fmt.Sprintf("%s/auth_check", path)

	raw, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "error creating body for provider auth check")
	}

	req, err := ps.client.newRequest(http.MethodGet, path, bytes.NewBuffer(raw), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for provider auth check")
	}

	authCheckResponse := &ProviderAuthCheckResponse{}
	_, err = ps.client.do(ctx, req, authCheckResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for provider auth check")
	}

	return authCheckResponse, nil
}

func (ps providersService) ListChannels(ctx context.Context, request *ProviderListChannelsRequest) (*ProviderListChannelsResponse, error) {
	path := providersAPIPath(request.ProviderName, request.ChannelId)
	path = fmt.Sprintf("%s/channels", path)

	raw, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "error creating body for provider list channels")
	}

	req, err := ps.client.newRequest(http.MethodGet, path, bytes.NewBuffer(raw), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for provider list channels")
	}

	listChannelsResponse := &ProviderListChannelsResponse{}
	_, err = ps.client.do(ctx, req, listChannelsResponse)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for provider list channels")
	}

	return listChannelsResponse, nil
}

func (ps providersService) RevokeAccess(ctx context.Context, request *ProviderRevokeAccessRequest) (bool, error) {
	path := providersAPIPath(request.ProviderName, request.ChannelId)
	path = fmt.Sprintf("%s/revoke_access", path)

	raw, err := json.Marshal(request)
	if err != nil {
		return false, errors.Wrap(err, "error creating body for provider revoke access")
	}

	req, err := ps.client.newRequest(http.MethodGet, path, bytes.NewBuffer(raw), nil)
	if err != nil {
		return false, errors.Wrap(err, "error creating request for provider revoke access")
	}

	revokeAccessResponse := struct {
		Ok bool `json:"ok"`
	}{}
	_, err = ps.client.do(ctx, req, revokeAccessResponse)
	if err != nil {
		return false, errors.Wrap(err, "error making request for provider revoke access")
	}

	return revokeAccessResponse.Ok, nil
}

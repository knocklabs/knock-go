package knock

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/http/httputil"
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
type ProviderAccessTokenObject struct {
	ObjectId   string `json:"object_id" url:"object_id"`
	Collection string `json:"collection" url:"collection"`
}

type ProviderAuthCheckRequest struct {
	ProviderName      string                    `json:"-" url:"-"`
	ChannelId         string                    `json:"channel_id" url:"channel_id"`
	AccessTokenObject ProviderAccessTokenObject `json:"access_token_object" url:"access_token_object"`
}

type ProviderAuthCheckResponse struct {
	Connection ProviderAuthCheckResponseConnection `json:"connection"`
	Error      string                              `json:"error,omitempty"`
}

type ProviderAuthCheckResponseConnection struct {
	BotId               string `json:"bot_id"`
	IsEnterpriseInstall bool   `json:"is_enterprise_install"`
	Ok                  bool   `json:"ok"`
	Team                string `json:"team"`
	TeamId              string `json:"team_id"`
	Url                 string `json:"url"`
	User                string `json:"user"`
	UserId              string `json:"user_id"`
}

type ProviderListChannelsRequest struct {
	ProviderName      string                    `json:"-" url:"-"`
	ChannelId         string                    `url:"channel_id"`
	AccessTokenObject ProviderAccessTokenObject `url:"access_token_object"`
	SlackQueryOptions *SlackQueryOptions        `url:"query_options,omitempty"`
}

type ProviderListChannelsResponse struct {
	SlackChannels []SlackChannel `json:"slack_channels"`
	NextCursor    string         `json:"next_cursor"`
}

type SlackQueryOptions struct {
	Cursor          string `url:"cursor,omitempty"`
	ExcludeArchived bool   `url:"exclude_archived,omitempty"`
	Limit           int    `url:"limit,omitempty"`
	TeamId          string `url:"team_id,omitempty"`
	Types           string `url:"types,omitempty"`
}

type SlackChannel struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	IsPrivate     bool   `json:"is_private"`
	IsIM          bool   `json:"is_im"`
	ContextTeamId string `json:"context_team_id"`
}

type ProviderRevokeAccessRequest struct {
	ProviderName      string                    `json:"-" url:"-"`
	AccessTokenObject ProviderAccessTokenObject `json:"access_token_object" url:"access_token_object"`
	ChannelId         string                    `json:"channel_id" url:"channel_id"`
}

func providersAPIPath(providerName, channelId string) string {
	return fmt.Sprintf("v1/providers/%s/%s", providerName, channelId)
}

func (ps providersService) AuthCheck(ctx context.Context, request *ProviderAuthCheckRequest) (*ProviderAuthCheckResponse, error) {
	queryString, err := query.Values(request)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing query params to provider auth check")
	}
	path := fmt.Sprintf("%s/auth_check?%s", providersAPIPath(request.ProviderName, request.ChannelId), queryString.Encode())

	req, err := ps.client.newRequest(http.MethodGet, path, nil, nil)
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
	baseUrl := fmt.Sprintf("%s/channels", providersAPIPath(request.ProviderName, request.ChannelId))

	result := &ProviderListChannelsResponse{}
	for {
		queryString, err := query.Values(request)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing query params to provider list channels")
		}
		path := fmt.Sprintf("%s?%s", baseUrl, queryString.Encode())

		req, err := ps.client.newRequest(http.MethodGet, path, nil, nil)
		if err != nil {
			return nil, errors.Wrap(err, "error creating request for provider list channels")
		}

		pageResponse := &ProviderListChannelsResponse{}
		if _, err = ps.client.do(ctx, req, pageResponse); err != nil {
			return nil, errors.Wrap(err, "error making request for provider list channels")
		}

		result.SlackChannels = append(result.SlackChannels, pageResponse.SlackChannels...)
		if pageResponse.NextCursor == "" {
			return result, nil
		}
	}
}

func (ps providersService) RevokeAccess(ctx context.Context, request *ProviderRevokeAccessRequest) (bool, error) {
	queryString, err := query.Values(request)
	if err != nil {
		return false, errors.Wrap(err, "error parsing query params to revoke access")
	}
	path := fmt.Sprintf("%s/revoke_access?%s", providersAPIPath(request.ProviderName, request.ChannelId), queryString.Encode())

	req, err := ps.client.newRequest(http.MethodPut, path, nil, nil)
	if err != nil {
		return false, errors.Wrap(err, "error creating request for provider revoke access")
	}

	raw, _ := httputil.DumpRequest(req, true)
	log.Println(string(raw))

	revokeAccessResponse := struct {
		Ok bool `json:"ok"`
	}{}
	_, err = ps.client.do(ctx, req, revokeAccessResponse)
	if err != nil {
		return false, errors.Wrap(err, "error making request for provider revoke access")
	}

	return revokeAccessResponse.Ok, nil
}

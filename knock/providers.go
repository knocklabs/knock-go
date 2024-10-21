package knock

import (
	"context"
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
	// TODO: query_options
}

type ProviderListChannelsResponse struct {
	// TODO:
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
	path := providersAPIPath(request.ProviderName, request.ChannelId)
	path = fmt.Sprintf("%s/channels", path)

	//TODO implement me
	panic("implement me")
}

func (ps providersService) RevokeAccess(ctx context.Context, request *ProviderRevokeAccessRequest) (bool, error) {
	//TODO implement me
	panic("implement me")
}

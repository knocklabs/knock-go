// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/apiquery"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/packages/pagination"
)

// ProviderSlackService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderSlackService] method instead.
type ProviderSlackService struct {
	Options []option.RequestOption
}

// NewProviderSlackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderSlackService(opts ...option.RequestOption) (r *ProviderSlackService) {
	r = &ProviderSlackService{}
	r.Options = opts
	return
}

// Check if a Slack channel is authenticated.
func (r *ProviderSlackService) CheckAuth(ctx context.Context, channelID string, query ProviderSlackCheckAuthParams, opts ...option.RequestOption) (res *ProviderSlackCheckAuthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/slack/%s/auth_check", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List Slack channels for a Slack workspace.
func (r *ProviderSlackService) ListChannels(ctx context.Context, channelID string, query ProviderSlackListChannelsParams, opts ...option.RequestOption) (res *pagination.SlackChannelsCursor[ProviderSlackListChannelsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/slack/%s/channels", channelID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Slack channels for a Slack workspace.
func (r *ProviderSlackService) ListChannelsAutoPaging(ctx context.Context, channelID string, query ProviderSlackListChannelsParams, opts ...option.RequestOption) *pagination.SlackChannelsCursorAutoPager[ProviderSlackListChannelsResponse] {
	return pagination.NewSlackChannelsCursorAutoPager(r.ListChannels(ctx, channelID, query, opts...))
}

// Revoke access for a Slack channel.
func (r *ProviderSlackService) RevokeAccess(ctx context.Context, channelID string, body ProviderSlackRevokeAccessParams, opts ...option.RequestOption) (res *ProviderSlackRevokeAccessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/slack/%s/revoke_access", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// The response from a Slack auth check request.
type ProviderSlackCheckAuthResponse struct {
	// A Slack connection object.
	Connection ProviderSlackCheckAuthResponseConnection `json:"connection,required"`
	JSON       providerSlackCheckAuthResponseJSON       `json:"-"`
}

// providerSlackCheckAuthResponseJSON contains the JSON metadata for the struct
// [ProviderSlackCheckAuthResponse]
type providerSlackCheckAuthResponseJSON struct {
	Connection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderSlackCheckAuthResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerSlackCheckAuthResponseJSON) RawJSON() string {
	return r.raw
}

// A Slack connection object.
type ProviderSlackCheckAuthResponseConnection struct {
	// Whether the Slack connection is valid.
	Ok bool `json:"ok,required"`
	// The reason for the Slack connection if it is not valid.
	Reason string                                       `json:"reason,nullable"`
	JSON   providerSlackCheckAuthResponseConnectionJSON `json:"-"`
}

// providerSlackCheckAuthResponseConnectionJSON contains the JSON metadata for the
// struct [ProviderSlackCheckAuthResponseConnection]
type providerSlackCheckAuthResponseConnectionJSON struct {
	Ok          apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderSlackCheckAuthResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerSlackCheckAuthResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// A Slack channel.
type ProviderSlackListChannelsResponse struct {
	// A Slack channel ID from the Slack provider.
	ID string `json:"id,required"`
	// The team ID that the Slack channel belongs to.
	ContextTeamID string `json:"context_team_id,required"`
	// Whether the Slack channel is an IM channel.
	IsIm bool `json:"is_im,required"`
	// Whether the Slack channel is private.
	IsPrivate bool `json:"is_private,required"`
	// Slack channel name.
	Name string                                `json:"name,required"`
	JSON providerSlackListChannelsResponseJSON `json:"-"`
}

// providerSlackListChannelsResponseJSON contains the JSON metadata for the struct
// [ProviderSlackListChannelsResponse]
type providerSlackListChannelsResponseJSON struct {
	ID            apijson.Field
	ContextTeamID apijson.Field
	IsIm          apijson.Field
	IsPrivate     apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProviderSlackListChannelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerSlackListChannelsResponseJSON) RawJSON() string {
	return r.raw
}

// A response indicating the operation was successful.
type ProviderSlackRevokeAccessResponse struct {
	// OK response.
	Ok   string                                `json:"ok"`
	JSON providerSlackRevokeAccessResponseJSON `json:"-"`
}

// providerSlackRevokeAccessResponseJSON contains the JSON metadata for the struct
// [ProviderSlackRevokeAccessResponse]
type providerSlackRevokeAccessResponseJSON struct {
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderSlackRevokeAccessResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerSlackRevokeAccessResponseJSON) RawJSON() string {
	return r.raw
}

type ProviderSlackCheckAuthParams struct {
	// A JSON encoded string containing the access token object reference.
	AccessTokenObject param.Field[string] `query:"access_token_object,required"`
}

// URLQuery serializes [ProviderSlackCheckAuthParams]'s query parameters as
// `url.Values`.
func (r ProviderSlackCheckAuthParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderSlackListChannelsParams struct {
	// A JSON encoded string containing the access token object reference.
	AccessTokenObject param.Field[string]                                      `query:"access_token_object,required"`
	QueryOptions      param.Field[ProviderSlackListChannelsParamsQueryOptions] `query:"query_options"`
}

// URLQuery serializes [ProviderSlackListChannelsParams]'s query parameters as
// `url.Values`.
func (r ProviderSlackListChannelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderSlackListChannelsParamsQueryOptions struct {
	// Paginate through collections of data by setting the cursor parameter to a
	// next_cursor attribute returned by a previous request's response_metadata.
	// Default value fetches the first "page" of the collection.
	Cursor param.Field[string] `query:"cursor"`
	// Set to true to exclude archived channels from the list. Defaults to `true` when
	// not explicitly provided.
	ExcludeArchived param.Field[bool] `query:"exclude_archived"`
	// The maximum number of channels to return. Defaults to 200.
	Limit param.Field[int64] `query:"limit"`
	// Encoded team ID (T1234) to list channels in, required if org token is used.
	TeamID param.Field[string] `query:"team_id"`
	// Mix and match channel types by providing a comma-separated list of any
	// combination of public_channel, private_channel, mpim, im. Defaults to
	// `"public_channel,private_channel"`. If the user's Slack ID is unavailable, this
	// option is ignored and only public channels are returned.
	Types param.Field[string] `query:"types"`
}

// URLQuery serializes [ProviderSlackListChannelsParamsQueryOptions]'s query
// parameters as `url.Values`.
func (r ProviderSlackListChannelsParamsQueryOptions) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderSlackRevokeAccessParams struct {
	// A JSON encoded string containing the access token object reference.
	AccessTokenObject param.Field[string] `query:"access_token_object,required"`
}

// URLQuery serializes [ProviderSlackRevokeAccessParams]'s query parameters as
// `url.Values`.
func (r ProviderSlackRevokeAccessParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

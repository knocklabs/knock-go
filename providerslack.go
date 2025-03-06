// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
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

// Check if a Slack channel is authenticated
func (r *ProviderSlackService) CheckAuth(ctx context.Context, channelID string, query ProviderSlackCheckAuthParams, opts ...option.RequestOption) (res *ProviderSlackCheckAuthResponse, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/slack/%s/auth_check", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Slack channels from a Slack workspace
func (r *ProviderSlackService) ListChannels(ctx context.Context, channelID string, query ProviderSlackListChannelsParams, opts ...option.RequestOption) (res *ProviderSlackListChannelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/slack/%s/channels", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Revoke access for a Slack channel
func (r *ProviderSlackService) RevokeAccess(ctx context.Context, channelID string, body ProviderSlackRevokeAccessParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/slack/%s/revoke_access", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// The response from a Slack auth check request
type ProviderSlackCheckAuthResponse struct {
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

type ProviderSlackCheckAuthResponseConnection struct {
	Ok     bool                                         `json:"ok,required"`
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

// The response from a Slack channels for provider request
type ProviderSlackListChannelsResponse struct {
	NextCursor    string                                          `json:"next_cursor,required,nullable"`
	SlackChannels []ProviderSlackListChannelsResponseSlackChannel `json:"slack_channels,required"`
	JSON          providerSlackListChannelsResponseJSON           `json:"-"`
}

// providerSlackListChannelsResponseJSON contains the JSON metadata for the struct
// [ProviderSlackListChannelsResponse]
type providerSlackListChannelsResponseJSON struct {
	NextCursor    apijson.Field
	SlackChannels apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProviderSlackListChannelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerSlackListChannelsResponseJSON) RawJSON() string {
	return r.raw
}

type ProviderSlackListChannelsResponseSlackChannel struct {
	ID            string                                            `json:"id,required"`
	ContextTeamID string                                            `json:"context_team_id,required"`
	IsIm          bool                                              `json:"is_im,required"`
	IsPrivate     bool                                              `json:"is_private,required"`
	Name          string                                            `json:"name,required"`
	JSON          providerSlackListChannelsResponseSlackChannelJSON `json:"-"`
}

// providerSlackListChannelsResponseSlackChannelJSON contains the JSON metadata for
// the struct [ProviderSlackListChannelsResponseSlackChannel]
type providerSlackListChannelsResponseSlackChannelJSON struct {
	ID            apijson.Field
	ContextTeamID apijson.Field
	IsIm          apijson.Field
	IsPrivate     apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProviderSlackListChannelsResponseSlackChannel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerSlackListChannelsResponseSlackChannelJSON) RawJSON() string {
	return r.raw
}

type ProviderSlackCheckAuthParams struct {
	// A JSON encoded string containing the access token object reference
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
	// A JSON encoded string containing the access token object reference
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
	// A cursor to paginate through the channels
	Cursor param.Field[string] `query:"cursor"`
	// Whether to exclude archived channels
	ExcludeArchived param.Field[string] `query:"exclude_archived"`
	// The number of channels to return
	Limit param.Field[string] `query:"limit"`
	// The ID of the Slack team to get channels for
	TeamID param.Field[string] `query:"team_id"`
	// The types of channels to return
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
	// A JSON encoded string containing the access token object reference
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

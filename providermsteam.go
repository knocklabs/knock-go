// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/apiquery"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/packages/pagination"
)

// ProviderMsTeamService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderMsTeamService] method instead.
type ProviderMsTeamService struct {
	Options []option.RequestOption
}

// NewProviderMsTeamService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderMsTeamService(opts ...option.RequestOption) (r *ProviderMsTeamService) {
	r = &ProviderMsTeamService{}
	r.Options = opts
	return
}

// Check if a connection to Microsoft Teams has been authorized for a given
// Microsoft Teams tenant object.
func (r *ProviderMsTeamService) CheckAuth(ctx context.Context, channelID string, query ProviderMsTeamCheckAuthParams, opts ...option.RequestOption) (res *ProviderMsTeamCheckAuthResponse, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/ms-teams/%s/auth_check", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List the Microsoft Teams channels within a team. By default, archived and
// private channels are excluded from the results.
func (r *ProviderMsTeamService) ListChannels(ctx context.Context, channelID string, query ProviderMsTeamListChannelsParams, opts ...option.RequestOption) (res *ProviderMsTeamListChannelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/ms-teams/%s/channels", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a list of teams belonging to the Microsoft Entra tenant. By default,
// archived and private channels are excluded from the results.
func (r *ProviderMsTeamService) ListTeams(ctx context.Context, channelID string, query ProviderMsTeamListTeamsParams, opts ...option.RequestOption) (res *pagination.MsTeamsPagination[ProviderMsTeamListTeamsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/ms-teams/%s/teams", channelID)
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

// Get a list of teams belonging to the Microsoft Entra tenant. By default,
// archived and private channels are excluded from the results.
func (r *ProviderMsTeamService) ListTeamsAutoPaging(ctx context.Context, channelID string, query ProviderMsTeamListTeamsParams, opts ...option.RequestOption) *pagination.MsTeamsPaginationAutoPager[ProviderMsTeamListTeamsResponse] {
	return pagination.NewMsTeamsPaginationAutoPager(r.ListTeams(ctx, channelID, query, opts...))
}

// Remove a Microsoft Entra tenant ID from a Microsoft Teams tenant object.
func (r *ProviderMsTeamService) RevokeAccess(ctx context.Context, channelID string, body ProviderMsTeamRevokeAccessParams, opts ...option.RequestOption) (res *ProviderMsTeamRevokeAccessResponse, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/providers/ms-teams/%s/revoke_access", channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// The response from a Microsoft Teams auth check request.
type ProviderMsTeamCheckAuthResponse struct {
	// A Microsoft Teams connection object.
	Connection ProviderMsTeamCheckAuthResponseConnection `json:"connection,required"`
	JSON       providerMsTeamCheckAuthResponseJSON       `json:"-"`
}

// providerMsTeamCheckAuthResponseJSON contains the JSON metadata for the struct
// [ProviderMsTeamCheckAuthResponse]
type providerMsTeamCheckAuthResponseJSON struct {
	Connection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderMsTeamCheckAuthResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerMsTeamCheckAuthResponseJSON) RawJSON() string {
	return r.raw
}

// A Microsoft Teams connection object.
type ProviderMsTeamCheckAuthResponseConnection struct {
	// Whether the Microsoft Teams connection is valid.
	Ok bool `json:"ok,required"`
	// The reason for the Microsoft Teams connection if it is not valid.
	Reason string                                        `json:"reason,nullable"`
	JSON   providerMsTeamCheckAuthResponseConnectionJSON `json:"-"`
}

// providerMsTeamCheckAuthResponseConnectionJSON contains the JSON metadata for the
// struct [ProviderMsTeamCheckAuthResponseConnection]
type providerMsTeamCheckAuthResponseConnectionJSON struct {
	Ok          apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderMsTeamCheckAuthResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerMsTeamCheckAuthResponseConnectionJSON) RawJSON() string {
	return r.raw
}

// The response from a Microsoft Teams provider request, containing a list of
// channels.
type ProviderMsTeamListChannelsResponse struct {
	// List of Microsoft Teams channels.
	MsTeamsChannels []ProviderMsTeamListChannelsResponseMsTeamsChannel `json:"ms_teams_channels,required"`
	JSON            providerMsTeamListChannelsResponseJSON             `json:"-"`
}

// providerMsTeamListChannelsResponseJSON contains the JSON metadata for the struct
// [ProviderMsTeamListChannelsResponse]
type providerMsTeamListChannelsResponseJSON struct {
	MsTeamsChannels apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderMsTeamListChannelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerMsTeamListChannelsResponseJSON) RawJSON() string {
	return r.raw
}

type ProviderMsTeamListChannelsResponseMsTeamsChannel struct {
	// Microsoft Teams channel ID.
	ID string `json:"id,required"`
	// Microsoft Teams channel name.
	DisplayName string `json:"displayName,required"`
	// Microsoft Teams channel created date and time.
	CreatedDateTime string `json:"createdDateTime"`
	// Microsoft Teams channel description.
	Description string `json:"description,nullable"`
	// Whether the Microsoft Teams channel is archived.
	IsArchived bool `json:"isArchived"`
	// Microsoft Teams channel membership type.
	MembershipType string                                               `json:"membershipType"`
	JSON           providerMsTeamListChannelsResponseMsTeamsChannelJSON `json:"-"`
}

// providerMsTeamListChannelsResponseMsTeamsChannelJSON contains the JSON metadata
// for the struct [ProviderMsTeamListChannelsResponseMsTeamsChannel]
type providerMsTeamListChannelsResponseMsTeamsChannelJSON struct {
	ID              apijson.Field
	DisplayName     apijson.Field
	CreatedDateTime apijson.Field
	Description     apijson.Field
	IsArchived      apijson.Field
	MembershipType  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderMsTeamListChannelsResponseMsTeamsChannel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerMsTeamListChannelsResponseMsTeamsChannelJSON) RawJSON() string {
	return r.raw
}

type ProviderMsTeamListTeamsResponse struct {
	// Microsoft Teams team ID.
	ID string `json:"id,required"`
	// Microsoft Teams team display name.
	DisplayName string `json:"displayName,required"`
	// Microsoft Teams team description.
	Description string                              `json:"description,nullable"`
	JSON        providerMsTeamListTeamsResponseJSON `json:"-"`
}

// providerMsTeamListTeamsResponseJSON contains the JSON metadata for the struct
// [ProviderMsTeamListTeamsResponse]
type providerMsTeamListTeamsResponseJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderMsTeamListTeamsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerMsTeamListTeamsResponseJSON) RawJSON() string {
	return r.raw
}

// A response indicating the operation was successful.
type ProviderMsTeamRevokeAccessResponse struct {
	// OK response.
	Ok   string                                 `json:"ok"`
	JSON providerMsTeamRevokeAccessResponseJSON `json:"-"`
}

// providerMsTeamRevokeAccessResponseJSON contains the JSON metadata for the struct
// [ProviderMsTeamRevokeAccessResponse]
type providerMsTeamRevokeAccessResponseJSON struct {
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderMsTeamRevokeAccessResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerMsTeamRevokeAccessResponseJSON) RawJSON() string {
	return r.raw
}

type ProviderMsTeamCheckAuthParams struct {
	// A JSON encoded string containing the Microsoft Teams tenant object reference.
	MsTeamsTenantObject param.Field[string] `query:"ms_teams_tenant_object,required"`
}

// URLQuery serializes [ProviderMsTeamCheckAuthParams]'s query parameters as
// `url.Values`.
func (r ProviderMsTeamCheckAuthParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderMsTeamListChannelsParams struct {
	// A JSON encoded string containing the Microsoft Teams tenant object reference.
	MsTeamsTenantObject param.Field[string] `query:"ms_teams_tenant_object,required"`
	// Microsoft Teams team ID.
	TeamID       param.Field[string]                                       `query:"team_id,required"`
	QueryOptions param.Field[ProviderMsTeamListChannelsParamsQueryOptions] `query:"query_options"`
}

// URLQuery serializes [ProviderMsTeamListChannelsParams]'s query parameters as
// `url.Values`.
func (r ProviderMsTeamListChannelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderMsTeamListChannelsParamsQueryOptions struct {
	// [OData param](https://learn.microsoft.com/en-us/graph/query-parameters) passed
	// to the Microsoft Graph API to filter channels.
	Filter param.Field[string] `query:"$filter"`
	// [OData param](https://learn.microsoft.com/en-us/graph/query-parameters) passed
	// to the Microsoft Graph API to select specific properties.
	Select param.Field[string] `query:"$select"`
}

// URLQuery serializes [ProviderMsTeamListChannelsParamsQueryOptions]'s query
// parameters as `url.Values`.
func (r ProviderMsTeamListChannelsParamsQueryOptions) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderMsTeamListTeamsParams struct {
	// A JSON encoded string containing the Microsoft Teams tenant object reference.
	MsTeamsTenantObject param.Field[string]                                    `query:"ms_teams_tenant_object,required"`
	QueryOptions        param.Field[ProviderMsTeamListTeamsParamsQueryOptions] `query:"query_options"`
}

// URLQuery serializes [ProviderMsTeamListTeamsParams]'s query parameters as
// `url.Values`.
func (r ProviderMsTeamListTeamsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderMsTeamListTeamsParamsQueryOptions struct {
	// [OData param](https://learn.microsoft.com/en-us/graph/query-parameters) passed
	// to the Microsoft Graph API to filter teams.
	Filter param.Field[string] `query:"$filter"`
	// [OData param](https://learn.microsoft.com/en-us/graph/query-parameters) passed
	// to the Microsoft Graph API to select fields on a team.
	Select param.Field[string] `query:"$select"`
	// [OData param](https://learn.microsoft.com/en-us/graph/query-parameters) passed
	// to the Microsoft Graph API to retrieve the next page of results.
	Skiptoken param.Field[string] `query:"$skiptoken"`
	// [OData param](https://learn.microsoft.com/en-us/graph/query-parameters) passed
	// to the Microsoft Graph API to limit the number of teams returned.
	Top param.Field[int64] `query:"$top"`
}

// URLQuery serializes [ProviderMsTeamListTeamsParamsQueryOptions]'s query
// parameters as `url.Values`.
func (r ProviderMsTeamListTeamsParamsQueryOptions) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderMsTeamRevokeAccessParams struct {
	// A JSON encoded string containing the Microsoft Teams tenant object reference.
	MsTeamsTenantObject param.Field[string] `query:"ms_teams_tenant_object,required"`
}

// URLQuery serializes [ProviderMsTeamRevokeAccessParams]'s query parameters as
// `url.Values`.
func (r ProviderMsTeamRevokeAccessParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
)

// TenantBulkService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantBulkService] method instead.
type TenantBulkService struct {
	Options []option.RequestOption
}

// NewTenantBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTenantBulkService(opts ...option.RequestOption) (r *TenantBulkService) {
	r = &TenantBulkService{}
	r.Options = opts
	return
}

// Bulk delete tenants
func (r *TenantBulkService) Delete(ctx context.Context, body TenantBulkDeleteParams, opts ...option.RequestOption) (res *TenantBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tenants/bulk/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk set tenants
func (r *TenantBulkService) Set(ctx context.Context, body TenantBulkSetParams, opts ...option.RequestOption) (res *TenantBulkSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tenants/bulk/set"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A bulk operation entity
type TenantBulkDeleteResponse struct {
	ID                 string                         `json:"id,required" format:"uuid"`
	Typename           string                         `json:"__typename,required"`
	EstimatedTotalRows int64                          `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                      `json:"inserted_at,required" format:"date-time"`
	Name               string                         `json:"name,required"`
	ProcessedRows      int64                          `json:"processed_rows,required"`
	Status             TenantBulkDeleteResponseStatus `json:"status,required"`
	SuccessCount       int64                          `json:"success_count,required"`
	UpdatedAt          time.Time                      `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                      `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                          `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []TenantBulkDeleteResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                           `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                           `json:"started_at,nullable" format:"date-time"`
	JSON       tenantBulkDeleteResponseJSON        `json:"-"`
}

// tenantBulkDeleteResponseJSON contains the JSON metadata for the struct
// [TenantBulkDeleteResponse]
type tenantBulkDeleteResponseJSON struct {
	ID                 apijson.Field
	Typename           apijson.Field
	EstimatedTotalRows apijson.Field
	InsertedAt         apijson.Field
	Name               apijson.Field
	ProcessedRows      apijson.Field
	Status             apijson.Field
	SuccessCount       apijson.Field
	UpdatedAt          apijson.Field
	CompletedAt        apijson.Field
	ErrorCount         apijson.Field
	ErrorItems         apijson.Field
	FailedAt           apijson.Field
	StartedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TenantBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TenantBulkDeleteResponseStatus string

const (
	TenantBulkDeleteResponseStatusQueued     TenantBulkDeleteResponseStatus = "queued"
	TenantBulkDeleteResponseStatusProcessing TenantBulkDeleteResponseStatus = "processing"
	TenantBulkDeleteResponseStatusCompleted  TenantBulkDeleteResponseStatus = "completed"
	TenantBulkDeleteResponseStatusFailed     TenantBulkDeleteResponseStatus = "failed"
)

func (r TenantBulkDeleteResponseStatus) IsKnown() bool {
	switch r {
	case TenantBulkDeleteResponseStatusQueued, TenantBulkDeleteResponseStatusProcessing, TenantBulkDeleteResponseStatusCompleted, TenantBulkDeleteResponseStatusFailed:
		return true
	}
	return false
}

type TenantBulkDeleteResponseErrorItem struct {
	ID         string                                `json:"id,required"`
	Collection string                                `json:"collection,nullable"`
	JSON       tenantBulkDeleteResponseErrorItemJSON `json:"-"`
}

// tenantBulkDeleteResponseErrorItemJSON contains the JSON metadata for the struct
// [TenantBulkDeleteResponseErrorItem]
type tenantBulkDeleteResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantBulkDeleteResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantBulkDeleteResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type TenantBulkSetResponse struct {
	ID                 string                      `json:"id,required" format:"uuid"`
	Typename           string                      `json:"__typename,required"`
	EstimatedTotalRows int64                       `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                   `json:"inserted_at,required" format:"date-time"`
	Name               string                      `json:"name,required"`
	ProcessedRows      int64                       `json:"processed_rows,required"`
	Status             TenantBulkSetResponseStatus `json:"status,required"`
	SuccessCount       int64                       `json:"success_count,required"`
	UpdatedAt          time.Time                   `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                   `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                       `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []TenantBulkSetResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                        `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                        `json:"started_at,nullable" format:"date-time"`
	JSON       tenantBulkSetResponseJSON        `json:"-"`
}

// tenantBulkSetResponseJSON contains the JSON metadata for the struct
// [TenantBulkSetResponse]
type tenantBulkSetResponseJSON struct {
	ID                 apijson.Field
	Typename           apijson.Field
	EstimatedTotalRows apijson.Field
	InsertedAt         apijson.Field
	Name               apijson.Field
	ProcessedRows      apijson.Field
	Status             apijson.Field
	SuccessCount       apijson.Field
	UpdatedAt          apijson.Field
	CompletedAt        apijson.Field
	ErrorCount         apijson.Field
	ErrorItems         apijson.Field
	FailedAt           apijson.Field
	StartedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TenantBulkSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantBulkSetResponseJSON) RawJSON() string {
	return r.raw
}

type TenantBulkSetResponseStatus string

const (
	TenantBulkSetResponseStatusQueued     TenantBulkSetResponseStatus = "queued"
	TenantBulkSetResponseStatusProcessing TenantBulkSetResponseStatus = "processing"
	TenantBulkSetResponseStatusCompleted  TenantBulkSetResponseStatus = "completed"
	TenantBulkSetResponseStatusFailed     TenantBulkSetResponseStatus = "failed"
)

func (r TenantBulkSetResponseStatus) IsKnown() bool {
	switch r {
	case TenantBulkSetResponseStatusQueued, TenantBulkSetResponseStatusProcessing, TenantBulkSetResponseStatusCompleted, TenantBulkSetResponseStatusFailed:
		return true
	}
	return false
}

type TenantBulkSetResponseErrorItem struct {
	ID         string                             `json:"id,required"`
	Collection string                             `json:"collection,nullable"`
	JSON       tenantBulkSetResponseErrorItemJSON `json:"-"`
}

// tenantBulkSetResponseErrorItemJSON contains the JSON metadata for the struct
// [TenantBulkSetResponseErrorItem]
type tenantBulkSetResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantBulkSetResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantBulkSetResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

type TenantBulkDeleteParams struct {
	// The IDs of the tenants to delete
	TenantIDs param.Field[[]string] `query:"tenant_ids,required"`
}

// URLQuery serializes [TenantBulkDeleteParams]'s query parameters as `url.Values`.
func (r TenantBulkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantBulkSetParams struct {
	Tenants param.Field[[]TenantBulkSetParamsTenantUnion] `json:"tenants,required"`
}

func (r TenantBulkSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An inline tenant request
//
// Satisfied by [shared.UnionString], [TenantBulkSetParamsTenantsTenantRequest].
type TenantBulkSetParamsTenantUnion interface {
	ImplementsTenantBulkSetParamsTenantUnion()
}

// A tenant to be set in the system
type TenantBulkSetParamsTenantsTenantRequest struct {
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]TenantBulkSetParamsTenantsTenantRequestChannelData] `json:"channel_data"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]TenantBulkSetParamsTenantsTenantRequestPreferences] `json:"preferences"`
	Settings    param.Field[TenantBulkSetParamsTenantsTenantRequestSettings]               `json:"settings"`
	ExtraFields map[string]interface{}                                                     `json:"-,extras"`
}

func (r TenantBulkSetParamsTenantsTenantRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequest) ImplementsTenantBulkSetParamsTenantUnion() {}

// Set channel data for a type of channel
type TenantBulkSetParamsTenantsTenantRequestChannelData struct {
	// Channel data for push providers
	Data param.Field[TenantBulkSetParamsTenantsTenantRequestChannelDataDataUnion] `json:"data,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type TenantBulkSetParamsTenantsTenantRequestChannelDataData struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataData) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataUnion() {
}

// Channel data for push providers
//
// Satisfied by
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataPushChannelData],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataOneSignalChannelData],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelData],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelData],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelData],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataData].
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataUnion interface {
	implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataUnion()
}

// Channel data for push providers
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataPushChannelData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataPushChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataPushChannelData) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataUnion() {
}

// OneSignal channel data
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataOneSignalChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataOneSignalChannelData) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataUnion() {
}

// Slack channel data
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelData struct {
	Connections param.Field[[]TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionUnion] `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token param.Field[TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataToken] `json:"token"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelData) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnection) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnection].
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataConnectionUnion() {
}

// A token that's used to store the access token for a Slack workspace.
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataToken struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataSlackChannelDataToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelData struct {
	Connections param.Field[[]TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionUnion] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelData) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataUnion() {
}

// Microsoft Teams token connection
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnection struct {
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnection) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams token connection
//
// Satisfied by
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnection].
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// Microsoft Teams token connection
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams incoming webhook connection
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// The incoming webhook
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Discord channel data
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelData struct {
	Connections param.Field[[]TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionUnion] `json:"connections,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelData) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataUnion() {
}

// Discord channel connection
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnection) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection],
// [TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnection].
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord incoming webhook connection
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) implementsTenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionUnion() {
}

// The incoming webhook
type TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type TenantBulkSetParamsTenantsTenantRequestPreferences struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesUnion() {
}

// Channel type preferences
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                        `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                        `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                         `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                         `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                        `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                        `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                             `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                        `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                        `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                       `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                   `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                   `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                    `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                    `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                     `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                     `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                    `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                    `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                         `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                         `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                    `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                    `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                   `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                   `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsUnion() {
}

// Channel type preferences
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                       `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                        `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                        `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                       `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                            `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                       `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                      `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                      `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                  `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                  `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

type TenantBulkSetParamsTenantsTenantRequestSettings struct {
	Branding param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsBranding] `json:"branding"`
	// Set preferences for a recipient
	PreferenceSet param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSet] `json:"preference_set"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsBranding struct {
	IconURL              param.Field[string] `json:"icon_url"`
	LogoURL              param.Field[string] `json:"logo_url"`
	PrimaryColor         param.Field[string] `json:"primary_color"`
	PrimaryColorContrast param.Field[string] `json:"primary_color_contrast"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsBranding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSet struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsUnion] `json:"workflows"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesUnion() {
}

// Channel type preferences
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                  `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                   `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                   `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                  `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                       `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                  `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                 `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                             `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                             `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                              `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                              `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                               `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                               `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                              `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                              `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                   `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                   `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                              `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                              `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                             `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                             `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsUnion() {
}

// Channel type preferences
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                 `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                  `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                 `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                      `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                      `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                 `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                            `json:"argument,required"`
	Operator param.Field[TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                            `json:"variable,required"`
}

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, TenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

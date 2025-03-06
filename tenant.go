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

// TenantService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantService] method instead.
type TenantService struct {
	Options []option.RequestOption
	Bulk    *TenantBulkService
}

// NewTenantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenantService(opts ...option.RequestOption) (r *TenantService) {
	r = &TenantService{}
	r.Options = opts
	r.Bulk = NewTenantBulkService(opts...)
	return
}

// List tenants
func (r *TenantService) List(ctx context.Context, query TenantListParams, opts ...option.RequestOption) (res *TenantListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tenants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a tenant
func (r *TenantService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tenants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a tenant
func (r *TenantService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TenantGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tenants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set a tenant
func (r *TenantService) Set(ctx context.Context, id string, body TenantSetParams, opts ...option.RequestOption) (res *TenantSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tenants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// A paginated list of tenants.
type TenantListResponse struct {
	// The list of tenants
	Entries []TenantListResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo TenantListResponsePageInfo `json:"page_info,required"`
	JSON     tenantListResponseJSON     `json:"-"`
}

// tenantListResponseJSON contains the JSON metadata for the struct
// [TenantListResponse]
type tenantListResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseJSON) RawJSON() string {
	return r.raw
}

// A tenant entity
type TenantListResponseEntry struct {
	ID          string                      `json:"id,required"`
	Typename    string                      `json:"__typename,required"`
	ExtraFields map[string]interface{}      `json:"-,extras"`
	JSON        tenantListResponseEntryJSON `json:"-"`
}

// tenantListResponseEntryJSON contains the JSON metadata for the struct
// [TenantListResponseEntry]
type tenantListResponseEntryJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseEntryJSON) RawJSON() string {
	return r.raw
}

// The information about a paginated result
type TenantListResponsePageInfo struct {
	Typename string                         `json:"__typename,required"`
	PageSize int64                          `json:"page_size,required"`
	After    string                         `json:"after,nullable"`
	Before   string                         `json:"before,nullable"`
	JSON     tenantListResponsePageInfoJSON `json:"-"`
}

// tenantListResponsePageInfoJSON contains the JSON metadata for the struct
// [TenantListResponsePageInfo]
type tenantListResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

// A tenant entity
type TenantGetResponse struct {
	ID          string                 `json:"id,required"`
	Typename    string                 `json:"__typename,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        tenantGetResponseJSON  `json:"-"`
}

// tenantGetResponseJSON contains the JSON metadata for the struct
// [TenantGetResponse]
type tenantGetResponseJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantGetResponseJSON) RawJSON() string {
	return r.raw
}

// A tenant entity
type TenantSetResponse struct {
	ID          string                 `json:"id,required"`
	Typename    string                 `json:"__typename,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        tenantSetResponseJSON  `json:"-"`
}

// tenantSetResponseJSON contains the JSON metadata for the struct
// [TenantSetResponse]
type tenantSetResponseJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantSetResponseJSON) RawJSON() string {
	return r.raw
}

type TenantListParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [TenantListParams]'s query parameters as `url.Values`.
func (r TenantListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantSetParams struct {
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]TenantSetParamsChannelData] `json:"channel_data"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]TenantSetParamsPreferences] `json:"preferences"`
	Settings    param.Field[TenantSetParamsSettings]               `json:"settings"`
}

func (r TenantSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set channel data for a type of channel
type TenantSetParamsChannelData struct {
	// Channel data for push providers
	Data param.Field[TenantSetParamsChannelDataDataUnion] `json:"data,required"`
}

func (r TenantSetParamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type TenantSetParamsChannelDataData struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r TenantSetParamsChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataData) implementsTenantSetParamsChannelDataDataUnion() {}

// Channel data for push providers
//
// Satisfied by [TenantSetParamsChannelDataDataPushChannelData],
// [TenantSetParamsChannelDataDataOneSignalChannelData],
// [TenantSetParamsChannelDataDataSlackChannelData],
// [TenantSetParamsChannelDataDataMsTeamsChannelData],
// [TenantSetParamsChannelDataDataDiscordChannelData],
// [TenantSetParamsChannelDataData].
type TenantSetParamsChannelDataDataUnion interface {
	implementsTenantSetParamsChannelDataDataUnion()
}

// Channel data for push providers
type TenantSetParamsChannelDataDataPushChannelData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r TenantSetParamsChannelDataDataPushChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataPushChannelData) implementsTenantSetParamsChannelDataDataUnion() {
}

// OneSignal channel data
type TenantSetParamsChannelDataDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r TenantSetParamsChannelDataDataOneSignalChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataOneSignalChannelData) implementsTenantSetParamsChannelDataDataUnion() {
}

// Slack channel data
type TenantSetParamsChannelDataDataSlackChannelData struct {
	Connections param.Field[[]TenantSetParamsChannelDataDataSlackChannelDataConnectionUnion] `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token param.Field[TenantSetParamsChannelDataDataSlackChannelDataToken] `json:"token"`
}

func (r TenantSetParamsChannelDataDataSlackChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataSlackChannelData) implementsTenantSetParamsChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type TenantSetParamsChannelDataDataSlackChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnection) implementsTenantSetParamsChannelDataDataSlackChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [TenantSetParamsChannelDataDataSlackChannelDataConnectionsSlackTokenConnection],
// [TenantSetParamsChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection],
// [TenantSetParamsChannelDataDataSlackChannelDataConnection].
type TenantSetParamsChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsTenantSetParamsChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type TenantSetParamsChannelDataDataSlackChannelDataConnectionsSlackTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) implementsTenantSetParamsChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type TenantSetParamsChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) implementsTenantSetParamsChannelDataDataSlackChannelDataConnectionUnion() {
}

// A token that's used to store the access token for a Slack workspace.
type TenantSetParamsChannelDataDataSlackChannelDataToken struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r TenantSetParamsChannelDataDataSlackChannelDataToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type TenantSetParamsChannelDataDataMsTeamsChannelData struct {
	Connections param.Field[[]TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelData) implementsTenantSetParamsChannelDataDataUnion() {
}

// Microsoft Teams token connection
type TenantSetParamsChannelDataDataMsTeamsChannelDataConnection struct {
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

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnection) implementsTenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams token connection
//
// Satisfied by
// [TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection],
// [TenantSetParamsChannelDataDataMsTeamsChannelDataConnection].
type TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsTenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// Microsoft Teams token connection
type TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) implementsTenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams incoming webhook connection
type TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) implementsTenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// The incoming webhook
type TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Discord channel data
type TenantSetParamsChannelDataDataDiscordChannelData struct {
	Connections param.Field[[]TenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion] `json:"connections,required"`
}

func (r TenantSetParamsChannelDataDataDiscordChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataDiscordChannelData) implementsTenantSetParamsChannelDataDataUnion() {
}

// Discord channel connection
type TenantSetParamsChannelDataDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnection) implementsTenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection],
// [TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection],
// [TenantSetParamsChannelDataDataDiscordChannelDataConnection].
type TenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsTenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) implementsTenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord incoming webhook connection
type TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) implementsTenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// The incoming webhook
type TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type TenantSetParamsPreferences struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]TenantSetParamsPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsPreferencesChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]TenantSetParamsPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r TenantSetParamsPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject].
type TenantSetParamsPreferencesCategoriesUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsTenantSetParamsPreferencesCategoriesUnion() {
}

// Channel type preferences
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                 `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                     `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                     `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                               `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                               `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                           `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type TenantSetParamsPreferencesChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantSetParamsPreferencesChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantSetParamsPreferencesChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantSetParamsPreferencesChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantSetParamsPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantSetParamsPreferencesChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantSetParamsPreferencesChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantSetParamsPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesChannelTypesChatUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                            `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                            `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                             `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                             `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                            `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                            `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                 `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesChannelTypesPushUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                            `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                            `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                           `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject].
type TenantSetParamsPreferencesWorkflowsUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsTenantSetParamsPreferencesWorkflowsUnion() {
}

// Channel type preferences
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                               `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                               `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                               `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                               `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                    `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                    `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                               `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                               `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                              `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                              `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                          `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

type TenantSetParamsSettings struct {
	Branding param.Field[TenantSetParamsSettingsBranding] `json:"branding"`
	// Set preferences for a recipient
	PreferenceSet param.Field[TenantSetParamsSettingsPreferenceSet] `json:"preference_set"`
}

func (r TenantSetParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsBranding struct {
	IconURL              param.Field[string] `json:"icon_url"`
	LogoURL              param.Field[string] `json:"logo_url"`
	PrimaryColor         param.Field[string] `json:"primary_color"`
	PrimaryColorContrast param.Field[string] `json:"primary_color_contrast"`
}

func (r TenantSetParamsSettingsBranding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type TenantSetParamsSettingsPreferenceSet struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]TenantSetParamsSettingsPreferenceSetCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsSettingsPreferenceSetChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]TenantSetParamsSettingsPreferenceSetWorkflowsUnion] `json:"workflows"`
}

func (r TenantSetParamsSettingsPreferenceSet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject].
type TenantSetParamsSettingsPreferenceSetCategoriesUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesUnion() {
}

// Channel type preferences
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                          `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                          `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                               `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                               `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                          `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                         `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                         `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                     `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                     `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type TenantSetParamsSettingsPreferenceSetChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetChannelTypesChatUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                      `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                      `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                       `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                       `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                      `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                      `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                           `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetChannelTypesPushUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                      `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                      `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                     `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                     `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject].
type TenantSetParamsSettingsPreferenceSetWorkflowsUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsUnion() {
}

// Channel type preferences
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                         `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                         `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                          `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                         `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                         `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                              `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                              `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                         `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                         `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                        `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                        `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                    `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                    `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

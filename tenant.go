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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
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
	Token       param.Field[TenantSetParamsChannelDataDataSlackChannelDataToken]             `json:"token"`
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
// [TenantSetParamsChannelDataDataSlackChannelDataConnectionsTokenConnection],
// [TenantSetParamsChannelDataDataSlackChannelDataConnectionsIncomingWebhookConnection],
// [TenantSetParamsChannelDataDataSlackChannelDataConnection].
type TenantSetParamsChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsTenantSetParamsChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type TenantSetParamsChannelDataDataSlackChannelDataConnectionsTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnectionsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnectionsTokenConnection) implementsTenantSetParamsChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type TenantSetParamsChannelDataDataSlackChannelDataConnectionsIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnectionsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataSlackChannelDataConnectionsIncomingWebhookConnection) implementsTenantSetParamsChannelDataDataSlackChannelDataConnectionUnion() {
}

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

// A Slack connection, which either includes a channel_id or a user_id
type TenantSetParamsChannelDataDataMsTeamsChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnection) implementsTenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsTokenConnection],
// [TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsIncomingWebhookConnection],
// [TenantSetParamsChannelDataDataMsTeamsChannelDataConnection].
type TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsTenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsTokenConnection) implementsTenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataMsTeamsChannelDataConnectionsIncomingWebhookConnection) implementsTenantSetParamsChannelDataDataMsTeamsChannelDataConnectionUnion() {
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
	ChannelID param.Field[string] `json:"channel_id"`
	URL       param.Field[string] `json:"url"`
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnection) implementsTenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [TenantSetParamsChannelDataDataDiscordChannelDataConnectionsChannelConnection],
// [TenantSetParamsChannelDataDataDiscordChannelDataConnectionsIncomingWebhookConnection],
// [TenantSetParamsChannelDataDataDiscordChannelDataConnection].
type TenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsTenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type TenantSetParamsChannelDataDataDiscordChannelDataConnectionsChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnectionsChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnectionsChannelConnection) implementsTenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type TenantSetParamsChannelDataDataDiscordChannelDataConnectionsIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnectionsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsChannelDataDataDiscordChannelDataConnectionsIncomingWebhookConnection) implementsTenantSetParamsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Set preferences for a recipient
type TenantSetParamsPreferences struct {
	Categories param.Field[map[string]TenantSetParamsPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsPreferencesChannelTypes]              `json:"channel_types"`
	Workflows    param.Field[map[string]TenantSetParamsPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r TenantSetParamsPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool], [TenantSetParamsPreferencesCategoriesObject].
type TenantSetParamsPreferencesCategoriesUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesUnion()
}

type TenantSetParamsPreferencesCategoriesObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantSetParamsPreferencesCategoriesObjectCondition]  `json:"conditions"`
}

func (r TenantSetParamsPreferencesCategoriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesObject) ImplementsTenantSetParamsPreferencesCategoriesUnion() {
}

// Channel type preferences
type TenantSetParamsPreferencesCategoriesObjectChannelTypes struct {
	Chat      param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPUnion]      `json:"http"`
	InAppFeed param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSUnion]       `json:"sms"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditions].
type TenantSetParamsPreferencesCategoriesObjectChannelTypesChatUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesChatUnion()
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditions) ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                 `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains             TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditions].
type TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesEmailUnion()
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditions) ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                                  `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                  `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains             TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditions].
type TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPUnion()
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditions) ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                 `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains             TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditions].
type TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion()
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditions) ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                                      `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                      `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditions].
type TenantSetParamsPreferencesCategoriesObjectChannelTypesPushUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesPushUnion()
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditions) ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                 `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains             TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditions].
type TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesSMSUnion()
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditions) ImplementsTenantSetParamsPreferencesCategoriesObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains             TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantSetParamsPreferencesCategoriesObjectCondition struct {
	Argument param.Field[string]                                                       `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesCategoriesObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                       `json:"variable,required"`
}

func (r TenantSetParamsPreferencesCategoriesObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesCategoriesObjectConditionsOperator string

const (
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorLessThan             TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorContains             TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorNotContains          TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorEmpty                TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesCategoriesObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesCategoriesObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesCategoriesObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorLessThan, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorContains, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorNotContains, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorEmpty, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesCategoriesObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type TenantSetParamsPreferencesChannelTypes struct {
	Chat      param.Field[TenantSetParamsPreferencesChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[TenantSetParamsPreferencesChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[TenantSetParamsPreferencesChannelTypesHTTPUnion]      `json:"http"`
	InAppFeed param.Field[TenantSetParamsPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[TenantSetParamsPreferencesChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[TenantSetParamsPreferencesChannelTypesSMSUnion]       `json:"sms"`
}

func (r TenantSetParamsPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesChatConditions].
type TenantSetParamsPreferencesChannelTypesChatUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesChatUnion()
}

type TenantSetParamsPreferencesChannelTypesChatConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesChatConditions) ImplementsTenantSetParamsPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                 `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesEmailConditions].
type TenantSetParamsPreferencesChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesEmailUnion()
}

type TenantSetParamsPreferencesChannelTypesEmailConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesEmailConditions) ImplementsTenantSetParamsPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                  `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                  `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesHTTPConditions].
type TenantSetParamsPreferencesChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesHTTPUnion()
}

type TenantSetParamsPreferencesChannelTypesHTTPConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesHTTPConditions) ImplementsTenantSetParamsPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                 `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesInAppFeedConditions].
type TenantSetParamsPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesInAppFeedUnion()
}

type TenantSetParamsPreferencesChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesInAppFeedConditions) ImplementsTenantSetParamsPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                      `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                      `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesPushConditions].
type TenantSetParamsPreferencesChannelTypesPushUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesPushUnion()
}

type TenantSetParamsPreferencesChannelTypesPushConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesPushConditions) ImplementsTenantSetParamsPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                 `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesChannelTypesSMSConditions].
type TenantSetParamsPreferencesChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsPreferencesChannelTypesSMSUnion()
}

type TenantSetParamsPreferencesChannelTypesSMSConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesChannelTypesSMSConditions) ImplementsTenantSetParamsPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                `json:"variable,required"`
}

func (r TenantSetParamsPreferencesChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContains             TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContains, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool], [TenantSetParamsPreferencesWorkflowsObject].
type TenantSetParamsPreferencesWorkflowsUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsUnion()
}

type TenantSetParamsPreferencesWorkflowsObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantSetParamsPreferencesWorkflowsObjectCondition]  `json:"conditions"`
}

func (r TenantSetParamsPreferencesWorkflowsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsObject) ImplementsTenantSetParamsPreferencesWorkflowsUnion() {
}

// Channel type preferences
type TenantSetParamsPreferencesWorkflowsObjectChannelTypes struct {
	Chat      param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion]      `json:"http"`
	InAppFeed param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSUnion]       `json:"sms"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditions].
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesChatUnion()
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditions) ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditions].
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailUnion()
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditions) ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                 `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditions].
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion()
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditions) ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditions].
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion()
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditions) ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                                     `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                     `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditions].
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesPushUnion()
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditions) ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditions].
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSUnion()
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditions struct {
	Conditions param.Field[[]TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditions) ImplementsTenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                               `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                               `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantSetParamsPreferencesWorkflowsObjectCondition struct {
	Argument param.Field[string]                                                      `json:"argument,required"`
	Operator param.Field[TenantSetParamsPreferencesWorkflowsObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                      `json:"variable,required"`
}

func (r TenantSetParamsPreferencesWorkflowsObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsPreferencesWorkflowsObjectConditionsOperator string

const (
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorEqualTo              TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "equal_to"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorNotEqualTo           TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "not_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThan          TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "greater_than"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorLessThan             TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "less_than"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorContains             TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "contains"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorNotContains          TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "not_contains"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorEmpty                TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "empty"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorNotEmpty             TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "not_empty"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorContainsAll          TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "contains_all"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestamp          TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "is_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsNotTimestamp       TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampAfter     TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampBefore    TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampBetween   TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsAudienceMember     TenantSetParamsPreferencesWorkflowsObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsPreferencesWorkflowsObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorEqualTo, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorNotEqualTo, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThan, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorLessThan, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorContains, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorNotContains, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorEmpty, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorNotEmpty, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorContainsAll, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestamp, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsNotTimestamp, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampAfter, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampBefore, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampBetween, TenantSetParamsPreferencesWorkflowsObjectConditionsOperatorIsAudienceMember:
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
	Categories param.Field[map[string]TenantSetParamsSettingsPreferenceSetCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsSettingsPreferenceSetChannelTypes]              `json:"channel_types"`
	Workflows    param.Field[map[string]TenantSetParamsSettingsPreferenceSetWorkflowsUnion] `json:"workflows"`
}

func (r TenantSetParamsSettingsPreferenceSet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesObject].
type TenantSetParamsSettingsPreferenceSetCategoriesUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesUnion()
}

type TenantSetParamsSettingsPreferenceSetCategoriesObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesObjectCondition]  `json:"conditions"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObject) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesUnion() {
}

// Channel type preferences
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypes struct {
	Chat      param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPUnion]      `json:"http"`
	InAppFeed param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSUnion]       `json:"sms"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditions].
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatUnion()
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditions) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                           `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditions].
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailUnion()
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditions) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                                            `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                            `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditions].
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPUnion()
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditions) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                           `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditions].
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedUnion()
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditions) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditions].
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushUnion()
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditions) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                           `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditions].
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSUnion()
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditions) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                          `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetCategoriesObjectCondition struct {
	Argument param.Field[string]                                                                 `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                 `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetCategoriesObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type TenantSetParamsSettingsPreferenceSetChannelTypes struct {
	Chat      param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion]      `json:"http"`
	InAppFeed param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion]       `json:"sms"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesChatConditions].
type TenantSetParamsSettingsPreferenceSetChannelTypesChatUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesChatUnion()
}

type TenantSetParamsSettingsPreferenceSetChannelTypesChatConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesChatConditions) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                           `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditions].
type TenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion()
}

type TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditions) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                            `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                            `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditions].
type TenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion()
}

type TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditions) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                           `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditions].
type TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion()
}

type TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditions) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesPushConditions].
type TenantSetParamsSettingsPreferenceSetChannelTypesPushUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesPushUnion()
}

type TenantSetParamsSettingsPreferenceSetChannelTypesPushConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesPushConditions) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                           `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditions].
type TenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion()
}

type TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditions) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                          `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsObject].
type TenantSetParamsSettingsPreferenceSetWorkflowsUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsUnion()
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObject struct {
	// Channel type preferences
	ChannelTypes param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsObjectCondition]  `json:"conditions"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObject) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsUnion() {
}

// Channel type preferences
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypes struct {
	Chat      param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPUnion]      `json:"http"`
	InAppFeed param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSUnion]       `json:"sms"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditions].
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatUnion()
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditions) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                          `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditions].
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailUnion()
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditions) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                                           `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                           `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditions].
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPUnion()
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditions) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                          `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditions].
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedUnion()
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditions) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                                               `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                               `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditions].
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushUnion()
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditions) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                                          `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                          `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditions].
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSUnion interface {
	ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSUnion()
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditions struct {
	Conditions param.Field[[]TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditions) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                                         `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                         `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type TenantSetParamsSettingsPreferenceSetWorkflowsObjectCondition struct {
	Argument param.Field[string]                                                                `json:"argument,required"`
	Operator param.Field[TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                `json:"variable,required"`
}

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator string

const (
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorEqualTo              TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorNotEqualTo           TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "not_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorGreaterThan          TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "greater_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorLessThan             TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "less_than"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "greater_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorLessThanOrEqualTo    TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "less_than_or_equal_to"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorContains             TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorNotContains          TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "not_contains"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorEmpty                TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorNotEmpty             TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "not_empty"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorContainsAll          TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "contains_all"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsTimestamp          TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "is_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsNotTimestamp       TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "is_not_timestamp"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsTimestampAfter     TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "is_timestamp_after"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsTimestampBefore    TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "is_timestamp_before"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsTimestampBetween   TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "is_timestamp_between"
	TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsAudienceMember     TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator = "is_audience_member"
)

func (r TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperator) IsKnown() bool {
	switch r {
	case TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorNotEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorGreaterThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorLessThan, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorLessThanOrEqualTo, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorNotContains, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorNotEmpty, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorContainsAll, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsNotTimestamp, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsTimestampAfter, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsTimestampBefore, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsTimestampBetween, TenantSetParamsSettingsPreferenceSetWorkflowsObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

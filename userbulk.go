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

// UserBulkService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserBulkService] method instead.
type UserBulkService struct {
	Options []option.RequestOption
}

// NewUserBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserBulkService(opts ...option.RequestOption) (r *UserBulkService) {
	r = &UserBulkService{}
	r.Options = opts
	return
}

// Bulk deletes users
func (r *UserBulkService) Delete(ctx context.Context, params UserBulkDeleteParams, opts ...option.RequestOption) (res *UserBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/bulk/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Bulk identifies users
func (r *UserBulkService) Identify(ctx context.Context, body UserBulkIdentifyParams, opts ...option.RequestOption) (res *UserBulkIdentifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/bulk/identify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk set user preferences
func (r *UserBulkService) SetPreferences(ctx context.Context, body UserBulkSetPreferencesParams, opts ...option.RequestOption) (res *UserBulkSetPreferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/bulk/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A bulk operation entity
type UserBulkDeleteResponse struct {
	ID                 string                       `json:"id,required" format:"uuid"`
	Typename           string                       `json:"__typename,required"`
	EstimatedTotalRows int64                        `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                    `json:"inserted_at,required" format:"date-time"`
	Name               string                       `json:"name,required"`
	ProcessedRows      int64                        `json:"processed_rows,required"`
	Status             UserBulkDeleteResponseStatus `json:"status,required"`
	SuccessCount       int64                        `json:"success_count,required"`
	UpdatedAt          time.Time                    `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                    `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                        `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []UserBulkDeleteResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                         `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                         `json:"started_at,nullable" format:"date-time"`
	JSON       userBulkDeleteResponseJSON        `json:"-"`
}

// userBulkDeleteResponseJSON contains the JSON metadata for the struct
// [UserBulkDeleteResponse]
type userBulkDeleteResponseJSON struct {
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

func (r *UserBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type UserBulkDeleteResponseStatus string

const (
	UserBulkDeleteResponseStatusQueued     UserBulkDeleteResponseStatus = "queued"
	UserBulkDeleteResponseStatusProcessing UserBulkDeleteResponseStatus = "processing"
	UserBulkDeleteResponseStatusCompleted  UserBulkDeleteResponseStatus = "completed"
	UserBulkDeleteResponseStatusFailed     UserBulkDeleteResponseStatus = "failed"
)

func (r UserBulkDeleteResponseStatus) IsKnown() bool {
	switch r {
	case UserBulkDeleteResponseStatusQueued, UserBulkDeleteResponseStatusProcessing, UserBulkDeleteResponseStatusCompleted, UserBulkDeleteResponseStatusFailed:
		return true
	}
	return false
}

type UserBulkDeleteResponseErrorItem struct {
	ID         string                              `json:"id,required"`
	Collection string                              `json:"collection,nullable"`
	JSON       userBulkDeleteResponseErrorItemJSON `json:"-"`
}

// userBulkDeleteResponseErrorItemJSON contains the JSON metadata for the struct
// [UserBulkDeleteResponseErrorItem]
type userBulkDeleteResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBulkDeleteResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkDeleteResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type UserBulkIdentifyResponse struct {
	ID                 string                         `json:"id,required" format:"uuid"`
	Typename           string                         `json:"__typename,required"`
	EstimatedTotalRows int64                          `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                      `json:"inserted_at,required" format:"date-time"`
	Name               string                         `json:"name,required"`
	ProcessedRows      int64                          `json:"processed_rows,required"`
	Status             UserBulkIdentifyResponseStatus `json:"status,required"`
	SuccessCount       int64                          `json:"success_count,required"`
	UpdatedAt          time.Time                      `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                      `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                          `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []UserBulkIdentifyResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                           `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                           `json:"started_at,nullable" format:"date-time"`
	JSON       userBulkIdentifyResponseJSON        `json:"-"`
}

// userBulkIdentifyResponseJSON contains the JSON metadata for the struct
// [UserBulkIdentifyResponse]
type userBulkIdentifyResponseJSON struct {
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

func (r *UserBulkIdentifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkIdentifyResponseJSON) RawJSON() string {
	return r.raw
}

type UserBulkIdentifyResponseStatus string

const (
	UserBulkIdentifyResponseStatusQueued     UserBulkIdentifyResponseStatus = "queued"
	UserBulkIdentifyResponseStatusProcessing UserBulkIdentifyResponseStatus = "processing"
	UserBulkIdentifyResponseStatusCompleted  UserBulkIdentifyResponseStatus = "completed"
	UserBulkIdentifyResponseStatusFailed     UserBulkIdentifyResponseStatus = "failed"
)

func (r UserBulkIdentifyResponseStatus) IsKnown() bool {
	switch r {
	case UserBulkIdentifyResponseStatusQueued, UserBulkIdentifyResponseStatusProcessing, UserBulkIdentifyResponseStatusCompleted, UserBulkIdentifyResponseStatusFailed:
		return true
	}
	return false
}

type UserBulkIdentifyResponseErrorItem struct {
	ID         string                                `json:"id,required"`
	Collection string                                `json:"collection,nullable"`
	JSON       userBulkIdentifyResponseErrorItemJSON `json:"-"`
}

// userBulkIdentifyResponseErrorItemJSON contains the JSON metadata for the struct
// [UserBulkIdentifyResponseErrorItem]
type userBulkIdentifyResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBulkIdentifyResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkIdentifyResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type UserBulkSetPreferencesResponse struct {
	ID                 string                               `json:"id,required" format:"uuid"`
	Typename           string                               `json:"__typename,required"`
	EstimatedTotalRows int64                                `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                            `json:"inserted_at,required" format:"date-time"`
	Name               string                               `json:"name,required"`
	ProcessedRows      int64                                `json:"processed_rows,required"`
	Status             UserBulkSetPreferencesResponseStatus `json:"status,required"`
	SuccessCount       int64                                `json:"success_count,required"`
	UpdatedAt          time.Time                            `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                            `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                                `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []UserBulkSetPreferencesResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                                 `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                                 `json:"started_at,nullable" format:"date-time"`
	JSON       userBulkSetPreferencesResponseJSON        `json:"-"`
}

// userBulkSetPreferencesResponseJSON contains the JSON metadata for the struct
// [UserBulkSetPreferencesResponse]
type userBulkSetPreferencesResponseJSON struct {
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

func (r *UserBulkSetPreferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkSetPreferencesResponseJSON) RawJSON() string {
	return r.raw
}

type UserBulkSetPreferencesResponseStatus string

const (
	UserBulkSetPreferencesResponseStatusQueued     UserBulkSetPreferencesResponseStatus = "queued"
	UserBulkSetPreferencesResponseStatusProcessing UserBulkSetPreferencesResponseStatus = "processing"
	UserBulkSetPreferencesResponseStatusCompleted  UserBulkSetPreferencesResponseStatus = "completed"
	UserBulkSetPreferencesResponseStatusFailed     UserBulkSetPreferencesResponseStatus = "failed"
)

func (r UserBulkSetPreferencesResponseStatus) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesResponseStatusQueued, UserBulkSetPreferencesResponseStatusProcessing, UserBulkSetPreferencesResponseStatusCompleted, UserBulkSetPreferencesResponseStatusFailed:
		return true
	}
	return false
}

type UserBulkSetPreferencesResponseErrorItem struct {
	ID         string                                      `json:"id,required"`
	Collection string                                      `json:"collection,nullable"`
	JSON       userBulkSetPreferencesResponseErrorItemJSON `json:"-"`
}

// userBulkSetPreferencesResponseErrorItemJSON contains the JSON metadata for the
// struct [UserBulkSetPreferencesResponseErrorItem]
type userBulkSetPreferencesResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBulkSetPreferencesResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkSetPreferencesResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

type UserBulkDeleteParams struct {
	// The IDs of the users to delete
	UserIDs1 param.Field[[]string] `query:"user_ids,required"`
	UserIDs2 param.Field[[]string] `json:"user_ids,required"`
}

func (r UserBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [UserBulkDeleteParams]'s query parameters as `url.Values`.
func (r UserBulkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserBulkIdentifyParams struct {
	Users param.Field[[]UserBulkIdentifyParamsUser] `json:"users,required"`
}

func (r UserBulkIdentifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of parameters to inline-identify a user with. Inline identifying the user
// will ensure that the user is available before the request is executed in Knock.
// It will perform an upsert against the user you're supplying, replacing any
// properties specified.
type UserBulkIdentifyParamsUser struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]UserBulkIdentifyParamsUsersChannelData] `json:"channel_data"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]UserBulkIdentifyParamsUsersPreferences] `json:"preferences"`
	ExtraFields map[string]interface{}                                         `json:"-,extras"`
}

func (r UserBulkIdentifyParamsUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set channel data for a type of channel
type UserBulkIdentifyParamsUsersChannelData struct {
	// Channel data for push providers
	Data param.Field[UserBulkIdentifyParamsUsersChannelDataDataUnion] `json:"data,required"`
}

func (r UserBulkIdentifyParamsUsersChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type UserBulkIdentifyParamsUsersChannelDataData struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r UserBulkIdentifyParamsUsersChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataData) implementsUserBulkIdentifyParamsUsersChannelDataDataUnion() {
}

// Channel data for push providers
//
// Satisfied by [UserBulkIdentifyParamsUsersChannelDataDataPushChannelData],
// [UserBulkIdentifyParamsUsersChannelDataDataOneSignalChannelData],
// [UserBulkIdentifyParamsUsersChannelDataDataSlackChannelData],
// [UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelData],
// [UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelData],
// [UserBulkIdentifyParamsUsersChannelDataData].
type UserBulkIdentifyParamsUsersChannelDataDataUnion interface {
	implementsUserBulkIdentifyParamsUsersChannelDataDataUnion()
}

// Channel data for push providers
type UserBulkIdentifyParamsUsersChannelDataDataPushChannelData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataPushChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataPushChannelData) implementsUserBulkIdentifyParamsUsersChannelDataDataUnion() {
}

// OneSignal channel data
type UserBulkIdentifyParamsUsersChannelDataDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataOneSignalChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataOneSignalChannelData) implementsUserBulkIdentifyParamsUsersChannelDataDataUnion() {
}

// Slack channel data
type UserBulkIdentifyParamsUsersChannelDataDataSlackChannelData struct {
	Connections param.Field[[]UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionUnion] `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token param.Field[UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataToken] `json:"token"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataSlackChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataSlackChannelData) implementsUserBulkIdentifyParamsUsersChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnection) implementsUserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionsSlackTokenConnection],
// [UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection],
// [UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnection].
type UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsUserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionsSlackTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) implementsUserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) implementsUserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataConnectionUnion() {
}

// A token that's used to store the access token for a Slack workspace.
type UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataToken struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataSlackChannelDataToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelData struct {
	Connections param.Field[[]UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionUnion] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelData) implementsUserBulkIdentifyParamsUsersChannelDataDataUnion() {
}

// Microsoft Teams token connection
type UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnection struct {
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

func (r UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnection) implementsUserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams token connection
//
// Satisfied by
// [UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection],
// [UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnection].
type UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsUserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// Microsoft Teams token connection
type UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) implementsUserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams incoming webhook connection
type UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) implementsUserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// The incoming webhook
type UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Discord channel data
type UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelData struct {
	Connections param.Field[[]UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionUnion] `json:"connections,required"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelData) implementsUserBulkIdentifyParamsUsersChannelDataDataUnion() {
}

// Discord channel connection
type UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnection) implementsUserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection],
// [UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection],
// [UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnection].
type UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsUserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) implementsUserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord incoming webhook connection
type UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) implementsUserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionUnion() {
}

// The incoming webhook
type UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r UserBulkIdentifyParamsUsersChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type UserBulkIdentifyParamsUsersPreferences struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]UserBulkIdentifyParamsUsersPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]UserBulkIdentifyParamsUsersPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r UserBulkIdentifyParamsUsersPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject].
type UserBulkIdentifyParamsUsersPreferencesCategoriesUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesUnion() {
}

// Channel type preferences
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                             `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                 `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                       `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                       `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type UserBulkIdentifyParamsUsersPreferencesChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSUnion] `json:"sms"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesChannelTypesChatUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                        `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                        `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                         `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                         `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                        `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                        `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                             `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                             `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesChannelTypesPushUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                        `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                        `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                       `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                       `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject].
type UserBulkIdentifyParamsUsersPreferencesWorkflowsUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsUnion() {
}

// Channel type preferences
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                          `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                          `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                      `json:"argument,required"`
	Operator param.Field[UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                      `json:"variable,required"`
}

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, UserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

type UserBulkSetPreferencesParams struct {
	// Set preferences for a recipient
	Preferences param.Field[UserBulkSetPreferencesParamsPreferences] `json:"preferences,required"`
	UserIDs     param.Field[[]string]                                `json:"user_ids,required"`
}

func (r UserBulkSetPreferencesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type UserBulkSetPreferencesParamsPreferences struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]UserBulkSetPreferencesParamsPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]UserBulkSetPreferencesParamsPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r UserBulkSetPreferencesParamsPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject].
type UserBulkSetPreferencesParamsPreferencesCategoriesUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesUnion() {
}

// Channel type preferences
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                             `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                              `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                              `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                             `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                  `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                             `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                        `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                        `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type UserBulkSetPreferencesParamsPreferencesChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesSMSUnion] `json:"sms"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesChannelTypesChatUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                         `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                         `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesChannelTypesEmailUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                          `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                          `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                         `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                         `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                              `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                              `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesChannelTypesPushUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                         `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                         `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesChannelTypesSMSUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                        `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                        `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject].
type UserBulkSetPreferencesParamsPreferencesWorkflowsUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsUnion() {
}

// Channel type preferences
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                             `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                 `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                       `json:"argument,required"`
	Operator param.Field[UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                       `json:"variable,required"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

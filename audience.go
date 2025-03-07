// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
)

// AudienceService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudienceService] method instead.
type AudienceService struct {
	Options []option.RequestOption
}

// NewAudienceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudienceService(opts ...option.RequestOption) (r *AudienceService) {
	r = &AudienceService{}
	r.Options = opts
	return
}

// Add members
func (r *AudienceService) AddMembers(ctx context.Context, key string, body AudienceAddMembersParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/members", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List members
func (r *AudienceService) ListMembers(ctx context.Context, key string, opts ...option.RequestOption) (res *AudienceListMembersResponse, err error) {
	opts = append(r.Options[:], opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/members", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove members
func (r *AudienceService) RemoveMembers(ctx context.Context, key string, body AudienceRemoveMembersParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/members", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// A response containing a list of audience members
type AudienceListMembersResponse struct {
	Entries []AudienceListMembersResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo AudienceListMembersResponsePageInfo `json:"page_info,required"`
	JSON     audienceListMembersResponseJSON     `json:"-"`
}

// audienceListMembersResponseJSON contains the JSON metadata for the struct
// [AudienceListMembersResponse]
type audienceListMembersResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudienceListMembersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audienceListMembersResponseJSON) RawJSON() string {
	return r.raw
}

// A user belonging to an audience
type AudienceListMembersResponseEntry struct {
	Typename string    `json:"__typename,required"`
	AddedAt  time.Time `json:"added_at,required" format:"date-time"`
	// A user object
	User   User                                 `json:"user,required"`
	UserID string                               `json:"user_id,required"`
	Tenant string                               `json:"tenant,nullable"`
	JSON   audienceListMembersResponseEntryJSON `json:"-"`
}

// audienceListMembersResponseEntryJSON contains the JSON metadata for the struct
// [AudienceListMembersResponseEntry]
type audienceListMembersResponseEntryJSON struct {
	Typename    apijson.Field
	AddedAt     apijson.Field
	User        apijson.Field
	UserID      apijson.Field
	Tenant      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudienceListMembersResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audienceListMembersResponseEntryJSON) RawJSON() string {
	return r.raw
}

// The information about a paginated result
type AudienceListMembersResponsePageInfo struct {
	Typename string                                  `json:"__typename,required"`
	PageSize int64                                   `json:"page_size,required"`
	After    string                                  `json:"after,nullable"`
	Before   string                                  `json:"before,nullable"`
	JSON     audienceListMembersResponsePageInfoJSON `json:"-"`
}

// audienceListMembersResponsePageInfoJSON contains the JSON metadata for the
// struct [AudienceListMembersResponsePageInfo]
type audienceListMembersResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudienceListMembersResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audienceListMembersResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

type AudienceAddMembersParams struct {
	Members param.Field[[]AudienceAddMembersParamsMember] `json:"members,required"`
}

func (r AudienceAddMembersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A request for an individual audience member
type AudienceAddMembersParamsMember struct {
	// A set of parameters to inline-identify a user with. Inline identifying the user
	// will ensure that the user is available before the request is executed in Knock.
	// It will perform an upsert against the user you're supplying, replacing any
	// properties specified.
	User   param.Field[AudienceAddMembersParamsMembersUser] `json:"user,required"`
	Tenant param.Field[string]                              `json:"tenant"`
}

func (r AudienceAddMembersParamsMember) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of parameters to inline-identify a user with. Inline identifying the user
// will ensure that the user is available before the request is executed in Knock.
// It will perform an upsert against the user you're supplying, replacing any
// properties specified.
type AudienceAddMembersParamsMembersUser struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]AudienceAddMembersParamsMembersUserChannelData] `json:"channel_data"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]AudienceAddMembersParamsMembersUserPreferences] `json:"preferences"`
	ExtraFields map[string]interface{}                                                 `json:"-,extras"`
}

func (r AudienceAddMembersParamsMembersUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set channel data for a type of channel
type AudienceAddMembersParamsMembersUserChannelData struct {
	// Channel data for push providers
	Data param.Field[AudienceAddMembersParamsMembersUserChannelDataDataUnion] `json:"data,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type AudienceAddMembersParamsMembersUserChannelDataData struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataData) implementsAudienceAddMembersParamsMembersUserChannelDataDataUnion() {
}

// Channel data for push providers
//
// Satisfied by
// [AudienceAddMembersParamsMembersUserChannelDataDataPushChannelData],
// [AudienceAddMembersParamsMembersUserChannelDataDataOneSignalChannelData],
// [AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelData],
// [AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelData],
// [AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelData],
// [AudienceAddMembersParamsMembersUserChannelDataData].
type AudienceAddMembersParamsMembersUserChannelDataDataUnion interface {
	implementsAudienceAddMembersParamsMembersUserChannelDataDataUnion()
}

// Channel data for push providers
type AudienceAddMembersParamsMembersUserChannelDataDataPushChannelData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataPushChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataPushChannelData) implementsAudienceAddMembersParamsMembersUserChannelDataDataUnion() {
}

// OneSignal channel data
type AudienceAddMembersParamsMembersUserChannelDataDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataOneSignalChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataOneSignalChannelData) implementsAudienceAddMembersParamsMembersUserChannelDataDataUnion() {
}

// Slack channel data
type AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelData struct {
	Connections param.Field[[]AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion] `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token param.Field[AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataToken] `json:"token"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelData) implementsAudienceAddMembersParamsMembersUserChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnection) implementsAudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackTokenConnection],
// [AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection],
// [AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnection].
type AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsAudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) implementsAudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) implementsAudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion() {
}

// A token that's used to store the access token for a Slack workspace.
type AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataToken struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataSlackChannelDataToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelData struct {
	Connections param.Field[[]AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelData) implementsAudienceAddMembersParamsMembersUserChannelDataDataUnion() {
}

// Microsoft Teams token connection
type AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnection struct {
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

func (r AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnection) implementsAudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams token connection
//
// Satisfied by
// [AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection],
// [AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnection].
type AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsAudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// Microsoft Teams token connection
type AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) implementsAudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams incoming webhook connection
type AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) implementsAudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// The incoming webhook
type AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Discord channel data
type AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelData struct {
	Connections param.Field[[]AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion] `json:"connections,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelData) implementsAudienceAddMembersParamsMembersUserChannelDataDataUnion() {
}

// Discord channel connection
type AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnection) implementsAudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection],
// [AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection],
// [AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnection].
type AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsAudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) implementsAudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord incoming webhook connection
type AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) implementsAudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion() {
}

// The incoming webhook
type AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r AudienceAddMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type AudienceAddMembersParamsMembersUserPreferences struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]AudienceAddMembersParamsMembersUserPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]AudienceAddMembersParamsMembersUserPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r AudienceAddMembersParamsMembersUserPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject].
type AudienceAddMembersParamsMembersUserPreferencesCategoriesUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesUnion() {
}

// Channel type preferences
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                    `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                    `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                     `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                     `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                    `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                    `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                         `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                         `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                    `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                    `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                   `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                   `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                               `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                               `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type AudienceAddMembersParamsMembersUserPreferencesChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSUnion] `json:"sms"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                 `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                 `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                     `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                     `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                               `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                               `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject].
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsUnion() {
}

// Channel type preferences
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                   `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                   `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                    `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                    `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                   `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                   `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                        `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                        `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                   `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                   `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                  `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                              `json:"argument,required"`
	Operator param.Field[AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                              `json:"variable,required"`
}

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

type AudienceRemoveMembersParams struct {
	Members param.Field[[]AudienceRemoveMembersParamsMember] `json:"members,required"`
}

func (r AudienceRemoveMembersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A request for an individual audience member
type AudienceRemoveMembersParamsMember struct {
	// A set of parameters to inline-identify a user with. Inline identifying the user
	// will ensure that the user is available before the request is executed in Knock.
	// It will perform an upsert against the user you're supplying, replacing any
	// properties specified.
	User   param.Field[AudienceRemoveMembersParamsMembersUser] `json:"user,required"`
	Tenant param.Field[string]                                 `json:"tenant"`
}

func (r AudienceRemoveMembersParamsMember) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of parameters to inline-identify a user with. Inline identifying the user
// will ensure that the user is available before the request is executed in Knock.
// It will perform an upsert against the user you're supplying, replacing any
// properties specified.
type AudienceRemoveMembersParamsMembersUser struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]AudienceRemoveMembersParamsMembersUserChannelData] `json:"channel_data"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]AudienceRemoveMembersParamsMembersUserPreferences] `json:"preferences"`
	ExtraFields map[string]interface{}                                                    `json:"-,extras"`
}

func (r AudienceRemoveMembersParamsMembersUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set channel data for a type of channel
type AudienceRemoveMembersParamsMembersUserChannelData struct {
	// Channel data for push providers
	Data param.Field[AudienceRemoveMembersParamsMembersUserChannelDataDataUnion] `json:"data,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type AudienceRemoveMembersParamsMembersUserChannelDataData struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataData) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataUnion() {
}

// Channel data for push providers
//
// Satisfied by
// [AudienceRemoveMembersParamsMembersUserChannelDataDataPushChannelData],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataOneSignalChannelData],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelData],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelData],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelData],
// [AudienceRemoveMembersParamsMembersUserChannelDataData].
type AudienceRemoveMembersParamsMembersUserChannelDataDataUnion interface {
	implementsAudienceRemoveMembersParamsMembersUserChannelDataDataUnion()
}

// Channel data for push providers
type AudienceRemoveMembersParamsMembersUserChannelDataDataPushChannelData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataPushChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataPushChannelData) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataUnion() {
}

// OneSignal channel data
type AudienceRemoveMembersParamsMembersUserChannelDataDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataOneSignalChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataOneSignalChannelData) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataUnion() {
}

// Slack channel data
type AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelData struct {
	Connections param.Field[[]AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion] `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token param.Field[AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataToken] `json:"token"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelData) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnection) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackTokenConnection],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnection].
type AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsAudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataConnectionUnion() {
}

// A token that's used to store the access token for a Slack workspace.
type AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataToken struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataSlackChannelDataToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelData struct {
	Connections param.Field[[]AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelData) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataUnion() {
}

// Microsoft Teams token connection
type AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnection struct {
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

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnection) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams token connection
//
// Satisfied by
// [AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnection].
type AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsAudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// Microsoft Teams token connection
type AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams incoming webhook connection
type AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// The incoming webhook
type AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Discord channel data
type AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelData struct {
	Connections param.Field[[]AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion] `json:"connections,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelData) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataUnion() {
}

// Discord channel connection
type AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnection) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection],
// [AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnection].
type AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsAudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord incoming webhook connection
type AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) implementsAudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionUnion() {
}

// The incoming webhook
type AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r AudienceRemoveMembersParamsMembersUserChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type AudienceRemoveMembersParamsMembersUserPreferences struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]AudienceRemoveMembersParamsMembersUserPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesUnion() {
}

// Channel type preferences
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                       `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                        `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                        `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                       `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                            `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                       `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                      `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                      `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                  `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                  `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSUnion] `json:"sms"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                   `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                   `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                    `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                    `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                   `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                   `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                        `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                        `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                   `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                   `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                  `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                  `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsUnion() {
}

// Channel type preferences
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                      `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                      `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                       `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                      `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                      `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                           `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                      `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                      `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                     `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                     `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                 `json:"argument,required"`
	Operator param.Field[AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                 `json:"variable,required"`
}

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
)

// UserService contains methods and other services that help with interacting with
// the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

// Identify a user
func (r *UserService) Update(ctx context.Context, userID string, body UserUpdateParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of users
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a user
func (r *UserService) Delete(ctx context.Context, userID string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns a user
func (r *UserService) Get(ctx context.Context, userID string, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Merges two users together
func (r *UserService) Merge(ctx context.Context, userID string, body UserMergeParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/merge", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A user object
type User struct {
	ID          string                 `json:"id,required"`
	Typename    string                 `json:"__typename,required"`
	UpdatedAt   time.Time              `json:"updated_at,required" format:"date-time"`
	Avatar      string                 `json:"avatar,nullable"`
	CreatedAt   time.Time              `json:"created_at,nullable" format:"date-time"`
	Email       string                 `json:"email,nullable" format:"email"`
	Name        string                 `json:"name,nullable"`
	PhoneNumber string                 `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                 `json:"timezone,nullable"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        userJSON               `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

// A paginated list of users.
type UserListResponse struct {
	// The list of users
	Entries []User `json:"entries"`
	// The information about a paginated result
	PageInfo UserListResponsePageInfo `json:"page_info"`
	JSON     userListResponseJSON     `json:"-"`
}

// userListResponseJSON contains the JSON metadata for the struct
// [UserListResponse]
type userListResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseJSON) RawJSON() string {
	return r.raw
}

// The information about a paginated result
type UserListResponsePageInfo struct {
	Typename string                       `json:"__typename,required"`
	PageSize int64                        `json:"page_size,required"`
	After    string                       `json:"after,nullable"`
	Before   string                       `json:"before,nullable"`
	JSON     userListResponsePageInfoJSON `json:"-"`
}

// userListResponsePageInfoJSON contains the JSON metadata for the struct
// [UserListResponsePageInfo]
type userListResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

type UserUpdateParams struct {
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]UserUpdateParamsChannelData] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                              `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]UserUpdateParamsPreferences] `json:"preferences"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set channel data for a type of channel
type UserUpdateParamsChannelData struct {
	// Channel data for push providers
	Data param.Field[UserUpdateParamsChannelDataDataUnion] `json:"data,required"`
}

func (r UserUpdateParamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type UserUpdateParamsChannelDataData struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r UserUpdateParamsChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataData) implementsUserUpdateParamsChannelDataDataUnion() {}

// Channel data for push providers
//
// Satisfied by [UserUpdateParamsChannelDataDataPushChannelData],
// [UserUpdateParamsChannelDataDataOneSignalChannelData],
// [UserUpdateParamsChannelDataDataSlackChannelData],
// [UserUpdateParamsChannelDataDataMsTeamsChannelData],
// [UserUpdateParamsChannelDataDataDiscordChannelData],
// [UserUpdateParamsChannelDataData].
type UserUpdateParamsChannelDataDataUnion interface {
	implementsUserUpdateParamsChannelDataDataUnion()
}

// Channel data for push providers
type UserUpdateParamsChannelDataDataPushChannelData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r UserUpdateParamsChannelDataDataPushChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataPushChannelData) implementsUserUpdateParamsChannelDataDataUnion() {
}

// OneSignal channel data
type UserUpdateParamsChannelDataDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r UserUpdateParamsChannelDataDataOneSignalChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataOneSignalChannelData) implementsUserUpdateParamsChannelDataDataUnion() {
}

// Slack channel data
type UserUpdateParamsChannelDataDataSlackChannelData struct {
	Connections param.Field[[]UserUpdateParamsChannelDataDataSlackChannelDataConnectionUnion] `json:"connections,required"`
	Token       param.Field[UserUpdateParamsChannelDataDataSlackChannelDataToken]             `json:"token"`
}

func (r UserUpdateParamsChannelDataDataSlackChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataSlackChannelData) implementsUserUpdateParamsChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type UserUpdateParamsChannelDataDataSlackChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r UserUpdateParamsChannelDataDataSlackChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataSlackChannelDataConnection) implementsUserUpdateParamsChannelDataDataSlackChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [UserUpdateParamsChannelDataDataSlackChannelDataConnectionsTokenConnection],
// [UserUpdateParamsChannelDataDataSlackChannelDataConnectionsIncomingWebhookConnection],
// [UserUpdateParamsChannelDataDataSlackChannelDataConnection].
type UserUpdateParamsChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsUserUpdateParamsChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type UserUpdateParamsChannelDataDataSlackChannelDataConnectionsTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r UserUpdateParamsChannelDataDataSlackChannelDataConnectionsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataSlackChannelDataConnectionsTokenConnection) implementsUserUpdateParamsChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type UserUpdateParamsChannelDataDataSlackChannelDataConnectionsIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r UserUpdateParamsChannelDataDataSlackChannelDataConnectionsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataSlackChannelDataConnectionsIncomingWebhookConnection) implementsUserUpdateParamsChannelDataDataSlackChannelDataConnectionUnion() {
}

type UserUpdateParamsChannelDataDataSlackChannelDataToken struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r UserUpdateParamsChannelDataDataSlackChannelDataToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type UserUpdateParamsChannelDataDataMsTeamsChannelData struct {
	Connections param.Field[[]UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionUnion] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r UserUpdateParamsChannelDataDataMsTeamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataMsTeamsChannelData) implementsUserUpdateParamsChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type UserUpdateParamsChannelDataDataMsTeamsChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r UserUpdateParamsChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataMsTeamsChannelDataConnection) implementsUserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionsTokenConnection],
// [UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionsIncomingWebhookConnection],
// [UserUpdateParamsChannelDataDataMsTeamsChannelDataConnection].
type UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsUserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionsTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionsTokenConnection) implementsUserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionsIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionsIncomingWebhookConnection) implementsUserUpdateParamsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Discord channel data
type UserUpdateParamsChannelDataDataDiscordChannelData struct {
	Connections param.Field[[]UserUpdateParamsChannelDataDataDiscordChannelDataConnectionUnion] `json:"connections,required"`
}

func (r UserUpdateParamsChannelDataDataDiscordChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataDiscordChannelData) implementsUserUpdateParamsChannelDataDataUnion() {
}

// Discord channel connection
type UserUpdateParamsChannelDataDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id"`
	URL       param.Field[string] `json:"url"`
}

func (r UserUpdateParamsChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataDiscordChannelDataConnection) implementsUserUpdateParamsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [UserUpdateParamsChannelDataDataDiscordChannelDataConnectionsChannelConnection],
// [UserUpdateParamsChannelDataDataDiscordChannelDataConnectionsIncomingWebhookConnection],
// [UserUpdateParamsChannelDataDataDiscordChannelDataConnection].
type UserUpdateParamsChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsUserUpdateParamsChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type UserUpdateParamsChannelDataDataDiscordChannelDataConnectionsChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r UserUpdateParamsChannelDataDataDiscordChannelDataConnectionsChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataDiscordChannelDataConnectionsChannelConnection) implementsUserUpdateParamsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type UserUpdateParamsChannelDataDataDiscordChannelDataConnectionsIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r UserUpdateParamsChannelDataDataDiscordChannelDataConnectionsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsChannelDataDataDiscordChannelDataConnectionsIncomingWebhookConnection) implementsUserUpdateParamsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Set preferences for a recipient
type UserUpdateParamsPreferences struct {
	Categories param.Field[map[string]UserUpdateParamsPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[UserUpdateParamsPreferencesChannelTypes]              `json:"channel_types"`
	Workflows    param.Field[map[string]UserUpdateParamsPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r UserUpdateParamsPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool], [UserUpdateParamsPreferencesCategoriesObject].
type UserUpdateParamsPreferencesCategoriesUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesUnion()
}

type UserUpdateParamsPreferencesCategoriesObject struct {
	// Channel type preferences
	ChannelTypes param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserUpdateParamsPreferencesCategoriesObjectCondition]  `json:"conditions"`
}

func (r UserUpdateParamsPreferencesCategoriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObject) ImplementsUserUpdateParamsPreferencesCategoriesUnion() {
}

// Channel type preferences
type UserUpdateParamsPreferencesCategoriesObjectChannelTypes struct {
	Chat      param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion]      `json:"http"`
	InApp     param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppUnion]     `json:"in_app"`
	InAppFeed param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion]       `json:"sms"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditions].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditions) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                                  `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                  `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditions].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditions) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                                   `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                   `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditions].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditions) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                                  `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                  `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditions].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditions) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsCondition struct {
	Argument param.Field[string]                                                                                   `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                   `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditions].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditions) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                                       `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                       `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditions].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditions) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                                  `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                  `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditions].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditions) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                                 `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                 `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserUpdateParamsPreferencesCategoriesObjectCondition struct {
	Argument param.Field[string]                                                        `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                        `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesCategoriesObjectConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesCategoriesObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorContainsAll, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type UserUpdateParamsPreferencesChannelTypes struct {
	Chat      param.Field[UserUpdateParamsPreferencesChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[UserUpdateParamsPreferencesChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[UserUpdateParamsPreferencesChannelTypesHTTPUnion]      `json:"http"`
	InApp     param.Field[UserUpdateParamsPreferencesChannelTypesInAppUnion]     `json:"in_app"`
	InAppFeed param.Field[UserUpdateParamsPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[UserUpdateParamsPreferencesChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[UserUpdateParamsPreferencesChannelTypesSMSUnion]       `json:"sms"`
}

func (r UserUpdateParamsPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesChatConditions].
type UserUpdateParamsPreferencesChannelTypesChatUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesChatUnion()
}

type UserUpdateParamsPreferencesChannelTypesChatConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesChatConditions) ImplementsUserUpdateParamsPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                  `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                  `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesEmailConditions].
type UserUpdateParamsPreferencesChannelTypesEmailUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesEmailUnion()
}

type UserUpdateParamsPreferencesChannelTypesEmailConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesEmailConditions) ImplementsUserUpdateParamsPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                   `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                   `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesHTTPConditions].
type UserUpdateParamsPreferencesChannelTypesHTTPUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesHTTPUnion()
}

type UserUpdateParamsPreferencesChannelTypesHTTPConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesHTTPConditions) ImplementsUserUpdateParamsPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                  `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                  `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesInAppConditions].
type UserUpdateParamsPreferencesChannelTypesInAppUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesInAppUnion()
}

type UserUpdateParamsPreferencesChannelTypesInAppConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesInAppConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesInAppConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesInAppConditions) ImplementsUserUpdateParamsPreferencesChannelTypesInAppUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesChannelTypesInAppConditionsCondition struct {
	Argument param.Field[string]                                                                   `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                   `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesInAppConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesChannelTypesInAppConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesInAppFeedConditions].
type UserUpdateParamsPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesInAppFeedUnion()
}

type UserUpdateParamsPreferencesChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesInAppFeedConditions) ImplementsUserUpdateParamsPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                       `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                       `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesPushConditions].
type UserUpdateParamsPreferencesChannelTypesPushUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesPushUnion()
}

type UserUpdateParamsPreferencesChannelTypesPushConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesPushConditions) ImplementsUserUpdateParamsPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                  `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                  `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesSMSConditions].
type UserUpdateParamsPreferencesChannelTypesSMSUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesSMSUnion()
}

type UserUpdateParamsPreferencesChannelTypesSMSConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesSMSConditions) ImplementsUserUpdateParamsPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                 `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                 `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool], [UserUpdateParamsPreferencesWorkflowsObject].
type UserUpdateParamsPreferencesWorkflowsUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsUnion()
}

type UserUpdateParamsPreferencesWorkflowsObject struct {
	// Channel type preferences
	ChannelTypes param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectCondition]  `json:"conditions"`
}

func (r UserUpdateParamsPreferencesWorkflowsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObject) ImplementsUserUpdateParamsPreferencesWorkflowsUnion() {
}

// Channel type preferences
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypes struct {
	Chat      param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion]      `json:"http"`
	InApp     param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppUnion]     `json:"in_app"`
	InAppFeed param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSUnion]       `json:"sms"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditions].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditions) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                                 `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                 `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditions].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditions) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                                  `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                  `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditions].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditions) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                                 `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                 `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditions].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditions) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsCondition struct {
	Argument param.Field[string]                                                                                  `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                  `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditions].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditions) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                                      `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                      `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditions].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditions) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                                 `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                 `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditions].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditions struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditions) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                                `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserUpdateParamsPreferencesWorkflowsObjectCondition struct {
	Argument param.Field[string]                                                       `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                       `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "contains_all"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestamp          UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "is_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsNotTimestamp       UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "is_not_timestamp"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampAfter     UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "is_timestamp_after"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampBefore    UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "is_timestamp_before"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampBetween   UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "is_timestamp_between"
	UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsAudienceMember     UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator = "is_audience_member"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorContainsAll, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestamp, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsNotTimestamp, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampAfter, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampBefore, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsTimestampBetween, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

type UserListParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserMergeParams struct {
	FromUserID param.Field[string] `json:"from_user_id"`
}

func (r UserMergeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

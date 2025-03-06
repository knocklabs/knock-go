// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
	"github.com/tidwall/gjson"
)

// UserService contains methods and other services that help with interacting with
// the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
	Feeds   *UserFeedService
	Bulk    *UserBulkService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.Feeds = NewUserFeedService(opts...)
	r.Bulk = NewUserBulkService(opts...)
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

// Get channel data for a user
func (r *UserService) GetChannelData(ctx context.Context, userID string, channelID string, opts ...option.RequestOption) (res *UserGetChannelDataResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/channel_data/%s", userID, channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a preference set for a user
func (r *UserService) GetPreferences(ctx context.Context, userID string, id string, query UserGetPreferencesParams, opts ...option.RequestOption) (res *UserGetPreferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/preferences/%s", userID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a paginated list of messages for a user
func (r *UserService) ListMessages(ctx context.Context, userID string, query UserListMessagesParams, opts ...option.RequestOption) (res *UserListMessagesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/messages", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List preference sets for a user
func (r *UserService) ListPreferences(ctx context.Context, userID string, opts ...option.RequestOption) (res *[]UserListPreferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/preferences", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List schedules for a user
func (r *UserService) ListSchedules(ctx context.Context, userID string, query UserListSchedulesParams, opts ...option.RequestOption) (res *UserListSchedulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/schedules", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List subscriptions for a user
func (r *UserService) ListSubscriptions(ctx context.Context, userID string, query UserListSubscriptionsParams, opts ...option.RequestOption) (res *UserListSubscriptionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/subscriptions", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

// Sets channel data for a user
func (r *UserService) SetChannelData(ctx context.Context, userID string, channelID string, opts ...option.RequestOption) (res *UserSetChannelDataResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/channel_data/%s", userID, channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Update a preference set for a user
func (r *UserService) SetPreferences(ctx context.Context, userID string, id string, body UserSetPreferencesParams, opts ...option.RequestOption) (res *UserSetPreferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/preferences/%s", userID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Unsets channel data for a user
func (r *UserService) UnsetChannelData(ctx context.Context, userID string, channelID string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/channel_data/%s", userID, channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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

func (r User) implementsUserListSchedulesResponseEntriesRecipient() {}

func (r User) implementsUserListSchedulesResponseEntriesActor() {}

func (r User) implementsUserListSubscriptionsResponseEntriesRecipient() {}

func (r User) implementsUserFeedGetResponseEntriesActivitiesActor() {}

func (r User) implementsUserFeedGetResponseEntriesActivitiesRecipient() {}

func (r User) implementsUserFeedGetResponseEntriesActor() {}

func (r User) implementsObjectAddSubscriptionsResponseRecipient() {}

func (r User) implementsObjectDeleteSubscriptionsResponseRecipient() {}

func (r User) implementsObjectListSchedulesResponseEntriesRecipient() {}

func (r User) implementsObjectListSchedulesResponseEntriesActor() {}

func (r User) implementsObjectListSubscriptionsResponseEntriesRecipient() {}

func (r User) implementsMessageListActivitiesResponseItemsActor() {}

func (r User) implementsMessageListActivitiesResponseItemsRecipient() {}

func (r User) implementsScheduleNewResponseRecipient() {}

func (r User) implementsScheduleNewResponseActor() {}

func (r User) implementsScheduleUpdateResponseRecipient() {}

func (r User) implementsScheduleUpdateResponseActor() {}

func (r User) implementsScheduleListResponseEntriesRecipient() {}

func (r User) implementsScheduleListResponseEntriesActor() {}

func (r User) implementsScheduleDeleteResponseRecipient() {}

func (r User) implementsScheduleDeleteResponseActor() {}

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

// Channel data for various channel types
type UserGetChannelDataResponse struct {
	Typename  string `json:"__typename,required"`
	ChannelID string `json:"channel_id,required" format:"uuid"`
	// Channel data for push providers
	Data UserGetChannelDataResponseData `json:"data,required"`
	JSON userGetChannelDataResponseJSON `json:"-"`
}

// userGetChannelDataResponseJSON contains the JSON metadata for the struct
// [UserGetChannelDataResponse]
type userGetChannelDataResponseJSON struct {
	Typename    apijson.Field
	ChannelID   apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseJSON) RawJSON() string {
	return r.raw
}

// Channel data for push providers
type UserGetChannelDataResponseData struct {
	// This field can have the runtime type of
	// [UserGetChannelDataResponseDataSlackChannelDataToken].
	Token interface{} `json:"token"`
	// This field can have the runtime type of
	// [[]UserGetChannelDataResponseDataSlackChannelDataConnection],
	// [[]UserGetChannelDataResponseDataMsTeamsChannelDataConnection],
	// [[]UserGetChannelDataResponseDataDiscordChannelDataConnection].
	Connections interface{} `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID string `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	// This field can have the runtime type of [[]string].
	PlayerIDs interface{} `json:"player_ids"`
	// This field can have the runtime type of [[]string].
	Tokens interface{}                        `json:"tokens"`
	JSON   userGetChannelDataResponseDataJSON `json:"-"`
	union  UserGetChannelDataResponseDataUnion
}

// userGetChannelDataResponseDataJSON contains the JSON metadata for the struct
// [UserGetChannelDataResponseData]
type userGetChannelDataResponseDataJSON struct {
	Token           apijson.Field
	Connections     apijson.Field
	MsTeamsTenantID apijson.Field
	PlayerIDs       apijson.Field
	Tokens          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r userGetChannelDataResponseDataJSON) RawJSON() string {
	return r.raw
}

func (r *UserGetChannelDataResponseData) UnmarshalJSON(data []byte) (err error) {
	*r = UserGetChannelDataResponseData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserGetChannelDataResponseDataUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserGetChannelDataResponseDataPushChannelData],
// [UserGetChannelDataResponseDataSlackChannelData],
// [UserGetChannelDataResponseDataMsTeamsChannelData],
// [UserGetChannelDataResponseDataDiscordChannelData],
// [UserGetChannelDataResponseDataOneSignalChannelData].
func (r UserGetChannelDataResponseData) AsUnion() UserGetChannelDataResponseDataUnion {
	return r.union
}

// Channel data for push providers
//
// Union satisfied by [UserGetChannelDataResponseDataPushChannelData],
// [UserGetChannelDataResponseDataSlackChannelData],
// [UserGetChannelDataResponseDataMsTeamsChannelData],
// [UserGetChannelDataResponseDataDiscordChannelData] or
// [UserGetChannelDataResponseDataOneSignalChannelData].
type UserGetChannelDataResponseDataUnion interface {
	implementsUserGetChannelDataResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetChannelDataResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataPushChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataSlackChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataMsTeamsChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataDiscordChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataOneSignalChannelData{}),
		},
	)
}

// Channel data for push providers
type UserGetChannelDataResponseDataPushChannelData struct {
	Tokens []string                                          `json:"tokens,required"`
	JSON   userGetChannelDataResponseDataPushChannelDataJSON `json:"-"`
}

// userGetChannelDataResponseDataPushChannelDataJSON contains the JSON metadata for
// the struct [UserGetChannelDataResponseDataPushChannelData]
type userGetChannelDataResponseDataPushChannelDataJSON struct {
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataPushChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataPushChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataPushChannelData) implementsUserGetChannelDataResponseData() {}

// Slack channel data
type UserGetChannelDataResponseDataSlackChannelData struct {
	Connections []UserGetChannelDataResponseDataSlackChannelDataConnection `json:"connections,required"`
	Token       UserGetChannelDataResponseDataSlackChannelDataToken        `json:"token,nullable"`
	JSON        userGetChannelDataResponseDataSlackChannelDataJSON         `json:"-"`
}

// userGetChannelDataResponseDataSlackChannelDataJSON contains the JSON metadata
// for the struct [UserGetChannelDataResponseDataSlackChannelData]
type userGetChannelDataResponseDataSlackChannelDataJSON struct {
	Connections apijson.Field
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataSlackChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataSlackChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataSlackChannelData) implementsUserGetChannelDataResponseData() {}

// A Slack connection, which either includes a channel_id or a user_id
type UserGetChannelDataResponseDataSlackChannelDataConnection struct {
	AccessToken string                                                       `json:"access_token,nullable"`
	ChannelID   string                                                       `json:"channel_id,nullable"`
	URL         string                                                       `json:"url"`
	UserID      string                                                       `json:"user_id,nullable"`
	JSON        userGetChannelDataResponseDataSlackChannelDataConnectionJSON `json:"-"`
	union       UserGetChannelDataResponseDataSlackChannelDataConnectionsUnion
}

// userGetChannelDataResponseDataSlackChannelDataConnectionJSON contains the JSON
// metadata for the struct
// [UserGetChannelDataResponseDataSlackChannelDataConnection]
type userGetChannelDataResponseDataSlackChannelDataConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	URL         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userGetChannelDataResponseDataSlackChannelDataConnectionJSON) RawJSON() string {
	return r.raw
}

func (r *UserGetChannelDataResponseDataSlackChannelDataConnection) UnmarshalJSON(data []byte) (err error) {
	*r = UserGetChannelDataResponseDataSlackChannelDataConnection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [UserGetChannelDataResponseDataSlackChannelDataConnectionsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection],
// [UserGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection].
func (r UserGetChannelDataResponseDataSlackChannelDataConnection) AsUnion() UserGetChannelDataResponseDataSlackChannelDataConnectionsUnion {
	return r.union
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Union satisfied by
// [UserGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection] or
// [UserGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection].
type UserGetChannelDataResponseDataSlackChannelDataConnectionsUnion interface {
	implementsUserGetChannelDataResponseDataSlackChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetChannelDataResponseDataSlackChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection{}),
		},
	)
}

// A Slack connection, which either includes a channel_id or a user_id
type UserGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection struct {
	AccessToken string                                                                       `json:"access_token,nullable"`
	ChannelID   string                                                                       `json:"channel_id,nullable"`
	UserID      string                                                                       `json:"user_id,nullable"`
	JSON        userGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnectionJSON `json:"-"`
}

// userGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnectionJSON
// contains the JSON metadata for the struct
// [UserGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection]
type userGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection) implementsUserGetChannelDataResponseDataSlackChannelDataConnection() {
}

// An incoming webhook Slack connection
type UserGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection struct {
	URL  string                                                                                 `json:"url,required"`
	JSON userGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnectionJSON `json:"-"`
}

// userGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnectionJSON
// contains the JSON metadata for the struct
// [UserGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection]
type userGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection) implementsUserGetChannelDataResponseDataSlackChannelDataConnection() {
}

type UserGetChannelDataResponseDataSlackChannelDataToken struct {
	AccessToken string                                                  `json:"access_token,required,nullable"`
	JSON        userGetChannelDataResponseDataSlackChannelDataTokenJSON `json:"-"`
}

// userGetChannelDataResponseDataSlackChannelDataTokenJSON contains the JSON
// metadata for the struct [UserGetChannelDataResponseDataSlackChannelDataToken]
type userGetChannelDataResponseDataSlackChannelDataTokenJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataSlackChannelDataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataSlackChannelDataTokenJSON) RawJSON() string {
	return r.raw
}

// Microsoft Teams channel data
type UserGetChannelDataResponseDataMsTeamsChannelData struct {
	Connections []UserGetChannelDataResponseDataMsTeamsChannelDataConnection `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID string                                               `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	JSON            userGetChannelDataResponseDataMsTeamsChannelDataJSON `json:"-"`
}

// userGetChannelDataResponseDataMsTeamsChannelDataJSON contains the JSON metadata
// for the struct [UserGetChannelDataResponseDataMsTeamsChannelData]
type userGetChannelDataResponseDataMsTeamsChannelDataJSON struct {
	Connections     apijson.Field
	MsTeamsTenantID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataMsTeamsChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataMsTeamsChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataMsTeamsChannelData) implementsUserGetChannelDataResponseData() {
}

// A Slack connection, which either includes a channel_id or a user_id
type UserGetChannelDataResponseDataMsTeamsChannelDataConnection struct {
	AccessToken string                                                         `json:"access_token,nullable"`
	ChannelID   string                                                         `json:"channel_id,nullable"`
	URL         string                                                         `json:"url"`
	UserID      string                                                         `json:"user_id,nullable"`
	JSON        userGetChannelDataResponseDataMsTeamsChannelDataConnectionJSON `json:"-"`
	union       UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion
}

// userGetChannelDataResponseDataMsTeamsChannelDataConnectionJSON contains the JSON
// metadata for the struct
// [UserGetChannelDataResponseDataMsTeamsChannelDataConnection]
type userGetChannelDataResponseDataMsTeamsChannelDataConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	URL         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userGetChannelDataResponseDataMsTeamsChannelDataConnectionJSON) RawJSON() string {
	return r.raw
}

func (r *UserGetChannelDataResponseDataMsTeamsChannelDataConnection) UnmarshalJSON(data []byte) (err error) {
	*r = UserGetChannelDataResponseDataMsTeamsChannelDataConnection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection],
// [UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection].
func (r UserGetChannelDataResponseDataMsTeamsChannelDataConnection) AsUnion() UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion {
	return r.union
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Union satisfied by
// [UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection] or
// [UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection].
type UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion interface {
	implementsUserGetChannelDataResponseDataMsTeamsChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection{}),
		},
	)
}

// A Slack connection, which either includes a channel_id or a user_id
type UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection struct {
	AccessToken string                                                                         `json:"access_token,nullable"`
	ChannelID   string                                                                         `json:"channel_id,nullable"`
	UserID      string                                                                         `json:"user_id,nullable"`
	JSON        userGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnectionJSON `json:"-"`
}

// userGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnectionJSON
// contains the JSON metadata for the struct
// [UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection]
type userGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection) implementsUserGetChannelDataResponseDataMsTeamsChannelDataConnection() {
}

// An incoming webhook Slack connection
type UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection struct {
	URL  string                                                                                   `json:"url,required"`
	JSON userGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnectionJSON `json:"-"`
}

// userGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnectionJSON
// contains the JSON metadata for the struct
// [UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection]
type userGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection) implementsUserGetChannelDataResponseDataMsTeamsChannelDataConnection() {
}

// Discord channel data
type UserGetChannelDataResponseDataDiscordChannelData struct {
	Connections []UserGetChannelDataResponseDataDiscordChannelDataConnection `json:"connections,required"`
	JSON        userGetChannelDataResponseDataDiscordChannelDataJSON         `json:"-"`
}

// userGetChannelDataResponseDataDiscordChannelDataJSON contains the JSON metadata
// for the struct [UserGetChannelDataResponseDataDiscordChannelData]
type userGetChannelDataResponseDataDiscordChannelDataJSON struct {
	Connections apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataDiscordChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataDiscordChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataDiscordChannelData) implementsUserGetChannelDataResponseData() {
}

// Discord channel connection
type UserGetChannelDataResponseDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID string                                                         `json:"channel_id"`
	URL       string                                                         `json:"url"`
	JSON      userGetChannelDataResponseDataDiscordChannelDataConnectionJSON `json:"-"`
	union     UserGetChannelDataResponseDataDiscordChannelDataConnectionsUnion
}

// userGetChannelDataResponseDataDiscordChannelDataConnectionJSON contains the JSON
// metadata for the struct
// [UserGetChannelDataResponseDataDiscordChannelDataConnection]
type userGetChannelDataResponseDataDiscordChannelDataConnectionJSON struct {
	ChannelID   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userGetChannelDataResponseDataDiscordChannelDataConnectionJSON) RawJSON() string {
	return r.raw
}

func (r *UserGetChannelDataResponseDataDiscordChannelDataConnection) UnmarshalJSON(data []byte) (err error) {
	*r = UserGetChannelDataResponseDataDiscordChannelDataConnection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [UserGetChannelDataResponseDataDiscordChannelDataConnectionsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection],
// [UserGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection].
func (r UserGetChannelDataResponseDataDiscordChannelDataConnection) AsUnion() UserGetChannelDataResponseDataDiscordChannelDataConnectionsUnion {
	return r.union
}

// Discord channel connection
//
// Union satisfied by
// [UserGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection]
// or
// [UserGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection].
type UserGetChannelDataResponseDataDiscordChannelDataConnectionsUnion interface {
	implementsUserGetChannelDataResponseDataDiscordChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetChannelDataResponseDataDiscordChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection{}),
		},
	)
}

// Discord channel connection
type UserGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection struct {
	// The Discord channel ID
	ChannelID string                                                                           `json:"channel_id,required"`
	JSON      userGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnectionJSON `json:"-"`
}

// userGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnectionJSON
// contains the JSON metadata for the struct
// [UserGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection]
type userGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnectionJSON struct {
	ChannelID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection) implementsUserGetChannelDataResponseDataDiscordChannelDataConnection() {
}

// An incoming webhook Slack connection
type UserGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection struct {
	URL  string                                                                                   `json:"url,required"`
	JSON userGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnectionJSON `json:"-"`
}

// userGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnectionJSON
// contains the JSON metadata for the struct
// [UserGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection]
type userGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection) implementsUserGetChannelDataResponseDataDiscordChannelDataConnection() {
}

// OneSignal channel data
type UserGetChannelDataResponseDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs []string                                               `json:"player_ids,required" format:"uuid"`
	JSON      userGetChannelDataResponseDataOneSignalChannelDataJSON `json:"-"`
}

// userGetChannelDataResponseDataOneSignalChannelDataJSON contains the JSON
// metadata for the struct [UserGetChannelDataResponseDataOneSignalChannelData]
type userGetChannelDataResponseDataOneSignalChannelDataJSON struct {
	PlayerIDs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetChannelDataResponseDataOneSignalChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetChannelDataResponseDataOneSignalChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserGetChannelDataResponseDataOneSignalChannelData) implementsUserGetChannelDataResponseData() {
}

// A preference set object.
type UserGetPreferencesResponse struct {
	ID         string                                               `json:"id,required"`
	Typename   string                                               `json:"__typename,required"`
	Categories map[string]UserGetPreferencesResponseCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes UserGetPreferencesResponseChannelTypes              `json:"channel_types,nullable"`
	Workflows    map[string]UserGetPreferencesResponseWorkflowsUnion `json:"workflows,nullable"`
	JSON         userGetPreferencesResponseJSON                      `json:"-"`
}

// userGetPreferencesResponseJSON contains the JSON metadata for the struct
// [UserGetPreferencesResponse]
type userGetPreferencesResponseJSON struct {
	ID           apijson.Field
	Typename     apijson.Field
	Categories   apijson.Field
	ChannelTypes apijson.Field
	Workflows    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserGetPreferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseJSON) RawJSON() string {
	return r.raw
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseCategoriesObject].
type UserGetPreferencesResponseCategoriesUnion interface {
	ImplementsUserGetPreferencesResponseCategoriesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseCategoriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseCategoriesObject{}),
		},
	)
}

type UserGetPreferencesResponseCategoriesObject struct {
	// Channel type preferences
	ChannelTypes UserGetPreferencesResponseCategoriesObjectChannelTypes `json:"channel_types,nullable"`
	Conditions   []UserGetPreferencesResponseCategoriesObjectCondition  `json:"conditions"`
	JSON         userGetPreferencesResponseCategoriesObjectJSON         `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectJSON contains the JSON metadata for
// the struct [UserGetPreferencesResponseCategoriesObject]
type userGetPreferencesResponseCategoriesObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseCategoriesObject) ImplementsUserGetPreferencesResponseCategoriesUnion() {
}

// Channel type preferences
type UserGetPreferencesResponseCategoriesObjectChannelTypes struct {
	Chat      UserGetPreferencesResponseCategoriesObjectChannelTypesChatUnion      `json:"chat"`
	Email     UserGetPreferencesResponseCategoriesObjectChannelTypesEmailUnion     `json:"email"`
	HTTP      UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion      `json:"http"`
	InAppFeed UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      UserGetPreferencesResponseCategoriesObjectChannelTypesPushUnion      `json:"push"`
	SMS       UserGetPreferencesResponseCategoriesObjectChannelTypesSMSUnion       `json:"sms"`
	JSON      userGetPreferencesResponseCategoriesObjectChannelTypesJSON           `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesJSON contains the JSON
// metadata for the struct [UserGetPreferencesResponseCategoriesObjectChannelTypes]
type userGetPreferencesResponseCategoriesObjectChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditions].
type UserGetPreferencesResponseCategoriesObjectChannelTypesChatUnion interface {
	ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseCategoriesObjectChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditions{}),
		},
	)
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditions struct {
	Conditions []UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditions]
type userGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditions) ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition]
type userGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains             UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditions].
type UserGetPreferencesResponseCategoriesObjectChannelTypesEmailUnion interface {
	ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseCategoriesObjectChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditions{}),
		},
	)
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditions struct {
	Conditions []UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditions]
type userGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditions) ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition struct {
	Argument string                                                                                  `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                  `json:"variable,required"`
	JSON     userGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition]
type userGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains             UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions].
type UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion interface {
	ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions{}),
		},
	)
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions struct {
	Conditions []UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions]
type userGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions) ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition]
type userGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions].
type UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions{}),
		},
	)
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions struct {
	Conditions []UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions]
type userGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions) ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument string                                                                                      `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                      `json:"variable,required"`
	JSON     userGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition]
type userGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditions].
type UserGetPreferencesResponseCategoriesObjectChannelTypesPushUnion interface {
	ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseCategoriesObjectChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditions{}),
		},
	)
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditions struct {
	Conditions []UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditions]
type userGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditions) ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition]
type userGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains             UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditions].
type UserGetPreferencesResponseCategoriesObjectChannelTypesSMSUnion interface {
	ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseCategoriesObjectChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditions{}),
		},
	)
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditions struct {
	Conditions []UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON contains
// the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditions]
type userGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditions) ImplementsUserGetPreferencesResponseCategoriesObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition struct {
	Argument string                                                                                `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                `json:"variable,required"`
	JSON     userGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition]
type userGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains             UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserGetPreferencesResponseCategoriesObjectCondition struct {
	Argument string                                                       `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseCategoriesObjectConditionsOperator `json:"operator,required"`
	Variable string                                                       `json:"variable,required"`
	JSON     userGetPreferencesResponseCategoriesObjectConditionJSON      `json:"-"`
}

// userGetPreferencesResponseCategoriesObjectConditionJSON contains the JSON
// metadata for the struct [UserGetPreferencesResponseCategoriesObjectCondition]
type userGetPreferencesResponseCategoriesObjectConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesObjectCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesObjectConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseCategoriesObjectConditionsOperator string

const (
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorEqualTo              UserGetPreferencesResponseCategoriesObjectConditionsOperator = "equal_to"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorNotEqualTo           UserGetPreferencesResponseCategoriesObjectConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorGreaterThan          UserGetPreferencesResponseCategoriesObjectConditionsOperator = "greater_than"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorLessThan             UserGetPreferencesResponseCategoriesObjectConditionsOperator = "less_than"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseCategoriesObjectConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseCategoriesObjectConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorContains             UserGetPreferencesResponseCategoriesObjectConditionsOperator = "contains"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorNotContains          UserGetPreferencesResponseCategoriesObjectConditionsOperator = "not_contains"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorEmpty                UserGetPreferencesResponseCategoriesObjectConditionsOperator = "empty"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorNotEmpty             UserGetPreferencesResponseCategoriesObjectConditionsOperator = "not_empty"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorContainsAll          UserGetPreferencesResponseCategoriesObjectConditionsOperator = "contains_all"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestamp          UserGetPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseCategoriesObjectConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsAudienceMember     UserGetPreferencesResponseCategoriesObjectConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseCategoriesObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseCategoriesObjectConditionsOperatorEqualTo, UserGetPreferencesResponseCategoriesObjectConditionsOperatorNotEqualTo, UserGetPreferencesResponseCategoriesObjectConditionsOperatorGreaterThan, UserGetPreferencesResponseCategoriesObjectConditionsOperatorLessThan, UserGetPreferencesResponseCategoriesObjectConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseCategoriesObjectConditionsOperatorContains, UserGetPreferencesResponseCategoriesObjectConditionsOperatorNotContains, UserGetPreferencesResponseCategoriesObjectConditionsOperatorEmpty, UserGetPreferencesResponseCategoriesObjectConditionsOperatorNotEmpty, UserGetPreferencesResponseCategoriesObjectConditionsOperatorContainsAll, UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestamp, UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseCategoriesObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type UserGetPreferencesResponseChannelTypes struct {
	Chat      UserGetPreferencesResponseChannelTypesChatUnion      `json:"chat"`
	Email     UserGetPreferencesResponseChannelTypesEmailUnion     `json:"email"`
	HTTP      UserGetPreferencesResponseChannelTypesHTTPUnion      `json:"http"`
	InAppFeed UserGetPreferencesResponseChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      UserGetPreferencesResponseChannelTypesPushUnion      `json:"push"`
	SMS       UserGetPreferencesResponseChannelTypesSMSUnion       `json:"sms"`
	JSON      userGetPreferencesResponseChannelTypesJSON           `json:"-"`
}

// userGetPreferencesResponseChannelTypesJSON contains the JSON metadata for the
// struct [UserGetPreferencesResponseChannelTypes]
type userGetPreferencesResponseChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseChannelTypesChatConditions].
type UserGetPreferencesResponseChannelTypesChatUnion interface {
	ImplementsUserGetPreferencesResponseChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseChannelTypesChatConditions{}),
		},
	)
}

type UserGetPreferencesResponseChannelTypesChatConditions struct {
	Conditions []UserGetPreferencesResponseChannelTypesChatConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseChannelTypesChatConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseChannelTypesChatConditionsJSON contains the JSON
// metadata for the struct [UserGetPreferencesResponseChannelTypesChatConditions]
type userGetPreferencesResponseChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseChannelTypesChatConditions) ImplementsUserGetPreferencesResponseChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseChannelTypesChatConditionsCondition struct {
	Argument string                                                                 `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                 `json:"variable,required"`
	JSON     userGetPreferencesResponseChannelTypesChatConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseChannelTypesChatConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserGetPreferencesResponseChannelTypesChatConditionsCondition]
type userGetPreferencesResponseChannelTypesChatConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesChatConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesChatConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator string

const (
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThan             UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorContains             UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotContains          UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorEmpty                UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThan, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorContains, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotContains, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorEmpty, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseChannelTypesEmailConditions].
type UserGetPreferencesResponseChannelTypesEmailUnion interface {
	ImplementsUserGetPreferencesResponseChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseChannelTypesEmailConditions{}),
		},
	)
}

type UserGetPreferencesResponseChannelTypesEmailConditions struct {
	Conditions []UserGetPreferencesResponseChannelTypesEmailConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseChannelTypesEmailConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseChannelTypesEmailConditionsJSON contains the JSON
// metadata for the struct [UserGetPreferencesResponseChannelTypesEmailConditions]
type userGetPreferencesResponseChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseChannelTypesEmailConditions) ImplementsUserGetPreferencesResponseChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseChannelTypesEmailConditionsCondition struct {
	Argument string                                                                  `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                  `json:"variable,required"`
	JSON     userGetPreferencesResponseChannelTypesEmailConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseChannelTypesEmailConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserGetPreferencesResponseChannelTypesEmailConditionsCondition]
type userGetPreferencesResponseChannelTypesEmailConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesEmailConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesEmailConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator string

const (
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThan             UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContains             UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotContains          UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEmpty                UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThan, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContains, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotContains, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEmpty, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseChannelTypesHTTPConditions].
type UserGetPreferencesResponseChannelTypesHTTPUnion interface {
	ImplementsUserGetPreferencesResponseChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseChannelTypesHTTPConditions{}),
		},
	)
}

type UserGetPreferencesResponseChannelTypesHTTPConditions struct {
	Conditions []UserGetPreferencesResponseChannelTypesHTTPConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseChannelTypesHTTPConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseChannelTypesHTTPConditionsJSON contains the JSON
// metadata for the struct [UserGetPreferencesResponseChannelTypesHTTPConditions]
type userGetPreferencesResponseChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseChannelTypesHTTPConditions) ImplementsUserGetPreferencesResponseChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseChannelTypesHTTPConditionsCondition struct {
	Argument string                                                                 `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                 `json:"variable,required"`
	JSON     userGetPreferencesResponseChannelTypesHTTPConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseChannelTypesHTTPConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserGetPreferencesResponseChannelTypesHTTPConditionsCondition]
type userGetPreferencesResponseChannelTypesHTTPConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesHTTPConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesHTTPConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator string

const (
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThan             UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContains             UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotContains          UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEmpty                UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThan, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContains, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotContains, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEmpty, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseChannelTypesInAppFeedConditions].
type UserGetPreferencesResponseChannelTypesInAppFeedUnion interface {
	ImplementsUserGetPreferencesResponseChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseChannelTypesInAppFeedConditions{}),
		},
	)
}

type UserGetPreferencesResponseChannelTypesInAppFeedConditions struct {
	Conditions []UserGetPreferencesResponseChannelTypesInAppFeedConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseChannelTypesInAppFeedConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseChannelTypesInAppFeedConditionsJSON contains the JSON
// metadata for the struct
// [UserGetPreferencesResponseChannelTypesInAppFeedConditions]
type userGetPreferencesResponseChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseChannelTypesInAppFeedConditions) ImplementsUserGetPreferencesResponseChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseChannelTypesInAppFeedConditionsCondition struct {
	Argument string                                                                      `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                      `json:"variable,required"`
	JSON     userGetPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON contains
// the JSON metadata for the struct
// [UserGetPreferencesResponseChannelTypesInAppFeedConditionsCondition]
type userGetPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesInAppFeedConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContains             UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContains, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseChannelTypesPushConditions].
type UserGetPreferencesResponseChannelTypesPushUnion interface {
	ImplementsUserGetPreferencesResponseChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseChannelTypesPushConditions{}),
		},
	)
}

type UserGetPreferencesResponseChannelTypesPushConditions struct {
	Conditions []UserGetPreferencesResponseChannelTypesPushConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseChannelTypesPushConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseChannelTypesPushConditionsJSON contains the JSON
// metadata for the struct [UserGetPreferencesResponseChannelTypesPushConditions]
type userGetPreferencesResponseChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseChannelTypesPushConditions) ImplementsUserGetPreferencesResponseChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseChannelTypesPushConditionsCondition struct {
	Argument string                                                                 `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                 `json:"variable,required"`
	JSON     userGetPreferencesResponseChannelTypesPushConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseChannelTypesPushConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserGetPreferencesResponseChannelTypesPushConditionsCondition]
type userGetPreferencesResponseChannelTypesPushConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesPushConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesPushConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator string

const (
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThan             UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorContains             UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotContains          UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorEmpty                UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThan, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorContains, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotContains, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorEmpty, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseChannelTypesSMSConditions].
type UserGetPreferencesResponseChannelTypesSMSUnion interface {
	ImplementsUserGetPreferencesResponseChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseChannelTypesSMSConditions{}),
		},
	)
}

type UserGetPreferencesResponseChannelTypesSMSConditions struct {
	Conditions []UserGetPreferencesResponseChannelTypesSMSConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseChannelTypesSMSConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseChannelTypesSMSConditionsJSON contains the JSON
// metadata for the struct [UserGetPreferencesResponseChannelTypesSMSConditions]
type userGetPreferencesResponseChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseChannelTypesSMSConditions) ImplementsUserGetPreferencesResponseChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseChannelTypesSMSConditionsCondition struct {
	Argument string                                                                `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                `json:"variable,required"`
	JSON     userGetPreferencesResponseChannelTypesSMSConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseChannelTypesSMSConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserGetPreferencesResponseChannelTypesSMSConditionsCondition]
type userGetPreferencesResponseChannelTypesSMSConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseChannelTypesSMSConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseChannelTypesSMSConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator string

const (
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThan             UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContains             UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotContains          UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEmpty                UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThan, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContains, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotContains, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEmpty, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseWorkflowsObject].
type UserGetPreferencesResponseWorkflowsUnion interface {
	ImplementsUserGetPreferencesResponseWorkflowsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseWorkflowsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseWorkflowsObject{}),
		},
	)
}

type UserGetPreferencesResponseWorkflowsObject struct {
	// Channel type preferences
	ChannelTypes UserGetPreferencesResponseWorkflowsObjectChannelTypes `json:"channel_types,nullable"`
	Conditions   []UserGetPreferencesResponseWorkflowsObjectCondition  `json:"conditions"`
	JSON         userGetPreferencesResponseWorkflowsObjectJSON         `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectJSON contains the JSON metadata for the
// struct [UserGetPreferencesResponseWorkflowsObject]
type userGetPreferencesResponseWorkflowsObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseWorkflowsObject) ImplementsUserGetPreferencesResponseWorkflowsUnion() {
}

// Channel type preferences
type UserGetPreferencesResponseWorkflowsObjectChannelTypes struct {
	Chat      UserGetPreferencesResponseWorkflowsObjectChannelTypesChatUnion      `json:"chat"`
	Email     UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion     `json:"email"`
	HTTP      UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion      `json:"http"`
	InAppFeed UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      UserGetPreferencesResponseWorkflowsObjectChannelTypesPushUnion      `json:"push"`
	SMS       UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion       `json:"sms"`
	JSON      userGetPreferencesResponseWorkflowsObjectChannelTypesJSON           `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesJSON contains the JSON
// metadata for the struct [UserGetPreferencesResponseWorkflowsObjectChannelTypes]
type userGetPreferencesResponseWorkflowsObjectChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditions].
type UserGetPreferencesResponseWorkflowsObjectChannelTypesChatUnion interface {
	ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseWorkflowsObjectChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditions{}),
		},
	)
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditions struct {
	Conditions []UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON contains
// the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditions]
type userGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditions) ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition struct {
	Argument string                                                                                `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                `json:"variable,required"`
	JSON     userGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition]
type userGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains             UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions].
type UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion interface {
	ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions{}),
		},
	)
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions struct {
	Conditions []UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions]
type userGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions) ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition]
type userGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains             UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions].
type UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion interface {
	ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions{}),
		},
	)
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions struct {
	Conditions []UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON contains
// the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions]
type userGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions) ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition struct {
	Argument string                                                                                `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                `json:"variable,required"`
	JSON     userGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition]
type userGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions].
type UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions{}),
		},
	)
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions struct {
	Conditions []UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions]
type userGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions) ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument string                                                                                     `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                     `json:"variable,required"`
	JSON     userGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition]
type userGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditions].
type UserGetPreferencesResponseWorkflowsObjectChannelTypesPushUnion interface {
	ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseWorkflowsObjectChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditions{}),
		},
	)
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditions struct {
	Conditions []UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON contains
// the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditions]
type userGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditions) ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition struct {
	Argument string                                                                                `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                `json:"variable,required"`
	JSON     userGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition]
type userGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains             UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions].
type UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion interface {
	ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions{}),
		},
	)
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions struct {
	Conditions []UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition `json:"conditions,required"`
	JSON       userGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON        `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON contains
// the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions]
type userGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions) ImplementsUserGetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition struct {
	Argument string                                                                               `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                               `json:"variable,required"`
	JSON     userGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON      `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition]
type userGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains             UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserGetPreferencesResponseWorkflowsObjectCondition struct {
	Argument string                                                      `json:"argument,required,nullable"`
	Operator UserGetPreferencesResponseWorkflowsObjectConditionsOperator `json:"operator,required"`
	Variable string                                                      `json:"variable,required"`
	JSON     userGetPreferencesResponseWorkflowsObjectConditionJSON      `json:"-"`
}

// userGetPreferencesResponseWorkflowsObjectConditionJSON contains the JSON
// metadata for the struct [UserGetPreferencesResponseWorkflowsObjectCondition]
type userGetPreferencesResponseWorkflowsObjectConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsObjectCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsObjectConditionJSON) RawJSON() string {
	return r.raw
}

type UserGetPreferencesResponseWorkflowsObjectConditionsOperator string

const (
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorEqualTo              UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "equal_to"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorNotEqualTo           UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "not_equal_to"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThan          UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "greater_than"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorLessThan             UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "less_than"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "greater_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorLessThanOrEqualTo    UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "less_than_or_equal_to"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorContains             UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "contains"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorNotContains          UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "not_contains"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorEmpty                UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "empty"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorNotEmpty             UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "not_empty"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorContainsAll          UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "contains_all"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestamp          UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsNotTimestamp       UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "is_not_timestamp"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampAfter     UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp_after"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBefore    UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp_before"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBetween   UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp_between"
	UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsAudienceMember     UserGetPreferencesResponseWorkflowsObjectConditionsOperator = "is_audience_member"
)

func (r UserGetPreferencesResponseWorkflowsObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserGetPreferencesResponseWorkflowsObjectConditionsOperatorEqualTo, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorNotEqualTo, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThan, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorLessThan, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorLessThanOrEqualTo, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorContains, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorNotContains, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorEmpty, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorNotEmpty, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorContainsAll, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestamp, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsNotTimestamp, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampAfter, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBefore, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBetween, UserGetPreferencesResponseWorkflowsObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A paginated list of messages.
type UserListMessagesResponse struct {
	// The list of messages
	Entries []UserListMessagesResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo UserListMessagesResponsePageInfo `json:"page_info,required"`
	JSON     userListMessagesResponseJSON     `json:"-"`
}

// userListMessagesResponseJSON contains the JSON metadata for the struct
// [UserListMessagesResponse]
type userListMessagesResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListMessagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single message that was generated by a workflow for a given
// channel.
type UserListMessagesResponseEntry struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []UserListMessagesResponseEntriesActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []UserListMessagesResponseEntriesEngagementStatus `json:"engagement_statuses"`
	// Timestamp of creation
	InsertedAt time.Time `json:"inserted_at" format:"date-time"`
	// Timestamp when message was interacted with
	InteractedAt time.Time `json:"interacted_at,nullable" format:"date-time"`
	// Timestamp when a link in the message was clicked
	LinkClickedAt time.Time `json:"link_clicked_at,nullable" format:"date-time"`
	// Message metadata
	Metadata interface{} `json:"metadata,nullable"`
	// Timestamp when message was read
	ReadAt time.Time `json:"read_at,nullable" format:"date-time"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (id, collection).
	Recipient UserListMessagesResponseEntriesRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source UserListMessagesResponseEntriesSource `json:"source"`
	// Message delivery status
	Status UserListMessagesResponseEntriesStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                            `json:"workflow,nullable"`
	JSON     userListMessagesResponseEntryJSON `json:"-"`
}

// userListMessagesResponseEntryJSON contains the JSON metadata for the struct
// [UserListMessagesResponseEntry]
type userListMessagesResponseEntryJSON struct {
	ID                 apijson.Field
	Typename           apijson.Field
	Actors             apijson.Field
	ArchivedAt         apijson.Field
	ChannelID          apijson.Field
	ClickedAt          apijson.Field
	Data               apijson.Field
	EngagementStatuses apijson.Field
	InsertedAt         apijson.Field
	InteractedAt       apijson.Field
	LinkClickedAt      apijson.Field
	Metadata           apijson.Field
	ReadAt             apijson.Field
	Recipient          apijson.Field
	ScheduledAt        apijson.Field
	SeenAt             apijson.Field
	Source             apijson.Field
	Status             apijson.Field
	Tenant             apijson.Field
	UpdatedAt          apijson.Field
	Workflow           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserListMessagesResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [UserListMessagesResponseEntriesActorsObject].
type UserListMessagesResponseEntriesActorsUnion interface {
	ImplementsUserListMessagesResponseEntriesActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListMessagesResponseEntriesActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListMessagesResponseEntriesActorsObject{}),
		},
	)
}

// An object reference to a recipient
type UserListMessagesResponseEntriesActorsObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                          `json:"collection,required"`
	JSON       userListMessagesResponseEntriesActorsObjectJSON `json:"-"`
}

// userListMessagesResponseEntriesActorsObjectJSON contains the JSON metadata for
// the struct [UserListMessagesResponseEntriesActorsObject]
type userListMessagesResponseEntriesActorsObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListMessagesResponseEntriesActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponseEntriesActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserListMessagesResponseEntriesActorsObject) ImplementsUserListMessagesResponseEntriesActorsUnion() {
}

type UserListMessagesResponseEntriesEngagementStatus string

const (
	UserListMessagesResponseEntriesEngagementStatusSeen        UserListMessagesResponseEntriesEngagementStatus = "seen"
	UserListMessagesResponseEntriesEngagementStatusRead        UserListMessagesResponseEntriesEngagementStatus = "read"
	UserListMessagesResponseEntriesEngagementStatusInteracted  UserListMessagesResponseEntriesEngagementStatus = "interacted"
	UserListMessagesResponseEntriesEngagementStatusLinkClicked UserListMessagesResponseEntriesEngagementStatus = "link_clicked"
	UserListMessagesResponseEntriesEngagementStatusArchived    UserListMessagesResponseEntriesEngagementStatus = "archived"
)

func (r UserListMessagesResponseEntriesEngagementStatus) IsKnown() bool {
	switch r {
	case UserListMessagesResponseEntriesEngagementStatusSeen, UserListMessagesResponseEntriesEngagementStatusRead, UserListMessagesResponseEntriesEngagementStatusInteracted, UserListMessagesResponseEntriesEngagementStatusLinkClicked, UserListMessagesResponseEntriesEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [UserListMessagesResponseEntriesRecipientObject].
type UserListMessagesResponseEntriesRecipientUnion interface {
	ImplementsUserListMessagesResponseEntriesRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListMessagesResponseEntriesRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListMessagesResponseEntriesRecipientObject{}),
		},
	)
}

// An object reference to a recipient
type UserListMessagesResponseEntriesRecipientObject struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                             `json:"collection,required"`
	JSON       userListMessagesResponseEntriesRecipientObjectJSON `json:"-"`
}

// userListMessagesResponseEntriesRecipientObjectJSON contains the JSON metadata
// for the struct [UserListMessagesResponseEntriesRecipientObject]
type userListMessagesResponseEntriesRecipientObjectJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListMessagesResponseEntriesRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponseEntriesRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserListMessagesResponseEntriesRecipientObject) ImplementsUserListMessagesResponseEntriesRecipientUnion() {
}

// Source information
type UserListMessagesResponseEntriesSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                                    `json:"version_id,required" format:"uuid"`
	JSON      userListMessagesResponseEntriesSourceJSON `json:"-"`
}

// userListMessagesResponseEntriesSourceJSON contains the JSON metadata for the
// struct [UserListMessagesResponseEntriesSource]
type userListMessagesResponseEntriesSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListMessagesResponseEntriesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponseEntriesSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type UserListMessagesResponseEntriesStatus string

const (
	UserListMessagesResponseEntriesStatusQueued            UserListMessagesResponseEntriesStatus = "queued"
	UserListMessagesResponseEntriesStatusSent              UserListMessagesResponseEntriesStatus = "sent"
	UserListMessagesResponseEntriesStatusDelivered         UserListMessagesResponseEntriesStatus = "delivered"
	UserListMessagesResponseEntriesStatusDeliveryAttempted UserListMessagesResponseEntriesStatus = "delivery_attempted"
	UserListMessagesResponseEntriesStatusUndelivered       UserListMessagesResponseEntriesStatus = "undelivered"
	UserListMessagesResponseEntriesStatusNotSent           UserListMessagesResponseEntriesStatus = "not_sent"
	UserListMessagesResponseEntriesStatusBounced           UserListMessagesResponseEntriesStatus = "bounced"
)

func (r UserListMessagesResponseEntriesStatus) IsKnown() bool {
	switch r {
	case UserListMessagesResponseEntriesStatusQueued, UserListMessagesResponseEntriesStatusSent, UserListMessagesResponseEntriesStatusDelivered, UserListMessagesResponseEntriesStatusDeliveryAttempted, UserListMessagesResponseEntriesStatusUndelivered, UserListMessagesResponseEntriesStatusNotSent, UserListMessagesResponseEntriesStatusBounced:
		return true
	}
	return false
}

// The information about a paginated result
type UserListMessagesResponsePageInfo struct {
	Typename string                               `json:"__typename,required"`
	PageSize int64                                `json:"page_size,required"`
	After    string                               `json:"after,nullable"`
	Before   string                               `json:"before,nullable"`
	JSON     userListMessagesResponsePageInfoJSON `json:"-"`
}

// userListMessagesResponsePageInfoJSON contains the JSON metadata for the struct
// [UserListMessagesResponsePageInfo]
type userListMessagesResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListMessagesResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

// A preference set object.
type UserListPreferencesResponse struct {
	ID         string                                                `json:"id,required"`
	Typename   string                                                `json:"__typename,required"`
	Categories map[string]UserListPreferencesResponseCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes UserListPreferencesResponseChannelTypes              `json:"channel_types,nullable"`
	Workflows    map[string]UserListPreferencesResponseWorkflowsUnion `json:"workflows,nullable"`
	JSON         userListPreferencesResponseJSON                      `json:"-"`
}

// userListPreferencesResponseJSON contains the JSON metadata for the struct
// [UserListPreferencesResponse]
type userListPreferencesResponseJSON struct {
	ID           apijson.Field
	Typename     apijson.Field
	Categories   apijson.Field
	ChannelTypes apijson.Field
	Workflows    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserListPreferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseJSON) RawJSON() string {
	return r.raw
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseCategoriesObject].
type UserListPreferencesResponseCategoriesUnion interface {
	ImplementsUserListPreferencesResponseCategoriesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseCategoriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseCategoriesObject{}),
		},
	)
}

type UserListPreferencesResponseCategoriesObject struct {
	// Channel type preferences
	ChannelTypes UserListPreferencesResponseCategoriesObjectChannelTypes `json:"channel_types,nullable"`
	Conditions   []UserListPreferencesResponseCategoriesObjectCondition  `json:"conditions"`
	JSON         userListPreferencesResponseCategoriesObjectJSON         `json:"-"`
}

// userListPreferencesResponseCategoriesObjectJSON contains the JSON metadata for
// the struct [UserListPreferencesResponseCategoriesObject]
type userListPreferencesResponseCategoriesObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseCategoriesObject) ImplementsUserListPreferencesResponseCategoriesUnion() {
}

// Channel type preferences
type UserListPreferencesResponseCategoriesObjectChannelTypes struct {
	Chat      UserListPreferencesResponseCategoriesObjectChannelTypesChatUnion      `json:"chat"`
	Email     UserListPreferencesResponseCategoriesObjectChannelTypesEmailUnion     `json:"email"`
	HTTP      UserListPreferencesResponseCategoriesObjectChannelTypesHTTPUnion      `json:"http"`
	InAppFeed UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      UserListPreferencesResponseCategoriesObjectChannelTypesPushUnion      `json:"push"`
	SMS       UserListPreferencesResponseCategoriesObjectChannelTypesSMSUnion       `json:"sms"`
	JSON      userListPreferencesResponseCategoriesObjectChannelTypesJSON           `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesJSON contains the JSON
// metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypes]
type userListPreferencesResponseCategoriesObjectChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseCategoriesObjectChannelTypesChatConditions].
type UserListPreferencesResponseCategoriesObjectChannelTypesChatUnion interface {
	ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseCategoriesObjectChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseCategoriesObjectChannelTypesChatConditions{}),
		},
	)
}

type UserListPreferencesResponseCategoriesObjectChannelTypesChatConditions struct {
	Conditions []UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON        `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesChatConditions]
type userListPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseCategoriesObjectChannelTypesChatConditions) ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition struct {
	Argument string                                                                                  `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                  `json:"variable,required"`
	JSON     userListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition]
type userListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains             UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditions].
type UserListPreferencesResponseCategoriesObjectChannelTypesEmailUnion interface {
	ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseCategoriesObjectChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditions{}),
		},
	)
}

type UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditions struct {
	Conditions []UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON        `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditions]
type userListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditions) ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition struct {
	Argument string                                                                                   `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                   `json:"variable,required"`
	JSON     userListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition]
type userListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains             UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditions].
type UserListPreferencesResponseCategoriesObjectChannelTypesHTTPUnion interface {
	ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseCategoriesObjectChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditions{}),
		},
	)
}

type UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditions struct {
	Conditions []UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON        `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditions]
type userListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditions) ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition struct {
	Argument string                                                                                  `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                  `json:"variable,required"`
	JSON     userListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition]
type userListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions].
type UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions{}),
		},
	)
}

type UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions struct {
	Conditions []UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON        `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions]
type userListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions) ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument string                                                                                       `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                       `json:"variable,required"`
	JSON     userListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition]
type userListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseCategoriesObjectChannelTypesPushConditions].
type UserListPreferencesResponseCategoriesObjectChannelTypesPushUnion interface {
	ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseCategoriesObjectChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseCategoriesObjectChannelTypesPushConditions{}),
		},
	)
}

type UserListPreferencesResponseCategoriesObjectChannelTypesPushConditions struct {
	Conditions []UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON        `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesPushConditions]
type userListPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseCategoriesObjectChannelTypesPushConditions) ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition struct {
	Argument string                                                                                  `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                  `json:"variable,required"`
	JSON     userListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition]
type userListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains             UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditions].
type UserListPreferencesResponseCategoriesObjectChannelTypesSMSUnion interface {
	ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseCategoriesObjectChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditions{}),
		},
	)
}

type UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditions struct {
	Conditions []UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON        `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditions]
type userListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditions) ImplementsUserListPreferencesResponseCategoriesObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition]
type userListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains             UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserListPreferencesResponseCategoriesObjectCondition struct {
	Argument string                                                        `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseCategoriesObjectConditionsOperator `json:"operator,required"`
	Variable string                                                        `json:"variable,required"`
	JSON     userListPreferencesResponseCategoriesObjectConditionJSON      `json:"-"`
}

// userListPreferencesResponseCategoriesObjectConditionJSON contains the JSON
// metadata for the struct [UserListPreferencesResponseCategoriesObjectCondition]
type userListPreferencesResponseCategoriesObjectConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesObjectCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesObjectConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseCategoriesObjectConditionsOperator string

const (
	UserListPreferencesResponseCategoriesObjectConditionsOperatorEqualTo              UserListPreferencesResponseCategoriesObjectConditionsOperator = "equal_to"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorNotEqualTo           UserListPreferencesResponseCategoriesObjectConditionsOperator = "not_equal_to"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorGreaterThan          UserListPreferencesResponseCategoriesObjectConditionsOperator = "greater_than"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorLessThan             UserListPreferencesResponseCategoriesObjectConditionsOperator = "less_than"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseCategoriesObjectConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseCategoriesObjectConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorContains             UserListPreferencesResponseCategoriesObjectConditionsOperator = "contains"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorNotContains          UserListPreferencesResponseCategoriesObjectConditionsOperator = "not_contains"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorEmpty                UserListPreferencesResponseCategoriesObjectConditionsOperator = "empty"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorNotEmpty             UserListPreferencesResponseCategoriesObjectConditionsOperator = "not_empty"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorContainsAll          UserListPreferencesResponseCategoriesObjectConditionsOperator = "contains_all"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorIsTimestamp          UserListPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorIsNotTimestamp       UserListPreferencesResponseCategoriesObjectConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampAfter     UserListPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBefore    UserListPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBetween   UserListPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseCategoriesObjectConditionsOperatorIsAudienceMember     UserListPreferencesResponseCategoriesObjectConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseCategoriesObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseCategoriesObjectConditionsOperatorEqualTo, UserListPreferencesResponseCategoriesObjectConditionsOperatorNotEqualTo, UserListPreferencesResponseCategoriesObjectConditionsOperatorGreaterThan, UserListPreferencesResponseCategoriesObjectConditionsOperatorLessThan, UserListPreferencesResponseCategoriesObjectConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseCategoriesObjectConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseCategoriesObjectConditionsOperatorContains, UserListPreferencesResponseCategoriesObjectConditionsOperatorNotContains, UserListPreferencesResponseCategoriesObjectConditionsOperatorEmpty, UserListPreferencesResponseCategoriesObjectConditionsOperatorNotEmpty, UserListPreferencesResponseCategoriesObjectConditionsOperatorContainsAll, UserListPreferencesResponseCategoriesObjectConditionsOperatorIsTimestamp, UserListPreferencesResponseCategoriesObjectConditionsOperatorIsNotTimestamp, UserListPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampAfter, UserListPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBefore, UserListPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBetween, UserListPreferencesResponseCategoriesObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type UserListPreferencesResponseChannelTypes struct {
	Chat      UserListPreferencesResponseChannelTypesChatUnion      `json:"chat"`
	Email     UserListPreferencesResponseChannelTypesEmailUnion     `json:"email"`
	HTTP      UserListPreferencesResponseChannelTypesHTTPUnion      `json:"http"`
	InAppFeed UserListPreferencesResponseChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      UserListPreferencesResponseChannelTypesPushUnion      `json:"push"`
	SMS       UserListPreferencesResponseChannelTypesSMSUnion       `json:"sms"`
	JSON      userListPreferencesResponseChannelTypesJSON           `json:"-"`
}

// userListPreferencesResponseChannelTypesJSON contains the JSON metadata for the
// struct [UserListPreferencesResponseChannelTypes]
type userListPreferencesResponseChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseChannelTypesChatConditions].
type UserListPreferencesResponseChannelTypesChatUnion interface {
	ImplementsUserListPreferencesResponseChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseChannelTypesChatConditions{}),
		},
	)
}

type UserListPreferencesResponseChannelTypesChatConditions struct {
	Conditions []UserListPreferencesResponseChannelTypesChatConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseChannelTypesChatConditionsJSON        `json:"-"`
}

// userListPreferencesResponseChannelTypesChatConditionsJSON contains the JSON
// metadata for the struct [UserListPreferencesResponseChannelTypesChatConditions]
type userListPreferencesResponseChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseChannelTypesChatConditions) ImplementsUserListPreferencesResponseChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseChannelTypesChatConditionsCondition struct {
	Argument string                                                                  `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                  `json:"variable,required"`
	JSON     userListPreferencesResponseChannelTypesChatConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseChannelTypesChatConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserListPreferencesResponseChannelTypesChatConditionsCondition]
type userListPreferencesResponseChannelTypesChatConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesChatConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesChatConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator string

const (
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorEqualTo              UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThan             UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorContains             UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "contains"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotContains          UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorEmpty                UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "empty"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorContainsAll          UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorEqualTo, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThan, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorContains, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotContains, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorEmpty, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorContainsAll, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseChannelTypesEmailConditions].
type UserListPreferencesResponseChannelTypesEmailUnion interface {
	ImplementsUserListPreferencesResponseChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseChannelTypesEmailConditions{}),
		},
	)
}

type UserListPreferencesResponseChannelTypesEmailConditions struct {
	Conditions []UserListPreferencesResponseChannelTypesEmailConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseChannelTypesEmailConditionsJSON        `json:"-"`
}

// userListPreferencesResponseChannelTypesEmailConditionsJSON contains the JSON
// metadata for the struct [UserListPreferencesResponseChannelTypesEmailConditions]
type userListPreferencesResponseChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseChannelTypesEmailConditions) ImplementsUserListPreferencesResponseChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseChannelTypesEmailConditionsCondition struct {
	Argument string                                                                   `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                   `json:"variable,required"`
	JSON     userListPreferencesResponseChannelTypesEmailConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseChannelTypesEmailConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserListPreferencesResponseChannelTypesEmailConditionsCondition]
type userListPreferencesResponseChannelTypesEmailConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesEmailConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesEmailConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator string

const (
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEqualTo              UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThan             UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContains             UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "contains"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotContains          UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEmpty                UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "empty"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContainsAll          UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEqualTo, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThan, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContains, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotContains, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEmpty, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContainsAll, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseChannelTypesHTTPConditions].
type UserListPreferencesResponseChannelTypesHTTPUnion interface {
	ImplementsUserListPreferencesResponseChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseChannelTypesHTTPConditions{}),
		},
	)
}

type UserListPreferencesResponseChannelTypesHTTPConditions struct {
	Conditions []UserListPreferencesResponseChannelTypesHTTPConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseChannelTypesHTTPConditionsJSON        `json:"-"`
}

// userListPreferencesResponseChannelTypesHTTPConditionsJSON contains the JSON
// metadata for the struct [UserListPreferencesResponseChannelTypesHTTPConditions]
type userListPreferencesResponseChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseChannelTypesHTTPConditions) ImplementsUserListPreferencesResponseChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseChannelTypesHTTPConditionsCondition struct {
	Argument string                                                                  `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                  `json:"variable,required"`
	JSON     userListPreferencesResponseChannelTypesHTTPConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseChannelTypesHTTPConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserListPreferencesResponseChannelTypesHTTPConditionsCondition]
type userListPreferencesResponseChannelTypesHTTPConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesHTTPConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesHTTPConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator string

const (
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThan             UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContains             UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotContains          UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEmpty                UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThan, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContains, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotContains, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEmpty, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseChannelTypesInAppFeedConditions].
type UserListPreferencesResponseChannelTypesInAppFeedUnion interface {
	ImplementsUserListPreferencesResponseChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseChannelTypesInAppFeedConditions{}),
		},
	)
}

type UserListPreferencesResponseChannelTypesInAppFeedConditions struct {
	Conditions []UserListPreferencesResponseChannelTypesInAppFeedConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseChannelTypesInAppFeedConditionsJSON        `json:"-"`
}

// userListPreferencesResponseChannelTypesInAppFeedConditionsJSON contains the JSON
// metadata for the struct
// [UserListPreferencesResponseChannelTypesInAppFeedConditions]
type userListPreferencesResponseChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseChannelTypesInAppFeedConditions) ImplementsUserListPreferencesResponseChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseChannelTypesInAppFeedConditionsCondition struct {
	Argument string                                                                       `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                       `json:"variable,required"`
	JSON     userListPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON contains
// the JSON metadata for the struct
// [UserListPreferencesResponseChannelTypesInAppFeedConditionsCondition]
type userListPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesInAppFeedConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContains             UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContains, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseChannelTypesPushConditions].
type UserListPreferencesResponseChannelTypesPushUnion interface {
	ImplementsUserListPreferencesResponseChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseChannelTypesPushConditions{}),
		},
	)
}

type UserListPreferencesResponseChannelTypesPushConditions struct {
	Conditions []UserListPreferencesResponseChannelTypesPushConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseChannelTypesPushConditionsJSON        `json:"-"`
}

// userListPreferencesResponseChannelTypesPushConditionsJSON contains the JSON
// metadata for the struct [UserListPreferencesResponseChannelTypesPushConditions]
type userListPreferencesResponseChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseChannelTypesPushConditions) ImplementsUserListPreferencesResponseChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseChannelTypesPushConditionsCondition struct {
	Argument string                                                                  `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                  `json:"variable,required"`
	JSON     userListPreferencesResponseChannelTypesPushConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseChannelTypesPushConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserListPreferencesResponseChannelTypesPushConditionsCondition]
type userListPreferencesResponseChannelTypesPushConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesPushConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesPushConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator string

const (
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorEqualTo              UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThan             UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorContains             UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "contains"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotContains          UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorEmpty                UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "empty"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorContainsAll          UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorEqualTo, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThan, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorContains, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotContains, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorEmpty, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorContainsAll, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseChannelTypesSMSConditions].
type UserListPreferencesResponseChannelTypesSMSUnion interface {
	ImplementsUserListPreferencesResponseChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseChannelTypesSMSConditions{}),
		},
	)
}

type UserListPreferencesResponseChannelTypesSMSConditions struct {
	Conditions []UserListPreferencesResponseChannelTypesSMSConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseChannelTypesSMSConditionsJSON        `json:"-"`
}

// userListPreferencesResponseChannelTypesSMSConditionsJSON contains the JSON
// metadata for the struct [UserListPreferencesResponseChannelTypesSMSConditions]
type userListPreferencesResponseChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseChannelTypesSMSConditions) ImplementsUserListPreferencesResponseChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseChannelTypesSMSConditionsCondition struct {
	Argument string                                                                 `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                 `json:"variable,required"`
	JSON     userListPreferencesResponseChannelTypesSMSConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseChannelTypesSMSConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserListPreferencesResponseChannelTypesSMSConditionsCondition]
type userListPreferencesResponseChannelTypesSMSConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseChannelTypesSMSConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseChannelTypesSMSConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator string

const (
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEqualTo              UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThan             UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContains             UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "contains"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotContains          UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEmpty                UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "empty"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContainsAll          UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEqualTo, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThan, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContains, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotContains, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEmpty, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContainsAll, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseWorkflowsObject].
type UserListPreferencesResponseWorkflowsUnion interface {
	ImplementsUserListPreferencesResponseWorkflowsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseWorkflowsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseWorkflowsObject{}),
		},
	)
}

type UserListPreferencesResponseWorkflowsObject struct {
	// Channel type preferences
	ChannelTypes UserListPreferencesResponseWorkflowsObjectChannelTypes `json:"channel_types,nullable"`
	Conditions   []UserListPreferencesResponseWorkflowsObjectCondition  `json:"conditions"`
	JSON         userListPreferencesResponseWorkflowsObjectJSON         `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectJSON contains the JSON metadata for
// the struct [UserListPreferencesResponseWorkflowsObject]
type userListPreferencesResponseWorkflowsObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseWorkflowsObject) ImplementsUserListPreferencesResponseWorkflowsUnion() {
}

// Channel type preferences
type UserListPreferencesResponseWorkflowsObjectChannelTypes struct {
	Chat      UserListPreferencesResponseWorkflowsObjectChannelTypesChatUnion      `json:"chat"`
	Email     UserListPreferencesResponseWorkflowsObjectChannelTypesEmailUnion     `json:"email"`
	HTTP      UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion      `json:"http"`
	InAppFeed UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      UserListPreferencesResponseWorkflowsObjectChannelTypesPushUnion      `json:"push"`
	SMS       UserListPreferencesResponseWorkflowsObjectChannelTypesSMSUnion       `json:"sms"`
	JSON      userListPreferencesResponseWorkflowsObjectChannelTypesJSON           `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesJSON contains the JSON
// metadata for the struct [UserListPreferencesResponseWorkflowsObjectChannelTypes]
type userListPreferencesResponseWorkflowsObjectChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditions].
type UserListPreferencesResponseWorkflowsObjectChannelTypesChatUnion interface {
	ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseWorkflowsObjectChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditions{}),
		},
	)
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditions struct {
	Conditions []UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON        `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditions]
type userListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditions) ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition]
type userListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains             UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditions].
type UserListPreferencesResponseWorkflowsObjectChannelTypesEmailUnion interface {
	ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseWorkflowsObjectChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditions{}),
		},
	)
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditions struct {
	Conditions []UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON        `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditions]
type userListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditions) ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition struct {
	Argument string                                                                                  `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                  `json:"variable,required"`
	JSON     userListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition]
type userListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains             UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions].
type UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion interface {
	ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions{}),
		},
	)
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions struct {
	Conditions []UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON        `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions]
type userListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions) ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition]
type userListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions].
type UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions{}),
		},
	)
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions struct {
	Conditions []UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON        `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions]
type userListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions) ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument string                                                                                      `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                      `json:"variable,required"`
	JSON     userListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition]
type userListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditions].
type UserListPreferencesResponseWorkflowsObjectChannelTypesPushUnion interface {
	ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseWorkflowsObjectChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditions{}),
		},
	)
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditions struct {
	Conditions []UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON        `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditions]
type userListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditions) ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition]
type userListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains             UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditions].
type UserListPreferencesResponseWorkflowsObjectChannelTypesSMSUnion interface {
	ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListPreferencesResponseWorkflowsObjectChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditions{}),
		},
	)
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditions struct {
	Conditions []UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition `json:"conditions,required"`
	JSON       userListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON        `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON contains
// the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditions]
type userListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditions) ImplementsUserListPreferencesResponseWorkflowsObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition struct {
	Argument string                                                                                `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                `json:"variable,required"`
	JSON     userListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON      `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition]
type userListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains             UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserListPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserListPreferencesResponseWorkflowsObjectCondition struct {
	Argument string                                                       `json:"argument,required,nullable"`
	Operator UserListPreferencesResponseWorkflowsObjectConditionsOperator `json:"operator,required"`
	Variable string                                                       `json:"variable,required"`
	JSON     userListPreferencesResponseWorkflowsObjectConditionJSON      `json:"-"`
}

// userListPreferencesResponseWorkflowsObjectConditionJSON contains the JSON
// metadata for the struct [UserListPreferencesResponseWorkflowsObjectCondition]
type userListPreferencesResponseWorkflowsObjectConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsObjectCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsObjectConditionJSON) RawJSON() string {
	return r.raw
}

type UserListPreferencesResponseWorkflowsObjectConditionsOperator string

const (
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorEqualTo              UserListPreferencesResponseWorkflowsObjectConditionsOperator = "equal_to"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorNotEqualTo           UserListPreferencesResponseWorkflowsObjectConditionsOperator = "not_equal_to"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThan          UserListPreferencesResponseWorkflowsObjectConditionsOperator = "greater_than"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorLessThan             UserListPreferencesResponseWorkflowsObjectConditionsOperator = "less_than"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo UserListPreferencesResponseWorkflowsObjectConditionsOperator = "greater_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorLessThanOrEqualTo    UserListPreferencesResponseWorkflowsObjectConditionsOperator = "less_than_or_equal_to"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorContains             UserListPreferencesResponseWorkflowsObjectConditionsOperator = "contains"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorNotContains          UserListPreferencesResponseWorkflowsObjectConditionsOperator = "not_contains"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorEmpty                UserListPreferencesResponseWorkflowsObjectConditionsOperator = "empty"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorNotEmpty             UserListPreferencesResponseWorkflowsObjectConditionsOperator = "not_empty"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorContainsAll          UserListPreferencesResponseWorkflowsObjectConditionsOperator = "contains_all"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestamp          UserListPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsNotTimestamp       UserListPreferencesResponseWorkflowsObjectConditionsOperator = "is_not_timestamp"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampAfter     UserListPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp_after"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBefore    UserListPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp_before"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBetween   UserListPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp_between"
	UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsAudienceMember     UserListPreferencesResponseWorkflowsObjectConditionsOperator = "is_audience_member"
)

func (r UserListPreferencesResponseWorkflowsObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserListPreferencesResponseWorkflowsObjectConditionsOperatorEqualTo, UserListPreferencesResponseWorkflowsObjectConditionsOperatorNotEqualTo, UserListPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThan, UserListPreferencesResponseWorkflowsObjectConditionsOperatorLessThan, UserListPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectConditionsOperatorLessThanOrEqualTo, UserListPreferencesResponseWorkflowsObjectConditionsOperatorContains, UserListPreferencesResponseWorkflowsObjectConditionsOperatorNotContains, UserListPreferencesResponseWorkflowsObjectConditionsOperatorEmpty, UserListPreferencesResponseWorkflowsObjectConditionsOperatorNotEmpty, UserListPreferencesResponseWorkflowsObjectConditionsOperatorContainsAll, UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestamp, UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsNotTimestamp, UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampAfter, UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBefore, UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBetween, UserListPreferencesResponseWorkflowsObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A paginated list of schedules in a collection.
type UserListSchedulesResponse struct {
	// The list of schedules
	Entries []UserListSchedulesResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo UserListSchedulesResponsePageInfo `json:"page_info,required"`
	JSON     userListSchedulesResponseJSON     `json:"-"`
}

// userListSchedulesResponseJSON contains the JSON metadata for the struct
// [UserListSchedulesResponse]
type userListSchedulesResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSchedulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSchedulesResponseJSON) RawJSON() string {
	return r.raw
}

// A schedule that represents a recurring workflow execution
type UserListSchedulesResponseEntry struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient UserListSchedulesResponseEntriesRecipient `json:"recipient,required"`
	Repeats   []UserListSchedulesResponseEntriesRepeat  `json:"repeats,required"`
	UpdatedAt time.Time                                 `json:"updated_at,required" format:"date-time"`
	Workflow  string                                    `json:"workflow,required"`
	Typename  string                                    `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            UserListSchedulesResponseEntriesActor `json:"actor,nullable"`
	Data             interface{}                           `json:"data,nullable"`
	LastOccurrenceAt time.Time                             `json:"last_occurrence_at,nullable" format:"date-time"`
	NextOccurrenceAt time.Time                             `json:"next_occurrence_at,nullable" format:"date-time"`
	Tenant           string                                `json:"tenant,nullable"`
	JSON             userListSchedulesResponseEntryJSON    `json:"-"`
}

// userListSchedulesResponseEntryJSON contains the JSON metadata for the struct
// [UserListSchedulesResponseEntry]
type userListSchedulesResponseEntryJSON struct {
	ID               apijson.Field
	InsertedAt       apijson.Field
	Recipient        apijson.Field
	Repeats          apijson.Field
	UpdatedAt        apijson.Field
	Workflow         apijson.Field
	Typename         apijson.Field
	Actor            apijson.Field
	Data             apijson.Field
	LastOccurrenceAt apijson.Field
	NextOccurrenceAt apijson.Field
	Tenant           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserListSchedulesResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSchedulesResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type UserListSchedulesResponseEntriesRecipient struct {
	ID          string                                        `json:"id,required"`
	Typename    string                                        `json:"__typename,required"`
	UpdatedAt   time.Time                                     `json:"updated_at,required" format:"date-time"`
	Avatar      string                                        `json:"avatar,nullable"`
	Collection  string                                        `json:"collection"`
	CreatedAt   time.Time                                     `json:"created_at,nullable" format:"date-time"`
	Email       string                                        `json:"email,nullable" format:"email"`
	Name        string                                        `json:"name,nullable"`
	PhoneNumber string                                        `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                                        `json:"timezone,nullable"`
	JSON        userListSchedulesResponseEntriesRecipientJSON `json:"-"`
	union       UserListSchedulesResponseEntriesRecipientUnion
}

// userListSchedulesResponseEntriesRecipientJSON contains the JSON metadata for the
// struct [UserListSchedulesResponseEntriesRecipient]
type userListSchedulesResponseEntriesRecipientJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userListSchedulesResponseEntriesRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *UserListSchedulesResponseEntriesRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = UserListSchedulesResponseEntriesRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserListSchedulesResponseEntriesRecipientUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [UserListSchedulesResponseEntriesRecipientObject].
func (r UserListSchedulesResponseEntriesRecipient) AsUnion() UserListSchedulesResponseEntriesRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [UserListSchedulesResponseEntriesRecipientObject].
type UserListSchedulesResponseEntriesRecipientUnion interface {
	implementsUserListSchedulesResponseEntriesRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListSchedulesResponseEntriesRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListSchedulesResponseEntriesRecipientObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type UserListSchedulesResponseEntriesRecipientObject struct {
	ID          string                                              `json:"id,required"`
	Typename    string                                              `json:"__typename,required"`
	Collection  string                                              `json:"collection,required"`
	UpdatedAt   time.Time                                           `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                           `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                              `json:"-,extras"`
	JSON        userListSchedulesResponseEntriesRecipientObjectJSON `json:"-"`
}

// userListSchedulesResponseEntriesRecipientObjectJSON contains the JSON metadata
// for the struct [UserListSchedulesResponseEntriesRecipientObject]
type userListSchedulesResponseEntriesRecipientObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSchedulesResponseEntriesRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSchedulesResponseEntriesRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserListSchedulesResponseEntriesRecipientObject) implementsUserListSchedulesResponseEntriesRecipient() {
}

// A schedule repeat rule
type UserListSchedulesResponseEntriesRepeat struct {
	Typename   string                                           `json:"__typename,required"`
	Frequency  UserListSchedulesResponseEntriesRepeatsFrequency `json:"frequency,required"`
	DayOfMonth int64                                            `json:"day_of_month,nullable"`
	Days       []UserListSchedulesResponseEntriesRepeatsDay     `json:"days,nullable"`
	Hours      int64                                            `json:"hours,nullable"`
	Interval   int64                                            `json:"interval"`
	Minutes    int64                                            `json:"minutes,nullable"`
	JSON       userListSchedulesResponseEntriesRepeatJSON       `json:"-"`
}

// userListSchedulesResponseEntriesRepeatJSON contains the JSON metadata for the
// struct [UserListSchedulesResponseEntriesRepeat]
type userListSchedulesResponseEntriesRepeatJSON struct {
	Typename    apijson.Field
	Frequency   apijson.Field
	DayOfMonth  apijson.Field
	Days        apijson.Field
	Hours       apijson.Field
	Interval    apijson.Field
	Minutes     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSchedulesResponseEntriesRepeat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSchedulesResponseEntriesRepeatJSON) RawJSON() string {
	return r.raw
}

type UserListSchedulesResponseEntriesRepeatsFrequency string

const (
	UserListSchedulesResponseEntriesRepeatsFrequencyDaily   UserListSchedulesResponseEntriesRepeatsFrequency = "daily"
	UserListSchedulesResponseEntriesRepeatsFrequencyWeekly  UserListSchedulesResponseEntriesRepeatsFrequency = "weekly"
	UserListSchedulesResponseEntriesRepeatsFrequencyMonthly UserListSchedulesResponseEntriesRepeatsFrequency = "monthly"
	UserListSchedulesResponseEntriesRepeatsFrequencyHourly  UserListSchedulesResponseEntriesRepeatsFrequency = "hourly"
)

func (r UserListSchedulesResponseEntriesRepeatsFrequency) IsKnown() bool {
	switch r {
	case UserListSchedulesResponseEntriesRepeatsFrequencyDaily, UserListSchedulesResponseEntriesRepeatsFrequencyWeekly, UserListSchedulesResponseEntriesRepeatsFrequencyMonthly, UserListSchedulesResponseEntriesRepeatsFrequencyHourly:
		return true
	}
	return false
}

type UserListSchedulesResponseEntriesRepeatsDay string

const (
	UserListSchedulesResponseEntriesRepeatsDayMon UserListSchedulesResponseEntriesRepeatsDay = "mon"
	UserListSchedulesResponseEntriesRepeatsDayTue UserListSchedulesResponseEntriesRepeatsDay = "tue"
	UserListSchedulesResponseEntriesRepeatsDayWed UserListSchedulesResponseEntriesRepeatsDay = "wed"
	UserListSchedulesResponseEntriesRepeatsDayThu UserListSchedulesResponseEntriesRepeatsDay = "thu"
	UserListSchedulesResponseEntriesRepeatsDayFri UserListSchedulesResponseEntriesRepeatsDay = "fri"
	UserListSchedulesResponseEntriesRepeatsDaySat UserListSchedulesResponseEntriesRepeatsDay = "sat"
	UserListSchedulesResponseEntriesRepeatsDaySun UserListSchedulesResponseEntriesRepeatsDay = "sun"
)

func (r UserListSchedulesResponseEntriesRepeatsDay) IsKnown() bool {
	switch r {
	case UserListSchedulesResponseEntriesRepeatsDayMon, UserListSchedulesResponseEntriesRepeatsDayTue, UserListSchedulesResponseEntriesRepeatsDayWed, UserListSchedulesResponseEntriesRepeatsDayThu, UserListSchedulesResponseEntriesRepeatsDayFri, UserListSchedulesResponseEntriesRepeatsDaySat, UserListSchedulesResponseEntriesRepeatsDaySun:
		return true
	}
	return false
}

// A recipient, which is either a user or an object
type UserListSchedulesResponseEntriesActor struct {
	ID          string                                    `json:"id,required"`
	Typename    string                                    `json:"__typename,required"`
	UpdatedAt   time.Time                                 `json:"updated_at,required" format:"date-time"`
	Avatar      string                                    `json:"avatar,nullable"`
	Collection  string                                    `json:"collection"`
	CreatedAt   time.Time                                 `json:"created_at,nullable" format:"date-time"`
	Email       string                                    `json:"email,nullable" format:"email"`
	Name        string                                    `json:"name,nullable"`
	PhoneNumber string                                    `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                                    `json:"timezone,nullable"`
	JSON        userListSchedulesResponseEntriesActorJSON `json:"-"`
	union       UserListSchedulesResponseEntriesActorUnion
}

// userListSchedulesResponseEntriesActorJSON contains the JSON metadata for the
// struct [UserListSchedulesResponseEntriesActor]
type userListSchedulesResponseEntriesActorJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userListSchedulesResponseEntriesActorJSON) RawJSON() string {
	return r.raw
}

func (r *UserListSchedulesResponseEntriesActor) UnmarshalJSON(data []byte) (err error) {
	*r = UserListSchedulesResponseEntriesActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserListSchedulesResponseEntriesActorUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [UserListSchedulesResponseEntriesActorObject].
func (r UserListSchedulesResponseEntriesActor) AsUnion() UserListSchedulesResponseEntriesActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [UserListSchedulesResponseEntriesActorObject].
type UserListSchedulesResponseEntriesActorUnion interface {
	implementsUserListSchedulesResponseEntriesActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListSchedulesResponseEntriesActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListSchedulesResponseEntriesActorObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type UserListSchedulesResponseEntriesActorObject struct {
	ID          string                                          `json:"id,required"`
	Typename    string                                          `json:"__typename,required"`
	Collection  string                                          `json:"collection,required"`
	UpdatedAt   time.Time                                       `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                       `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                          `json:"-,extras"`
	JSON        userListSchedulesResponseEntriesActorObjectJSON `json:"-"`
}

// userListSchedulesResponseEntriesActorObjectJSON contains the JSON metadata for
// the struct [UserListSchedulesResponseEntriesActorObject]
type userListSchedulesResponseEntriesActorObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSchedulesResponseEntriesActorObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSchedulesResponseEntriesActorObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserListSchedulesResponseEntriesActorObject) implementsUserListSchedulesResponseEntriesActor() {
}

// The information about a paginated result
type UserListSchedulesResponsePageInfo struct {
	Typename string                                `json:"__typename,required"`
	PageSize int64                                 `json:"page_size,required"`
	After    string                                `json:"after,nullable"`
	Before   string                                `json:"before,nullable"`
	JSON     userListSchedulesResponsePageInfoJSON `json:"-"`
}

// userListSchedulesResponsePageInfoJSON contains the JSON metadata for the struct
// [UserListSchedulesResponsePageInfo]
type userListSchedulesResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSchedulesResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSchedulesResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

// A paginated list of subscriptions for an object.
type UserListSubscriptionsResponse struct {
	// The list of subscriptions
	Entries []UserListSubscriptionsResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo UserListSubscriptionsResponsePageInfo `json:"page_info,required"`
	JSON     userListSubscriptionsResponseJSON     `json:"-"`
}

// userListSubscriptionsResponseJSON contains the JSON metadata for the struct
// [UserListSubscriptionsResponse]
type userListSubscriptionsResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSubscriptionsResponseJSON) RawJSON() string {
	return r.raw
}

// A subscription object
type UserListSubscriptionsResponseEntry struct {
	Typename   string    `json:"__typename,required"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A custom-object entity which belongs to a collection.
	Object UserListSubscriptionsResponseEntriesObject `json:"object,required"`
	// A recipient, which is either a user or an object
	Recipient UserListSubscriptionsResponseEntriesRecipient `json:"recipient,required"`
	UpdatedAt time.Time                                     `json:"updated_at,required" format:"date-time"`
	// The custom properties associated with the subscription
	Properties map[string]interface{}                 `json:"properties,nullable"`
	JSON       userListSubscriptionsResponseEntryJSON `json:"-"`
}

// userListSubscriptionsResponseEntryJSON contains the JSON metadata for the struct
// [UserListSubscriptionsResponseEntry]
type userListSubscriptionsResponseEntryJSON struct {
	Typename    apijson.Field
	InsertedAt  apijson.Field
	Object      apijson.Field
	Recipient   apijson.Field
	UpdatedAt   apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSubscriptionsResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSubscriptionsResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A custom-object entity which belongs to a collection.
type UserListSubscriptionsResponseEntriesObject struct {
	ID          string                                         `json:"id,required"`
	Typename    string                                         `json:"__typename,required"`
	Collection  string                                         `json:"collection,required"`
	UpdatedAt   time.Time                                      `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                      `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                         `json:"-,extras"`
	JSON        userListSubscriptionsResponseEntriesObjectJSON `json:"-"`
}

// userListSubscriptionsResponseEntriesObjectJSON contains the JSON metadata for
// the struct [UserListSubscriptionsResponseEntriesObject]
type userListSubscriptionsResponseEntriesObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSubscriptionsResponseEntriesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSubscriptionsResponseEntriesObjectJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type UserListSubscriptionsResponseEntriesRecipient struct {
	ID          string                                            `json:"id,required"`
	Typename    string                                            `json:"__typename,required"`
	UpdatedAt   time.Time                                         `json:"updated_at,required" format:"date-time"`
	Avatar      string                                            `json:"avatar,nullable"`
	Collection  string                                            `json:"collection"`
	CreatedAt   time.Time                                         `json:"created_at,nullable" format:"date-time"`
	Email       string                                            `json:"email,nullable" format:"email"`
	Name        string                                            `json:"name,nullable"`
	PhoneNumber string                                            `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                                            `json:"timezone,nullable"`
	JSON        userListSubscriptionsResponseEntriesRecipientJSON `json:"-"`
	union       UserListSubscriptionsResponseEntriesRecipientUnion
}

// userListSubscriptionsResponseEntriesRecipientJSON contains the JSON metadata for
// the struct [UserListSubscriptionsResponseEntriesRecipient]
type userListSubscriptionsResponseEntriesRecipientJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userListSubscriptionsResponseEntriesRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *UserListSubscriptionsResponseEntriesRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = UserListSubscriptionsResponseEntriesRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserListSubscriptionsResponseEntriesRecipientUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [UserListSubscriptionsResponseEntriesRecipientObject].
func (r UserListSubscriptionsResponseEntriesRecipient) AsUnion() UserListSubscriptionsResponseEntriesRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or
// [UserListSubscriptionsResponseEntriesRecipientObject].
type UserListSubscriptionsResponseEntriesRecipientUnion interface {
	implementsUserListSubscriptionsResponseEntriesRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListSubscriptionsResponseEntriesRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListSubscriptionsResponseEntriesRecipientObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type UserListSubscriptionsResponseEntriesRecipientObject struct {
	ID          string                                                  `json:"id,required"`
	Typename    string                                                  `json:"__typename,required"`
	Collection  string                                                  `json:"collection,required"`
	UpdatedAt   time.Time                                               `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                               `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                                  `json:"-,extras"`
	JSON        userListSubscriptionsResponseEntriesRecipientObjectJSON `json:"-"`
}

// userListSubscriptionsResponseEntriesRecipientObjectJSON contains the JSON
// metadata for the struct [UserListSubscriptionsResponseEntriesRecipientObject]
type userListSubscriptionsResponseEntriesRecipientObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSubscriptionsResponseEntriesRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSubscriptionsResponseEntriesRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserListSubscriptionsResponseEntriesRecipientObject) implementsUserListSubscriptionsResponseEntriesRecipient() {
}

// The information about a paginated result
type UserListSubscriptionsResponsePageInfo struct {
	Typename string                                    `json:"__typename,required"`
	PageSize int64                                     `json:"page_size,required"`
	After    string                                    `json:"after,nullable"`
	Before   string                                    `json:"before,nullable"`
	JSON     userListSubscriptionsResponsePageInfoJSON `json:"-"`
}

// userListSubscriptionsResponsePageInfoJSON contains the JSON metadata for the
// struct [UserListSubscriptionsResponsePageInfo]
type userListSubscriptionsResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSubscriptionsResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSubscriptionsResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

// Channel data for various channel types
type UserSetChannelDataResponse struct {
	Typename  string `json:"__typename,required"`
	ChannelID string `json:"channel_id,required" format:"uuid"`
	// Channel data for push providers
	Data UserSetChannelDataResponseData `json:"data,required"`
	JSON userSetChannelDataResponseJSON `json:"-"`
}

// userSetChannelDataResponseJSON contains the JSON metadata for the struct
// [UserSetChannelDataResponse]
type userSetChannelDataResponseJSON struct {
	Typename    apijson.Field
	ChannelID   apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseJSON) RawJSON() string {
	return r.raw
}

// Channel data for push providers
type UserSetChannelDataResponseData struct {
	// This field can have the runtime type of
	// [UserSetChannelDataResponseDataSlackChannelDataToken].
	Token interface{} `json:"token"`
	// This field can have the runtime type of
	// [[]UserSetChannelDataResponseDataSlackChannelDataConnection],
	// [[]UserSetChannelDataResponseDataMsTeamsChannelDataConnection],
	// [[]UserSetChannelDataResponseDataDiscordChannelDataConnection].
	Connections interface{} `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID string `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	// This field can have the runtime type of [[]string].
	PlayerIDs interface{} `json:"player_ids"`
	// This field can have the runtime type of [[]string].
	Tokens interface{}                        `json:"tokens"`
	JSON   userSetChannelDataResponseDataJSON `json:"-"`
	union  UserSetChannelDataResponseDataUnion
}

// userSetChannelDataResponseDataJSON contains the JSON metadata for the struct
// [UserSetChannelDataResponseData]
type userSetChannelDataResponseDataJSON struct {
	Token           apijson.Field
	Connections     apijson.Field
	MsTeamsTenantID apijson.Field
	PlayerIDs       apijson.Field
	Tokens          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r userSetChannelDataResponseDataJSON) RawJSON() string {
	return r.raw
}

func (r *UserSetChannelDataResponseData) UnmarshalJSON(data []byte) (err error) {
	*r = UserSetChannelDataResponseData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSetChannelDataResponseDataUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSetChannelDataResponseDataPushChannelData],
// [UserSetChannelDataResponseDataSlackChannelData],
// [UserSetChannelDataResponseDataMsTeamsChannelData],
// [UserSetChannelDataResponseDataDiscordChannelData],
// [UserSetChannelDataResponseDataOneSignalChannelData].
func (r UserSetChannelDataResponseData) AsUnion() UserSetChannelDataResponseDataUnion {
	return r.union
}

// Channel data for push providers
//
// Union satisfied by [UserSetChannelDataResponseDataPushChannelData],
// [UserSetChannelDataResponseDataSlackChannelData],
// [UserSetChannelDataResponseDataMsTeamsChannelData],
// [UserSetChannelDataResponseDataDiscordChannelData] or
// [UserSetChannelDataResponseDataOneSignalChannelData].
type UserSetChannelDataResponseDataUnion interface {
	implementsUserSetChannelDataResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetChannelDataResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataPushChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataSlackChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataMsTeamsChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataDiscordChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataOneSignalChannelData{}),
		},
	)
}

// Channel data for push providers
type UserSetChannelDataResponseDataPushChannelData struct {
	Tokens []string                                          `json:"tokens,required"`
	JSON   userSetChannelDataResponseDataPushChannelDataJSON `json:"-"`
}

// userSetChannelDataResponseDataPushChannelDataJSON contains the JSON metadata for
// the struct [UserSetChannelDataResponseDataPushChannelData]
type userSetChannelDataResponseDataPushChannelDataJSON struct {
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataPushChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataPushChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataPushChannelData) implementsUserSetChannelDataResponseData() {}

// Slack channel data
type UserSetChannelDataResponseDataSlackChannelData struct {
	Connections []UserSetChannelDataResponseDataSlackChannelDataConnection `json:"connections,required"`
	Token       UserSetChannelDataResponseDataSlackChannelDataToken        `json:"token,nullable"`
	JSON        userSetChannelDataResponseDataSlackChannelDataJSON         `json:"-"`
}

// userSetChannelDataResponseDataSlackChannelDataJSON contains the JSON metadata
// for the struct [UserSetChannelDataResponseDataSlackChannelData]
type userSetChannelDataResponseDataSlackChannelDataJSON struct {
	Connections apijson.Field
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataSlackChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataSlackChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataSlackChannelData) implementsUserSetChannelDataResponseData() {}

// A Slack connection, which either includes a channel_id or a user_id
type UserSetChannelDataResponseDataSlackChannelDataConnection struct {
	AccessToken string                                                       `json:"access_token,nullable"`
	ChannelID   string                                                       `json:"channel_id,nullable"`
	URL         string                                                       `json:"url"`
	UserID      string                                                       `json:"user_id,nullable"`
	JSON        userSetChannelDataResponseDataSlackChannelDataConnectionJSON `json:"-"`
	union       UserSetChannelDataResponseDataSlackChannelDataConnectionsUnion
}

// userSetChannelDataResponseDataSlackChannelDataConnectionJSON contains the JSON
// metadata for the struct
// [UserSetChannelDataResponseDataSlackChannelDataConnection]
type userSetChannelDataResponseDataSlackChannelDataConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	URL         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userSetChannelDataResponseDataSlackChannelDataConnectionJSON) RawJSON() string {
	return r.raw
}

func (r *UserSetChannelDataResponseDataSlackChannelDataConnection) UnmarshalJSON(data []byte) (err error) {
	*r = UserSetChannelDataResponseDataSlackChannelDataConnection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [UserSetChannelDataResponseDataSlackChannelDataConnectionsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection],
// [UserSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection].
func (r UserSetChannelDataResponseDataSlackChannelDataConnection) AsUnion() UserSetChannelDataResponseDataSlackChannelDataConnectionsUnion {
	return r.union
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Union satisfied by
// [UserSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection] or
// [UserSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection].
type UserSetChannelDataResponseDataSlackChannelDataConnectionsUnion interface {
	implementsUserSetChannelDataResponseDataSlackChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetChannelDataResponseDataSlackChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection{}),
		},
	)
}

// A Slack connection, which either includes a channel_id or a user_id
type UserSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection struct {
	AccessToken string                                                                       `json:"access_token,nullable"`
	ChannelID   string                                                                       `json:"channel_id,nullable"`
	UserID      string                                                                       `json:"user_id,nullable"`
	JSON        userSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnectionJSON `json:"-"`
}

// userSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnectionJSON
// contains the JSON metadata for the struct
// [UserSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection]
type userSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataSlackChannelDataConnectionsTokenConnection) implementsUserSetChannelDataResponseDataSlackChannelDataConnection() {
}

// An incoming webhook Slack connection
type UserSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection struct {
	URL  string                                                                                 `json:"url,required"`
	JSON userSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnectionJSON `json:"-"`
}

// userSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnectionJSON
// contains the JSON metadata for the struct
// [UserSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection]
type userSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataSlackChannelDataConnectionsIncomingWebhookConnection) implementsUserSetChannelDataResponseDataSlackChannelDataConnection() {
}

type UserSetChannelDataResponseDataSlackChannelDataToken struct {
	AccessToken string                                                  `json:"access_token,required,nullable"`
	JSON        userSetChannelDataResponseDataSlackChannelDataTokenJSON `json:"-"`
}

// userSetChannelDataResponseDataSlackChannelDataTokenJSON contains the JSON
// metadata for the struct [UserSetChannelDataResponseDataSlackChannelDataToken]
type userSetChannelDataResponseDataSlackChannelDataTokenJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataSlackChannelDataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataSlackChannelDataTokenJSON) RawJSON() string {
	return r.raw
}

// Microsoft Teams channel data
type UserSetChannelDataResponseDataMsTeamsChannelData struct {
	Connections []UserSetChannelDataResponseDataMsTeamsChannelDataConnection `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID string                                               `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	JSON            userSetChannelDataResponseDataMsTeamsChannelDataJSON `json:"-"`
}

// userSetChannelDataResponseDataMsTeamsChannelDataJSON contains the JSON metadata
// for the struct [UserSetChannelDataResponseDataMsTeamsChannelData]
type userSetChannelDataResponseDataMsTeamsChannelDataJSON struct {
	Connections     apijson.Field
	MsTeamsTenantID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataMsTeamsChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataMsTeamsChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataMsTeamsChannelData) implementsUserSetChannelDataResponseData() {
}

// A Slack connection, which either includes a channel_id or a user_id
type UserSetChannelDataResponseDataMsTeamsChannelDataConnection struct {
	AccessToken string                                                         `json:"access_token,nullable"`
	ChannelID   string                                                         `json:"channel_id,nullable"`
	URL         string                                                         `json:"url"`
	UserID      string                                                         `json:"user_id,nullable"`
	JSON        userSetChannelDataResponseDataMsTeamsChannelDataConnectionJSON `json:"-"`
	union       UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion
}

// userSetChannelDataResponseDataMsTeamsChannelDataConnectionJSON contains the JSON
// metadata for the struct
// [UserSetChannelDataResponseDataMsTeamsChannelDataConnection]
type userSetChannelDataResponseDataMsTeamsChannelDataConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	URL         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userSetChannelDataResponseDataMsTeamsChannelDataConnectionJSON) RawJSON() string {
	return r.raw
}

func (r *UserSetChannelDataResponseDataMsTeamsChannelDataConnection) UnmarshalJSON(data []byte) (err error) {
	*r = UserSetChannelDataResponseDataMsTeamsChannelDataConnection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection],
// [UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection].
func (r UserSetChannelDataResponseDataMsTeamsChannelDataConnection) AsUnion() UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion {
	return r.union
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Union satisfied by
// [UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection] or
// [UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection].
type UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion interface {
	implementsUserSetChannelDataResponseDataMsTeamsChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection{}),
		},
	)
}

// A Slack connection, which either includes a channel_id or a user_id
type UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection struct {
	AccessToken string                                                                         `json:"access_token,nullable"`
	ChannelID   string                                                                         `json:"channel_id,nullable"`
	UserID      string                                                                         `json:"user_id,nullable"`
	JSON        userSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnectionJSON `json:"-"`
}

// userSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnectionJSON
// contains the JSON metadata for the struct
// [UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection]
type userSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsTokenConnection) implementsUserSetChannelDataResponseDataMsTeamsChannelDataConnection() {
}

// An incoming webhook Slack connection
type UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection struct {
	URL  string                                                                                   `json:"url,required"`
	JSON userSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnectionJSON `json:"-"`
}

// userSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnectionJSON
// contains the JSON metadata for the struct
// [UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection]
type userSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataMsTeamsChannelDataConnectionsIncomingWebhookConnection) implementsUserSetChannelDataResponseDataMsTeamsChannelDataConnection() {
}

// Discord channel data
type UserSetChannelDataResponseDataDiscordChannelData struct {
	Connections []UserSetChannelDataResponseDataDiscordChannelDataConnection `json:"connections,required"`
	JSON        userSetChannelDataResponseDataDiscordChannelDataJSON         `json:"-"`
}

// userSetChannelDataResponseDataDiscordChannelDataJSON contains the JSON metadata
// for the struct [UserSetChannelDataResponseDataDiscordChannelData]
type userSetChannelDataResponseDataDiscordChannelDataJSON struct {
	Connections apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataDiscordChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataDiscordChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataDiscordChannelData) implementsUserSetChannelDataResponseData() {
}

// Discord channel connection
type UserSetChannelDataResponseDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID string                                                         `json:"channel_id"`
	URL       string                                                         `json:"url"`
	JSON      userSetChannelDataResponseDataDiscordChannelDataConnectionJSON `json:"-"`
	union     UserSetChannelDataResponseDataDiscordChannelDataConnectionsUnion
}

// userSetChannelDataResponseDataDiscordChannelDataConnectionJSON contains the JSON
// metadata for the struct
// [UserSetChannelDataResponseDataDiscordChannelDataConnection]
type userSetChannelDataResponseDataDiscordChannelDataConnectionJSON struct {
	ChannelID   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userSetChannelDataResponseDataDiscordChannelDataConnectionJSON) RawJSON() string {
	return r.raw
}

func (r *UserSetChannelDataResponseDataDiscordChannelDataConnection) UnmarshalJSON(data []byte) (err error) {
	*r = UserSetChannelDataResponseDataDiscordChannelDataConnection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [UserSetChannelDataResponseDataDiscordChannelDataConnectionsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection],
// [UserSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection].
func (r UserSetChannelDataResponseDataDiscordChannelDataConnection) AsUnion() UserSetChannelDataResponseDataDiscordChannelDataConnectionsUnion {
	return r.union
}

// Discord channel connection
//
// Union satisfied by
// [UserSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection]
// or
// [UserSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection].
type UserSetChannelDataResponseDataDiscordChannelDataConnectionsUnion interface {
	implementsUserSetChannelDataResponseDataDiscordChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetChannelDataResponseDataDiscordChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection{}),
		},
	)
}

// Discord channel connection
type UserSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection struct {
	// The Discord channel ID
	ChannelID string                                                                           `json:"channel_id,required"`
	JSON      userSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnectionJSON `json:"-"`
}

// userSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnectionJSON
// contains the JSON metadata for the struct
// [UserSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection]
type userSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnectionJSON struct {
	ChannelID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataDiscordChannelDataConnectionsChannelConnection) implementsUserSetChannelDataResponseDataDiscordChannelDataConnection() {
}

// An incoming webhook Slack connection
type UserSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection struct {
	URL  string                                                                                   `json:"url,required"`
	JSON userSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnectionJSON `json:"-"`
}

// userSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnectionJSON
// contains the JSON metadata for the struct
// [UserSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection]
type userSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataDiscordChannelDataConnectionsIncomingWebhookConnection) implementsUserSetChannelDataResponseDataDiscordChannelDataConnection() {
}

// OneSignal channel data
type UserSetChannelDataResponseDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs []string                                               `json:"player_ids,required" format:"uuid"`
	JSON      userSetChannelDataResponseDataOneSignalChannelDataJSON `json:"-"`
}

// userSetChannelDataResponseDataOneSignalChannelDataJSON contains the JSON
// metadata for the struct [UserSetChannelDataResponseDataOneSignalChannelData]
type userSetChannelDataResponseDataOneSignalChannelDataJSON struct {
	PlayerIDs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetChannelDataResponseDataOneSignalChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetChannelDataResponseDataOneSignalChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r UserSetChannelDataResponseDataOneSignalChannelData) implementsUserSetChannelDataResponseData() {
}

// A preference set object.
type UserSetPreferencesResponse struct {
	ID         string                                               `json:"id,required"`
	Typename   string                                               `json:"__typename,required"`
	Categories map[string]UserSetPreferencesResponseCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes UserSetPreferencesResponseChannelTypes              `json:"channel_types,nullable"`
	Workflows    map[string]UserSetPreferencesResponseWorkflowsUnion `json:"workflows,nullable"`
	JSON         userSetPreferencesResponseJSON                      `json:"-"`
}

// userSetPreferencesResponseJSON contains the JSON metadata for the struct
// [UserSetPreferencesResponse]
type userSetPreferencesResponseJSON struct {
	ID           apijson.Field
	Typename     apijson.Field
	Categories   apijson.Field
	ChannelTypes apijson.Field
	Workflows    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserSetPreferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseJSON) RawJSON() string {
	return r.raw
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseCategoriesObject].
type UserSetPreferencesResponseCategoriesUnion interface {
	ImplementsUserSetPreferencesResponseCategoriesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseCategoriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseCategoriesObject{}),
		},
	)
}

type UserSetPreferencesResponseCategoriesObject struct {
	// Channel type preferences
	ChannelTypes UserSetPreferencesResponseCategoriesObjectChannelTypes `json:"channel_types,nullable"`
	Conditions   []UserSetPreferencesResponseCategoriesObjectCondition  `json:"conditions"`
	JSON         userSetPreferencesResponseCategoriesObjectJSON         `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectJSON contains the JSON metadata for
// the struct [UserSetPreferencesResponseCategoriesObject]
type userSetPreferencesResponseCategoriesObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseCategoriesObject) ImplementsUserSetPreferencesResponseCategoriesUnion() {
}

// Channel type preferences
type UserSetPreferencesResponseCategoriesObjectChannelTypes struct {
	Chat      UserSetPreferencesResponseCategoriesObjectChannelTypesChatUnion      `json:"chat"`
	Email     UserSetPreferencesResponseCategoriesObjectChannelTypesEmailUnion     `json:"email"`
	HTTP      UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion      `json:"http"`
	InAppFeed UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      UserSetPreferencesResponseCategoriesObjectChannelTypesPushUnion      `json:"push"`
	SMS       UserSetPreferencesResponseCategoriesObjectChannelTypesSMSUnion       `json:"sms"`
	JSON      userSetPreferencesResponseCategoriesObjectChannelTypesJSON           `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesJSON contains the JSON
// metadata for the struct [UserSetPreferencesResponseCategoriesObjectChannelTypes]
type userSetPreferencesResponseCategoriesObjectChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditions].
type UserSetPreferencesResponseCategoriesObjectChannelTypesChatUnion interface {
	ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseCategoriesObjectChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditions{}),
		},
	)
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditions struct {
	Conditions []UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditions]
type userSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditions) ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition]
type userSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains             UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditions].
type UserSetPreferencesResponseCategoriesObjectChannelTypesEmailUnion interface {
	ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseCategoriesObjectChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditions{}),
		},
	)
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditions struct {
	Conditions []UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditions]
type userSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditions) ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition struct {
	Argument string                                                                                  `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                  `json:"variable,required"`
	JSON     userSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition]
type userSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains             UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions].
type UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion interface {
	ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions{}),
		},
	)
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions struct {
	Conditions []UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions]
type userSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditions) ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition]
type userSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions].
type UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions{}),
		},
	)
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions struct {
	Conditions []UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions]
type userSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditions) ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument string                                                                                      `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                      `json:"variable,required"`
	JSON     userSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition]
type userSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditions].
type UserSetPreferencesResponseCategoriesObjectChannelTypesPushUnion interface {
	ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseCategoriesObjectChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditions{}),
		},
	)
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditions struct {
	Conditions []UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditions]
type userSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditions) ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition]
type userSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains             UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditions].
type UserSetPreferencesResponseCategoriesObjectChannelTypesSMSUnion interface {
	ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseCategoriesObjectChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditions{}),
		},
	)
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditions struct {
	Conditions []UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON contains
// the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditions]
type userSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditions) ImplementsUserSetPreferencesResponseCategoriesObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition struct {
	Argument string                                                                                `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                `json:"variable,required"`
	JSON     userSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition]
type userSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains             UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserSetPreferencesResponseCategoriesObjectCondition struct {
	Argument string                                                       `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseCategoriesObjectConditionsOperator `json:"operator,required"`
	Variable string                                                       `json:"variable,required"`
	JSON     userSetPreferencesResponseCategoriesObjectConditionJSON      `json:"-"`
}

// userSetPreferencesResponseCategoriesObjectConditionJSON contains the JSON
// metadata for the struct [UserSetPreferencesResponseCategoriesObjectCondition]
type userSetPreferencesResponseCategoriesObjectConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesObjectCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesObjectConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseCategoriesObjectConditionsOperator string

const (
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorEqualTo              UserSetPreferencesResponseCategoriesObjectConditionsOperator = "equal_to"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorNotEqualTo           UserSetPreferencesResponseCategoriesObjectConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorGreaterThan          UserSetPreferencesResponseCategoriesObjectConditionsOperator = "greater_than"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorLessThan             UserSetPreferencesResponseCategoriesObjectConditionsOperator = "less_than"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseCategoriesObjectConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseCategoriesObjectConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorContains             UserSetPreferencesResponseCategoriesObjectConditionsOperator = "contains"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorNotContains          UserSetPreferencesResponseCategoriesObjectConditionsOperator = "not_contains"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorEmpty                UserSetPreferencesResponseCategoriesObjectConditionsOperator = "empty"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorNotEmpty             UserSetPreferencesResponseCategoriesObjectConditionsOperator = "not_empty"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorContainsAll          UserSetPreferencesResponseCategoriesObjectConditionsOperator = "contains_all"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestamp          UserSetPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseCategoriesObjectConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseCategoriesObjectConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsAudienceMember     UserSetPreferencesResponseCategoriesObjectConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseCategoriesObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseCategoriesObjectConditionsOperatorEqualTo, UserSetPreferencesResponseCategoriesObjectConditionsOperatorNotEqualTo, UserSetPreferencesResponseCategoriesObjectConditionsOperatorGreaterThan, UserSetPreferencesResponseCategoriesObjectConditionsOperatorLessThan, UserSetPreferencesResponseCategoriesObjectConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseCategoriesObjectConditionsOperatorContains, UserSetPreferencesResponseCategoriesObjectConditionsOperatorNotContains, UserSetPreferencesResponseCategoriesObjectConditionsOperatorEmpty, UserSetPreferencesResponseCategoriesObjectConditionsOperatorNotEmpty, UserSetPreferencesResponseCategoriesObjectConditionsOperatorContainsAll, UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestamp, UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseCategoriesObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type UserSetPreferencesResponseChannelTypes struct {
	Chat      UserSetPreferencesResponseChannelTypesChatUnion      `json:"chat"`
	Email     UserSetPreferencesResponseChannelTypesEmailUnion     `json:"email"`
	HTTP      UserSetPreferencesResponseChannelTypesHTTPUnion      `json:"http"`
	InAppFeed UserSetPreferencesResponseChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      UserSetPreferencesResponseChannelTypesPushUnion      `json:"push"`
	SMS       UserSetPreferencesResponseChannelTypesSMSUnion       `json:"sms"`
	JSON      userSetPreferencesResponseChannelTypesJSON           `json:"-"`
}

// userSetPreferencesResponseChannelTypesJSON contains the JSON metadata for the
// struct [UserSetPreferencesResponseChannelTypes]
type userSetPreferencesResponseChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseChannelTypesChatConditions].
type UserSetPreferencesResponseChannelTypesChatUnion interface {
	ImplementsUserSetPreferencesResponseChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseChannelTypesChatConditions{}),
		},
	)
}

type UserSetPreferencesResponseChannelTypesChatConditions struct {
	Conditions []UserSetPreferencesResponseChannelTypesChatConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseChannelTypesChatConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseChannelTypesChatConditionsJSON contains the JSON
// metadata for the struct [UserSetPreferencesResponseChannelTypesChatConditions]
type userSetPreferencesResponseChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseChannelTypesChatConditions) ImplementsUserSetPreferencesResponseChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseChannelTypesChatConditionsCondition struct {
	Argument string                                                                 `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                 `json:"variable,required"`
	JSON     userSetPreferencesResponseChannelTypesChatConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseChannelTypesChatConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserSetPreferencesResponseChannelTypesChatConditionsCondition]
type userSetPreferencesResponseChannelTypesChatConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesChatConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesChatConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator string

const (
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThan             UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorContains             UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotContains          UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorEmpty                UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThan, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorContains, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotContains, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorEmpty, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseChannelTypesEmailConditions].
type UserSetPreferencesResponseChannelTypesEmailUnion interface {
	ImplementsUserSetPreferencesResponseChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseChannelTypesEmailConditions{}),
		},
	)
}

type UserSetPreferencesResponseChannelTypesEmailConditions struct {
	Conditions []UserSetPreferencesResponseChannelTypesEmailConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseChannelTypesEmailConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseChannelTypesEmailConditionsJSON contains the JSON
// metadata for the struct [UserSetPreferencesResponseChannelTypesEmailConditions]
type userSetPreferencesResponseChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseChannelTypesEmailConditions) ImplementsUserSetPreferencesResponseChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseChannelTypesEmailConditionsCondition struct {
	Argument string                                                                  `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                  `json:"variable,required"`
	JSON     userSetPreferencesResponseChannelTypesEmailConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseChannelTypesEmailConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserSetPreferencesResponseChannelTypesEmailConditionsCondition]
type userSetPreferencesResponseChannelTypesEmailConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesEmailConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesEmailConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator string

const (
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThan             UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContains             UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotContains          UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEmpty                UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThan, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContains, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotContains, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorEmpty, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseChannelTypesHTTPConditions].
type UserSetPreferencesResponseChannelTypesHTTPUnion interface {
	ImplementsUserSetPreferencesResponseChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseChannelTypesHTTPConditions{}),
		},
	)
}

type UserSetPreferencesResponseChannelTypesHTTPConditions struct {
	Conditions []UserSetPreferencesResponseChannelTypesHTTPConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseChannelTypesHTTPConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseChannelTypesHTTPConditionsJSON contains the JSON
// metadata for the struct [UserSetPreferencesResponseChannelTypesHTTPConditions]
type userSetPreferencesResponseChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseChannelTypesHTTPConditions) ImplementsUserSetPreferencesResponseChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseChannelTypesHTTPConditionsCondition struct {
	Argument string                                                                 `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                 `json:"variable,required"`
	JSON     userSetPreferencesResponseChannelTypesHTTPConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseChannelTypesHTTPConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserSetPreferencesResponseChannelTypesHTTPConditionsCondition]
type userSetPreferencesResponseChannelTypesHTTPConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesHTTPConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesHTTPConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator string

const (
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThan             UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContains             UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotContains          UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEmpty                UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThan, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContains, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotContains, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorEmpty, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseChannelTypesInAppFeedConditions].
type UserSetPreferencesResponseChannelTypesInAppFeedUnion interface {
	ImplementsUserSetPreferencesResponseChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseChannelTypesInAppFeedConditions{}),
		},
	)
}

type UserSetPreferencesResponseChannelTypesInAppFeedConditions struct {
	Conditions []UserSetPreferencesResponseChannelTypesInAppFeedConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseChannelTypesInAppFeedConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseChannelTypesInAppFeedConditionsJSON contains the JSON
// metadata for the struct
// [UserSetPreferencesResponseChannelTypesInAppFeedConditions]
type userSetPreferencesResponseChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseChannelTypesInAppFeedConditions) ImplementsUserSetPreferencesResponseChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseChannelTypesInAppFeedConditionsCondition struct {
	Argument string                                                                      `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                      `json:"variable,required"`
	JSON     userSetPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON contains
// the JSON metadata for the struct
// [UserSetPreferencesResponseChannelTypesInAppFeedConditionsCondition]
type userSetPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesInAppFeedConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesInAppFeedConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContains             UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContains, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseChannelTypesPushConditions].
type UserSetPreferencesResponseChannelTypesPushUnion interface {
	ImplementsUserSetPreferencesResponseChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseChannelTypesPushConditions{}),
		},
	)
}

type UserSetPreferencesResponseChannelTypesPushConditions struct {
	Conditions []UserSetPreferencesResponseChannelTypesPushConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseChannelTypesPushConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseChannelTypesPushConditionsJSON contains the JSON
// metadata for the struct [UserSetPreferencesResponseChannelTypesPushConditions]
type userSetPreferencesResponseChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseChannelTypesPushConditions) ImplementsUserSetPreferencesResponseChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseChannelTypesPushConditionsCondition struct {
	Argument string                                                                 `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                 `json:"variable,required"`
	JSON     userSetPreferencesResponseChannelTypesPushConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseChannelTypesPushConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserSetPreferencesResponseChannelTypesPushConditionsCondition]
type userSetPreferencesResponseChannelTypesPushConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesPushConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesPushConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator string

const (
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThan             UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorContains             UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotContains          UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorEmpty                UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThan, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorContains, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotContains, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorEmpty, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseChannelTypesSMSConditions].
type UserSetPreferencesResponseChannelTypesSMSUnion interface {
	ImplementsUserSetPreferencesResponseChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseChannelTypesSMSConditions{}),
		},
	)
}

type UserSetPreferencesResponseChannelTypesSMSConditions struct {
	Conditions []UserSetPreferencesResponseChannelTypesSMSConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseChannelTypesSMSConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseChannelTypesSMSConditionsJSON contains the JSON
// metadata for the struct [UserSetPreferencesResponseChannelTypesSMSConditions]
type userSetPreferencesResponseChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseChannelTypesSMSConditions) ImplementsUserSetPreferencesResponseChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseChannelTypesSMSConditionsCondition struct {
	Argument string                                                                `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                `json:"variable,required"`
	JSON     userSetPreferencesResponseChannelTypesSMSConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseChannelTypesSMSConditionsConditionJSON contains the
// JSON metadata for the struct
// [UserSetPreferencesResponseChannelTypesSMSConditionsCondition]
type userSetPreferencesResponseChannelTypesSMSConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseChannelTypesSMSConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseChannelTypesSMSConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator string

const (
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThan             UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContains             UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotContains          UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEmpty                UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThan, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContains, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotContains, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorEmpty, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseWorkflowsObject].
type UserSetPreferencesResponseWorkflowsUnion interface {
	ImplementsUserSetPreferencesResponseWorkflowsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseWorkflowsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseWorkflowsObject{}),
		},
	)
}

type UserSetPreferencesResponseWorkflowsObject struct {
	// Channel type preferences
	ChannelTypes UserSetPreferencesResponseWorkflowsObjectChannelTypes `json:"channel_types,nullable"`
	Conditions   []UserSetPreferencesResponseWorkflowsObjectCondition  `json:"conditions"`
	JSON         userSetPreferencesResponseWorkflowsObjectJSON         `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectJSON contains the JSON metadata for the
// struct [UserSetPreferencesResponseWorkflowsObject]
type userSetPreferencesResponseWorkflowsObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseWorkflowsObject) ImplementsUserSetPreferencesResponseWorkflowsUnion() {
}

// Channel type preferences
type UserSetPreferencesResponseWorkflowsObjectChannelTypes struct {
	Chat      UserSetPreferencesResponseWorkflowsObjectChannelTypesChatUnion      `json:"chat"`
	Email     UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion     `json:"email"`
	HTTP      UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion      `json:"http"`
	InAppFeed UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      UserSetPreferencesResponseWorkflowsObjectChannelTypesPushUnion      `json:"push"`
	SMS       UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion       `json:"sms"`
	JSON      userSetPreferencesResponseWorkflowsObjectChannelTypesJSON           `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesJSON contains the JSON
// metadata for the struct [UserSetPreferencesResponseWorkflowsObjectChannelTypes]
type userSetPreferencesResponseWorkflowsObjectChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditions].
type UserSetPreferencesResponseWorkflowsObjectChannelTypesChatUnion interface {
	ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseWorkflowsObjectChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditions{}),
		},
	)
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditions struct {
	Conditions []UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON contains
// the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditions]
type userSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditions) ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition struct {
	Argument string                                                                                `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                `json:"variable,required"`
	JSON     userSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition]
type userSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains             UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions].
type UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion interface {
	ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions{}),
		},
	)
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions struct {
	Conditions []UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions]
type userSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditions) ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition struct {
	Argument string                                                                                 `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                 `json:"variable,required"`
	JSON     userSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition]
type userSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains             UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions].
type UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion interface {
	ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions{}),
		},
	)
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions struct {
	Conditions []UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON contains
// the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions]
type userSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditions) ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition struct {
	Argument string                                                                                `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                `json:"variable,required"`
	JSON     userSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition]
type userSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions].
type UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions{}),
		},
	)
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions struct {
	Conditions []UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions]
type userSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditions) ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument string                                                                                     `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                     `json:"variable,required"`
	JSON     userSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition]
type userSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditions].
type UserSetPreferencesResponseWorkflowsObjectChannelTypesPushUnion interface {
	ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseWorkflowsObjectChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditions{}),
		},
	)
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditions struct {
	Conditions []UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON contains
// the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditions]
type userSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditions) ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition struct {
	Argument string                                                                                `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                                `json:"variable,required"`
	JSON     userSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition]
type userSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains             UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions].
type UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion interface {
	ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions{}),
		},
	)
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions struct {
	Conditions []UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition `json:"conditions,required"`
	JSON       userSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON        `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON contains
// the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions]
type userSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditions) ImplementsUserSetPreferencesResponseWorkflowsObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition struct {
	Argument string                                                                               `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator `json:"operator,required"`
	Variable string                                                                               `json:"variable,required"`
	JSON     userSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON      `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition]
type userSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains             UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserSetPreferencesResponseWorkflowsObjectCondition struct {
	Argument string                                                      `json:"argument,required,nullable"`
	Operator UserSetPreferencesResponseWorkflowsObjectConditionsOperator `json:"operator,required"`
	Variable string                                                      `json:"variable,required"`
	JSON     userSetPreferencesResponseWorkflowsObjectConditionJSON      `json:"-"`
}

// userSetPreferencesResponseWorkflowsObjectConditionJSON contains the JSON
// metadata for the struct [UserSetPreferencesResponseWorkflowsObjectCondition]
type userSetPreferencesResponseWorkflowsObjectConditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsObjectCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsObjectConditionJSON) RawJSON() string {
	return r.raw
}

type UserSetPreferencesResponseWorkflowsObjectConditionsOperator string

const (
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorEqualTo              UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "equal_to"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorNotEqualTo           UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "not_equal_to"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThan          UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "greater_than"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorLessThan             UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "less_than"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorLessThanOrEqualTo    UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorContains             UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "contains"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorNotContains          UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "not_contains"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorEmpty                UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "empty"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorNotEmpty             UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "not_empty"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorContainsAll          UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "contains_all"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestamp          UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsNotTimestamp       UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "is_not_timestamp"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampAfter     UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp_after"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBefore    UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp_before"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBetween   UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "is_timestamp_between"
	UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsAudienceMember     UserSetPreferencesResponseWorkflowsObjectConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesResponseWorkflowsObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesResponseWorkflowsObjectConditionsOperatorEqualTo, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorNotEqualTo, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThan, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorLessThan, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorLessThanOrEqualTo, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorContains, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorNotContains, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorEmpty, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorNotEmpty, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorContainsAll, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestamp, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsNotTimestamp, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampAfter, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBefore, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsTimestampBetween, UserSetPreferencesResponseWorkflowsObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserGetPreferencesParams struct {
	// Tenant ID
	Tenant param.Field[string] `query:"tenant"`
}

// URLQuery serializes [UserGetPreferencesParams]'s query parameters as
// `url.Values`.
func (r UserGetPreferencesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserListMessagesParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The channel ID
	ChannelID param.Field[string] `query:"channel_id"`
	// The engagement status of the message
	EngagementStatus param.Field[[]UserListMessagesParamsEngagementStatus] `query:"engagement_status"`
	// The message IDs to filter messages by
	MessageIDs param.Field[[]string] `query:"message_ids"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// The source of the message (workflow key)
	Source param.Field[string] `query:"source"`
	// The status of the message
	Status param.Field[[]UserListMessagesParamsStatus] `query:"status"`
	// The tenant ID
	Tenant param.Field[string] `query:"tenant"`
	// The trigger data to filter messages by. Must be a valid JSON object.
	TriggerData param.Field[string] `query:"trigger_data"`
	// The workflow categories to filter messages by
	WorkflowCategories param.Field[[]string] `query:"workflow_categories"`
	// The workflow recipient run ID to filter messages by
	WorkflowRecipientRunID param.Field[string] `query:"workflow_recipient_run_id" format:"uuid"`
	// The workflow run ID to filter messages by
	WorkflowRunID param.Field[string] `query:"workflow_run_id" format:"uuid"`
}

// URLQuery serializes [UserListMessagesParams]'s query parameters as `url.Values`.
func (r UserListMessagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserListMessagesParamsEngagementStatus string

const (
	UserListMessagesParamsEngagementStatusSeen        UserListMessagesParamsEngagementStatus = "seen"
	UserListMessagesParamsEngagementStatusRead        UserListMessagesParamsEngagementStatus = "read"
	UserListMessagesParamsEngagementStatusInteracted  UserListMessagesParamsEngagementStatus = "interacted"
	UserListMessagesParamsEngagementStatusLinkClicked UserListMessagesParamsEngagementStatus = "link_clicked"
	UserListMessagesParamsEngagementStatusArchived    UserListMessagesParamsEngagementStatus = "archived"
)

func (r UserListMessagesParamsEngagementStatus) IsKnown() bool {
	switch r {
	case UserListMessagesParamsEngagementStatusSeen, UserListMessagesParamsEngagementStatusRead, UserListMessagesParamsEngagementStatusInteracted, UserListMessagesParamsEngagementStatusLinkClicked, UserListMessagesParamsEngagementStatusArchived:
		return true
	}
	return false
}

type UserListMessagesParamsStatus string

const (
	UserListMessagesParamsStatusQueued            UserListMessagesParamsStatus = "queued"
	UserListMessagesParamsStatusSent              UserListMessagesParamsStatus = "sent"
	UserListMessagesParamsStatusDelivered         UserListMessagesParamsStatus = "delivered"
	UserListMessagesParamsStatusDeliveryAttempted UserListMessagesParamsStatus = "delivery_attempted"
	UserListMessagesParamsStatusUndelivered       UserListMessagesParamsStatus = "undelivered"
	UserListMessagesParamsStatusNotSent           UserListMessagesParamsStatus = "not_sent"
	UserListMessagesParamsStatusBounced           UserListMessagesParamsStatus = "bounced"
)

func (r UserListMessagesParamsStatus) IsKnown() bool {
	switch r {
	case UserListMessagesParamsStatusQueued, UserListMessagesParamsStatusSent, UserListMessagesParamsStatusDelivered, UserListMessagesParamsStatusDeliveryAttempted, UserListMessagesParamsStatusUndelivered, UserListMessagesParamsStatusNotSent, UserListMessagesParamsStatusBounced:
		return true
	}
	return false
}

type UserListSchedulesParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// The ID of the tenant to list schedules for
	Tenant param.Field[string] `query:"tenant"`
	// The ID of the workflow to list schedules for
	Workflow param.Field[string] `query:"workflow"`
}

// URLQuery serializes [UserListSchedulesParams]'s query parameters as
// `url.Values`.
func (r UserListSchedulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserListSubscriptionsParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// Objects to filter by
	Objects param.Field[[]UserListSubscriptionsParamsObjectUnion] `query:"objects"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [UserListSubscriptionsParams]'s query parameters as
// `url.Values`.
func (r UserListSubscriptionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Satisfied by [shared.UnionString], [UserListSubscriptionsParamsObjectsObject].
type UserListSubscriptionsParamsObjectUnion interface {
	ImplementsUserListSubscriptionsParamsObjectUnion()
}

// An object reference to a recipient
type UserListSubscriptionsParamsObjectsObject struct {
	// An object identifier
	ID param.Field[string] `query:"id,required"`
	// The collection the object belongs to
	Collection param.Field[string] `query:"collection,required"`
}

// URLQuery serializes [UserListSubscriptionsParamsObjectsObject]'s query
// parameters as `url.Values`.
func (r UserListSubscriptionsParamsObjectsObject) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (r UserListSubscriptionsParamsObjectsObject) ImplementsUserListSubscriptionsParamsObjectUnion() {
}

type UserMergeParams struct {
	FromUserID param.Field[string] `json:"from_user_id"`
}

func (r UserMergeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParams struct {
	Categories param.Field[map[string]UserSetPreferencesParamsCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[UserSetPreferencesParamsChannelTypes]              `json:"channel_types"`
	Workflows    param.Field[map[string]UserSetPreferencesParamsWorkflowsUnion] `json:"workflows"`
}

func (r UserSetPreferencesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool], [UserSetPreferencesParamsCategoriesObject].
type UserSetPreferencesParamsCategoriesUnion interface {
	ImplementsUserSetPreferencesParamsCategoriesUnion()
}

type UserSetPreferencesParamsCategoriesObject struct {
	// Channel type preferences
	ChannelTypes param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserSetPreferencesParamsCategoriesObjectCondition]  `json:"conditions"`
}

func (r UserSetPreferencesParamsCategoriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsCategoriesObject) ImplementsUserSetPreferencesParamsCategoriesUnion() {
}

// Channel type preferences
type UserSetPreferencesParamsCategoriesObjectChannelTypes struct {
	Chat      param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPUnion]      `json:"http"`
	InAppFeed param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesSMSUnion]       `json:"sms"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditions].
type UserSetPreferencesParamsCategoriesObjectChannelTypesChatUnion interface {
	ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesChatUnion()
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditions) ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                               `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                               `json:"variable,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains             UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsCategoriesObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditions].
type UserSetPreferencesParamsCategoriesObjectChannelTypesEmailUnion interface {
	ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesEmailUnion()
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditions) ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                                `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                `json:"variable,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains             UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsCategoriesObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditions].
type UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPUnion interface {
	ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesHTTPUnion()
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditions) ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                               `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                               `json:"variable,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditions].
type UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedUnion()
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditions) ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                                    `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                    `json:"variable,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditions].
type UserSetPreferencesParamsCategoriesObjectChannelTypesPushUnion interface {
	ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesPushUnion()
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditions) ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                               `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                               `json:"variable,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains             UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsCategoriesObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditions].
type UserSetPreferencesParamsCategoriesObjectChannelTypesSMSUnion interface {
	ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesSMSUnion()
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditions) ImplementsUserSetPreferencesParamsCategoriesObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                              `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                              `json:"variable,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains             UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsCategoriesObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserSetPreferencesParamsCategoriesObjectCondition struct {
	Argument param.Field[string]                                                     `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsCategoriesObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                     `json:"variable,required"`
}

func (r UserSetPreferencesParamsCategoriesObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsCategoriesObjectConditionsOperator string

const (
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorEqualTo              UserSetPreferencesParamsCategoriesObjectConditionsOperator = "equal_to"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorNotEqualTo           UserSetPreferencesParamsCategoriesObjectConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorGreaterThan          UserSetPreferencesParamsCategoriesObjectConditionsOperator = "greater_than"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorLessThan             UserSetPreferencesParamsCategoriesObjectConditionsOperator = "less_than"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsCategoriesObjectConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsCategoriesObjectConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorContains             UserSetPreferencesParamsCategoriesObjectConditionsOperator = "contains"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorNotContains          UserSetPreferencesParamsCategoriesObjectConditionsOperator = "not_contains"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorEmpty                UserSetPreferencesParamsCategoriesObjectConditionsOperator = "empty"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorNotEmpty             UserSetPreferencesParamsCategoriesObjectConditionsOperator = "not_empty"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorContainsAll          UserSetPreferencesParamsCategoriesObjectConditionsOperator = "contains_all"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsTimestamp          UserSetPreferencesParamsCategoriesObjectConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsCategoriesObjectConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsCategoriesObjectConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsCategoriesObjectConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsCategoriesObjectConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsAudienceMember     UserSetPreferencesParamsCategoriesObjectConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsCategoriesObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsCategoriesObjectConditionsOperatorEqualTo, UserSetPreferencesParamsCategoriesObjectConditionsOperatorNotEqualTo, UserSetPreferencesParamsCategoriesObjectConditionsOperatorGreaterThan, UserSetPreferencesParamsCategoriesObjectConditionsOperatorLessThan, UserSetPreferencesParamsCategoriesObjectConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsCategoriesObjectConditionsOperatorContains, UserSetPreferencesParamsCategoriesObjectConditionsOperatorNotContains, UserSetPreferencesParamsCategoriesObjectConditionsOperatorEmpty, UserSetPreferencesParamsCategoriesObjectConditionsOperatorNotEmpty, UserSetPreferencesParamsCategoriesObjectConditionsOperatorContainsAll, UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsTimestamp, UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsCategoriesObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type UserSetPreferencesParamsChannelTypes struct {
	Chat      param.Field[UserSetPreferencesParamsChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[UserSetPreferencesParamsChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[UserSetPreferencesParamsChannelTypesHTTPUnion]      `json:"http"`
	InAppFeed param.Field[UserSetPreferencesParamsChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[UserSetPreferencesParamsChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[UserSetPreferencesParamsChannelTypesSMSUnion]       `json:"sms"`
}

func (r UserSetPreferencesParamsChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsChannelTypesChatConditions].
type UserSetPreferencesParamsChannelTypesChatUnion interface {
	ImplementsUserSetPreferencesParamsChannelTypesChatUnion()
}

type UserSetPreferencesParamsChannelTypesChatConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsChannelTypesChatConditions) ImplementsUserSetPreferencesParamsChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                               `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                               `json:"variable,required"`
}

func (r UserSetPreferencesParamsChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator string

const (
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorLessThan             UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorContains             UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorNotContains          UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorEmpty                UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorLessThan, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorContains, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorNotContains, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorEmpty, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsChannelTypesEmailConditions].
type UserSetPreferencesParamsChannelTypesEmailUnion interface {
	ImplementsUserSetPreferencesParamsChannelTypesEmailUnion()
}

type UserSetPreferencesParamsChannelTypesEmailConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsChannelTypesEmailConditions) ImplementsUserSetPreferencesParamsChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                `json:"variable,required"`
}

func (r UserSetPreferencesParamsChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator string

const (
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorLessThan             UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorContains             UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorNotContains          UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorEmpty                UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorLessThan, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorContains, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorNotContains, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorEmpty, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsChannelTypesHTTPConditions].
type UserSetPreferencesParamsChannelTypesHTTPUnion interface {
	ImplementsUserSetPreferencesParamsChannelTypesHTTPUnion()
}

type UserSetPreferencesParamsChannelTypesHTTPConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsChannelTypesHTTPConditions) ImplementsUserSetPreferencesParamsChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                               `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                               `json:"variable,required"`
}

func (r UserSetPreferencesParamsChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator string

const (
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorLessThan             UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorContains             UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorNotContains          UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorEmpty                UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorLessThan, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorContains, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorNotContains, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorEmpty, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsChannelTypesInAppFeedConditions].
type UserSetPreferencesParamsChannelTypesInAppFeedUnion interface {
	ImplementsUserSetPreferencesParamsChannelTypesInAppFeedUnion()
}

type UserSetPreferencesParamsChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsChannelTypesInAppFeedConditions) ImplementsUserSetPreferencesParamsChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                    `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                    `json:"variable,required"`
}

func (r UserSetPreferencesParamsChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorContains             UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorContains, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsChannelTypesPushConditions].
type UserSetPreferencesParamsChannelTypesPushUnion interface {
	ImplementsUserSetPreferencesParamsChannelTypesPushUnion()
}

type UserSetPreferencesParamsChannelTypesPushConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsChannelTypesPushConditions) ImplementsUserSetPreferencesParamsChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                               `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                               `json:"variable,required"`
}

func (r UserSetPreferencesParamsChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator string

const (
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorLessThan             UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorContains             UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorNotContains          UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorEmpty                UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorLessThan, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorContains, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorNotContains, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorEmpty, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsChannelTypesSMSConditions].
type UserSetPreferencesParamsChannelTypesSMSUnion interface {
	ImplementsUserSetPreferencesParamsChannelTypesSMSUnion()
}

type UserSetPreferencesParamsChannelTypesSMSConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsChannelTypesSMSConditions) ImplementsUserSetPreferencesParamsChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                              `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                              `json:"variable,required"`
}

func (r UserSetPreferencesParamsChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator string

const (
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorLessThan             UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorContains             UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorNotContains          UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorEmpty                UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorLessThan, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorContains, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorNotContains, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorEmpty, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool], [UserSetPreferencesParamsWorkflowsObject].
type UserSetPreferencesParamsWorkflowsUnion interface {
	ImplementsUserSetPreferencesParamsWorkflowsUnion()
}

type UserSetPreferencesParamsWorkflowsObject struct {
	// Channel type preferences
	ChannelTypes param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserSetPreferencesParamsWorkflowsObjectCondition]  `json:"conditions"`
}

func (r UserSetPreferencesParamsWorkflowsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsWorkflowsObject) ImplementsUserSetPreferencesParamsWorkflowsUnion() {}

// Channel type preferences
type UserSetPreferencesParamsWorkflowsObjectChannelTypes struct {
	Chat      param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesChatUnion]      `json:"chat"`
	Email     param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailUnion]     `json:"email"`
	HTTP      param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPUnion]      `json:"http"`
	InAppFeed param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	Push      param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesPushUnion]      `json:"push"`
	SMS       param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSUnion]       `json:"sms"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditions].
type UserSetPreferencesParamsWorkflowsObjectChannelTypesChatUnion interface {
	ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesChatUnion()
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditions) ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsCondition struct {
	Argument param.Field[string]                                                                              `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                              `json:"variable,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator string

const (
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan             UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains             UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains          UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty                UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsWorkflowsObjectChannelTypesChatConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditions].
type UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailUnion interface {
	ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesEmailUnion()
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditions) ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsCondition struct {
	Argument param.Field[string]                                                                               `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                               `json:"variable,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator string

const (
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan             UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains             UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains          UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty                UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditions].
type UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPUnion interface {
	ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPUnion()
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditions) ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsCondition struct {
	Argument param.Field[string]                                                                              `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                              `json:"variable,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator string

const (
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan             UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains             UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains          UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty                UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditions].
type UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedUnion()
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditions) ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsCondition struct {
	Argument param.Field[string]                                                                                   `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                   `json:"variable,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator string

const (
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan             UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains             UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains          UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty                UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditions].
type UserSetPreferencesParamsWorkflowsObjectChannelTypesPushUnion interface {
	ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesPushUnion()
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditions) ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsCondition struct {
	Argument param.Field[string]                                                                              `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                              `json:"variable,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator string

const (
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan             UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains             UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains          UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty                UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsWorkflowsObjectChannelTypesPushConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool],
// [UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditions].
type UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSUnion interface {
	ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesSMSUnion()
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditions struct {
	Conditions param.Field[[]UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsCondition] `json:"conditions,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditions) ImplementsUserSetPreferencesParamsWorkflowsObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsCondition struct {
	Argument param.Field[string]                                                                             `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                             `json:"variable,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator string

const (
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo              UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo           UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan          UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan             UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains             UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains          UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_contains"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty                UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty             UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "not_empty"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll          UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "contains_all"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp          UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember     UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type UserSetPreferencesParamsWorkflowsObjectCondition struct {
	Argument param.Field[string]                                                    `json:"argument,required"`
	Operator param.Field[UserSetPreferencesParamsWorkflowsObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                    `json:"variable,required"`
}

func (r UserSetPreferencesParamsWorkflowsObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetPreferencesParamsWorkflowsObjectConditionsOperator string

const (
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorEqualTo              UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "equal_to"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorNotEqualTo           UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "not_equal_to"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorGreaterThan          UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "greater_than"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorLessThan             UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "less_than"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "greater_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorLessThanOrEqualTo    UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "less_than_or_equal_to"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorContains             UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "contains"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorNotContains          UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "not_contains"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorEmpty                UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "empty"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorNotEmpty             UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "not_empty"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorContainsAll          UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "contains_all"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsTimestamp          UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "is_timestamp"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsNotTimestamp       UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "is_not_timestamp"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsTimestampAfter     UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "is_timestamp_after"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsTimestampBefore    UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "is_timestamp_before"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsTimestampBetween   UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "is_timestamp_between"
	UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsAudienceMember     UserSetPreferencesParamsWorkflowsObjectConditionsOperator = "is_audience_member"
)

func (r UserSetPreferencesParamsWorkflowsObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserSetPreferencesParamsWorkflowsObjectConditionsOperatorEqualTo, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorNotEqualTo, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorGreaterThan, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorLessThan, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorLessThanOrEqualTo, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorContains, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorNotContains, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorEmpty, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorNotEmpty, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorContainsAll, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsTimestamp, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsNotTimestamp, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsTimestampAfter, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsTimestampBefore, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsTimestampBetween, UserSetPreferencesParamsWorkflowsObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

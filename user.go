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
	"github.com/stainless-sdks/knock-go/packages/pagination"
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

// Identify user
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

// List users
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[User], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/users"
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

// List users
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[User] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete user
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

// Get user
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

// Get channel data
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

// Get preference set
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

// List messages
func (r *UserService) ListMessages(ctx context.Context, userID string, query UserListMessagesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[UserListMessagesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/messages", userID)
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

// List messages
func (r *UserService) ListMessagesAutoPaging(ctx context.Context, userID string, query UserListMessagesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[UserListMessagesResponse] {
	return pagination.NewEntriesCursorAutoPager(r.ListMessages(ctx, userID, query, opts...))
}

// List preference sets
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

// List schedules
func (r *UserService) ListSchedules(ctx context.Context, userID string, query UserListSchedulesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[UserListSchedulesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/schedules", userID)
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

// List schedules
func (r *UserService) ListSchedulesAutoPaging(ctx context.Context, userID string, query UserListSchedulesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[UserListSchedulesResponse] {
	return pagination.NewEntriesCursorAutoPager(r.ListSchedules(ctx, userID, query, opts...))
}

// List subscriptions
func (r *UserService) ListSubscriptions(ctx context.Context, userID string, query UserListSubscriptionsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[UserListSubscriptionsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/subscriptions", userID)
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

// List subscriptions
func (r *UserService) ListSubscriptionsAutoPaging(ctx context.Context, userID string, query UserListSubscriptionsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[UserListSubscriptionsResponse] {
	return pagination.NewEntriesCursorAutoPager(r.ListSubscriptions(ctx, userID, query, opts...))
}

// Merge two users together, where the user specified with the `from_user_id` param
// will be merged into the user specified by `user_id`.
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

// Set channel data
func (r *UserService) SetChannelData(ctx context.Context, userID string, channelID string, body UserSetChannelDataParams, opts ...option.RequestOption) (res *UserSetChannelDataResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Updates a complete preference set for a user. This is a destructive operation
// that will replace the existing preference set for the user.
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

// Unset channel data
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
	Email       string                 `json:"email,nullable"`
	Name        string                 `json:"name,nullable"`
	PhoneNumber string                 `json:"phone_number,nullable"`
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

func (r User) ImplementsUserListSchedulesResponseRecipient() {}

func (r User) ImplementsUserListSchedulesResponseActor() {}

func (r User) ImplementsUserListSubscriptionsResponseRecipient() {}

func (r User) ImplementsUserFeedListItemsResponseActivitiesActor() {}

func (r User) ImplementsUserFeedListItemsResponseActivitiesRecipient() {}

func (r User) ImplementsUserFeedListItemsResponseActor() {}

func (r User) ImplementsObjectAddSubscriptionsResponseRecipient() {}

func (r User) ImplementsObjectDeleteSubscriptionsResponseRecipient() {}

func (r User) ImplementsObjectListSchedulesResponseRecipient() {}

func (r User) ImplementsObjectListSchedulesResponseActor() {}

func (r User) ImplementsObjectListSubscriptionsResponseRecipient() {}

func (r User) ImplementsMessageListActivitiesResponseActor() {}

func (r User) ImplementsMessageListActivitiesResponseRecipient() {}

func (r User) ImplementsScheduleNewResponseRecipient() {}

func (r User) ImplementsScheduleNewResponseActor() {}

func (r User) ImplementsScheduleUpdateResponseRecipient() {}

func (r User) ImplementsScheduleUpdateResponseActor() {}

func (r User) ImplementsScheduleListResponseRecipient() {}

func (r User) ImplementsScheduleListResponseActor() {}

func (r User) ImplementsScheduleDeleteResponseRecipient() {}

func (r User) ImplementsScheduleDeleteResponseActor() {}

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
	// This field can have the runtime type of [shared.SlackChannelDataToken].
	Token interface{} `json:"token"`
	// This field can have the runtime type of [[]shared.SlackChannelDataConnection],
	// [[]shared.MsTeamsChannelDataConnection],
	// [[]shared.DiscordChannelDataConnection].
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
// Possible runtime types of the union are [shared.PushChannelData],
// [shared.SlackChannelData], [shared.MsTeamsChannelData],
// [shared.DiscordChannelData], [shared.OneSignalChannelData].
func (r UserGetChannelDataResponseData) AsUnion() UserGetChannelDataResponseDataUnion {
	return r.union
}

// Channel data for push providers
//
// Union satisfied by [shared.PushChannelData], [shared.SlackChannelData],
// [shared.MsTeamsChannelData], [shared.DiscordChannelData] or
// [shared.OneSignalChannelData].
type UserGetChannelDataResponseDataUnion interface {
	ImplementsUserGetChannelDataResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetChannelDataResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.PushChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.SlackChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.MsTeamsChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.DiscordChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.OneSignalChannelData{}),
		},
	)
}

// A preference set object.
type UserGetPreferencesResponse struct {
	ID       string `json:"id,required"`
	Typename string `json:"__typename,required"`
	// A map of categories and their settings
	Categories map[string]UserGetPreferencesResponseCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// A map of workflows and their settings
	Workflows map[string]UserGetPreferencesResponseWorkflowsUnion `json:"workflows,nullable"`
	JSON      userGetPreferencesResponseJSON                      `json:"-"`
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
// [UserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject].
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
			Type:       reflect.TypeOf(UserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                   `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                 `json:"conditions,nullable"`
	JSON         userGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// userGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject]
type userGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsUserGetPreferencesResponseCategoriesUnion() {
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [UserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject].
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
			Type:       reflect.TypeOf(UserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                  `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                `json:"conditions,nullable"`
	JSON         userGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// userGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [UserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject]
type userGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsUserGetPreferencesResponseWorkflowsUnion() {
}

// Represents a single message that was generated by a workflow for a given
// channel.
type UserListMessagesResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []UserListMessagesResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data map[string]interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []UserListMessagesResponseEngagementStatus `json:"engagement_statuses"`
	// Timestamp of creation
	InsertedAt time.Time `json:"inserted_at" format:"date-time"`
	// Timestamp when message was interacted with
	InteractedAt time.Time `json:"interacted_at,nullable" format:"date-time"`
	// Timestamp when a link in the message was clicked
	LinkClickedAt time.Time `json:"link_clicked_at,nullable" format:"date-time"`
	// Message metadata
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Timestamp when message was read
	ReadAt time.Time `json:"read_at,nullable" format:"date-time"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (id, collection).
	Recipient UserListMessagesResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source UserListMessagesResponseSource `json:"source"`
	// Message delivery status
	Status UserListMessagesResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                       `json:"workflow,nullable"`
	JSON     userListMessagesResponseJSON `json:"-"`
}

// userListMessagesResponseJSON contains the JSON metadata for the struct
// [UserListMessagesResponse]
type userListMessagesResponseJSON struct {
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

func (r *UserListMessagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [UserListMessagesResponseActorsObjectReference].
type UserListMessagesResponseActorsUnion interface {
	ImplementsUserListMessagesResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListMessagesResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListMessagesResponseActorsObjectReference{}),
		},
	)
}

// An object reference to a recipient
type UserListMessagesResponseActorsObjectReference struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                            `json:"collection,required"`
	JSON       userListMessagesResponseActorsObjectReferenceJSON `json:"-"`
}

// userListMessagesResponseActorsObjectReferenceJSON contains the JSON metadata for
// the struct [UserListMessagesResponseActorsObjectReference]
type userListMessagesResponseActorsObjectReferenceJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListMessagesResponseActorsObjectReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponseActorsObjectReferenceJSON) RawJSON() string {
	return r.raw
}

func (r UserListMessagesResponseActorsObjectReference) ImplementsUserListMessagesResponseActorsUnion() {
}

type UserListMessagesResponseEngagementStatus string

const (
	UserListMessagesResponseEngagementStatusSeen        UserListMessagesResponseEngagementStatus = "seen"
	UserListMessagesResponseEngagementStatusRead        UserListMessagesResponseEngagementStatus = "read"
	UserListMessagesResponseEngagementStatusInteracted  UserListMessagesResponseEngagementStatus = "interacted"
	UserListMessagesResponseEngagementStatusLinkClicked UserListMessagesResponseEngagementStatus = "link_clicked"
	UserListMessagesResponseEngagementStatusArchived    UserListMessagesResponseEngagementStatus = "archived"
)

func (r UserListMessagesResponseEngagementStatus) IsKnown() bool {
	switch r {
	case UserListMessagesResponseEngagementStatusSeen, UserListMessagesResponseEngagementStatusRead, UserListMessagesResponseEngagementStatusInteracted, UserListMessagesResponseEngagementStatusLinkClicked, UserListMessagesResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [UserListMessagesResponseRecipientObjectReference].
type UserListMessagesResponseRecipientUnion interface {
	ImplementsUserListMessagesResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListMessagesResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserListMessagesResponseRecipientObjectReference{}),
		},
	)
}

// An object reference to a recipient
type UserListMessagesResponseRecipientObjectReference struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                               `json:"collection,required"`
	JSON       userListMessagesResponseRecipientObjectReferenceJSON `json:"-"`
}

// userListMessagesResponseRecipientObjectReferenceJSON contains the JSON metadata
// for the struct [UserListMessagesResponseRecipientObjectReference]
type userListMessagesResponseRecipientObjectReferenceJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListMessagesResponseRecipientObjectReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponseRecipientObjectReferenceJSON) RawJSON() string {
	return r.raw
}

func (r UserListMessagesResponseRecipientObjectReference) ImplementsUserListMessagesResponseRecipientUnion() {
}

// Source information
type UserListMessagesResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                             `json:"version_id,required" format:"uuid"`
	JSON      userListMessagesResponseSourceJSON `json:"-"`
}

// userListMessagesResponseSourceJSON contains the JSON metadata for the struct
// [UserListMessagesResponseSource]
type userListMessagesResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListMessagesResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListMessagesResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type UserListMessagesResponseStatus string

const (
	UserListMessagesResponseStatusQueued            UserListMessagesResponseStatus = "queued"
	UserListMessagesResponseStatusSent              UserListMessagesResponseStatus = "sent"
	UserListMessagesResponseStatusDelivered         UserListMessagesResponseStatus = "delivered"
	UserListMessagesResponseStatusDeliveryAttempted UserListMessagesResponseStatus = "delivery_attempted"
	UserListMessagesResponseStatusUndelivered       UserListMessagesResponseStatus = "undelivered"
	UserListMessagesResponseStatusNotSent           UserListMessagesResponseStatus = "not_sent"
	UserListMessagesResponseStatusBounced           UserListMessagesResponseStatus = "bounced"
)

func (r UserListMessagesResponseStatus) IsKnown() bool {
	switch r {
	case UserListMessagesResponseStatusQueued, UserListMessagesResponseStatusSent, UserListMessagesResponseStatusDelivered, UserListMessagesResponseStatusDeliveryAttempted, UserListMessagesResponseStatusUndelivered, UserListMessagesResponseStatusNotSent, UserListMessagesResponseStatusBounced:
		return true
	}
	return false
}

// A preference set object.
type UserListPreferencesResponse struct {
	ID       string `json:"id,required"`
	Typename string `json:"__typename,required"`
	// A map of categories and their settings
	Categories map[string]UserListPreferencesResponseCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// A map of workflows and their settings
	Workflows map[string]UserListPreferencesResponseWorkflowsUnion `json:"workflows,nullable"`
	JSON      userListPreferencesResponseJSON                      `json:"-"`
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
// [UserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject].
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
			Type:       reflect.TypeOf(UserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                    `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                  `json:"conditions,nullable"`
	JSON         userListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// userListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject]
type userListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsUserListPreferencesResponseCategoriesUnion() {
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [UserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject].
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
			Type:       reflect.TypeOf(UserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                   `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                 `json:"conditions,nullable"`
	JSON         userListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// userListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [UserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject]
type userListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsUserListPreferencesResponseWorkflowsUnion() {
}

// A schedule that represents a recurring workflow execution
type UserListSchedulesResponse struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient UserListSchedulesResponseRecipient `json:"recipient,required"`
	Repeats   []shared.ScheduleRepeatRule        `json:"repeats,required"`
	UpdatedAt time.Time                          `json:"updated_at,required" format:"date-time"`
	Workflow  string                             `json:"workflow,required"`
	Typename  string                             `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            UserListSchedulesResponseActor `json:"actor,nullable"`
	Data             map[string]interface{}         `json:"data,nullable"`
	LastOccurrenceAt time.Time                      `json:"last_occurrence_at,nullable" format:"date-time"`
	NextOccurrenceAt time.Time                      `json:"next_occurrence_at,nullable" format:"date-time"`
	Tenant           string                         `json:"tenant,nullable"`
	JSON             userListSchedulesResponseJSON  `json:"-"`
}

// userListSchedulesResponseJSON contains the JSON metadata for the struct
// [UserListSchedulesResponse]
type userListSchedulesResponseJSON struct {
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

func (r *UserListSchedulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSchedulesResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type UserListSchedulesResponseRecipient struct {
	ID          string                                 `json:"id,required"`
	Typename    string                                 `json:"__typename,required"`
	UpdatedAt   time.Time                              `json:"updated_at,required" format:"date-time"`
	Avatar      string                                 `json:"avatar,nullable"`
	Collection  string                                 `json:"collection"`
	CreatedAt   time.Time                              `json:"created_at,nullable" format:"date-time"`
	Email       string                                 `json:"email,nullable"`
	Name        string                                 `json:"name,nullable"`
	PhoneNumber string                                 `json:"phone_number,nullable"`
	Timezone    string                                 `json:"timezone,nullable"`
	JSON        userListSchedulesResponseRecipientJSON `json:"-"`
	union       UserListSchedulesResponseRecipientUnion
}

// userListSchedulesResponseRecipientJSON contains the JSON metadata for the struct
// [UserListSchedulesResponseRecipient]
type userListSchedulesResponseRecipientJSON struct {
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

func (r userListSchedulesResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *UserListSchedulesResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = UserListSchedulesResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserListSchedulesResponseRecipientUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r UserListSchedulesResponseRecipient) AsUnion() UserListSchedulesResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type UserListSchedulesResponseRecipientUnion interface {
	ImplementsUserListSchedulesResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListSchedulesResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A recipient, which is either a user or an object
type UserListSchedulesResponseActor struct {
	ID          string                             `json:"id,required"`
	Typename    string                             `json:"__typename,required"`
	UpdatedAt   time.Time                          `json:"updated_at,required" format:"date-time"`
	Avatar      string                             `json:"avatar,nullable"`
	Collection  string                             `json:"collection"`
	CreatedAt   time.Time                          `json:"created_at,nullable" format:"date-time"`
	Email       string                             `json:"email,nullable"`
	Name        string                             `json:"name,nullable"`
	PhoneNumber string                             `json:"phone_number,nullable"`
	Timezone    string                             `json:"timezone,nullable"`
	JSON        userListSchedulesResponseActorJSON `json:"-"`
	union       UserListSchedulesResponseActorUnion
}

// userListSchedulesResponseActorJSON contains the JSON metadata for the struct
// [UserListSchedulesResponseActor]
type userListSchedulesResponseActorJSON struct {
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

func (r userListSchedulesResponseActorJSON) RawJSON() string {
	return r.raw
}

func (r *UserListSchedulesResponseActor) UnmarshalJSON(data []byte) (err error) {
	*r = UserListSchedulesResponseActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserListSchedulesResponseActorUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r UserListSchedulesResponseActor) AsUnion() UserListSchedulesResponseActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type UserListSchedulesResponseActorUnion interface {
	ImplementsUserListSchedulesResponseActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListSchedulesResponseActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A subscription object
type UserListSubscriptionsResponse struct {
	Typename   string    `json:"__typename,required"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A custom-object entity which belongs to a collection.
	Object shared.Object `json:"object,required"`
	// A recipient, which is either a user or an object
	Recipient UserListSubscriptionsResponseRecipient `json:"recipient,required"`
	UpdatedAt time.Time                              `json:"updated_at,required" format:"date-time"`
	// The custom properties associated with the subscription
	Properties map[string]interface{}            `json:"properties,nullable"`
	JSON       userListSubscriptionsResponseJSON `json:"-"`
}

// userListSubscriptionsResponseJSON contains the JSON metadata for the struct
// [UserListSubscriptionsResponse]
type userListSubscriptionsResponseJSON struct {
	Typename    apijson.Field
	InsertedAt  apijson.Field
	Object      apijson.Field
	Recipient   apijson.Field
	UpdatedAt   apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListSubscriptionsResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type UserListSubscriptionsResponseRecipient struct {
	ID          string                                     `json:"id,required"`
	Typename    string                                     `json:"__typename,required"`
	UpdatedAt   time.Time                                  `json:"updated_at,required" format:"date-time"`
	Avatar      string                                     `json:"avatar,nullable"`
	Collection  string                                     `json:"collection"`
	CreatedAt   time.Time                                  `json:"created_at,nullable" format:"date-time"`
	Email       string                                     `json:"email,nullable"`
	Name        string                                     `json:"name,nullable"`
	PhoneNumber string                                     `json:"phone_number,nullable"`
	Timezone    string                                     `json:"timezone,nullable"`
	JSON        userListSubscriptionsResponseRecipientJSON `json:"-"`
	union       UserListSubscriptionsResponseRecipientUnion
}

// userListSubscriptionsResponseRecipientJSON contains the JSON metadata for the
// struct [UserListSubscriptionsResponseRecipient]
type userListSubscriptionsResponseRecipientJSON struct {
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

func (r userListSubscriptionsResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *UserListSubscriptionsResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = UserListSubscriptionsResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserListSubscriptionsResponseRecipientUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r UserListSubscriptionsResponseRecipient) AsUnion() UserListSubscriptionsResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type UserListSubscriptionsResponseRecipientUnion interface {
	ImplementsUserListSubscriptionsResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListSubscriptionsResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
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
	// This field can have the runtime type of [shared.SlackChannelDataToken].
	Token interface{} `json:"token"`
	// This field can have the runtime type of [[]shared.SlackChannelDataConnection],
	// [[]shared.MsTeamsChannelDataConnection],
	// [[]shared.DiscordChannelDataConnection].
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
// Possible runtime types of the union are [shared.PushChannelData],
// [shared.SlackChannelData], [shared.MsTeamsChannelData],
// [shared.DiscordChannelData], [shared.OneSignalChannelData].
func (r UserSetChannelDataResponseData) AsUnion() UserSetChannelDataResponseDataUnion {
	return r.union
}

// Channel data for push providers
//
// Union satisfied by [shared.PushChannelData], [shared.SlackChannelData],
// [shared.MsTeamsChannelData], [shared.DiscordChannelData] or
// [shared.OneSignalChannelData].
type UserSetChannelDataResponseDataUnion interface {
	ImplementsUserSetChannelDataResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSetChannelDataResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.PushChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.SlackChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.MsTeamsChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.DiscordChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.OneSignalChannelData{}),
		},
	)
}

// A preference set object.
type UserSetPreferencesResponse struct {
	ID       string `json:"id,required"`
	Typename string `json:"__typename,required"`
	// A map of categories and their settings
	Categories map[string]UserSetPreferencesResponseCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// A map of workflows and their settings
	Workflows map[string]UserSetPreferencesResponseWorkflowsUnion `json:"workflows,nullable"`
	JSON      userSetPreferencesResponseJSON                      `json:"-"`
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
// [UserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject].
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
			Type:       reflect.TypeOf(UserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                   `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                 `json:"conditions,nullable"`
	JSON         userSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// userSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject]
type userSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsUserSetPreferencesResponseCategoriesUnion() {
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [UserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject].
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
			Type:       reflect.TypeOf(UserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type UserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                  `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                `json:"conditions,nullable"`
	JSON         userSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// userSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [UserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject]
type userSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsUserSetPreferencesResponseWorkflowsUnion() {
}

type UserUpdateParams struct {
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[shared.InlineChannelDataRequestParam] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                            `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[shared.InlinePreferenceSetRequestParam] `json:"preferences"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type UserMergeParams struct {
	// The user ID to merge from
	FromUserID param.Field[string] `json:"from_user_id,required"`
}

func (r UserMergeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetChannelDataParams struct {
	// Set channel data for a type of channel
	ChannelDataRequest shared.ChannelDataRequestParam `json:"channel_data_request,required"`
}

func (r UserSetChannelDataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ChannelDataRequest)
}

type UserSetPreferencesParams struct {
	// Set preferences for a recipient
	PreferenceSetRequest shared.PreferenceSetRequestParam `json:"preference_set_request,required"`
}

func (r UserSetPreferencesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PreferenceSetRequest)
}

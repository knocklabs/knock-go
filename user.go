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
	"github.com/stainless-sdks/knock-go/packages/pagination"
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

// Create or update a user with the provided identification data.
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

// Retrieve a paginated list of users in the environment.
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

// Retrieve a paginated list of users in the environment.
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[User] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete a user and all associated data.
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

// Retrieve a specific user by their ID.
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

// Retrieves the channel data for a specific user and channel ID.
func (r *UserService) GetChannelData(ctx context.Context, userID string, channelID string, opts ...option.RequestOption) (res *ChannelData, err error) {
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

// Retrieves a specific preference set for a user identified by the preference set
// ID.
func (r *UserService) GetPreferences(ctx context.Context, userID string, preferenceSetID string, query UserGetPreferencesParams, opts ...option.RequestOption) (res *PreferenceSet, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if preferenceSetID == "" {
		err = errors.New("missing required preference_set_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/preferences/%s", userID, preferenceSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a paginated list of messages for a specific user. Allows filtering by
// message status and provides various sorting options.
func (r *UserService) ListMessages(ctx context.Context, userID string, query UserListMessagesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Message], err error) {
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

// Returns a paginated list of messages for a specific user. Allows filtering by
// message status and provides various sorting options.
func (r *UserService) ListMessagesAutoPaging(ctx context.Context, userID string, query UserListMessagesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Message] {
	return pagination.NewEntriesCursorAutoPager(r.ListMessages(ctx, userID, query, opts...))
}

// Retrieves a list of all preference sets for a specific user.
func (r *UserService) ListPreferences(ctx context.Context, userID string, opts ...option.RequestOption) (res *[]PreferenceSet, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/preferences", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a paginated list of schedules for a specific user. Can be filtered by
// workflow and tenant.
func (r *UserService) ListSchedules(ctx context.Context, userID string, query UserListSchedulesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Schedule], err error) {
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

// Returns a paginated list of schedules for a specific user. Can be filtered by
// workflow and tenant.
func (r *UserService) ListSchedulesAutoPaging(ctx context.Context, userID string, query UserListSchedulesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Schedule] {
	return pagination.NewEntriesCursorAutoPager(r.ListSchedules(ctx, userID, query, opts...))
}

// Retrieves a paginated list of subscriptions for a specific user. Allows
// filtering by objects and includes optional preference data.
func (r *UserService) ListSubscriptions(ctx context.Context, userID string, query UserListSubscriptionsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Subscription], err error) {
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

// Retrieves a paginated list of subscriptions for a specific user. Allows
// filtering by objects and includes optional preference data.
func (r *UserService) ListSubscriptionsAutoPaging(ctx context.Context, userID string, query UserListSubscriptionsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Subscription] {
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

// Updates or creates channel data for a specific user and channel ID.
func (r *UserService) SetChannelData(ctx context.Context, userID string, channelID string, body UserSetChannelDataParams, opts ...option.RequestOption) (res *ChannelData, err error) {
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
func (r *UserService) SetPreferences(ctx context.Context, userID string, preferenceSetID string, body UserSetPreferencesParams, opts ...option.RequestOption) (res *PreferenceSet, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if preferenceSetID == "" {
		err = errors.New("missing required preference_set_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/preferences/%s", userID, preferenceSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes channel data for a specific user and channel ID.
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

// A set of parameters to identify a user with. Does not include the user ID, as
// that's specified elsewhere in the request. You can supply any additional
// properties you'd like to upsert against the user.
type IdentifyUserRequestParam struct {
	// A request to set channel data for a type of channel inline.
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
}

func (r IdentifyUserRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of parameters to inline-identify a user with. Inline identifying the user
// will ensure that the user is available before the request is executed in Knock.
// It will perform an upsert against the user you're supplying, replacing any
// properties specified.
type InlineIdentifyUserRequestParam struct {
	// The unique identifier for the user.
	ID param.Field[string] `json:"id"`
	// A request to set channel data for a type of channel inline.
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
}

func (r InlineIdentifyUserRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InlineIdentifyUserRequestParam) ImplementsRecipientRequestUnionParam() {}

// A user object.
type User struct {
	// The unique identifier for the user.
	ID string `json:"id,required"`
	// The type name of the schema.
	Typename string `json:"__typename,required"`
	// The timestamp when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// URL to the user's avatar image.
	Avatar string `json:"avatar,nullable"`
	// Timestamp when the resource was created.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// The email address of the user.
	Email string `json:"email,nullable"`
	// Display name of the user.
	Name string `json:"name,nullable"`
	// Phone number of the user.
	PhoneNumber string `json:"phone_number,nullable"`
	// Timezone of the user.
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

func (r User) implementsRecipient() {}

type UserUpdateParams struct {
	// A set of parameters to identify a user with. Does not include the user ID, as
	// that's specified elsewhere in the request. You can supply any additional
	// properties you'd like to upsert against the user.
	IdentifyUserRequest IdentifyUserRequestParam `json:"identify_user_request,required"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IdentifyUserRequest)
}

type UserListParams struct {
	// The cursor to fetch entries after..
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before..
	Before param.Field[string] `query:"before"`
	// Includes preferences of the users in the response.
	Include param.Field[[]UserListParamsInclude] `query:"include"`
	// The number of items per page..
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserListParamsInclude string

const (
	UserListParamsIncludePreferences UserListParamsInclude = "preferences"
)

func (r UserListParamsInclude) IsKnown() bool {
	switch r {
	case UserListParamsIncludePreferences:
		return true
	}
	return false
}

type UserGetPreferencesParams struct {
	// The unique identifier for the tenant.
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
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// The unique identifier for the channel.
	ChannelID param.Field[string] `query:"channel_id"`
	// The engagement status to filter messages by.
	EngagementStatus param.Field[[]UserListMessagesParamsEngagementStatus] `query:"engagement_status"`
	// The message IDs to filter messages by.
	MessageIDs param.Field[[]string] `query:"message_ids"`
	// The number of items per page.
	PageSize param.Field[int64] `query:"page_size"`
	// The source of the message (workflow key).
	Source param.Field[string] `query:"source"`
	// The delivery status to filter messages by.
	Status param.Field[[]UserListMessagesParamsStatus] `query:"status"`
	// The unique identifier for the tenant.
	Tenant param.Field[string] `query:"tenant"`
	// The trigger data to filter messages by. Must be a valid JSON object.
	TriggerData param.Field[string] `query:"trigger_data"`
	// The workflow categories to filter messages by.
	WorkflowCategories param.Field[[]string] `query:"workflow_categories"`
	// The workflow recipient run ID to filter messages by.
	WorkflowRecipientRunID param.Field[string] `query:"workflow_recipient_run_id" format:"uuid"`
	// The workflow run ID to filter messages by.
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
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// The number of items per page.
	PageSize param.Field[int64] `query:"page_size"`
	// The ID of the tenant to list schedules for.
	Tenant param.Field[string] `query:"tenant"`
	// The ID of the workflow to list schedules for.
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
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// Includes preferences of the recipient subscribers in the response.
	Include param.Field[[]UserListSubscriptionsParamsInclude] `query:"include"`
	// Objects to filter by.
	Objects param.Field[[]UserListSubscriptionsParamsObjectUnion] `query:"objects"`
	// The number of items per page.
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

type UserListSubscriptionsParamsInclude string

const (
	UserListSubscriptionsParamsIncludePreferences UserListSubscriptionsParamsInclude = "preferences"
)

func (r UserListSubscriptionsParamsInclude) IsKnown() bool {
	switch r {
	case UserListSubscriptionsParamsIncludePreferences:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Satisfied by [shared.UnionString],
// [UserListSubscriptionsParamsObjectsRecipientReference].
type UserListSubscriptionsParamsObjectUnion interface {
	ImplementsUserListSubscriptionsParamsObjectUnion()
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
type UserListSubscriptionsParamsObjectsRecipientReference struct {
	// An identifier for the recipient object.
	ID param.Field[string] `query:"id"`
	// The collection the recipient object belongs to.
	Collection param.Field[string] `query:"collection"`
}

// URLQuery serializes [UserListSubscriptionsParamsObjectsRecipientReference]'s
// query parameters as `url.Values`.
func (r UserListSubscriptionsParamsObjectsRecipientReference) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (r UserListSubscriptionsParamsObjectsRecipientReference) ImplementsUserListSubscriptionsParamsObjectUnion() {
}

type UserMergeParams struct {
	// The user ID to merge from.
	FromUserID param.Field[string] `json:"from_user_id,required"`
}

func (r UserMergeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSetChannelDataParams struct {
	// A request to set channel data for a type of channel.
	ChannelDataRequest ChannelDataRequestParam `json:"channel_data_request,required"`
}

func (r UserSetChannelDataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ChannelDataRequest)
}

type UserSetPreferencesParams struct {
	// A request to set a preference set for a recipient.
	PreferenceSetRequest PreferenceSetRequestParam `json:"preference_set_request,required"`
}

func (r UserSetPreferencesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PreferenceSetRequest)
}

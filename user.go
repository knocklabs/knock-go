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
func (r *UserService) Update(ctx context.Context, userID string, body UserUpdateParams, opts ...option.RequestOption) (res *shared.User, err error) {
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
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[shared.User], err error) {
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
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[shared.User] {
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
func (r *UserService) Get(ctx context.Context, userID string, opts ...option.RequestOption) (res *shared.User, err error) {
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
func (r *UserService) GetChannelData(ctx context.Context, userID string, channelID string, opts ...option.RequestOption) (res *shared.ChannelData, err error) {
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
func (r *UserService) GetPreferences(ctx context.Context, userID string, id string, query UserGetPreferencesParams, opts ...option.RequestOption) (res *shared.PreferenceSet, err error) {
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
func (r *UserService) ListPreferences(ctx context.Context, userID string, opts ...option.RequestOption) (res *[]shared.PreferenceSet, err error) {
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
func (r *UserService) ListSchedules(ctx context.Context, userID string, query UserListSchedulesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[shared.Schedule], err error) {
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
func (r *UserService) ListSchedulesAutoPaging(ctx context.Context, userID string, query UserListSchedulesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[shared.Schedule] {
	return pagination.NewEntriesCursorAutoPager(r.ListSchedules(ctx, userID, query, opts...))
}

// List subscriptions
func (r *UserService) ListSubscriptions(ctx context.Context, userID string, query UserListSubscriptionsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[shared.Subscription], err error) {
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
func (r *UserService) ListSubscriptionsAutoPaging(ctx context.Context, userID string, query UserListSubscriptionsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[shared.Subscription] {
	return pagination.NewEntriesCursorAutoPager(r.ListSubscriptions(ctx, userID, query, opts...))
}

// Merge two users together, where the user specified with the `from_user_id` param
// will be merged into the user specified by `user_id`.
func (r *UserService) Merge(ctx context.Context, userID string, body UserMergeParams, opts ...option.RequestOption) (res *shared.User, err error) {
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
func (r *UserService) SetChannelData(ctx context.Context, userID string, channelID string, body UserSetChannelDataParams, opts ...option.RequestOption) (res *shared.ChannelData, err error) {
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
func (r *UserService) SetPreferences(ctx context.Context, userID string, id string, body UserSetPreferencesParams, opts ...option.RequestOption) (res *shared.PreferenceSet, err error) {
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

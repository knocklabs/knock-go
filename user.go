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
	ID          string            `json:"id,required"`
	UpdatedAt   time.Time         `json:"updated_at,required" format:"date-time"`
	Typename    UserTypename      `json:"__typename"`
	CreatedAt   time.Time         `json:"created_at" format:"date-time"`
	ExtraFields map[string]string `json:"-,extras"`
	JSON        userJSON          `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID          apijson.Field
	UpdatedAt   apijson.Field
	Typename    apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type UserTypename string

const (
	UserTypenameUser UserTypename = "User"
)

func (r UserTypename) IsKnown() bool {
	switch r {
	case UserTypenameUser:
		return true
	}
	return false
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
	Typename UserListResponsePageInfoTypename `json:"__typename"`
	After    string                           `json:"after"`
	Before   string                           `json:"before"`
	PageSize int64                            `json:"page_size"`
	JSON     userListResponsePageInfoJSON     `json:"-"`
}

// userListResponsePageInfoJSON contains the JSON metadata for the struct
// [UserListResponsePageInfo]
type userListResponsePageInfoJSON struct {
	Typename    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	PageSize    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

type UserListResponsePageInfoTypename string

const (
	UserListResponsePageInfoTypenamePageInfo UserListResponsePageInfoTypename = "PageInfo"
)

func (r UserListResponsePageInfoTypename) IsKnown() bool {
	switch r {
	case UserListResponsePageInfoTypenamePageInfo:
		return true
	}
	return false
}

type UserUpdateParams struct {
	ChannelData param.Field[map[string]UserUpdateParamsChannelData] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                              `json:"created_at" format:"date-time"`
	Preferences param.Field[map[string]UserUpdateParamsPreferences] `json:"preferences"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsChannelData struct {
	Data param.Field[UserUpdateParamsChannelDataData] `json:"data"`
}

func (r UserUpdateParamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsChannelDataData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r UserUpdateParamsChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferences struct {
	Categories   param.Field[UserUpdateParamsPreferencesCategoriesUnion] `json:"categories"`
	ChannelTypes param.Field[UserUpdateParamsPreferencesChannelTypes]    `json:"channel_types"`
	Workflows    param.Field[UserUpdateParamsPreferencesWorkflowsUnion]  `json:"workflows"`
}

func (r UserUpdateParamsPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategories struct {
	ChannelTypes param.Field[interface{}] `json:"channel_types,required"`
	Conditions   param.Field[interface{}] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategories) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategories) ImplementsUserUpdateParamsPreferencesCategoriesUnion() {
}

// Satisfied by [shared.UnionBool], [UserUpdateParamsPreferencesCategoriesObject],
// [UserUpdateParamsPreferencesCategories].
type UserUpdateParamsPreferencesCategoriesUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesUnion()
}

type UserUpdateParamsPreferencesCategoriesObject struct {
	ChannelTypes param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserUpdateParamsPreferencesCategoriesObjectCondition]  `json:"conditions"`
}

func (r UserUpdateParamsPreferencesCategoriesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObject) ImplementsUserUpdateParamsPreferencesCategoriesUnion() {
}

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

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesChat struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChat) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObject],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesChat].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObject) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesChatUnion() {
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectCondition struct {
	Argument param.Field[string]                                                                              `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                              `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmail struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmail) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObject],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmail].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObject) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailUnion() {
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectCondition struct {
	Argument param.Field[string]                                                                               `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                               `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTP struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTP) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObject],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTP].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObject) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPUnion() {
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectCondition struct {
	Argument param.Field[string]                                                                              `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                              `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeed struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeed) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObject],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeed].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObject) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedUnion() {
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectCondition struct {
	Argument param.Field[string]                                                                                   `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                   `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesPush struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPush) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPush) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObject],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesPush].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObject) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesPushUnion() {
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectCondition struct {
	Argument param.Field[string]                                                                              `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                              `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMS struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMS) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObject],
// [UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMS].
type UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion interface {
	ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion()
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObject) ImplementsUserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSUnion() {
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectCondition struct {
	Argument param.Field[string]                                                                             `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                             `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorContains             UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

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
)

func (r UserUpdateParamsPreferencesCategoriesObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

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

type UserUpdateParamsPreferencesChannelTypesChat struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesChannelTypesChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesChat) ImplementsUserUpdateParamsPreferencesChannelTypesChatUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesChatObject],
// [UserUpdateParamsPreferencesChannelTypesChat].
type UserUpdateParamsPreferencesChannelTypesChatUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesChatUnion()
}

type UserUpdateParamsPreferencesChannelTypesChatObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesChatObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesChatObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesChatObject) ImplementsUserUpdateParamsPreferencesChannelTypesChatUnion() {
}

type UserUpdateParamsPreferencesChannelTypesChatObjectCondition struct {
	Argument param.Field[string]                                                              `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                              `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesChatObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesChatObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesChannelTypesEmail struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesChannelTypesEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesEmail) ImplementsUserUpdateParamsPreferencesChannelTypesEmailUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesEmailObject],
// [UserUpdateParamsPreferencesChannelTypesEmail].
type UserUpdateParamsPreferencesChannelTypesEmailUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesEmailUnion()
}

type UserUpdateParamsPreferencesChannelTypesEmailObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesEmailObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesEmailObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesEmailObject) ImplementsUserUpdateParamsPreferencesChannelTypesEmailUnion() {
}

type UserUpdateParamsPreferencesChannelTypesEmailObjectCondition struct {
	Argument param.Field[string]                                                               `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                               `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesEmailObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesEmailObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesChannelTypesHTTP struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesChannelTypesHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesHTTP) ImplementsUserUpdateParamsPreferencesChannelTypesHTTPUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesHTTPObject],
// [UserUpdateParamsPreferencesChannelTypesHTTP].
type UserUpdateParamsPreferencesChannelTypesHTTPUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesHTTPUnion()
}

type UserUpdateParamsPreferencesChannelTypesHTTPObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesHTTPObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesHTTPObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesHTTPObject) ImplementsUserUpdateParamsPreferencesChannelTypesHTTPUnion() {
}

type UserUpdateParamsPreferencesChannelTypesHTTPObjectCondition struct {
	Argument param.Field[string]                                                              `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                              `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesHTTPObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesHTTPObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesChannelTypesInAppFeed struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesChannelTypesInAppFeed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesInAppFeed) ImplementsUserUpdateParamsPreferencesChannelTypesInAppFeedUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesInAppFeedObject],
// [UserUpdateParamsPreferencesChannelTypesInAppFeed].
type UserUpdateParamsPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesInAppFeedUnion()
}

type UserUpdateParamsPreferencesChannelTypesInAppFeedObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesInAppFeedObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesInAppFeedObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesInAppFeedObject) ImplementsUserUpdateParamsPreferencesChannelTypesInAppFeedUnion() {
}

type UserUpdateParamsPreferencesChannelTypesInAppFeedObjectCondition struct {
	Argument param.Field[string]                                                                   `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                   `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesInAppFeedObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesInAppFeedObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesChannelTypesPush struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesChannelTypesPush) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesPush) ImplementsUserUpdateParamsPreferencesChannelTypesPushUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesPushObject],
// [UserUpdateParamsPreferencesChannelTypesPush].
type UserUpdateParamsPreferencesChannelTypesPushUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesPushUnion()
}

type UserUpdateParamsPreferencesChannelTypesPushObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesPushObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesPushObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesPushObject) ImplementsUserUpdateParamsPreferencesChannelTypesPushUnion() {
}

type UserUpdateParamsPreferencesChannelTypesPushObjectCondition struct {
	Argument param.Field[string]                                                              `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                              `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesPushObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesPushObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesChannelTypesSMS struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesChannelTypesSMS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesSMS) ImplementsUserUpdateParamsPreferencesChannelTypesSMSUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesChannelTypesSMSObject],
// [UserUpdateParamsPreferencesChannelTypesSMS].
type UserUpdateParamsPreferencesChannelTypesSMSUnion interface {
	ImplementsUserUpdateParamsPreferencesChannelTypesSMSUnion()
}

type UserUpdateParamsPreferencesChannelTypesSMSObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesChannelTypesSMSObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesSMSObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesChannelTypesSMSObject) ImplementsUserUpdateParamsPreferencesChannelTypesSMSUnion() {
}

type UserUpdateParamsPreferencesChannelTypesSMSObjectCondition struct {
	Argument param.Field[string]                                                             `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                             `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesChannelTypesSMSObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorContains             UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesSMSObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesWorkflows struct {
	ChannelTypes param.Field[interface{}] `json:"channel_types,required"`
	Conditions   param.Field[interface{}] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflows) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflows) ImplementsUserUpdateParamsPreferencesWorkflowsUnion() {}

// Satisfied by [shared.UnionBool], [UserUpdateParamsPreferencesWorkflowsObject],
// [UserUpdateParamsPreferencesWorkflows].
type UserUpdateParamsPreferencesWorkflowsUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsUnion()
}

type UserUpdateParamsPreferencesWorkflowsObject struct {
	ChannelTypes param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectCondition]  `json:"conditions"`
}

func (r UserUpdateParamsPreferencesWorkflowsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObject) ImplementsUserUpdateParamsPreferencesWorkflowsUnion() {
}

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

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChat struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChat) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObject],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChat].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObject) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatUnion() {
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectCondition struct {
	Argument param.Field[string]                                                                             `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                             `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmail struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmail) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObject],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmail].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObject) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailUnion() {
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectCondition struct {
	Argument param.Field[string]                                                                              `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                              `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTP struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTP) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObject],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTP].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObject) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPUnion() {
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectCondition struct {
	Argument param.Field[string]                                                                             `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                             `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeed struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeed) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObject],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeed].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObject) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedUnion() {
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectCondition struct {
	Argument param.Field[string]                                                                                  `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                  `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPush struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPush) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPush) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObject],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPush].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObject) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushUnion() {
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectCondition struct {
	Argument param.Field[string]                                                                             `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                             `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMS struct {
	Conditions param.Field[interface{}] `json:"conditions"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMS) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSUnion() {
}

// Satisfied by [shared.UnionBool],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObject],
// [UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMS].
type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSUnion interface {
	ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSUnion()
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObject struct {
	Conditions param.Field[[]UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectCondition] `json:"conditions,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObject) ImplementsUserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSUnion() {
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectCondition struct {
	Argument param.Field[string]                                                                            `json:"argument,required"`
	Operator param.Field[UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                            `json:"variable,required"`
}

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator string

const (
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorEqualTo              UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorNotEqualTo           UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "not_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorGreaterThan          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "greater_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorLessThan             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "less_than"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorGreaterThanOrEqualTo UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "greater_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorLessThanOrEqualTo    UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "less_than_or_equal_to"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorContains             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorNotContains          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "not_contains"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorEmpty                UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorNotEmpty             UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "not_empty"
	UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorContainsAll          UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator = "contains_all"
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSObjectConditionsOperatorContainsAll:
		return true
	}
	return false
}

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
)

func (r UserUpdateParamsPreferencesWorkflowsObjectConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectConditionsOperatorContainsAll:
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

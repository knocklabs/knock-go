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
	Typename    User_Typename     `json:"__typename"`
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

type User_Typename string

const (
	User_TypenameUser User_Typename = "User"
)

func (r User_Typename) IsKnown() bool {
	switch r {
	case User_TypenameUser:
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

// Satisfied by [shared.UnionBool], [UserUpdateParamsPreferencesCategoriesObject].
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
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesChatConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesEmailConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesPushConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesCategoriesObjectChannelTypesSMSConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesChatConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesEmailConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesHTTPConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesInAppFeedConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesPushConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContains, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesChannelTypesSMSConditionsConditionsOperatorContainsAll:
		return true
	}
	return false
}

// Satisfied by [shared.UnionBool], [UserUpdateParamsPreferencesWorkflowsObject].
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
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesChatConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesEmailConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesHTTPConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesInAppFeedConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesPushConditionsConditionsOperatorContainsAll:
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
)

func (r UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperator) IsKnown() bool {
	switch r {
	case UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThan, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorGreaterThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorLessThanOrEqualTo, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotContains, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorNotEmpty, UserUpdateParamsPreferencesWorkflowsObjectChannelTypesSMSConditionsConditionsOperatorContainsAll:
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

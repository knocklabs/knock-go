// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"slices"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/shared"
)

// A bulk operation is a set of changes applied across zero or more records
// triggered via a call to the Knock API and performed asynchronously.
//
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

// Permanently deletes up to 1,000 users at a time.
func (r *UserBulkService) Delete(ctx context.Context, body UserBulkDeleteParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/bulk/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Identifies multiple users in a single operation. Allows creating or updating up
// to 1,000 users in a single batch with various properties, preferences, and
// channel data.
func (r *UserBulkService) Identify(ctx context.Context, body UserBulkIdentifyParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/bulk/identify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Bulk sets the preferences for up to 1,000 users at a time. The preference
// set `:id` can be either `default` or a `tenant.id`. Learn more
// about [per-tenant preferences](/preferences/tenant-preferences). Note that this
// is a destructive operation and will replace any existing users' preferences with
// the preferences sent.
func (r *UserBulkService) SetPreferences(ctx context.Context, body UserBulkSetPreferencesParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/bulk/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type UserBulkDeleteParams struct {
	// A list of user IDs.
	UserIDs param.Field[[]string] `json:"user_ids" api:"required"`
}

func (r UserBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParams struct {
	// A list of users.
	Users param.Field[[]InlineIdentifyUserRequestParam] `json:"users" api:"required"`
}

func (r UserBulkIdentifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParams struct {
	// A preference set to apply in a bulk operation. Always replaces existing
	// preferences for the specified set.
	Preferences param.Field[UserBulkSetPreferencesParamsPreferences] `json:"preferences" api:"required"`
	// A list of user IDs.
	UserIDs param.Field[[]string] `json:"user_ids" api:"required"`
}

func (r UserBulkSetPreferencesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A preference set to apply in a bulk operation. Always replaces existing
// preferences for the specified set.
type UserBulkSetPreferencesParamsPreferences struct {
	// Identifier for the preference set to update. Can be `default` or a tenant ID.
	ID param.Field[string] `json:"id"`
	// An object where the key is the category and the values are the preference
	// settings for that category.
	Categories param.Field[map[string]UserBulkSetPreferencesParamsPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences.
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// Channel preferences.
	Channels param.Field[map[string]UserBulkSetPreferencesParamsPreferencesChannelsUnion] `json:"channels"`
	// Whether the recipient is subscribed to commercial communications. When false,
	// the recipient will not receive commercial workflow notifications.
	CommercialSubscribed param.Field[bool] `json:"commercial_subscribed"`
	// An object where the key is the workflow key and the values are the preference
	// settings for that workflow.
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
	// Channel type preferences.
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// Channel preferences.
	Channels param.Field[map[string]UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelsUnion] `json:"channels"`
	// A list of conditions to apply to a channel type.
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions"`
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesUnion() {
}

// Whether the specific channel (by channel_id) is enabled for the preference set,
// or a settings object with conditions.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelSettingParam].
type UserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelsUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelsUnion()
}

// Whether the specific channel (by channel_id) is enabled for the preference set,
// or a settings object with conditions.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelSettingParam].
type UserBulkSetPreferencesParamsPreferencesChannelsUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesChannelsUnion()
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
	// Channel type preferences.
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// Channel preferences.
	Channels param.Field[map[string]UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelsUnion] `json:"channels"`
	// A list of conditions to apply to a channel type.
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions"`
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsUnion() {
}

// Whether the specific channel (by channel_id) is enabled for the preference set,
// or a settings object with conditions.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelSettingParam].
type UserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelsUnion interface {
	ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelsUnion()
}

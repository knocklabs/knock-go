// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"reflect"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
	"github.com/tidwall/gjson"
)

// RecipientPreferenceService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientPreferenceService] method instead.
type RecipientPreferenceService struct {
	Options []option.RequestOption
}

// NewRecipientPreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRecipientPreferenceService(opts ...option.RequestOption) (r *RecipientPreferenceService) {
	r = &RecipientPreferenceService{}
	r.Options = opts
	return
}

type InlinePreferenceSetRequestParam struct {
	// Unique identifier for the preference set.
	ID param.Field[string] `json:"id,required"`
	// An object where the key is the category and the values are the preference
	// settings for that category.
	Categories param.Field[map[string]InlinePreferenceSetRequestCategoriesUnionParam] `json:"categories"`
	// Channel type preferences.
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// An object where the key is the workflow key and the values are the preference
	// settings for that workflow.
	Workflows param.Field[map[string]InlinePreferenceSetRequestWorkflowsUnionParam] `json:"workflows"`
}

func (r InlinePreferenceSetRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [InlinePreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam].
type InlinePreferenceSetRequestCategoriesUnionParam interface {
	ImplementsInlinePreferenceSetRequestCategoriesUnionParam()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type InlinePreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam struct {
	// Channel type preferences.
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// A list of conditions to apply to a channel type.
	Conditions param.Field[[]ConditionParam] `json:"conditions"`
}

func (r InlinePreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InlinePreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam) ImplementsInlinePreferenceSetRequestCategoriesUnionParam() {
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [InlinePreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam].
type InlinePreferenceSetRequestWorkflowsUnionParam interface {
	ImplementsInlinePreferenceSetRequestWorkflowsUnionParam()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type InlinePreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam struct {
	// Channel type preferences.
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// A list of conditions to apply to a channel type.
	Conditions param.Field[[]ConditionParam] `json:"conditions"`
}

func (r InlinePreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InlinePreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam) ImplementsInlinePreferenceSetRequestWorkflowsUnionParam() {
}

// A preference set represents a specific set of notification preferences for a
// recipient. A recipient can have multiple preference sets.
type PreferenceSet struct {
	// Unique identifier for the preference set.
	ID string `json:"id,required"`
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// An object where the key is the category and the values are the preference
	// settings for that category.
	Categories map[string]PreferenceSetCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences.
	ChannelTypes PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// An object where the key is the workflow key and the values are the preference
	// settings for that workflow.
	Workflows map[string]PreferenceSetWorkflowsUnion `json:"workflows,nullable"`
	JSON      preferenceSetJSON                      `json:"-"`
}

// preferenceSetJSON contains the JSON metadata for the struct [PreferenceSet]
type preferenceSetJSON struct {
	ID           apijson.Field
	Typename     apijson.Field
	Categories   apijson.Field
	ChannelTypes apijson.Field
	Workflows    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PreferenceSet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetJSON) RawJSON() string {
	return r.raw
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [PreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject].
type PreferenceSetCategoriesUnion interface {
	ImplementsPreferenceSetCategoriesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetCategoriesUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(PreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type PreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences.
	ChannelTypes PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// A list of conditions to apply to a channel type.
	Conditions []Condition                                                           `json:"conditions,nullable"`
	JSON       preferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// preferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectJSON contains
// the JSON metadata for the struct
// [PreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject]
type preferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsPreferenceSetCategoriesUnion() {
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [PreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject].
type PreferenceSetWorkflowsUnion interface {
	ImplementsPreferenceSetWorkflowsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetWorkflowsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(PreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type PreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences.
	ChannelTypes PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// A list of conditions to apply to a channel type.
	Conditions []Condition                                                          `json:"conditions,nullable"`
	JSON       preferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// preferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON contains
// the JSON metadata for the struct
// [PreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject]
type preferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsPreferenceSetWorkflowsUnion() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypeSetting struct {
	// A list of conditions to apply to a channel type.
	Conditions []Condition                         `json:"conditions,required"`
	JSON       preferenceSetChannelTypeSettingJSON `json:"-"`
}

// preferenceSetChannelTypeSettingJSON contains the JSON metadata for the struct
// [PreferenceSetChannelTypeSetting]
type preferenceSetChannelTypeSettingJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypeSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypeSettingJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypeSetting) ImplementsPreferenceSetChannelTypesChatUnion() {}

func (r PreferenceSetChannelTypeSetting) ImplementsPreferenceSetChannelTypesEmailUnion() {}

func (r PreferenceSetChannelTypeSetting) ImplementsPreferenceSetChannelTypesHTTPUnion() {}

func (r PreferenceSetChannelTypeSetting) ImplementsPreferenceSetChannelTypesInAppFeedUnion() {}

func (r PreferenceSetChannelTypeSetting) ImplementsPreferenceSetChannelTypesPushUnion() {}

func (r PreferenceSetChannelTypeSetting) ImplementsPreferenceSetChannelTypesSMSUnion() {}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypeSettingParam struct {
	// A list of conditions to apply to a channel type.
	Conditions param.Field[[]ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypeSettingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypeSettingParam) ImplementsPreferenceSetChannelTypesChatUnionParam() {}

func (r PreferenceSetChannelTypeSettingParam) ImplementsPreferenceSetChannelTypesEmailUnionParam() {}

func (r PreferenceSetChannelTypeSettingParam) ImplementsPreferenceSetChannelTypesHTTPUnionParam() {}

func (r PreferenceSetChannelTypeSettingParam) ImplementsPreferenceSetChannelTypesInAppFeedUnionParam() {
}

func (r PreferenceSetChannelTypeSettingParam) ImplementsPreferenceSetChannelTypesPushUnionParam() {}

func (r PreferenceSetChannelTypeSettingParam) ImplementsPreferenceSetChannelTypesSMSUnionParam() {}

// Channel type preferences.
type PreferenceSetChannelTypes struct {
	// Either a boolean or a setting for the given channel type.
	Chat PreferenceSetChannelTypesChatUnion `json:"chat"`
	// Either a boolean or a setting for the given channel type.
	Email PreferenceSetChannelTypesEmailUnion `json:"email"`
	// Either a boolean or a setting for the given channel type.
	HTTP PreferenceSetChannelTypesHTTPUnion `json:"http"`
	// Either a boolean or a setting for the given channel type.
	InAppFeed PreferenceSetChannelTypesInAppFeedUnion `json:"in_app_feed"`
	// Either a boolean or a setting for the given channel type.
	Push PreferenceSetChannelTypesPushUnion `json:"push"`
	// Either a boolean or a setting for the given channel type.
	SMS  PreferenceSetChannelTypesSMSUnion `json:"sms"`
	JSON preferenceSetChannelTypesJSON     `json:"-"`
}

// preferenceSetChannelTypesJSON contains the JSON metadata for the struct
// [PreferenceSetChannelTypes]
type preferenceSetChannelTypesJSON struct {
	Chat        apijson.Field
	Email       apijson.Field
	HTTP        apijson.Field
	InAppFeed   apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesJSON) RawJSON() string {
	return r.raw
}

// Either a boolean or a setting for the given channel type.
//
// Union satisfied by [shared.UnionBool] or [PreferenceSetChannelTypeSetting].
type PreferenceSetChannelTypesChatUnion interface {
	ImplementsPreferenceSetChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesChatUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypeSetting{}),
		},
	)
}

// Either a boolean or a setting for the given channel type.
//
// Union satisfied by [shared.UnionBool] or [PreferenceSetChannelTypeSetting].
type PreferenceSetChannelTypesEmailUnion interface {
	ImplementsPreferenceSetChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesEmailUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypeSetting{}),
		},
	)
}

// Either a boolean or a setting for the given channel type.
//
// Union satisfied by [shared.UnionBool] or [PreferenceSetChannelTypeSetting].
type PreferenceSetChannelTypesHTTPUnion interface {
	ImplementsPreferenceSetChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesHTTPUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypeSetting{}),
		},
	)
}

// Either a boolean or a setting for the given channel type.
//
// Union satisfied by [shared.UnionBool] or [PreferenceSetChannelTypeSetting].
type PreferenceSetChannelTypesInAppFeedUnion interface {
	ImplementsPreferenceSetChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesInAppFeedUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypeSetting{}),
		},
	)
}

// Either a boolean or a setting for the given channel type.
//
// Union satisfied by [shared.UnionBool] or [PreferenceSetChannelTypeSetting].
type PreferenceSetChannelTypesPushUnion interface {
	ImplementsPreferenceSetChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesPushUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypeSetting{}),
		},
	)
}

// Either a boolean or a setting for the given channel type.
//
// Union satisfied by [shared.UnionBool] or [PreferenceSetChannelTypeSetting].
type PreferenceSetChannelTypesSMSUnion interface {
	ImplementsPreferenceSetChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesSMSUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypeSetting{}),
		},
	)
}

// Channel type preferences.
type PreferenceSetChannelTypesParam struct {
	// Either a boolean or a setting for the given channel type.
	Chat param.Field[PreferenceSetChannelTypesChatUnionParam] `json:"chat"`
	// Either a boolean or a setting for the given channel type.
	Email param.Field[PreferenceSetChannelTypesEmailUnionParam] `json:"email"`
	// Either a boolean or a setting for the given channel type.
	HTTP param.Field[PreferenceSetChannelTypesHTTPUnionParam] `json:"http"`
	// Either a boolean or a setting for the given channel type.
	InAppFeed param.Field[PreferenceSetChannelTypesInAppFeedUnionParam] `json:"in_app_feed"`
	// Either a boolean or a setting for the given channel type.
	Push param.Field[PreferenceSetChannelTypesPushUnionParam] `json:"push"`
	// Either a boolean or a setting for the given channel type.
	SMS param.Field[PreferenceSetChannelTypesSMSUnionParam] `json:"sms"`
}

func (r PreferenceSetChannelTypesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either a boolean or a setting for the given channel type.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesChatUnionParam interface {
	ImplementsPreferenceSetChannelTypesChatUnionParam()
}

// Either a boolean or a setting for the given channel type.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesEmailUnionParam interface {
	ImplementsPreferenceSetChannelTypesEmailUnionParam()
}

// Either a boolean or a setting for the given channel type.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesHTTPUnionParam interface {
	ImplementsPreferenceSetChannelTypesHTTPUnionParam()
}

// Either a boolean or a setting for the given channel type.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesInAppFeedUnionParam interface {
	ImplementsPreferenceSetChannelTypesInAppFeedUnionParam()
}

// Either a boolean or a setting for the given channel type.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesPushUnionParam interface {
	ImplementsPreferenceSetChannelTypesPushUnionParam()
}

// Either a boolean or a setting for the given channel type.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesSMSUnionParam interface {
	ImplementsPreferenceSetChannelTypesSMSUnionParam()
}

// A request to set a preference set for a recipient.
type PreferenceSetRequestParam struct {
	// An object where the key is the category and the values are the preference
	// settings for that category.
	Categories param.Field[map[string]PreferenceSetRequestCategoriesUnionParam] `json:"categories"`
	// Channel type preferences.
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// An object where the key is the workflow key and the values are the preference
	// settings for that workflow.
	Workflows param.Field[map[string]PreferenceSetRequestWorkflowsUnionParam] `json:"workflows"`
}

func (r PreferenceSetRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam].
type PreferenceSetRequestCategoriesUnionParam interface {
	ImplementsPreferenceSetRequestCategoriesUnionParam()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam struct {
	// Channel type preferences.
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// A list of conditions to apply to a channel type.
	Conditions param.Field[[]ConditionParam] `json:"conditions"`
}

func (r PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam) ImplementsPreferenceSetRequestCategoriesUnionParam() {
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam].
type PreferenceSetRequestWorkflowsUnionParam interface {
	ImplementsPreferenceSetRequestWorkflowsUnionParam()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam struct {
	// Channel type preferences.
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// A list of conditions to apply to a channel type.
	Conditions param.Field[[]ConditionParam] `json:"conditions"`
}

func (r PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam) ImplementsPreferenceSetRequestWorkflowsUnionParam() {
}

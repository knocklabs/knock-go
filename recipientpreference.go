// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"reflect"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/shared"
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

type InlinePreferenceSetRequestParam map[string]PreferenceSetRequestParam

// A preference set represents a specific set of notification preferences for a
// recipient. A recipient can have multiple preference sets.
type PreferenceSet struct {
	// Unique identifier for the preference set.
	ID string `json:"id,required"`
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
	Conditions []shared.Condition                                                    `json:"conditions,nullable"`
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
	Conditions []shared.Condition                                                   `json:"conditions,nullable"`
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
	Conditions []shared.Condition                  `json:"conditions,required"`
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
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions,required"`
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
	// Whether the channel type is enabled for the preference set.
	Chat PreferenceSetChannelTypesChatUnion `json:"chat"`
	// Whether the channel type is enabled for the preference set.
	Email PreferenceSetChannelTypesEmailUnion `json:"email"`
	// Whether the channel type is enabled for the preference set.
	HTTP PreferenceSetChannelTypesHTTPUnion `json:"http"`
	// Whether the channel type is enabled for the preference set.
	InAppFeed PreferenceSetChannelTypesInAppFeedUnion `json:"in_app_feed"`
	// Whether the channel type is enabled for the preference set.
	Push PreferenceSetChannelTypesPushUnion `json:"push"`
	// Whether the channel type is enabled for the preference set.
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

// Whether the channel type is enabled for the preference set.
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

// Whether the channel type is enabled for the preference set.
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

// Whether the channel type is enabled for the preference set.
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

// Whether the channel type is enabled for the preference set.
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

// Whether the channel type is enabled for the preference set.
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

// Whether the channel type is enabled for the preference set.
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
	// Whether the channel type is enabled for the preference set.
	Chat param.Field[PreferenceSetChannelTypesChatUnionParam] `json:"chat"`
	// Whether the channel type is enabled for the preference set.
	Email param.Field[PreferenceSetChannelTypesEmailUnionParam] `json:"email"`
	// Whether the channel type is enabled for the preference set.
	HTTP param.Field[PreferenceSetChannelTypesHTTPUnionParam] `json:"http"`
	// Whether the channel type is enabled for the preference set.
	InAppFeed param.Field[PreferenceSetChannelTypesInAppFeedUnionParam] `json:"in_app_feed"`
	// Whether the channel type is enabled for the preference set.
	Push param.Field[PreferenceSetChannelTypesPushUnionParam] `json:"push"`
	// Whether the channel type is enabled for the preference set.
	SMS param.Field[PreferenceSetChannelTypesSMSUnionParam] `json:"sms"`
}

func (r PreferenceSetChannelTypesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the channel type is enabled for the preference set.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesChatUnionParam interface {
	ImplementsPreferenceSetChannelTypesChatUnionParam()
}

// Whether the channel type is enabled for the preference set.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesEmailUnionParam interface {
	ImplementsPreferenceSetChannelTypesEmailUnionParam()
}

// Whether the channel type is enabled for the preference set.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesHTTPUnionParam interface {
	ImplementsPreferenceSetChannelTypesHTTPUnionParam()
}

// Whether the channel type is enabled for the preference set.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesInAppFeedUnionParam interface {
	ImplementsPreferenceSetChannelTypesInAppFeedUnionParam()
}

// Whether the channel type is enabled for the preference set.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesPushUnionParam interface {
	ImplementsPreferenceSetChannelTypesPushUnionParam()
}

// Whether the channel type is enabled for the preference set.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesSMSUnionParam interface {
	ImplementsPreferenceSetChannelTypesSMSUnionParam()
}

// A request to set a preference set for a recipient.
type PreferenceSetRequestParam struct {
	// Controls how the preference set is persisted. 'replace' will completely replace
	// the preference set, 'merge' will merge with existing preferences.
	PersistenceStrategy param.Field[PreferenceSetRequest_PersistenceStrategy] `json:"__persistence_strategy__"`
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

// Controls how the preference set is persisted. 'replace' will completely replace
// the preference set, 'merge' will merge with existing preferences.
type PreferenceSetRequest_PersistenceStrategy string

const (
	PreferenceSetRequest_PersistenceStrategyMerge   PreferenceSetRequest_PersistenceStrategy = "merge"
	PreferenceSetRequest_PersistenceStrategyReplace PreferenceSetRequest_PersistenceStrategy = "replace"
)

func (r PreferenceSetRequest_PersistenceStrategy) IsKnown() bool {
	switch r {
	case PreferenceSetRequest_PersistenceStrategyMerge, PreferenceSetRequest_PersistenceStrategyReplace:
		return true
	}
	return false
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
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions"`
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
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions"`
}

func (r PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam) ImplementsPreferenceSetRequestWorkflowsUnionParam() {
}

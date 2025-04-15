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

type InlinePreferenceSetRequestParam map[string]PreferenceSetRequestParam

// A preference set object.
type PreferenceSet struct {
	ID         string                                  `json:"id,required"`
	Typename   string                                  `json:"__typename,required"`
	Categories map[string]PreferenceSetCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes PreferenceSetChannelTypes              `json:"channel_types,nullable"`
	Workflows    map[string]PreferenceSetWorkflowsUnion `json:"workflows,nullable"`
	JSON         preferenceSetJSON                      `json:"-"`
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
// Union satisfied by [shared.UnionBool] or [PreferenceSetCategoriesObject].
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
			Type:       reflect.TypeOf(PreferenceSetCategoriesObject{}),
		},
	)
}

type PreferenceSetCategoriesObject struct {
	// Channel type preferences
	ChannelTypes PreferenceSetChannelTypes         `json:"channel_types,nullable"`
	Conditions   []shared.Condition                `json:"conditions"`
	JSON         preferenceSetCategoriesObjectJSON `json:"-"`
}

// preferenceSetCategoriesObjectJSON contains the JSON metadata for the struct
// [PreferenceSetCategoriesObject]
type preferenceSetCategoriesObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PreferenceSetCategoriesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetCategoriesObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetCategoriesObject) ImplementsPreferenceSetCategoriesUnion() {}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or [PreferenceSetWorkflowsObject].
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
			Type:       reflect.TypeOf(PreferenceSetWorkflowsObject{}),
		},
	)
}

type PreferenceSetWorkflowsObject struct {
	// Channel type preferences
	ChannelTypes PreferenceSetChannelTypes        `json:"channel_types,nullable"`
	Conditions   []shared.Condition               `json:"conditions"`
	JSON         preferenceSetWorkflowsObjectJSON `json:"-"`
}

// preferenceSetWorkflowsObjectJSON contains the JSON metadata for the struct
// [PreferenceSetWorkflowsObject]
type preferenceSetWorkflowsObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PreferenceSetWorkflowsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetWorkflowsObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetWorkflowsObject) ImplementsPreferenceSetWorkflowsUnion() {}

// Channel type preferences
type PreferenceSetChannelTypes struct {
	Chat      PreferenceSetChannelTypesChatUnion      `json:"chat"`
	Email     PreferenceSetChannelTypesEmailUnion     `json:"email"`
	HTTP      PreferenceSetChannelTypesHTTPUnion      `json:"http"`
	InAppFeed PreferenceSetChannelTypesInAppFeedUnion `json:"in_app_feed"`
	Push      PreferenceSetChannelTypesPushUnion      `json:"push"`
	SMS       PreferenceSetChannelTypesSMSUnion       `json:"sms"`
	JSON      preferenceSetChannelTypesJSON           `json:"-"`
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

// Union satisfied by [shared.UnionBool] or
// [PreferenceSetChannelTypesChatConditions].
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypesChatConditions{}),
		},
	)
}

type PreferenceSetChannelTypesChatConditions struct {
	Conditions []shared.Condition                          `json:"conditions,required"`
	JSON       preferenceSetChannelTypesChatConditionsJSON `json:"-"`
}

// preferenceSetChannelTypesChatConditionsJSON contains the JSON metadata for the
// struct [PreferenceSetChannelTypesChatConditions]
type preferenceSetChannelTypesChatConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesChatConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesChatConditionsJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesChatConditions) ImplementsPreferenceSetChannelTypesChatUnion() {}

// Union satisfied by [shared.UnionBool] or
// [PreferenceSetChannelTypesEmailConditions].
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypesEmailConditions{}),
		},
	)
}

type PreferenceSetChannelTypesEmailConditions struct {
	Conditions []shared.Condition                           `json:"conditions,required"`
	JSON       preferenceSetChannelTypesEmailConditionsJSON `json:"-"`
}

// preferenceSetChannelTypesEmailConditionsJSON contains the JSON metadata for the
// struct [PreferenceSetChannelTypesEmailConditions]
type preferenceSetChannelTypesEmailConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesEmailConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesEmailConditionsJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesEmailConditions) ImplementsPreferenceSetChannelTypesEmailUnion() {}

// Union satisfied by [shared.UnionBool] or
// [PreferenceSetChannelTypesHTTPConditions].
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypesHTTPConditions{}),
		},
	)
}

type PreferenceSetChannelTypesHTTPConditions struct {
	Conditions []shared.Condition                          `json:"conditions,required"`
	JSON       preferenceSetChannelTypesHTTPConditionsJSON `json:"-"`
}

// preferenceSetChannelTypesHTTPConditionsJSON contains the JSON metadata for the
// struct [PreferenceSetChannelTypesHTTPConditions]
type preferenceSetChannelTypesHTTPConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesHTTPConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesHTTPConditionsJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesHTTPConditions) ImplementsPreferenceSetChannelTypesHTTPUnion() {}

// Union satisfied by [shared.UnionBool] or
// [PreferenceSetChannelTypesInAppFeedConditions].
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypesInAppFeedConditions{}),
		},
	)
}

type PreferenceSetChannelTypesInAppFeedConditions struct {
	Conditions []shared.Condition                               `json:"conditions,required"`
	JSON       preferenceSetChannelTypesInAppFeedConditionsJSON `json:"-"`
}

// preferenceSetChannelTypesInAppFeedConditionsJSON contains the JSON metadata for
// the struct [PreferenceSetChannelTypesInAppFeedConditions]
type preferenceSetChannelTypesInAppFeedConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesInAppFeedConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesInAppFeedConditionsJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesInAppFeedConditions) ImplementsPreferenceSetChannelTypesInAppFeedUnion() {
}

// Union satisfied by [shared.UnionBool] or
// [PreferenceSetChannelTypesPushConditions].
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypesPushConditions{}),
		},
	)
}

type PreferenceSetChannelTypesPushConditions struct {
	Conditions []shared.Condition                          `json:"conditions,required"`
	JSON       preferenceSetChannelTypesPushConditionsJSON `json:"-"`
}

// preferenceSetChannelTypesPushConditionsJSON contains the JSON metadata for the
// struct [PreferenceSetChannelTypesPushConditions]
type preferenceSetChannelTypesPushConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesPushConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesPushConditionsJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesPushConditions) ImplementsPreferenceSetChannelTypesPushUnion() {}

// Union satisfied by [shared.UnionBool] or
// [PreferenceSetChannelTypesSMSConditions].
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
			Type:       reflect.TypeOf(PreferenceSetChannelTypesSMSConditions{}),
		},
	)
}

type PreferenceSetChannelTypesSMSConditions struct {
	Conditions []shared.Condition                         `json:"conditions,required"`
	JSON       preferenceSetChannelTypesSMSConditionsJSON `json:"-"`
}

// preferenceSetChannelTypesSMSConditionsJSON contains the JSON metadata for the
// struct [PreferenceSetChannelTypesSMSConditions]
type preferenceSetChannelTypesSMSConditionsJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesSMSConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesSMSConditionsJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesSMSConditions) ImplementsPreferenceSetChannelTypesSMSUnion() {}

// Channel type preferences
type PreferenceSetChannelTypesParam struct {
	Chat      param.Field[PreferenceSetChannelTypesChatUnionParam]      `json:"chat"`
	Email     param.Field[PreferenceSetChannelTypesEmailUnionParam]     `json:"email"`
	HTTP      param.Field[PreferenceSetChannelTypesHTTPUnionParam]      `json:"http"`
	InAppFeed param.Field[PreferenceSetChannelTypesInAppFeedUnionParam] `json:"in_app_feed"`
	Push      param.Field[PreferenceSetChannelTypesPushUnionParam]      `json:"push"`
	SMS       param.Field[PreferenceSetChannelTypesSMSUnionParam]       `json:"sms"`
}

func (r PreferenceSetChannelTypesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypesChatConditionsParam].
type PreferenceSetChannelTypesChatUnionParam interface {
	ImplementsPreferenceSetChannelTypesChatUnionParam()
}

type PreferenceSetChannelTypesChatConditionsParam struct {
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesChatConditionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesChatConditionsParam) ImplementsPreferenceSetChannelTypesChatUnionParam() {
}

// Satisfied by [shared.UnionBool],
// [PreferenceSetChannelTypesEmailConditionsParam].
type PreferenceSetChannelTypesEmailUnionParam interface {
	ImplementsPreferenceSetChannelTypesEmailUnionParam()
}

type PreferenceSetChannelTypesEmailConditionsParam struct {
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesEmailConditionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesEmailConditionsParam) ImplementsPreferenceSetChannelTypesEmailUnionParam() {
}

// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypesHTTPConditionsParam].
type PreferenceSetChannelTypesHTTPUnionParam interface {
	ImplementsPreferenceSetChannelTypesHTTPUnionParam()
}

type PreferenceSetChannelTypesHTTPConditionsParam struct {
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesHTTPConditionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesHTTPConditionsParam) ImplementsPreferenceSetChannelTypesHTTPUnionParam() {
}

// Satisfied by [shared.UnionBool],
// [PreferenceSetChannelTypesInAppFeedConditionsParam].
type PreferenceSetChannelTypesInAppFeedUnionParam interface {
	ImplementsPreferenceSetChannelTypesInAppFeedUnionParam()
}

type PreferenceSetChannelTypesInAppFeedConditionsParam struct {
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesInAppFeedConditionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesInAppFeedConditionsParam) ImplementsPreferenceSetChannelTypesInAppFeedUnionParam() {
}

// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypesPushConditionsParam].
type PreferenceSetChannelTypesPushUnionParam interface {
	ImplementsPreferenceSetChannelTypesPushUnionParam()
}

type PreferenceSetChannelTypesPushConditionsParam struct {
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesPushConditionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesPushConditionsParam) ImplementsPreferenceSetChannelTypesPushUnionParam() {
}

// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypesSMSConditionsParam].
type PreferenceSetChannelTypesSMSUnionParam interface {
	ImplementsPreferenceSetChannelTypesSMSUnionParam()
}

type PreferenceSetChannelTypesSMSConditionsParam struct {
	Conditions param.Field[[]shared.ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesSMSConditionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesSMSConditionsParam) ImplementsPreferenceSetChannelTypesSMSUnionParam() {
}

// Set preferences for a recipient
type PreferenceSetRequestParam struct {
	Categories param.Field[map[string]PreferenceSetRequestCategoriesUnionParam] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[PreferenceSetChannelTypesParam]                     `json:"channel_types"`
	Workflows    param.Field[map[string]PreferenceSetRequestWorkflowsUnionParam] `json:"workflows"`
}

func (r PreferenceSetRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool], [PreferenceSetRequestCategoriesObjectParam].
type PreferenceSetRequestCategoriesUnionParam interface {
	ImplementsPreferenceSetRequestCategoriesUnionParam()
}

type PreferenceSetRequestCategoriesObjectParam struct {
	// Channel type preferences
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	Conditions   param.Field[[]shared.ConditionParam]        `json:"conditions"`
}

func (r PreferenceSetRequestCategoriesObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetRequestCategoriesObjectParam) ImplementsPreferenceSetRequestCategoriesUnionParam() {
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool], [PreferenceSetRequestWorkflowsObjectParam].
type PreferenceSetRequestWorkflowsUnionParam interface {
	ImplementsPreferenceSetRequestWorkflowsUnionParam()
}

type PreferenceSetRequestWorkflowsObjectParam struct {
	// Channel type preferences
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	Conditions   param.Field[[]shared.ConditionParam]        `json:"conditions"`
}

func (r PreferenceSetRequestWorkflowsObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetRequestWorkflowsObjectParam) ImplementsPreferenceSetRequestWorkflowsUnionParam() {
}

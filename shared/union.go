// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsRecipientReferenceUnionParam()        {}
func (UnionString) ImplementsRecipientReferenceUnion()             {}
func (UnionString) ImplementsRecipientRequestUnionParam()          {}
func (UnionString) ImplementsInlineTenantRequestUnionParam()       {}
func (UnionString) ImplementsMessageDeliveryLogRequestBodyUnion()  {}
func (UnionString) ImplementsMessageDeliveryLogResponseBodyUnion() {}

type UnionBool bool

func (UnionBool) ImplementsPreferenceSetCategoriesUnion() {}
func (UnionBool) ImplementsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelsUnion() {
}
func (UnionBool) ImplementsPreferenceSetChannelsUnion()  {}
func (UnionBool) ImplementsPreferenceSetWorkflowsUnion() {}
func (UnionBool) ImplementsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelsUnion() {
}
func (UnionBool) ImplementsPreferenceSetChannelTypesChatUnionParam()      {}
func (UnionBool) ImplementsPreferenceSetChannelTypesChatUnion()           {}
func (UnionBool) ImplementsPreferenceSetChannelTypesEmailUnionParam()     {}
func (UnionBool) ImplementsPreferenceSetChannelTypesEmailUnion()          {}
func (UnionBool) ImplementsPreferenceSetChannelTypesHTTPUnionParam()      {}
func (UnionBool) ImplementsPreferenceSetChannelTypesHTTPUnion()           {}
func (UnionBool) ImplementsPreferenceSetChannelTypesInAppFeedUnionParam() {}
func (UnionBool) ImplementsPreferenceSetChannelTypesInAppFeedUnion()      {}
func (UnionBool) ImplementsPreferenceSetChannelTypesPushUnionParam()      {}
func (UnionBool) ImplementsPreferenceSetChannelTypesPushUnion()           {}
func (UnionBool) ImplementsPreferenceSetChannelTypesSMSUnionParam()       {}
func (UnionBool) ImplementsPreferenceSetChannelTypesSMSUnion()            {}
func (UnionBool) ImplementsPreferenceSetRequestCategoriesUnionParam()     {}
func (UnionBool) ImplementsPreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectChannelsUnionParam() {
}
func (UnionBool) ImplementsPreferenceSetRequestChannelsUnionParam()  {}
func (UnionBool) ImplementsPreferenceSetRequestWorkflowsUnionParam() {}
func (UnionBool) ImplementsPreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelsUnionParam() {
}

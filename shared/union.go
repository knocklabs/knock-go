// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsUserListSubscriptionsParamsObjectUnion()                      {}
func (UnionString) ImplementsObjectAddSubscriptionsParamsRecipientUnion()                  {}
func (UnionString) ImplementsObjectDeleteSubscriptionsParamsRecipientUnion()               {}
func (UnionString) ImplementsObjectListSubscriptionsParamsObjectUnion()                    {}
func (UnionString) ImplementsObjectListSubscriptionsParamsRecipientUnion()                 {}
func (UnionString) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion() {}
func (UnionString) ImplementsInlineTenantRequestUnionParam()                               {}
func (UnionString) ImplementsMessageActorsUnion()                                          {}
func (UnionString) ImplementsMessageRecipientUnion()                                       {}
func (UnionString) ImplementsMessageDeliveryLogRequestBodyUnion()                          {}
func (UnionString) ImplementsMessageDeliveryLogResponseBodyUnion()                         {}
func (UnionString) ImplementsMessageEventRecipientUnion()                                  {}
func (UnionString) ImplementsWorkflowCancelParamsRecipientUnion()                          {}
func (UnionString) ImplementsWorkflowTriggerParamsRecipientUnion()                         {}
func (UnionString) ImplementsWorkflowTriggerParamsActorUnion()                             {}
func (UnionString) ImplementsScheduleNewParamsRecipientUnion()                             {}
func (UnionString) ImplementsScheduleUpdateParamsActorUnion()                              {}

type UnionBool bool

func (UnionBool) ImplementsInlinePreferenceSetRequestItemCategoriesUnionParam() {}
func (UnionBool) ImplementsInlinePreferenceSetRequestItemWorkflowsUnionParam()  {}
func (UnionBool) ImplementsPreferenceSetCategoriesUnion()                       {}
func (UnionBool) ImplementsPreferenceSetWorkflowsUnion()                        {}
func (UnionBool) ImplementsPreferenceSetChannelTypesChatUnionParam()            {}
func (UnionBool) ImplementsPreferenceSetChannelTypesChatUnion()                 {}
func (UnionBool) ImplementsPreferenceSetChannelTypesEmailUnionParam()           {}
func (UnionBool) ImplementsPreferenceSetChannelTypesEmailUnion()                {}
func (UnionBool) ImplementsPreferenceSetChannelTypesHTTPUnionParam()            {}
func (UnionBool) ImplementsPreferenceSetChannelTypesHTTPUnion()                 {}
func (UnionBool) ImplementsPreferenceSetChannelTypesInAppFeedUnionParam()       {}
func (UnionBool) ImplementsPreferenceSetChannelTypesInAppFeedUnion()            {}
func (UnionBool) ImplementsPreferenceSetChannelTypesPushUnionParam()            {}
func (UnionBool) ImplementsPreferenceSetChannelTypesPushUnion()                 {}
func (UnionBool) ImplementsPreferenceSetChannelTypesSMSUnionParam()             {}
func (UnionBool) ImplementsPreferenceSetChannelTypesSMSUnion()                  {}
func (UnionBool) ImplementsPreferenceSetRequestCategoriesUnionParam()           {}
func (UnionBool) ImplementsPreferenceSetRequestWorkflowsUnionParam()            {}

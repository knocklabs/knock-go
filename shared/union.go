// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsInlineTenantRequestUnionParam()                               {}
func (UnionString) ImplementsUserListMessagesResponseActorsUnion()                         {}
func (UnionString) ImplementsUserListMessagesResponseRecipientUnion()                      {}
func (UnionString) ImplementsObjectListMessagesResponseActorsUnion()                       {}
func (UnionString) ImplementsObjectListMessagesResponseRecipientUnion()                    {}
func (UnionString) ImplementsObjectAddSubscriptionsParamsRecipientUnion()                  {}
func (UnionString) ImplementsObjectDeleteSubscriptionsParamsRecipientUnion()               {}
func (UnionString) ImplementsObjectListSubscriptionsParamsRecipientUnion()                 {}
func (UnionString) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion() {}
func (UnionString) ImplementsMessageListResponseActorsUnion()                              {}
func (UnionString) ImplementsMessageListResponseRecipientUnion()                           {}
func (UnionString) ImplementsMessageArchiveResponseActorsUnion()                           {}
func (UnionString) ImplementsMessageArchiveResponseRecipientUnion()                        {}
func (UnionString) ImplementsMessageGetResponseActorsUnion()                               {}
func (UnionString) ImplementsMessageGetResponseRecipientUnion()                            {}
func (UnionString) ImplementsMessageListDeliveryLogsResponseRequestBodyUnion()             {}
func (UnionString) ImplementsMessageListDeliveryLogsResponseResponseBodyUnion()            {}
func (UnionString) ImplementsMessageListEventsResponseRecipientUnion()                     {}
func (UnionString) ImplementsMessageMarkAsInteractedResponseActorsUnion()                  {}
func (UnionString) ImplementsMessageMarkAsInteractedResponseRecipientUnion()               {}
func (UnionString) ImplementsMessageMarkAsReadResponseActorsUnion()                        {}
func (UnionString) ImplementsMessageMarkAsReadResponseRecipientUnion()                     {}
func (UnionString) ImplementsMessageMarkAsSeenResponseActorsUnion()                        {}
func (UnionString) ImplementsMessageMarkAsSeenResponseRecipientUnion()                     {}
func (UnionString) ImplementsMessageMarkAsUnreadResponseActorsUnion()                      {}
func (UnionString) ImplementsMessageMarkAsUnreadResponseRecipientUnion()                   {}
func (UnionString) ImplementsMessageMarkAsUnseenResponseActorsUnion()                      {}
func (UnionString) ImplementsMessageMarkAsUnseenResponseRecipientUnion()                   {}
func (UnionString) ImplementsMessageUnarchiveResponseActorsUnion()                         {}
func (UnionString) ImplementsMessageUnarchiveResponseRecipientUnion()                      {}
func (UnionString) ImplementsMessageBatchArchiveResponseActorsUnion()                      {}
func (UnionString) ImplementsMessageBatchArchiveResponseRecipientUnion()                   {}
func (UnionString) ImplementsMessageBatchMarkAsInteractedResponseActorsUnion()             {}
func (UnionString) ImplementsMessageBatchMarkAsInteractedResponseRecipientUnion()          {}
func (UnionString) ImplementsMessageBatchMarkAsReadResponseActorsUnion()                   {}
func (UnionString) ImplementsMessageBatchMarkAsReadResponseRecipientUnion()                {}
func (UnionString) ImplementsMessageBatchMarkAsSeenResponseActorsUnion()                   {}
func (UnionString) ImplementsMessageBatchMarkAsSeenResponseRecipientUnion()                {}
func (UnionString) ImplementsMessageBatchMarkAsUnreadResponseActorsUnion()                 {}
func (UnionString) ImplementsMessageBatchMarkAsUnreadResponseRecipientUnion()              {}
func (UnionString) ImplementsMessageBatchMarkAsUnseenResponseActorsUnion()                 {}
func (UnionString) ImplementsMessageBatchMarkAsUnseenResponseRecipientUnion()              {}
func (UnionString) ImplementsMessageBatchUnarchiveResponseActorsUnion()                    {}
func (UnionString) ImplementsMessageBatchUnarchiveResponseRecipientUnion()                 {}
func (UnionString) ImplementsWorkflowTriggerParamsActorUnion()                             {}
func (UnionString) ImplementsWorkflowTriggerParamsRecipientUnion()                         {}
func (UnionString) ImplementsScheduleNewParamsRecipientUnion()                             {}
func (UnionString) ImplementsScheduleUpdateParamsActorUnion()                              {}
func (UnionString) ImplementsScheduleListParamsRecipientUnion()                            {}

type UnionBool bool

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
func (UnionBool) ImplementsPreferenceSetRequestWorkflowsUnionParam()      {}
func (UnionBool) ImplementsUserGetPreferencesResponseCategoriesUnion()    {}
func (UnionBool) ImplementsUserGetPreferencesResponseWorkflowsUnion()     {}
func (UnionBool) ImplementsUserListPreferencesResponseCategoriesUnion()   {}
func (UnionBool) ImplementsUserListPreferencesResponseWorkflowsUnion()    {}
func (UnionBool) ImplementsUserSetPreferencesResponseCategoriesUnion()    {}
func (UnionBool) ImplementsUserSetPreferencesResponseWorkflowsUnion()     {}
func (UnionBool) ImplementsObjectGetPreferencesResponseCategoriesUnion()  {}
func (UnionBool) ImplementsObjectGetPreferencesResponseWorkflowsUnion()   {}
func (UnionBool) ImplementsObjectListPreferencesResponseCategoriesUnion() {}
func (UnionBool) ImplementsObjectListPreferencesResponseWorkflowsUnion()  {}
func (UnionBool) ImplementsObjectSetPreferencesResponseCategoriesUnion()  {}
func (UnionBool) ImplementsObjectSetPreferencesResponseWorkflowsUnion()   {}

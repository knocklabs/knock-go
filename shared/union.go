// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsUserListMessagesResponseActorsUnion()                         {}
func (UnionString) ImplementsUserListMessagesResponseRecipientUnion()                      {}
func (UnionString) ImplementsObjectListMessagesResponseActorsUnion()                       {}
func (UnionString) ImplementsObjectListMessagesResponseRecipientUnion()                    {}
func (UnionString) ImplementsObjectAddSubscriptionsParamsRecipientUnion()                  {}
func (UnionString) ImplementsObjectDeleteSubscriptionsParamsRecipientUnion()               {}
func (UnionString) ImplementsObjectListSubscriptionsParamsRecipientUnion()                 {}
func (UnionString) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion() {}
func (UnionString) ImplementsTenantBulkSetParamsTenantUnion()                              {}
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
func (UnionString) ImplementsWorkflowTriggerParamsTenantUnion()                            {}
func (UnionString) ImplementsScheduleNewParamsRecipientUnion()                             {}
func (UnionString) ImplementsScheduleNewParamsTenantUnion()                                {}
func (UnionString) ImplementsScheduleUpdateParamsActorUnion()                              {}
func (UnionString) ImplementsScheduleUpdateParamsTenantUnion()                             {}
func (UnionString) ImplementsScheduleListParamsRecipientUnion()                            {}

type UnionBool bool

func (UnionBool) ImplementsUserGetPreferencesResponseCategoriesUnion() {}
func (UnionBool) ImplementsUserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseChannelTypesChatUnion()      {}
func (UnionBool) ImplementsUserGetPreferencesResponseChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsUserGetPreferencesResponseChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsUserGetPreferencesResponseChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsUserGetPreferencesResponseChannelTypesPushUnion()      {}
func (UnionBool) ImplementsUserGetPreferencesResponseChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsUserGetPreferencesResponseWorkflowsUnion()             {}
func (UnionBool) ImplementsUserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseCategoriesUnion() {}
func (UnionBool) ImplementsUserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseChannelTypesChatUnion()      {}
func (UnionBool) ImplementsUserListPreferencesResponseChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsUserListPreferencesResponseChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsUserListPreferencesResponseChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsUserListPreferencesResponseChannelTypesPushUnion()      {}
func (UnionBool) ImplementsUserListPreferencesResponseChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsUserListPreferencesResponseWorkflowsUnion()             {}
func (UnionBool) ImplementsUserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseCategoriesUnion() {}
func (UnionBool) ImplementsUserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseChannelTypesChatUnion()      {}
func (UnionBool) ImplementsUserSetPreferencesResponseChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsUserSetPreferencesResponseChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsUserSetPreferencesResponseChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsUserSetPreferencesResponseChannelTypesPushUnion()      {}
func (UnionBool) ImplementsUserSetPreferencesResponseChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsUserSetPreferencesResponseWorkflowsUnion()             {}
func (UnionBool) ImplementsUserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsUserUpdateParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesChannelTypesChatUnion()      {}
func (UnionBool) ImplementsUserUpdateParamsPreferencesChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsUserUpdateParamsPreferencesChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsUserUpdateParamsPreferencesChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsUserUpdateParamsPreferencesChannelTypesPushUnion()      {}
func (UnionBool) ImplementsUserUpdateParamsPreferencesChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsUserUpdateParamsPreferencesWorkflowsUnion()             {}
func (UnionBool) ImplementsUserUpdateParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserUpdateParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsCategoriesUnion() {}
func (UnionBool) ImplementsUserSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsChannelTypesChatUnion()      {}
func (UnionBool) ImplementsUserSetPreferencesParamsChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsUserSetPreferencesParamsChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsUserSetPreferencesParamsChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsUserSetPreferencesParamsChannelTypesPushUnion()      {}
func (UnionBool) ImplementsUserSetPreferencesParamsChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsUserSetPreferencesParamsWorkflowsUnion()             {}
func (UnionBool) ImplementsUserSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesChatUnion()      {}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesPushUnion()      {}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsUnion()             {}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserBulkIdentifyParamsUsersPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesChatUnion()      {}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesPushUnion()      {}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsUnion()             {}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsUserBulkSetPreferencesParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseCategoriesUnion() {}
func (UnionBool) ImplementsObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseChannelTypesChatUnion()      {}
func (UnionBool) ImplementsObjectGetPreferencesResponseChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsObjectGetPreferencesResponseChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsObjectGetPreferencesResponseChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsObjectGetPreferencesResponseChannelTypesPushUnion()      {}
func (UnionBool) ImplementsObjectGetPreferencesResponseChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsObjectGetPreferencesResponseWorkflowsUnion()             {}
func (UnionBool) ImplementsObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseCategoriesUnion() {}
func (UnionBool) ImplementsObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseChannelTypesChatUnion()      {}
func (UnionBool) ImplementsObjectListPreferencesResponseChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsObjectListPreferencesResponseChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsObjectListPreferencesResponseChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsObjectListPreferencesResponseChannelTypesPushUnion()      {}
func (UnionBool) ImplementsObjectListPreferencesResponseChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsObjectListPreferencesResponseWorkflowsUnion()             {}
func (UnionBool) ImplementsObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseCategoriesUnion() {}
func (UnionBool) ImplementsObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseChannelTypesChatUnion()      {}
func (UnionBool) ImplementsObjectSetPreferencesResponseChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsObjectSetPreferencesResponseChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsObjectSetPreferencesResponseChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsObjectSetPreferencesResponseChannelTypesPushUnion()      {}
func (UnionBool) ImplementsObjectSetPreferencesResponseChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsObjectSetPreferencesResponseWorkflowsUnion()             {}
func (UnionBool) ImplementsObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesChannelTypesChatUnion()      {}
func (UnionBool) ImplementsObjectSetParamsPreferencesChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsObjectSetParamsPreferencesChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsObjectSetParamsPreferencesChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsObjectSetParamsPreferencesChannelTypesPushUnion()      {}
func (UnionBool) ImplementsObjectSetParamsPreferencesChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsObjectSetParamsPreferencesWorkflowsUnion()             {}
func (UnionBool) ImplementsObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsCategoriesUnion() {}
func (UnionBool) ImplementsObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsChannelTypesChatUnion()      {}
func (UnionBool) ImplementsObjectSetPreferencesParamsChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsObjectSetPreferencesParamsChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsObjectSetPreferencesParamsChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsObjectSetPreferencesParamsChannelTypesPushUnion()      {}
func (UnionBool) ImplementsObjectSetPreferencesParamsChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsObjectSetPreferencesParamsWorkflowsUnion()             {}
func (UnionBool) ImplementsObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectSetPreferencesParamsWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesChatUnion()      {}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesPushUnion()      {}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsUnion()             {}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesChannelTypesChatUnion()      {}
func (UnionBool) ImplementsTenantSetParamsPreferencesChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsTenantSetParamsPreferencesChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsTenantSetParamsPreferencesChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsTenantSetParamsPreferencesChannelTypesPushUnion()      {}
func (UnionBool) ImplementsTenantSetParamsPreferencesChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsTenantSetParamsPreferencesWorkflowsUnion()             {}
func (UnionBool) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantSetParamsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesUnion() {}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesChatUnion()      {}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesEmailUnion()     {}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesHTTPUnion()      {}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesInAppFeedUnion() {}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesPushUnion()      {}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetChannelTypesSMSUnion()       {}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsUnion()             {}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantSetParamsSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesChannelTypesSMSUnion() {}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsUnion()       {}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsTenantBulkSetParamsTenantsTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesWorkflowsUnion() {}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetCategoriesUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetWorkflowsUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsWorkflowTriggerParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesChannelTypesChatUnion()  {}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesChannelTypesEmailUnion() {}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesChannelTypesHTTPUnion()  {}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesChannelTypesPushUnion() {}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesChannelTypesSMSUnion()  {}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesWorkflowsUnion()        {}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetCategoriesUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetWorkflowsUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleNewParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesCategoriesUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsActorInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesChannelTypesSMSUnion() {}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesWorkflowsUnion()       {}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetCategoriesUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetWorkflowsUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsScheduleUpdateParamsTenantTenantRequestSettingsPreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesChatUnion()  {}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailUnion() {}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPUnion()  {}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesPushUnion() {}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSUnion()  {}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsUnion()        {}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsAudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesUnion() {}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatUnion() {}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPUnion() {}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushUnion() {}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSUnion()  {}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsUnion()        {}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}
func (UnionBool) ImplementsAudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

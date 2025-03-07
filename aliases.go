// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"github.com/stainless-sdks/knock-go/internal/apierror"
	"github.com/stainless-sdks/knock-go/shared"
)

type Error = apierror.Error

// Channel data for various channel types
//
// This is an alias to an internal type.
type ChannelData = shared.ChannelData

// Channel data for push providers
//
// This is an alias to an internal type.
type ChannelDataData = shared.ChannelDataData

// Set channel data for a type of channel
//
// This is an alias to an internal type.
type ChannelDataRequestParam = shared.ChannelDataRequestParam

// Channel data for push providers
//
// This is an alias to an internal type.
type ChannelDataRequestDataUnionParam = shared.ChannelDataRequestDataUnionParam

// A condition to be evaluated
//
// This is an alias to an internal type.
type Condition = shared.Condition

// This is an alias to an internal type.
type ConditionOperator = shared.ConditionOperator

// This is an alias to an internal value.
const ConditionOperatorEqualTo = shared.ConditionOperatorEqualTo

// This is an alias to an internal value.
const ConditionOperatorNotEqualTo = shared.ConditionOperatorNotEqualTo

// This is an alias to an internal value.
const ConditionOperatorGreaterThan = shared.ConditionOperatorGreaterThan

// This is an alias to an internal value.
const ConditionOperatorLessThan = shared.ConditionOperatorLessThan

// This is an alias to an internal value.
const ConditionOperatorGreaterThanOrEqualTo = shared.ConditionOperatorGreaterThanOrEqualTo

// This is an alias to an internal value.
const ConditionOperatorLessThanOrEqualTo = shared.ConditionOperatorLessThanOrEqualTo

// This is an alias to an internal value.
const ConditionOperatorContains = shared.ConditionOperatorContains

// This is an alias to an internal value.
const ConditionOperatorNotContains = shared.ConditionOperatorNotContains

// This is an alias to an internal value.
const ConditionOperatorEmpty = shared.ConditionOperatorEmpty

// This is an alias to an internal value.
const ConditionOperatorNotEmpty = shared.ConditionOperatorNotEmpty

// This is an alias to an internal value.
const ConditionOperatorContainsAll = shared.ConditionOperatorContainsAll

// This is an alias to an internal value.
const ConditionOperatorIsTimestamp = shared.ConditionOperatorIsTimestamp

// This is an alias to an internal value.
const ConditionOperatorIsNotTimestamp = shared.ConditionOperatorIsNotTimestamp

// This is an alias to an internal value.
const ConditionOperatorIsTimestampAfter = shared.ConditionOperatorIsTimestampAfter

// This is an alias to an internal value.
const ConditionOperatorIsTimestampBefore = shared.ConditionOperatorIsTimestampBefore

// This is an alias to an internal value.
const ConditionOperatorIsTimestampBetween = shared.ConditionOperatorIsTimestampBetween

// This is an alias to an internal value.
const ConditionOperatorIsAudienceMember = shared.ConditionOperatorIsAudienceMember

// A condition to be evaluated
//
// This is an alias to an internal type.
type ConditionParam = shared.ConditionParam

// Discord channel data
//
// This is an alias to an internal type.
type DiscordChannelData = shared.DiscordChannelData

// Discord channel connection
//
// This is an alias to an internal type.
type DiscordChannelDataConnection = shared.DiscordChannelDataConnection

// Discord channel connection
//
// This is an alias to an internal type.
type DiscordChannelDataConnectionsDiscordChannelConnection = shared.DiscordChannelDataConnectionsDiscordChannelConnection

// Discord incoming webhook connection
//
// This is an alias to an internal type.
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnection = shared.DiscordChannelDataConnectionsDiscordIncomingWebhookConnection

// The incoming webhook
//
// This is an alias to an internal type.
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook = shared.DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook

// Discord channel data
//
// This is an alias to an internal type.
type DiscordChannelDataParam = shared.DiscordChannelDataParam

// Discord channel connection
//
// This is an alias to an internal type.
type DiscordChannelDataConnectionsUnionParam = shared.DiscordChannelDataConnectionsUnionParam

// Discord channel connection
//
// This is an alias to an internal type.
type DiscordChannelDataConnectionsDiscordChannelConnectionParam = shared.DiscordChannelDataConnectionsDiscordChannelConnectionParam

// Discord incoming webhook connection
//
// This is an alias to an internal type.
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam = shared.DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam

// The incoming webhook
//
// This is an alias to an internal type.
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookParam = shared.DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookParam

// Allows inline setting channel data for a recipient
//
// This is an alias to an internal type.
type InlineChannelDataRequestParam = shared.InlineChannelDataRequestParam

// A set of parameters to inline-identify a user with. Inline identifying the user
// will ensure that the user is available before the request is executed in Knock.
// It will perform an upsert against the user you're supplying, replacing any
// properties specified.
//
// This is an alias to an internal type.
type InlineIdentifyUserRequestParam = shared.InlineIdentifyUserRequestParam

// Inline identifies a custom object belonging to a collection
//
// This is an alias to an internal type.
type InlineObjectRequestParam = shared.InlineObjectRequestParam

// Inline set preferences for a recipient, where the key is the preference set name
//
// This is an alias to an internal type.
type InlinePreferenceSetRequestParam = shared.InlinePreferenceSetRequestParam

// An inline tenant request
//
// This is an alias to an internal type.
type InlineTenantRequestUnionParam = shared.InlineTenantRequestUnionParam

// Microsoft Teams channel data
//
// This is an alias to an internal type.
type MsTeamsChannelData = shared.MsTeamsChannelData

// Microsoft Teams token connection
//
// This is an alias to an internal type.
type MsTeamsChannelDataConnection = shared.MsTeamsChannelDataConnection

// Microsoft Teams token connection
//
// This is an alias to an internal type.
type MsTeamsChannelDataConnectionsMsTeamsTokenConnection = shared.MsTeamsChannelDataConnectionsMsTeamsTokenConnection

// Microsoft Teams incoming webhook connection
//
// This is an alias to an internal type.
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection = shared.MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection

// The incoming webhook
//
// This is an alias to an internal type.
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook = shared.MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook

// Microsoft Teams channel data
//
// This is an alias to an internal type.
type MsTeamsChannelDataParam = shared.MsTeamsChannelDataParam

// Microsoft Teams token connection
//
// This is an alias to an internal type.
type MsTeamsChannelDataConnectionsUnionParam = shared.MsTeamsChannelDataConnectionsUnionParam

// Microsoft Teams token connection
//
// This is an alias to an internal type.
type MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam = shared.MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam

// Microsoft Teams incoming webhook connection
//
// This is an alias to an internal type.
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam = shared.MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam

// The incoming webhook
//
// This is an alias to an internal type.
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookParam = shared.MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookParam

// A custom-object entity which belongs to a collection.
//
// This is an alias to an internal type.
type Object = shared.Object

// OneSignal channel data
//
// This is an alias to an internal type.
type OneSignalChannelData = shared.OneSignalChannelData

// OneSignal channel data
//
// This is an alias to an internal type.
type OneSignalChannelDataParam = shared.OneSignalChannelDataParam

// A preference set object.
//
// This is an alias to an internal type.
type PreferenceSet = shared.PreferenceSet

// Workflow or category preferences within a preference set
//
// This is an alias to an internal type.
type PreferenceSetCategoriesUnion = shared.PreferenceSetCategoriesUnion

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
//
// This is an alias to an internal type.
type PreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject = shared.PreferenceSetCategoriesPreferenceSetWorkflowCategorySettingObject

// Workflow or category preferences within a preference set
//
// This is an alias to an internal type.
type PreferenceSetWorkflowsUnion = shared.PreferenceSetWorkflowsUnion

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
//
// This is an alias to an internal type.
type PreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject = shared.PreferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObject

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypeSetting = shared.PreferenceSetChannelTypeSetting

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypeSettingParam = shared.PreferenceSetChannelTypeSettingParam

// Channel type preferences
//
// This is an alias to an internal type.
type PreferenceSetChannelTypes = shared.PreferenceSetChannelTypes

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesChatUnion = shared.PreferenceSetChannelTypesChatUnion

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesEmailUnion = shared.PreferenceSetChannelTypesEmailUnion

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesHTTPUnion = shared.PreferenceSetChannelTypesHTTPUnion

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesInAppFeedUnion = shared.PreferenceSetChannelTypesInAppFeedUnion

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesPushUnion = shared.PreferenceSetChannelTypesPushUnion

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesSMSUnion = shared.PreferenceSetChannelTypesSMSUnion

// Channel type preferences
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesParam = shared.PreferenceSetChannelTypesParam

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesChatUnionParam = shared.PreferenceSetChannelTypesChatUnionParam

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesEmailUnionParam = shared.PreferenceSetChannelTypesEmailUnionParam

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesHTTPUnionParam = shared.PreferenceSetChannelTypesHTTPUnionParam

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesInAppFeedUnionParam = shared.PreferenceSetChannelTypesInAppFeedUnionParam

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesPushUnionParam = shared.PreferenceSetChannelTypesPushUnionParam

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// This is an alias to an internal type.
type PreferenceSetChannelTypesSMSUnionParam = shared.PreferenceSetChannelTypesSMSUnionParam

// Set preferences for a recipient
//
// This is an alias to an internal type.
type PreferenceSetRequestParam = shared.PreferenceSetRequestParam

// Workflow or category preferences within a preference set
//
// This is an alias to an internal type.
type PreferenceSetRequestCategoriesUnionParam = shared.PreferenceSetRequestCategoriesUnionParam

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
//
// This is an alias to an internal type.
type PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam = shared.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam

// Workflow or category preferences within a preference set
//
// This is an alias to an internal type.
type PreferenceSetRequestWorkflowsUnionParam = shared.PreferenceSetRequestWorkflowsUnionParam

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
//
// This is an alias to an internal type.
type PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam = shared.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam

// Channel data for push providers
//
// This is an alias to an internal type.
type PushChannelData = shared.PushChannelData

// Channel data for push providers
//
// This is an alias to an internal type.
type PushChannelDataParam = shared.PushChannelDataParam

// A recipient, which is either a user or an object
//
// This is an alias to an internal type.
type Recipient = shared.Recipient

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// This is an alias to an internal type.
type RecipientRequestUnionParam = shared.RecipientRequestUnionParam

// A schedule that represents a recurring workflow execution
//
// This is an alias to an internal type.
type Schedule = shared.Schedule

// A schedule repeat rule
//
// This is an alias to an internal type.
type ScheduleRepeatRule = shared.ScheduleRepeatRule

// This is an alias to an internal type.
type ScheduleRepeatRuleFrequency = shared.ScheduleRepeatRuleFrequency

// This is an alias to an internal value.
const ScheduleRepeatRuleFrequencyDaily = shared.ScheduleRepeatRuleFrequencyDaily

// This is an alias to an internal value.
const ScheduleRepeatRuleFrequencyWeekly = shared.ScheduleRepeatRuleFrequencyWeekly

// This is an alias to an internal value.
const ScheduleRepeatRuleFrequencyMonthly = shared.ScheduleRepeatRuleFrequencyMonthly

// This is an alias to an internal value.
const ScheduleRepeatRuleFrequencyHourly = shared.ScheduleRepeatRuleFrequencyHourly

// This is an alias to an internal type.
type ScheduleRepeatRuleDay = shared.ScheduleRepeatRuleDay

// This is an alias to an internal value.
const ScheduleRepeatRuleDayMon = shared.ScheduleRepeatRuleDayMon

// This is an alias to an internal value.
const ScheduleRepeatRuleDayTue = shared.ScheduleRepeatRuleDayTue

// This is an alias to an internal value.
const ScheduleRepeatRuleDayWed = shared.ScheduleRepeatRuleDayWed

// This is an alias to an internal value.
const ScheduleRepeatRuleDayThu = shared.ScheduleRepeatRuleDayThu

// This is an alias to an internal value.
const ScheduleRepeatRuleDayFri = shared.ScheduleRepeatRuleDayFri

// This is an alias to an internal value.
const ScheduleRepeatRuleDaySat = shared.ScheduleRepeatRuleDaySat

// This is an alias to an internal value.
const ScheduleRepeatRuleDaySun = shared.ScheduleRepeatRuleDaySun

// A schedule repeat rule
//
// This is an alias to an internal type.
type ScheduleRepeatRuleParam = shared.ScheduleRepeatRuleParam

// Slack channel data
//
// This is an alias to an internal type.
type SlackChannelData = shared.SlackChannelData

// A Slack connection, which either includes a channel_id or a user_id
//
// This is an alias to an internal type.
type SlackChannelDataConnection = shared.SlackChannelDataConnection

// A Slack connection, which either includes a channel_id or a user_id
//
// This is an alias to an internal type.
type SlackChannelDataConnectionsSlackTokenConnection = shared.SlackChannelDataConnectionsSlackTokenConnection

// An incoming webhook Slack connection
//
// This is an alias to an internal type.
type SlackChannelDataConnectionsSlackIncomingWebhookConnection = shared.SlackChannelDataConnectionsSlackIncomingWebhookConnection

// A token that's used to store the access token for a Slack workspace.
//
// This is an alias to an internal type.
type SlackChannelDataToken = shared.SlackChannelDataToken

// Slack channel data
//
// This is an alias to an internal type.
type SlackChannelDataParam = shared.SlackChannelDataParam

// A Slack connection, which either includes a channel_id or a user_id
//
// This is an alias to an internal type.
type SlackChannelDataConnectionsUnionParam = shared.SlackChannelDataConnectionsUnionParam

// A Slack connection, which either includes a channel_id or a user_id
//
// This is an alias to an internal type.
type SlackChannelDataConnectionsSlackTokenConnectionParam = shared.SlackChannelDataConnectionsSlackTokenConnectionParam

// An incoming webhook Slack connection
//
// This is an alias to an internal type.
type SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam = shared.SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam

// A token that's used to store the access token for a Slack workspace.
//
// This is an alias to an internal type.
type SlackChannelDataTokenParam = shared.SlackChannelDataTokenParam

// A subscription object
//
// This is an alias to an internal type.
type Subscription = shared.Subscription

// A tenant entity
//
// This is an alias to an internal type.
type Tenant = shared.Tenant

// A tenant to be set in the system
//
// This is an alias to an internal type.
type TenantRequestParam = shared.TenantRequestParam

// This is an alias to an internal type.
type TenantRequestSettingsParam = shared.TenantRequestSettingsParam

// This is an alias to an internal type.
type TenantRequestSettingsBrandingParam = shared.TenantRequestSettingsBrandingParam

// A user object
//
// This is an alias to an internal type.
type User = shared.User

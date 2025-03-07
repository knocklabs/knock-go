// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/tidwall/gjson"
)

// Set channel data for a type of channel
type ChannelDataRequestParam struct {
	// Channel data for push providers
	Data param.Field[ChannelDataRequestDataUnionParam] `json:"data,required"`
}

func (r ChannelDataRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type ChannelDataRequestDataParam struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r ChannelDataRequestDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChannelDataRequestDataParam) ImplementsChannelDataRequestDataUnionParam() {}

// Channel data for push providers
//
// Satisfied by [shared.PushChannelDataParam], [shared.OneSignalChannelDataParam],
// [shared.SlackChannelDataParam], [shared.MsTeamsChannelDataParam],
// [shared.DiscordChannelDataParam], [ChannelDataRequestDataParam].
type ChannelDataRequestDataUnionParam interface {
	ImplementsChannelDataRequestDataUnionParam()
}

// A condition to be evaluated
type Condition struct {
	Argument string            `json:"argument,required,nullable"`
	Operator ConditionOperator `json:"operator,required"`
	Variable string            `json:"variable,required"`
	JSON     conditionJSON     `json:"-"`
}

// conditionJSON contains the JSON metadata for the struct [Condition]
type conditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Condition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionJSON) RawJSON() string {
	return r.raw
}

type ConditionOperator string

const (
	ConditionOperatorEqualTo              ConditionOperator = "equal_to"
	ConditionOperatorNotEqualTo           ConditionOperator = "not_equal_to"
	ConditionOperatorGreaterThan          ConditionOperator = "greater_than"
	ConditionOperatorLessThan             ConditionOperator = "less_than"
	ConditionOperatorGreaterThanOrEqualTo ConditionOperator = "greater_than_or_equal_to"
	ConditionOperatorLessThanOrEqualTo    ConditionOperator = "less_than_or_equal_to"
	ConditionOperatorContains             ConditionOperator = "contains"
	ConditionOperatorNotContains          ConditionOperator = "not_contains"
	ConditionOperatorEmpty                ConditionOperator = "empty"
	ConditionOperatorNotEmpty             ConditionOperator = "not_empty"
	ConditionOperatorContainsAll          ConditionOperator = "contains_all"
	ConditionOperatorIsTimestamp          ConditionOperator = "is_timestamp"
	ConditionOperatorIsNotTimestamp       ConditionOperator = "is_not_timestamp"
	ConditionOperatorIsTimestampAfter     ConditionOperator = "is_timestamp_after"
	ConditionOperatorIsTimestampBefore    ConditionOperator = "is_timestamp_before"
	ConditionOperatorIsTimestampBetween   ConditionOperator = "is_timestamp_between"
	ConditionOperatorIsAudienceMember     ConditionOperator = "is_audience_member"
)

func (r ConditionOperator) IsKnown() bool {
	switch r {
	case ConditionOperatorEqualTo, ConditionOperatorNotEqualTo, ConditionOperatorGreaterThan, ConditionOperatorLessThan, ConditionOperatorGreaterThanOrEqualTo, ConditionOperatorLessThanOrEqualTo, ConditionOperatorContains, ConditionOperatorNotContains, ConditionOperatorEmpty, ConditionOperatorNotEmpty, ConditionOperatorContainsAll, ConditionOperatorIsTimestamp, ConditionOperatorIsNotTimestamp, ConditionOperatorIsTimestampAfter, ConditionOperatorIsTimestampBefore, ConditionOperatorIsTimestampBetween, ConditionOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type ConditionParam struct {
	Argument param.Field[string]            `json:"argument,required"`
	Operator param.Field[ConditionOperator] `json:"operator,required"`
	Variable param.Field[string]            `json:"variable,required"`
}

func (r ConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Discord channel data
type DiscordChannelData struct {
	Connections []DiscordChannelDataConnection `json:"connections,required"`
	JSON        discordChannelDataJSON         `json:"-"`
}

// discordChannelDataJSON contains the JSON metadata for the struct
// [DiscordChannelData]
type discordChannelDataJSON struct {
	Connections apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscordChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discordChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r DiscordChannelData) ImplementsUserGetChannelDataResponseData() {}

func (r DiscordChannelData) ImplementsUserSetChannelDataResponseData() {}

func (r DiscordChannelData) ImplementsObjectGetChannelDataResponseData() {}

func (r DiscordChannelData) ImplementsObjectSetChannelDataResponseData() {}

// Discord channel connection
type DiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID string `json:"channel_id"`
	// This field can have the runtime type of
	// [DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook].
	IncomingWebhook interface{}                      `json:"incoming_webhook"`
	JSON            discordChannelDataConnectionJSON `json:"-"`
	union           DiscordChannelDataConnectionsUnion
}

// discordChannelDataConnectionJSON contains the JSON metadata for the struct
// [DiscordChannelDataConnection]
type discordChannelDataConnectionJSON struct {
	ChannelID       apijson.Field
	IncomingWebhook apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r discordChannelDataConnectionJSON) RawJSON() string {
	return r.raw
}

func (r *DiscordChannelDataConnection) UnmarshalJSON(data []byte) (err error) {
	*r = DiscordChannelDataConnection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DiscordChannelDataConnectionsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.DiscordChannelDataConnectionsDiscordChannelConnection],
// [shared.DiscordChannelDataConnectionsDiscordIncomingWebhookConnection].
func (r DiscordChannelDataConnection) AsUnion() DiscordChannelDataConnectionsUnion {
	return r.union
}

// Discord channel connection
//
// Union satisfied by
// [shared.DiscordChannelDataConnectionsDiscordChannelConnection] or
// [shared.DiscordChannelDataConnectionsDiscordIncomingWebhookConnection].
type DiscordChannelDataConnectionsUnion interface {
	implementsDiscordChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DiscordChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DiscordChannelDataConnectionsDiscordChannelConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DiscordChannelDataConnectionsDiscordIncomingWebhookConnection{}),
		},
	)
}

// Discord channel connection
type DiscordChannelDataConnectionsDiscordChannelConnection struct {
	// The Discord channel ID
	ChannelID string                                                    `json:"channel_id,required"`
	JSON      discordChannelDataConnectionsDiscordChannelConnectionJSON `json:"-"`
}

// discordChannelDataConnectionsDiscordChannelConnectionJSON contains the JSON
// metadata for the struct [DiscordChannelDataConnectionsDiscordChannelConnection]
type discordChannelDataConnectionsDiscordChannelConnectionJSON struct {
	ChannelID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscordChannelDataConnectionsDiscordChannelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discordChannelDataConnectionsDiscordChannelConnectionJSON) RawJSON() string {
	return r.raw
}

func (r DiscordChannelDataConnectionsDiscordChannelConnection) implementsDiscordChannelDataConnection() {
}

// Discord incoming webhook connection
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook `json:"incoming_webhook,required"`
	JSON            discordChannelDataConnectionsDiscordIncomingWebhookConnectionJSON            `json:"-"`
}

// discordChannelDataConnectionsDiscordIncomingWebhookConnectionJSON contains the
// JSON metadata for the struct
// [DiscordChannelDataConnectionsDiscordIncomingWebhookConnection]
type discordChannelDataConnectionsDiscordIncomingWebhookConnectionJSON struct {
	IncomingWebhook apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DiscordChannelDataConnectionsDiscordIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discordChannelDataConnectionsDiscordIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r DiscordChannelDataConnectionsDiscordIncomingWebhookConnection) implementsDiscordChannelDataConnection() {
}

// The incoming webhook
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL  string                                                                           `json:"url,required"`
	JSON discordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookJSON `json:"-"`
}

// discordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookJSON
// contains the JSON metadata for the struct
// [DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook]
type discordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookJSON) RawJSON() string {
	return r.raw
}

// Discord channel data
type DiscordChannelDataParam struct {
	Connections param.Field[[]DiscordChannelDataConnectionsUnionParam] `json:"connections,required"`
}

func (r DiscordChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataParam) ImplementsChannelDataRequestDataUnionParam() {}

// Discord channel connection
type DiscordChannelDataConnectionParam struct {
	// The Discord channel ID
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r DiscordChannelDataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataConnectionParam) implementsDiscordChannelDataConnectionsUnionParam() {}

// Discord channel connection
//
// Satisfied by
// [shared.DiscordChannelDataConnectionsDiscordChannelConnectionParam],
// [shared.DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam],
// [DiscordChannelDataConnectionParam].
type DiscordChannelDataConnectionsUnionParam interface {
	implementsDiscordChannelDataConnectionsUnionParam()
}

// Discord channel connection
type DiscordChannelDataConnectionsDiscordChannelConnectionParam struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r DiscordChannelDataConnectionsDiscordChannelConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataConnectionsDiscordChannelConnectionParam) implementsDiscordChannelDataConnectionsUnionParam() {
}

// Discord incoming webhook connection
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam struct {
	// The incoming webhook
	IncomingWebhook param.Field[DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookParam] `json:"incoming_webhook,required"`
}

func (r DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam) implementsDiscordChannelDataConnectionsUnionParam() {
}

// The incoming webhook
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookParam struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type MsTeamsChannelData struct {
	Connections []MsTeamsChannelDataConnection `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID string                 `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	JSON            msTeamsChannelDataJSON `json:"-"`
}

// msTeamsChannelDataJSON contains the JSON metadata for the struct
// [MsTeamsChannelData]
type msTeamsChannelDataJSON struct {
	Connections     apijson.Field
	MsTeamsTenantID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MsTeamsChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r msTeamsChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r MsTeamsChannelData) ImplementsUserGetChannelDataResponseData() {}

func (r MsTeamsChannelData) ImplementsUserSetChannelDataResponseData() {}

func (r MsTeamsChannelData) ImplementsObjectGetChannelDataResponseData() {}

func (r MsTeamsChannelData) ImplementsObjectSetChannelDataResponseData() {}

// Microsoft Teams token connection
type MsTeamsChannelDataConnection struct {
	// This field can have the runtime type of
	// [MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook].
	IncomingWebhook interface{} `json:"incoming_webhook"`
	// The Microsoft Teams channel ID
	MsTeamsChannelID string `json:"ms_teams_channel_id,nullable" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID string `json:"ms_teams_team_id,nullable" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID string `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID string                           `json:"ms_teams_user_id,nullable" format:"uuid"`
	JSON          msTeamsChannelDataConnectionJSON `json:"-"`
	union         MsTeamsChannelDataConnectionsUnion
}

// msTeamsChannelDataConnectionJSON contains the JSON metadata for the struct
// [MsTeamsChannelDataConnection]
type msTeamsChannelDataConnectionJSON struct {
	IncomingWebhook  apijson.Field
	MsTeamsChannelID apijson.Field
	MsTeamsTeamID    apijson.Field
	MsTeamsTenantID  apijson.Field
	MsTeamsUserID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r msTeamsChannelDataConnectionJSON) RawJSON() string {
	return r.raw
}

func (r *MsTeamsChannelDataConnection) UnmarshalJSON(data []byte) (err error) {
	*r = MsTeamsChannelDataConnection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MsTeamsChannelDataConnectionsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.MsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [shared.MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection].
func (r MsTeamsChannelDataConnection) AsUnion() MsTeamsChannelDataConnectionsUnion {
	return r.union
}

// Microsoft Teams token connection
//
// Union satisfied by [shared.MsTeamsChannelDataConnectionsMsTeamsTokenConnection]
// or [shared.MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection].
type MsTeamsChannelDataConnectionsUnion interface {
	implementsMsTeamsChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MsTeamsChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MsTeamsChannelDataConnectionsMsTeamsTokenConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection{}),
		},
	)
}

// Microsoft Teams token connection
type MsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID string `json:"ms_teams_channel_id,nullable" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID string `json:"ms_teams_team_id,nullable" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID string `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID string                                                  `json:"ms_teams_user_id,nullable" format:"uuid"`
	JSON          msTeamsChannelDataConnectionsMsTeamsTokenConnectionJSON `json:"-"`
}

// msTeamsChannelDataConnectionsMsTeamsTokenConnectionJSON contains the JSON
// metadata for the struct [MsTeamsChannelDataConnectionsMsTeamsTokenConnection]
type msTeamsChannelDataConnectionsMsTeamsTokenConnectionJSON struct {
	MsTeamsChannelID apijson.Field
	MsTeamsTeamID    apijson.Field
	MsTeamsTenantID  apijson.Field
	MsTeamsUserID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MsTeamsChannelDataConnectionsMsTeamsTokenConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r msTeamsChannelDataConnectionsMsTeamsTokenConnectionJSON) RawJSON() string {
	return r.raw
}

func (r MsTeamsChannelDataConnectionsMsTeamsTokenConnection) implementsMsTeamsChannelDataConnection() {
}

// Microsoft Teams incoming webhook connection
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook `json:"incoming_webhook,required"`
	JSON            msTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionJSON            `json:"-"`
}

// msTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionJSON contains the
// JSON metadata for the struct
// [MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection]
type msTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionJSON struct {
	IncomingWebhook apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r msTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) implementsMsTeamsChannelDataConnection() {
}

// The incoming webhook
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL  string                                                                           `json:"url,required"`
	JSON msTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookJSON `json:"-"`
}

// msTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookJSON
// contains the JSON metadata for the struct
// [MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook]
type msTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r msTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookJSON) RawJSON() string {
	return r.raw
}

// Microsoft Teams channel data
type MsTeamsChannelDataParam struct {
	Connections param.Field[[]MsTeamsChannelDataConnectionsUnionParam] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r MsTeamsChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataParam) ImplementsChannelDataRequestDataUnionParam() {}

// Microsoft Teams token connection
type MsTeamsChannelDataConnectionParam struct {
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r MsTeamsChannelDataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataConnectionParam) implementsMsTeamsChannelDataConnectionsUnionParam() {}

// Microsoft Teams token connection
//
// Satisfied by [shared.MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam],
// [shared.MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam],
// [MsTeamsChannelDataConnectionParam].
type MsTeamsChannelDataConnectionsUnionParam interface {
	implementsMsTeamsChannelDataConnectionsUnionParam()
}

// Microsoft Teams token connection
type MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam) implementsMsTeamsChannelDataConnectionsUnionParam() {
}

// Microsoft Teams incoming webhook connection
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam struct {
	// The incoming webhook
	IncomingWebhook param.Field[MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookParam] `json:"incoming_webhook,required"`
}

func (r MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam) implementsMsTeamsChannelDataConnectionsUnionParam() {
}

// The incoming webhook
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookParam struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom-object entity which belongs to a collection.
type Object struct {
	ID          string                 `json:"id,required"`
	Typename    string                 `json:"__typename,required"`
	Collection  string                 `json:"collection,required"`
	UpdatedAt   time.Time              `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time              `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        objectJSON             `json:"-"`
}

// objectJSON contains the JSON metadata for the struct [Object]
type objectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Object) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectJSON) RawJSON() string {
	return r.raw
}

func (r Object) ImplementsUserListSchedulesResponseRecipient() {}

func (r Object) ImplementsUserListSchedulesResponseActor() {}

func (r Object) ImplementsUserListSubscriptionsResponseRecipient() {}

func (r Object) ImplementsUserFeedListItemsResponseActivitiesActor() {}

func (r Object) ImplementsUserFeedListItemsResponseActivitiesRecipient() {}

func (r Object) ImplementsUserFeedListItemsResponseActor() {}

func (r Object) ImplementsObjectAddSubscriptionsResponseRecipient() {}

func (r Object) ImplementsObjectDeleteSubscriptionsResponseRecipient() {}

func (r Object) ImplementsObjectListSchedulesResponseRecipient() {}

func (r Object) ImplementsObjectListSchedulesResponseActor() {}

func (r Object) ImplementsObjectListSubscriptionsResponseRecipient() {}

func (r Object) ImplementsMessageListActivitiesResponseActor() {}

func (r Object) ImplementsMessageListActivitiesResponseRecipient() {}

func (r Object) ImplementsScheduleNewResponseRecipient() {}

func (r Object) ImplementsScheduleNewResponseActor() {}

func (r Object) ImplementsScheduleUpdateResponseRecipient() {}

func (r Object) ImplementsScheduleUpdateResponseActor() {}

func (r Object) ImplementsScheduleListResponseRecipient() {}

func (r Object) ImplementsScheduleListResponseActor() {}

func (r Object) ImplementsScheduleDeleteResponseRecipient() {}

func (r Object) ImplementsScheduleDeleteResponseActor() {}

// OneSignal channel data
type OneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs []string                 `json:"player_ids,required" format:"uuid"`
	JSON      oneSignalChannelDataJSON `json:"-"`
}

// oneSignalChannelDataJSON contains the JSON metadata for the struct
// [OneSignalChannelData]
type oneSignalChannelDataJSON struct {
	PlayerIDs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OneSignalChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oneSignalChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r OneSignalChannelData) ImplementsUserGetChannelDataResponseData() {}

func (r OneSignalChannelData) ImplementsUserSetChannelDataResponseData() {}

func (r OneSignalChannelData) ImplementsObjectGetChannelDataResponseData() {}

func (r OneSignalChannelData) ImplementsObjectSetChannelDataResponseData() {}

// OneSignal channel data
type OneSignalChannelDataParam struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r OneSignalChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OneSignalChannelDataParam) ImplementsChannelDataRequestDataUnionParam() {}

// Channel type preferences
type PreferenceSetChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat PreferenceSetChannelTypesChatUnion `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email PreferenceSetChannelTypesEmailUnion `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP PreferenceSetChannelTypesHTTPUnion `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed PreferenceSetChannelTypesInAppFeedUnion `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push PreferenceSetChannelTypesPushUnion `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
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

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Union satisfied by [shared.UnionBool] or
// [shared.PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject].
type PreferenceSetChannelTypesChatUnion interface {
	ImplementsPreferenceSetChannelTypesChatUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesChatUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject{}),
		},
	)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions []Condition                                                            `json:"conditions,required"`
	JSON       preferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectJSON `json:"-"`
}

// preferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectJSON contains
// the JSON metadata for the struct
// [PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject]
type preferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsPreferenceSetChannelTypesChatUnion() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Union satisfied by [shared.UnionBool] or
// [shared.PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type PreferenceSetChannelTypesEmailUnion interface {
	ImplementsPreferenceSetChannelTypesEmailUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesEmailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject{}),
		},
	)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions []Condition                                                             `json:"conditions,required"`
	JSON       preferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectJSON `json:"-"`
}

// preferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectJSON contains
// the JSON metadata for the struct
// [PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject]
type preferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsPreferenceSetChannelTypesEmailUnion() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Union satisfied by [shared.UnionBool] or
// [shared.PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type PreferenceSetChannelTypesHTTPUnion interface {
	ImplementsPreferenceSetChannelTypesHTTPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesHTTPUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject{}),
		},
	)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions []Condition                                                            `json:"conditions,required"`
	JSON       preferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectJSON `json:"-"`
}

// preferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectJSON contains
// the JSON metadata for the struct
// [PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject]
type preferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsPreferenceSetChannelTypesHTTPUnion() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Union satisfied by [shared.UnionBool] or
// [shared.PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type PreferenceSetChannelTypesInAppFeedUnion interface {
	ImplementsPreferenceSetChannelTypesInAppFeedUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesInAppFeedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject{}),
		},
	)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions []Condition                                                                 `json:"conditions,required"`
	JSON       preferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectJSON `json:"-"`
}

// preferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectJSON
// contains the JSON metadata for the struct
// [PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject]
type preferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsPreferenceSetChannelTypesInAppFeedUnion() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Union satisfied by [shared.UnionBool] or
// [shared.PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject].
type PreferenceSetChannelTypesPushUnion interface {
	ImplementsPreferenceSetChannelTypesPushUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesPushUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject{}),
		},
	)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions []Condition                                                            `json:"conditions,required"`
	JSON       preferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectJSON `json:"-"`
}

// preferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectJSON contains
// the JSON metadata for the struct
// [PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject]
type preferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsPreferenceSetChannelTypesPushUnion() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Union satisfied by [shared.UnionBool] or
// [shared.PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type PreferenceSetChannelTypesSMSUnion interface {
	ImplementsPreferenceSetChannelTypesSMSUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PreferenceSetChannelTypesSMSUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject{}),
		},
	)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions []Condition                                                           `json:"conditions,required"`
	JSON       preferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectJSON `json:"-"`
}

// preferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectJSON contains
// the JSON metadata for the struct
// [PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject]
type preferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectJSON struct {
	Conditions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsPreferenceSetChannelTypesSMSUnion() {
}

// Channel type preferences
type PreferenceSetChannelTypesParam struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[PreferenceSetChannelTypesChatUnionParam] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[PreferenceSetChannelTypesEmailUnionParam] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[PreferenceSetChannelTypesHTTPUnionParam] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[PreferenceSetChannelTypesInAppFeedUnionParam] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[PreferenceSetChannelTypesPushUnionParam] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[PreferenceSetChannelTypesSMSUnionParam] `json:"sms"`
}

func (r PreferenceSetChannelTypesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [shared.PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectParam].
type PreferenceSetChannelTypesChatUnionParam interface {
	ImplementsPreferenceSetChannelTypesChatUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectParam struct {
	Conditions param.Field[[]ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesChatPreferenceSetChannelTypeSettingObjectParam) ImplementsPreferenceSetChannelTypesChatUnionParam() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [shared.PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectParam].
type PreferenceSetChannelTypesEmailUnionParam interface {
	ImplementsPreferenceSetChannelTypesEmailUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectParam struct {
	Conditions param.Field[[]ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesEmailPreferenceSetChannelTypeSettingObjectParam) ImplementsPreferenceSetChannelTypesEmailUnionParam() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [shared.PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectParam].
type PreferenceSetChannelTypesHTTPUnionParam interface {
	ImplementsPreferenceSetChannelTypesHTTPUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectParam struct {
	Conditions param.Field[[]ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesHTTPPreferenceSetChannelTypeSettingObjectParam) ImplementsPreferenceSetChannelTypesHTTPUnionParam() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [shared.PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectParam].
type PreferenceSetChannelTypesInAppFeedUnionParam interface {
	ImplementsPreferenceSetChannelTypesInAppFeedUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectParam struct {
	Conditions param.Field[[]ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectParam) ImplementsPreferenceSetChannelTypesInAppFeedUnionParam() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [shared.PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectParam].
type PreferenceSetChannelTypesPushUnionParam interface {
	ImplementsPreferenceSetChannelTypesPushUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectParam struct {
	Conditions param.Field[[]ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesPushPreferenceSetChannelTypeSettingObjectParam) ImplementsPreferenceSetChannelTypesPushUnionParam() {
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [shared.PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectParam].
type PreferenceSetChannelTypesSMSUnionParam interface {
	ImplementsPreferenceSetChannelTypesSMSUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectParam struct {
	Conditions param.Field[[]ConditionParam] `json:"conditions,required"`
}

func (r PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetChannelTypesSMSPreferenceSetChannelTypeSettingObjectParam) ImplementsPreferenceSetChannelTypesSMSUnionParam() {
}

// Set preferences for a recipient
type PreferenceSetRequestParam struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]PreferenceSetRequestCategoriesUnionParam] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]PreferenceSetRequestWorkflowsUnionParam] `json:"workflows"`
}

func (r PreferenceSetRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [shared.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam].
type PreferenceSetRequestCategoriesUnionParam interface {
	ImplementsPreferenceSetRequestCategoriesUnionParam()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam struct {
	// Channel type preferences
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	Conditions   param.Field[[]ConditionParam]               `json:"conditions"`
}

func (r PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam) ImplementsPreferenceSetRequestCategoriesUnionParam() {
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [shared.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam].
type PreferenceSetRequestWorkflowsUnionParam interface {
	ImplementsPreferenceSetRequestWorkflowsUnionParam()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam struct {
	// Channel type preferences
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	Conditions   param.Field[[]ConditionParam]               `json:"conditions"`
}

func (r PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam) ImplementsPreferenceSetRequestWorkflowsUnionParam() {
}

// Channel data for push providers
type PushChannelData struct {
	Tokens []string            `json:"tokens,required"`
	JSON   pushChannelDataJSON `json:"-"`
}

// pushChannelDataJSON contains the JSON metadata for the struct [PushChannelData]
type pushChannelDataJSON struct {
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PushChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pushChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r PushChannelData) ImplementsUserGetChannelDataResponseData() {}

func (r PushChannelData) ImplementsUserSetChannelDataResponseData() {}

func (r PushChannelData) ImplementsObjectGetChannelDataResponseData() {}

func (r PushChannelData) ImplementsObjectSetChannelDataResponseData() {}

// Channel data for push providers
type PushChannelDataParam struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r PushChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PushChannelDataParam) ImplementsChannelDataRequestDataUnionParam() {}

// A schedule repeat rule
type ScheduleRepeatRule struct {
	Typename   string                      `json:"__typename,required"`
	Frequency  ScheduleRepeatRuleFrequency `json:"frequency,required"`
	DayOfMonth int64                       `json:"day_of_month,nullable"`
	Days       []ScheduleRepeatRuleDay     `json:"days,nullable"`
	Hours      int64                       `json:"hours,nullable"`
	Interval   int64                       `json:"interval"`
	Minutes    int64                       `json:"minutes,nullable"`
	JSON       scheduleRepeatRuleJSON      `json:"-"`
}

// scheduleRepeatRuleJSON contains the JSON metadata for the struct
// [ScheduleRepeatRule]
type scheduleRepeatRuleJSON struct {
	Typename    apijson.Field
	Frequency   apijson.Field
	DayOfMonth  apijson.Field
	Days        apijson.Field
	Hours       apijson.Field
	Interval    apijson.Field
	Minutes     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleRepeatRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleRepeatRuleJSON) RawJSON() string {
	return r.raw
}

type ScheduleRepeatRuleFrequency string

const (
	ScheduleRepeatRuleFrequencyDaily   ScheduleRepeatRuleFrequency = "daily"
	ScheduleRepeatRuleFrequencyWeekly  ScheduleRepeatRuleFrequency = "weekly"
	ScheduleRepeatRuleFrequencyMonthly ScheduleRepeatRuleFrequency = "monthly"
	ScheduleRepeatRuleFrequencyHourly  ScheduleRepeatRuleFrequency = "hourly"
)

func (r ScheduleRepeatRuleFrequency) IsKnown() bool {
	switch r {
	case ScheduleRepeatRuleFrequencyDaily, ScheduleRepeatRuleFrequencyWeekly, ScheduleRepeatRuleFrequencyMonthly, ScheduleRepeatRuleFrequencyHourly:
		return true
	}
	return false
}

type ScheduleRepeatRuleDay string

const (
	ScheduleRepeatRuleDayMon ScheduleRepeatRuleDay = "mon"
	ScheduleRepeatRuleDayTue ScheduleRepeatRuleDay = "tue"
	ScheduleRepeatRuleDayWed ScheduleRepeatRuleDay = "wed"
	ScheduleRepeatRuleDayThu ScheduleRepeatRuleDay = "thu"
	ScheduleRepeatRuleDayFri ScheduleRepeatRuleDay = "fri"
	ScheduleRepeatRuleDaySat ScheduleRepeatRuleDay = "sat"
	ScheduleRepeatRuleDaySun ScheduleRepeatRuleDay = "sun"
)

func (r ScheduleRepeatRuleDay) IsKnown() bool {
	switch r {
	case ScheduleRepeatRuleDayMon, ScheduleRepeatRuleDayTue, ScheduleRepeatRuleDayWed, ScheduleRepeatRuleDayThu, ScheduleRepeatRuleDayFri, ScheduleRepeatRuleDaySat, ScheduleRepeatRuleDaySun:
		return true
	}
	return false
}

// A schedule repeat rule
type ScheduleRepeatRuleParam struct {
	Typename   param.Field[string]                      `json:"__typename,required"`
	Frequency  param.Field[ScheduleRepeatRuleFrequency] `json:"frequency,required"`
	DayOfMonth param.Field[int64]                       `json:"day_of_month"`
	Days       param.Field[[]ScheduleRepeatRuleDay]     `json:"days"`
	Hours      param.Field[int64]                       `json:"hours"`
	Interval   param.Field[int64]                       `json:"interval"`
	Minutes    param.Field[int64]                       `json:"minutes"`
}

func (r ScheduleRepeatRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Slack channel data
type SlackChannelData struct {
	Connections []SlackChannelDataConnection `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token SlackChannelDataToken `json:"token,nullable"`
	JSON  slackChannelDataJSON  `json:"-"`
}

// slackChannelDataJSON contains the JSON metadata for the struct
// [SlackChannelData]
type slackChannelDataJSON struct {
	Connections apijson.Field
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlackChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slackChannelDataJSON) RawJSON() string {
	return r.raw
}

func (r SlackChannelData) ImplementsUserGetChannelDataResponseData() {}

func (r SlackChannelData) ImplementsUserSetChannelDataResponseData() {}

func (r SlackChannelData) ImplementsObjectGetChannelDataResponseData() {}

func (r SlackChannelData) ImplementsObjectSetChannelDataResponseData() {}

// A Slack connection, which either includes a channel_id or a user_id
type SlackChannelDataConnection struct {
	AccessToken string                         `json:"access_token,nullable"`
	ChannelID   string                         `json:"channel_id,nullable"`
	URL         string                         `json:"url"`
	UserID      string                         `json:"user_id,nullable"`
	JSON        slackChannelDataConnectionJSON `json:"-"`
	union       SlackChannelDataConnectionsUnion
}

// slackChannelDataConnectionJSON contains the JSON metadata for the struct
// [SlackChannelDataConnection]
type slackChannelDataConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	URL         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r slackChannelDataConnectionJSON) RawJSON() string {
	return r.raw
}

func (r *SlackChannelDataConnection) UnmarshalJSON(data []byte) (err error) {
	*r = SlackChannelDataConnection{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SlackChannelDataConnectionsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.SlackChannelDataConnectionsSlackTokenConnection],
// [shared.SlackChannelDataConnectionsSlackIncomingWebhookConnection].
func (r SlackChannelDataConnection) AsUnion() SlackChannelDataConnectionsUnion {
	return r.union
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Union satisfied by [shared.SlackChannelDataConnectionsSlackTokenConnection] or
// [shared.SlackChannelDataConnectionsSlackIncomingWebhookConnection].
type SlackChannelDataConnectionsUnion interface {
	implementsSlackChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SlackChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SlackChannelDataConnectionsSlackTokenConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SlackChannelDataConnectionsSlackIncomingWebhookConnection{}),
		},
	)
}

// A Slack connection, which either includes a channel_id or a user_id
type SlackChannelDataConnectionsSlackTokenConnection struct {
	AccessToken string                                              `json:"access_token,nullable"`
	ChannelID   string                                              `json:"channel_id,nullable"`
	UserID      string                                              `json:"user_id,nullable"`
	JSON        slackChannelDataConnectionsSlackTokenConnectionJSON `json:"-"`
}

// slackChannelDataConnectionsSlackTokenConnectionJSON contains the JSON metadata
// for the struct [SlackChannelDataConnectionsSlackTokenConnection]
type slackChannelDataConnectionsSlackTokenConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlackChannelDataConnectionsSlackTokenConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slackChannelDataConnectionsSlackTokenConnectionJSON) RawJSON() string {
	return r.raw
}

func (r SlackChannelDataConnectionsSlackTokenConnection) implementsSlackChannelDataConnection() {}

// An incoming webhook Slack connection
type SlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	URL  string                                                        `json:"url,required"`
	JSON slackChannelDataConnectionsSlackIncomingWebhookConnectionJSON `json:"-"`
}

// slackChannelDataConnectionsSlackIncomingWebhookConnectionJSON contains the JSON
// metadata for the struct
// [SlackChannelDataConnectionsSlackIncomingWebhookConnection]
type slackChannelDataConnectionsSlackIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlackChannelDataConnectionsSlackIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slackChannelDataConnectionsSlackIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r SlackChannelDataConnectionsSlackIncomingWebhookConnection) implementsSlackChannelDataConnection() {
}

// A token that's used to store the access token for a Slack workspace.
type SlackChannelDataToken struct {
	AccessToken string                    `json:"access_token,required,nullable"`
	JSON        slackChannelDataTokenJSON `json:"-"`
}

// slackChannelDataTokenJSON contains the JSON metadata for the struct
// [SlackChannelDataToken]
type slackChannelDataTokenJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlackChannelDataToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slackChannelDataTokenJSON) RawJSON() string {
	return r.raw
}

// Slack channel data
type SlackChannelDataParam struct {
	Connections param.Field[[]SlackChannelDataConnectionsUnionParam] `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token param.Field[SlackChannelDataTokenParam] `json:"token"`
}

func (r SlackChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataParam) ImplementsChannelDataRequestDataUnionParam() {}

// A Slack connection, which either includes a channel_id or a user_id
type SlackChannelDataConnectionParam struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r SlackChannelDataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataConnectionParam) implementsSlackChannelDataConnectionsUnionParam() {}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by [shared.SlackChannelDataConnectionsSlackTokenConnectionParam],
// [shared.SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam],
// [SlackChannelDataConnectionParam].
type SlackChannelDataConnectionsUnionParam interface {
	implementsSlackChannelDataConnectionsUnionParam()
}

// A Slack connection, which either includes a channel_id or a user_id
type SlackChannelDataConnectionsSlackTokenConnectionParam struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r SlackChannelDataConnectionsSlackTokenConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataConnectionsSlackTokenConnectionParam) implementsSlackChannelDataConnectionsUnionParam() {
}

// An incoming webhook Slack connection
type SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam struct {
	URL param.Field[string] `json:"url,required"`
}

func (r SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam) implementsSlackChannelDataConnectionsUnionParam() {
}

// A token that's used to store the access token for a Slack workspace.
type SlackChannelDataTokenParam struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r SlackChannelDataTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A tenant to be set in the system
type TenantRequestParam struct {
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]ChannelDataRequestParam] `json:"channel_data"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]PreferenceSetRequestParam] `json:"preferences"`
	Settings    param.Field[TenantRequestSettingsParam]           `json:"settings"`
	ExtraFields map[string]interface{}                            `json:"-,extras"`
}

func (r TenantRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantRequestParam) ImplementsTenantBulkSetParamsTenantUnion() {}

func (r TenantRequestParam) ImplementsWorkflowTriggerParamsTenantUnion() {}

func (r TenantRequestParam) ImplementsScheduleNewParamsTenantUnion() {}

func (r TenantRequestParam) ImplementsScheduleUpdateParamsTenantUnion() {}

type TenantRequestSettingsParam struct {
	Branding param.Field[TenantRequestSettingsBrandingParam] `json:"branding"`
	// Set preferences for a recipient
	PreferenceSet param.Field[PreferenceSetRequestParam] `json:"preference_set"`
}

func (r TenantRequestSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantRequestSettingsBrandingParam struct {
	IconURL              param.Field[string] `json:"icon_url"`
	LogoURL              param.Field[string] `json:"logo_url"`
	PrimaryColor         param.Field[string] `json:"primary_color"`
	PrimaryColorContrast param.Field[string] `json:"primary_color_contrast"`
}

func (r TenantRequestSettingsBrandingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

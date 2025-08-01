// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"reflect"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/option"
	"github.com/tidwall/gjson"
)

// RecipientChannelDataService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientChannelDataService] method instead.
type RecipientChannelDataService struct {
	Options []option.RequestOption
}

// NewRecipientChannelDataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRecipientChannelDataService(opts ...option.RequestOption) (r *RecipientChannelDataService) {
	r = &RecipientChannelDataService{}
	r.Options = opts
	return
}

// Channel data for a given channel type.
type ChannelData struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The unique identifier for the channel.
	ChannelID string `json:"channel_id,required" format:"uuid"`
	// Channel data for a given channel type.
	Data ChannelDataData `json:"data,required"`
	// The type of provider.
	Provider ChannelDataProvider `json:"provider"`
	JSON     channelDataJSON     `json:"-"`
}

// channelDataJSON contains the JSON metadata for the struct [ChannelData]
type channelDataJSON struct {
	Typename    apijson.Field
	ChannelID   apijson.Field
	Data        apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r channelDataJSON) RawJSON() string {
	return r.raw
}

// Channel data for a given channel type.
type ChannelDataData struct {
	// This field can have the runtime type of [SlackChannelDataToken].
	Token interface{} `json:"token"`
	// This field can have the runtime type of [[]SlackChannelDataConnection],
	// [[]MsTeamsChannelDataConnection], [[]DiscordChannelDataConnection].
	Connections interface{} `json:"connections"`
	// Microsoft Teams tenant ID.
	MsTeamsTenantID string `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	// This field can have the runtime type of [[]string].
	PlayerIDs interface{} `json:"player_ids"`
	// This field can have the runtime type of [[]string].
	Tokens interface{}         `json:"tokens"`
	JSON   channelDataDataJSON `json:"-"`
	union  ChannelDataDataUnion
}

// channelDataDataJSON contains the JSON metadata for the struct [ChannelDataData]
type channelDataDataJSON struct {
	Token           apijson.Field
	Connections     apijson.Field
	MsTeamsTenantID apijson.Field
	PlayerIDs       apijson.Field
	Tokens          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r channelDataDataJSON) RawJSON() string {
	return r.raw
}

func (r *ChannelDataData) UnmarshalJSON(data []byte) (err error) {
	*r = ChannelDataData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ChannelDataDataUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [PushChannelData], [SlackChannelData],
// [MsTeamsChannelData], [DiscordChannelData], [OneSignalChannelData].
func (r ChannelDataData) AsUnion() ChannelDataDataUnion {
	return r.union
}

// Channel data for a given channel type.
//
// Union satisfied by [PushChannelData], [SlackChannelData], [MsTeamsChannelData],
// [DiscordChannelData] or [OneSignalChannelData].
type ChannelDataDataUnion interface {
	implementsChannelDataData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ChannelDataDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PushChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SlackChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MsTeamsChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DiscordChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OneSignalChannelData{}),
		},
	)
}

// The type of provider.
type ChannelDataProvider string

const (
	ChannelDataProviderPushFcm          ChannelDataProvider = "push_fcm"
	ChannelDataProviderPushApns         ChannelDataProvider = "push_apns"
	ChannelDataProviderPushExpo         ChannelDataProvider = "push_expo"
	ChannelDataProviderPushOneSignal    ChannelDataProvider = "push_one_signal"
	ChannelDataProviderChatSlack        ChannelDataProvider = "chat_slack"
	ChannelDataProviderChatMsTeams      ChannelDataProvider = "chat_ms_teams"
	ChannelDataProviderChatDiscord      ChannelDataProvider = "chat_discord"
	ChannelDataProviderHTTPKnockWebhook ChannelDataProvider = "http_knock_webhook"
)

func (r ChannelDataProvider) IsKnown() bool {
	switch r {
	case ChannelDataProviderPushFcm, ChannelDataProviderPushApns, ChannelDataProviderPushExpo, ChannelDataProviderPushOneSignal, ChannelDataProviderChatSlack, ChannelDataProviderChatMsTeams, ChannelDataProviderChatDiscord, ChannelDataProviderHTTPKnockWebhook:
		return true
	}
	return false
}

// A request to set channel data for a type of channel.
type ChannelDataRequestParam struct {
	// Channel data for a given channel type.
	Data param.Field[ChannelDataRequestDataUnionParam] `json:"data,required"`
}

func (r ChannelDataRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for a given channel type.
type ChannelDataRequestDataParam struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// Microsoft Teams tenant ID.
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r ChannelDataRequestDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChannelDataRequestDataParam) implementsChannelDataRequestDataUnionParam() {}

// Channel data for a given channel type.
//
// Satisfied by [PushChannelDataParam], [OneSignalChannelDataParam],
// [SlackChannelDataParam], [MsTeamsChannelDataParam], [DiscordChannelDataParam],
// [ChannelDataRequestDataParam].
type ChannelDataRequestDataUnionParam interface {
	implementsChannelDataRequestDataUnionParam()
}

// Discord channel data.
type DiscordChannelData struct {
	// List of Discord channel connections.
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

func (r DiscordChannelData) implementsChannelDataData() {}

// Discord channel connection, either a channel connection or an incoming webhook
// connection.
type DiscordChannelDataConnection struct {
	// Discord channel ID.
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
// [DiscordChannelDataConnectionsDiscordChannelConnection],
// [DiscordChannelDataConnectionsDiscordIncomingWebhookConnection].
func (r DiscordChannelDataConnection) AsUnion() DiscordChannelDataConnectionsUnion {
	return r.union
}

// Discord channel connection, either a channel connection or an incoming webhook
// connection.
//
// Union satisfied by [DiscordChannelDataConnectionsDiscordChannelConnection] or
// [DiscordChannelDataConnectionsDiscordIncomingWebhookConnection].
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

// Discord channel connection.
type DiscordChannelDataConnectionsDiscordChannelConnection struct {
	// Discord channel ID.
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

// Discord incoming webhook connection.
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// Discord incoming webhook object.
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

// Discord incoming webhook object.
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// Incoming webhook URL.
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

// Discord channel data.
type DiscordChannelDataParam struct {
	// List of Discord channel connections.
	Connections param.Field[[]DiscordChannelDataConnectionsUnionParam] `json:"connections,required"`
}

func (r DiscordChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

func (r DiscordChannelDataParam) implementsInlineChannelDataRequestItemUnionParam() {}

// Discord channel connection, either a channel connection or an incoming webhook
// connection.
type DiscordChannelDataConnectionParam struct {
	// Discord channel ID.
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r DiscordChannelDataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataConnectionParam) implementsDiscordChannelDataConnectionsUnionParam() {}

// Discord channel connection, either a channel connection or an incoming webhook
// connection.
//
// Satisfied by [DiscordChannelDataConnectionsDiscordChannelConnectionParam],
// [DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam],
// [DiscordChannelDataConnectionParam].
type DiscordChannelDataConnectionsUnionParam interface {
	implementsDiscordChannelDataConnectionsUnionParam()
}

// Discord channel connection.
type DiscordChannelDataConnectionsDiscordChannelConnectionParam struct {
	// Discord channel ID.
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r DiscordChannelDataConnectionsDiscordChannelConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataConnectionsDiscordChannelConnectionParam) implementsDiscordChannelDataConnectionsUnionParam() {
}

// Discord incoming webhook connection.
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam struct {
	// Discord incoming webhook object.
	IncomingWebhook param.Field[DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookParam] `json:"incoming_webhook,required"`
}

func (r DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam) implementsDiscordChannelDataConnectionsUnionParam() {
}

// Discord incoming webhook object.
type DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookParam struct {
	// Incoming webhook URL.
	URL param.Field[string] `json:"url,required"`
}

func (r DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InlineChannelDataRequestParam map[string]InlineChannelDataRequestItemUnionParam

// Channel data for a given channel type.
type InlineChannelDataRequestItemParam struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// Microsoft Teams tenant ID.
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r InlineChannelDataRequestItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InlineChannelDataRequestItemParam) implementsInlineChannelDataRequestItemUnionParam() {}

// Channel data for a given channel type.
//
// Satisfied by [PushChannelDataParam], [OneSignalChannelDataParam],
// [SlackChannelDataParam], [MsTeamsChannelDataParam], [DiscordChannelDataParam],
// [InlineChannelDataRequestItemParam].
type InlineChannelDataRequestItemUnionParam interface {
	implementsInlineChannelDataRequestItemUnionParam()
}

// Microsoft Teams channel data.
type MsTeamsChannelData struct {
	// List of Microsoft Teams connections.
	Connections []MsTeamsChannelDataConnection `json:"connections,required"`
	// Microsoft Teams tenant ID.
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

func (r MsTeamsChannelData) implementsChannelDataData() {}

// Microsoft Teams token connection.
type MsTeamsChannelDataConnection struct {
	// This field can have the runtime type of
	// [MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook].
	IncomingWebhook interface{} `json:"incoming_webhook"`
	// Microsoft Teams channel ID.
	MsTeamsChannelID string `json:"ms_teams_channel_id,nullable" format:"uuid"`
	// Microsoft Teams team ID.
	MsTeamsTeamID string `json:"ms_teams_team_id,nullable" format:"uuid"`
	// Microsoft Teams tenant ID.
	MsTeamsTenantID string `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	// Microsoft Teams user ID.
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
// [MsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection].
func (r MsTeamsChannelDataConnection) AsUnion() MsTeamsChannelDataConnectionsUnion {
	return r.union
}

// Microsoft Teams token connection.
//
// Union satisfied by [MsTeamsChannelDataConnectionsMsTeamsTokenConnection] or
// [MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection].
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

// Microsoft Teams token connection.
type MsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// Microsoft Teams channel ID.
	MsTeamsChannelID string `json:"ms_teams_channel_id,nullable" format:"uuid"`
	// Microsoft Teams team ID.
	MsTeamsTeamID string `json:"ms_teams_team_id,nullable" format:"uuid"`
	// Microsoft Teams tenant ID.
	MsTeamsTenantID string `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	// Microsoft Teams user ID.
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

// Microsoft Teams incoming webhook connection.
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// Microsoft Teams incoming webhook.
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

// Microsoft Teams incoming webhook.
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// Microsoft Teams incoming webhook URL.
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

// Microsoft Teams channel data.
type MsTeamsChannelDataParam struct {
	// List of Microsoft Teams connections.
	Connections param.Field[[]MsTeamsChannelDataConnectionsUnionParam] `json:"connections,required"`
	// Microsoft Teams tenant ID.
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r MsTeamsChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

func (r MsTeamsChannelDataParam) implementsInlineChannelDataRequestItemUnionParam() {}

// Microsoft Teams token connection.
type MsTeamsChannelDataConnectionParam struct {
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
	// Microsoft Teams channel ID.
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// Microsoft Teams team ID.
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// Microsoft Teams tenant ID.
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// Microsoft Teams user ID.
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r MsTeamsChannelDataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataConnectionParam) implementsMsTeamsChannelDataConnectionsUnionParam() {}

// Microsoft Teams token connection.
//
// Satisfied by [MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam],
// [MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam],
// [MsTeamsChannelDataConnectionParam].
type MsTeamsChannelDataConnectionsUnionParam interface {
	implementsMsTeamsChannelDataConnectionsUnionParam()
}

// Microsoft Teams token connection.
type MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam struct {
	// Microsoft Teams channel ID.
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// Microsoft Teams team ID.
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// Microsoft Teams tenant ID.
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// Microsoft Teams user ID.
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam) implementsMsTeamsChannelDataConnectionsUnionParam() {
}

// Microsoft Teams incoming webhook connection.
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam struct {
	// Microsoft Teams incoming webhook.
	IncomingWebhook param.Field[MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookParam] `json:"incoming_webhook,required"`
}

func (r MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam) implementsMsTeamsChannelDataConnectionsUnionParam() {
}

// Microsoft Teams incoming webhook.
type MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookParam struct {
	// Microsoft Teams incoming webhook URL.
	URL param.Field[string] `json:"url,required"`
}

func (r MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// OneSignal channel data.
type OneSignalChannelData struct {
	// A list of OneSignal player IDs.
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

func (r OneSignalChannelData) implementsChannelDataData() {}

// OneSignal channel data.
type OneSignalChannelDataParam struct {
	// A list of OneSignal player IDs.
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r OneSignalChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OneSignalChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

func (r OneSignalChannelDataParam) implementsInlineChannelDataRequestItemUnionParam() {}

// Push channel data.
type PushChannelData struct {
	// A list of push channel tokens.
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

func (r PushChannelData) implementsChannelDataData() {}

// Push channel data.
type PushChannelDataParam struct {
	// A list of push channel tokens.
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r PushChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PushChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

func (r PushChannelDataParam) implementsInlineChannelDataRequestItemUnionParam() {}

// Slack channel data.
type SlackChannelData struct {
	// List of Slack channel connections.
	Connections []SlackChannelDataConnection `json:"connections,required"`
	// A Slack connection token.
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

func (r SlackChannelData) implementsChannelDataData() {}

// A Slack connection, either an access token or an incoming webhook
type SlackChannelDataConnection struct {
	// A Slack access token.
	AccessToken string `json:"access_token,nullable"`
	// A Slack channel ID from the Slack provider.
	ChannelID string `json:"channel_id,nullable"`
	// The URL of the incoming webhook for a Slack connection.
	URL string `json:"url"`
	// A Slack user ID from the Slack provider.
	UserID string                         `json:"user_id,nullable"`
	JSON   slackChannelDataConnectionJSON `json:"-"`
	union  SlackChannelDataConnectionsUnion
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
// [SlackChannelDataConnectionsSlackTokenConnection],
// [SlackChannelDataConnectionsSlackIncomingWebhookConnection].
func (r SlackChannelDataConnection) AsUnion() SlackChannelDataConnectionsUnion {
	return r.union
}

// A Slack connection, either an access token or an incoming webhook
//
// Union satisfied by [SlackChannelDataConnectionsSlackTokenConnection] or
// [SlackChannelDataConnectionsSlackIncomingWebhookConnection].
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

// A Slack connection token.
type SlackChannelDataConnectionsSlackTokenConnection struct {
	// A Slack access token.
	AccessToken string `json:"access_token,nullable"`
	// A Slack channel ID from the Slack provider.
	ChannelID string `json:"channel_id,nullable"`
	// A Slack user ID from the Slack provider.
	UserID string                                              `json:"user_id,nullable"`
	JSON   slackChannelDataConnectionsSlackTokenConnectionJSON `json:"-"`
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

// A Slack connection incoming webhook.
type SlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	// The URL of the incoming webhook for a Slack connection.
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

// A Slack connection token.
type SlackChannelDataToken struct {
	// A Slack access token.
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

// Slack channel data.
type SlackChannelDataParam struct {
	// List of Slack channel connections.
	Connections param.Field[[]SlackChannelDataConnectionsUnionParam] `json:"connections,required"`
	// A Slack connection token.
	Token param.Field[SlackChannelDataTokenParam] `json:"token"`
}

func (r SlackChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

func (r SlackChannelDataParam) implementsInlineChannelDataRequestItemUnionParam() {}

// A Slack connection, either an access token or an incoming webhook
type SlackChannelDataConnectionParam struct {
	// A Slack access token.
	AccessToken param.Field[string] `json:"access_token"`
	// A Slack channel ID from the Slack provider.
	ChannelID param.Field[string] `json:"channel_id"`
	// The URL of the incoming webhook for a Slack connection.
	URL param.Field[string] `json:"url"`
	// A Slack user ID from the Slack provider.
	UserID param.Field[string] `json:"user_id"`
}

func (r SlackChannelDataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataConnectionParam) implementsSlackChannelDataConnectionsUnionParam() {}

// A Slack connection, either an access token or an incoming webhook
//
// Satisfied by [SlackChannelDataConnectionsSlackTokenConnectionParam],
// [SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam],
// [SlackChannelDataConnectionParam].
type SlackChannelDataConnectionsUnionParam interface {
	implementsSlackChannelDataConnectionsUnionParam()
}

// A Slack connection token.
type SlackChannelDataConnectionsSlackTokenConnectionParam struct {
	// A Slack access token.
	AccessToken param.Field[string] `json:"access_token"`
	// A Slack channel ID from the Slack provider.
	ChannelID param.Field[string] `json:"channel_id"`
	// A Slack user ID from the Slack provider.
	UserID param.Field[string] `json:"user_id"`
}

func (r SlackChannelDataConnectionsSlackTokenConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataConnectionsSlackTokenConnectionParam) implementsSlackChannelDataConnectionsUnionParam() {
}

// A Slack connection incoming webhook.
type SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam struct {
	// The URL of the incoming webhook for a Slack connection.
	URL param.Field[string] `json:"url,required"`
}

func (r SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam) implementsSlackChannelDataConnectionsUnionParam() {
}

// A Slack connection token.
type SlackChannelDataTokenParam struct {
	// A Slack access token.
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r SlackChannelDataTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"reflect"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
	"github.com/tidwall/gjson"
)

// RecipientService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientService] method instead.
type RecipientService struct {
	Options []option.RequestOption
}

// NewRecipientService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecipientService(opts ...option.RequestOption) (r *RecipientService) {
	r = &RecipientService{}
	r.Options = opts
	return
}

// Channel data for various channel types
type ChannelData struct {
	Typename  string `json:"__typename,required"`
	ChannelID string `json:"channel_id,required" format:"uuid"`
	// Channel data for push providers
	Data ChannelDataData `json:"data,required"`
	JSON channelDataJSON `json:"-"`
}

// channelDataJSON contains the JSON metadata for the struct [ChannelData]
type channelDataJSON struct {
	Typename    apijson.Field
	ChannelID   apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChannelData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r channelDataJSON) RawJSON() string {
	return r.raw
}

// Channel data for push providers
type ChannelDataData struct {
	// This field can have the runtime type of [SlackChannelDataToken].
	Token interface{} `json:"token"`
	// This field can have the runtime type of [[]SlackChannelDataConnection],
	// [[]MsTeamsChannelDataConnection], [[]DiscordChannelDataConnection].
	Connections interface{} `json:"connections"`
	// The Microsoft Teams tenant ID
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

// Channel data for push providers
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

func (r ChannelDataRequestDataParam) implementsChannelDataRequestDataUnionParam() {}

// Channel data for push providers
//
// Satisfied by [PushChannelDataParam], [OneSignalChannelDataParam],
// [SlackChannelDataParam], [MsTeamsChannelDataParam], [DiscordChannelDataParam],
// [ChannelDataRequestDataParam].
type ChannelDataRequestDataUnionParam interface {
	implementsChannelDataRequestDataUnionParam()
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

func (r DiscordChannelData) implementsChannelDataData() {}

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
// [DiscordChannelDataConnectionsDiscordChannelConnection],
// [DiscordChannelDataConnectionsDiscordIncomingWebhookConnection].
func (r DiscordChannelDataConnection) AsUnion() DiscordChannelDataConnectionsUnion {
	return r.union
}

// Discord channel connection
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

func (r DiscordChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

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
// Satisfied by [DiscordChannelDataConnectionsDiscordChannelConnectionParam],
// [DiscordChannelDataConnectionsDiscordIncomingWebhookConnectionParam],
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

type InlineChannelDataRequestParam map[string]ChannelDataRequestParam

type InlinePreferenceSetRequestParam map[string]PreferenceSetRequestParam

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

func (r MsTeamsChannelData) implementsChannelDataData() {}

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
// [MsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection].
func (r MsTeamsChannelDataConnection) AsUnion() MsTeamsChannelDataConnectionsUnion {
	return r.union
}

// Microsoft Teams token connection
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

func (r MsTeamsChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

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
// Satisfied by [MsTeamsChannelDataConnectionsMsTeamsTokenConnectionParam],
// [MsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionParam],
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

func (r OneSignalChannelData) implementsChannelDataData() {}

// OneSignal channel data
type OneSignalChannelDataParam struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r OneSignalChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OneSignalChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

// A preference set object.
type PreferenceSet struct {
	ID       string `json:"id,required"`
	Typename string `json:"__typename,required"`
	// A map of categories and their settings
	Categories map[string]PreferenceSetCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// A map of workflows and their settings
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
	// Channel type preferences
	ChannelTypes PreferenceSetChannelTypes                                             `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                    `json:"conditions,nullable"`
	JSON         preferenceSetCategoriesPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
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
	// Channel type preferences
	ChannelTypes PreferenceSetChannelTypes                                            `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                   `json:"conditions,nullable"`
	JSON         preferenceSetWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
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

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
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

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
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

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
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

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
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

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
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
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesChatUnionParam interface {
	ImplementsPreferenceSetChannelTypesChatUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesEmailUnionParam interface {
	ImplementsPreferenceSetChannelTypesEmailUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesHTTPUnionParam interface {
	ImplementsPreferenceSetChannelTypesHTTPUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesInAppFeedUnionParam interface {
	ImplementsPreferenceSetChannelTypesInAppFeedUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesPushUnionParam interface {
	ImplementsPreferenceSetChannelTypesPushUnionParam()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool], [PreferenceSetChannelTypeSettingParam].
type PreferenceSetChannelTypesSMSUnionParam interface {
	ImplementsPreferenceSetChannelTypesSMSUnionParam()
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
// [PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam].
type PreferenceSetRequestCategoriesUnionParam interface {
	ImplementsPreferenceSetRequestCategoriesUnionParam()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam struct {
	// Channel type preferences
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	Conditions   param.Field[[]shared.ConditionParam]        `json:"conditions"`
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
	// Channel type preferences
	ChannelTypes param.Field[PreferenceSetChannelTypesParam] `json:"channel_types"`
	Conditions   param.Field[[]shared.ConditionParam]        `json:"conditions"`
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

func (r PushChannelData) implementsChannelDataData() {}

// Channel data for push providers
type PushChannelDataParam struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r PushChannelDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PushChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

// A recipient, which is either a user or an object
type Recipient struct {
	ID          string        `json:"id,required"`
	Typename    string        `json:"__typename,required"`
	UpdatedAt   time.Time     `json:"updated_at,required" format:"date-time"`
	Avatar      string        `json:"avatar,nullable"`
	Collection  string        `json:"collection"`
	CreatedAt   time.Time     `json:"created_at,nullable" format:"date-time"`
	Email       string        `json:"email,nullable"`
	Name        string        `json:"name,nullable"`
	PhoneNumber string        `json:"phone_number,nullable"`
	Timezone    string        `json:"timezone,nullable"`
	JSON        recipientJSON `json:"-"`
	union       RecipientUnion
}

// recipientJSON contains the JSON metadata for the struct [Recipient]
type recipientJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r recipientJSON) RawJSON() string {
	return r.raw
}

func (r *Recipient) UnmarshalJSON(data []byte) (err error) {
	*r = Recipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecipientUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [User], [Object].
func (r Recipient) AsUnion() RecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [Object].
type RecipientUnion interface {
	implementsRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Object{}),
		},
	)
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type RecipientRequestParam struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	Collection  param.Field[string]                        `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
}

func (r RecipientRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecipientRequestParam) ImplementsRecipientRequestUnionParam() {}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// Satisfied by [shared.UnionString], [InlineIdentifyUserRequestParam],
// [InlineObjectRequestParam], [RecipientRequestParam].
type RecipientRequestUnionParam interface {
	ImplementsRecipientRequestUnionParam()
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

func (r SlackChannelData) implementsChannelDataData() {}

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
// [SlackChannelDataConnectionsSlackTokenConnection],
// [SlackChannelDataConnectionsSlackIncomingWebhookConnection].
func (r SlackChannelDataConnection) AsUnion() SlackChannelDataConnectionsUnion {
	return r.union
}

// A Slack connection, which either includes a channel_id or a user_id
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

func (r SlackChannelDataParam) implementsChannelDataRequestDataUnionParam() {}

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
// Satisfied by [SlackChannelDataConnectionsSlackTokenConnectionParam],
// [SlackChannelDataConnectionsSlackIncomingWebhookConnectionParam],
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

// A subscription object
type Subscription struct {
	Typename   string    `json:"__typename,required"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A custom-object entity which belongs to a collection.
	Object Object `json:"object,required"`
	// A recipient, which is either a user or an object
	Recipient Recipient `json:"recipient,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The custom properties associated with the subscription
	Properties map[string]interface{} `json:"properties,nullable"`
	JSON       subscriptionJSON       `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	Typename    apijson.Field
	InsertedAt  apijson.Field
	Object      apijson.Field
	Recipient   apijson.Field
	UpdatedAt   apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

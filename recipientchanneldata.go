// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"reflect"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/option"
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
	ChannelID string                           `json:"channel_id"`
	URL       string                           `json:"url"`
	JSON      discordChannelDataConnectionJSON `json:"-"`
	union     DiscordChannelDataConnectionsUnion
}

// discordChannelDataConnectionJSON contains the JSON metadata for the struct
// [DiscordChannelDataConnection]
type discordChannelDataConnectionJSON struct {
	ChannelID   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// [DiscordChannelDataConnectionsChannelConnection],
// [DiscordChannelDataConnectionsIncomingWebhookConnection].
func (r DiscordChannelDataConnection) AsUnion() DiscordChannelDataConnectionsUnion {
	return r.union
}

// Discord channel connection
//
// Union satisfied by [DiscordChannelDataConnectionsChannelConnection] or
// [DiscordChannelDataConnectionsIncomingWebhookConnection].
type DiscordChannelDataConnectionsUnion interface {
	implementsDiscordChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DiscordChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DiscordChannelDataConnectionsChannelConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DiscordChannelDataConnectionsIncomingWebhookConnection{}),
		},
	)
}

// Discord channel connection
type DiscordChannelDataConnectionsChannelConnection struct {
	// The Discord channel ID
	ChannelID string                                             `json:"channel_id,required"`
	JSON      discordChannelDataConnectionsChannelConnectionJSON `json:"-"`
}

// discordChannelDataConnectionsChannelConnectionJSON contains the JSON metadata
// for the struct [DiscordChannelDataConnectionsChannelConnection]
type discordChannelDataConnectionsChannelConnectionJSON struct {
	ChannelID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscordChannelDataConnectionsChannelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discordChannelDataConnectionsChannelConnectionJSON) RawJSON() string {
	return r.raw
}

func (r DiscordChannelDataConnectionsChannelConnection) implementsDiscordChannelDataConnection() {}

// An incoming webhook Slack connection
type DiscordChannelDataConnectionsIncomingWebhookConnection struct {
	URL  string                                                     `json:"url,required"`
	JSON discordChannelDataConnectionsIncomingWebhookConnectionJSON `json:"-"`
}

// discordChannelDataConnectionsIncomingWebhookConnectionJSON contains the JSON
// metadata for the struct [DiscordChannelDataConnectionsIncomingWebhookConnection]
type discordChannelDataConnectionsIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscordChannelDataConnectionsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discordChannelDataConnectionsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r DiscordChannelDataConnectionsIncomingWebhookConnection) implementsDiscordChannelDataConnection() {
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
	ChannelID param.Field[string] `json:"channel_id"`
	URL       param.Field[string] `json:"url"`
}

func (r DiscordChannelDataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataConnectionParam) implementsDiscordChannelDataConnectionsUnionParam() {}

// Discord channel connection
//
// Satisfied by [DiscordChannelDataConnectionsChannelConnectionParam],
// [DiscordChannelDataConnectionsIncomingWebhookConnectionParam],
// [DiscordChannelDataConnectionParam].
type DiscordChannelDataConnectionsUnionParam interface {
	implementsDiscordChannelDataConnectionsUnionParam()
}

// Discord channel connection
type DiscordChannelDataConnectionsChannelConnectionParam struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r DiscordChannelDataConnectionsChannelConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataConnectionsChannelConnectionParam) implementsDiscordChannelDataConnectionsUnionParam() {
}

// An incoming webhook Slack connection
type DiscordChannelDataConnectionsIncomingWebhookConnectionParam struct {
	URL param.Field[string] `json:"url,required"`
}

func (r DiscordChannelDataConnectionsIncomingWebhookConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DiscordChannelDataConnectionsIncomingWebhookConnectionParam) implementsDiscordChannelDataConnectionsUnionParam() {
}

type InlineChannelDataRequestParam map[string]ChannelDataRequestParam

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

// A Slack connection, which either includes a channel_id or a user_id
type MsTeamsChannelDataConnection struct {
	AccessToken string                           `json:"access_token,nullable"`
	ChannelID   string                           `json:"channel_id,nullable"`
	URL         string                           `json:"url"`
	UserID      string                           `json:"user_id,nullable"`
	JSON        msTeamsChannelDataConnectionJSON `json:"-"`
	union       MsTeamsChannelDataConnectionsUnion
}

// msTeamsChannelDataConnectionJSON contains the JSON metadata for the struct
// [MsTeamsChannelDataConnection]
type msTeamsChannelDataConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	URL         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// [MsTeamsChannelDataConnectionsTokenConnection],
// [MsTeamsChannelDataConnectionsIncomingWebhookConnection].
func (r MsTeamsChannelDataConnection) AsUnion() MsTeamsChannelDataConnectionsUnion {
	return r.union
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Union satisfied by [MsTeamsChannelDataConnectionsTokenConnection] or
// [MsTeamsChannelDataConnectionsIncomingWebhookConnection].
type MsTeamsChannelDataConnectionsUnion interface {
	implementsMsTeamsChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MsTeamsChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MsTeamsChannelDataConnectionsTokenConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MsTeamsChannelDataConnectionsIncomingWebhookConnection{}),
		},
	)
}

// A Slack connection, which either includes a channel_id or a user_id
type MsTeamsChannelDataConnectionsTokenConnection struct {
	AccessToken string                                           `json:"access_token,nullable"`
	ChannelID   string                                           `json:"channel_id,nullable"`
	UserID      string                                           `json:"user_id,nullable"`
	JSON        msTeamsChannelDataConnectionsTokenConnectionJSON `json:"-"`
}

// msTeamsChannelDataConnectionsTokenConnectionJSON contains the JSON metadata for
// the struct [MsTeamsChannelDataConnectionsTokenConnection]
type msTeamsChannelDataConnectionsTokenConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MsTeamsChannelDataConnectionsTokenConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r msTeamsChannelDataConnectionsTokenConnectionJSON) RawJSON() string {
	return r.raw
}

func (r MsTeamsChannelDataConnectionsTokenConnection) implementsMsTeamsChannelDataConnection() {}

// An incoming webhook Slack connection
type MsTeamsChannelDataConnectionsIncomingWebhookConnection struct {
	URL  string                                                     `json:"url,required"`
	JSON msTeamsChannelDataConnectionsIncomingWebhookConnectionJSON `json:"-"`
}

// msTeamsChannelDataConnectionsIncomingWebhookConnectionJSON contains the JSON
// metadata for the struct [MsTeamsChannelDataConnectionsIncomingWebhookConnection]
type msTeamsChannelDataConnectionsIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MsTeamsChannelDataConnectionsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r msTeamsChannelDataConnectionsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r MsTeamsChannelDataConnectionsIncomingWebhookConnection) implementsMsTeamsChannelDataConnection() {
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

// A Slack connection, which either includes a channel_id or a user_id
type MsTeamsChannelDataConnectionParam struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r MsTeamsChannelDataConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataConnectionParam) implementsMsTeamsChannelDataConnectionsUnionParam() {}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by [MsTeamsChannelDataConnectionsTokenConnectionParam],
// [MsTeamsChannelDataConnectionsIncomingWebhookConnectionParam],
// [MsTeamsChannelDataConnectionParam].
type MsTeamsChannelDataConnectionsUnionParam interface {
	implementsMsTeamsChannelDataConnectionsUnionParam()
}

// A Slack connection, which either includes a channel_id or a user_id
type MsTeamsChannelDataConnectionsTokenConnectionParam struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r MsTeamsChannelDataConnectionsTokenConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataConnectionsTokenConnectionParam) implementsMsTeamsChannelDataConnectionsUnionParam() {
}

// An incoming webhook Slack connection
type MsTeamsChannelDataConnectionsIncomingWebhookConnectionParam struct {
	URL param.Field[string] `json:"url,required"`
}

func (r MsTeamsChannelDataConnectionsIncomingWebhookConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MsTeamsChannelDataConnectionsIncomingWebhookConnectionParam) implementsMsTeamsChannelDataConnectionsUnionParam() {
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

// Slack channel data
type SlackChannelData struct {
	Connections []SlackChannelDataConnection `json:"connections,required"`
	Token       SlackChannelDataToken        `json:"token,nullable"`
	JSON        slackChannelDataJSON         `json:"-"`
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
// [SlackChannelDataConnectionsTokenConnection],
// [SlackChannelDataConnectionsIncomingWebhookConnection].
func (r SlackChannelDataConnection) AsUnion() SlackChannelDataConnectionsUnion {
	return r.union
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Union satisfied by [SlackChannelDataConnectionsTokenConnection] or
// [SlackChannelDataConnectionsIncomingWebhookConnection].
type SlackChannelDataConnectionsUnion interface {
	implementsSlackChannelDataConnection()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SlackChannelDataConnectionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SlackChannelDataConnectionsTokenConnection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SlackChannelDataConnectionsIncomingWebhookConnection{}),
		},
	)
}

// A Slack connection, which either includes a channel_id or a user_id
type SlackChannelDataConnectionsTokenConnection struct {
	AccessToken string                                         `json:"access_token,nullable"`
	ChannelID   string                                         `json:"channel_id,nullable"`
	UserID      string                                         `json:"user_id,nullable"`
	JSON        slackChannelDataConnectionsTokenConnectionJSON `json:"-"`
}

// slackChannelDataConnectionsTokenConnectionJSON contains the JSON metadata for
// the struct [SlackChannelDataConnectionsTokenConnection]
type slackChannelDataConnectionsTokenConnectionJSON struct {
	AccessToken apijson.Field
	ChannelID   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlackChannelDataConnectionsTokenConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slackChannelDataConnectionsTokenConnectionJSON) RawJSON() string {
	return r.raw
}

func (r SlackChannelDataConnectionsTokenConnection) implementsSlackChannelDataConnection() {}

// An incoming webhook Slack connection
type SlackChannelDataConnectionsIncomingWebhookConnection struct {
	URL  string                                                   `json:"url,required"`
	JSON slackChannelDataConnectionsIncomingWebhookConnectionJSON `json:"-"`
}

// slackChannelDataConnectionsIncomingWebhookConnectionJSON contains the JSON
// metadata for the struct [SlackChannelDataConnectionsIncomingWebhookConnection]
type slackChannelDataConnectionsIncomingWebhookConnectionJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlackChannelDataConnectionsIncomingWebhookConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slackChannelDataConnectionsIncomingWebhookConnectionJSON) RawJSON() string {
	return r.raw
}

func (r SlackChannelDataConnectionsIncomingWebhookConnection) implementsSlackChannelDataConnection() {
}

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
	Token       param.Field[SlackChannelDataTokenParam]              `json:"token"`
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
// Satisfied by [SlackChannelDataConnectionsTokenConnectionParam],
// [SlackChannelDataConnectionsIncomingWebhookConnectionParam],
// [SlackChannelDataConnectionParam].
type SlackChannelDataConnectionsUnionParam interface {
	implementsSlackChannelDataConnectionsUnionParam()
}

// A Slack connection, which either includes a channel_id or a user_id
type SlackChannelDataConnectionsTokenConnectionParam struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r SlackChannelDataConnectionsTokenConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataConnectionsTokenConnectionParam) implementsSlackChannelDataConnectionsUnionParam() {
}

// An incoming webhook Slack connection
type SlackChannelDataConnectionsIncomingWebhookConnectionParam struct {
	URL param.Field[string] `json:"url,required"`
}

func (r SlackChannelDataConnectionsIncomingWebhookConnectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SlackChannelDataConnectionsIncomingWebhookConnectionParam) implementsSlackChannelDataConnectionsUnionParam() {
}

type SlackChannelDataTokenParam struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r SlackChannelDataTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

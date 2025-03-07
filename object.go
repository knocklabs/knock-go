// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/packages/pagination"
	"github.com/stainless-sdks/knock-go/shared"
	"github.com/tidwall/gjson"
)

// ObjectService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectService] method instead.
type ObjectService struct {
	Options []option.RequestOption
	Bulk    *ObjectBulkService
}

// NewObjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewObjectService(opts ...option.RequestOption) (r *ObjectService) {
	r = &ObjectService{}
	r.Options = opts
	r.Bulk = NewObjectBulkService(opts...)
	return
}

// List objects in a collection
func (r *ObjectService) List(ctx context.Context, collection string, query ObjectListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[shared.Object], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s", collection)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List objects in a collection
func (r *ObjectService) ListAutoPaging(ctx context.Context, collection string, query ObjectListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[shared.Object] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, collection, query, opts...))
}

// Delete an object
func (r *ObjectService) Delete(ctx context.Context, collection string, id string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s", collection, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add subscriptions for an object. If a subscription already exists, it will be
// updated.
func (r *ObjectService) AddSubscriptions(ctx context.Context, collection string, objectID string, body ObjectAddSubscriptionsParams, opts ...option.RequestOption) (res *[]ObjectAddSubscriptionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required object_id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/subscriptions", collection, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete subscriptions
func (r *ObjectService) DeleteSubscriptions(ctx context.Context, collection string, objectID string, body ObjectDeleteSubscriptionsParams, opts ...option.RequestOption) (res *[]ObjectDeleteSubscriptionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required object_id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/subscriptions", collection, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Get an object
func (r *ObjectService) Get(ctx context.Context, collection string, id string, opts ...option.RequestOption) (res *shared.Object, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s", collection, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get channel data
func (r *ObjectService) GetChannelData(ctx context.Context, collection string, objectID string, channelID string, opts ...option.RequestOption) (res *ObjectGetChannelDataResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required object_id parameter")
		return
	}
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/channel_data/%s", collection, objectID, channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a preference set
func (r *ObjectService) GetPreferences(ctx context.Context, collection string, objectID string, id string, query ObjectGetPreferencesParams, opts ...option.RequestOption) (res *ObjectGetPreferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required object_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/preferences/%s", collection, objectID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List messages
func (r *ObjectService) ListMessages(ctx context.Context, collection string, id string, query ObjectListMessagesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[ObjectListMessagesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/messages", collection, id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List messages
func (r *ObjectService) ListMessagesAutoPaging(ctx context.Context, collection string, id string, query ObjectListMessagesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[ObjectListMessagesResponse] {
	return pagination.NewEntriesCursorAutoPager(r.ListMessages(ctx, collection, id, query, opts...))
}

// List preference sets
func (r *ObjectService) ListPreferences(ctx context.Context, collection string, objectID string, opts ...option.RequestOption) (res *[]ObjectListPreferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required object_id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/preferences", collection, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List schedules
func (r *ObjectService) ListSchedules(ctx context.Context, collection string, id string, query ObjectListSchedulesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[ObjectListSchedulesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/schedules", collection, id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List schedules
func (r *ObjectService) ListSchedulesAutoPaging(ctx context.Context, collection string, id string, query ObjectListSchedulesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[ObjectListSchedulesResponse] {
	return pagination.NewEntriesCursorAutoPager(r.ListSchedules(ctx, collection, id, query, opts...))
}

// List subscriptions for an object. Either list all subscriptions that belong to
// the object, or all subscriptions that this object has. Determined by the `mode`
// query parameter.
func (r *ObjectService) ListSubscriptions(ctx context.Context, collection string, objectID string, query ObjectListSubscriptionsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[ObjectListSubscriptionsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required object_id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/subscriptions", collection, objectID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List subscriptions for an object. Either list all subscriptions that belong to
// the object, or all subscriptions that this object has. Determined by the `mode`
// query parameter.
func (r *ObjectService) ListSubscriptionsAutoPaging(ctx context.Context, collection string, objectID string, query ObjectListSubscriptionsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[ObjectListSubscriptionsResponse] {
	return pagination.NewEntriesCursorAutoPager(r.ListSubscriptions(ctx, collection, objectID, query, opts...))
}

// Set (identify) an object
func (r *ObjectService) Set(ctx context.Context, collection string, id string, body ObjectSetParams, opts ...option.RequestOption) (res *shared.Object, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s", collection, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Set channel data
func (r *ObjectService) SetChannelData(ctx context.Context, collection string, objectID string, channelID string, body ObjectSetChannelDataParams, opts ...option.RequestOption) (res *ObjectSetChannelDataResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required object_id parameter")
		return
	}
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/channel_data/%s", collection, objectID, channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Update a preference set
func (r *ObjectService) SetPreferences(ctx context.Context, collection string, objectID string, id string, body ObjectSetPreferencesParams, opts ...option.RequestOption) (res *ObjectSetPreferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required object_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/preferences/%s", collection, objectID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Unset channel data
func (r *ObjectService) UnsetChannelData(ctx context.Context, collection string, objectID string, channelID string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required object_id parameter")
		return
	}
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/%s/channel_data/%s", collection, objectID, channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A subscription object
type ObjectAddSubscriptionsResponse struct {
	Typename   string    `json:"__typename,required"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A custom-object entity which belongs to a collection.
	Object shared.Object `json:"object,required"`
	// A recipient, which is either a user or an object
	Recipient ObjectAddSubscriptionsResponseRecipient `json:"recipient,required"`
	UpdatedAt time.Time                               `json:"updated_at,required" format:"date-time"`
	// The custom properties associated with the subscription
	Properties map[string]interface{}             `json:"properties,nullable"`
	JSON       objectAddSubscriptionsResponseJSON `json:"-"`
}

// objectAddSubscriptionsResponseJSON contains the JSON metadata for the struct
// [ObjectAddSubscriptionsResponse]
type objectAddSubscriptionsResponseJSON struct {
	Typename    apijson.Field
	InsertedAt  apijson.Field
	Object      apijson.Field
	Recipient   apijson.Field
	UpdatedAt   apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectAddSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectAddSubscriptionsResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type ObjectAddSubscriptionsResponseRecipient struct {
	ID          string                                      `json:"id,required"`
	Typename    string                                      `json:"__typename,required"`
	UpdatedAt   time.Time                                   `json:"updated_at,required" format:"date-time"`
	Avatar      string                                      `json:"avatar,nullable"`
	Collection  string                                      `json:"collection"`
	CreatedAt   time.Time                                   `json:"created_at,nullable" format:"date-time"`
	Email       string                                      `json:"email,nullable"`
	Name        string                                      `json:"name,nullable"`
	PhoneNumber string                                      `json:"phone_number,nullable"`
	Timezone    string                                      `json:"timezone,nullable"`
	JSON        objectAddSubscriptionsResponseRecipientJSON `json:"-"`
	union       ObjectAddSubscriptionsResponseRecipientUnion
}

// objectAddSubscriptionsResponseRecipientJSON contains the JSON metadata for the
// struct [ObjectAddSubscriptionsResponseRecipient]
type objectAddSubscriptionsResponseRecipientJSON struct {
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

func (r objectAddSubscriptionsResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *ObjectAddSubscriptionsResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = ObjectAddSubscriptionsResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObjectAddSubscriptionsResponseRecipientUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r ObjectAddSubscriptionsResponseRecipient) AsUnion() ObjectAddSubscriptionsResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ObjectAddSubscriptionsResponseRecipientUnion interface {
	ImplementsObjectAddSubscriptionsResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectAddSubscriptionsResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A subscription object
type ObjectDeleteSubscriptionsResponse struct {
	Typename   string    `json:"__typename,required"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A custom-object entity which belongs to a collection.
	Object shared.Object `json:"object,required"`
	// A recipient, which is either a user or an object
	Recipient ObjectDeleteSubscriptionsResponseRecipient `json:"recipient,required"`
	UpdatedAt time.Time                                  `json:"updated_at,required" format:"date-time"`
	// The custom properties associated with the subscription
	Properties map[string]interface{}                `json:"properties,nullable"`
	JSON       objectDeleteSubscriptionsResponseJSON `json:"-"`
}

// objectDeleteSubscriptionsResponseJSON contains the JSON metadata for the struct
// [ObjectDeleteSubscriptionsResponse]
type objectDeleteSubscriptionsResponseJSON struct {
	Typename    apijson.Field
	InsertedAt  apijson.Field
	Object      apijson.Field
	Recipient   apijson.Field
	UpdatedAt   apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectDeleteSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectDeleteSubscriptionsResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type ObjectDeleteSubscriptionsResponseRecipient struct {
	ID          string                                         `json:"id,required"`
	Typename    string                                         `json:"__typename,required"`
	UpdatedAt   time.Time                                      `json:"updated_at,required" format:"date-time"`
	Avatar      string                                         `json:"avatar,nullable"`
	Collection  string                                         `json:"collection"`
	CreatedAt   time.Time                                      `json:"created_at,nullable" format:"date-time"`
	Email       string                                         `json:"email,nullable"`
	Name        string                                         `json:"name,nullable"`
	PhoneNumber string                                         `json:"phone_number,nullable"`
	Timezone    string                                         `json:"timezone,nullable"`
	JSON        objectDeleteSubscriptionsResponseRecipientJSON `json:"-"`
	union       ObjectDeleteSubscriptionsResponseRecipientUnion
}

// objectDeleteSubscriptionsResponseRecipientJSON contains the JSON metadata for
// the struct [ObjectDeleteSubscriptionsResponseRecipient]
type objectDeleteSubscriptionsResponseRecipientJSON struct {
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

func (r objectDeleteSubscriptionsResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *ObjectDeleteSubscriptionsResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = ObjectDeleteSubscriptionsResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObjectDeleteSubscriptionsResponseRecipientUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r ObjectDeleteSubscriptionsResponseRecipient) AsUnion() ObjectDeleteSubscriptionsResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ObjectDeleteSubscriptionsResponseRecipientUnion interface {
	ImplementsObjectDeleteSubscriptionsResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectDeleteSubscriptionsResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// Channel data for various channel types
type ObjectGetChannelDataResponse struct {
	Typename  string `json:"__typename,required"`
	ChannelID string `json:"channel_id,required" format:"uuid"`
	// Channel data for push providers
	Data ObjectGetChannelDataResponseData `json:"data,required"`
	JSON objectGetChannelDataResponseJSON `json:"-"`
}

// objectGetChannelDataResponseJSON contains the JSON metadata for the struct
// [ObjectGetChannelDataResponse]
type objectGetChannelDataResponseJSON struct {
	Typename    apijson.Field
	ChannelID   apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectGetChannelDataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectGetChannelDataResponseJSON) RawJSON() string {
	return r.raw
}

// Channel data for push providers
type ObjectGetChannelDataResponseData struct {
	// This field can have the runtime type of [shared.SlackChannelDataToken].
	Token interface{} `json:"token"`
	// This field can have the runtime type of [[]shared.SlackChannelDataConnection],
	// [[]shared.MsTeamsChannelDataConnection],
	// [[]shared.DiscordChannelDataConnection].
	Connections interface{} `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID string `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	// This field can have the runtime type of [[]string].
	PlayerIDs interface{} `json:"player_ids"`
	// This field can have the runtime type of [[]string].
	Tokens interface{}                          `json:"tokens"`
	JSON   objectGetChannelDataResponseDataJSON `json:"-"`
	union  ObjectGetChannelDataResponseDataUnion
}

// objectGetChannelDataResponseDataJSON contains the JSON metadata for the struct
// [ObjectGetChannelDataResponseData]
type objectGetChannelDataResponseDataJSON struct {
	Token           apijson.Field
	Connections     apijson.Field
	MsTeamsTenantID apijson.Field
	PlayerIDs       apijson.Field
	Tokens          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r objectGetChannelDataResponseDataJSON) RawJSON() string {
	return r.raw
}

func (r *ObjectGetChannelDataResponseData) UnmarshalJSON(data []byte) (err error) {
	*r = ObjectGetChannelDataResponseData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObjectGetChannelDataResponseDataUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.PushChannelData],
// [shared.SlackChannelData], [shared.MsTeamsChannelData],
// [shared.DiscordChannelData], [shared.OneSignalChannelData].
func (r ObjectGetChannelDataResponseData) AsUnion() ObjectGetChannelDataResponseDataUnion {
	return r.union
}

// Channel data for push providers
//
// Union satisfied by [shared.PushChannelData], [shared.SlackChannelData],
// [shared.MsTeamsChannelData], [shared.DiscordChannelData] or
// [shared.OneSignalChannelData].
type ObjectGetChannelDataResponseDataUnion interface {
	ImplementsObjectGetChannelDataResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectGetChannelDataResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.PushChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.SlackChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.MsTeamsChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.DiscordChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.OneSignalChannelData{}),
		},
	)
}

// A preference set object.
type ObjectGetPreferencesResponse struct {
	ID       string `json:"id,required"`
	Typename string `json:"__typename,required"`
	// A map of categories and their settings
	Categories map[string]ObjectGetPreferencesResponseCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// A map of workflows and their settings
	Workflows map[string]ObjectGetPreferencesResponseWorkflowsUnion `json:"workflows,nullable"`
	JSON      objectGetPreferencesResponseJSON                      `json:"-"`
}

// objectGetPreferencesResponseJSON contains the JSON metadata for the struct
// [ObjectGetPreferencesResponse]
type objectGetPreferencesResponseJSON struct {
	ID           apijson.Field
	Typename     apijson.Field
	Categories   apijson.Field
	ChannelTypes apijson.Field
	Workflows    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObjectGetPreferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectGetPreferencesResponseJSON) RawJSON() string {
	return r.raw
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [ObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject].
type ObjectGetPreferencesResponseCategoriesUnion interface {
	ImplementsObjectGetPreferencesResponseCategoriesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectGetPreferencesResponseCategoriesUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(ObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                     `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                   `json:"conditions,nullable"`
	JSON         objectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// objectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [ObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject]
type objectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObjectGetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsObjectGetPreferencesResponseCategoriesUnion() {
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [ObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject].
type ObjectGetPreferencesResponseWorkflowsUnion interface {
	ImplementsObjectGetPreferencesResponseWorkflowsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectGetPreferencesResponseWorkflowsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(ObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                    `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                  `json:"conditions,nullable"`
	JSON         objectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// objectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [ObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject]
type objectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObjectGetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsObjectGetPreferencesResponseWorkflowsUnion() {
}

// Represents a single message that was generated by a workflow for a given
// channel.
type ObjectListMessagesResponse struct {
	// The message ID
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A list of actor representations associated with the message (up to 10)
	Actors []ObjectListMessagesResponseActorsUnion `json:"actors"`
	// Timestamp when message was archived
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Channel ID associated with the message
	ChannelID string `json:"channel_id" format:"uuid"`
	// Timestamp when message was clicked
	ClickedAt time.Time `json:"clicked_at,nullable" format:"date-time"`
	// Additional message data
	Data map[string]interface{} `json:"data,nullable"`
	// List of engagement statuses
	EngagementStatuses []ObjectListMessagesResponseEngagementStatus `json:"engagement_statuses"`
	// Timestamp of creation
	InsertedAt time.Time `json:"inserted_at" format:"date-time"`
	// Timestamp when message was interacted with
	InteractedAt time.Time `json:"interacted_at,nullable" format:"date-time"`
	// Timestamp when a link in the message was clicked
	LinkClickedAt time.Time `json:"link_clicked_at,nullable" format:"date-time"`
	// Message metadata
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// Timestamp when message was read
	ReadAt time.Time `json:"read_at,nullable" format:"date-time"`
	// A reference to a recipient, either a user identifier (string) or an object
	// reference (id, collection).
	Recipient ObjectListMessagesResponseRecipientUnion `json:"recipient"`
	// Timestamp when message was scheduled for
	ScheduledAt time.Time `json:"scheduled_at,nullable" format:"date-time"`
	// Timestamp when message was seen
	SeenAt time.Time `json:"seen_at,nullable" format:"date-time"`
	// Source information
	Source ObjectListMessagesResponseSource `json:"source"`
	// Message delivery status
	Status ObjectListMessagesResponseStatus `json:"status"`
	// Tenant ID that the message belongs to
	Tenant string `json:"tenant,nullable"`
	// Timestamp of last update
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Workflow key used to create the message
	//
	// Deprecated: deprecated
	Workflow string                         `json:"workflow,nullable"`
	JSON     objectListMessagesResponseJSON `json:"-"`
}

// objectListMessagesResponseJSON contains the JSON metadata for the struct
// [ObjectListMessagesResponse]
type objectListMessagesResponseJSON struct {
	ID                 apijson.Field
	Typename           apijson.Field
	Actors             apijson.Field
	ArchivedAt         apijson.Field
	ChannelID          apijson.Field
	ClickedAt          apijson.Field
	Data               apijson.Field
	EngagementStatuses apijson.Field
	InsertedAt         apijson.Field
	InteractedAt       apijson.Field
	LinkClickedAt      apijson.Field
	Metadata           apijson.Field
	ReadAt             apijson.Field
	Recipient          apijson.Field
	ScheduledAt        apijson.Field
	SeenAt             apijson.Field
	Source             apijson.Field
	Status             apijson.Field
	Tenant             apijson.Field
	UpdatedAt          apijson.Field
	Workflow           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ObjectListMessagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectListMessagesResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [ObjectListMessagesResponseActorsObjectReference].
type ObjectListMessagesResponseActorsUnion interface {
	ImplementsObjectListMessagesResponseActorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectListMessagesResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObjectListMessagesResponseActorsObjectReference{}),
		},
	)
}

// An object reference to a recipient
type ObjectListMessagesResponseActorsObjectReference struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                              `json:"collection,required"`
	JSON       objectListMessagesResponseActorsObjectReferenceJSON `json:"-"`
}

// objectListMessagesResponseActorsObjectReferenceJSON contains the JSON metadata
// for the struct [ObjectListMessagesResponseActorsObjectReference]
type objectListMessagesResponseActorsObjectReferenceJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectListMessagesResponseActorsObjectReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectListMessagesResponseActorsObjectReferenceJSON) RawJSON() string {
	return r.raw
}

func (r ObjectListMessagesResponseActorsObjectReference) ImplementsObjectListMessagesResponseActorsUnion() {
}

type ObjectListMessagesResponseEngagementStatus string

const (
	ObjectListMessagesResponseEngagementStatusSeen        ObjectListMessagesResponseEngagementStatus = "seen"
	ObjectListMessagesResponseEngagementStatusRead        ObjectListMessagesResponseEngagementStatus = "read"
	ObjectListMessagesResponseEngagementStatusInteracted  ObjectListMessagesResponseEngagementStatus = "interacted"
	ObjectListMessagesResponseEngagementStatusLinkClicked ObjectListMessagesResponseEngagementStatus = "link_clicked"
	ObjectListMessagesResponseEngagementStatusArchived    ObjectListMessagesResponseEngagementStatus = "archived"
)

func (r ObjectListMessagesResponseEngagementStatus) IsKnown() bool {
	switch r {
	case ObjectListMessagesResponseEngagementStatusSeen, ObjectListMessagesResponseEngagementStatusRead, ObjectListMessagesResponseEngagementStatusInteracted, ObjectListMessagesResponseEngagementStatusLinkClicked, ObjectListMessagesResponseEngagementStatusArchived:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Union satisfied by [shared.UnionString] or
// [ObjectListMessagesResponseRecipientObjectReference].
type ObjectListMessagesResponseRecipientUnion interface {
	ImplementsObjectListMessagesResponseRecipientUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectListMessagesResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObjectListMessagesResponseRecipientObjectReference{}),
		},
	)
}

// An object reference to a recipient
type ObjectListMessagesResponseRecipientObjectReference struct {
	// An object identifier
	ID string `json:"id,required"`
	// The collection the object belongs to
	Collection string                                                 `json:"collection,required"`
	JSON       objectListMessagesResponseRecipientObjectReferenceJSON `json:"-"`
}

// objectListMessagesResponseRecipientObjectReferenceJSON contains the JSON
// metadata for the struct [ObjectListMessagesResponseRecipientObjectReference]
type objectListMessagesResponseRecipientObjectReferenceJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectListMessagesResponseRecipientObjectReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectListMessagesResponseRecipientObjectReferenceJSON) RawJSON() string {
	return r.raw
}

func (r ObjectListMessagesResponseRecipientObjectReference) ImplementsObjectListMessagesResponseRecipientUnion() {
}

// Source information
type ObjectListMessagesResponseSource struct {
	Typename string `json:"__typename,required"`
	// The workflow categories
	Categories []string `json:"categories,required"`
	// The workflow key
	Key string `json:"key,required"`
	// The source version ID
	VersionID string                               `json:"version_id,required" format:"uuid"`
	JSON      objectListMessagesResponseSourceJSON `json:"-"`
}

// objectListMessagesResponseSourceJSON contains the JSON metadata for the struct
// [ObjectListMessagesResponseSource]
type objectListMessagesResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectListMessagesResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectListMessagesResponseSourceJSON) RawJSON() string {
	return r.raw
}

// Message delivery status
type ObjectListMessagesResponseStatus string

const (
	ObjectListMessagesResponseStatusQueued            ObjectListMessagesResponseStatus = "queued"
	ObjectListMessagesResponseStatusSent              ObjectListMessagesResponseStatus = "sent"
	ObjectListMessagesResponseStatusDelivered         ObjectListMessagesResponseStatus = "delivered"
	ObjectListMessagesResponseStatusDeliveryAttempted ObjectListMessagesResponseStatus = "delivery_attempted"
	ObjectListMessagesResponseStatusUndelivered       ObjectListMessagesResponseStatus = "undelivered"
	ObjectListMessagesResponseStatusNotSent           ObjectListMessagesResponseStatus = "not_sent"
	ObjectListMessagesResponseStatusBounced           ObjectListMessagesResponseStatus = "bounced"
)

func (r ObjectListMessagesResponseStatus) IsKnown() bool {
	switch r {
	case ObjectListMessagesResponseStatusQueued, ObjectListMessagesResponseStatusSent, ObjectListMessagesResponseStatusDelivered, ObjectListMessagesResponseStatusDeliveryAttempted, ObjectListMessagesResponseStatusUndelivered, ObjectListMessagesResponseStatusNotSent, ObjectListMessagesResponseStatusBounced:
		return true
	}
	return false
}

// A preference set object.
type ObjectListPreferencesResponse struct {
	ID       string `json:"id,required"`
	Typename string `json:"__typename,required"`
	// A map of categories and their settings
	Categories map[string]ObjectListPreferencesResponseCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// A map of workflows and their settings
	Workflows map[string]ObjectListPreferencesResponseWorkflowsUnion `json:"workflows,nullable"`
	JSON      objectListPreferencesResponseJSON                      `json:"-"`
}

// objectListPreferencesResponseJSON contains the JSON metadata for the struct
// [ObjectListPreferencesResponse]
type objectListPreferencesResponseJSON struct {
	ID           apijson.Field
	Typename     apijson.Field
	Categories   apijson.Field
	ChannelTypes apijson.Field
	Workflows    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObjectListPreferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectListPreferencesResponseJSON) RawJSON() string {
	return r.raw
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [ObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject].
type ObjectListPreferencesResponseCategoriesUnion interface {
	ImplementsObjectListPreferencesResponseCategoriesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectListPreferencesResponseCategoriesUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(ObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                      `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                    `json:"conditions,nullable"`
	JSON         objectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// objectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [ObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject]
type objectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObjectListPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsObjectListPreferencesResponseCategoriesUnion() {
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [ObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject].
type ObjectListPreferencesResponseWorkflowsUnion interface {
	ImplementsObjectListPreferencesResponseWorkflowsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectListPreferencesResponseWorkflowsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(ObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                     `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                   `json:"conditions,nullable"`
	JSON         objectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// objectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [ObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject]
type objectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObjectListPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsObjectListPreferencesResponseWorkflowsUnion() {
}

// A schedule that represents a recurring workflow execution
type ObjectListSchedulesResponse struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient ObjectListSchedulesResponseRecipient `json:"recipient,required"`
	Repeats   []shared.ScheduleRepeatRule          `json:"repeats,required"`
	UpdatedAt time.Time                            `json:"updated_at,required" format:"date-time"`
	Workflow  string                               `json:"workflow,required"`
	Typename  string                               `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            ObjectListSchedulesResponseActor `json:"actor,nullable"`
	Data             map[string]interface{}           `json:"data,nullable"`
	LastOccurrenceAt time.Time                        `json:"last_occurrence_at,nullable" format:"date-time"`
	NextOccurrenceAt time.Time                        `json:"next_occurrence_at,nullable" format:"date-time"`
	Tenant           string                           `json:"tenant,nullable"`
	JSON             objectListSchedulesResponseJSON  `json:"-"`
}

// objectListSchedulesResponseJSON contains the JSON metadata for the struct
// [ObjectListSchedulesResponse]
type objectListSchedulesResponseJSON struct {
	ID               apijson.Field
	InsertedAt       apijson.Field
	Recipient        apijson.Field
	Repeats          apijson.Field
	UpdatedAt        apijson.Field
	Workflow         apijson.Field
	Typename         apijson.Field
	Actor            apijson.Field
	Data             apijson.Field
	LastOccurrenceAt apijson.Field
	NextOccurrenceAt apijson.Field
	Tenant           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ObjectListSchedulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectListSchedulesResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type ObjectListSchedulesResponseRecipient struct {
	ID          string                                   `json:"id,required"`
	Typename    string                                   `json:"__typename,required"`
	UpdatedAt   time.Time                                `json:"updated_at,required" format:"date-time"`
	Avatar      string                                   `json:"avatar,nullable"`
	Collection  string                                   `json:"collection"`
	CreatedAt   time.Time                                `json:"created_at,nullable" format:"date-time"`
	Email       string                                   `json:"email,nullable"`
	Name        string                                   `json:"name,nullable"`
	PhoneNumber string                                   `json:"phone_number,nullable"`
	Timezone    string                                   `json:"timezone,nullable"`
	JSON        objectListSchedulesResponseRecipientJSON `json:"-"`
	union       ObjectListSchedulesResponseRecipientUnion
}

// objectListSchedulesResponseRecipientJSON contains the JSON metadata for the
// struct [ObjectListSchedulesResponseRecipient]
type objectListSchedulesResponseRecipientJSON struct {
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

func (r objectListSchedulesResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *ObjectListSchedulesResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = ObjectListSchedulesResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObjectListSchedulesResponseRecipientUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r ObjectListSchedulesResponseRecipient) AsUnion() ObjectListSchedulesResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ObjectListSchedulesResponseRecipientUnion interface {
	ImplementsObjectListSchedulesResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectListSchedulesResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A recipient, which is either a user or an object
type ObjectListSchedulesResponseActor struct {
	ID          string                               `json:"id,required"`
	Typename    string                               `json:"__typename,required"`
	UpdatedAt   time.Time                            `json:"updated_at,required" format:"date-time"`
	Avatar      string                               `json:"avatar,nullable"`
	Collection  string                               `json:"collection"`
	CreatedAt   time.Time                            `json:"created_at,nullable" format:"date-time"`
	Email       string                               `json:"email,nullable"`
	Name        string                               `json:"name,nullable"`
	PhoneNumber string                               `json:"phone_number,nullable"`
	Timezone    string                               `json:"timezone,nullable"`
	JSON        objectListSchedulesResponseActorJSON `json:"-"`
	union       ObjectListSchedulesResponseActorUnion
}

// objectListSchedulesResponseActorJSON contains the JSON metadata for the struct
// [ObjectListSchedulesResponseActor]
type objectListSchedulesResponseActorJSON struct {
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

func (r objectListSchedulesResponseActorJSON) RawJSON() string {
	return r.raw
}

func (r *ObjectListSchedulesResponseActor) UnmarshalJSON(data []byte) (err error) {
	*r = ObjectListSchedulesResponseActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObjectListSchedulesResponseActorUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r ObjectListSchedulesResponseActor) AsUnion() ObjectListSchedulesResponseActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ObjectListSchedulesResponseActorUnion interface {
	ImplementsObjectListSchedulesResponseActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectListSchedulesResponseActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A subscription object
type ObjectListSubscriptionsResponse struct {
	Typename   string    `json:"__typename,required"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A custom-object entity which belongs to a collection.
	Object shared.Object `json:"object,required"`
	// A recipient, which is either a user or an object
	Recipient ObjectListSubscriptionsResponseRecipient `json:"recipient,required"`
	UpdatedAt time.Time                                `json:"updated_at,required" format:"date-time"`
	// The custom properties associated with the subscription
	Properties map[string]interface{}              `json:"properties,nullable"`
	JSON       objectListSubscriptionsResponseJSON `json:"-"`
}

// objectListSubscriptionsResponseJSON contains the JSON metadata for the struct
// [ObjectListSubscriptionsResponse]
type objectListSubscriptionsResponseJSON struct {
	Typename    apijson.Field
	InsertedAt  apijson.Field
	Object      apijson.Field
	Recipient   apijson.Field
	UpdatedAt   apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectListSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectListSubscriptionsResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type ObjectListSubscriptionsResponseRecipient struct {
	ID          string                                       `json:"id,required"`
	Typename    string                                       `json:"__typename,required"`
	UpdatedAt   time.Time                                    `json:"updated_at,required" format:"date-time"`
	Avatar      string                                       `json:"avatar,nullable"`
	Collection  string                                       `json:"collection"`
	CreatedAt   time.Time                                    `json:"created_at,nullable" format:"date-time"`
	Email       string                                       `json:"email,nullable"`
	Name        string                                       `json:"name,nullable"`
	PhoneNumber string                                       `json:"phone_number,nullable"`
	Timezone    string                                       `json:"timezone,nullable"`
	JSON        objectListSubscriptionsResponseRecipientJSON `json:"-"`
	union       ObjectListSubscriptionsResponseRecipientUnion
}

// objectListSubscriptionsResponseRecipientJSON contains the JSON metadata for the
// struct [ObjectListSubscriptionsResponseRecipient]
type objectListSubscriptionsResponseRecipientJSON struct {
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

func (r objectListSubscriptionsResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *ObjectListSubscriptionsResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = ObjectListSubscriptionsResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObjectListSubscriptionsResponseRecipientUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r ObjectListSubscriptionsResponseRecipient) AsUnion() ObjectListSubscriptionsResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ObjectListSubscriptionsResponseRecipientUnion interface {
	ImplementsObjectListSubscriptionsResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectListSubscriptionsResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// Channel data for various channel types
type ObjectSetChannelDataResponse struct {
	Typename  string `json:"__typename,required"`
	ChannelID string `json:"channel_id,required" format:"uuid"`
	// Channel data for push providers
	Data ObjectSetChannelDataResponseData `json:"data,required"`
	JSON objectSetChannelDataResponseJSON `json:"-"`
}

// objectSetChannelDataResponseJSON contains the JSON metadata for the struct
// [ObjectSetChannelDataResponse]
type objectSetChannelDataResponseJSON struct {
	Typename    apijson.Field
	ChannelID   apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectSetChannelDataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectSetChannelDataResponseJSON) RawJSON() string {
	return r.raw
}

// Channel data for push providers
type ObjectSetChannelDataResponseData struct {
	// This field can have the runtime type of [shared.SlackChannelDataToken].
	Token interface{} `json:"token"`
	// This field can have the runtime type of [[]shared.SlackChannelDataConnection],
	// [[]shared.MsTeamsChannelDataConnection],
	// [[]shared.DiscordChannelDataConnection].
	Connections interface{} `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID string `json:"ms_teams_tenant_id,nullable" format:"uuid"`
	// This field can have the runtime type of [[]string].
	PlayerIDs interface{} `json:"player_ids"`
	// This field can have the runtime type of [[]string].
	Tokens interface{}                          `json:"tokens"`
	JSON   objectSetChannelDataResponseDataJSON `json:"-"`
	union  ObjectSetChannelDataResponseDataUnion
}

// objectSetChannelDataResponseDataJSON contains the JSON metadata for the struct
// [ObjectSetChannelDataResponseData]
type objectSetChannelDataResponseDataJSON struct {
	Token           apijson.Field
	Connections     apijson.Field
	MsTeamsTenantID apijson.Field
	PlayerIDs       apijson.Field
	Tokens          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r objectSetChannelDataResponseDataJSON) RawJSON() string {
	return r.raw
}

func (r *ObjectSetChannelDataResponseData) UnmarshalJSON(data []byte) (err error) {
	*r = ObjectSetChannelDataResponseData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObjectSetChannelDataResponseDataUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.PushChannelData],
// [shared.SlackChannelData], [shared.MsTeamsChannelData],
// [shared.DiscordChannelData], [shared.OneSignalChannelData].
func (r ObjectSetChannelDataResponseData) AsUnion() ObjectSetChannelDataResponseDataUnion {
	return r.union
}

// Channel data for push providers
//
// Union satisfied by [shared.PushChannelData], [shared.SlackChannelData],
// [shared.MsTeamsChannelData], [shared.DiscordChannelData] or
// [shared.OneSignalChannelData].
type ObjectSetChannelDataResponseDataUnion interface {
	ImplementsObjectSetChannelDataResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectSetChannelDataResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.PushChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.SlackChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.MsTeamsChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.DiscordChannelData{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.OneSignalChannelData{}),
		},
	)
}

// A preference set object.
type ObjectSetPreferencesResponse struct {
	ID       string `json:"id,required"`
	Typename string `json:"__typename,required"`
	// A map of categories and their settings
	Categories map[string]ObjectSetPreferencesResponseCategoriesUnion `json:"categories,nullable"`
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes `json:"channel_types,nullable"`
	// A map of workflows and their settings
	Workflows map[string]ObjectSetPreferencesResponseWorkflowsUnion `json:"workflows,nullable"`
	JSON      objectSetPreferencesResponseJSON                      `json:"-"`
}

// objectSetPreferencesResponseJSON contains the JSON metadata for the struct
// [ObjectSetPreferencesResponse]
type objectSetPreferencesResponseJSON struct {
	ID           apijson.Field
	Typename     apijson.Field
	Categories   apijson.Field
	ChannelTypes apijson.Field
	Workflows    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObjectSetPreferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectSetPreferencesResponseJSON) RawJSON() string {
	return r.raw
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [ObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject].
type ObjectSetPreferencesResponseCategoriesUnion interface {
	ImplementsObjectSetPreferencesResponseCategoriesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectSetPreferencesResponseCategoriesUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(ObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                     `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                   `json:"conditions,nullable"`
	JSON         objectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// objectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [ObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject]
type objectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObjectSetPreferencesResponseCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsObjectSetPreferencesResponseCategoriesUnion() {
}

// Workflow or category preferences within a preference set
//
// Union satisfied by [shared.UnionBool] or
// [ObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject].
type ObjectSetPreferencesResponseWorkflowsUnion interface {
	ImplementsObjectSetPreferencesResponseWorkflowsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObjectSetPreferencesResponseWorkflowsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(ObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject{}),
		},
	)
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes shared.PreferenceSetChannelTypes                                                    `json:"channel_types,nullable"`
	Conditions   []shared.Condition                                                                  `json:"conditions,nullable"`
	JSON         objectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON `json:"-"`
}

// objectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON
// contains the JSON metadata for the struct
// [ObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject]
type objectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON struct {
	ChannelTypes apijson.Field
	Conditions   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObjectSetPreferencesResponseWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsObjectSetPreferencesResponseWorkflowsUnion() {
}

type ObjectListParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [ObjectListParams]'s query parameters as `url.Values`.
func (r ObjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectAddSubscriptionsParams struct {
	// The recipients to subscribe to the object
	Recipients param.Field[[]ObjectAddSubscriptionsParamsRecipientUnion] `json:"recipients,required"`
	// The custom properties associated with the subscription
	Properties param.Field[map[string]interface{}] `json:"properties"`
}

func (r ObjectAddSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type ObjectAddSubscriptionsParamsRecipient struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID          param.Field[string]      `json:"id,required"`
	ChannelData param.Field[interface{}] `json:"channel_data"`
	Collection  param.Field[string]      `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt   param.Field[time.Time]   `json:"created_at" format:"date-time"`
	Preferences param.Field[interface{}] `json:"preferences"`
}

func (r ObjectAddSubscriptionsParamsRecipient) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectAddSubscriptionsParamsRecipient) ImplementsObjectAddSubscriptionsParamsRecipientUnion() {
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// Satisfied by [shared.UnionString], [shared.InlineIdentifyUserRequestParam],
// [ObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequest],
// [ObjectAddSubscriptionsParamsRecipient].
type ObjectAddSubscriptionsParamsRecipientUnion interface {
	ImplementsObjectAddSubscriptionsParamsRecipientUnion()
}

// Inline identifies a custom object belonging to a collection
type ObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequest struct {
	ID         param.Field[string] `json:"id,required"`
	Collection param.Field[string] `json:"collection,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]shared.ChannelDataRequestParam] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                                 `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]shared.PreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                                   `json:"-,extras"`
}

func (r ObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectAddSubscriptionsParamsRecipientsInlineIdentifyObjectRequest) ImplementsObjectAddSubscriptionsParamsRecipientUnion() {
}

type ObjectDeleteSubscriptionsParams struct {
	Recipients param.Field[[]ObjectDeleteSubscriptionsParamsRecipientUnion] `json:"recipients,required"`
}

func (r ObjectDeleteSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type ObjectDeleteSubscriptionsParamsRecipient struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID          param.Field[string]      `json:"id,required"`
	ChannelData param.Field[interface{}] `json:"channel_data"`
	Collection  param.Field[string]      `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt   param.Field[time.Time]   `json:"created_at" format:"date-time"`
	Preferences param.Field[interface{}] `json:"preferences"`
}

func (r ObjectDeleteSubscriptionsParamsRecipient) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectDeleteSubscriptionsParamsRecipient) ImplementsObjectDeleteSubscriptionsParamsRecipientUnion() {
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// Satisfied by [shared.UnionString], [shared.InlineIdentifyUserRequestParam],
// [ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequest],
// [ObjectDeleteSubscriptionsParamsRecipient].
type ObjectDeleteSubscriptionsParamsRecipientUnion interface {
	ImplementsObjectDeleteSubscriptionsParamsRecipientUnion()
}

// Inline identifies a custom object belonging to a collection
type ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequest struct {
	ID         param.Field[string] `json:"id,required"`
	Collection param.Field[string] `json:"collection,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]shared.ChannelDataRequestParam] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                                 `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]shared.PreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                                   `json:"-,extras"`
}

func (r ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectDeleteSubscriptionsParamsRecipientsInlineIdentifyObjectRequest) ImplementsObjectDeleteSubscriptionsParamsRecipientUnion() {
}

type ObjectGetPreferencesParams struct {
	// Tenant ID
	Tenant param.Field[string] `query:"tenant"`
}

// URLQuery serializes [ObjectGetPreferencesParams]'s query parameters as
// `url.Values`.
func (r ObjectGetPreferencesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectListMessagesParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The channel ID
	ChannelID param.Field[string] `query:"channel_id"`
	// The engagement status of the message
	EngagementStatus param.Field[[]ObjectListMessagesParamsEngagementStatus] `query:"engagement_status"`
	// The message IDs to filter messages by
	MessageIDs param.Field[[]string] `query:"message_ids"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// The source of the message (workflow key)
	Source param.Field[string] `query:"source"`
	// The status of the message
	Status param.Field[[]ObjectListMessagesParamsStatus] `query:"status"`
	// The tenant ID
	Tenant param.Field[string] `query:"tenant"`
	// The trigger data to filter messages by. Must be a valid JSON object.
	TriggerData param.Field[string] `query:"trigger_data"`
	// The workflow categories to filter messages by
	WorkflowCategories param.Field[[]string] `query:"workflow_categories"`
	// The workflow recipient run ID to filter messages by
	WorkflowRecipientRunID param.Field[string] `query:"workflow_recipient_run_id" format:"uuid"`
	// The workflow run ID to filter messages by
	WorkflowRunID param.Field[string] `query:"workflow_run_id" format:"uuid"`
}

// URLQuery serializes [ObjectListMessagesParams]'s query parameters as
// `url.Values`.
func (r ObjectListMessagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectListMessagesParamsEngagementStatus string

const (
	ObjectListMessagesParamsEngagementStatusSeen        ObjectListMessagesParamsEngagementStatus = "seen"
	ObjectListMessagesParamsEngagementStatusRead        ObjectListMessagesParamsEngagementStatus = "read"
	ObjectListMessagesParamsEngagementStatusInteracted  ObjectListMessagesParamsEngagementStatus = "interacted"
	ObjectListMessagesParamsEngagementStatusLinkClicked ObjectListMessagesParamsEngagementStatus = "link_clicked"
	ObjectListMessagesParamsEngagementStatusArchived    ObjectListMessagesParamsEngagementStatus = "archived"
)

func (r ObjectListMessagesParamsEngagementStatus) IsKnown() bool {
	switch r {
	case ObjectListMessagesParamsEngagementStatusSeen, ObjectListMessagesParamsEngagementStatusRead, ObjectListMessagesParamsEngagementStatusInteracted, ObjectListMessagesParamsEngagementStatusLinkClicked, ObjectListMessagesParamsEngagementStatusArchived:
		return true
	}
	return false
}

type ObjectListMessagesParamsStatus string

const (
	ObjectListMessagesParamsStatusQueued            ObjectListMessagesParamsStatus = "queued"
	ObjectListMessagesParamsStatusSent              ObjectListMessagesParamsStatus = "sent"
	ObjectListMessagesParamsStatusDelivered         ObjectListMessagesParamsStatus = "delivered"
	ObjectListMessagesParamsStatusDeliveryAttempted ObjectListMessagesParamsStatus = "delivery_attempted"
	ObjectListMessagesParamsStatusUndelivered       ObjectListMessagesParamsStatus = "undelivered"
	ObjectListMessagesParamsStatusNotSent           ObjectListMessagesParamsStatus = "not_sent"
	ObjectListMessagesParamsStatusBounced           ObjectListMessagesParamsStatus = "bounced"
)

func (r ObjectListMessagesParamsStatus) IsKnown() bool {
	switch r {
	case ObjectListMessagesParamsStatusQueued, ObjectListMessagesParamsStatusSent, ObjectListMessagesParamsStatusDelivered, ObjectListMessagesParamsStatusDeliveryAttempted, ObjectListMessagesParamsStatusUndelivered, ObjectListMessagesParamsStatusNotSent, ObjectListMessagesParamsStatusBounced:
		return true
	}
	return false
}

type ObjectListSchedulesParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// The ID of the tenant to list schedules for
	Tenant param.Field[string] `query:"tenant"`
	// The ID of the workflow to list schedules for
	Workflow param.Field[string] `query:"workflow"`
}

// URLQuery serializes [ObjectListSchedulesParams]'s query parameters as
// `url.Values`.
func (r ObjectListSchedulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectListSubscriptionsParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// Mode of the request
	Mode param.Field[ObjectListSubscriptionsParamsMode] `query:"mode"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// Recipients to filter by (only used if mode is `object`)
	Recipients param.Field[[]ObjectListSubscriptionsParamsRecipientUnion] `query:"recipients"`
}

// URLQuery serializes [ObjectListSubscriptionsParams]'s query parameters as
// `url.Values`.
func (r ObjectListSubscriptionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Mode of the request
type ObjectListSubscriptionsParamsMode string

const (
	ObjectListSubscriptionsParamsModeRecipient ObjectListSubscriptionsParamsMode = "recipient"
	ObjectListSubscriptionsParamsModeObject    ObjectListSubscriptionsParamsMode = "object"
)

func (r ObjectListSubscriptionsParamsMode) IsKnown() bool {
	switch r {
	case ObjectListSubscriptionsParamsModeRecipient, ObjectListSubscriptionsParamsModeObject:
		return true
	}
	return false
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Satisfied by [shared.UnionString],
// [ObjectListSubscriptionsParamsRecipientsObjectReference].
type ObjectListSubscriptionsParamsRecipientUnion interface {
	ImplementsObjectListSubscriptionsParamsRecipientUnion()
}

// An object reference to a recipient
type ObjectListSubscriptionsParamsRecipientsObjectReference struct {
	// An object identifier
	ID param.Field[string] `query:"id,required"`
	// The collection the object belongs to
	Collection param.Field[string] `query:"collection,required"`
}

// URLQuery serializes [ObjectListSubscriptionsParamsRecipientsObjectReference]'s
// query parameters as `url.Values`.
func (r ObjectListSubscriptionsParamsRecipientsObjectReference) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (r ObjectListSubscriptionsParamsRecipientsObjectReference) ImplementsObjectListSubscriptionsParamsRecipientUnion() {
}

type ObjectSetParams struct {
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]shared.ChannelDataRequestParam] `json:"channel_data"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]shared.PreferenceSetRequestParam] `json:"preferences"`
}

func (r ObjectSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectSetChannelDataParams struct {
	// Set channel data for a type of channel
	ChannelDataRequest shared.ChannelDataRequestParam `json:"channel_data_request,required"`
}

func (r ObjectSetChannelDataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ChannelDataRequest)
}

type ObjectSetPreferencesParams struct {
	// Set preferences for a recipient
	PreferenceSetRequest shared.PreferenceSetRequestParam `json:"preference_set_request,required"`
}

func (r ObjectSetPreferencesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PreferenceSetRequest)
}

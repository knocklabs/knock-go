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
func (r *ObjectService) AddSubscriptions(ctx context.Context, collection string, objectID string, body ObjectAddSubscriptionsParams, opts ...option.RequestOption) (res *[]shared.Subscription, err error) {
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
func (r *ObjectService) DeleteSubscriptions(ctx context.Context, collection string, objectID string, body ObjectDeleteSubscriptionsParams, opts ...option.RequestOption) (res *[]shared.Subscription, err error) {
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
func (r *ObjectService) GetChannelData(ctx context.Context, collection string, objectID string, channelID string, opts ...option.RequestOption) (res *shared.ChannelData, err error) {
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
func (r *ObjectService) GetPreferences(ctx context.Context, collection string, objectID string, id string, query ObjectGetPreferencesParams, opts ...option.RequestOption) (res *shared.PreferenceSet, err error) {
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
func (r *ObjectService) ListPreferences(ctx context.Context, collection string, objectID string, opts ...option.RequestOption) (res *[]shared.PreferenceSet, err error) {
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
func (r *ObjectService) ListSchedules(ctx context.Context, collection string, id string, query ObjectListSchedulesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[shared.Schedule], err error) {
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
func (r *ObjectService) ListSchedulesAutoPaging(ctx context.Context, collection string, id string, query ObjectListSchedulesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[shared.Schedule] {
	return pagination.NewEntriesCursorAutoPager(r.ListSchedules(ctx, collection, id, query, opts...))
}

// List subscriptions for an object. Either list all subscriptions that belong to
// the object, or all subscriptions that this object has. Determined by the `mode`
// query parameter.
func (r *ObjectService) ListSubscriptions(ctx context.Context, collection string, objectID string, query ObjectListSubscriptionsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[shared.Subscription], err error) {
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
func (r *ObjectService) ListSubscriptionsAutoPaging(ctx context.Context, collection string, objectID string, query ObjectListSubscriptionsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[shared.Subscription] {
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
func (r *ObjectService) SetChannelData(ctx context.Context, collection string, objectID string, channelID string, body ObjectSetChannelDataParams, opts ...option.RequestOption) (res *shared.ChannelData, err error) {
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
func (r *ObjectService) SetPreferences(ctx context.Context, collection string, objectID string, id string, body ObjectSetPreferencesParams, opts ...option.RequestOption) (res *shared.PreferenceSet, err error) {
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
	Recipients param.Field[[]shared.RecipientRequestUnionParam] `json:"recipients,required"`
	// The custom properties associated with the subscription
	Properties param.Field[map[string]interface{}] `json:"properties"`
}

func (r ObjectAddSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectDeleteSubscriptionsParams struct {
	Recipients param.Field[[]shared.RecipientRequestUnionParam] `json:"recipients,required"`
}

func (r ObjectDeleteSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	ChannelData param.Field[shared.InlineChannelDataRequestParam] `json:"channel_data"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[shared.InlinePreferenceSetRequestParam] `json:"preferences"`
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

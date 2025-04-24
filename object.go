// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/packages/pagination"
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

// Returns a paginated list of objects from the specified collection. Optionally
// includes preference data for the objects.
func (r *ObjectService) List(ctx context.Context, collection string, query ObjectListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Object], err error) {
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

// Returns a paginated list of objects from the specified collection. Optionally
// includes preference data for the objects.
func (r *ObjectService) ListAutoPaging(ctx context.Context, collection string, query ObjectListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Object] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, collection, query, opts...))
}

// Permanently removes an object from the specified collection. This operation
// cannot be undone.
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
// updated. This endpoint also handles
// [inline identifications](/managing-recipients/identifying-recipients#inline-identifying-recipients)
// for the `recipient`.
func (r *ObjectService) AddSubscriptions(ctx context.Context, collection string, objectID string, body ObjectAddSubscriptionsParams, opts ...option.RequestOption) (res *[]Subscription, err error) {
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

// Delete subscriptions for the specified recipients from an object. Returns the
// list of deleted subscriptions.
func (r *ObjectService) DeleteSubscriptions(ctx context.Context, collection string, objectID string, body ObjectDeleteSubscriptionsParams, opts ...option.RequestOption) (res *[]Subscription, err error) {
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

// Retrieves a specific object by its ID from the specified collection. Returns the
// object with all its properties.
func (r *ObjectService) Get(ctx context.Context, collection string, id string, opts ...option.RequestOption) (res *Object, err error) {
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

// Returns the channel data for the specified object and channel.
func (r *ObjectService) GetChannelData(ctx context.Context, collection string, objectID string, channelID string, opts ...option.RequestOption) (res *ChannelData, err error) {
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

// Returns the preference set for the specified object and preference set `id`.
func (r *ObjectService) GetPreferences(ctx context.Context, collection string, objectID string, id string, opts ...option.RequestOption) (res *PreferenceSet, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a paginated list of messages for a specific object in the given
// collection. Allows filtering by message status and provides various sorting
// options.
func (r *ObjectService) ListMessages(ctx context.Context, collection string, id string, query ObjectListMessagesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Message], err error) {
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

// Returns a paginated list of messages for a specific object in the given
// collection. Allows filtering by message status and provides various sorting
// options.
func (r *ObjectService) ListMessagesAutoPaging(ctx context.Context, collection string, id string, query ObjectListMessagesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Message] {
	return pagination.NewEntriesCursorAutoPager(r.ListMessages(ctx, collection, id, query, opts...))
}

// Returns a paginated list of preference sets for the specified object.
func (r *ObjectService) ListPreferences(ctx context.Context, collection string, objectID string, opts ...option.RequestOption) (res *[]PreferenceSet, err error) {
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

// Returns a paginated list of schedules for an object.
func (r *ObjectService) ListSchedules(ctx context.Context, collection string, id string, query ObjectListSchedulesParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Schedule], err error) {
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

// Returns a paginated list of schedules for an object.
func (r *ObjectService) ListSchedulesAutoPaging(ctx context.Context, collection string, id string, query ObjectListSchedulesParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Schedule] {
	return pagination.NewEntriesCursorAutoPager(r.ListSchedules(ctx, collection, id, query, opts...))
}

// List subscriptions for an object. Either list the recipients that subscribe to
// the provided object, or list the objects that the provided object is subscribed
// to. Determined by the `mode` query parameter.
func (r *ObjectService) ListSubscriptions(ctx context.Context, collection string, objectID string, query ObjectListSubscriptionsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Subscription], err error) {
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

// List subscriptions for an object. Either list the recipients that subscribe to
// the provided object, or list the objects that the provided object is subscribed
// to. Determined by the `mode` query parameter.
func (r *ObjectService) ListSubscriptionsAutoPaging(ctx context.Context, collection string, objectID string, query ObjectListSubscriptionsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Subscription] {
	return pagination.NewEntriesCursorAutoPager(r.ListSubscriptions(ctx, collection, objectID, query, opts...))
}

// Creates a new object or updates an existing one in the specified collection.
// This operation is used to identify objects with their properties, as well as
// optional preferences and channel data.
func (r *ObjectService) Set(ctx context.Context, collection string, id string, body ObjectSetParams, opts ...option.RequestOption) (res *Object, err error) {
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

// Sets the channel data for the specified object and channel. If no object exists
// in the current environment for the given `collection` and `object_id`, Knock
// will create the object as part of this request.
func (r *ObjectService) SetChannelData(ctx context.Context, collection string, objectID string, channelID string, body ObjectSetChannelDataParams, opts ...option.RequestOption) (res *ChannelData, err error) {
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

// Sets preferences within the given preference set. This is a destructive
// operation and will replace any existing preferences with the preferences given.
// If no object exists in the current environment for the given `:collection` and
// `:object_id`, Knock will create the object as part of this request. The
// preference set `:id` can be either `default` or a `tenant.id`. Learn more about
// [per-tenant preferences](/preferences/tenant-preferences).
func (r *ObjectService) SetPreferences(ctx context.Context, collection string, objectID string, id string, body ObjectSetPreferencesParams, opts ...option.RequestOption) (res *PreferenceSet, err error) {
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

// Unsets the channel data for the specified object and channel.
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

// A custom [Object](/concepts/objects) entity which belongs to a collection.
type InlineObjectRequestParam struct {
	// Unique identifier for the object.
	ID param.Field[string] `json:"id,required"`
	// The collection this object belongs to.
	Collection param.Field[string] `json:"collection,required"`
	// A request to set channel data for a type of channel inline.
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	// Timestamp when the resource was created.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// An optional set of [preferences](/concepts/preferences) for the object.
	Preferences param.Field[[]InlinePreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                         `json:"-,extras"`
}

func (r InlineObjectRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InlineObjectRequestParam) ImplementsRecipientRequestUnionParam() {}

// A custom [Object](/concepts/objects) entity which belongs to a collection.
type Object struct {
	// Unique identifier for the object.
	ID string `json:"id,required"`
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The collection this object belongs to.
	Collection string `json:"collection,required"`
	// The timestamp when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Timestamp when the resource was created.
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

func (r Object) implementsRecipient() {}

type ObjectListParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// Includes preferences of the objects in the response.
	Include param.Field[[]ObjectListParamsInclude] `query:"include"`
	// The number of items per page.
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [ObjectListParams]'s query parameters as `url.Values`.
func (r ObjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectListParamsInclude string

const (
	ObjectListParamsIncludePreferences ObjectListParamsInclude = "preferences"
)

func (r ObjectListParamsInclude) IsKnown() bool {
	switch r {
	case ObjectListParamsIncludePreferences:
		return true
	}
	return false
}

type ObjectAddSubscriptionsParams struct {
	// The recipients of the subscription. You can subscribe up to 100 recipients to an
	// object at a time.
	Recipients param.Field[[]RecipientRequestUnionParam] `json:"recipients,required"`
	// The custom properties associated with the subscription relationship.
	Properties param.Field[map[string]interface{}] `json:"properties"`
}

func (r ObjectAddSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectDeleteSubscriptionsParams struct {
	// The recipients of the subscription. You can subscribe up to 100 recipients to an
	// object at a time.
	Recipients param.Field[[]RecipientReferenceUnionParam] `json:"recipients,required"`
}

func (r ObjectDeleteSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectListMessagesParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// Limits the results to items with the corresponding channel ID.
	ChannelID param.Field[string] `query:"channel_id"`
	// Limits the results to messages with the given engagement status.
	EngagementStatus param.Field[[]ObjectListMessagesParamsEngagementStatus] `query:"engagement_status"`
	InsertedAt       param.Field[ObjectListMessagesParamsInsertedAt]         `query:"inserted_at"`
	// Limits the results to only the message IDs given (max 50). Note: when using this
	// option, the results will be subject to any other filters applied to the query.
	MessageIDs param.Field[[]string] `query:"message_ids"`
	// The number of items per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Limits the results to messages triggered by the given workflow key.
	Source param.Field[string] `query:"source"`
	// Limits the results to messages with the given delivery status.
	Status param.Field[[]ObjectListMessagesParamsStatus] `query:"status"`
	// Limits the results to items with the corresponding tenant.
	Tenant param.Field[string] `query:"tenant"`
	// Limits the results to only messages that were generated with the given data. See
	// [trigger data filtering](/api-reference/overview/trigger-data-filtering) for
	// more information.
	TriggerData param.Field[string] `query:"trigger_data"`
	// Limits the results to messages related to any of the provided categories.
	WorkflowCategories param.Field[[]string] `query:"workflow_categories"`
	// Limits the results to messages for a specific recipient's workflow run.
	WorkflowRecipientRunID param.Field[string] `query:"workflow_recipient_run_id" format:"uuid"`
	// Limits the results to messages associated with the top-level workflow run ID
	// returned by the workflow trigger request.
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

type ObjectListMessagesParamsInsertedAt struct {
	// Limits the results to messages inserted after the given date.
	Gt param.Field[string] `query:"gt"`
	// Limits the results to messages inserted after or on the given date.
	Gte param.Field[string] `query:"gte"`
	// Limits the results to messages inserted before the given date.
	Lt param.Field[string] `query:"lt"`
	// Limits the results to messages inserted before or on the given date.
	Lte param.Field[string] `query:"lte"`
}

// URLQuery serializes [ObjectListMessagesParamsInsertedAt]'s query parameters as
// `url.Values`.
func (r ObjectListMessagesParamsInsertedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// The number of items per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Filter schedules by tenant id.
	Tenant param.Field[string] `query:"tenant"`
	// Filter schedules by workflow id.
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
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// Additional fields to include in the response.
	Include param.Field[[]ObjectListSubscriptionsParamsInclude] `query:"include"`
	// Mode of the request. `recipient` to list the objects that the provided object is
	// subscribed to, `object` to list the recipients that subscribe to the provided
	// object.
	Mode param.Field[ObjectListSubscriptionsParamsMode] `query:"mode"`
	// Objects to filter by (only used if mode is `recipient`).
	Objects param.Field[[]ObjectListSubscriptionsParamsObject] `query:"objects"`
	// The number of items per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Recipients to filter by (only used if mode is `object`).
	Recipients param.Field[[]RecipientReferenceUnionParam] `query:"recipients"`
}

// URLQuery serializes [ObjectListSubscriptionsParams]'s query parameters as
// `url.Values`.
func (r ObjectListSubscriptionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectListSubscriptionsParamsInclude string

const (
	ObjectListSubscriptionsParamsIncludePreferences ObjectListSubscriptionsParamsInclude = "preferences"
)

func (r ObjectListSubscriptionsParamsInclude) IsKnown() bool {
	switch r {
	case ObjectListSubscriptionsParamsIncludePreferences:
		return true
	}
	return false
}

// Mode of the request. `recipient` to list the objects that the provided object is
// subscribed to, `object` to list the recipients that subscribe to the provided
// object.
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

// A reference to a recipient object.
type ObjectListSubscriptionsParamsObject struct {
	// An identifier for the recipient object.
	ID param.Field[string] `query:"id"`
	// The collection the recipient object belongs to.
	Collection param.Field[string] `query:"collection"`
}

// URLQuery serializes [ObjectListSubscriptionsParamsObject]'s query parameters as
// `url.Values`.
func (r ObjectListSubscriptionsParamsObject) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectSetParams struct {
	// A request to set channel data for a type of channel inline.
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	// The locale of the object. Used for
	// [message localization](/concepts/translations).
	Locale param.Field[string] `json:"locale"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
	// The timezone of the object. Must be a valid
	// [tz database time zone string](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	// Used for
	// [recurring schedules](/concepts/schedules#scheduling-workflows-with-recurring-schedules-for-recipients).
	Timezone param.Field[string] `json:"timezone"`
}

func (r ObjectSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectSetChannelDataParams struct {
	// A request to set channel data for a type of channel.
	ChannelDataRequest ChannelDataRequestParam `json:"channel_data_request,required"`
}

func (r ObjectSetChannelDataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ChannelDataRequest)
}

type ObjectSetPreferencesParams struct {
	// A request to set a preference set for a recipient.
	PreferenceSetRequest PreferenceSetRequestParam `json:"preference_set_request,required"`
}

func (r ObjectSetPreferencesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PreferenceSetRequest)
}

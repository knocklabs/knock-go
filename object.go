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

// List objects in a collection
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

// List objects in a collection
func (r *ObjectService) ListAutoPaging(ctx context.Context, collection string, query ObjectListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Object] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, collection, query, opts...))
}

// Upsert subscriptions for an object
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

// Delete subscriptions for an object
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

// Get channel data for an object
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

// List subscriptions for an object. Either list all subscriptions that belong to
// the object, or all subscriptions that this object has. Determined by the `mode`
// query parameter.
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

// List subscriptions for an object. Either list all subscriptions that belong to
// the object, or all subscriptions that this object has. Determined by the `mode`
// query parameter.
func (r *ObjectService) ListSubscriptionsAutoPaging(ctx context.Context, collection string, objectID string, query ObjectListSubscriptionsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Subscription] {
	return pagination.NewEntriesCursorAutoPager(r.ListSubscriptions(ctx, collection, objectID, query, opts...))
}

// Set channel data for an object
func (r *ObjectService) SetChannelData(ctx context.Context, collection string, objectID string, channelID string, opts ...option.RequestOption) (res *ChannelData, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Unset channel data for an object
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

// Inline identifies a custom object belonging to a collection
type InlineObjectRequestParam struct {
	ID         param.Field[string] `json:"id,required"`
	Collection param.Field[string] `json:"collection,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                     `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
}

func (r InlineObjectRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InlineObjectRequestParam) ImplementsRecipientRequestUnionParam() {}

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

func (r Object) implementsRecipient() {}

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
	Recipients param.Field[[]RecipientRequestUnionParam] `json:"recipients,required"`
	// The custom properties associated with the subscription
	Properties param.Field[map[string]interface{}] `json:"properties"`
}

func (r ObjectAddSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectDeleteSubscriptionsParams struct {
	Recipients param.Field[[]RecipientRequestUnionParam] `json:"recipients,required"`
}

func (r ObjectDeleteSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectListSubscriptionsParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// Mode of the request
	Mode param.Field[ObjectListSubscriptionsParamsMode] `query:"mode"`
	// Objects to filter by (only used if mode is `recipient`)
	Objects param.Field[[]ObjectListSubscriptionsParamsObjectUnion] `query:"objects"`
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
// Satisfied by [shared.UnionString], [ObjectListSubscriptionsParamsObjectsObject].
type ObjectListSubscriptionsParamsObjectUnion interface {
	ImplementsObjectListSubscriptionsParamsObjectUnion()
}

// An object reference to a recipient
type ObjectListSubscriptionsParamsObjectsObject struct {
	// An object identifier
	ID param.Field[string] `query:"id,required"`
	// The collection the object belongs to
	Collection param.Field[string] `query:"collection,required"`
}

// URLQuery serializes [ObjectListSubscriptionsParamsObjectsObject]'s query
// parameters as `url.Values`.
func (r ObjectListSubscriptionsParamsObjectsObject) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (r ObjectListSubscriptionsParamsObjectsObject) ImplementsObjectListSubscriptionsParamsObjectUnion() {
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Satisfied by [shared.UnionString],
// [ObjectListSubscriptionsParamsRecipientsObject].
type ObjectListSubscriptionsParamsRecipientUnion interface {
	ImplementsObjectListSubscriptionsParamsRecipientUnion()
}

// An object reference to a recipient
type ObjectListSubscriptionsParamsRecipientsObject struct {
	// An object identifier
	ID param.Field[string] `query:"id,required"`
	// The collection the object belongs to
	Collection param.Field[string] `query:"collection,required"`
}

// URLQuery serializes [ObjectListSubscriptionsParamsRecipientsObject]'s query
// parameters as `url.Values`.
func (r ObjectListSubscriptionsParamsRecipientsObject) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (r ObjectListSubscriptionsParamsRecipientsObject) ImplementsObjectListSubscriptionsParamsRecipientUnion() {
}

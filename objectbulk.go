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
	"github.com/stainless-sdks/knock-go/shared"
)

// ObjectBulkService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectBulkService] method instead.
type ObjectBulkService struct {
	Options []option.RequestOption
}

// NewObjectBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectBulkService(opts ...option.RequestOption) (r *ObjectBulkService) {
	r = &ObjectBulkService{}
	r.Options = opts
	return
}

// Bulk delete objects
func (r *ObjectBulkService) Delete(ctx context.Context, collection string, body ObjectBulkDeleteParams, opts ...option.RequestOption) (res *ObjectBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/delete", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add subscriptions for a set of objects in a single collection. If a subscription
// already exists, it will be updated.
func (r *ObjectBulkService) AddSubscriptions(ctx context.Context, collection string, body ObjectBulkAddSubscriptionsParams, opts ...option.RequestOption) (res *ObjectBulkAddSubscriptionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/subscriptions/add", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk set objects
func (r *ObjectBulkService) Set(ctx context.Context, collection string, body ObjectBulkSetParams, opts ...option.RequestOption) (res *ObjectBulkSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/set", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A bulk operation entity
type ObjectBulkDeleteResponse struct {
	ID                 string                         `json:"id,required" format:"uuid"`
	Typename           string                         `json:"__typename,required"`
	EstimatedTotalRows int64                          `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                      `json:"inserted_at,required" format:"date-time"`
	Name               string                         `json:"name,required"`
	ProcessedRows      int64                          `json:"processed_rows,required"`
	Status             ObjectBulkDeleteResponseStatus `json:"status,required"`
	SuccessCount       int64                          `json:"success_count,required"`
	UpdatedAt          time.Time                      `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                      `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                          `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []ObjectBulkDeleteResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                           `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                           `json:"started_at,nullable" format:"date-time"`
	JSON       objectBulkDeleteResponseJSON        `json:"-"`
}

// objectBulkDeleteResponseJSON contains the JSON metadata for the struct
// [ObjectBulkDeleteResponse]
type objectBulkDeleteResponseJSON struct {
	ID                 apijson.Field
	Typename           apijson.Field
	EstimatedTotalRows apijson.Field
	InsertedAt         apijson.Field
	Name               apijson.Field
	ProcessedRows      apijson.Field
	Status             apijson.Field
	SuccessCount       apijson.Field
	UpdatedAt          apijson.Field
	CompletedAt        apijson.Field
	ErrorCount         apijson.Field
	ErrorItems         apijson.Field
	FailedAt           apijson.Field
	StartedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ObjectBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ObjectBulkDeleteResponseStatus string

const (
	ObjectBulkDeleteResponseStatusQueued     ObjectBulkDeleteResponseStatus = "queued"
	ObjectBulkDeleteResponseStatusProcessing ObjectBulkDeleteResponseStatus = "processing"
	ObjectBulkDeleteResponseStatusCompleted  ObjectBulkDeleteResponseStatus = "completed"
	ObjectBulkDeleteResponseStatusFailed     ObjectBulkDeleteResponseStatus = "failed"
)

func (r ObjectBulkDeleteResponseStatus) IsKnown() bool {
	switch r {
	case ObjectBulkDeleteResponseStatusQueued, ObjectBulkDeleteResponseStatusProcessing, ObjectBulkDeleteResponseStatusCompleted, ObjectBulkDeleteResponseStatusFailed:
		return true
	}
	return false
}

type ObjectBulkDeleteResponseErrorItem struct {
	ID         string                                `json:"id,required"`
	Collection string                                `json:"collection,nullable"`
	JSON       objectBulkDeleteResponseErrorItemJSON `json:"-"`
}

// objectBulkDeleteResponseErrorItemJSON contains the JSON metadata for the struct
// [ObjectBulkDeleteResponseErrorItem]
type objectBulkDeleteResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectBulkDeleteResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkDeleteResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type ObjectBulkAddSubscriptionsResponse struct {
	ID                 string                                   `json:"id,required" format:"uuid"`
	Typename           string                                   `json:"__typename,required"`
	EstimatedTotalRows int64                                    `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                                `json:"inserted_at,required" format:"date-time"`
	Name               string                                   `json:"name,required"`
	ProcessedRows      int64                                    `json:"processed_rows,required"`
	Status             ObjectBulkAddSubscriptionsResponseStatus `json:"status,required"`
	SuccessCount       int64                                    `json:"success_count,required"`
	UpdatedAt          time.Time                                `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                                `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                                    `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []ObjectBulkAddSubscriptionsResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                                     `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                                     `json:"started_at,nullable" format:"date-time"`
	JSON       objectBulkAddSubscriptionsResponseJSON        `json:"-"`
}

// objectBulkAddSubscriptionsResponseJSON contains the JSON metadata for the struct
// [ObjectBulkAddSubscriptionsResponse]
type objectBulkAddSubscriptionsResponseJSON struct {
	ID                 apijson.Field
	Typename           apijson.Field
	EstimatedTotalRows apijson.Field
	InsertedAt         apijson.Field
	Name               apijson.Field
	ProcessedRows      apijson.Field
	Status             apijson.Field
	SuccessCount       apijson.Field
	UpdatedAt          apijson.Field
	CompletedAt        apijson.Field
	ErrorCount         apijson.Field
	ErrorItems         apijson.Field
	FailedAt           apijson.Field
	StartedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ObjectBulkAddSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkAddSubscriptionsResponseJSON) RawJSON() string {
	return r.raw
}

type ObjectBulkAddSubscriptionsResponseStatus string

const (
	ObjectBulkAddSubscriptionsResponseStatusQueued     ObjectBulkAddSubscriptionsResponseStatus = "queued"
	ObjectBulkAddSubscriptionsResponseStatusProcessing ObjectBulkAddSubscriptionsResponseStatus = "processing"
	ObjectBulkAddSubscriptionsResponseStatusCompleted  ObjectBulkAddSubscriptionsResponseStatus = "completed"
	ObjectBulkAddSubscriptionsResponseStatusFailed     ObjectBulkAddSubscriptionsResponseStatus = "failed"
)

func (r ObjectBulkAddSubscriptionsResponseStatus) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsResponseStatusQueued, ObjectBulkAddSubscriptionsResponseStatusProcessing, ObjectBulkAddSubscriptionsResponseStatusCompleted, ObjectBulkAddSubscriptionsResponseStatusFailed:
		return true
	}
	return false
}

type ObjectBulkAddSubscriptionsResponseErrorItem struct {
	ID         string                                          `json:"id,required"`
	Collection string                                          `json:"collection,nullable"`
	JSON       objectBulkAddSubscriptionsResponseErrorItemJSON `json:"-"`
}

// objectBulkAddSubscriptionsResponseErrorItemJSON contains the JSON metadata for
// the struct [ObjectBulkAddSubscriptionsResponseErrorItem]
type objectBulkAddSubscriptionsResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectBulkAddSubscriptionsResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkAddSubscriptionsResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type ObjectBulkSetResponse struct {
	ID                 string                      `json:"id,required" format:"uuid"`
	Typename           string                      `json:"__typename,required"`
	EstimatedTotalRows int64                       `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                   `json:"inserted_at,required" format:"date-time"`
	Name               string                      `json:"name,required"`
	ProcessedRows      int64                       `json:"processed_rows,required"`
	Status             ObjectBulkSetResponseStatus `json:"status,required"`
	SuccessCount       int64                       `json:"success_count,required"`
	UpdatedAt          time.Time                   `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                   `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                       `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []ObjectBulkSetResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                        `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                        `json:"started_at,nullable" format:"date-time"`
	JSON       objectBulkSetResponseJSON        `json:"-"`
}

// objectBulkSetResponseJSON contains the JSON metadata for the struct
// [ObjectBulkSetResponse]
type objectBulkSetResponseJSON struct {
	ID                 apijson.Field
	Typename           apijson.Field
	EstimatedTotalRows apijson.Field
	InsertedAt         apijson.Field
	Name               apijson.Field
	ProcessedRows      apijson.Field
	Status             apijson.Field
	SuccessCount       apijson.Field
	UpdatedAt          apijson.Field
	CompletedAt        apijson.Field
	ErrorCount         apijson.Field
	ErrorItems         apijson.Field
	FailedAt           apijson.Field
	StartedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ObjectBulkSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkSetResponseJSON) RawJSON() string {
	return r.raw
}

type ObjectBulkSetResponseStatus string

const (
	ObjectBulkSetResponseStatusQueued     ObjectBulkSetResponseStatus = "queued"
	ObjectBulkSetResponseStatusProcessing ObjectBulkSetResponseStatus = "processing"
	ObjectBulkSetResponseStatusCompleted  ObjectBulkSetResponseStatus = "completed"
	ObjectBulkSetResponseStatusFailed     ObjectBulkSetResponseStatus = "failed"
)

func (r ObjectBulkSetResponseStatus) IsKnown() bool {
	switch r {
	case ObjectBulkSetResponseStatusQueued, ObjectBulkSetResponseStatusProcessing, ObjectBulkSetResponseStatusCompleted, ObjectBulkSetResponseStatusFailed:
		return true
	}
	return false
}

type ObjectBulkSetResponseErrorItem struct {
	ID         string                             `json:"id,required"`
	Collection string                             `json:"collection,nullable"`
	JSON       objectBulkSetResponseErrorItemJSON `json:"-"`
}

// objectBulkSetResponseErrorItemJSON contains the JSON metadata for the struct
// [ObjectBulkSetResponseErrorItem]
type objectBulkSetResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectBulkSetResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkSetResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

type ObjectBulkDeleteParams struct {
	// The IDs of the objects to delete
	ObjectIDs param.Field[[]string] `query:"object_ids,required"`
}

// URLQuery serializes [ObjectBulkDeleteParams]'s query parameters as `url.Values`.
func (r ObjectBulkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectBulkAddSubscriptionsParams struct {
	Subscriptions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscription] `json:"subscriptions,required"`
}

func (r ObjectBulkAddSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscription struct {
	ID         param.Field[string]                                                        `json:"id,required"`
	Recipients param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion] `json:"recipients,required"`
	Properties param.Field[map[string]interface{}]                                        `json:"properties"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[shared.InlineChannelDataRequestParam] `json:"channel_data"`
	Collection  param.Field[string]                               `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[shared.InlinePreferenceSetRequestParam] `json:"preferences"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion() {
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// Satisfied by [shared.UnionString], [shared.InlineIdentifyUserRequestParam],
// [shared.InlineIdentifyObjectRequestParam],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion()
}

type ObjectBulkSetParams struct {
	Objects param.Field[[]shared.InlineIdentifyObjectRequestParam] `json:"objects,required"`
}

func (r ObjectBulkSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

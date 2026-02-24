// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
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

// Bulk deletes objects from the specified collection.
func (r *ObjectBulkService) Delete(ctx context.Context, collection string, body ObjectBulkDeleteParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/delete", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add subscriptions for all objects in a single collection. If a subscription for
// an object in the collection already exists, it will be updated. This endpoint
// also handles
// [inline identifications](/managing-recipients/identifying-recipients#inline-identifying-recipients)
// for the `recipient` field.
func (r *ObjectBulkService) AddSubscriptions(ctx context.Context, collection string, body ObjectBulkAddSubscriptionsParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/subscriptions/add", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete subscriptions for many objects in a single collection type. If a
// subscription for an object in the collection doesn't exist, it will be skipped.
func (r *ObjectBulkService) DeleteSubscriptions(ctx context.Context, collection string, body ObjectBulkDeleteSubscriptionsParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/subscriptions/delete", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk sets up to 1,000 objects at a time in the specified collection.
func (r *ObjectBulkService) Set(ctx context.Context, collection string, body ObjectBulkSetParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/set", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObjectBulkDeleteParams struct {
	// List of object IDs to delete.
	ObjectIDs param.Field[[]string] `json:"object_ids" api:"required"`
}

func (r ObjectBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParams struct {
	// A nested list of subscriptions.
	Subscriptions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscription] `json:"subscriptions" api:"required"`
}

func (r ObjectBulkAddSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of subscriptions. 1 subscribed-to id, and N subscriber recipients.
type ObjectBulkAddSubscriptionsParamsSubscription struct {
	// Unique identifier for the object.
	ID param.Field[string] `json:"id" api:"required"`
	// The recipients of the subscription. You can subscribe up to 100 recipients to an
	// object at a time.
	Recipients param.Field[[]RecipientRequestUnionParam] `json:"recipients" api:"required"`
	// The custom properties associated with the subscription relationship.
	Properties param.Field[map[string]interface{}] `json:"properties"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkDeleteSubscriptionsParams struct {
	// A nested list of subscriptions.
	Subscriptions param.Field[[]ObjectBulkDeleteSubscriptionsParamsSubscription] `json:"subscriptions" api:"required"`
}

func (r ObjectBulkDeleteSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of subscriptions. 1 subscribed-to id, and N subscriber recipients.
type ObjectBulkDeleteSubscriptionsParamsSubscription struct {
	// Unique identifier for the object.
	ID param.Field[string] `json:"id" api:"required"`
	// The recipients of the subscription. You can subscribe up to 100 recipients to an
	// object at a time.
	Recipients param.Field[[]RecipientReferenceUnionParam] `json:"recipients" api:"required"`
}

func (r ObjectBulkDeleteSubscriptionsParamsSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParams struct {
	// A list of objects.
	Objects param.Field[[]ObjectBulkSetParamsObject] `json:"objects" api:"required"`
}

func (r ObjectBulkSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom [Object](/concepts/objects) entity which belongs to a collection.
type ObjectBulkSetParamsObject struct {
	// Unique identifier for the object.
	ID param.Field[string] `json:"id" api:"required"`
	// A request to set channel data for a type of channel inline.
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	// Timestamp when the resource was created.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// An optional name for the object.
	Name param.Field[string] `json:"name"`
	// Inline set preferences for a recipient, where the key is the preference set id.
	// Preferences that are set inline will be merged into any existing preferences
	// rather than replacing them.
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
}

func (r ObjectBulkSetParamsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/subscriptions/add", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk sets up to 1,000 objects at a time in the specified collection.
func (r *ObjectBulkService) Set(ctx context.Context, collection string, body ObjectBulkSetParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/set", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObjectBulkDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ObjectBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ObjectBulkAddSubscriptionsParams struct {
	// A list of subscriptions.
	Subscriptions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscription] `json:"subscriptions,required"`
}

func (r ObjectBulkAddSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscription struct {
	// The recipients of the subscription. You can subscribe up to 100 recipients to an
	// object at a time.
	Recipients param.Field[[]RecipientRequestUnionParam] `json:"recipients,required"`
	// The custom properties associated with the subscription relationship.
	Properties param.Field[map[string]interface{}] `json:"properties"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParams struct {
	// A list of objects.
	Objects param.Field[[]InlineObjectRequestParam] `json:"objects,required"`
}

func (r ObjectBulkSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

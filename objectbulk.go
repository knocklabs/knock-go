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
// an object in the collection already exists, it will be updated.
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

// Bulk sets objects in the specified collection.
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
	// A list of object IDs.
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
	// A list of subscriptions.
	Subscriptions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscription] `json:"subscriptions,required"`
}

func (r ObjectBulkAddSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscription struct {
	// Unique identifier for the subscription.
	ID param.Field[string] `json:"id,required"`
	// The recipients of the subscription.
	Recipients param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion] `json:"recipients,required"`
	// The custom properties associated with the recipients of the subscription.
	Properties param.Field[map[string]interface{}] `json:"properties"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient struct {
	// The ID for the user that you set when identifying them in Knock.
	ID param.Field[string] `json:"id,required"`
	// A request to set channel data for a type of channel inline.
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	// The collection this object belongs to.
	Collection param.Field[string] `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
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
// Satisfied by [shared.UnionString], [InlineIdentifyUserRequestParam],
// [InlineObjectRequestParam],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion()
}

type ObjectBulkSetParams struct {
	// A list of objects.
	Objects param.Field[[]InlineObjectRequestParam] `json:"objects,required"`
}

func (r ObjectBulkSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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

// Deletes objects in bulk for a given collection
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

// Bulk upserts subscriptions for a set of objects in a single collection
func (r *ObjectBulkService) AddSubscriptions(ctx context.Context, collection string, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/subscriptions/add", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Sets objects in bulk for a given collection
func (r *ObjectBulkService) Set(ctx context.Context, collection string, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/set", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
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

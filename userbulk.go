// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
)

// UserBulkService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserBulkService] method instead.
type UserBulkService struct {
	Options []option.RequestOption
}

// NewUserBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserBulkService(opts ...option.RequestOption) (r *UserBulkService) {
	r = &UserBulkService{}
	r.Options = opts
	return
}

// Bulk deletes a list of users
func (r *UserBulkService) Delete(ctx context.Context, body UserBulkDeleteParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/bulk/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk identifies a list of users
func (r *UserBulkService) Identify(ctx context.Context, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/bulk/identify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Bulk sets user preferences
func (r *UserBulkService) SetPreferences(ctx context.Context, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/bulk/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type UserBulkDeleteParams struct {
	// The IDs of the users to delete
	UserIDs param.Field[[]string] `query:"user_ids,required"`
}

// URLQuery serializes [UserBulkDeleteParams]'s query parameters as `url.Values`.
func (r UserBulkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

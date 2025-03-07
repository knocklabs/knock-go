// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
)

// TenantBulkService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantBulkService] method instead.
type TenantBulkService struct {
	Options []option.RequestOption
}

// NewTenantBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTenantBulkService(opts ...option.RequestOption) (r *TenantBulkService) {
	r = &TenantBulkService{}
	r.Options = opts
	return
}

// Bulk delete tenants
func (r *TenantBulkService) Delete(ctx context.Context, body TenantBulkDeleteParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tenants/bulk/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk set tenants
func (r *TenantBulkService) Set(ctx context.Context, body TenantBulkSetParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tenants/bulk/set"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TenantBulkDeleteParams struct {
	// The IDs of the tenants to delete
	TenantIDs param.Field[[]string] `query:"tenant_ids,required"`
}

// URLQuery serializes [TenantBulkDeleteParams]'s query parameters as `url.Values`.
func (r TenantBulkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantBulkSetParams struct {
	Tenants param.Field[[]shared.InlineTenantRequestUnionParam] `json:"tenants,required"`
}

func (r TenantBulkSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
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
func (r *TenantBulkService) Delete(ctx context.Context, body TenantBulkDeleteParams, opts ...option.RequestOption) (res *TenantBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tenants/bulk/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk set tenants
func (r *TenantBulkService) Set(ctx context.Context, body TenantBulkSetParams, opts ...option.RequestOption) (res *TenantBulkSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tenants/bulk/set"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A bulk operation entity
type TenantBulkDeleteResponse struct {
	ID                 string                         `json:"id,required" format:"uuid"`
	Typename           string                         `json:"__typename,required"`
	EstimatedTotalRows int64                          `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                      `json:"inserted_at,required" format:"date-time"`
	Name               string                         `json:"name,required"`
	ProcessedRows      int64                          `json:"processed_rows,required"`
	Status             TenantBulkDeleteResponseStatus `json:"status,required"`
	SuccessCount       int64                          `json:"success_count,required"`
	UpdatedAt          time.Time                      `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                      `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                          `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []TenantBulkDeleteResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                           `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                           `json:"started_at,nullable" format:"date-time"`
	JSON       tenantBulkDeleteResponseJSON        `json:"-"`
}

// tenantBulkDeleteResponseJSON contains the JSON metadata for the struct
// [TenantBulkDeleteResponse]
type tenantBulkDeleteResponseJSON struct {
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

func (r *TenantBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TenantBulkDeleteResponseStatus string

const (
	TenantBulkDeleteResponseStatusQueued     TenantBulkDeleteResponseStatus = "queued"
	TenantBulkDeleteResponseStatusProcessing TenantBulkDeleteResponseStatus = "processing"
	TenantBulkDeleteResponseStatusCompleted  TenantBulkDeleteResponseStatus = "completed"
	TenantBulkDeleteResponseStatusFailed     TenantBulkDeleteResponseStatus = "failed"
)

func (r TenantBulkDeleteResponseStatus) IsKnown() bool {
	switch r {
	case TenantBulkDeleteResponseStatusQueued, TenantBulkDeleteResponseStatusProcessing, TenantBulkDeleteResponseStatusCompleted, TenantBulkDeleteResponseStatusFailed:
		return true
	}
	return false
}

type TenantBulkDeleteResponseErrorItem struct {
	ID         string                                `json:"id,required"`
	Collection string                                `json:"collection,nullable"`
	JSON       tenantBulkDeleteResponseErrorItemJSON `json:"-"`
}

// tenantBulkDeleteResponseErrorItemJSON contains the JSON metadata for the struct
// [TenantBulkDeleteResponseErrorItem]
type tenantBulkDeleteResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantBulkDeleteResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantBulkDeleteResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type TenantBulkSetResponse struct {
	ID                 string                      `json:"id,required" format:"uuid"`
	Typename           string                      `json:"__typename,required"`
	EstimatedTotalRows int64                       `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                   `json:"inserted_at,required" format:"date-time"`
	Name               string                      `json:"name,required"`
	ProcessedRows      int64                       `json:"processed_rows,required"`
	Status             TenantBulkSetResponseStatus `json:"status,required"`
	SuccessCount       int64                       `json:"success_count,required"`
	UpdatedAt          time.Time                   `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                   `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                       `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []TenantBulkSetResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                        `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                        `json:"started_at,nullable" format:"date-time"`
	JSON       tenantBulkSetResponseJSON        `json:"-"`
}

// tenantBulkSetResponseJSON contains the JSON metadata for the struct
// [TenantBulkSetResponse]
type tenantBulkSetResponseJSON struct {
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

func (r *TenantBulkSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantBulkSetResponseJSON) RawJSON() string {
	return r.raw
}

type TenantBulkSetResponseStatus string

const (
	TenantBulkSetResponseStatusQueued     TenantBulkSetResponseStatus = "queued"
	TenantBulkSetResponseStatusProcessing TenantBulkSetResponseStatus = "processing"
	TenantBulkSetResponseStatusCompleted  TenantBulkSetResponseStatus = "completed"
	TenantBulkSetResponseStatusFailed     TenantBulkSetResponseStatus = "failed"
)

func (r TenantBulkSetResponseStatus) IsKnown() bool {
	switch r {
	case TenantBulkSetResponseStatusQueued, TenantBulkSetResponseStatusProcessing, TenantBulkSetResponseStatusCompleted, TenantBulkSetResponseStatusFailed:
		return true
	}
	return false
}

type TenantBulkSetResponseErrorItem struct {
	ID         string                             `json:"id,required"`
	Collection string                             `json:"collection,nullable"`
	JSON       tenantBulkSetResponseErrorItemJSON `json:"-"`
}

// tenantBulkSetResponseErrorItemJSON contains the JSON metadata for the struct
// [TenantBulkSetResponseErrorItem]
type tenantBulkSetResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantBulkSetResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantBulkSetResponseErrorItemJSON) RawJSON() string {
	return r.raw
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
	Tenants param.Field[[]TenantBulkSetParamsTenantUnion] `json:"tenants,required"`
}

func (r TenantBulkSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An inline tenant request
//
// Satisfied by [shared.UnionString], [shared.TenantRequestParam].
type TenantBulkSetParamsTenantUnion interface {
	ImplementsTenantBulkSetParamsTenantUnion()
}

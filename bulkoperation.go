// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
)

// BulkOperationService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBulkOperationService] method instead.
type BulkOperationService struct {
	Options []option.RequestOption
}

// NewBulkOperationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBulkOperationService(opts ...option.RequestOption) (r *BulkOperationService) {
	r = &BulkOperationService{}
	r.Options = opts
	return
}

// Retrieves a bulk operation (if it exists) and displays the current state of it.
func (r *BulkOperationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BulkOperationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/bulk_operations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A bulk operation entity
type BulkOperationGetResponse struct {
	ID                 string                         `json:"id,required" format:"uuid"`
	Typename           string                         `json:"__typename,required"`
	EstimatedTotalRows int64                          `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                      `json:"inserted_at,required" format:"date-time"`
	Name               string                         `json:"name,required"`
	ProcessedRows      int64                          `json:"processed_rows,required"`
	Status             BulkOperationGetResponseStatus `json:"status,required"`
	SuccessCount       int64                          `json:"success_count,required"`
	UpdatedAt          time.Time                      `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                      `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                          `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []BulkOperationGetResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                           `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                           `json:"started_at,nullable" format:"date-time"`
	JSON       bulkOperationGetResponseJSON        `json:"-"`
}

// bulkOperationGetResponseJSON contains the JSON metadata for the struct
// [BulkOperationGetResponse]
type bulkOperationGetResponseJSON struct {
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

func (r *BulkOperationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkOperationGetResponseJSON) RawJSON() string {
	return r.raw
}

type BulkOperationGetResponseStatus string

const (
	BulkOperationGetResponseStatusQueued     BulkOperationGetResponseStatus = "queued"
	BulkOperationGetResponseStatusProcessing BulkOperationGetResponseStatus = "processing"
	BulkOperationGetResponseStatusCompleted  BulkOperationGetResponseStatus = "completed"
	BulkOperationGetResponseStatusFailed     BulkOperationGetResponseStatus = "failed"
)

func (r BulkOperationGetResponseStatus) IsKnown() bool {
	switch r {
	case BulkOperationGetResponseStatusQueued, BulkOperationGetResponseStatusProcessing, BulkOperationGetResponseStatusCompleted, BulkOperationGetResponseStatusFailed:
		return true
	}
	return false
}

type BulkOperationGetResponseErrorItem struct {
	ID         string                                `json:"id,required"`
	Collection string                                `json:"collection,nullable"`
	JSON       bulkOperationGetResponseErrorItemJSON `json:"-"`
}

// bulkOperationGetResponseErrorItemJSON contains the JSON metadata for the struct
// [BulkOperationGetResponseErrorItem]
type bulkOperationGetResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkOperationGetResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkOperationGetResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

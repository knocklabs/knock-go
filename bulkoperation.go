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
func (r *BulkOperationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BulkOperation, err error) {
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
type BulkOperation struct {
	ID                 string              `json:"id,required" format:"uuid"`
	Typename           string              `json:"__typename,required"`
	EstimatedTotalRows int64               `json:"estimated_total_rows,required"`
	InsertedAt         time.Time           `json:"inserted_at,required" format:"date-time"`
	Name               string              `json:"name,required"`
	ProcessedRows      int64               `json:"processed_rows,required"`
	Status             BulkOperationStatus `json:"status,required"`
	SuccessCount       int64               `json:"success_count,required"`
	UpdatedAt          time.Time           `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time           `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64               `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []BulkOperationErrorItem `json:"error_items"`
	FailedAt   time.Time                `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                `json:"started_at,nullable" format:"date-time"`
	JSON       bulkOperationJSON        `json:"-"`
}

// bulkOperationJSON contains the JSON metadata for the struct [BulkOperation]
type bulkOperationJSON struct {
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

func (r *BulkOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkOperationJSON) RawJSON() string {
	return r.raw
}

type BulkOperationStatus string

const (
	BulkOperationStatusQueued     BulkOperationStatus = "queued"
	BulkOperationStatusProcessing BulkOperationStatus = "processing"
	BulkOperationStatusCompleted  BulkOperationStatus = "completed"
	BulkOperationStatusFailed     BulkOperationStatus = "failed"
)

func (r BulkOperationStatus) IsKnown() bool {
	switch r {
	case BulkOperationStatusQueued, BulkOperationStatusProcessing, BulkOperationStatusCompleted, BulkOperationStatusFailed:
		return true
	}
	return false
}

type BulkOperationErrorItem struct {
	ID         string                     `json:"id,required"`
	Collection string                     `json:"collection,nullable"`
	JSON       bulkOperationErrorItemJSON `json:"-"`
}

// bulkOperationErrorItemJSON contains the JSON metadata for the struct
// [BulkOperationErrorItem]
type bulkOperationErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkOperationErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkOperationErrorItemJSON) RawJSON() string {
	return r.raw
}

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
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
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
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/bulk_operations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A bulk operation entity.
type BulkOperation struct {
	// Unique identifier for the bulk operation.
	ID string `json:"id" api:"required" format:"uuid"`
	// The typename of the schema.
	Typename string `json:"__typename" api:"required"`
	// The estimated total number of rows to process.
	EstimatedTotalRows int64 `json:"estimated_total_rows" api:"required"`
	// Timestamp when the resource was created.
	InsertedAt time.Time `json:"inserted_at" api:"required" format:"date-time"`
	// The name of the bulk operation.
	Name string `json:"name" api:"required"`
	// The number of rows processed so far.
	ProcessedRows int64 `json:"processed_rows" api:"required"`
	// The status of the bulk operation.
	Status BulkOperationStatus `json:"status" api:"required"`
	// The number of successful operations.
	SuccessCount int64 `json:"success_count" api:"required"`
	// The timestamp when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Timestamp when the bulk operation was completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// The number of failed operations.
	ErrorCount int64 `json:"error_count"`
	// A list of items that failed to be processed.
	ErrorItems []BulkOperationErrorItem `json:"error_items"`
	// Timestamp when the bulk operation failed.
	FailedAt time.Time `json:"failed_at" api:"nullable" format:"date-time"`
	// The URI to the bulk operation's progress.
	ProgressPath string `json:"progress_path" format:"uri"`
	// Timestamp when the bulk operation was started.
	StartedAt time.Time         `json:"started_at" api:"nullable" format:"date-time"`
	JSON      bulkOperationJSON `json:"-"`
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
	ProgressPath       apijson.Field
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

// The status of the bulk operation.
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
	// Unique identifier for the object.
	ID string `json:"id" api:"required"`
	// The collection this object belongs to.
	Collection string                     `json:"collection" api:"nullable"`
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

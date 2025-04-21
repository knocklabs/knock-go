// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
)

// ScheduleBulkService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScheduleBulkService] method instead.
type ScheduleBulkService struct {
	Options []option.RequestOption
}

// NewScheduleBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScheduleBulkService(opts ...option.RequestOption) (r *ScheduleBulkService) {
	r = &ScheduleBulkService{}
	r.Options = opts
	return
}

// Bulk creates up to 1,000 schedules at a time. This endpoint also handles
// [inline identifications](/managing-recipients/identifying-recipients#inline-identifying-recipients)
// for the `actor`, `recipient`, and `tenant` fields.
func (r *ScheduleBulkService) New(ctx context.Context, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules/bulk/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

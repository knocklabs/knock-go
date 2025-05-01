// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"time"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
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
func (r *ScheduleBulkService) New(ctx context.Context, body ScheduleBulkNewParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules/bulk/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ScheduleBulkNewParams struct {
	// A list of schedules.
	Schedules param.Field[[]ScheduleBulkNewParamsSchedule] `json:"schedules,required"`
}

func (r ScheduleBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A schedule represents a recurring workflow execution.
type ScheduleBulkNewParamsSchedule struct {
	// The key of the workflow.
	Workflow param.Field[string] `json:"workflow,required"`
	// Specifies a recipient in a request. This can either be a user identifier
	// (string), an inline user request (object), or an inline object request, which is
	// determined by the presence of a `collection` property.
	Actor param.Field[RecipientRequestUnionParam] `json:"actor"`
	// An optional map of data to pass into the workflow execution.
	Data param.Field[map[string]interface{}] `json:"data"`
	// The ending date and time for the schedule.
	EndingAt param.Field[time.Time] `json:"ending_at" format:"date-time"`
	// Specifies a recipient in a request. This can either be a user identifier
	// (string), an inline user request (object), or an inline object request, which is
	// determined by the presence of a `collection` property.
	Recipient param.Field[RecipientRequestUnionParam] `json:"recipient"`
	// The repeat rule for the schedule.
	Repeats param.Field[[]ScheduleRepeatRuleParam] `json:"repeats"`
	// The starting date and time for the schedule.
	ScheduledAt param.Field[time.Time] `json:"scheduled_at" format:"date-time"`
	// An request to set a tenant inline.
	Tenant param.Field[InlineTenantRequestUnionParam] `json:"tenant"`
}

func (r ScheduleBulkNewParamsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

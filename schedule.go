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
	"github.com/stainless-sdks/knock-go/packages/pagination"
)

// ScheduleService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScheduleService] method instead.
type ScheduleService struct {
	Options []option.RequestOption
}

// NewScheduleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScheduleService(opts ...option.RequestOption) (r *ScheduleService) {
	r = &ScheduleService{}
	r.Options = opts
	return
}

// Create schedules
func (r *ScheduleService) New(ctx context.Context, body ScheduleNewParams, opts ...option.RequestOption) (res *[]Schedule, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update schedules
func (r *ScheduleService) Update(ctx context.Context, body ScheduleUpdateParams, opts ...option.RequestOption) (res *[]Schedule, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List schedules
func (r *ScheduleService) List(ctx context.Context, query ScheduleListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Schedule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/schedules"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List schedules
func (r *ScheduleService) ListAutoPaging(ctx context.Context, query ScheduleListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Schedule] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete schedules
func (r *ScheduleService) Delete(ctx context.Context, body ScheduleDeleteParams, opts ...option.RequestOption) (res *[]Schedule, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// A schedule that represents a recurring workflow execution
type Schedule struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient Recipient            `json:"recipient,required"`
	Repeats   []ScheduleRepeatRule `json:"repeats,required"`
	UpdatedAt time.Time            `json:"updated_at,required" format:"date-time"`
	Workflow  string               `json:"workflow,required"`
	Typename  string               `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            Recipient              `json:"actor,nullable"`
	Data             map[string]interface{} `json:"data,nullable"`
	LastOccurrenceAt time.Time              `json:"last_occurrence_at,nullable" format:"date-time"`
	NextOccurrenceAt time.Time              `json:"next_occurrence_at,nullable" format:"date-time"`
	Tenant           string                 `json:"tenant,nullable"`
	JSON             scheduleJSON           `json:"-"`
}

// scheduleJSON contains the JSON metadata for the struct [Schedule]
type scheduleJSON struct {
	ID               apijson.Field
	InsertedAt       apijson.Field
	Recipient        apijson.Field
	Repeats          apijson.Field
	UpdatedAt        apijson.Field
	Workflow         apijson.Field
	Typename         apijson.Field
	Actor            apijson.Field
	Data             apijson.Field
	LastOccurrenceAt apijson.Field
	NextOccurrenceAt apijson.Field
	Tenant           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Schedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleJSON) RawJSON() string {
	return r.raw
}

// A schedule repeat rule
type ScheduleRepeatRule struct {
	Typename   string                      `json:"__typename,required"`
	Frequency  ScheduleRepeatRuleFrequency `json:"frequency,required"`
	DayOfMonth int64                       `json:"day_of_month,nullable"`
	Days       []ScheduleRepeatRuleDay     `json:"days,nullable"`
	Hours      int64                       `json:"hours,nullable"`
	Interval   int64                       `json:"interval"`
	Minutes    int64                       `json:"minutes,nullable"`
	JSON       scheduleRepeatRuleJSON      `json:"-"`
}

// scheduleRepeatRuleJSON contains the JSON metadata for the struct
// [ScheduleRepeatRule]
type scheduleRepeatRuleJSON struct {
	Typename    apijson.Field
	Frequency   apijson.Field
	DayOfMonth  apijson.Field
	Days        apijson.Field
	Hours       apijson.Field
	Interval    apijson.Field
	Minutes     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleRepeatRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleRepeatRuleJSON) RawJSON() string {
	return r.raw
}

type ScheduleRepeatRuleFrequency string

const (
	ScheduleRepeatRuleFrequencyDaily   ScheduleRepeatRuleFrequency = "daily"
	ScheduleRepeatRuleFrequencyWeekly  ScheduleRepeatRuleFrequency = "weekly"
	ScheduleRepeatRuleFrequencyMonthly ScheduleRepeatRuleFrequency = "monthly"
	ScheduleRepeatRuleFrequencyHourly  ScheduleRepeatRuleFrequency = "hourly"
)

func (r ScheduleRepeatRuleFrequency) IsKnown() bool {
	switch r {
	case ScheduleRepeatRuleFrequencyDaily, ScheduleRepeatRuleFrequencyWeekly, ScheduleRepeatRuleFrequencyMonthly, ScheduleRepeatRuleFrequencyHourly:
		return true
	}
	return false
}

type ScheduleRepeatRuleDay string

const (
	ScheduleRepeatRuleDayMon ScheduleRepeatRuleDay = "mon"
	ScheduleRepeatRuleDayTue ScheduleRepeatRuleDay = "tue"
	ScheduleRepeatRuleDayWed ScheduleRepeatRuleDay = "wed"
	ScheduleRepeatRuleDayThu ScheduleRepeatRuleDay = "thu"
	ScheduleRepeatRuleDayFri ScheduleRepeatRuleDay = "fri"
	ScheduleRepeatRuleDaySat ScheduleRepeatRuleDay = "sat"
	ScheduleRepeatRuleDaySun ScheduleRepeatRuleDay = "sun"
)

func (r ScheduleRepeatRuleDay) IsKnown() bool {
	switch r {
	case ScheduleRepeatRuleDayMon, ScheduleRepeatRuleDayTue, ScheduleRepeatRuleDayWed, ScheduleRepeatRuleDayThu, ScheduleRepeatRuleDayFri, ScheduleRepeatRuleDaySat, ScheduleRepeatRuleDaySun:
		return true
	}
	return false
}

// A schedule repeat rule
type ScheduleRepeatRuleParam struct {
	Typename   param.Field[string]                      `json:"__typename,required"`
	Frequency  param.Field[ScheduleRepeatRuleFrequency] `json:"frequency,required"`
	DayOfMonth param.Field[int64]                       `json:"day_of_month"`
	Days       param.Field[[]ScheduleRepeatRuleDay]     `json:"days"`
	Hours      param.Field[int64]                       `json:"hours"`
	Interval   param.Field[int64]                       `json:"interval"`
	Minutes    param.Field[int64]                       `json:"minutes"`
}

func (r ScheduleRepeatRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScheduleNewParams struct {
	Recipients  param.Field[[]ScheduleNewParamsRecipientUnion] `json:"recipients,required"`
	Repeats     param.Field[[]ScheduleRepeatRuleParam]         `json:"repeats,required"`
	Workflow    param.Field[string]                            `json:"workflow,required"`
	Data        param.Field[map[string]interface{}]            `json:"data"`
	EndingAt    param.Field[time.Time]                         `json:"ending_at" format:"date-time"`
	ScheduledAt param.Field[time.Time]                         `json:"scheduled_at" format:"date-time"`
	// An inline tenant request
	Tenant param.Field[InlineTenantRequestUnionParam] `json:"tenant"`
}

func (r ScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Satisfied by [shared.UnionString], [ScheduleNewParamsRecipientsObjectReference].
type ScheduleNewParamsRecipientUnion interface {
	ImplementsScheduleNewParamsRecipientUnion()
}

// An object reference to a recipient
type ScheduleNewParamsRecipientsObjectReference struct {
	// An object identifier
	ID param.Field[string] `json:"id,required"`
	// The collection the object belongs to
	Collection param.Field[string] `json:"collection,required"`
}

func (r ScheduleNewParamsRecipientsObjectReference) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScheduleNewParamsRecipientsObjectReference) ImplementsScheduleNewParamsRecipientUnion() {}

type ScheduleUpdateParams struct {
	ScheduleIDs param.Field[[]string] `json:"schedule_ids,required" format:"uuid"`
	// Specifies a recipient in a request. This can either be a user identifier
	// (string), an inline user request (object), or an inline object request, which is
	// determined by the presence of a `collection` property.
	Actor       param.Field[RecipientRequestUnionParam] `json:"actor"`
	Data        param.Field[map[string]interface{}]     `json:"data"`
	EndingAt    param.Field[time.Time]                  `json:"ending_at" format:"date-time"`
	Repeats     param.Field[[]ScheduleRepeatRuleParam]  `json:"repeats"`
	ScheduledAt param.Field[time.Time]                  `json:"scheduled_at" format:"date-time"`
	// An inline tenant request
	Tenant param.Field[InlineTenantRequestUnionParam] `json:"tenant"`
}

func (r ScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScheduleListParams struct {
	// Filter by workflow
	Workflow param.Field[string] `query:"workflow,required"`
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by recipient
	Recipients param.Field[[]ScheduleListParamsRecipientUnion] `query:"recipients"`
	// Filter by tenant
	Tenant param.Field[string] `query:"tenant"`
}

// URLQuery serializes [ScheduleListParams]'s query parameters as `url.Values`.
func (r ScheduleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (id, collection).
//
// Satisfied by [shared.UnionString],
// [ScheduleListParamsRecipientsObjectReference].
type ScheduleListParamsRecipientUnion interface {
	ImplementsScheduleListParamsRecipientUnion()
}

// An object reference to a recipient
type ScheduleListParamsRecipientsObjectReference struct {
	// An object identifier
	ID param.Field[string] `query:"id,required"`
	// The collection the object belongs to
	Collection param.Field[string] `query:"collection,required"`
}

// URLQuery serializes [ScheduleListParamsRecipientsObjectReference]'s query
// parameters as `url.Values`.
func (r ScheduleListParamsRecipientsObjectReference) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (r ScheduleListParamsRecipientsObjectReference) ImplementsScheduleListParamsRecipientUnion() {}

type ScheduleDeleteParams struct {
	ScheduleIDs param.Field[[]string] `json:"schedule_ids,required"`
}

func (r ScheduleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

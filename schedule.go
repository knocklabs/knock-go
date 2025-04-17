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

// Creates one or more schedules for a workflow with the specified recipients,
// timing, and data. Schedules can be one-time or recurring.
func (r *ScheduleService) New(ctx context.Context, body ScheduleNewParams, opts ...option.RequestOption) (res *[]Schedule, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates one or more existing schedules with new timing, data, or other
// properties. All specified schedule IDs will be updated with the same values.
func (r *ScheduleService) Update(ctx context.Context, body ScheduleUpdateParams, opts ...option.RequestOption) (res *[]Schedule, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a paginated list of schedules for the current environment, filtered by
// workflow and optionally by recipients and tenant.
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

// Returns a paginated list of schedules for the current environment, filtered by
// workflow and optionally by recipients and tenant.
func (r *ScheduleService) ListAutoPaging(ctx context.Context, query ScheduleListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Schedule] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes one or more schedules identified by the provided schedule
// IDs. This operation cannot be undone.
func (r *ScheduleService) Delete(ctx context.Context, body ScheduleDeleteParams, opts ...option.RequestOption) (res *[]Schedule, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// A schedule represents a recurring workflow execution.
type Schedule struct {
	// Unique identifier for the schedule.
	ID string `json:"id,required" format:"uuid"`
	// Timestamp when the resource was created.
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient of a notification, which is either a user or an object.
	Recipient Recipient `json:"recipient,required"`
	// The repeat rule for the schedule.
	Repeats []ScheduleRepeatRule `json:"repeats,required"`
	// The timestamp when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The workflow the schedule is applied to.
	Workflow string `json:"workflow,required"`
	// The typename of the schema.
	Typename string `json:"__typename"`
	// A recipient of a notification, which is either a user or an object.
	Actor Recipient `json:"actor,nullable"`
	// An optional map of data to pass into the workflow execution.
	Data map[string]interface{} `json:"data,nullable"`
	// The last occurrence of the schedule.
	LastOccurrenceAt time.Time `json:"last_occurrence_at,nullable" format:"date-time"`
	// The next occurrence of the schedule.
	NextOccurrenceAt time.Time `json:"next_occurrence_at,nullable" format:"date-time"`
	// The tenant to trigger the workflow for. Triggering with a tenant will use any
	// tenant-level overrides associated with the tenant object, and all messages
	// produced from workflow runs will be tagged with the tenant.
	Tenant string       `json:"tenant,nullable"`
	JSON   scheduleJSON `json:"-"`
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

// The repeat rule for the schedule.
type ScheduleRepeatRule struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The frequency of the schedule.
	Frequency ScheduleRepeatRuleFrequency `json:"frequency,required"`
	// The day of the month to repeat the schedule.
	DayOfMonth int64 `json:"day_of_month,nullable"`
	// The days of the week to repeat the schedule.
	Days []ScheduleRepeatRuleDay `json:"days,nullable"`
	// The hour of the day to repeat the schedule.
	Hours int64 `json:"hours,nullable"`
	// The interval of the schedule.
	Interval int64 `json:"interval"`
	// The minute of the hour to repeat the schedule.
	Minutes int64                  `json:"minutes,nullable"`
	JSON    scheduleRepeatRuleJSON `json:"-"`
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

// The frequency of the schedule.
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

// An identifier for a day of the week.
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

// The repeat rule for the schedule.
type ScheduleRepeatRuleParam struct {
	// The typename of the schema.
	Typename param.Field[string] `json:"__typename,required"`
	// The frequency of the schedule.
	Frequency param.Field[ScheduleRepeatRuleFrequency] `json:"frequency,required"`
	// The day of the month to repeat the schedule.
	DayOfMonth param.Field[int64] `json:"day_of_month"`
	// The days of the week to repeat the schedule.
	Days param.Field[[]ScheduleRepeatRuleDay] `json:"days"`
	// The hour of the day to repeat the schedule.
	Hours param.Field[int64] `json:"hours"`
	// The interval of the schedule.
	Interval param.Field[int64] `json:"interval"`
	// The minute of the hour to repeat the schedule.
	Minutes param.Field[int64] `json:"minutes"`
}

func (r ScheduleRepeatRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScheduleNewParams struct {
	// The recipients to trigger the workflow for. Can inline identify users, objects,
	// or use a list of user ids. Cannot exceed 1000 recipients in a single trigger.
	Recipients param.Field[[]ScheduleNewParamsRecipientUnion] `json:"recipients,required"`
	// The repeat rule for the schedule.
	Repeats param.Field[[]ScheduleRepeatRuleParam] `json:"repeats,required"`
	// The key of the workflow.
	Workflow param.Field[string] `json:"workflow,required"`
	// An optional map of data to pass into the workflow execution.
	Data param.Field[map[string]interface{}] `json:"data"`
	// The ending date and time for the schedule.
	EndingAt param.Field[time.Time] `json:"ending_at" format:"date-time"`
	// The starting date and time for the schedule.
	ScheduledAt param.Field[time.Time] `json:"scheduled_at" format:"date-time"`
	// An request to set a tenant inline.
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

// A reference to a recipient object.
type ScheduleNewParamsRecipientsObjectReference struct {
	// An identifier for the recipient object.
	ID param.Field[string] `json:"id"`
	// The collection the recipient object belongs to.
	Collection param.Field[string] `json:"collection"`
}

func (r ScheduleNewParamsRecipientsObjectReference) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScheduleNewParamsRecipientsObjectReference) ImplementsScheduleNewParamsRecipientUnion() {}

type ScheduleUpdateParams struct {
	// A list of schedule IDs.
	ScheduleIDs param.Field[[]string] `json:"schedule_ids,required" format:"uuid"`
	// Specifies a recipient in a request. This can either be a user identifier
	// (string), an inline user request (object), or an inline object request, which is
	// determined by the presence of a `collection` property.
	Actor param.Field[RecipientRequestUnionParam] `json:"actor"`
	// An optional map of data to pass into the workflow execution.
	Data param.Field[map[string]interface{}] `json:"data"`
	// The ending date and time for the schedule.
	EndingAt param.Field[time.Time] `json:"ending_at" format:"date-time"`
	// The repeat rule for the schedule.
	Repeats param.Field[[]ScheduleRepeatRuleParam] `json:"repeats"`
	// The starting date and time for the schedule.
	ScheduledAt param.Field[time.Time] `json:"scheduled_at" format:"date-time"`
	// An request to set a tenant inline.
	Tenant param.Field[InlineTenantRequestUnionParam] `json:"tenant"`
}

func (r ScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScheduleListParams struct {
	// Filter by workflow key.
	Workflow param.Field[string] `query:"workflow,required"`
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// The number of items per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by recipient IDs.
	Recipients param.Field[[]string] `query:"recipients"`
	// Filter by tenant ID.
	Tenant param.Field[string] `query:"tenant"`
}

// URLQuery serializes [ScheduleListParams]'s query parameters as `url.Values`.
func (r ScheduleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScheduleDeleteParams struct {
	// A list of schedule IDs.
	ScheduleIDs param.Field[[]string] `json:"schedule_ids,required"`
}

func (r ScheduleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

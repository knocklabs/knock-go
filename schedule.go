// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/tidwall/gjson"
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
func (r *ScheduleService) New(ctx context.Context, opts ...option.RequestOption) (res *[]ScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Update schedules
func (r *ScheduleService) Update(ctx context.Context, opts ...option.RequestOption) (res *[]ScheduleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// List schedules
func (r *ScheduleService) List(ctx context.Context, query ScheduleListParams, opts ...option.RequestOption) (res *ScheduleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete schedules
func (r *ScheduleService) Delete(ctx context.Context, opts ...option.RequestOption) (res *[]ScheduleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A schedule that represents a recurring workflow execution
type ScheduleNewResponse struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient ScheduleNewResponseRecipient `json:"recipient,required"`
	Repeats   []ScheduleNewResponseRepeat  `json:"repeats,required"`
	UpdatedAt time.Time                    `json:"updated_at,required" format:"date-time"`
	Workflow  string                       `json:"workflow,required"`
	Typename  string                       `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            ScheduleNewResponseActor `json:"actor,nullable"`
	Data             interface{}              `json:"data,nullable"`
	LastOccurrenceAt time.Time                `json:"last_occurrence_at,nullable" format:"date-time"`
	NextOccurrenceAt time.Time                `json:"next_occurrence_at,nullable" format:"date-time"`
	Tenant           string                   `json:"tenant,nullable"`
	JSON             scheduleNewResponseJSON  `json:"-"`
}

// scheduleNewResponseJSON contains the JSON metadata for the struct
// [ScheduleNewResponse]
type scheduleNewResponseJSON struct {
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

func (r *ScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type ScheduleNewResponseRecipient struct {
	ID          string                           `json:"id,required"`
	Typename    string                           `json:"__typename,required"`
	UpdatedAt   time.Time                        `json:"updated_at,required" format:"date-time"`
	Avatar      string                           `json:"avatar,nullable"`
	Collection  string                           `json:"collection"`
	CreatedAt   time.Time                        `json:"created_at,nullable" format:"date-time"`
	Email       string                           `json:"email,nullable" format:"email"`
	Name        string                           `json:"name,nullable"`
	PhoneNumber string                           `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                           `json:"timezone,nullable"`
	JSON        scheduleNewResponseRecipientJSON `json:"-"`
	union       ScheduleNewResponseRecipientUnion
}

// scheduleNewResponseRecipientJSON contains the JSON metadata for the struct
// [ScheduleNewResponseRecipient]
type scheduleNewResponseRecipientJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scheduleNewResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleNewResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleNewResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleNewResponseRecipientUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [ScheduleNewResponseRecipientObject].
func (r ScheduleNewResponseRecipient) AsUnion() ScheduleNewResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [ScheduleNewResponseRecipientObject].
type ScheduleNewResponseRecipientUnion interface {
	implementsScheduleNewResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleNewResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScheduleNewResponseRecipientObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type ScheduleNewResponseRecipientObject struct {
	ID          string                                 `json:"id,required"`
	Typename    string                                 `json:"__typename,required"`
	Collection  string                                 `json:"collection,required"`
	UpdatedAt   time.Time                              `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                              `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                 `json:"-,extras"`
	JSON        scheduleNewResponseRecipientObjectJSON `json:"-"`
}

// scheduleNewResponseRecipientObjectJSON contains the JSON metadata for the struct
// [ScheduleNewResponseRecipientObject]
type scheduleNewResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleNewResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleNewResponseRecipientObject) implementsScheduleNewResponseRecipient() {}

// A schedule repeat rule
type ScheduleNewResponseRepeat struct {
	Typename   string                              `json:"__typename,required"`
	Frequency  ScheduleNewResponseRepeatsFrequency `json:"frequency,required"`
	DayOfMonth int64                               `json:"day_of_month,nullable"`
	Days       []ScheduleNewResponseRepeatsDay     `json:"days,nullable"`
	Hours      int64                               `json:"hours,nullable"`
	Interval   int64                               `json:"interval"`
	Minutes    int64                               `json:"minutes,nullable"`
	JSON       scheduleNewResponseRepeatJSON       `json:"-"`
}

// scheduleNewResponseRepeatJSON contains the JSON metadata for the struct
// [ScheduleNewResponseRepeat]
type scheduleNewResponseRepeatJSON struct {
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

func (r *ScheduleNewResponseRepeat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseRepeatJSON) RawJSON() string {
	return r.raw
}

type ScheduleNewResponseRepeatsFrequency string

const (
	ScheduleNewResponseRepeatsFrequencyDaily   ScheduleNewResponseRepeatsFrequency = "daily"
	ScheduleNewResponseRepeatsFrequencyWeekly  ScheduleNewResponseRepeatsFrequency = "weekly"
	ScheduleNewResponseRepeatsFrequencyMonthly ScheduleNewResponseRepeatsFrequency = "monthly"
	ScheduleNewResponseRepeatsFrequencyHourly  ScheduleNewResponseRepeatsFrequency = "hourly"
)

func (r ScheduleNewResponseRepeatsFrequency) IsKnown() bool {
	switch r {
	case ScheduleNewResponseRepeatsFrequencyDaily, ScheduleNewResponseRepeatsFrequencyWeekly, ScheduleNewResponseRepeatsFrequencyMonthly, ScheduleNewResponseRepeatsFrequencyHourly:
		return true
	}
	return false
}

type ScheduleNewResponseRepeatsDay string

const (
	ScheduleNewResponseRepeatsDayMon ScheduleNewResponseRepeatsDay = "mon"
	ScheduleNewResponseRepeatsDayTue ScheduleNewResponseRepeatsDay = "tue"
	ScheduleNewResponseRepeatsDayWed ScheduleNewResponseRepeatsDay = "wed"
	ScheduleNewResponseRepeatsDayThu ScheduleNewResponseRepeatsDay = "thu"
	ScheduleNewResponseRepeatsDayFri ScheduleNewResponseRepeatsDay = "fri"
	ScheduleNewResponseRepeatsDaySat ScheduleNewResponseRepeatsDay = "sat"
	ScheduleNewResponseRepeatsDaySun ScheduleNewResponseRepeatsDay = "sun"
)

func (r ScheduleNewResponseRepeatsDay) IsKnown() bool {
	switch r {
	case ScheduleNewResponseRepeatsDayMon, ScheduleNewResponseRepeatsDayTue, ScheduleNewResponseRepeatsDayWed, ScheduleNewResponseRepeatsDayThu, ScheduleNewResponseRepeatsDayFri, ScheduleNewResponseRepeatsDaySat, ScheduleNewResponseRepeatsDaySun:
		return true
	}
	return false
}

// A recipient, which is either a user or an object
type ScheduleNewResponseActor struct {
	ID          string                       `json:"id,required"`
	Typename    string                       `json:"__typename,required"`
	UpdatedAt   time.Time                    `json:"updated_at,required" format:"date-time"`
	Avatar      string                       `json:"avatar,nullable"`
	Collection  string                       `json:"collection"`
	CreatedAt   time.Time                    `json:"created_at,nullable" format:"date-time"`
	Email       string                       `json:"email,nullable" format:"email"`
	Name        string                       `json:"name,nullable"`
	PhoneNumber string                       `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                       `json:"timezone,nullable"`
	JSON        scheduleNewResponseActorJSON `json:"-"`
	union       ScheduleNewResponseActorUnion
}

// scheduleNewResponseActorJSON contains the JSON metadata for the struct
// [ScheduleNewResponseActor]
type scheduleNewResponseActorJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scheduleNewResponseActorJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleNewResponseActor) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleNewResponseActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleNewResponseActorUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [ScheduleNewResponseActorObject].
func (r ScheduleNewResponseActor) AsUnion() ScheduleNewResponseActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [ScheduleNewResponseActorObject].
type ScheduleNewResponseActorUnion interface {
	implementsScheduleNewResponseActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleNewResponseActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScheduleNewResponseActorObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type ScheduleNewResponseActorObject struct {
	ID          string                             `json:"id,required"`
	Typename    string                             `json:"__typename,required"`
	Collection  string                             `json:"collection,required"`
	UpdatedAt   time.Time                          `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                          `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}             `json:"-,extras"`
	JSON        scheduleNewResponseActorObjectJSON `json:"-"`
}

// scheduleNewResponseActorObjectJSON contains the JSON metadata for the struct
// [ScheduleNewResponseActorObject]
type scheduleNewResponseActorObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleNewResponseActorObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleNewResponseActorObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleNewResponseActorObject) implementsScheduleNewResponseActor() {}

// A schedule that represents a recurring workflow execution
type ScheduleUpdateResponse struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient ScheduleUpdateResponseRecipient `json:"recipient,required"`
	Repeats   []ScheduleUpdateResponseRepeat  `json:"repeats,required"`
	UpdatedAt time.Time                       `json:"updated_at,required" format:"date-time"`
	Workflow  string                          `json:"workflow,required"`
	Typename  string                          `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            ScheduleUpdateResponseActor `json:"actor,nullable"`
	Data             interface{}                 `json:"data,nullable"`
	LastOccurrenceAt time.Time                   `json:"last_occurrence_at,nullable" format:"date-time"`
	NextOccurrenceAt time.Time                   `json:"next_occurrence_at,nullable" format:"date-time"`
	Tenant           string                      `json:"tenant,nullable"`
	JSON             scheduleUpdateResponseJSON  `json:"-"`
}

// scheduleUpdateResponseJSON contains the JSON metadata for the struct
// [ScheduleUpdateResponse]
type scheduleUpdateResponseJSON struct {
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

func (r *ScheduleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type ScheduleUpdateResponseRecipient struct {
	ID          string                              `json:"id,required"`
	Typename    string                              `json:"__typename,required"`
	UpdatedAt   time.Time                           `json:"updated_at,required" format:"date-time"`
	Avatar      string                              `json:"avatar,nullable"`
	Collection  string                              `json:"collection"`
	CreatedAt   time.Time                           `json:"created_at,nullable" format:"date-time"`
	Email       string                              `json:"email,nullable" format:"email"`
	Name        string                              `json:"name,nullable"`
	PhoneNumber string                              `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                              `json:"timezone,nullable"`
	JSON        scheduleUpdateResponseRecipientJSON `json:"-"`
	union       ScheduleUpdateResponseRecipientUnion
}

// scheduleUpdateResponseRecipientJSON contains the JSON metadata for the struct
// [ScheduleUpdateResponseRecipient]
type scheduleUpdateResponseRecipientJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scheduleUpdateResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleUpdateResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleUpdateResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleUpdateResponseRecipientUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [ScheduleUpdateResponseRecipientObject].
func (r ScheduleUpdateResponseRecipient) AsUnion() ScheduleUpdateResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [ScheduleUpdateResponseRecipientObject].
type ScheduleUpdateResponseRecipientUnion interface {
	implementsScheduleUpdateResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleUpdateResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScheduleUpdateResponseRecipientObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type ScheduleUpdateResponseRecipientObject struct {
	ID          string                                    `json:"id,required"`
	Typename    string                                    `json:"__typename,required"`
	Collection  string                                    `json:"collection,required"`
	UpdatedAt   time.Time                                 `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                 `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
	JSON        scheduleUpdateResponseRecipientObjectJSON `json:"-"`
}

// scheduleUpdateResponseRecipientObjectJSON contains the JSON metadata for the
// struct [ScheduleUpdateResponseRecipientObject]
type scheduleUpdateResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleUpdateResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleUpdateResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleUpdateResponseRecipientObject) implementsScheduleUpdateResponseRecipient() {}

// A schedule repeat rule
type ScheduleUpdateResponseRepeat struct {
	Typename   string                                 `json:"__typename,required"`
	Frequency  ScheduleUpdateResponseRepeatsFrequency `json:"frequency,required"`
	DayOfMonth int64                                  `json:"day_of_month,nullable"`
	Days       []ScheduleUpdateResponseRepeatsDay     `json:"days,nullable"`
	Hours      int64                                  `json:"hours,nullable"`
	Interval   int64                                  `json:"interval"`
	Minutes    int64                                  `json:"minutes,nullable"`
	JSON       scheduleUpdateResponseRepeatJSON       `json:"-"`
}

// scheduleUpdateResponseRepeatJSON contains the JSON metadata for the struct
// [ScheduleUpdateResponseRepeat]
type scheduleUpdateResponseRepeatJSON struct {
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

func (r *ScheduleUpdateResponseRepeat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleUpdateResponseRepeatJSON) RawJSON() string {
	return r.raw
}

type ScheduleUpdateResponseRepeatsFrequency string

const (
	ScheduleUpdateResponseRepeatsFrequencyDaily   ScheduleUpdateResponseRepeatsFrequency = "daily"
	ScheduleUpdateResponseRepeatsFrequencyWeekly  ScheduleUpdateResponseRepeatsFrequency = "weekly"
	ScheduleUpdateResponseRepeatsFrequencyMonthly ScheduleUpdateResponseRepeatsFrequency = "monthly"
	ScheduleUpdateResponseRepeatsFrequencyHourly  ScheduleUpdateResponseRepeatsFrequency = "hourly"
)

func (r ScheduleUpdateResponseRepeatsFrequency) IsKnown() bool {
	switch r {
	case ScheduleUpdateResponseRepeatsFrequencyDaily, ScheduleUpdateResponseRepeatsFrequencyWeekly, ScheduleUpdateResponseRepeatsFrequencyMonthly, ScheduleUpdateResponseRepeatsFrequencyHourly:
		return true
	}
	return false
}

type ScheduleUpdateResponseRepeatsDay string

const (
	ScheduleUpdateResponseRepeatsDayMon ScheduleUpdateResponseRepeatsDay = "mon"
	ScheduleUpdateResponseRepeatsDayTue ScheduleUpdateResponseRepeatsDay = "tue"
	ScheduleUpdateResponseRepeatsDayWed ScheduleUpdateResponseRepeatsDay = "wed"
	ScheduleUpdateResponseRepeatsDayThu ScheduleUpdateResponseRepeatsDay = "thu"
	ScheduleUpdateResponseRepeatsDayFri ScheduleUpdateResponseRepeatsDay = "fri"
	ScheduleUpdateResponseRepeatsDaySat ScheduleUpdateResponseRepeatsDay = "sat"
	ScheduleUpdateResponseRepeatsDaySun ScheduleUpdateResponseRepeatsDay = "sun"
)

func (r ScheduleUpdateResponseRepeatsDay) IsKnown() bool {
	switch r {
	case ScheduleUpdateResponseRepeatsDayMon, ScheduleUpdateResponseRepeatsDayTue, ScheduleUpdateResponseRepeatsDayWed, ScheduleUpdateResponseRepeatsDayThu, ScheduleUpdateResponseRepeatsDayFri, ScheduleUpdateResponseRepeatsDaySat, ScheduleUpdateResponseRepeatsDaySun:
		return true
	}
	return false
}

// A recipient, which is either a user or an object
type ScheduleUpdateResponseActor struct {
	ID          string                          `json:"id,required"`
	Typename    string                          `json:"__typename,required"`
	UpdatedAt   time.Time                       `json:"updated_at,required" format:"date-time"`
	Avatar      string                          `json:"avatar,nullable"`
	Collection  string                          `json:"collection"`
	CreatedAt   time.Time                       `json:"created_at,nullable" format:"date-time"`
	Email       string                          `json:"email,nullable" format:"email"`
	Name        string                          `json:"name,nullable"`
	PhoneNumber string                          `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                          `json:"timezone,nullable"`
	JSON        scheduleUpdateResponseActorJSON `json:"-"`
	union       ScheduleUpdateResponseActorUnion
}

// scheduleUpdateResponseActorJSON contains the JSON metadata for the struct
// [ScheduleUpdateResponseActor]
type scheduleUpdateResponseActorJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scheduleUpdateResponseActorJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleUpdateResponseActor) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleUpdateResponseActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleUpdateResponseActorUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [ScheduleUpdateResponseActorObject].
func (r ScheduleUpdateResponseActor) AsUnion() ScheduleUpdateResponseActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [ScheduleUpdateResponseActorObject].
type ScheduleUpdateResponseActorUnion interface {
	implementsScheduleUpdateResponseActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleUpdateResponseActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScheduleUpdateResponseActorObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type ScheduleUpdateResponseActorObject struct {
	ID          string                                `json:"id,required"`
	Typename    string                                `json:"__typename,required"`
	Collection  string                                `json:"collection,required"`
	UpdatedAt   time.Time                             `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                             `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                `json:"-,extras"`
	JSON        scheduleUpdateResponseActorObjectJSON `json:"-"`
}

// scheduleUpdateResponseActorObjectJSON contains the JSON metadata for the struct
// [ScheduleUpdateResponseActorObject]
type scheduleUpdateResponseActorObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleUpdateResponseActorObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleUpdateResponseActorObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleUpdateResponseActorObject) implementsScheduleUpdateResponseActor() {}

// A paginated list of schedules in a collection.
type ScheduleListResponse struct {
	// The list of schedules
	Entries []ScheduleListResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo ScheduleListResponsePageInfo `json:"page_info,required"`
	JSON     scheduleListResponseJSON     `json:"-"`
}

// scheduleListResponseJSON contains the JSON metadata for the struct
// [ScheduleListResponse]
type scheduleListResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleListResponseJSON) RawJSON() string {
	return r.raw
}

// A schedule that represents a recurring workflow execution
type ScheduleListResponseEntry struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient ScheduleListResponseEntriesRecipient `json:"recipient,required"`
	Repeats   []ScheduleListResponseEntriesRepeat  `json:"repeats,required"`
	UpdatedAt time.Time                            `json:"updated_at,required" format:"date-time"`
	Workflow  string                               `json:"workflow,required"`
	Typename  string                               `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            ScheduleListResponseEntriesActor `json:"actor,nullable"`
	Data             interface{}                      `json:"data,nullable"`
	LastOccurrenceAt time.Time                        `json:"last_occurrence_at,nullable" format:"date-time"`
	NextOccurrenceAt time.Time                        `json:"next_occurrence_at,nullable" format:"date-time"`
	Tenant           string                           `json:"tenant,nullable"`
	JSON             scheduleListResponseEntryJSON    `json:"-"`
}

// scheduleListResponseEntryJSON contains the JSON metadata for the struct
// [ScheduleListResponseEntry]
type scheduleListResponseEntryJSON struct {
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

func (r *ScheduleListResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleListResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type ScheduleListResponseEntriesRecipient struct {
	ID          string                                   `json:"id,required"`
	Typename    string                                   `json:"__typename,required"`
	UpdatedAt   time.Time                                `json:"updated_at,required" format:"date-time"`
	Avatar      string                                   `json:"avatar,nullable"`
	Collection  string                                   `json:"collection"`
	CreatedAt   time.Time                                `json:"created_at,nullable" format:"date-time"`
	Email       string                                   `json:"email,nullable" format:"email"`
	Name        string                                   `json:"name,nullable"`
	PhoneNumber string                                   `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                                   `json:"timezone,nullable"`
	JSON        scheduleListResponseEntriesRecipientJSON `json:"-"`
	union       ScheduleListResponseEntriesRecipientUnion
}

// scheduleListResponseEntriesRecipientJSON contains the JSON metadata for the
// struct [ScheduleListResponseEntriesRecipient]
type scheduleListResponseEntriesRecipientJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scheduleListResponseEntriesRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleListResponseEntriesRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleListResponseEntriesRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleListResponseEntriesRecipientUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [ScheduleListResponseEntriesRecipientObject].
func (r ScheduleListResponseEntriesRecipient) AsUnion() ScheduleListResponseEntriesRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [ScheduleListResponseEntriesRecipientObject].
type ScheduleListResponseEntriesRecipientUnion interface {
	implementsScheduleListResponseEntriesRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleListResponseEntriesRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScheduleListResponseEntriesRecipientObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type ScheduleListResponseEntriesRecipientObject struct {
	ID          string                                         `json:"id,required"`
	Typename    string                                         `json:"__typename,required"`
	Collection  string                                         `json:"collection,required"`
	UpdatedAt   time.Time                                      `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                      `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                         `json:"-,extras"`
	JSON        scheduleListResponseEntriesRecipientObjectJSON `json:"-"`
}

// scheduleListResponseEntriesRecipientObjectJSON contains the JSON metadata for
// the struct [ScheduleListResponseEntriesRecipientObject]
type scheduleListResponseEntriesRecipientObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleListResponseEntriesRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleListResponseEntriesRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleListResponseEntriesRecipientObject) implementsScheduleListResponseEntriesRecipient() {
}

// A schedule repeat rule
type ScheduleListResponseEntriesRepeat struct {
	Typename   string                                      `json:"__typename,required"`
	Frequency  ScheduleListResponseEntriesRepeatsFrequency `json:"frequency,required"`
	DayOfMonth int64                                       `json:"day_of_month,nullable"`
	Days       []ScheduleListResponseEntriesRepeatsDay     `json:"days,nullable"`
	Hours      int64                                       `json:"hours,nullable"`
	Interval   int64                                       `json:"interval"`
	Minutes    int64                                       `json:"minutes,nullable"`
	JSON       scheduleListResponseEntriesRepeatJSON       `json:"-"`
}

// scheduleListResponseEntriesRepeatJSON contains the JSON metadata for the struct
// [ScheduleListResponseEntriesRepeat]
type scheduleListResponseEntriesRepeatJSON struct {
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

func (r *ScheduleListResponseEntriesRepeat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleListResponseEntriesRepeatJSON) RawJSON() string {
	return r.raw
}

type ScheduleListResponseEntriesRepeatsFrequency string

const (
	ScheduleListResponseEntriesRepeatsFrequencyDaily   ScheduleListResponseEntriesRepeatsFrequency = "daily"
	ScheduleListResponseEntriesRepeatsFrequencyWeekly  ScheduleListResponseEntriesRepeatsFrequency = "weekly"
	ScheduleListResponseEntriesRepeatsFrequencyMonthly ScheduleListResponseEntriesRepeatsFrequency = "monthly"
	ScheduleListResponseEntriesRepeatsFrequencyHourly  ScheduleListResponseEntriesRepeatsFrequency = "hourly"
)

func (r ScheduleListResponseEntriesRepeatsFrequency) IsKnown() bool {
	switch r {
	case ScheduleListResponseEntriesRepeatsFrequencyDaily, ScheduleListResponseEntriesRepeatsFrequencyWeekly, ScheduleListResponseEntriesRepeatsFrequencyMonthly, ScheduleListResponseEntriesRepeatsFrequencyHourly:
		return true
	}
	return false
}

type ScheduleListResponseEntriesRepeatsDay string

const (
	ScheduleListResponseEntriesRepeatsDayMon ScheduleListResponseEntriesRepeatsDay = "mon"
	ScheduleListResponseEntriesRepeatsDayTue ScheduleListResponseEntriesRepeatsDay = "tue"
	ScheduleListResponseEntriesRepeatsDayWed ScheduleListResponseEntriesRepeatsDay = "wed"
	ScheduleListResponseEntriesRepeatsDayThu ScheduleListResponseEntriesRepeatsDay = "thu"
	ScheduleListResponseEntriesRepeatsDayFri ScheduleListResponseEntriesRepeatsDay = "fri"
	ScheduleListResponseEntriesRepeatsDaySat ScheduleListResponseEntriesRepeatsDay = "sat"
	ScheduleListResponseEntriesRepeatsDaySun ScheduleListResponseEntriesRepeatsDay = "sun"
)

func (r ScheduleListResponseEntriesRepeatsDay) IsKnown() bool {
	switch r {
	case ScheduleListResponseEntriesRepeatsDayMon, ScheduleListResponseEntriesRepeatsDayTue, ScheduleListResponseEntriesRepeatsDayWed, ScheduleListResponseEntriesRepeatsDayThu, ScheduleListResponseEntriesRepeatsDayFri, ScheduleListResponseEntriesRepeatsDaySat, ScheduleListResponseEntriesRepeatsDaySun:
		return true
	}
	return false
}

// A recipient, which is either a user or an object
type ScheduleListResponseEntriesActor struct {
	ID          string                               `json:"id,required"`
	Typename    string                               `json:"__typename,required"`
	UpdatedAt   time.Time                            `json:"updated_at,required" format:"date-time"`
	Avatar      string                               `json:"avatar,nullable"`
	Collection  string                               `json:"collection"`
	CreatedAt   time.Time                            `json:"created_at,nullable" format:"date-time"`
	Email       string                               `json:"email,nullable" format:"email"`
	Name        string                               `json:"name,nullable"`
	PhoneNumber string                               `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                               `json:"timezone,nullable"`
	JSON        scheduleListResponseEntriesActorJSON `json:"-"`
	union       ScheduleListResponseEntriesActorUnion
}

// scheduleListResponseEntriesActorJSON contains the JSON metadata for the struct
// [ScheduleListResponseEntriesActor]
type scheduleListResponseEntriesActorJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scheduleListResponseEntriesActorJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleListResponseEntriesActor) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleListResponseEntriesActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleListResponseEntriesActorUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [ScheduleListResponseEntriesActorObject].
func (r ScheduleListResponseEntriesActor) AsUnion() ScheduleListResponseEntriesActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [ScheduleListResponseEntriesActorObject].
type ScheduleListResponseEntriesActorUnion interface {
	implementsScheduleListResponseEntriesActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleListResponseEntriesActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScheduleListResponseEntriesActorObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type ScheduleListResponseEntriesActorObject struct {
	ID          string                                     `json:"id,required"`
	Typename    string                                     `json:"__typename,required"`
	Collection  string                                     `json:"collection,required"`
	UpdatedAt   time.Time                                  `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                  `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                     `json:"-,extras"`
	JSON        scheduleListResponseEntriesActorObjectJSON `json:"-"`
}

// scheduleListResponseEntriesActorObjectJSON contains the JSON metadata for the
// struct [ScheduleListResponseEntriesActorObject]
type scheduleListResponseEntriesActorObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleListResponseEntriesActorObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleListResponseEntriesActorObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleListResponseEntriesActorObject) implementsScheduleListResponseEntriesActor() {}

// The information about a paginated result
type ScheduleListResponsePageInfo struct {
	Typename string                           `json:"__typename,required"`
	PageSize int64                            `json:"page_size,required"`
	After    string                           `json:"after,nullable"`
	Before   string                           `json:"before,nullable"`
	JSON     scheduleListResponsePageInfoJSON `json:"-"`
}

// scheduleListResponsePageInfoJSON contains the JSON metadata for the struct
// [ScheduleListResponsePageInfo]
type scheduleListResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleListResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleListResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

// A schedule that represents a recurring workflow execution
type ScheduleDeleteResponse struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient ScheduleDeleteResponseRecipient `json:"recipient,required"`
	Repeats   []ScheduleDeleteResponseRepeat  `json:"repeats,required"`
	UpdatedAt time.Time                       `json:"updated_at,required" format:"date-time"`
	Workflow  string                          `json:"workflow,required"`
	Typename  string                          `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            ScheduleDeleteResponseActor `json:"actor,nullable"`
	Data             interface{}                 `json:"data,nullable"`
	LastOccurrenceAt time.Time                   `json:"last_occurrence_at,nullable" format:"date-time"`
	NextOccurrenceAt time.Time                   `json:"next_occurrence_at,nullable" format:"date-time"`
	Tenant           string                      `json:"tenant,nullable"`
	JSON             scheduleDeleteResponseJSON  `json:"-"`
}

// scheduleDeleteResponseJSON contains the JSON metadata for the struct
// [ScheduleDeleteResponse]
type scheduleDeleteResponseJSON struct {
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

func (r *ScheduleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type ScheduleDeleteResponseRecipient struct {
	ID          string                              `json:"id,required"`
	Typename    string                              `json:"__typename,required"`
	UpdatedAt   time.Time                           `json:"updated_at,required" format:"date-time"`
	Avatar      string                              `json:"avatar,nullable"`
	Collection  string                              `json:"collection"`
	CreatedAt   time.Time                           `json:"created_at,nullable" format:"date-time"`
	Email       string                              `json:"email,nullable" format:"email"`
	Name        string                              `json:"name,nullable"`
	PhoneNumber string                              `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                              `json:"timezone,nullable"`
	JSON        scheduleDeleteResponseRecipientJSON `json:"-"`
	union       ScheduleDeleteResponseRecipientUnion
}

// scheduleDeleteResponseRecipientJSON contains the JSON metadata for the struct
// [ScheduleDeleteResponseRecipient]
type scheduleDeleteResponseRecipientJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scheduleDeleteResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleDeleteResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleDeleteResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleDeleteResponseRecipientUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [ScheduleDeleteResponseRecipientObject].
func (r ScheduleDeleteResponseRecipient) AsUnion() ScheduleDeleteResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [ScheduleDeleteResponseRecipientObject].
type ScheduleDeleteResponseRecipientUnion interface {
	implementsScheduleDeleteResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleDeleteResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScheduleDeleteResponseRecipientObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type ScheduleDeleteResponseRecipientObject struct {
	ID          string                                    `json:"id,required"`
	Typename    string                                    `json:"__typename,required"`
	Collection  string                                    `json:"collection,required"`
	UpdatedAt   time.Time                                 `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                 `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
	JSON        scheduleDeleteResponseRecipientObjectJSON `json:"-"`
}

// scheduleDeleteResponseRecipientObjectJSON contains the JSON metadata for the
// struct [ScheduleDeleteResponseRecipientObject]
type scheduleDeleteResponseRecipientObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleDeleteResponseRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleDeleteResponseRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleDeleteResponseRecipientObject) implementsScheduleDeleteResponseRecipient() {}

// A schedule repeat rule
type ScheduleDeleteResponseRepeat struct {
	Typename   string                                 `json:"__typename,required"`
	Frequency  ScheduleDeleteResponseRepeatsFrequency `json:"frequency,required"`
	DayOfMonth int64                                  `json:"day_of_month,nullable"`
	Days       []ScheduleDeleteResponseRepeatsDay     `json:"days,nullable"`
	Hours      int64                                  `json:"hours,nullable"`
	Interval   int64                                  `json:"interval"`
	Minutes    int64                                  `json:"minutes,nullable"`
	JSON       scheduleDeleteResponseRepeatJSON       `json:"-"`
}

// scheduleDeleteResponseRepeatJSON contains the JSON metadata for the struct
// [ScheduleDeleteResponseRepeat]
type scheduleDeleteResponseRepeatJSON struct {
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

func (r *ScheduleDeleteResponseRepeat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleDeleteResponseRepeatJSON) RawJSON() string {
	return r.raw
}

type ScheduleDeleteResponseRepeatsFrequency string

const (
	ScheduleDeleteResponseRepeatsFrequencyDaily   ScheduleDeleteResponseRepeatsFrequency = "daily"
	ScheduleDeleteResponseRepeatsFrequencyWeekly  ScheduleDeleteResponseRepeatsFrequency = "weekly"
	ScheduleDeleteResponseRepeatsFrequencyMonthly ScheduleDeleteResponseRepeatsFrequency = "monthly"
	ScheduleDeleteResponseRepeatsFrequencyHourly  ScheduleDeleteResponseRepeatsFrequency = "hourly"
)

func (r ScheduleDeleteResponseRepeatsFrequency) IsKnown() bool {
	switch r {
	case ScheduleDeleteResponseRepeatsFrequencyDaily, ScheduleDeleteResponseRepeatsFrequencyWeekly, ScheduleDeleteResponseRepeatsFrequencyMonthly, ScheduleDeleteResponseRepeatsFrequencyHourly:
		return true
	}
	return false
}

type ScheduleDeleteResponseRepeatsDay string

const (
	ScheduleDeleteResponseRepeatsDayMon ScheduleDeleteResponseRepeatsDay = "mon"
	ScheduleDeleteResponseRepeatsDayTue ScheduleDeleteResponseRepeatsDay = "tue"
	ScheduleDeleteResponseRepeatsDayWed ScheduleDeleteResponseRepeatsDay = "wed"
	ScheduleDeleteResponseRepeatsDayThu ScheduleDeleteResponseRepeatsDay = "thu"
	ScheduleDeleteResponseRepeatsDayFri ScheduleDeleteResponseRepeatsDay = "fri"
	ScheduleDeleteResponseRepeatsDaySat ScheduleDeleteResponseRepeatsDay = "sat"
	ScheduleDeleteResponseRepeatsDaySun ScheduleDeleteResponseRepeatsDay = "sun"
)

func (r ScheduleDeleteResponseRepeatsDay) IsKnown() bool {
	switch r {
	case ScheduleDeleteResponseRepeatsDayMon, ScheduleDeleteResponseRepeatsDayTue, ScheduleDeleteResponseRepeatsDayWed, ScheduleDeleteResponseRepeatsDayThu, ScheduleDeleteResponseRepeatsDayFri, ScheduleDeleteResponseRepeatsDaySat, ScheduleDeleteResponseRepeatsDaySun:
		return true
	}
	return false
}

// A recipient, which is either a user or an object
type ScheduleDeleteResponseActor struct {
	ID          string                          `json:"id,required"`
	Typename    string                          `json:"__typename,required"`
	UpdatedAt   time.Time                       `json:"updated_at,required" format:"date-time"`
	Avatar      string                          `json:"avatar,nullable"`
	Collection  string                          `json:"collection"`
	CreatedAt   time.Time                       `json:"created_at,nullable" format:"date-time"`
	Email       string                          `json:"email,nullable" format:"email"`
	Name        string                          `json:"name,nullable"`
	PhoneNumber string                          `json:"phone_number,nullable" format:"phone-number"`
	Timezone    string                          `json:"timezone,nullable"`
	JSON        scheduleDeleteResponseActorJSON `json:"-"`
	union       ScheduleDeleteResponseActorUnion
}

// scheduleDeleteResponseActorJSON contains the JSON metadata for the struct
// [ScheduleDeleteResponseActor]
type scheduleDeleteResponseActorJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scheduleDeleteResponseActorJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleDeleteResponseActor) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleDeleteResponseActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleDeleteResponseActorUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [ScheduleDeleteResponseActorObject].
func (r ScheduleDeleteResponseActor) AsUnion() ScheduleDeleteResponseActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [ScheduleDeleteResponseActorObject].
type ScheduleDeleteResponseActorUnion interface {
	implementsScheduleDeleteResponseActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleDeleteResponseActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScheduleDeleteResponseActorObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type ScheduleDeleteResponseActorObject struct {
	ID          string                                `json:"id,required"`
	Typename    string                                `json:"__typename,required"`
	Collection  string                                `json:"collection,required"`
	UpdatedAt   time.Time                             `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                             `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                `json:"-,extras"`
	JSON        scheduleDeleteResponseActorObjectJSON `json:"-"`
}

// scheduleDeleteResponseActorObjectJSON contains the JSON metadata for the struct
// [ScheduleDeleteResponseActorObject]
type scheduleDeleteResponseActorObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleDeleteResponseActorObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleDeleteResponseActorObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScheduleDeleteResponseActorObject) implementsScheduleDeleteResponseActor() {}

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
// Satisfied by [shared.UnionString], [ScheduleListParamsRecipientsObject].
type ScheduleListParamsRecipientUnion interface {
	ImplementsScheduleListParamsRecipientUnion()
}

// An object reference to a recipient
type ScheduleListParamsRecipientsObject struct {
	// An object identifier
	ID param.Field[string] `query:"id,required"`
	// The collection the object belongs to
	Collection param.Field[string] `query:"collection,required"`
}

// URLQuery serializes [ScheduleListParamsRecipientsObject]'s query parameters as
// `url.Values`.
func (r ScheduleListParamsRecipientsObject) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (r ScheduleListParamsRecipientsObject) ImplementsScheduleListParamsRecipientUnion() {}

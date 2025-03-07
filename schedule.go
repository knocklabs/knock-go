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
	"github.com/stainless-sdks/knock-go/packages/pagination"
	"github.com/stainless-sdks/knock-go/shared"
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
func (r *ScheduleService) New(ctx context.Context, body ScheduleNewParams, opts ...option.RequestOption) (res *[]ScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update schedules
func (r *ScheduleService) Update(ctx context.Context, body ScheduleUpdateParams, opts ...option.RequestOption) (res *[]ScheduleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List schedules
func (r *ScheduleService) List(ctx context.Context, query ScheduleListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[ScheduleListResponse], err error) {
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
func (r *ScheduleService) ListAutoPaging(ctx context.Context, query ScheduleListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[ScheduleListResponse] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete schedules
func (r *ScheduleService) Delete(ctx context.Context, body ScheduleDeleteParams, opts ...option.RequestOption) (res *[]ScheduleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// A schedule that represents a recurring workflow execution
type ScheduleNewResponse struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient ScheduleNewResponseRecipient `json:"recipient,required"`
	Repeats   []shared.ScheduleRepeatRule  `json:"repeats,required"`
	UpdatedAt time.Time                    `json:"updated_at,required" format:"date-time"`
	Workflow  string                       `json:"workflow,required"`
	Typename  string                       `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            ScheduleNewResponseActor `json:"actor,nullable"`
	Data             map[string]interface{}   `json:"data,nullable"`
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
	Email       string                           `json:"email,nullable"`
	Name        string                           `json:"name,nullable"`
	PhoneNumber string                           `json:"phone_number,nullable"`
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
// Possible runtime types of the union are [User], [shared.Object].
func (r ScheduleNewResponseRecipient) AsUnion() ScheduleNewResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ScheduleNewResponseRecipientUnion interface {
	ImplementsScheduleNewResponseRecipient()
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
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A recipient, which is either a user or an object
type ScheduleNewResponseActor struct {
	ID          string                       `json:"id,required"`
	Typename    string                       `json:"__typename,required"`
	UpdatedAt   time.Time                    `json:"updated_at,required" format:"date-time"`
	Avatar      string                       `json:"avatar,nullable"`
	Collection  string                       `json:"collection"`
	CreatedAt   time.Time                    `json:"created_at,nullable" format:"date-time"`
	Email       string                       `json:"email,nullable"`
	Name        string                       `json:"name,nullable"`
	PhoneNumber string                       `json:"phone_number,nullable"`
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
// Possible runtime types of the union are [User], [shared.Object].
func (r ScheduleNewResponseActor) AsUnion() ScheduleNewResponseActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ScheduleNewResponseActorUnion interface {
	ImplementsScheduleNewResponseActor()
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
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A schedule that represents a recurring workflow execution
type ScheduleUpdateResponse struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient ScheduleUpdateResponseRecipient `json:"recipient,required"`
	Repeats   []shared.ScheduleRepeatRule     `json:"repeats,required"`
	UpdatedAt time.Time                       `json:"updated_at,required" format:"date-time"`
	Workflow  string                          `json:"workflow,required"`
	Typename  string                          `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            ScheduleUpdateResponseActor `json:"actor,nullable"`
	Data             map[string]interface{}      `json:"data,nullable"`
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
	Email       string                              `json:"email,nullable"`
	Name        string                              `json:"name,nullable"`
	PhoneNumber string                              `json:"phone_number,nullable"`
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
// Possible runtime types of the union are [User], [shared.Object].
func (r ScheduleUpdateResponseRecipient) AsUnion() ScheduleUpdateResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ScheduleUpdateResponseRecipientUnion interface {
	ImplementsScheduleUpdateResponseRecipient()
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
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A recipient, which is either a user or an object
type ScheduleUpdateResponseActor struct {
	ID          string                          `json:"id,required"`
	Typename    string                          `json:"__typename,required"`
	UpdatedAt   time.Time                       `json:"updated_at,required" format:"date-time"`
	Avatar      string                          `json:"avatar,nullable"`
	Collection  string                          `json:"collection"`
	CreatedAt   time.Time                       `json:"created_at,nullable" format:"date-time"`
	Email       string                          `json:"email,nullable"`
	Name        string                          `json:"name,nullable"`
	PhoneNumber string                          `json:"phone_number,nullable"`
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
// Possible runtime types of the union are [User], [shared.Object].
func (r ScheduleUpdateResponseActor) AsUnion() ScheduleUpdateResponseActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ScheduleUpdateResponseActorUnion interface {
	ImplementsScheduleUpdateResponseActor()
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
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A schedule that represents a recurring workflow execution
type ScheduleListResponse struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient ScheduleListResponseRecipient `json:"recipient,required"`
	Repeats   []shared.ScheduleRepeatRule   `json:"repeats,required"`
	UpdatedAt time.Time                     `json:"updated_at,required" format:"date-time"`
	Workflow  string                        `json:"workflow,required"`
	Typename  string                        `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            ScheduleListResponseActor `json:"actor,nullable"`
	Data             map[string]interface{}    `json:"data,nullable"`
	LastOccurrenceAt time.Time                 `json:"last_occurrence_at,nullable" format:"date-time"`
	NextOccurrenceAt time.Time                 `json:"next_occurrence_at,nullable" format:"date-time"`
	Tenant           string                    `json:"tenant,nullable"`
	JSON             scheduleListResponseJSON  `json:"-"`
}

// scheduleListResponseJSON contains the JSON metadata for the struct
// [ScheduleListResponse]
type scheduleListResponseJSON struct {
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

func (r *ScheduleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleListResponseJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type ScheduleListResponseRecipient struct {
	ID          string                            `json:"id,required"`
	Typename    string                            `json:"__typename,required"`
	UpdatedAt   time.Time                         `json:"updated_at,required" format:"date-time"`
	Avatar      string                            `json:"avatar,nullable"`
	Collection  string                            `json:"collection"`
	CreatedAt   time.Time                         `json:"created_at,nullable" format:"date-time"`
	Email       string                            `json:"email,nullable"`
	Name        string                            `json:"name,nullable"`
	PhoneNumber string                            `json:"phone_number,nullable"`
	Timezone    string                            `json:"timezone,nullable"`
	JSON        scheduleListResponseRecipientJSON `json:"-"`
	union       ScheduleListResponseRecipientUnion
}

// scheduleListResponseRecipientJSON contains the JSON metadata for the struct
// [ScheduleListResponseRecipient]
type scheduleListResponseRecipientJSON struct {
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

func (r scheduleListResponseRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleListResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleListResponseRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleListResponseRecipientUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r ScheduleListResponseRecipient) AsUnion() ScheduleListResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ScheduleListResponseRecipientUnion interface {
	ImplementsScheduleListResponseRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleListResponseRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A recipient, which is either a user or an object
type ScheduleListResponseActor struct {
	ID          string                        `json:"id,required"`
	Typename    string                        `json:"__typename,required"`
	UpdatedAt   time.Time                     `json:"updated_at,required" format:"date-time"`
	Avatar      string                        `json:"avatar,nullable"`
	Collection  string                        `json:"collection"`
	CreatedAt   time.Time                     `json:"created_at,nullable" format:"date-time"`
	Email       string                        `json:"email,nullable"`
	Name        string                        `json:"name,nullable"`
	PhoneNumber string                        `json:"phone_number,nullable"`
	Timezone    string                        `json:"timezone,nullable"`
	JSON        scheduleListResponseActorJSON `json:"-"`
	union       ScheduleListResponseActorUnion
}

// scheduleListResponseActorJSON contains the JSON metadata for the struct
// [ScheduleListResponseActor]
type scheduleListResponseActorJSON struct {
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

func (r scheduleListResponseActorJSON) RawJSON() string {
	return r.raw
}

func (r *ScheduleListResponseActor) UnmarshalJSON(data []byte) (err error) {
	*r = ScheduleListResponseActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScheduleListResponseActorUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [User], [shared.Object].
func (r ScheduleListResponseActor) AsUnion() ScheduleListResponseActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ScheduleListResponseActorUnion interface {
	ImplementsScheduleListResponseActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleListResponseActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A schedule that represents a recurring workflow execution
type ScheduleDeleteResponse struct {
	ID         string    `json:"id,required" format:"uuid"`
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient ScheduleDeleteResponseRecipient `json:"recipient,required"`
	Repeats   []shared.ScheduleRepeatRule     `json:"repeats,required"`
	UpdatedAt time.Time                       `json:"updated_at,required" format:"date-time"`
	Workflow  string                          `json:"workflow,required"`
	Typename  string                          `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor            ScheduleDeleteResponseActor `json:"actor,nullable"`
	Data             map[string]interface{}      `json:"data,nullable"`
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
	Email       string                              `json:"email,nullable"`
	Name        string                              `json:"name,nullable"`
	PhoneNumber string                              `json:"phone_number,nullable"`
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
// Possible runtime types of the union are [User], [shared.Object].
func (r ScheduleDeleteResponseRecipient) AsUnion() ScheduleDeleteResponseRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ScheduleDeleteResponseRecipientUnion interface {
	ImplementsScheduleDeleteResponseRecipient()
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
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

// A recipient, which is either a user or an object
type ScheduleDeleteResponseActor struct {
	ID          string                          `json:"id,required"`
	Typename    string                          `json:"__typename,required"`
	UpdatedAt   time.Time                       `json:"updated_at,required" format:"date-time"`
	Avatar      string                          `json:"avatar,nullable"`
	Collection  string                          `json:"collection"`
	CreatedAt   time.Time                       `json:"created_at,nullable" format:"date-time"`
	Email       string                          `json:"email,nullable"`
	Name        string                          `json:"name,nullable"`
	PhoneNumber string                          `json:"phone_number,nullable"`
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
// Possible runtime types of the union are [User], [shared.Object].
func (r ScheduleDeleteResponseActor) AsUnion() ScheduleDeleteResponseActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [shared.Object].
type ScheduleDeleteResponseActorUnion interface {
	ImplementsScheduleDeleteResponseActor()
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
			Type:       reflect.TypeOf(shared.Object{}),
		},
	)
}

type ScheduleNewParams struct {
	Recipients  param.Field[[]ScheduleNewParamsRecipientUnion] `json:"recipients,required"`
	Repeats     param.Field[[]shared.ScheduleRepeatRuleParam]  `json:"repeats,required"`
	Workflow    param.Field[string]                            `json:"workflow,required"`
	Data        param.Field[map[string]interface{}]            `json:"data"`
	EndingAt    param.Field[time.Time]                         `json:"ending_at" format:"date-time"`
	ScheduledAt param.Field[time.Time]                         `json:"scheduled_at" format:"date-time"`
	// An inline tenant request
	Tenant param.Field[ScheduleNewParamsTenantUnion] `json:"tenant"`
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

// An inline tenant request
//
// Satisfied by [shared.UnionString], [shared.TenantRequestParam].
type ScheduleNewParamsTenantUnion interface {
	ImplementsScheduleNewParamsTenantUnion()
}

type ScheduleUpdateParams struct {
	ScheduleIDs param.Field[[]string] `json:"schedule_ids,required" format:"uuid"`
	// Specifies a recipient in a request. This can either be a user identifier
	// (string), an inline user request (object), or an inline object request, which is
	// determined by the presence of a `collection` property.
	Actor       param.Field[ScheduleUpdateParamsActorUnion]   `json:"actor"`
	Data        param.Field[map[string]interface{}]           `json:"data"`
	EndingAt    param.Field[time.Time]                        `json:"ending_at" format:"date-time"`
	Repeats     param.Field[[]shared.ScheduleRepeatRuleParam] `json:"repeats"`
	ScheduledAt param.Field[time.Time]                        `json:"scheduled_at" format:"date-time"`
	// An inline tenant request
	Tenant param.Field[ScheduleUpdateParamsTenantUnion] `json:"tenant"`
}

func (r ScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type ScheduleUpdateParamsActor struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID          param.Field[string]      `json:"id,required"`
	ChannelData param.Field[interface{}] `json:"channel_data"`
	Collection  param.Field[string]      `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt   param.Field[time.Time]   `json:"created_at" format:"date-time"`
	Preferences param.Field[interface{}] `json:"preferences"`
}

func (r ScheduleUpdateParamsActor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScheduleUpdateParamsActor) ImplementsScheduleUpdateParamsActorUnion() {}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// Satisfied by [shared.UnionString], [shared.InlineIdentifyUserRequestParam],
// [ScheduleUpdateParamsActorInlineIdentifyObjectRequest],
// [ScheduleUpdateParamsActor].
type ScheduleUpdateParamsActorUnion interface {
	ImplementsScheduleUpdateParamsActorUnion()
}

// Inline identifies a custom object belonging to a collection
type ScheduleUpdateParamsActorInlineIdentifyObjectRequest struct {
	ID         param.Field[string] `json:"id,required"`
	Collection param.Field[string] `json:"collection,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]shared.ChannelDataRequestParam] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                                 `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]shared.PreferenceSetRequestParam] `json:"preferences"`
	ExtraFields map[string]interface{}                                   `json:"-,extras"`
}

func (r ScheduleUpdateParamsActorInlineIdentifyObjectRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScheduleUpdateParamsActorInlineIdentifyObjectRequest) ImplementsScheduleUpdateParamsActorUnion() {
}

// An inline tenant request
//
// Satisfied by [shared.UnionString], [shared.TenantRequestParam].
type ScheduleUpdateParamsTenantUnion interface {
	ImplementsScheduleUpdateParamsTenantUnion()
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

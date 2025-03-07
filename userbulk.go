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
	"github.com/stainless-sdks/knock-go/shared"
)

// UserBulkService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserBulkService] method instead.
type UserBulkService struct {
	Options []option.RequestOption
}

// NewUserBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserBulkService(opts ...option.RequestOption) (r *UserBulkService) {
	r = &UserBulkService{}
	r.Options = opts
	return
}

// Bulk delete users
func (r *UserBulkService) Delete(ctx context.Context, params UserBulkDeleteParams, opts ...option.RequestOption) (res *UserBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/bulk/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Bulk identifies users
func (r *UserBulkService) Identify(ctx context.Context, body UserBulkIdentifyParams, opts ...option.RequestOption) (res *UserBulkIdentifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/bulk/identify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk set preferences
func (r *UserBulkService) SetPreferences(ctx context.Context, body UserBulkSetPreferencesParams, opts ...option.RequestOption) (res *UserBulkSetPreferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/users/bulk/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A bulk operation entity
type UserBulkDeleteResponse struct {
	ID                 string                       `json:"id,required" format:"uuid"`
	Typename           string                       `json:"__typename,required"`
	EstimatedTotalRows int64                        `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                    `json:"inserted_at,required" format:"date-time"`
	Name               string                       `json:"name,required"`
	ProcessedRows      int64                        `json:"processed_rows,required"`
	Status             UserBulkDeleteResponseStatus `json:"status,required"`
	SuccessCount       int64                        `json:"success_count,required"`
	UpdatedAt          time.Time                    `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                    `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                        `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []UserBulkDeleteResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                         `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                         `json:"started_at,nullable" format:"date-time"`
	JSON       userBulkDeleteResponseJSON        `json:"-"`
}

// userBulkDeleteResponseJSON contains the JSON metadata for the struct
// [UserBulkDeleteResponse]
type userBulkDeleteResponseJSON struct {
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

func (r *UserBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type UserBulkDeleteResponseStatus string

const (
	UserBulkDeleteResponseStatusQueued     UserBulkDeleteResponseStatus = "queued"
	UserBulkDeleteResponseStatusProcessing UserBulkDeleteResponseStatus = "processing"
	UserBulkDeleteResponseStatusCompleted  UserBulkDeleteResponseStatus = "completed"
	UserBulkDeleteResponseStatusFailed     UserBulkDeleteResponseStatus = "failed"
)

func (r UserBulkDeleteResponseStatus) IsKnown() bool {
	switch r {
	case UserBulkDeleteResponseStatusQueued, UserBulkDeleteResponseStatusProcessing, UserBulkDeleteResponseStatusCompleted, UserBulkDeleteResponseStatusFailed:
		return true
	}
	return false
}

type UserBulkDeleteResponseErrorItem struct {
	ID         string                              `json:"id,required"`
	Collection string                              `json:"collection,nullable"`
	JSON       userBulkDeleteResponseErrorItemJSON `json:"-"`
}

// userBulkDeleteResponseErrorItemJSON contains the JSON metadata for the struct
// [UserBulkDeleteResponseErrorItem]
type userBulkDeleteResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBulkDeleteResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkDeleteResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type UserBulkIdentifyResponse struct {
	ID                 string                         `json:"id,required" format:"uuid"`
	Typename           string                         `json:"__typename,required"`
	EstimatedTotalRows int64                          `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                      `json:"inserted_at,required" format:"date-time"`
	Name               string                         `json:"name,required"`
	ProcessedRows      int64                          `json:"processed_rows,required"`
	Status             UserBulkIdentifyResponseStatus `json:"status,required"`
	SuccessCount       int64                          `json:"success_count,required"`
	UpdatedAt          time.Time                      `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                      `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                          `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []UserBulkIdentifyResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                           `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                           `json:"started_at,nullable" format:"date-time"`
	JSON       userBulkIdentifyResponseJSON        `json:"-"`
}

// userBulkIdentifyResponseJSON contains the JSON metadata for the struct
// [UserBulkIdentifyResponse]
type userBulkIdentifyResponseJSON struct {
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

func (r *UserBulkIdentifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkIdentifyResponseJSON) RawJSON() string {
	return r.raw
}

type UserBulkIdentifyResponseStatus string

const (
	UserBulkIdentifyResponseStatusQueued     UserBulkIdentifyResponseStatus = "queued"
	UserBulkIdentifyResponseStatusProcessing UserBulkIdentifyResponseStatus = "processing"
	UserBulkIdentifyResponseStatusCompleted  UserBulkIdentifyResponseStatus = "completed"
	UserBulkIdentifyResponseStatusFailed     UserBulkIdentifyResponseStatus = "failed"
)

func (r UserBulkIdentifyResponseStatus) IsKnown() bool {
	switch r {
	case UserBulkIdentifyResponseStatusQueued, UserBulkIdentifyResponseStatusProcessing, UserBulkIdentifyResponseStatusCompleted, UserBulkIdentifyResponseStatusFailed:
		return true
	}
	return false
}

type UserBulkIdentifyResponseErrorItem struct {
	ID         string                                `json:"id,required"`
	Collection string                                `json:"collection,nullable"`
	JSON       userBulkIdentifyResponseErrorItemJSON `json:"-"`
}

// userBulkIdentifyResponseErrorItemJSON contains the JSON metadata for the struct
// [UserBulkIdentifyResponseErrorItem]
type userBulkIdentifyResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBulkIdentifyResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkIdentifyResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type UserBulkSetPreferencesResponse struct {
	ID                 string                               `json:"id,required" format:"uuid"`
	Typename           string                               `json:"__typename,required"`
	EstimatedTotalRows int64                                `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                            `json:"inserted_at,required" format:"date-time"`
	Name               string                               `json:"name,required"`
	ProcessedRows      int64                                `json:"processed_rows,required"`
	Status             UserBulkSetPreferencesResponseStatus `json:"status,required"`
	SuccessCount       int64                                `json:"success_count,required"`
	UpdatedAt          time.Time                            `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                            `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                                `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []UserBulkSetPreferencesResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                                 `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                                 `json:"started_at,nullable" format:"date-time"`
	JSON       userBulkSetPreferencesResponseJSON        `json:"-"`
}

// userBulkSetPreferencesResponseJSON contains the JSON metadata for the struct
// [UserBulkSetPreferencesResponse]
type userBulkSetPreferencesResponseJSON struct {
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

func (r *UserBulkSetPreferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkSetPreferencesResponseJSON) RawJSON() string {
	return r.raw
}

type UserBulkSetPreferencesResponseStatus string

const (
	UserBulkSetPreferencesResponseStatusQueued     UserBulkSetPreferencesResponseStatus = "queued"
	UserBulkSetPreferencesResponseStatusProcessing UserBulkSetPreferencesResponseStatus = "processing"
	UserBulkSetPreferencesResponseStatusCompleted  UserBulkSetPreferencesResponseStatus = "completed"
	UserBulkSetPreferencesResponseStatusFailed     UserBulkSetPreferencesResponseStatus = "failed"
)

func (r UserBulkSetPreferencesResponseStatus) IsKnown() bool {
	switch r {
	case UserBulkSetPreferencesResponseStatusQueued, UserBulkSetPreferencesResponseStatusProcessing, UserBulkSetPreferencesResponseStatusCompleted, UserBulkSetPreferencesResponseStatusFailed:
		return true
	}
	return false
}

type UserBulkSetPreferencesResponseErrorItem struct {
	ID         string                                      `json:"id,required"`
	Collection string                                      `json:"collection,nullable"`
	JSON       userBulkSetPreferencesResponseErrorItemJSON `json:"-"`
}

// userBulkSetPreferencesResponseErrorItemJSON contains the JSON metadata for the
// struct [UserBulkSetPreferencesResponseErrorItem]
type userBulkSetPreferencesResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBulkSetPreferencesResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBulkSetPreferencesResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

type UserBulkDeleteParams struct {
	// The IDs of the users to delete
	QueryUserIDs param.Field[[]string] `query:"user_ids,required"`
	BodyUserIDs  param.Field[[]string] `json:"user_ids,required"`
}

func (r UserBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [UserBulkDeleteParams]'s query parameters as `url.Values`.
func (r UserBulkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserBulkIdentifyParams struct {
	Users param.Field[[]shared.InlineIdentifyUserRequestParam] `json:"users,required"`
}

func (r UserBulkIdentifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParams struct {
	// Set preferences for a recipient
	Preferences param.Field[shared.PreferenceSetRequestParam] `json:"preferences,required"`
	UserIDs     param.Field[[]string]                         `json:"user_ids,required"`
}

func (r UserBulkSetPreferencesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

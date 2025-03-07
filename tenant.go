// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/packages/pagination"
	"github.com/stainless-sdks/knock-go/shared"
)

// TenantService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantService] method instead.
type TenantService struct {
	Options []option.RequestOption
	Bulk    *TenantBulkService
}

// NewTenantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenantService(opts ...option.RequestOption) (r *TenantService) {
	r = &TenantService{}
	r.Options = opts
	r.Bulk = NewTenantBulkService(opts...)
	return
}

// List tenants
func (r *TenantService) List(ctx context.Context, query TenantListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[TenantListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/tenants"
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

// List tenants
func (r *TenantService) ListAutoPaging(ctx context.Context, query TenantListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[TenantListResponse] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a tenant
func (r *TenantService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tenants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a tenant
func (r *TenantService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TenantGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tenants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set a tenant
func (r *TenantService) Set(ctx context.Context, id string, body TenantSetParams, opts ...option.RequestOption) (res *TenantSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tenants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// A tenant entity
type TenantListResponse struct {
	ID          string                 `json:"id,required"`
	Typename    string                 `json:"__typename,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        tenantListResponseJSON `json:"-"`
}

// tenantListResponseJSON contains the JSON metadata for the struct
// [TenantListResponse]
type tenantListResponseJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseJSON) RawJSON() string {
	return r.raw
}

// A tenant entity
type TenantGetResponse struct {
	ID          string                 `json:"id,required"`
	Typename    string                 `json:"__typename,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        tenantGetResponseJSON  `json:"-"`
}

// tenantGetResponseJSON contains the JSON metadata for the struct
// [TenantGetResponse]
type tenantGetResponseJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantGetResponseJSON) RawJSON() string {
	return r.raw
}

// A tenant entity
type TenantSetResponse struct {
	ID          string                 `json:"id,required"`
	Typename    string                 `json:"__typename,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        tenantSetResponseJSON  `json:"-"`
}

// tenantSetResponseJSON contains the JSON metadata for the struct
// [TenantSetResponse]
type tenantSetResponseJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantSetResponseJSON) RawJSON() string {
	return r.raw
}

type TenantListParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [TenantListParams]'s query parameters as `url.Values`.
func (r TenantListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantSetParams struct {
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[shared.InlineChannelDataRequestParam] `json:"channel_data"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[shared.InlinePreferenceSetRequestParam] `json:"preferences"`
	Settings    param.Field[TenantSetParamsSettings]                `json:"settings"`
}

func (r TenantSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettings struct {
	Branding param.Field[TenantSetParamsSettingsBranding] `json:"branding"`
	// Set preferences for a recipient
	PreferenceSet param.Field[shared.PreferenceSetRequestParam] `json:"preference_set"`
}

func (r TenantSetParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantSetParamsSettingsBranding struct {
	IconURL              param.Field[string] `json:"icon_url"`
	LogoURL              param.Field[string] `json:"logo_url"`
	PrimaryColor         param.Field[string] `json:"primary_color"`
	PrimaryColorContrast param.Field[string] `json:"primary_color_contrast"`
}

func (r TenantSetParamsSettingsBranding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

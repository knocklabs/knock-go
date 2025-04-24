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

// List tenants for the current environment.
func (r *TenantService) List(ctx context.Context, query TenantListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Tenant], err error) {
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

// List tenants for the current environment.
func (r *TenantService) ListAutoPaging(ctx context.Context, query TenantListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Tenant] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a tenant and all associated data. This operation cannot be undone.
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

// Get a tenant by ID.
func (r *TenantService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Tenant, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tenants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets a tenant within an environment, performing an upsert operation. Any
// existing properties will be merged with the incoming properties.
func (r *TenantService) Set(ctx context.Context, id string, body TenantSetParams, opts ...option.RequestOption) (res *Tenant, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tenants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// An request to set a tenant inline.
//
// Satisfied by [shared.UnionString], [TenantRequestParam].
type InlineTenantRequestUnionParam interface {
	ImplementsInlineTenantRequestUnionParam()
}

// A tenant entity.
type Tenant struct {
	// The unique identifier for the tenant.
	ID string `json:"id,required"`
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// An optional name for the tenant.
	Name string `json:"name,nullable"`
	// The settings for the tenant. Includes branding and preference set.
	Settings    TenantSettings         `json:"settings,nullable"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        tenantJSON             `json:"-"`
}

// tenantJSON contains the JSON metadata for the struct [Tenant]
type tenantJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Name        apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tenant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantJSON) RawJSON() string {
	return r.raw
}

// The settings for the tenant. Includes branding and preference set.
type TenantSettings struct {
	// The branding for the tenant.
	Branding TenantSettingsBranding `json:"branding,nullable"`
	// A preference set represents a specific set of notification preferences for a
	// recipient. A recipient can have multiple preference sets.
	PreferenceSet PreferenceSet      `json:"preference_set,nullable"`
	JSON          tenantSettingsJSON `json:"-"`
}

// tenantSettingsJSON contains the JSON metadata for the struct [TenantSettings]
type tenantSettingsJSON struct {
	Branding      apijson.Field
	PreferenceSet apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TenantSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantSettingsJSON) RawJSON() string {
	return r.raw
}

// The branding for the tenant.
type TenantSettingsBranding struct {
	// The icon URL for the tenant. Must point to a valid image with an image MIME
	// type.
	IconURL string `json:"icon_url,nullable" format:"uri"`
	// The logo URL for the tenant. Must point to a valid image with an image MIME
	// type.
	LogoURL string `json:"logo_url,nullable" format:"uri"`
	// The primary color for the tenant, provided as a hex value.
	PrimaryColor string `json:"primary_color,nullable"`
	// The primary color contrast for the tenant, provided as a hex value.
	PrimaryColorContrast string                     `json:"primary_color_contrast,nullable"`
	JSON                 tenantSettingsBrandingJSON `json:"-"`
}

// tenantSettingsBrandingJSON contains the JSON metadata for the struct
// [TenantSettingsBranding]
type tenantSettingsBrandingJSON struct {
	IconURL              apijson.Field
	LogoURL              apijson.Field
	PrimaryColor         apijson.Field
	PrimaryColorContrast apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TenantSettingsBranding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantSettingsBrandingJSON) RawJSON() string {
	return r.raw
}

// A tenant to be set in the system. You can supply any additional properties on
// the tenant object.
type TenantRequestParam struct {
	// The unique identifier for the tenant.
	ID param.Field[string] `json:"id,required"`
	// A request to set channel data for a type of channel inline.
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
	// The settings for the tenant. Includes branding and preference set.
	Settings    param.Field[TenantRequestSettingsParam] `json:"settings"`
	ExtraFields map[string]interface{}                  `json:"-,extras"`
}

func (r TenantRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TenantRequestParam) ImplementsInlineTenantRequestUnionParam() {}

// The settings for the tenant. Includes branding and preference set.
type TenantRequestSettingsParam struct {
	// The branding for the tenant.
	Branding param.Field[TenantRequestSettingsBrandingParam] `json:"branding"`
	// A request to set a preference set for a recipient.
	PreferenceSet param.Field[PreferenceSetRequestParam] `json:"preference_set"`
}

func (r TenantRequestSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The branding for the tenant.
type TenantRequestSettingsBrandingParam struct {
	// The icon URL for the tenant. Must point to a valid image with an image MIME
	// type.
	IconURL param.Field[string] `json:"icon_url"`
	// The logo URL for the tenant. Must point to a valid image with an image MIME
	// type.
	LogoURL param.Field[string] `json:"logo_url"`
	// The primary color for the tenant, provided as a hex value.
	PrimaryColor param.Field[string] `json:"primary_color"`
	// The primary color contrast for the tenant, provided as a hex value.
	PrimaryColorContrast param.Field[string] `json:"primary_color_contrast"`
}

func (r TenantRequestSettingsBrandingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TenantListParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// Filter tenants by name.
	Name param.Field[string] `query:"name"`
	// The number of items per page.
	PageSize param.Field[int64] `query:"page_size"`
	// Filter tenants by ID.
	TenantID param.Field[string] `query:"tenant_id"`
}

// URLQuery serializes [TenantListParams]'s query parameters as `url.Values`.
func (r TenantListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantSetParams struct {
	// A request to set channel data for a type of channel inline.
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
	// The settings for the tenant. Includes branding and preference set.
	Settings param.Field[TenantSetParamsSettings] `json:"settings"`
}

func (r TenantSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settings for the tenant. Includes branding and preference set.
type TenantSetParamsSettings struct {
	// The branding for the tenant.
	Branding param.Field[TenantSetParamsSettingsBranding] `json:"branding"`
	// A request to set a preference set for a recipient.
	PreferenceSet param.Field[PreferenceSetRequestParam] `json:"preference_set"`
}

func (r TenantSetParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The branding for the tenant.
type TenantSetParamsSettingsBranding struct {
	// The icon URL for the tenant. Must point to a valid image with an image MIME
	// type.
	IconURL param.Field[string] `json:"icon_url"`
	// The logo URL for the tenant. Must point to a valid image with an image MIME
	// type.
	LogoURL param.Field[string] `json:"logo_url"`
	// The primary color for the tenant, provided as a hex value.
	PrimaryColor param.Field[string] `json:"primary_color"`
	// The primary color contrast for the tenant, provided as a hex value.
	PrimaryColorContrast param.Field[string] `json:"primary_color_contrast"`
}

func (r TenantSetParamsSettingsBranding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

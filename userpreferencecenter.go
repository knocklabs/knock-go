// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
)

// The preference center is a hosted page where users can manage their notification
// preferences.
//
// UserPreferenceCenterService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserPreferenceCenterService] method instead.
type UserPreferenceCenterService struct {
	Options []option.RequestOption
}

// NewUserPreferenceCenterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserPreferenceCenterService(opts ...option.RequestOption) (r *UserPreferenceCenterService) {
	r = &UserPreferenceCenterService{}
	r.Options = opts
	return
}

// Generates a signed preference center URL and token for the given user in the
// current environment.
func (r *UserPreferenceCenterService) GenerateSignedURL(ctx context.Context, userID string, opts ...option.RequestOption) (res *UserPreferenceCenterGenerateSignedURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/preference_center/signed_url", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the preference center config with environment metadata for the given
// user.
func (r *UserPreferenceCenterService) GetConfig(ctx context.Context, userID string, opts ...option.RequestOption) (res *UserPreferenceCenterGetConfigResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/preference_center/config", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The branding for the preference center, sourced from public environment
// variables.
type PreferenceCenterBrandingConfig struct {
	// The icon URL for the preference center. Must point to a valid image with an
	// image MIME type.
	IconURL string `json:"icon_url" api:"nullable" format:"uri"`
	// The logo URL for the preference center. Must point to a valid image with an
	// image MIME type.
	LogoURL string `json:"logo_url" api:"nullable" format:"uri"`
	// The primary color for the preference center, provided as a hex value.
	PrimaryColor string `json:"primary_color" api:"nullable"`
	// The primary color contrast for the preference center, provided as a hex value.
	PrimaryColorContrast string                             `json:"primary_color_contrast" api:"nullable"`
	JSON                 preferenceCenterBrandingConfigJSON `json:"-"`
}

// preferenceCenterBrandingConfigJSON contains the JSON metadata for the struct
// [PreferenceCenterBrandingConfig]
type preferenceCenterBrandingConfigJSON struct {
	IconURL              apijson.Field
	LogoURL              apijson.Field
	PrimaryColor         apijson.Field
	PrimaryColorContrast apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PreferenceCenterBrandingConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r preferenceCenterBrandingConfigJSON) RawJSON() string {
	return r.raw
}

// A signed preference center URL and token for a user.
type UserPreferenceCenterGenerateSignedURLResponse struct {
	// The signed JWT token for the preference center, usable as the `/p/{token}` path
	// segment.
	Token string `json:"token" api:"required"`
	// The full URL to the preference center for the user.
	URL  string                                            `json:"url" api:"required"`
	JSON userPreferenceCenterGenerateSignedURLResponseJSON `json:"-"`
}

// userPreferenceCenterGenerateSignedURLResponseJSON contains the JSON metadata for
// the struct [UserPreferenceCenterGenerateSignedURLResponse]
type userPreferenceCenterGenerateSignedURLResponseJSON struct {
	Token       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserPreferenceCenterGenerateSignedURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPreferenceCenterGenerateSignedURLResponseJSON) RawJSON() string {
	return r.raw
}

// The preference center configuration for an environment. Controls whether the
// preference center is enabled and defines the rows displayed in the UI.
type UserPreferenceCenterGetConfigResponse struct {
	// The name of the account that the preference center is associated with.
	AccountName string `json:"account_name" api:"required,nullable"`
	// The branding for the preference center, sourced from public environment
	// variables.
	Branding UserPreferenceCenterGetConfigResponseBranding `json:"branding" api:"required"`
	// The preference center configuration data containing the rows to display.
	Config UserPreferenceCenterGetConfigResponseConfig `json:"config" api:"required"`
	// Whether the preference center is enabled for this environment.
	Enabled bool `json:"enabled" api:"required"`
	// A display label for the user that the preference center is associated with,
	// resolved as email, then user id.
	UserEmail string `json:"user_email" api:"required,nullable"`
	// Whether Knock branding is required in the preference center.
	KnockBrandingRequired bool                                      `json:"knock_branding_required"`
	JSON                  userPreferenceCenterGetConfigResponseJSON `json:"-"`
}

// userPreferenceCenterGetConfigResponseJSON contains the JSON metadata for the
// struct [UserPreferenceCenterGetConfigResponse]
type userPreferenceCenterGetConfigResponseJSON struct {
	AccountName           apijson.Field
	Branding              apijson.Field
	Config                apijson.Field
	Enabled               apijson.Field
	UserEmail             apijson.Field
	KnockBrandingRequired apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *UserPreferenceCenterGetConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPreferenceCenterGetConfigResponseJSON) RawJSON() string {
	return r.raw
}

// The branding for the preference center, sourced from public environment
// variables.
type UserPreferenceCenterGetConfigResponseBranding struct {
	// The branding for the preference center, sourced from public environment
	// variables.
	Dark PreferenceCenterBrandingConfig                    `json:"dark"`
	JSON userPreferenceCenterGetConfigResponseBrandingJSON `json:"-"`
	PreferenceCenterBrandingConfig
}

// userPreferenceCenterGetConfigResponseBrandingJSON contains the JSON metadata for
// the struct [UserPreferenceCenterGetConfigResponseBranding]
type userPreferenceCenterGetConfigResponseBrandingJSON struct {
	Dark        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserPreferenceCenterGetConfigResponseBranding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPreferenceCenterGetConfigResponseBrandingJSON) RawJSON() string {
	return r.raw
}

// The preference center configuration data containing the rows to display.
type UserPreferenceCenterGetConfigResponseConfig struct {
	// The body text displayed below the title.
	Body string `json:"body" api:"required"`
	// An ordered list of rows to display in the preference center.
	Rows []UserPreferenceCenterGetConfigResponseConfigRow `json:"rows" api:"required"`
	// The title displayed at the top of the preference center.
	Title string `json:"title" api:"required"`
	// Whether the account name should be displayed in the preference center.
	ShowAccountName bool                                            `json:"show_account_name"`
	JSON            userPreferenceCenterGetConfigResponseConfigJSON `json:"-"`
}

// userPreferenceCenterGetConfigResponseConfigJSON contains the JSON metadata for
// the struct [UserPreferenceCenterGetConfigResponseConfig]
type userPreferenceCenterGetConfigResponseConfigJSON struct {
	Body            apijson.Field
	Rows            apijson.Field
	Title           apijson.Field
	ShowAccountName apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserPreferenceCenterGetConfigResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPreferenceCenterGetConfigResponseConfigJSON) RawJSON() string {
	return r.raw
}

// A preference row in the preference center configuration.
type UserPreferenceCenterGetConfigResponseConfigRow struct {
	// The display name of the preference row.
	Name string `json:"name" api:"required"`
	// The type of this preference row. `workflow` targets a workflow, `channel`
	// targets a specific channel, `category` targets a workflow category,
	// `channel_types` controls per-channel-type opt-in/out, and
	// `commercial_subscribed` is the commercial notification toggle.
	Type UserPreferenceCenterGetConfigResponseConfigRowsType `json:"type" api:"required"`
	// The list of channel types this preference is scoped to. An empty list (or
	// `null`) means the preference applies to all channel types. Present for
	// `workflow`, `category`, and `channel_types` types.
	ChannelTypes []UserPreferenceCenterGetConfigResponseConfigRowsChannelType `json:"channel_types" api:"nullable"`
	// A description shown below the preference row name.
	Description string `json:"description"`
	// The category name, workflow key, or channel ID this row controls (e.g.
	// `marketing`, `new-project-mentions`, or a channel UUID). Present for `workflow`,
	// `channel`, and `category` types.
	Identifier string                                             `json:"identifier" api:"nullable"`
	JSON       userPreferenceCenterGetConfigResponseConfigRowJSON `json:"-"`
}

// userPreferenceCenterGetConfigResponseConfigRowJSON contains the JSON metadata
// for the struct [UserPreferenceCenterGetConfigResponseConfigRow]
type userPreferenceCenterGetConfigResponseConfigRowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	ChannelTypes apijson.Field
	Description  apijson.Field
	Identifier   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserPreferenceCenterGetConfigResponseConfigRow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPreferenceCenterGetConfigResponseConfigRowJSON) RawJSON() string {
	return r.raw
}

// The type of this preference row. `workflow` targets a workflow, `channel`
// targets a specific channel, `category` targets a workflow category,
// `channel_types` controls per-channel-type opt-in/out, and
// `commercial_subscribed` is the commercial notification toggle.
type UserPreferenceCenterGetConfigResponseConfigRowsType string

const (
	UserPreferenceCenterGetConfigResponseConfigRowsTypeWorkflow             UserPreferenceCenterGetConfigResponseConfigRowsType = "workflow"
	UserPreferenceCenterGetConfigResponseConfigRowsTypeChannel              UserPreferenceCenterGetConfigResponseConfigRowsType = "channel"
	UserPreferenceCenterGetConfigResponseConfigRowsTypeCategory             UserPreferenceCenterGetConfigResponseConfigRowsType = "category"
	UserPreferenceCenterGetConfigResponseConfigRowsTypeChannelTypes         UserPreferenceCenterGetConfigResponseConfigRowsType = "channel_types"
	UserPreferenceCenterGetConfigResponseConfigRowsTypeCommercialSubscribed UserPreferenceCenterGetConfigResponseConfigRowsType = "commercial_subscribed"
)

func (r UserPreferenceCenterGetConfigResponseConfigRowsType) IsKnown() bool {
	switch r {
	case UserPreferenceCenterGetConfigResponseConfigRowsTypeWorkflow, UserPreferenceCenterGetConfigResponseConfigRowsTypeChannel, UserPreferenceCenterGetConfigResponseConfigRowsTypeCategory, UserPreferenceCenterGetConfigResponseConfigRowsTypeChannelTypes, UserPreferenceCenterGetConfigResponseConfigRowsTypeCommercialSubscribed:
		return true
	}
	return false
}

type UserPreferenceCenterGetConfigResponseConfigRowsChannelType string

const (
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeEmail       UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "email"
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeInApp       UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "in_app"
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeInAppFeed   UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "in_app_feed"
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeInAppGuide  UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "in_app_guide"
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeSMS         UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "sms"
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypePush        UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "push"
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeChat        UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "chat"
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeHTTP        UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "http"
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeLog         UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "log"
	UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeDeferredLog UserPreferenceCenterGetConfigResponseConfigRowsChannelType = "deferred_log"
)

func (r UserPreferenceCenterGetConfigResponseConfigRowsChannelType) IsKnown() bool {
	switch r {
	case UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeEmail, UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeInApp, UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeInAppFeed, UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeInAppGuide, UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeSMS, UserPreferenceCenterGetConfigResponseConfigRowsChannelTypePush, UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeChat, UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeHTTP, UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeLog, UserPreferenceCenterGetConfigResponseConfigRowsChannelTypeDeferredLog:
		return true
	}
	return false
}

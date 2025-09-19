// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"
	"slices"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
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

// Permanently deletes up to 1,000 users at a time.
func (r *UserBulkService) Delete(ctx context.Context, body UserBulkDeleteParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/bulk/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Identifies multiple users in a single operation. Allows creating or updating up
// to 1,000 users in a single batch with various properties, preferences, and
// channel data.
func (r *UserBulkService) Identify(ctx context.Context, body UserBulkIdentifyParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/bulk/identify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk sets the preferences for up to 1,000 users at a time. The preference
// set `:id` can be either `default` or a `tenant.id`. Learn more
// about [per-tenant preferences](/preferences/tenant-preferences). Note that this
// is a destructive operation and will replace any existing users' preferences with
// the preferences sent.
func (r *UserBulkService) SetPreferences(ctx context.Context, body UserBulkSetPreferencesParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/users/bulk/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserBulkDeleteParams struct {
	// A list of user IDs.
	UserIDs param.Field[[]string] `json:"user_ids,required"`
}

func (r UserBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkIdentifyParams struct {
	// A list of users.
	Users param.Field[[]InlineIdentifyUserRequestParam] `json:"users,required"`
}

func (r UserBulkIdentifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserBulkSetPreferencesParams struct {
	// A request to set a preference set for a recipient.
	Preferences param.Field[PreferenceSetRequestParam] `json:"preferences,required"`
	// A list of user IDs.
	UserIDs param.Field[[]string] `json:"user_ids,required"`
}

func (r UserBulkSetPreferencesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

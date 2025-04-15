// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"github.com/stainless-sdks/knock-go/option"
)

// UserFeedService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserFeedService] method instead.
type UserFeedService struct {
	Options []option.RequestOption
}

// NewUserFeedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserFeedService(opts ...option.RequestOption) (r *UserFeedService) {
	r = &UserFeedService{}
	r.Options = opts
	return
}

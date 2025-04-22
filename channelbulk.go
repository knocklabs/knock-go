// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"github.com/stainless-sdks/knock-go/option"
)

// ChannelBulkService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelBulkService] method instead.
type ChannelBulkService struct {
	Options []option.RequestOption
}

// NewChannelBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChannelBulkService(opts ...option.RequestOption) (r *ChannelBulkService) {
	r = &ChannelBulkService{}
	r.Options = opts
	return
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"github.com/stainless-sdks/knock-go/option"
)

// IntegrationService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationService] method instead.
type IntegrationService struct {
	Options   []option.RequestOption
	Census    *IntegrationCensusService
	Hightouch *IntegrationHightouchService
}

// NewIntegrationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntegrationService(opts ...option.RequestOption) (r *IntegrationService) {
	r = &IntegrationService{}
	r.Options = opts
	r.Census = NewIntegrationCensusService(opts...)
	r.Hightouch = NewIntegrationHightouchService(opts...)
	return
}

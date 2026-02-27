// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"github.com/knocklabs/knock-go/option"
)

// ProviderService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderService] method instead.
type ProviderService struct {
	Options []option.RequestOption
	// A provider represents a third-party service that Knock integrates with and is
	// configured via a channel.
	Slack *ProviderSlackService
	// A provider represents a third-party service that Knock integrates with and is
	// configured via a channel.
	MsTeams *ProviderMsTeamService
}

// NewProviderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderService(opts ...option.RequestOption) (r *ProviderService) {
	r = &ProviderService{}
	r.Options = opts
	r.Slack = NewProviderSlackService(opts...)
	r.MsTeams = NewProviderMsTeamService(opts...)
	return
}

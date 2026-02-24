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

// IntegrationHightouchService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationHightouchService] method instead.
type IntegrationHightouchService struct {
	Options []option.RequestOption
}

// NewIntegrationHightouchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntegrationHightouchService(opts ...option.RequestOption) (r *IntegrationHightouchService) {
	r = &IntegrationHightouchService{}
	r.Options = opts
	return
}

// Processes a Hightouch embedded destination RPC request.
func (r *IntegrationHightouchService) EmbeddedDestination(ctx context.Context, body IntegrationHightouchEmbeddedDestinationParams, opts ...option.RequestOption) (res *IntegrationHightouchEmbeddedDestinationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/integrations/hightouch/embedded-destination"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IntegrationHightouchEmbeddedDestinationResponse struct {
	// The request ID.
	ID string `json:"id"`
	// The result of the RPC call.
	Result map[string]interface{}                              `json:"result"`
	JSON   integrationHightouchEmbeddedDestinationResponseJSON `json:"-"`
}

// integrationHightouchEmbeddedDestinationResponseJSON contains the JSON metadata
// for the struct [IntegrationHightouchEmbeddedDestinationResponse]
type integrationHightouchEmbeddedDestinationResponseJSON struct {
	ID          apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationHightouchEmbeddedDestinationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationHightouchEmbeddedDestinationResponseJSON) RawJSON() string {
	return r.raw
}

type IntegrationHightouchEmbeddedDestinationParams struct {
	// The unique identifier for the RPC request.
	ID param.Field[string] `json:"id" api:"required"`
	// The JSON-RPC version.
	Jsonrpc param.Field[string] `json:"jsonrpc" api:"required"`
	// The method name to execute.
	Method param.Field[string] `json:"method" api:"required"`
	// The parameters for the method.
	Params param.Field[map[string]interface{}] `json:"params"`
}

func (r IntegrationHightouchEmbeddedDestinationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
)

// IntegrationCensusService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationCensusService] method instead.
type IntegrationCensusService struct {
	Options []option.RequestOption
}

// NewIntegrationCensusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntegrationCensusService(opts ...option.RequestOption) (r *IntegrationCensusService) {
	r = &IntegrationCensusService{}
	r.Options = opts
	return
}

// Processes a Census custom destination RPC request.
func (r *IntegrationCensusService) CustomDestination(ctx context.Context, body IntegrationCensusCustomDestinationParams, opts ...option.RequestOption) (res *IntegrationCensusCustomDestinationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/integrations/census/custom-destination"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IntegrationCensusCustomDestinationResponse struct {
	// The request ID.
	ID string `json:"id"`
	// The result of the RPC call.
	Result interface{}                                    `json:"result"`
	JSON   integrationCensusCustomDestinationResponseJSON `json:"-"`
}

// integrationCensusCustomDestinationResponseJSON contains the JSON metadata for
// the struct [IntegrationCensusCustomDestinationResponse]
type integrationCensusCustomDestinationResponseJSON struct {
	ID          apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationCensusCustomDestinationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationCensusCustomDestinationResponseJSON) RawJSON() string {
	return r.raw
}

type IntegrationCensusCustomDestinationParams struct {
	// The unique identifier for the RPC request.
	ID param.Field[string] `json:"id,required"`
	// The JSON-RPC version.
	Jsonrpc param.Field[string] `json:"jsonrpc,required"`
	// The method name to execute.
	Method param.Field[string] `json:"method,required"`
	// The parameters for the method.
	Params param.Field[interface{}] `json:"params"`
}

func (r IntegrationCensusCustomDestinationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

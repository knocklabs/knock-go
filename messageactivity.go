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
)

// MessageActivityService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageActivityService] method instead.
type MessageActivityService struct {
	Options []option.RequestOption
}

// NewMessageActivityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessageActivityService(opts ...option.RequestOption) (r *MessageActivityService) {
	r = &MessageActivityService{}
	r.Options = opts
	return
}

// Returns a paginated list of activities for the specified message.
func (r *MessageActivityService) List(ctx context.Context, messageID string, query MessageActivityListParams, opts ...option.RequestOption) (res *MessageActivityListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/activities", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a paginated list of `activities` associated with a given message. For
// messages produced after a [batch step](/designing-workflows/batch-function),
// this will contain one or more activities. Non-batched messages will always
// return a single activity.
type MessageActivityListResponse struct {
	// A list of activities.
	Entries []Activity `json:"entries,required"`
	// Pagination information for a list of resources.
	PageInfo PageInfo                        `json:"page_info,required"`
	JSON     messageActivityListResponseJSON `json:"-"`
}

// messageActivityListResponseJSON contains the JSON metadata for the struct
// [MessageActivityListResponse]
type messageActivityListResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageActivityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageActivityListResponseJSON) RawJSON() string {
	return r.raw
}

type MessageActivityListParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// The number of items per page.
	PageSize param.Field[int64] `query:"page_size"`
	// The trigger data to filter activities by.
	TriggerData param.Field[string] `query:"trigger_data"`
}

// URLQuery serializes [MessageActivityListParams]'s query parameters as
// `url.Values`.
func (r MessageActivityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

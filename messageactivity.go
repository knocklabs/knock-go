// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/knocklabs/knock-go/internal/apiquery"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/packages/pagination"
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
func (r *MessageActivityService) List(ctx context.Context, messageID string, query MessageActivityListParams, opts ...option.RequestOption) (res *pagination.ItemsCursor[Activity], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/activities", messageID)
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

// Returns a paginated list of activities for the specified message.
func (r *MessageActivityService) ListAutoPaging(ctx context.Context, messageID string, query MessageActivityListParams, opts ...option.RequestOption) *pagination.ItemsCursorAutoPager[Activity] {
	return pagination.NewItemsCursorAutoPager(r.List(ctx, messageID, query, opts...))
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

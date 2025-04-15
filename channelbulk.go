// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/knock-go/internal/requestconfig"
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

// Bulk update messages for a specific channel
func (r *ChannelBulkService) UpdateMessageStatus(ctx context.Context, channelID string, action ChannelBulkUpdateMessageStatusParamsAction, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/channels/%s/messages/bulk/%v", channelID, action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ChannelBulkUpdateMessageStatusParamsAction string

const (
	ChannelBulkUpdateMessageStatusParamsActionSeen       ChannelBulkUpdateMessageStatusParamsAction = "seen"
	ChannelBulkUpdateMessageStatusParamsActionUnseen     ChannelBulkUpdateMessageStatusParamsAction = "unseen"
	ChannelBulkUpdateMessageStatusParamsActionRead       ChannelBulkUpdateMessageStatusParamsAction = "read"
	ChannelBulkUpdateMessageStatusParamsActionUnread     ChannelBulkUpdateMessageStatusParamsAction = "unread"
	ChannelBulkUpdateMessageStatusParamsActionArchived   ChannelBulkUpdateMessageStatusParamsAction = "archived"
	ChannelBulkUpdateMessageStatusParamsActionUnarchived ChannelBulkUpdateMessageStatusParamsAction = "unarchived"
	ChannelBulkUpdateMessageStatusParamsActionInteracted ChannelBulkUpdateMessageStatusParamsAction = "interacted"
	ChannelBulkUpdateMessageStatusParamsActionArchive    ChannelBulkUpdateMessageStatusParamsAction = "archive"
	ChannelBulkUpdateMessageStatusParamsActionUnarchive  ChannelBulkUpdateMessageStatusParamsAction = "unarchive"
	ChannelBulkUpdateMessageStatusParamsActionDelete     ChannelBulkUpdateMessageStatusParamsAction = "delete"
)

func (r ChannelBulkUpdateMessageStatusParamsAction) IsKnown() bool {
	switch r {
	case ChannelBulkUpdateMessageStatusParamsActionSeen, ChannelBulkUpdateMessageStatusParamsActionUnseen, ChannelBulkUpdateMessageStatusParamsActionRead, ChannelBulkUpdateMessageStatusParamsActionUnread, ChannelBulkUpdateMessageStatusParamsActionArchived, ChannelBulkUpdateMessageStatusParamsActionUnarchived, ChannelBulkUpdateMessageStatusParamsActionInteracted, ChannelBulkUpdateMessageStatusParamsActionArchive, ChannelBulkUpdateMessageStatusParamsActionUnarchive, ChannelBulkUpdateMessageStatusParamsActionDelete:
		return true
	}
	return false
}

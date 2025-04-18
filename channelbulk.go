// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
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

// Bulk update messages for a specific channel. The channel is specified by the
// `channel_id` parameter. The action to perform is specified by the `action`
// parameter, where the action is a status change action (e.g. `archive`,
// `unarchive`).
func (r *ChannelBulkService) UpdateMessageStatus(ctx context.Context, channelID string, action ChannelBulkUpdateMessageStatusParamsAction, body ChannelBulkUpdateMessageStatusParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/channels/%s/messages/bulk/%v", channelID, action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChannelBulkUpdateMessageStatusParams struct {
	// The archived status to filter messages by.
	Archived param.Field[ChannelBulkUpdateMessageStatusParamsArchived] `json:"archived"`
	// The delivery status to filter messages by.
	DeliveryStatus param.Field[ChannelBulkUpdateMessageStatusParamsDeliveryStatus] `json:"delivery_status"`
	// The engagement status to filter messages by.
	EngagementStatus param.Field[ChannelBulkUpdateMessageStatusParamsEngagementStatus] `json:"engagement_status"`
	// Whether to include only messages that have a tenant or not.
	HasTenant param.Field[bool] `json:"has_tenant"`
	// The timestamp to filter messages by. Only include messages created after this
	// timestamp.
	NewerThan param.Field[time.Time] `json:"newer_than" format:"date-time"`
	// The timestamp to filter messages by. Only include messages created before this
	// timestamp.
	OlderThan param.Field[time.Time] `json:"older_than" format:"date-time"`
	// The recipient GIDs to filter messages by.
	RecipientGids param.Field[[]string] `json:"recipient_gids"`
	// The recipient IDs to filter messages by.
	RecipientIDs param.Field[[]string] `json:"recipient_ids"`
	// The tenant IDs to filter messages by.
	Tenants param.Field[[]string] `json:"tenants"`
	// The trigger data to filter messages by. Must be a valid JSON object.
	TriggerData param.Field[string] `json:"trigger_data"`
	// The workflow keys to filter messages by.
	Workflows param.Field[[]string] `json:"workflows"`
}

func (r ChannelBulkUpdateMessageStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// The archived status to filter messages by.
type ChannelBulkUpdateMessageStatusParamsArchived string

const (
	ChannelBulkUpdateMessageStatusParamsArchivedExclude ChannelBulkUpdateMessageStatusParamsArchived = "exclude"
	ChannelBulkUpdateMessageStatusParamsArchivedInclude ChannelBulkUpdateMessageStatusParamsArchived = "include"
	ChannelBulkUpdateMessageStatusParamsArchivedOnly    ChannelBulkUpdateMessageStatusParamsArchived = "only"
)

func (r ChannelBulkUpdateMessageStatusParamsArchived) IsKnown() bool {
	switch r {
	case ChannelBulkUpdateMessageStatusParamsArchivedExclude, ChannelBulkUpdateMessageStatusParamsArchivedInclude, ChannelBulkUpdateMessageStatusParamsArchivedOnly:
		return true
	}
	return false
}

// The delivery status to filter messages by.
type ChannelBulkUpdateMessageStatusParamsDeliveryStatus string

const (
	ChannelBulkUpdateMessageStatusParamsDeliveryStatusQueued            ChannelBulkUpdateMessageStatusParamsDeliveryStatus = "queued"
	ChannelBulkUpdateMessageStatusParamsDeliveryStatusSent              ChannelBulkUpdateMessageStatusParamsDeliveryStatus = "sent"
	ChannelBulkUpdateMessageStatusParamsDeliveryStatusDelivered         ChannelBulkUpdateMessageStatusParamsDeliveryStatus = "delivered"
	ChannelBulkUpdateMessageStatusParamsDeliveryStatusDeliveryAttempted ChannelBulkUpdateMessageStatusParamsDeliveryStatus = "delivery_attempted"
	ChannelBulkUpdateMessageStatusParamsDeliveryStatusUndelivered       ChannelBulkUpdateMessageStatusParamsDeliveryStatus = "undelivered"
	ChannelBulkUpdateMessageStatusParamsDeliveryStatusNotSent           ChannelBulkUpdateMessageStatusParamsDeliveryStatus = "not_sent"
	ChannelBulkUpdateMessageStatusParamsDeliveryStatusBounced           ChannelBulkUpdateMessageStatusParamsDeliveryStatus = "bounced"
)

func (r ChannelBulkUpdateMessageStatusParamsDeliveryStatus) IsKnown() bool {
	switch r {
	case ChannelBulkUpdateMessageStatusParamsDeliveryStatusQueued, ChannelBulkUpdateMessageStatusParamsDeliveryStatusSent, ChannelBulkUpdateMessageStatusParamsDeliveryStatusDelivered, ChannelBulkUpdateMessageStatusParamsDeliveryStatusDeliveryAttempted, ChannelBulkUpdateMessageStatusParamsDeliveryStatusUndelivered, ChannelBulkUpdateMessageStatusParamsDeliveryStatusNotSent, ChannelBulkUpdateMessageStatusParamsDeliveryStatusBounced:
		return true
	}
	return false
}

// The engagement status to filter messages by.
type ChannelBulkUpdateMessageStatusParamsEngagementStatus string

const (
	ChannelBulkUpdateMessageStatusParamsEngagementStatusSeen        ChannelBulkUpdateMessageStatusParamsEngagementStatus = "seen"
	ChannelBulkUpdateMessageStatusParamsEngagementStatusUnseen      ChannelBulkUpdateMessageStatusParamsEngagementStatus = "unseen"
	ChannelBulkUpdateMessageStatusParamsEngagementStatusRead        ChannelBulkUpdateMessageStatusParamsEngagementStatus = "read"
	ChannelBulkUpdateMessageStatusParamsEngagementStatusUnread      ChannelBulkUpdateMessageStatusParamsEngagementStatus = "unread"
	ChannelBulkUpdateMessageStatusParamsEngagementStatusArchived    ChannelBulkUpdateMessageStatusParamsEngagementStatus = "archived"
	ChannelBulkUpdateMessageStatusParamsEngagementStatusUnarchived  ChannelBulkUpdateMessageStatusParamsEngagementStatus = "unarchived"
	ChannelBulkUpdateMessageStatusParamsEngagementStatusLinkClicked ChannelBulkUpdateMessageStatusParamsEngagementStatus = "link_clicked"
	ChannelBulkUpdateMessageStatusParamsEngagementStatusInteracted  ChannelBulkUpdateMessageStatusParamsEngagementStatus = "interacted"
)

func (r ChannelBulkUpdateMessageStatusParamsEngagementStatus) IsKnown() bool {
	switch r {
	case ChannelBulkUpdateMessageStatusParamsEngagementStatusSeen, ChannelBulkUpdateMessageStatusParamsEngagementStatusUnseen, ChannelBulkUpdateMessageStatusParamsEngagementStatusRead, ChannelBulkUpdateMessageStatusParamsEngagementStatusUnread, ChannelBulkUpdateMessageStatusParamsEngagementStatusArchived, ChannelBulkUpdateMessageStatusParamsEngagementStatusUnarchived, ChannelBulkUpdateMessageStatusParamsEngagementStatusLinkClicked, ChannelBulkUpdateMessageStatusParamsEngagementStatusInteracted:
		return true
	}
	return false
}

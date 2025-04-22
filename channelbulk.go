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

// Bulk update the status of messages for a specific channel. The channel is
// specified by the `channel_id` parameter. The action to perform is specified by
// the `action` parameter, where the action is a status change action (e.g.
// `archive`, `unarchive`).
func (r *ChannelBulkService) UpdateMessageStatus(ctx context.Context, channelID string, status ChannelBulkUpdateMessageStatusParamsStatus, body ChannelBulkUpdateMessageStatusParams, opts ...option.RequestOption) (res *BulkOperation, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/channels/%s/messages/bulk/%v", channelID, status)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChannelBulkUpdateMessageStatusParams struct {
	// Limits the results to messages with the given archived status.
	Archived param.Field[ChannelBulkUpdateMessageStatusParamsArchived] `json:"archived"`
	// Limits the results to messages with the given delivery status.
	DeliveryStatus param.Field[ChannelBulkUpdateMessageStatusParamsDeliveryStatus] `json:"delivery_status"`
	// Limits the results to messages with the given engagement status.
	EngagementStatus param.Field[ChannelBulkUpdateMessageStatusParamsEngagementStatus] `json:"engagement_status"`
	// Limits the results to messages that have a tenant or not.
	HasTenant param.Field[bool] `json:"has_tenant"`
	// Limits the results to messages inserted after the given date.
	NewerThan param.Field[time.Time] `json:"newer_than" format:"date-time"`
	// Limits the results to messages inserted before the given date.
	OlderThan param.Field[time.Time] `json:"older_than" format:"date-time"`
	// Limits the results to messages with the given recipient IDs.
	RecipientIDs param.Field[[]string] `json:"recipient_ids"`
	// Limits the results to messages with the given tenant IDs.
	Tenants param.Field[[]string] `json:"tenants"`
	// Limits the results to only messages that were generated with the given data. See
	// [trigger data filtering](/api-reference/overview/trigger-data-filtering) for
	// more information.
	TriggerData param.Field[string] `json:"trigger_data"`
	// Limits the results to messages with the given workflow keys.
	Workflows param.Field[[]string] `json:"workflows"`
}

func (r ChannelBulkUpdateMessageStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target status to be applied to the messages.
type ChannelBulkUpdateMessageStatusParamsStatus string

const (
	ChannelBulkUpdateMessageStatusParamsStatusSeen       ChannelBulkUpdateMessageStatusParamsStatus = "seen"
	ChannelBulkUpdateMessageStatusParamsStatusUnseen     ChannelBulkUpdateMessageStatusParamsStatus = "unseen"
	ChannelBulkUpdateMessageStatusParamsStatusRead       ChannelBulkUpdateMessageStatusParamsStatus = "read"
	ChannelBulkUpdateMessageStatusParamsStatusUnread     ChannelBulkUpdateMessageStatusParamsStatus = "unread"
	ChannelBulkUpdateMessageStatusParamsStatusArchived   ChannelBulkUpdateMessageStatusParamsStatus = "archived"
	ChannelBulkUpdateMessageStatusParamsStatusUnarchived ChannelBulkUpdateMessageStatusParamsStatus = "unarchived"
	ChannelBulkUpdateMessageStatusParamsStatusInteracted ChannelBulkUpdateMessageStatusParamsStatus = "interacted"
	ChannelBulkUpdateMessageStatusParamsStatusArchive    ChannelBulkUpdateMessageStatusParamsStatus = "archive"
	ChannelBulkUpdateMessageStatusParamsStatusUnarchive  ChannelBulkUpdateMessageStatusParamsStatus = "unarchive"
	ChannelBulkUpdateMessageStatusParamsStatusDelete     ChannelBulkUpdateMessageStatusParamsStatus = "delete"
)

func (r ChannelBulkUpdateMessageStatusParamsStatus) IsKnown() bool {
	switch r {
	case ChannelBulkUpdateMessageStatusParamsStatusSeen, ChannelBulkUpdateMessageStatusParamsStatusUnseen, ChannelBulkUpdateMessageStatusParamsStatusRead, ChannelBulkUpdateMessageStatusParamsStatusUnread, ChannelBulkUpdateMessageStatusParamsStatusArchived, ChannelBulkUpdateMessageStatusParamsStatusUnarchived, ChannelBulkUpdateMessageStatusParamsStatusInteracted, ChannelBulkUpdateMessageStatusParamsStatusArchive, ChannelBulkUpdateMessageStatusParamsStatusUnarchive, ChannelBulkUpdateMessageStatusParamsStatusDelete:
		return true
	}
	return false
}

// Limits the results to messages with the given archived status.
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

// Limits the results to messages with the given delivery status.
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

// Limits the results to messages with the given engagement status.
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

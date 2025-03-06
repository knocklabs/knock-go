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

// Bulk update messages for a specific channel
func (r *ChannelBulkService) UpdateMessageStatus(ctx context.Context, channelID string, action ChannelBulkUpdateMessageStatusParamsAction, body ChannelBulkUpdateMessageStatusParams, opts ...option.RequestOption) (res *ChannelBulkUpdateMessageStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/channels/%s/messages/bulk/%v", channelID, action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A bulk operation entity
type ChannelBulkUpdateMessageStatusResponse struct {
	ID                 string                                       `json:"id,required" format:"uuid"`
	Typename           string                                       `json:"__typename,required"`
	EstimatedTotalRows int64                                        `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                                    `json:"inserted_at,required" format:"date-time"`
	Name               string                                       `json:"name,required"`
	ProcessedRows      int64                                        `json:"processed_rows,required"`
	Status             ChannelBulkUpdateMessageStatusResponseStatus `json:"status,required"`
	SuccessCount       int64                                        `json:"success_count,required"`
	UpdatedAt          time.Time                                    `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                                    `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                                        `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []ChannelBulkUpdateMessageStatusResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                                         `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                                         `json:"started_at,nullable" format:"date-time"`
	JSON       channelBulkUpdateMessageStatusResponseJSON        `json:"-"`
}

// channelBulkUpdateMessageStatusResponseJSON contains the JSON metadata for the
// struct [ChannelBulkUpdateMessageStatusResponse]
type channelBulkUpdateMessageStatusResponseJSON struct {
	ID                 apijson.Field
	Typename           apijson.Field
	EstimatedTotalRows apijson.Field
	InsertedAt         apijson.Field
	Name               apijson.Field
	ProcessedRows      apijson.Field
	Status             apijson.Field
	SuccessCount       apijson.Field
	UpdatedAt          apijson.Field
	CompletedAt        apijson.Field
	ErrorCount         apijson.Field
	ErrorItems         apijson.Field
	FailedAt           apijson.Field
	StartedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ChannelBulkUpdateMessageStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r channelBulkUpdateMessageStatusResponseJSON) RawJSON() string {
	return r.raw
}

type ChannelBulkUpdateMessageStatusResponseStatus string

const (
	ChannelBulkUpdateMessageStatusResponseStatusQueued     ChannelBulkUpdateMessageStatusResponseStatus = "queued"
	ChannelBulkUpdateMessageStatusResponseStatusProcessing ChannelBulkUpdateMessageStatusResponseStatus = "processing"
	ChannelBulkUpdateMessageStatusResponseStatusCompleted  ChannelBulkUpdateMessageStatusResponseStatus = "completed"
	ChannelBulkUpdateMessageStatusResponseStatusFailed     ChannelBulkUpdateMessageStatusResponseStatus = "failed"
)

func (r ChannelBulkUpdateMessageStatusResponseStatus) IsKnown() bool {
	switch r {
	case ChannelBulkUpdateMessageStatusResponseStatusQueued, ChannelBulkUpdateMessageStatusResponseStatusProcessing, ChannelBulkUpdateMessageStatusResponseStatusCompleted, ChannelBulkUpdateMessageStatusResponseStatusFailed:
		return true
	}
	return false
}

type ChannelBulkUpdateMessageStatusResponseErrorItem struct {
	ID         string                                              `json:"id,required"`
	Collection string                                              `json:"collection,nullable"`
	JSON       channelBulkUpdateMessageStatusResponseErrorItemJSON `json:"-"`
}

// channelBulkUpdateMessageStatusResponseErrorItemJSON contains the JSON metadata
// for the struct [ChannelBulkUpdateMessageStatusResponseErrorItem]
type channelBulkUpdateMessageStatusResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChannelBulkUpdateMessageStatusResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r channelBulkUpdateMessageStatusResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

type ChannelBulkUpdateMessageStatusParams struct {
	Archived         param.Field[ChannelBulkUpdateMessageStatusParamsArchived]         `json:"archived"`
	DeliveryStatus   param.Field[ChannelBulkUpdateMessageStatusParamsDeliveryStatus]   `json:"delivery_status"`
	EngagementStatus param.Field[ChannelBulkUpdateMessageStatusParamsEngagementStatus] `json:"engagement_status"`
	HasTenant        param.Field[bool]                                                 `json:"has_tenant"`
	NewerThan        param.Field[time.Time]                                            `json:"newer_than" format:"date-time"`
	OlderThan        param.Field[time.Time]                                            `json:"older_than" format:"date-time"`
	RecipientIDs     param.Field[[]string]                                             `json:"recipient_ids"`
	Tenants          param.Field[[]string]                                             `json:"tenants"`
	TriggerData      param.Field[string]                                               `json:"trigger_data"`
	Workflows        param.Field[[]string]                                             `json:"workflows"`
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

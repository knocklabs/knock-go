// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
)

// ObjectBulkService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectBulkService] method instead.
type ObjectBulkService struct {
	Options []option.RequestOption
}

// NewObjectBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectBulkService(opts ...option.RequestOption) (r *ObjectBulkService) {
	r = &ObjectBulkService{}
	r.Options = opts
	return
}

// Bulk delete objects
func (r *ObjectBulkService) Delete(ctx context.Context, collection string, body ObjectBulkDeleteParams, opts ...option.RequestOption) (res *ObjectBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/delete", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add subscriptions for a set of objects in a single collection. If a subscription
// already exists, it will be updated.
func (r *ObjectBulkService) AddSubscriptions(ctx context.Context, collection string, body ObjectBulkAddSubscriptionsParams, opts ...option.RequestOption) (res *ObjectBulkAddSubscriptionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/subscriptions/add", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Bulk set objects
func (r *ObjectBulkService) Set(ctx context.Context, collection string, body ObjectBulkSetParams, opts ...option.RequestOption) (res *ObjectBulkSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if collection == "" {
		err = errors.New("missing required collection parameter")
		return
	}
	path := fmt.Sprintf("v1/objects/%s/bulk/set", collection)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A bulk operation entity
type ObjectBulkDeleteResponse struct {
	ID                 string                         `json:"id,required" format:"uuid"`
	Typename           string                         `json:"__typename,required"`
	EstimatedTotalRows int64                          `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                      `json:"inserted_at,required" format:"date-time"`
	Name               string                         `json:"name,required"`
	ProcessedRows      int64                          `json:"processed_rows,required"`
	Status             ObjectBulkDeleteResponseStatus `json:"status,required"`
	SuccessCount       int64                          `json:"success_count,required"`
	UpdatedAt          time.Time                      `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                      `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                          `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []ObjectBulkDeleteResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                           `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                           `json:"started_at,nullable" format:"date-time"`
	JSON       objectBulkDeleteResponseJSON        `json:"-"`
}

// objectBulkDeleteResponseJSON contains the JSON metadata for the struct
// [ObjectBulkDeleteResponse]
type objectBulkDeleteResponseJSON struct {
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

func (r *ObjectBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ObjectBulkDeleteResponseStatus string

const (
	ObjectBulkDeleteResponseStatusQueued     ObjectBulkDeleteResponseStatus = "queued"
	ObjectBulkDeleteResponseStatusProcessing ObjectBulkDeleteResponseStatus = "processing"
	ObjectBulkDeleteResponseStatusCompleted  ObjectBulkDeleteResponseStatus = "completed"
	ObjectBulkDeleteResponseStatusFailed     ObjectBulkDeleteResponseStatus = "failed"
)

func (r ObjectBulkDeleteResponseStatus) IsKnown() bool {
	switch r {
	case ObjectBulkDeleteResponseStatusQueued, ObjectBulkDeleteResponseStatusProcessing, ObjectBulkDeleteResponseStatusCompleted, ObjectBulkDeleteResponseStatusFailed:
		return true
	}
	return false
}

type ObjectBulkDeleteResponseErrorItem struct {
	ID         string                                `json:"id,required"`
	Collection string                                `json:"collection,nullable"`
	JSON       objectBulkDeleteResponseErrorItemJSON `json:"-"`
}

// objectBulkDeleteResponseErrorItemJSON contains the JSON metadata for the struct
// [ObjectBulkDeleteResponseErrorItem]
type objectBulkDeleteResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectBulkDeleteResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkDeleteResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type ObjectBulkAddSubscriptionsResponse struct {
	ID                 string                                   `json:"id,required" format:"uuid"`
	Typename           string                                   `json:"__typename,required"`
	EstimatedTotalRows int64                                    `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                                `json:"inserted_at,required" format:"date-time"`
	Name               string                                   `json:"name,required"`
	ProcessedRows      int64                                    `json:"processed_rows,required"`
	Status             ObjectBulkAddSubscriptionsResponseStatus `json:"status,required"`
	SuccessCount       int64                                    `json:"success_count,required"`
	UpdatedAt          time.Time                                `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                                `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                                    `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []ObjectBulkAddSubscriptionsResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                                     `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                                     `json:"started_at,nullable" format:"date-time"`
	JSON       objectBulkAddSubscriptionsResponseJSON        `json:"-"`
}

// objectBulkAddSubscriptionsResponseJSON contains the JSON metadata for the struct
// [ObjectBulkAddSubscriptionsResponse]
type objectBulkAddSubscriptionsResponseJSON struct {
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

func (r *ObjectBulkAddSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkAddSubscriptionsResponseJSON) RawJSON() string {
	return r.raw
}

type ObjectBulkAddSubscriptionsResponseStatus string

const (
	ObjectBulkAddSubscriptionsResponseStatusQueued     ObjectBulkAddSubscriptionsResponseStatus = "queued"
	ObjectBulkAddSubscriptionsResponseStatusProcessing ObjectBulkAddSubscriptionsResponseStatus = "processing"
	ObjectBulkAddSubscriptionsResponseStatusCompleted  ObjectBulkAddSubscriptionsResponseStatus = "completed"
	ObjectBulkAddSubscriptionsResponseStatusFailed     ObjectBulkAddSubscriptionsResponseStatus = "failed"
)

func (r ObjectBulkAddSubscriptionsResponseStatus) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsResponseStatusQueued, ObjectBulkAddSubscriptionsResponseStatusProcessing, ObjectBulkAddSubscriptionsResponseStatusCompleted, ObjectBulkAddSubscriptionsResponseStatusFailed:
		return true
	}
	return false
}

type ObjectBulkAddSubscriptionsResponseErrorItem struct {
	ID         string                                          `json:"id,required"`
	Collection string                                          `json:"collection,nullable"`
	JSON       objectBulkAddSubscriptionsResponseErrorItemJSON `json:"-"`
}

// objectBulkAddSubscriptionsResponseErrorItemJSON contains the JSON metadata for
// the struct [ObjectBulkAddSubscriptionsResponseErrorItem]
type objectBulkAddSubscriptionsResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectBulkAddSubscriptionsResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkAddSubscriptionsResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

// A bulk operation entity
type ObjectBulkSetResponse struct {
	ID                 string                      `json:"id,required" format:"uuid"`
	Typename           string                      `json:"__typename,required"`
	EstimatedTotalRows int64                       `json:"estimated_total_rows,required"`
	InsertedAt         time.Time                   `json:"inserted_at,required" format:"date-time"`
	Name               string                      `json:"name,required"`
	ProcessedRows      int64                       `json:"processed_rows,required"`
	Status             ObjectBulkSetResponseStatus `json:"status,required"`
	SuccessCount       int64                       `json:"success_count,required"`
	UpdatedAt          time.Time                   `json:"updated_at,required" format:"date-time"`
	CompletedAt        time.Time                   `json:"completed_at,nullable" format:"date-time"`
	ErrorCount         int64                       `json:"error_count"`
	// A list of items that failed to be processed
	ErrorItems []ObjectBulkSetResponseErrorItem `json:"error_items"`
	FailedAt   time.Time                        `json:"failed_at,nullable" format:"date-time"`
	StartedAt  time.Time                        `json:"started_at,nullable" format:"date-time"`
	JSON       objectBulkSetResponseJSON        `json:"-"`
}

// objectBulkSetResponseJSON contains the JSON metadata for the struct
// [ObjectBulkSetResponse]
type objectBulkSetResponseJSON struct {
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

func (r *ObjectBulkSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkSetResponseJSON) RawJSON() string {
	return r.raw
}

type ObjectBulkSetResponseStatus string

const (
	ObjectBulkSetResponseStatusQueued     ObjectBulkSetResponseStatus = "queued"
	ObjectBulkSetResponseStatusProcessing ObjectBulkSetResponseStatus = "processing"
	ObjectBulkSetResponseStatusCompleted  ObjectBulkSetResponseStatus = "completed"
	ObjectBulkSetResponseStatusFailed     ObjectBulkSetResponseStatus = "failed"
)

func (r ObjectBulkSetResponseStatus) IsKnown() bool {
	switch r {
	case ObjectBulkSetResponseStatusQueued, ObjectBulkSetResponseStatusProcessing, ObjectBulkSetResponseStatusCompleted, ObjectBulkSetResponseStatusFailed:
		return true
	}
	return false
}

type ObjectBulkSetResponseErrorItem struct {
	ID         string                             `json:"id,required"`
	Collection string                             `json:"collection,nullable"`
	JSON       objectBulkSetResponseErrorItemJSON `json:"-"`
}

// objectBulkSetResponseErrorItemJSON contains the JSON metadata for the struct
// [ObjectBulkSetResponseErrorItem]
type objectBulkSetResponseErrorItemJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectBulkSetResponseErrorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectBulkSetResponseErrorItemJSON) RawJSON() string {
	return r.raw
}

type ObjectBulkDeleteParams struct {
	// The IDs of the objects to delete
	ObjectIDs param.Field[[]string] `query:"object_ids,required"`
}

// URLQuery serializes [ObjectBulkDeleteParams]'s query parameters as `url.Values`.
func (r ObjectBulkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectBulkAddSubscriptionsParams struct {
	Subscriptions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscription] `json:"subscriptions,required"`
}

func (r ObjectBulkAddSubscriptionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscription struct {
	ID         param.Field[string]                                                        `json:"id,required"`
	Recipients param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion] `json:"recipients,required"`
	Properties param.Field[map[string]interface{}]                                        `json:"properties"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID          param.Field[string]      `json:"id,required"`
	ChannelData param.Field[interface{}] `json:"channel_data"`
	Collection  param.Field[string]      `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt   param.Field[time.Time]   `json:"created_at" format:"date-time"`
	Preferences param.Field[interface{}] `json:"preferences"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion() {
}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// Satisfied by [shared.UnionString],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequest],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequest],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipient].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion()
}

// A set of parameters to inline-identify a user with. Inline identifying the user
// will ensure that the user is available before the request is executed in Knock.
// It will perform an upsert against the user you're supplying, replacing any
// properties specified.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequest struct {
	// The ID of the user to identify. This is an ID that you supply.
	ID param.Field[string] `json:"id,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelData] `json:"channel_data"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferences] `json:"preferences"`
	ExtraFields map[string]interface{}                                                                                              `json:"-,extras"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequest) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion() {
}

// Set channel data for a type of channel
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelData struct {
	// Channel data for push providers
	Data param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion] `json:"data,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataData struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion() {
}

// Channel data for push providers
//
// Satisfied by
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataPushChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataOneSignalChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataData].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion interface {
	implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion()
}

// Channel data for push providers
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataPushChannelData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataPushChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataPushChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion() {
}

// OneSignal channel data
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataOneSignalChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataOneSignalChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion() {
}

// Slack channel data
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelData struct {
	Connections param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionUnion] `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataToken] `json:"token"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnection].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataConnectionUnion() {
}

// A token that's used to store the access token for a Slack workspace.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataToken struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataSlackChannelDataToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelData struct {
	Connections param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionUnion] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion() {
}

// Microsoft Teams token connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnection struct {
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams token connection
//
// Satisfied by
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnection].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// Microsoft Teams token connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams incoming webhook connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// The incoming webhook
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Discord channel data
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelData struct {
	Connections param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionUnion] `json:"connections,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataUnion() {
}

// Discord channel connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnection].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord incoming webhook connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionUnion() {
}

// The incoming webhook
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferences struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesUnion() {
}

// Channel type preferences
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                 `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                  `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                 `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                      `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                      `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                 `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion] `json:"sms"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                             `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                              `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                              `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                             `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                  `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                             `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsUnion() {
}

// Channel type preferences
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                 `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                     `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                     `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                               `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                               `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Inline identifies a custom object belonging to a collection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequest struct {
	ID         param.Field[string] `json:"id,required"`
	Collection param.Field[string] `json:"collection,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelData] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                                                                                                `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferences] `json:"preferences"`
	ExtraFields map[string]interface{}                                                                                                `json:"-,extras"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequest) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion() {
}

// Set channel data for a type of channel
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelData struct {
	// Channel data for push providers
	Data param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataUnion] `json:"data,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataData struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataUnion() {
}

// Channel data for push providers
//
// Satisfied by
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataPushChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataOneSignalChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelData],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataData].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataUnion interface {
	implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataUnion()
}

// Channel data for push providers
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataPushChannelData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataPushChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataPushChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataUnion() {
}

// OneSignal channel data
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataOneSignalChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataOneSignalChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataUnion() {
}

// Slack channel data
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelData struct {
	Connections param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionUnion] `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataToken] `json:"token"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnection].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataConnectionUnion() {
}

// A token that's used to store the access token for a Slack workspace.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataToken struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataSlackChannelDataToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelData struct {
	Connections param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionUnion] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataUnion() {
}

// Microsoft Teams token connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnection struct {
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams token connection
//
// Satisfied by
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnection].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// Microsoft Teams token connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams incoming webhook connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// The incoming webhook
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Discord channel data
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelData struct {
	Connections param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionUnion] `json:"connections,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelData) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataUnion() {
}

// Discord channel connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnection].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord incoming webhook connection
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) implementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionUnion() {
}

// The incoming webhook
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferences struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesUnion() {
}

// Channel type preferences
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                   `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                   `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                    `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                    `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                   `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                   `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                        `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                        `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                   `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                   `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                  `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                              `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                              `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion] `json:"sms"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                               `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                               `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                               `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                               `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                    `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                    `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                               `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                               `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                              `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                              `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsUnion() {
}

// Channel type preferences
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                  `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                   `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                   `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                  `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                       `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                       `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                  `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                  `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                                                                 `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                                                                 `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                             `json:"argument,required"`
	Operator param.Field[ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                             `json:"variable,required"`
}

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyObjectRequestPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

type ObjectBulkSetParams struct {
	Objects param.Field[[]ObjectBulkSetParamsObject] `json:"objects,required"`
}

func (r ObjectBulkSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Inline identifies a custom object belonging to a collection
type ObjectBulkSetParamsObject struct {
	ID         param.Field[string] `json:"id,required"`
	Collection param.Field[string] `json:"collection,required"`
	// Allows inline setting channel data for a recipient
	ChannelData param.Field[map[string]ObjectBulkSetParamsObjectsChannelData] `json:"channel_data"`
	CreatedAt   param.Field[time.Time]                                        `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient, where the key is the preference set name
	Preferences param.Field[map[string]ObjectBulkSetParamsObjectsPreferences] `json:"preferences"`
	ExtraFields map[string]interface{}                                        `json:"-,extras"`
}

func (r ObjectBulkSetParamsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set channel data for a type of channel
type ObjectBulkSetParamsObjectsChannelData struct {
	// Channel data for push providers
	Data param.Field[ObjectBulkSetParamsObjectsChannelDataDataUnion] `json:"data,required"`
}

func (r ObjectBulkSetParamsObjectsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Channel data for push providers
type ObjectBulkSetParamsObjectsChannelDataData struct {
	Token       param.Field[interface{}] `json:"token"`
	Connections param.Field[interface{}] `json:"connections"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string]      `json:"ms_teams_tenant_id" format:"uuid"`
	PlayerIDs       param.Field[interface{}] `json:"player_ids"`
	Tokens          param.Field[interface{}] `json:"tokens"`
}

func (r ObjectBulkSetParamsObjectsChannelDataData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataData) implementsObjectBulkSetParamsObjectsChannelDataDataUnion() {
}

// Channel data for push providers
//
// Satisfied by [ObjectBulkSetParamsObjectsChannelDataDataPushChannelData],
// [ObjectBulkSetParamsObjectsChannelDataDataOneSignalChannelData],
// [ObjectBulkSetParamsObjectsChannelDataDataSlackChannelData],
// [ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelData],
// [ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelData],
// [ObjectBulkSetParamsObjectsChannelDataData].
type ObjectBulkSetParamsObjectsChannelDataDataUnion interface {
	implementsObjectBulkSetParamsObjectsChannelDataDataUnion()
}

// Channel data for push providers
type ObjectBulkSetParamsObjectsChannelDataDataPushChannelData struct {
	Tokens param.Field[[]string] `json:"tokens,required"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataPushChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataPushChannelData) implementsObjectBulkSetParamsObjectsChannelDataDataUnion() {
}

// OneSignal channel data
type ObjectBulkSetParamsObjectsChannelDataDataOneSignalChannelData struct {
	// The OneSignal player IDs
	PlayerIDs param.Field[[]string] `json:"player_ids,required" format:"uuid"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataOneSignalChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataOneSignalChannelData) implementsObjectBulkSetParamsObjectsChannelDataDataUnion() {
}

// Slack channel data
type ObjectBulkSetParamsObjectsChannelDataDataSlackChannelData struct {
	Connections param.Field[[]ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionUnion] `json:"connections,required"`
	// A token that's used to store the access token for a Slack workspace.
	Token param.Field[ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataToken] `json:"token"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataSlackChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataSlackChannelData) implementsObjectBulkSetParamsObjectsChannelDataDataUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
type ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	URL         param.Field[string] `json:"url"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnection) implementsObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionUnion() {
}

// A Slack connection, which either includes a channel_id or a user_id
//
// Satisfied by
// [ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionsSlackTokenConnection],
// [ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection],
// [ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnection].
type ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionUnion interface {
	implementsObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionUnion()
}

// A Slack connection, which either includes a channel_id or a user_id
type ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionsSlackTokenConnection struct {
	AccessToken param.Field[string] `json:"access_token"`
	ChannelID   param.Field[string] `json:"channel_id"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionsSlackTokenConnection) implementsObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionUnion() {
}

// An incoming webhook Slack connection
type ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection struct {
	URL param.Field[string] `json:"url,required"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionsSlackIncomingWebhookConnection) implementsObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataConnectionUnion() {
}

// A token that's used to store the access token for a Slack workspace.
type ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataToken struct {
	AccessToken param.Field[string] `json:"access_token,required"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataSlackChannelDataToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Microsoft Teams channel data
type ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelData struct {
	Connections param.Field[[]ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionUnion] `json:"connections,required"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelData) implementsObjectBulkSetParamsObjectsChannelDataDataUnion() {
}

// Microsoft Teams token connection
type ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnection struct {
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnection) implementsObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams token connection
//
// Satisfied by
// [ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection],
// [ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection],
// [ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnection].
type ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionUnion interface {
	implementsObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionUnion()
}

// Microsoft Teams token connection
type ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection struct {
	// The Microsoft Teams channel ID
	MsTeamsChannelID param.Field[string] `json:"ms_teams_channel_id" format:"uuid"`
	// The Microsoft Teams team ID
	MsTeamsTeamID param.Field[string] `json:"ms_teams_team_id" format:"uuid"`
	// The Microsoft Teams tenant ID
	MsTeamsTenantID param.Field[string] `json:"ms_teams_tenant_id" format:"uuid"`
	// The Microsoft Teams user ID
	MsTeamsUserID param.Field[string] `json:"ms_teams_user_id" format:"uuid"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsTokenConnection) implementsObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// Microsoft Teams incoming webhook connection
type ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnection) implementsObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionUnion() {
}

// The incoming webhook
type ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataMsTeamsChannelDataConnectionsMsTeamsIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Discord channel data
type ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelData struct {
	Connections param.Field[[]ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionUnion] `json:"connections,required"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelData) implementsObjectBulkSetParamsObjectsChannelDataDataUnion() {
}

// Discord channel connection
type ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnection struct {
	// The Discord channel ID
	ChannelID       param.Field[string]      `json:"channel_id"`
	IncomingWebhook param.Field[interface{}] `json:"incoming_webhook"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnection) implementsObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord channel connection
//
// Satisfied by
// [ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection],
// [ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection],
// [ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnection].
type ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionUnion interface {
	implementsObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionUnion()
}

// Discord channel connection
type ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection struct {
	// The Discord channel ID
	ChannelID param.Field[string] `json:"channel_id,required"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordChannelConnection) implementsObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// Discord incoming webhook connection
type ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection struct {
	// The incoming webhook
	IncomingWebhook param.Field[ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook] `json:"incoming_webhook,required"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnection) implementsObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionUnion() {
}

// The incoming webhook
type ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook struct {
	// The URL of the incoming webhook
	URL param.Field[string] `json:"url,required"`
}

func (r ObjectBulkSetParamsObjectsChannelDataDataDiscordChannelDataConnectionsDiscordIncomingWebhookConnectionIncomingWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set preferences for a recipient
type ObjectBulkSetParamsObjectsPreferences struct {
	// A setting for a preference set, where the key in the object is the category, and
	// the values are the preference settings for that category.
	Categories param.Field[map[string]ObjectBulkSetParamsObjectsPreferencesCategoriesUnion] `json:"categories"`
	// Channel type preferences
	ChannelTypes param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypes] `json:"channel_types"`
	// A setting for a preference set, where the key in the object is the workflow key,
	// and the values are the preference settings for that workflow.
	Workflows param.Field[map[string]ObjectBulkSetParamsObjectsPreferencesWorkflowsUnion] `json:"workflows"`
}

func (r ObjectBulkSetParamsObjectsPreferences) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject].
type ObjectBulkSetParamsObjectsPreferencesCategoriesUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesUnion() {
}

// Channel type preferences
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                            `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                            `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                                `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                                `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                          `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                          `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                      `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                      `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Channel type preferences
type ObjectBulkSetParamsObjectsPreferencesChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSUnion] `json:"sms"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesChannelTypesChatUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesChatUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                       `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                       `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesEmailUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                        `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                        `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                       `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                       `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                            `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                            `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesChannelTypesPushUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesPushUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                       `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                       `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesChannelTypesSMSUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                      `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                      `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// Workflow or category preferences within a preference set
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject].
type ObjectBulkSetParamsObjectsPreferencesWorkflowsUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsUnion()
}

// The settings object for a workflow or category, where you can specify channel
// types or conditions.
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject struct {
	// Channel type preferences
	ChannelTypes param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes] `json:"channel_types"`
	Conditions   param.Field[[]ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition]  `json:"conditions"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsUnion() {
}

// Channel type preferences
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes struct {
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Chat param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion] `json:"chat"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Email param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion] `json:"email"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	HTTP param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion] `json:"http"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	InAppFeed param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion] `json:"in_app_feed"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	Push param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion] `json:"push"`
	// A set of settings for a channel type. Currently, this can only be a list of
	// conditions to apply.
	SMS param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion] `json:"sms"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                          `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                          `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                           `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                           `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                          `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                          `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                               `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                               `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                          `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                          `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
//
// Satisfied by [shared.UnionBool],
// [ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject].
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion interface {
	ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion()
}

// A set of settings for a channel type. Currently, this can only be a list of
// conditions to apply.
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject struct {
	Conditions param.Field[[]ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition] `json:"conditions,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObject) ImplementsObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion() {
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition struct {
	Argument param.Field[string]                                                                                                                                                         `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                                                                         `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSPreferenceSetChannelTypeSettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated
type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition struct {
	Argument param.Field[string]                                                                                                     `json:"argument,required"`
	Operator param.Field[ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator] `json:"operator,required"`
	Variable param.Field[string]                                                                                                     `json:"variable,required"`
}

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator string

const (
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo              ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo           ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "greater_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "less_than_or_equal_to"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_contains"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty                ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty             ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "not_empty"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "contains_all"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp          ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp       ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_not_timestamp"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_after"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore    ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_before"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween   ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_timestamp_between"
	ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember     ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator = "is_audience_member"
)

func (r ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperator) IsKnown() bool {
	switch r {
	case ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThan, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorGreaterThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorLessThanOrEqualTo, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotContains, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorNotEmpty, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorContainsAll, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsNotTimestamp, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampAfter, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBefore, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsTimestampBetween, ObjectBulkSetParamsObjectsPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorIsAudienceMember:
		return true
	}
	return false
}

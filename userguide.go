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

// UserGuideService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserGuideService] method instead.
type UserGuideService struct {
	Options []option.RequestOption
}

// NewUserGuideService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserGuideService(opts ...option.RequestOption) (r *UserGuideService) {
	r = &UserGuideService{}
	r.Options = opts
	return
}

// Returns a list of eligible in-app guides for a specific user and channel.
func (r *UserGuideService) GetChannel(ctx context.Context, userID string, channelID string, query UserGuideGetChannelParams, opts ...option.RequestOption) (res *UserGuideGetChannelResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if channelID == "" {
		err = errors.New("missing required channel_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/guides/%s", userID, channelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Records that a guide has been archived by a user, triggering any associated
// archived events.
func (r *UserGuideService) MarkMessageAsArchived(ctx context.Context, userID string, messageID string, body UserGuideMarkMessageAsArchivedParams, opts ...option.RequestOption) (res *UserGuideMarkMessageAsArchivedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/guides/messages/%s/archived", userID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Records that a user has interacted with a guide, triggering any associated
// interacted events.
func (r *UserGuideService) MarkMessageAsInteracted(ctx context.Context, userID string, messageID string, body UserGuideMarkMessageAsInteractedParams, opts ...option.RequestOption) (res *UserGuideMarkMessageAsInteractedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/guides/messages/%s/interacted", userID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Records that a guide has been seen by a user, triggering any associated seen
// events.
func (r *UserGuideService) MarkMessageAsSeen(ctx context.Context, userID string, messageID string, body UserGuideMarkMessageAsSeenParams, opts ...option.RequestOption) (res *UserGuideMarkMessageAsSeenResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/guides/messages/%s/seen", userID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// A response for a list of guides.
type UserGuideGetChannelResponse struct {
	// A list of guides.
	Guides []UserGuideGetChannelResponseGuide `json:"guides,required"`
	// The recipient of the guide.
	Recipient UserGuideGetChannelResponseRecipient `json:"recipient,nullable"`
	JSON      userGuideGetChannelResponseJSON      `json:"-"`
}

// userGuideGetChannelResponseJSON contains the JSON metadata for the struct
// [UserGuideGetChannelResponse]
type userGuideGetChannelResponseJSON struct {
	Guides      apijson.Field
	Recipient   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGuideGetChannelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseJSON) RawJSON() string {
	return r.raw
}

type UserGuideGetChannelResponseGuide struct {
	// The unique identifier for the guide.
	ID string `json:"id" format:"uuid"`
	// The content of the guide.
	Content string `json:"content"`
	// The metadata of the guide.
	Metadata map[string]interface{} `json:"metadata"`
	// The title of the guide.
	Title string                               `json:"title"`
	JSON  userGuideGetChannelResponseGuideJSON `json:"-"`
}

// userGuideGetChannelResponseGuideJSON contains the JSON metadata for the struct
// [UserGuideGetChannelResponseGuide]
type userGuideGetChannelResponseGuideJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Metadata    apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGuideGetChannelResponseGuide) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseGuideJSON) RawJSON() string {
	return r.raw
}

// The recipient of the guide.
type UserGuideGetChannelResponseRecipient struct {
	// Unique identifier for the recipient.
	ID   string                                   `json:"id"`
	JSON userGuideGetChannelResponseRecipientJSON `json:"-"`
}

// userGuideGetChannelResponseRecipientJSON contains the JSON metadata for the
// struct [UserGuideGetChannelResponseRecipient]
type userGuideGetChannelResponseRecipientJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGuideGetChannelResponseRecipient) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseRecipientJSON) RawJSON() string {
	return r.raw
}

// A response for a guide action.
type UserGuideMarkMessageAsArchivedResponse struct {
	// The status of a guide's action.
	Status string                                     `json:"status,required"`
	JSON   userGuideMarkMessageAsArchivedResponseJSON `json:"-"`
}

// userGuideMarkMessageAsArchivedResponseJSON contains the JSON metadata for the
// struct [UserGuideMarkMessageAsArchivedResponse]
type userGuideMarkMessageAsArchivedResponseJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGuideMarkMessageAsArchivedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideMarkMessageAsArchivedResponseJSON) RawJSON() string {
	return r.raw
}

// A response for a guide action.
type UserGuideMarkMessageAsInteractedResponse struct {
	// The status of a guide's action.
	Status string                                       `json:"status,required"`
	JSON   userGuideMarkMessageAsInteractedResponseJSON `json:"-"`
}

// userGuideMarkMessageAsInteractedResponseJSON contains the JSON metadata for the
// struct [UserGuideMarkMessageAsInteractedResponse]
type userGuideMarkMessageAsInteractedResponseJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGuideMarkMessageAsInteractedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideMarkMessageAsInteractedResponseJSON) RawJSON() string {
	return r.raw
}

// A response for a guide action.
type UserGuideMarkMessageAsSeenResponse struct {
	// The status of a guide's action.
	Status string                                 `json:"status,required"`
	JSON   userGuideMarkMessageAsSeenResponseJSON `json:"-"`
}

// userGuideMarkMessageAsSeenResponseJSON contains the JSON metadata for the struct
// [UserGuideMarkMessageAsSeenResponse]
type userGuideMarkMessageAsSeenResponseJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGuideMarkMessageAsSeenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideMarkMessageAsSeenResponseJSON) RawJSON() string {
	return r.raw
}

type UserGuideGetChannelParams struct {
	// The data to filter guides by.
	Data param.Field[string] `query:"data"`
	// The tenant ID to filter guides by.
	Tenant param.Field[string] `query:"tenant"`
	// The type of guides to filter by.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [UserGuideGetChannelParams]'s query parameters as
// `url.Values`.
func (r UserGuideGetChannelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserGuideMarkMessageAsArchivedParams struct {
	// The unique identifier for the channel.
	ChannelID param.Field[string] `json:"channel_id,required" format:"uuid"`
	// The unique identifier for the guide.
	GuideID param.Field[string] `json:"guide_id,required" format:"uuid"`
	// The key of the guide.
	GuideKey param.Field[string] `json:"guide_key,required"`
	// The step reference of the guide.
	GuideStepRef param.Field[string] `json:"guide_step_ref,required"`
	// The content of the guide.
	Content param.Field[interface{}] `json:"content"`
	// The data of the guide.
	Data param.Field[interface{}] `json:"data"`
	// Whether the guide is final.
	IsFinal param.Field[bool] `json:"is_final"`
	// The metadata of the guide.
	Metadata param.Field[interface{}] `json:"metadata"`
	// The tenant ID of the guide.
	Tenant param.Field[string] `json:"tenant"`
}

func (r UserGuideMarkMessageAsArchivedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGuideMarkMessageAsInteractedParams struct {
	// The unique identifier for the channel.
	ChannelID param.Field[string] `json:"channel_id,required" format:"uuid"`
	// The unique identifier for the guide.
	GuideID param.Field[string] `json:"guide_id,required" format:"uuid"`
	// The key of the guide.
	GuideKey param.Field[string] `json:"guide_key,required"`
	// The step reference of the guide.
	GuideStepRef param.Field[string] `json:"guide_step_ref,required"`
	// The content of the guide.
	Content param.Field[interface{}] `json:"content"`
	// The data of the guide.
	Data param.Field[interface{}] `json:"data"`
	// Whether the guide is final.
	IsFinal param.Field[bool] `json:"is_final"`
	// The metadata of the guide.
	Metadata param.Field[interface{}] `json:"metadata"`
	// The tenant ID of the guide.
	Tenant param.Field[string] `json:"tenant"`
}

func (r UserGuideMarkMessageAsInteractedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGuideMarkMessageAsSeenParams struct {
	// The unique identifier for the channel.
	ChannelID param.Field[string] `json:"channel_id,required" format:"uuid"`
	// The unique identifier for the guide.
	GuideID param.Field[string] `json:"guide_id,required" format:"uuid"`
	// The key of the guide.
	GuideKey param.Field[string] `json:"guide_key,required"`
	// The step reference of the guide.
	GuideStepRef param.Field[string] `json:"guide_step_ref,required"`
	// The content of the guide.
	Content param.Field[interface{}] `json:"content"`
	// The data of the guide.
	Data param.Field[interface{}] `json:"data"`
	// Whether the guide is final.
	IsFinal param.Field[bool] `json:"is_final"`
	// The metadata of the guide.
	Metadata param.Field[interface{}] `json:"metadata"`
	// The tenant ID of the guide.
	Tenant param.Field[string] `json:"tenant"`
}

func (r UserGuideMarkMessageAsSeenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/apiquery"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
)

// A user is an individual from your system, represented in Knock. They are most
// commonly a recipient of a notification.
//
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	Entries []UserGuideGetChannelResponseEntry `json:"entries" api:"required"`
	// A map of guide group keys to their last display timestamps.
	GuideGroupDisplayLogs map[string]time.Time `json:"guide_group_display_logs" api:"required" format:"date-time"`
	// A list of guide groups with their display sequences and intervals.
	GuideGroups []UserGuideGetChannelResponseGuideGroup `json:"guide_groups" api:"required"`
	// Markers for guides the user is not eligible to see.
	IneligibleGuides []UserGuideGetChannelResponseIneligibleGuide `json:"ineligible_guides" api:"required"`
	JSON             userGuideGetChannelResponseJSON              `json:"-"`
}

// userGuideGetChannelResponseJSON contains the JSON metadata for the struct
// [UserGuideGetChannelResponse]
type userGuideGetChannelResponseJSON struct {
	Entries               apijson.Field
	GuideGroupDisplayLogs apijson.Field
	GuideGroups           apijson.Field
	IneligibleGuides      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *UserGuideGetChannelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseJSON) RawJSON() string {
	return r.raw
}

type UserGuideGetChannelResponseEntry struct {
	// The unique identifier for the guide.
	ID string `json:"id" format:"uuid"`
	// The typename of the schema.
	Typename string `json:"__typename"`
	// A list of URL Patterns to evaluate user's current location to activate the
	// guide, if matched
	ActivationURLPatterns []UserGuideGetChannelResponseEntriesActivationURLPattern `json:"activation_url_patterns"`
	// A list of URL rules to evaluate user's current location to activate the guide,
	// if matched
	ActivationURLRules []UserGuideGetChannelResponseEntriesActivationURLRule `json:"activation_url_rules"`
	// Whether the guide is active.
	Active                 bool      `json:"active"`
	BypassGlobalGroupLimit bool      `json:"bypass_global_group_limit"`
	ChannelID              string    `json:"channel_id" format:"uuid"`
	InsertedAt             time.Time `json:"inserted_at" format:"date-time"`
	// The key of the guide.
	Key    string                                   `json:"key"`
	Semver string                                   `json:"semver"`
	Steps  []UserGuideGetChannelResponseEntriesStep `json:"steps"`
	// The type of the guide.
	Type      string                               `json:"type"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      userGuideGetChannelResponseEntryJSON `json:"-"`
}

// userGuideGetChannelResponseEntryJSON contains the JSON metadata for the struct
// [UserGuideGetChannelResponseEntry]
type userGuideGetChannelResponseEntryJSON struct {
	ID                     apijson.Field
	Typename               apijson.Field
	ActivationURLPatterns  apijson.Field
	ActivationURLRules     apijson.Field
	Active                 apijson.Field
	BypassGlobalGroupLimit apijson.Field
	ChannelID              apijson.Field
	InsertedAt             apijson.Field
	Key                    apijson.Field
	Semver                 apijson.Field
	Steps                  apijson.Field
	Type                   apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *UserGuideGetChannelResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseEntryJSON) RawJSON() string {
	return r.raw
}

type UserGuideGetChannelResponseEntriesActivationURLPattern struct {
	// The directive for the URL pattern ('allow' or 'block')
	Directive string `json:"directive"`
	// The pathname pattern to match (supports wildcards like /\*)
	Pathname string `json:"pathname"`
	// The search query params to match
	Search string                                                     `json:"search"`
	JSON   userGuideGetChannelResponseEntriesActivationURLPatternJSON `json:"-"`
}

// userGuideGetChannelResponseEntriesActivationURLPatternJSON contains the JSON
// metadata for the struct [UserGuideGetChannelResponseEntriesActivationURLPattern]
type userGuideGetChannelResponseEntriesActivationURLPatternJSON struct {
	Directive   apijson.Field
	Pathname    apijson.Field
	Search      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGuideGetChannelResponseEntriesActivationURLPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseEntriesActivationURLPatternJSON) RawJSON() string {
	return r.raw
}

type UserGuideGetChannelResponseEntriesActivationURLRule struct {
	// The value to compare against
	Argument string `json:"argument"`
	// The directive for the URL rule ('allow' or 'block')
	Directive string `json:"directive"`
	// The comparison operator ('contains' or 'equal_to')
	Operator string `json:"operator"`
	// The variable to evaluate ('pathname')
	Variable string                                                  `json:"variable"`
	JSON     userGuideGetChannelResponseEntriesActivationURLRuleJSON `json:"-"`
}

// userGuideGetChannelResponseEntriesActivationURLRuleJSON contains the JSON
// metadata for the struct [UserGuideGetChannelResponseEntriesActivationURLRule]
type userGuideGetChannelResponseEntriesActivationURLRuleJSON struct {
	Argument    apijson.Field
	Directive   apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGuideGetChannelResponseEntriesActivationURLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseEntriesActivationURLRuleJSON) RawJSON() string {
	return r.raw
}

type UserGuideGetChannelResponseEntriesStep struct {
	Content          map[string]interface{}                         `json:"content"`
	Message          UserGuideGetChannelResponseEntriesStepsMessage `json:"message"`
	Ref              string                                         `json:"ref"`
	SchemaKey        string                                         `json:"schema_key"`
	SchemaSemver     string                                         `json:"schema_semver"`
	SchemaVariantKey string                                         `json:"schema_variant_key"`
	JSON             userGuideGetChannelResponseEntriesStepJSON     `json:"-"`
}

// userGuideGetChannelResponseEntriesStepJSON contains the JSON metadata for the
// struct [UserGuideGetChannelResponseEntriesStep]
type userGuideGetChannelResponseEntriesStepJSON struct {
	Content          apijson.Field
	Message          apijson.Field
	Ref              apijson.Field
	SchemaKey        apijson.Field
	SchemaSemver     apijson.Field
	SchemaVariantKey apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGuideGetChannelResponseEntriesStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseEntriesStepJSON) RawJSON() string {
	return r.raw
}

type UserGuideGetChannelResponseEntriesStepsMessage struct {
	ID            string                                             `json:"id" api:"nullable"`
	ArchivedAt    time.Time                                          `json:"archived_at" api:"nullable" format:"date-time"`
	InteractedAt  time.Time                                          `json:"interacted_at" api:"nullable" format:"date-time"`
	LinkClickedAt time.Time                                          `json:"link_clicked_at" api:"nullable" format:"date-time"`
	ReadAt        time.Time                                          `json:"read_at" api:"nullable" format:"date-time"`
	SeenAt        time.Time                                          `json:"seen_at" api:"nullable" format:"date-time"`
	JSON          userGuideGetChannelResponseEntriesStepsMessageJSON `json:"-"`
}

// userGuideGetChannelResponseEntriesStepsMessageJSON contains the JSON metadata
// for the struct [UserGuideGetChannelResponseEntriesStepsMessage]
type userGuideGetChannelResponseEntriesStepsMessageJSON struct {
	ID            apijson.Field
	ArchivedAt    apijson.Field
	InteractedAt  apijson.Field
	LinkClickedAt apijson.Field
	ReadAt        apijson.Field
	SeenAt        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserGuideGetChannelResponseEntriesStepsMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseEntriesStepsMessageJSON) RawJSON() string {
	return r.raw
}

type UserGuideGetChannelResponseGuideGroup struct {
	Typename        string                                    `json:"__typename"`
	DisplayInterval int64                                     `json:"display_interval"`
	DisplaySequence []string                                  `json:"display_sequence"`
	InsertedAt      time.Time                                 `json:"inserted_at" format:"date-time"`
	Key             string                                    `json:"key"`
	UpdatedAt       time.Time                                 `json:"updated_at" format:"date-time"`
	JSON            userGuideGetChannelResponseGuideGroupJSON `json:"-"`
}

// userGuideGetChannelResponseGuideGroupJSON contains the JSON metadata for the
// struct [UserGuideGetChannelResponseGuideGroup]
type userGuideGetChannelResponseGuideGroupJSON struct {
	Typename        apijson.Field
	DisplayInterval apijson.Field
	DisplaySequence apijson.Field
	InsertedAt      apijson.Field
	Key             apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserGuideGetChannelResponseGuideGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseGuideGroupJSON) RawJSON() string {
	return r.raw
}

type UserGuideGetChannelResponseIneligibleGuide struct {
	// The guide's key identifier
	Key string `json:"key" api:"required"`
	// Human-readable explanation of ineligibility
	Message string `json:"message" api:"required"`
	// Reason code for ineligibility
	Reason UserGuideGetChannelResponseIneligibleGuidesReason `json:"reason" api:"required"`
	JSON   userGuideGetChannelResponseIneligibleGuideJSON    `json:"-"`
}

// userGuideGetChannelResponseIneligibleGuideJSON contains the JSON metadata for
// the struct [UserGuideGetChannelResponseIneligibleGuide]
type userGuideGetChannelResponseIneligibleGuideJSON struct {
	Key         apijson.Field
	Message     apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGuideGetChannelResponseIneligibleGuide) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGuideGetChannelResponseIneligibleGuideJSON) RawJSON() string {
	return r.raw
}

// Reason code for ineligibility
type UserGuideGetChannelResponseIneligibleGuidesReason string

const (
	UserGuideGetChannelResponseIneligibleGuidesReasonGuideNotActive         UserGuideGetChannelResponseIneligibleGuidesReason = "guide_not_active"
	UserGuideGetChannelResponseIneligibleGuidesReasonMarkedAsArchived       UserGuideGetChannelResponseIneligibleGuidesReason = "marked_as_archived"
	UserGuideGetChannelResponseIneligibleGuidesReasonTargetConditionsNotMet UserGuideGetChannelResponseIneligibleGuidesReason = "target_conditions_not_met"
	UserGuideGetChannelResponseIneligibleGuidesReasonNotInTargetAudience    UserGuideGetChannelResponseIneligibleGuidesReason = "not_in_target_audience"
)

func (r UserGuideGetChannelResponseIneligibleGuidesReason) IsKnown() bool {
	switch r {
	case UserGuideGetChannelResponseIneligibleGuidesReasonGuideNotActive, UserGuideGetChannelResponseIneligibleGuidesReasonMarkedAsArchived, UserGuideGetChannelResponseIneligibleGuidesReasonTargetConditionsNotMet, UserGuideGetChannelResponseIneligibleGuidesReasonNotInTargetAudience:
		return true
	}
	return false
}

// A response for a guide action.
type UserGuideMarkMessageAsArchivedResponse struct {
	// The status of a guide's action.
	Status string                                     `json:"status" api:"required"`
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
	Status string                                       `json:"status" api:"required"`
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
	Status string                                 `json:"status" api:"required"`
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
	// The data (JSON encoded object) to use for targeting and rendering guides.
	Data param.Field[string] `query:"data"`
	// The tenant ID to use for targeting and rendering guides.
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
	ChannelID param.Field[string] `json:"channel_id" api:"required" format:"uuid"`
	// The unique identifier for the guide.
	GuideID param.Field[string] `json:"guide_id" api:"required" format:"uuid"`
	// The key of the guide.
	GuideKey param.Field[string] `json:"guide_key" api:"required"`
	// The step reference of the guide.
	GuideStepRef param.Field[string] `json:"guide_step_ref" api:"required"`
	// The content of the guide.
	Content param.Field[map[string]interface{}] `json:"content"`
	// The data of the guide.
	Data param.Field[map[string]interface{}] `json:"data"`
	// Whether the guide is final.
	IsFinal param.Field[bool] `json:"is_final"`
	// The metadata of the guide.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// The tenant ID of the guide.
	Tenant param.Field[string] `json:"tenant"`
}

func (r UserGuideMarkMessageAsArchivedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGuideMarkMessageAsInteractedParams struct {
	// The unique identifier for the channel.
	ChannelID param.Field[string] `json:"channel_id" api:"required" format:"uuid"`
	// The unique identifier for the guide.
	GuideID param.Field[string] `json:"guide_id" api:"required" format:"uuid"`
	// The key of the guide.
	GuideKey param.Field[string] `json:"guide_key" api:"required"`
	// The step reference of the guide.
	GuideStepRef param.Field[string] `json:"guide_step_ref" api:"required"`
	// The content of the guide.
	Content param.Field[map[string]interface{}] `json:"content"`
	// The data of the guide.
	Data param.Field[map[string]interface{}] `json:"data"`
	// Whether the guide is final.
	IsFinal param.Field[bool] `json:"is_final"`
	// The metadata of the guide.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// The tenant ID of the guide.
	Tenant param.Field[string] `json:"tenant"`
}

func (r UserGuideMarkMessageAsInteractedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGuideMarkMessageAsSeenParams struct {
	// The unique identifier for the channel.
	ChannelID param.Field[string] `json:"channel_id" api:"required" format:"uuid"`
	// The unique identifier for the guide.
	GuideID param.Field[string] `json:"guide_id" api:"required" format:"uuid"`
	// The key of the guide.
	GuideKey param.Field[string] `json:"guide_key" api:"required"`
	// The step reference of the guide.
	GuideStepRef param.Field[string] `json:"guide_step_ref" api:"required"`
	// The content of the guide.
	Content param.Field[map[string]interface{}] `json:"content"`
	// The data of the guide.
	Data param.Field[map[string]interface{}] `json:"data"`
	// Whether the guide is final.
	IsFinal param.Field[bool] `json:"is_final"`
	// The metadata of the guide.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// The tenant ID of the guide.
	Tenant param.Field[string] `json:"tenant"`
}

func (r UserGuideMarkMessageAsSeenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

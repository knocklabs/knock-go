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
	"github.com/knocklabs/knock-go/shared"
)

// AudienceService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudienceService] method instead.
type AudienceService struct {
	Options []option.RequestOption
}

// NewAudienceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudienceService(opts ...option.RequestOption) (r *AudienceService) {
	r = &AudienceService{}
	r.Options = opts
	return
}

// Adds one or more members to the specified audience.
func (r *AudienceService) AddMembers(ctx context.Context, key string, params AudienceAddMembersParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/members", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Returns a paginated list of members for the specified audience.
func (r *AudienceService) ListMembers(ctx context.Context, key string, opts ...option.RequestOption) (res *AudienceListMembersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/members", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes one or more members from the specified audience.
func (r *AudienceService) RemoveMembers(ctx context.Context, key string, body AudienceRemoveMembersParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/members", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// An audience member.
type AudienceMember struct {
	// The typename of the schema.
	Typename string `json:"__typename" api:"required"`
	// Timestamp when the resource was created.
	AddedAt time.Time `json:"added_at" api:"required" format:"date-time"`
	// A [User](/concepts/users) represents an individual in your system who can
	// receive notifications through Knock. Users are the most common recipients of
	// notifications and are always referenced by your internal identifier.
	User User `json:"user" api:"required"`
	// The unique identifier of the user.
	UserID string `json:"user_id" api:"required"`
	// The unique identifier for the tenant.
	Tenant string             `json:"tenant" api:"nullable"`
	JSON   audienceMemberJSON `json:"-"`
}

// audienceMemberJSON contains the JSON metadata for the struct [AudienceMember]
type audienceMemberJSON struct {
	Typename    apijson.Field
	AddedAt     apijson.Field
	User        apijson.Field
	UserID      apijson.Field
	Tenant      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudienceMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audienceMemberJSON) RawJSON() string {
	return r.raw
}

// A paginated list of audience members.
type AudienceListMembersResponse struct {
	// A list of audience members.
	Entries []AudienceMember `json:"entries" api:"required"`
	// Pagination information for a list of resources.
	PageInfo shared.PageInfo                 `json:"page_info" api:"required"`
	JSON     audienceListMembersResponseJSON `json:"-"`
}

// audienceListMembersResponseJSON contains the JSON metadata for the struct
// [AudienceListMembersResponse]
type audienceListMembersResponseJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudienceListMembersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audienceListMembersResponseJSON) RawJSON() string {
	return r.raw
}

type AudienceAddMembersParams struct {
	// A list of audience members to add. You can add up to 1,000 members per request.
	Members param.Field[[]AudienceAddMembersParamsMember] `json:"members" api:"required"`
	// Create the audience if it does not exist.
	CreateAudience param.Field[bool] `query:"create_audience"`
}

func (r AudienceAddMembersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AudienceAddMembersParams]'s query parameters as
// `url.Values`.
func (r AudienceAddMembersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// An audience member.
type AudienceAddMembersParamsMember struct {
	// A set of parameters to inline-identify a user with. Inline identifying the user
	// will ensure that the user is available before the request is executed in Knock.
	// It will perform an upsert for the user you're supplying, replacing any
	// properties specified.
	User param.Field[InlineIdentifyUserRequestParam] `json:"user" api:"required"`
	// The unique identifier for the tenant.
	Tenant param.Field[string] `json:"tenant"`
}

func (r AudienceAddMembersParamsMember) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AudienceRemoveMembersParams struct {
	// A list of audience members to remove. You can remove up to 1,000 members per
	// request.
	Members param.Field[[]AudienceRemoveMembersParamsMember] `json:"members" api:"required"`
}

func (r AudienceRemoveMembersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An audience member.
type AudienceRemoveMembersParamsMember struct {
	// A set of parameters to inline-identify a user with. Inline identifying the user
	// will ensure that the user is available before the request is executed in Knock.
	// It will perform an upsert for the user you're supplying, replacing any
	// properties specified.
	User param.Field[InlineIdentifyUserRequestParam] `json:"user" api:"required"`
	// The unique identifier for the tenant.
	Tenant param.Field[string] `json:"tenant"`
}

func (r AudienceRemoveMembersParamsMember) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

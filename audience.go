// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
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

// Add members to an audience
func (r *AudienceService) AddMembers(ctx context.Context, key string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/members", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List members of an audience
func (r *AudienceService) ListMembers(ctx context.Context, key string, opts ...option.RequestOption) (res *AudienceListMembersResponse, err error) {
	opts = append(r.Options[:], opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/members", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove members from an audience
func (r *AudienceService) RemoveMembers(ctx context.Context, key string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/members", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A response containing a list of audience members
type AudienceListMembersResponse struct {
	Entries []AudienceListMembersResponseEntry `json:"entries,required"`
	// The information about a paginated result
	PageInfo AudienceListMembersResponsePageInfo `json:"page_info,required"`
	JSON     audienceListMembersResponseJSON     `json:"-"`
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

// A user belonging to an audience
type AudienceListMembersResponseEntry struct {
	Typename string `json:"__typename,required"`
	AddedAt  string `json:"added_at,required" format:"date_time"`
	// A user object
	User   User                                 `json:"user,required"`
	UserID string                               `json:"user_id,required"`
	Tenant string                               `json:"tenant,nullable"`
	JSON   audienceListMembersResponseEntryJSON `json:"-"`
}

// audienceListMembersResponseEntryJSON contains the JSON metadata for the struct
// [AudienceListMembersResponseEntry]
type audienceListMembersResponseEntryJSON struct {
	Typename    apijson.Field
	AddedAt     apijson.Field
	User        apijson.Field
	UserID      apijson.Field
	Tenant      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudienceListMembersResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audienceListMembersResponseEntryJSON) RawJSON() string {
	return r.raw
}

// The information about a paginated result
type AudienceListMembersResponsePageInfo struct {
	Typename string                                  `json:"__typename,required"`
	PageSize int64                                   `json:"page_size,required"`
	After    string                                  `json:"after,nullable"`
	Before   string                                  `json:"before,nullable"`
	JSON     audienceListMembersResponsePageInfoJSON `json:"-"`
}

// audienceListMembersResponsePageInfoJSON contains the JSON metadata for the
// struct [AudienceListMembersResponsePageInfo]
type audienceListMembersResponsePageInfoJSON struct {
	Typename    apijson.Field
	PageSize    apijson.Field
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudienceListMembersResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audienceListMembersResponsePageInfoJSON) RawJSON() string {
	return r.raw
}

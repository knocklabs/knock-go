// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/apiquery"
	"github.com/knocklabs/knock-go/internal/param"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/packages/pagination"
	"github.com/tidwall/gjson"
)

// UserFeedService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserFeedService] method instead.
type UserFeedService struct {
	Options []option.RequestOption
}

// NewUserFeedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserFeedService(opts ...option.RequestOption) (r *UserFeedService) {
	r = &UserFeedService{}
	r.Options = opts
	return
}

// Returns the feed settings for a user.
func (r *UserFeedService) GetSettings(ctx context.Context, userID string, id string, opts ...option.RequestOption) (res *UserFeedGetSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/feeds/%s/settings", userID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a paginated list of feed items for a user in reverse chronological
// order, including metadata about the feed. If the user has not yet been
// identified within Knock, an empty feed will be returned.
//
// You can customize the response using
// [response filters](/integrations/in-app/knock#customizing-api-response-content)
// to exclude or only include specific properties on your resources.
//
// **Notes:**
//
//   - When making this call from a client-side environment, use your publishable key
//     along with a user token.
//   - This endpoint’s rate limit is always scoped per-user and per-environment. This
//     is true even for requests made without a signed user token.
//   - Any [attachments](/integrations/email/attachments) present in trigger data are
//     automatically excluded from both the `data` and `activities` fields of
//     `UserInAppFeedResponse`.
func (r *UserFeedService) ListItems(ctx context.Context, userID string, id string, query UserFeedListItemsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[UserFeedListItemsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/feeds/%s", userID, id)
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

// Returns a paginated list of feed items for a user in reverse chronological
// order, including metadata about the feed. If the user has not yet been
// identified within Knock, an empty feed will be returned.
//
// You can customize the response using
// [response filters](/integrations/in-app/knock#customizing-api-response-content)
// to exclude or only include specific properties on your resources.
//
// **Notes:**
//
//   - When making this call from a client-side environment, use your publishable key
//     along with a user token.
//   - This endpoint’s rate limit is always scoped per-user and per-environment. This
//     is true even for requests made without a signed user token.
//   - Any [attachments](/integrations/email/attachments) present in trigger data are
//     automatically excluded from both the `data` and `activities` fields of
//     `UserInAppFeedResponse`.
func (r *UserFeedService) ListItemsAutoPaging(ctx context.Context, userID string, id string, query UserFeedListItemsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[UserFeedListItemsResponse] {
	return pagination.NewEntriesCursorAutoPager(r.ListItems(ctx, userID, id, query, opts...))
}

// The response for the user's feed settings.
type UserFeedGetSettingsResponse struct {
	// Features configuration for the user's feed.
	Features UserFeedGetSettingsResponseFeatures `json:"features,required"`
	JSON     userFeedGetSettingsResponseJSON     `json:"-"`
}

// userFeedGetSettingsResponseJSON contains the JSON metadata for the struct
// [UserFeedGetSettingsResponse]
type userFeedGetSettingsResponseJSON struct {
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetSettingsResponseJSON) RawJSON() string {
	return r.raw
}

// Features configuration for the user's feed.
type UserFeedGetSettingsResponseFeatures struct {
	// Whether branding is required for the user's feed.
	BrandingRequired bool                                    `json:"branding_required,required"`
	JSON             userFeedGetSettingsResponseFeaturesJSON `json:"-"`
}

// userFeedGetSettingsResponseFeaturesJSON contains the JSON metadata for the
// struct [UserFeedGetSettingsResponseFeatures]
type userFeedGetSettingsResponseFeaturesJSON struct {
	BrandingRequired apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserFeedGetSettingsResponseFeatures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetSettingsResponseFeaturesJSON) RawJSON() string {
	return r.raw
}

// An in-app feed message in a user's feed.
type UserFeedListItemsResponse struct {
	// Unique identifier for the feed.
	ID string `json:"id,required"`
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// List of activities associated with this feed item.
	Activities []Activity `json:"activities,required"`
	// List of actors associated with this feed item.
	Actors []Recipient `json:"actors,required"`
	// Content blocks that make up the feed item.
	Blocks []UserFeedListItemsResponseBlock `json:"blocks,required"`
	// Additional data associated with the feed item.
	Data map[string]interface{} `json:"data,required,nullable"`
	// Timestamp when the resource was created.
	InsertedAt string `json:"inserted_at,required"`
	// Source information for the feed item.
	Source UserFeedListItemsResponseSource `json:"source,required"`
	// Tenant ID that the feed item belongs to.
	Tenant string `json:"tenant,required,nullable"`
	// Total number of activities related to this feed item.
	TotalActivities int64 `json:"total_activities,required"`
	// Total number of actors related to this feed item.
	TotalActors int64 `json:"total_actors,required"`
	// The timestamp when the resource was last updated.
	UpdatedAt string `json:"updated_at,required"`
	// Timestamp when the feed item was archived.
	ArchivedAt string `json:"archived_at,nullable"`
	// Timestamp when the feed item was clicked.
	ClickedAt string `json:"clicked_at,nullable"`
	// Timestamp when the feed item was interacted with.
	InteractedAt string `json:"interacted_at,nullable"`
	// Timestamp when a link within the feed item was clicked.
	LinkClickedAt string `json:"link_clicked_at,nullable"`
	// Timestamp when the feed item was marked as read.
	ReadAt string `json:"read_at,nullable"`
	// Timestamp when the feed item was marked as seen.
	SeenAt string                        `json:"seen_at,nullable"`
	JSON   userFeedListItemsResponseJSON `json:"-"`
}

// userFeedListItemsResponseJSON contains the JSON metadata for the struct
// [UserFeedListItemsResponse]
type userFeedListItemsResponseJSON struct {
	ID              apijson.Field
	Typename        apijson.Field
	Activities      apijson.Field
	Actors          apijson.Field
	Blocks          apijson.Field
	Data            apijson.Field
	InsertedAt      apijson.Field
	Source          apijson.Field
	Tenant          apijson.Field
	TotalActivities apijson.Field
	TotalActors     apijson.Field
	UpdatedAt       apijson.Field
	ArchivedAt      apijson.Field
	ClickedAt       apijson.Field
	InteractedAt    apijson.Field
	LinkClickedAt   apijson.Field
	ReadAt          apijson.Field
	SeenAt          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserFeedListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedListItemsResponseJSON) RawJSON() string {
	return r.raw
}

// A content block for the feed, can be content or a button set.
type UserFeedListItemsResponseBlock struct {
	// The name of the block in a message in an app feed.
	Name string `json:"name,required"`
	// The type of block in a message in an app feed.
	Type UserFeedListItemsResponseBlocksType `json:"type,required"`
	// This field can have the runtime type of
	// [[]UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButton].
	Buttons interface{} `json:"buttons"`
	// The content of the block in a message in an app feed.
	Content string `json:"content"`
	// The rendered HTML version of the content.
	Rendered string                             `json:"rendered"`
	JSON     userFeedListItemsResponseBlockJSON `json:"-"`
	union    UserFeedListItemsResponseBlocksUnion
}

// userFeedListItemsResponseBlockJSON contains the JSON metadata for the struct
// [UserFeedListItemsResponseBlock]
type userFeedListItemsResponseBlockJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Buttons     apijson.Field
	Content     apijson.Field
	Rendered    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userFeedListItemsResponseBlockJSON) RawJSON() string {
	return r.raw
}

func (r *UserFeedListItemsResponseBlock) UnmarshalJSON(data []byte) (err error) {
	*r = UserFeedListItemsResponseBlock{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserFeedListItemsResponseBlocksUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserFeedListItemsResponseBlocksMessageInAppFeedContentBlock],
// [UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlock].
func (r UserFeedListItemsResponseBlock) AsUnion() UserFeedListItemsResponseBlocksUnion {
	return r.union
}

// A content block for the feed, can be content or a button set.
//
// Union satisfied by [UserFeedListItemsResponseBlocksMessageInAppFeedContentBlock]
// or [UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlock].
type UserFeedListItemsResponseBlocksUnion interface {
	implementsUserFeedListItemsResponseBlock()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFeedListItemsResponseBlocksUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedListItemsResponseBlocksMessageInAppFeedContentBlock{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlock{}),
		},
	)
}

// A block in a message in an app feed.
type UserFeedListItemsResponseBlocksMessageInAppFeedContentBlock struct {
	// The content of the block in a message in an app feed.
	Content string `json:"content,required"`
	// The name of the block in a message in an app feed.
	Name string `json:"name,required"`
	// The rendered HTML version of the content.
	Rendered string `json:"rendered,required"`
	// The type of block in a message in an app feed.
	Type UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockType `json:"type,required"`
	JSON userFeedListItemsResponseBlocksMessageInAppFeedContentBlockJSON `json:"-"`
}

// userFeedListItemsResponseBlocksMessageInAppFeedContentBlockJSON contains the
// JSON metadata for the struct
// [UserFeedListItemsResponseBlocksMessageInAppFeedContentBlock]
type userFeedListItemsResponseBlocksMessageInAppFeedContentBlockJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Rendered    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedListItemsResponseBlocksMessageInAppFeedContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedListItemsResponseBlocksMessageInAppFeedContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedListItemsResponseBlocksMessageInAppFeedContentBlock) implementsUserFeedListItemsResponseBlock() {
}

// The type of block in a message in an app feed.
type UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockType string

const (
	UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockTypeMarkdown UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockType = "markdown"
	UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockTypeText     UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockType = "text"
)

func (r UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockType) IsKnown() bool {
	switch r {
	case UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockTypeMarkdown, UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockTypeText:
		return true
	}
	return false
}

// A button set block in a message in an app feed.
type UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlock struct {
	// A list of buttons in an in app feed message.
	Buttons []UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButton `json:"buttons,required"`
	// The name of the button set in a message in an app feed.
	Name string `json:"name,required"`
	// The type of block in a message in an app feed.
	Type UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockType `json:"type,required"`
	JSON userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockJSON `json:"-"`
}

// userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockJSON contains the
// JSON metadata for the struct
// [UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlock]
type userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockJSON struct {
	Buttons     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlock) implementsUserFeedListItemsResponseBlock() {
}

// A button in an in app feed message.
type UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButton struct {
	// The action to take when the button is clicked.
	Action string `json:"action,required"`
	// The label of the button.
	Label string `json:"label,required"`
	// The name of the button.
	Name string                                                                  `json:"name,required"`
	JSON userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButtonJSON `json:"-"`
}

// userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButtonJSON contains
// the JSON metadata for the struct
// [UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButton]
type userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButtonJSON struct {
	Action      apijson.Field
	Label       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButton) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButtonJSON) RawJSON() string {
	return r.raw
}

// The type of block in a message in an app feed.
type UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockType string

const (
	UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockTypeButtonSet UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockType = "button_set"
)

func (r UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockType) IsKnown() bool {
	switch r {
	case UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockTypeButtonSet:
		return true
	}
	return false
}

// The type of block in a message in an app feed.
type UserFeedListItemsResponseBlocksType string

const (
	UserFeedListItemsResponseBlocksTypeMarkdown  UserFeedListItemsResponseBlocksType = "markdown"
	UserFeedListItemsResponseBlocksTypeText      UserFeedListItemsResponseBlocksType = "text"
	UserFeedListItemsResponseBlocksTypeButtonSet UserFeedListItemsResponseBlocksType = "button_set"
)

func (r UserFeedListItemsResponseBlocksType) IsKnown() bool {
	switch r {
	case UserFeedListItemsResponseBlocksTypeMarkdown, UserFeedListItemsResponseBlocksTypeText, UserFeedListItemsResponseBlocksTypeButtonSet:
		return true
	}
	return false
}

// Source information for the feed item.
type UserFeedListItemsResponseSource struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// Categories this workflow belongs to.
	Categories []string `json:"categories,required"`
	// The key of the workflow.
	Key string `json:"key,required"`
	// The workflow version ID.
	VersionID string                              `json:"version_id,required" format:"uuid"`
	JSON      userFeedListItemsResponseSourceJSON `json:"-"`
}

// userFeedListItemsResponseSourceJSON contains the JSON metadata for the struct
// [UserFeedListItemsResponseSource]
type userFeedListItemsResponseSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedListItemsResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedListItemsResponseSourceJSON) RawJSON() string {
	return r.raw
}

type UserFeedListItemsParams struct {
	// The cursor to fetch entries after.
	After param.Field[string] `query:"after"`
	// The archived status of the feed items.
	Archived param.Field[UserFeedListItemsParamsArchived] `query:"archived"`
	// The cursor to fetch entries before.
	Before param.Field[string] `query:"before"`
	// Comma-separated list of field paths to exclude from the response. Use dot
	// notation for nested fields (e.g., `entries.archived_at`). Limited to 3 levels
	// deep.
	Exclude param.Field[string] `query:"exclude"`
	// Whether the feed items have a tenant.
	HasTenant  param.Field[bool]                              `query:"has_tenant"`
	InsertedAt param.Field[UserFeedListItemsParamsInsertedAt] `query:"inserted_at"`
	// The locale to render the feed items in. Must be in the IETF 5646 format (e.g.
	// `en-US`). When not provided, will default to the locale that the feed items were
	// rendered in. Only available for enterprise plan customers using custom
	// translations.
	Locale param.Field[string] `query:"locale"`
	// The number of items per page (defaults to 50).
	PageSize param.Field[int64] `query:"page_size"`
	// The workflow key associated with the message in the feed.
	Source param.Field[string] `query:"source"`
	// The status of the feed items.
	Status param.Field[UserFeedListItemsParamsStatus] `query:"status"`
	// The tenant associated with the feed items.
	Tenant param.Field[string] `query:"tenant"`
	// The trigger data of the feed items (as a JSON string).
	TriggerData param.Field[string] `query:"trigger_data"`
	// The workflow categories of the feed items.
	WorkflowCategories param.Field[[]string] `query:"workflow_categories"`
}

// URLQuery serializes [UserFeedListItemsParams]'s query parameters as
// `url.Values`.
func (r UserFeedListItemsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The archived status of the feed items.
type UserFeedListItemsParamsArchived string

const (
	UserFeedListItemsParamsArchivedExclude UserFeedListItemsParamsArchived = "exclude"
	UserFeedListItemsParamsArchivedInclude UserFeedListItemsParamsArchived = "include"
	UserFeedListItemsParamsArchivedOnly    UserFeedListItemsParamsArchived = "only"
)

func (r UserFeedListItemsParamsArchived) IsKnown() bool {
	switch r {
	case UserFeedListItemsParamsArchivedExclude, UserFeedListItemsParamsArchivedInclude, UserFeedListItemsParamsArchivedOnly:
		return true
	}
	return false
}

type UserFeedListItemsParamsInsertedAt struct {
	// Limits the results to items inserted after the given date.
	Gt param.Field[string] `query:"gt"`
	// Limits the results to items inserted after or on the given date.
	Gte param.Field[string] `query:"gte"`
	// Limits the results to items inserted before the given date.
	Lt param.Field[string] `query:"lt"`
	// Limits the results to items inserted before or on the given date.
	Lte param.Field[string] `query:"lte"`
}

// URLQuery serializes [UserFeedListItemsParamsInsertedAt]'s query parameters as
// `url.Values`.
func (r UserFeedListItemsParamsInsertedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The status of the feed items.
type UserFeedListItemsParamsStatus string

const (
	UserFeedListItemsParamsStatusUnread UserFeedListItemsParamsStatus = "unread"
	UserFeedListItemsParamsStatusRead   UserFeedListItemsParamsStatus = "read"
	UserFeedListItemsParamsStatusUnseen UserFeedListItemsParamsStatus = "unseen"
	UserFeedListItemsParamsStatusSeen   UserFeedListItemsParamsStatus = "seen"
	UserFeedListItemsParamsStatusAll    UserFeedListItemsParamsStatus = "all"
)

func (r UserFeedListItemsParamsStatus) IsKnown() bool {
	switch r {
	case UserFeedListItemsParamsStatusUnread, UserFeedListItemsParamsStatusRead, UserFeedListItemsParamsStatusUnseen, UserFeedListItemsParamsStatusSeen, UserFeedListItemsParamsStatusAll:
		return true
	}
	return false
}

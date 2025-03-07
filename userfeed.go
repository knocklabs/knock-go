// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/apiquery"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/packages/pagination"
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
	opts = append(r.Options[:], opts...)
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

// Returns a paginated list of feed items for a user, including metadata about the
// feed.
func (r *UserFeedService) ListItems(ctx context.Context, userID string, id string, query UserFeedListItemsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[UserFeedListItemsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// Returns a paginated list of feed items for a user, including metadata about the
// feed.
func (r *UserFeedService) ListItemsAutoPaging(ctx context.Context, userID string, id string, query UserFeedListItemsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[UserFeedListItemsResponse] {
	return pagination.NewEntriesCursorAutoPager(r.ListItems(ctx, userID, id, query, opts...))
}

// The response for the user's feed settings
type UserFeedGetSettingsResponse struct {
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

type UserFeedGetSettingsResponseFeatures struct {
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

// An in-app feed message in a user's feed
type UserFeedListItemsResponse struct {
	ID              string                              `json:"id,required"`
	Typename        string                              `json:"__typename,required"`
	Activities      []UserFeedListItemsResponseActivity `json:"activities,required"`
	Actors          []UserFeedListItemsResponseActor    `json:"actors,required"`
	Blocks          []UserFeedListItemsResponseBlock    `json:"blocks,required"`
	Data            map[string]interface{}              `json:"data,required,nullable"`
	InsertedAt      string                              `json:"inserted_at,required"`
	Source          UserFeedListItemsResponseSource     `json:"source,required"`
	Tenant          string                              `json:"tenant,required,nullable"`
	TotalActivities int64                               `json:"total_activities,required"`
	TotalActors     int64                               `json:"total_actors,required"`
	UpdatedAt       string                              `json:"updated_at,required"`
	ArchivedAt      string                              `json:"archived_at,nullable"`
	ClickedAt       string                              `json:"clicked_at,nullable"`
	InteractedAt    string                              `json:"interacted_at,nullable"`
	LinkClickedAt   string                              `json:"link_clicked_at,nullable"`
	ReadAt          string                              `json:"read_at,nullable"`
	SeenAt          string                              `json:"seen_at,nullable"`
	JSON            userFeedListItemsResponseJSON       `json:"-"`
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

// An activity associated with a workflow run
type UserFeedListItemsResponseActivity struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor UserFeedListItemsResponseActivitiesActor `json:"actor,nullable"`
	// The data associated with the activity
	Data       map[string]interface{} `json:"data,nullable"`
	InsertedAt time.Time              `json:"inserted_at" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient UserFeedListItemsResponseActivitiesRecipient `json:"recipient"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      userFeedListItemsResponseActivityJSON        `json:"-"`
}

// userFeedListItemsResponseActivityJSON contains the JSON metadata for the struct
// [UserFeedListItemsResponseActivity]
type userFeedListItemsResponseActivityJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Actor       apijson.Field
	Data        apijson.Field
	InsertedAt  apijson.Field
	Recipient   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedListItemsResponseActivity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedListItemsResponseActivityJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type UserFeedListItemsResponseActivitiesActor struct {
	ID          string                                       `json:"id,required"`
	Typename    string                                       `json:"__typename,required"`
	UpdatedAt   time.Time                                    `json:"updated_at,required" format:"date-time"`
	Avatar      string                                       `json:"avatar,nullable"`
	Collection  string                                       `json:"collection"`
	CreatedAt   time.Time                                    `json:"created_at,nullable" format:"date-time"`
	Email       string                                       `json:"email,nullable"`
	Name        string                                       `json:"name,nullable"`
	PhoneNumber string                                       `json:"phone_number,nullable"`
	Timezone    string                                       `json:"timezone,nullable"`
	JSON        userFeedListItemsResponseActivitiesActorJSON `json:"-"`
	union       UserFeedListItemsResponseActivitiesActorUnion
}

// userFeedListItemsResponseActivitiesActorJSON contains the JSON metadata for the
// struct [UserFeedListItemsResponseActivitiesActor]
type userFeedListItemsResponseActivitiesActorJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userFeedListItemsResponseActivitiesActorJSON) RawJSON() string {
	return r.raw
}

func (r *UserFeedListItemsResponseActivitiesActor) UnmarshalJSON(data []byte) (err error) {
	*r = UserFeedListItemsResponseActivitiesActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserFeedListItemsResponseActivitiesActorUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [UserFeedListItemsResponseActivitiesActorObject].
func (r UserFeedListItemsResponseActivitiesActor) AsUnion() UserFeedListItemsResponseActivitiesActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [UserFeedListItemsResponseActivitiesActorObject].
type UserFeedListItemsResponseActivitiesActorUnion interface {
	implementsUserFeedListItemsResponseActivitiesActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFeedListItemsResponseActivitiesActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedListItemsResponseActivitiesActorObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type UserFeedListItemsResponseActivitiesActorObject struct {
	ID          string                                             `json:"id,required"`
	Typename    string                                             `json:"__typename,required"`
	Collection  string                                             `json:"collection,required"`
	UpdatedAt   time.Time                                          `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                          `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                             `json:"-,extras"`
	JSON        userFeedListItemsResponseActivitiesActorObjectJSON `json:"-"`
}

// userFeedListItemsResponseActivitiesActorObjectJSON contains the JSON metadata
// for the struct [UserFeedListItemsResponseActivitiesActorObject]
type userFeedListItemsResponseActivitiesActorObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedListItemsResponseActivitiesActorObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedListItemsResponseActivitiesActorObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedListItemsResponseActivitiesActorObject) implementsUserFeedListItemsResponseActivitiesActor() {
}

// A recipient, which is either a user or an object
type UserFeedListItemsResponseActivitiesRecipient struct {
	ID          string                                           `json:"id,required"`
	Typename    string                                           `json:"__typename,required"`
	UpdatedAt   time.Time                                        `json:"updated_at,required" format:"date-time"`
	Avatar      string                                           `json:"avatar,nullable"`
	Collection  string                                           `json:"collection"`
	CreatedAt   time.Time                                        `json:"created_at,nullable" format:"date-time"`
	Email       string                                           `json:"email,nullable"`
	Name        string                                           `json:"name,nullable"`
	PhoneNumber string                                           `json:"phone_number,nullable"`
	Timezone    string                                           `json:"timezone,nullable"`
	JSON        userFeedListItemsResponseActivitiesRecipientJSON `json:"-"`
	union       UserFeedListItemsResponseActivitiesRecipientUnion
}

// userFeedListItemsResponseActivitiesRecipientJSON contains the JSON metadata for
// the struct [UserFeedListItemsResponseActivitiesRecipient]
type userFeedListItemsResponseActivitiesRecipientJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userFeedListItemsResponseActivitiesRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *UserFeedListItemsResponseActivitiesRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = UserFeedListItemsResponseActivitiesRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserFeedListItemsResponseActivitiesRecipientUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [UserFeedListItemsResponseActivitiesRecipientObject].
func (r UserFeedListItemsResponseActivitiesRecipient) AsUnion() UserFeedListItemsResponseActivitiesRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or
// [UserFeedListItemsResponseActivitiesRecipientObject].
type UserFeedListItemsResponseActivitiesRecipientUnion interface {
	implementsUserFeedListItemsResponseActivitiesRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFeedListItemsResponseActivitiesRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedListItemsResponseActivitiesRecipientObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type UserFeedListItemsResponseActivitiesRecipientObject struct {
	ID          string                                                 `json:"id,required"`
	Typename    string                                                 `json:"__typename,required"`
	Collection  string                                                 `json:"collection,required"`
	UpdatedAt   time.Time                                              `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                              `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                                 `json:"-,extras"`
	JSON        userFeedListItemsResponseActivitiesRecipientObjectJSON `json:"-"`
}

// userFeedListItemsResponseActivitiesRecipientObjectJSON contains the JSON
// metadata for the struct [UserFeedListItemsResponseActivitiesRecipientObject]
type userFeedListItemsResponseActivitiesRecipientObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedListItemsResponseActivitiesRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedListItemsResponseActivitiesRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedListItemsResponseActivitiesRecipientObject) implementsUserFeedListItemsResponseActivitiesRecipient() {
}

// A recipient, which is either a user or an object
type UserFeedListItemsResponseActor struct {
	ID          string                             `json:"id,required"`
	Typename    string                             `json:"__typename,required"`
	UpdatedAt   time.Time                          `json:"updated_at,required" format:"date-time"`
	Avatar      string                             `json:"avatar,nullable"`
	Collection  string                             `json:"collection"`
	CreatedAt   time.Time                          `json:"created_at,nullable" format:"date-time"`
	Email       string                             `json:"email,nullable"`
	Name        string                             `json:"name,nullable"`
	PhoneNumber string                             `json:"phone_number,nullable"`
	Timezone    string                             `json:"timezone,nullable"`
	JSON        userFeedListItemsResponseActorJSON `json:"-"`
	union       UserFeedListItemsResponseActorsUnion
}

// userFeedListItemsResponseActorJSON contains the JSON metadata for the struct
// [UserFeedListItemsResponseActor]
type userFeedListItemsResponseActorJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	UpdatedAt   apijson.Field
	Avatar      apijson.Field
	Collection  apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userFeedListItemsResponseActorJSON) RawJSON() string {
	return r.raw
}

func (r *UserFeedListItemsResponseActor) UnmarshalJSON(data []byte) (err error) {
	*r = UserFeedListItemsResponseActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserFeedListItemsResponseActorsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [UserFeedListItemsResponseActorsObject].
func (r UserFeedListItemsResponseActor) AsUnion() UserFeedListItemsResponseActorsUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [UserFeedListItemsResponseActorsObject].
type UserFeedListItemsResponseActorsUnion interface {
	implementsUserFeedListItemsResponseActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFeedListItemsResponseActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedListItemsResponseActorsObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type UserFeedListItemsResponseActorsObject struct {
	ID          string                                    `json:"id,required"`
	Typename    string                                    `json:"__typename,required"`
	Collection  string                                    `json:"collection,required"`
	UpdatedAt   time.Time                                 `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                 `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
	JSON        userFeedListItemsResponseActorsObjectJSON `json:"-"`
}

// userFeedListItemsResponseActorsObjectJSON contains the JSON metadata for the
// struct [UserFeedListItemsResponseActorsObject]
type userFeedListItemsResponseActorsObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedListItemsResponseActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedListItemsResponseActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedListItemsResponseActorsObject) implementsUserFeedListItemsResponseActor() {}

// A content (text or markdown) block in a message in an app feed
type UserFeedListItemsResponseBlock struct {
	Name string                              `json:"name,required"`
	Type UserFeedListItemsResponseBlocksType `json:"type,required"`
	// This field can have the runtime type of
	// [[]UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButton].
	Buttons  interface{}                        `json:"buttons"`
	Content  string                             `json:"content"`
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

// A content (text or markdown) block in a message in an app feed
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

// A content (text or markdown) block in a message in an app feed
type UserFeedListItemsResponseBlocksMessageInAppFeedContentBlock struct {
	Content  string                                                          `json:"content,required"`
	Name     string                                                          `json:"name,required"`
	Rendered string                                                          `json:"rendered,required"`
	Type     UserFeedListItemsResponseBlocksMessageInAppFeedContentBlockType `json:"type,required"`
	JSON     userFeedListItemsResponseBlocksMessageInAppFeedContentBlockJSON `json:"-"`
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

// A set of buttons in a message in an app feed
type UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlock struct {
	Buttons []UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButton `json:"buttons,required"`
	Name    string                                                                `json:"name,required"`
	Type    UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockType     `json:"type,required"`
	JSON    userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockJSON     `json:"-"`
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

// A button in a set of buttons
type UserFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButton struct {
	Action string                                                                  `json:"action,required"`
	Label  string                                                                  `json:"label,required"`
	Name   string                                                                  `json:"name,required"`
	JSON   userFeedListItemsResponseBlocksMessageInAppFeedButtonSetBlockButtonJSON `json:"-"`
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

type UserFeedListItemsResponseSource struct {
	Typename   string                              `json:"__typename,required"`
	Categories []string                            `json:"categories,required"`
	Key        string                              `json:"key,required"`
	VersionID  string                              `json:"version_id,required" format:"uuid"`
	JSON       userFeedListItemsResponseSourceJSON `json:"-"`
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
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The archived status of the feed items to return
	Archived param.Field[UserFeedListItemsParamsArchived] `query:"archived"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// Whether the feed items have a tenant
	HasTenant param.Field[bool] `query:"has_tenant"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// The source of the feed items to return
	Source param.Field[string] `query:"source"`
	// The status of the feed items to return
	Status param.Field[UserFeedListItemsParamsStatus] `query:"status"`
	// The tenant of the feed items to return
	Tenant param.Field[string] `query:"tenant"`
	// The trigger data of the feed items to return (as a JSON string)
	TriggerData param.Field[string] `query:"trigger_data"`
	// The workflow categories of the feed items to return
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

// The archived status of the feed items to return
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

// The status of the feed items to return
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

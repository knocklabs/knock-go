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

// Get a user's feed of in-app messages
func (r *UserFeedService) Get(ctx context.Context, userID string, id string, query UserFeedGetParams, opts ...option.RequestOption) (res *UserFeedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/users/%s/feeds/%s", userID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a user's in-app feed settings
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

// The response for the user's feed
type UserFeedGetResponse struct {
	Entries  []UserFeedGetResponseEntry  `json:"entries,required"`
	Meta     UserFeedGetResponseMeta     `json:"meta,required"`
	PageInfo UserFeedGetResponsePageInfo `json:"page_info,required"`
	Vars     map[string]interface{}      `json:"vars,required"`
	JSON     userFeedGetResponseJSON     `json:"-"`
}

// userFeedGetResponseJSON contains the JSON metadata for the struct
// [UserFeedGetResponse]
type userFeedGetResponseJSON struct {
	Entries     apijson.Field
	Meta        apijson.Field
	PageInfo    apijson.Field
	Vars        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseJSON) RawJSON() string {
	return r.raw
}

// An in-app feed message in a user's feed
type UserFeedGetResponseEntry struct {
	ID              string                               `json:"id,required"`
	Typename        string                               `json:"__typename,required"`
	Activities      []UserFeedGetResponseEntriesActivity `json:"activities,required"`
	Actors          []UserFeedGetResponseEntriesActor    `json:"actors,required"`
	Blocks          []UserFeedGetResponseEntriesBlock    `json:"blocks,required"`
	Data            map[string]interface{}               `json:"data,required,nullable"`
	InsertedAt      string                               `json:"inserted_at,required"`
	Source          UserFeedGetResponseEntriesSource     `json:"source,required"`
	Tenant          string                               `json:"tenant,required,nullable"`
	TotalActivities int64                                `json:"total_activities,required"`
	TotalActors     int64                                `json:"total_actors,required"`
	UpdatedAt       string                               `json:"updated_at,required"`
	ArchivedAt      string                               `json:"archived_at,nullable"`
	ClickedAt       string                               `json:"clicked_at,nullable"`
	InteractedAt    string                               `json:"interacted_at,nullable"`
	LinkClickedAt   string                               `json:"link_clicked_at,nullable"`
	ReadAt          string                               `json:"read_at,nullable"`
	SeenAt          string                               `json:"seen_at,nullable"`
	JSON            userFeedGetResponseEntryJSON         `json:"-"`
}

// userFeedGetResponseEntryJSON contains the JSON metadata for the struct
// [UserFeedGetResponseEntry]
type userFeedGetResponseEntryJSON struct {
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

func (r *UserFeedGetResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseEntryJSON) RawJSON() string {
	return r.raw
}

// An activity associated with a workflow run
type UserFeedGetResponseEntriesActivity struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
	// A recipient, which is either a user or an object
	Actor UserFeedGetResponseEntriesActivitiesActor `json:"actor,nullable"`
	// The data associated with the activity
	Data       map[string]interface{} `json:"data,nullable"`
	InsertedAt time.Time              `json:"inserted_at" format:"date-time"`
	// A recipient, which is either a user or an object
	Recipient UserFeedGetResponseEntriesActivitiesRecipient `json:"recipient"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      userFeedGetResponseEntriesActivityJSON        `json:"-"`
}

// userFeedGetResponseEntriesActivityJSON contains the JSON metadata for the struct
// [UserFeedGetResponseEntriesActivity]
type userFeedGetResponseEntriesActivityJSON struct {
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

func (r *UserFeedGetResponseEntriesActivity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseEntriesActivityJSON) RawJSON() string {
	return r.raw
}

// A recipient, which is either a user or an object
type UserFeedGetResponseEntriesActivitiesActor struct {
	ID          string                                        `json:"id,required"`
	Typename    string                                        `json:"__typename,required"`
	UpdatedAt   time.Time                                     `json:"updated_at,required" format:"date-time"`
	Avatar      string                                        `json:"avatar,nullable"`
	Collection  string                                        `json:"collection"`
	CreatedAt   time.Time                                     `json:"created_at,nullable" format:"date-time"`
	Email       string                                        `json:"email,nullable"`
	Name        string                                        `json:"name,nullable"`
	PhoneNumber string                                        `json:"phone_number,nullable"`
	Timezone    string                                        `json:"timezone,nullable"`
	JSON        userFeedGetResponseEntriesActivitiesActorJSON `json:"-"`
	union       UserFeedGetResponseEntriesActivitiesActorUnion
}

// userFeedGetResponseEntriesActivitiesActorJSON contains the JSON metadata for the
// struct [UserFeedGetResponseEntriesActivitiesActor]
type userFeedGetResponseEntriesActivitiesActorJSON struct {
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

func (r userFeedGetResponseEntriesActivitiesActorJSON) RawJSON() string {
	return r.raw
}

func (r *UserFeedGetResponseEntriesActivitiesActor) UnmarshalJSON(data []byte) (err error) {
	*r = UserFeedGetResponseEntriesActivitiesActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserFeedGetResponseEntriesActivitiesActorUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [UserFeedGetResponseEntriesActivitiesActorObject].
func (r UserFeedGetResponseEntriesActivitiesActor) AsUnion() UserFeedGetResponseEntriesActivitiesActorUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [UserFeedGetResponseEntriesActivitiesActorObject].
type UserFeedGetResponseEntriesActivitiesActorUnion interface {
	implementsUserFeedGetResponseEntriesActivitiesActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFeedGetResponseEntriesActivitiesActorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedGetResponseEntriesActivitiesActorObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type UserFeedGetResponseEntriesActivitiesActorObject struct {
	ID          string                                              `json:"id,required"`
	Typename    string                                              `json:"__typename,required"`
	Collection  string                                              `json:"collection,required"`
	UpdatedAt   time.Time                                           `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                           `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                              `json:"-,extras"`
	JSON        userFeedGetResponseEntriesActivitiesActorObjectJSON `json:"-"`
}

// userFeedGetResponseEntriesActivitiesActorObjectJSON contains the JSON metadata
// for the struct [UserFeedGetResponseEntriesActivitiesActorObject]
type userFeedGetResponseEntriesActivitiesActorObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetResponseEntriesActivitiesActorObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseEntriesActivitiesActorObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedGetResponseEntriesActivitiesActorObject) implementsUserFeedGetResponseEntriesActivitiesActor() {
}

// A recipient, which is either a user or an object
type UserFeedGetResponseEntriesActivitiesRecipient struct {
	ID          string                                            `json:"id,required"`
	Typename    string                                            `json:"__typename,required"`
	UpdatedAt   time.Time                                         `json:"updated_at,required" format:"date-time"`
	Avatar      string                                            `json:"avatar,nullable"`
	Collection  string                                            `json:"collection"`
	CreatedAt   time.Time                                         `json:"created_at,nullable" format:"date-time"`
	Email       string                                            `json:"email,nullable"`
	Name        string                                            `json:"name,nullable"`
	PhoneNumber string                                            `json:"phone_number,nullable"`
	Timezone    string                                            `json:"timezone,nullable"`
	JSON        userFeedGetResponseEntriesActivitiesRecipientJSON `json:"-"`
	union       UserFeedGetResponseEntriesActivitiesRecipientUnion
}

// userFeedGetResponseEntriesActivitiesRecipientJSON contains the JSON metadata for
// the struct [UserFeedGetResponseEntriesActivitiesRecipient]
type userFeedGetResponseEntriesActivitiesRecipientJSON struct {
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

func (r userFeedGetResponseEntriesActivitiesRecipientJSON) RawJSON() string {
	return r.raw
}

func (r *UserFeedGetResponseEntriesActivitiesRecipient) UnmarshalJSON(data []byte) (err error) {
	*r = UserFeedGetResponseEntriesActivitiesRecipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserFeedGetResponseEntriesActivitiesRecipientUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [UserFeedGetResponseEntriesActivitiesRecipientObject].
func (r UserFeedGetResponseEntriesActivitiesRecipient) AsUnion() UserFeedGetResponseEntriesActivitiesRecipientUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or
// [UserFeedGetResponseEntriesActivitiesRecipientObject].
type UserFeedGetResponseEntriesActivitiesRecipientUnion interface {
	implementsUserFeedGetResponseEntriesActivitiesRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFeedGetResponseEntriesActivitiesRecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedGetResponseEntriesActivitiesRecipientObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type UserFeedGetResponseEntriesActivitiesRecipientObject struct {
	ID          string                                                  `json:"id,required"`
	Typename    string                                                  `json:"__typename,required"`
	Collection  string                                                  `json:"collection,required"`
	UpdatedAt   time.Time                                               `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                               `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                                  `json:"-,extras"`
	JSON        userFeedGetResponseEntriesActivitiesRecipientObjectJSON `json:"-"`
}

// userFeedGetResponseEntriesActivitiesRecipientObjectJSON contains the JSON
// metadata for the struct [UserFeedGetResponseEntriesActivitiesRecipientObject]
type userFeedGetResponseEntriesActivitiesRecipientObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetResponseEntriesActivitiesRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseEntriesActivitiesRecipientObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedGetResponseEntriesActivitiesRecipientObject) implementsUserFeedGetResponseEntriesActivitiesRecipient() {
}

// A recipient, which is either a user or an object
type UserFeedGetResponseEntriesActor struct {
	ID          string                              `json:"id,required"`
	Typename    string                              `json:"__typename,required"`
	UpdatedAt   time.Time                           `json:"updated_at,required" format:"date-time"`
	Avatar      string                              `json:"avatar,nullable"`
	Collection  string                              `json:"collection"`
	CreatedAt   time.Time                           `json:"created_at,nullable" format:"date-time"`
	Email       string                              `json:"email,nullable"`
	Name        string                              `json:"name,nullable"`
	PhoneNumber string                              `json:"phone_number,nullable"`
	Timezone    string                              `json:"timezone,nullable"`
	JSON        userFeedGetResponseEntriesActorJSON `json:"-"`
	union       UserFeedGetResponseEntriesActorsUnion
}

// userFeedGetResponseEntriesActorJSON contains the JSON metadata for the struct
// [UserFeedGetResponseEntriesActor]
type userFeedGetResponseEntriesActorJSON struct {
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

func (r userFeedGetResponseEntriesActorJSON) RawJSON() string {
	return r.raw
}

func (r *UserFeedGetResponseEntriesActor) UnmarshalJSON(data []byte) (err error) {
	*r = UserFeedGetResponseEntriesActor{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserFeedGetResponseEntriesActorsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [User],
// [UserFeedGetResponseEntriesActorsObject].
func (r UserFeedGetResponseEntriesActor) AsUnion() UserFeedGetResponseEntriesActorsUnion {
	return r.union
}

// A recipient, which is either a user or an object
//
// Union satisfied by [User] or [UserFeedGetResponseEntriesActorsObject].
type UserFeedGetResponseEntriesActorsUnion interface {
	implementsUserFeedGetResponseEntriesActor()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFeedGetResponseEntriesActorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedGetResponseEntriesActorsObject{}),
		},
	)
}

// A custom-object entity which belongs to a collection.
type UserFeedGetResponseEntriesActorsObject struct {
	ID          string                                     `json:"id,required"`
	Typename    string                                     `json:"__typename,required"`
	Collection  string                                     `json:"collection,required"`
	UpdatedAt   time.Time                                  `json:"updated_at,required" format:"date-time"`
	CreatedAt   time.Time                                  `json:"created_at,nullable" format:"date-time"`
	ExtraFields map[string]interface{}                     `json:"-,extras"`
	JSON        userFeedGetResponseEntriesActorsObjectJSON `json:"-"`
}

// userFeedGetResponseEntriesActorsObjectJSON contains the JSON metadata for the
// struct [UserFeedGetResponseEntriesActorsObject]
type userFeedGetResponseEntriesActorsObjectJSON struct {
	ID          apijson.Field
	Typename    apijson.Field
	Collection  apijson.Field
	UpdatedAt   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetResponseEntriesActorsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseEntriesActorsObjectJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedGetResponseEntriesActorsObject) implementsUserFeedGetResponseEntriesActor() {}

// A content (text or markdown) block in a message in an app feed
type UserFeedGetResponseEntriesBlock struct {
	Name string                               `json:"name,required"`
	Type UserFeedGetResponseEntriesBlocksType `json:"type,required"`
	// This field can have the runtime type of
	// [[]UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockButton].
	Buttons  interface{}                         `json:"buttons"`
	Content  string                              `json:"content"`
	Rendered string                              `json:"rendered"`
	JSON     userFeedGetResponseEntriesBlockJSON `json:"-"`
	union    UserFeedGetResponseEntriesBlocksUnion
}

// userFeedGetResponseEntriesBlockJSON contains the JSON metadata for the struct
// [UserFeedGetResponseEntriesBlock]
type userFeedGetResponseEntriesBlockJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Buttons     apijson.Field
	Content     apijson.Field
	Rendered    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userFeedGetResponseEntriesBlockJSON) RawJSON() string {
	return r.raw
}

func (r *UserFeedGetResponseEntriesBlock) UnmarshalJSON(data []byte) (err error) {
	*r = UserFeedGetResponseEntriesBlock{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserFeedGetResponseEntriesBlocksUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlock],
// [UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlock].
func (r UserFeedGetResponseEntriesBlock) AsUnion() UserFeedGetResponseEntriesBlocksUnion {
	return r.union
}

// A content (text or markdown) block in a message in an app feed
//
// Union satisfied by
// [UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlock] or
// [UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlock].
type UserFeedGetResponseEntriesBlocksUnion interface {
	implementsUserFeedGetResponseEntriesBlock()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserFeedGetResponseEntriesBlocksUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlock{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlock{}),
		},
	)
}

// A content (text or markdown) block in a message in an app feed
type UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlock struct {
	Content  string                                                           `json:"content,required"`
	Name     string                                                           `json:"name,required"`
	Rendered string                                                           `json:"rendered,required"`
	Type     UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockType `json:"type,required"`
	JSON     userFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockJSON `json:"-"`
}

// userFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockJSON contains the
// JSON metadata for the struct
// [UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlock]
type userFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Rendered    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlock) implementsUserFeedGetResponseEntriesBlock() {
}

type UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockType string

const (
	UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockTypeMarkdown UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockType = "markdown"
	UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockTypeText     UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockType = "text"
)

func (r UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockType) IsKnown() bool {
	switch r {
	case UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockTypeMarkdown, UserFeedGetResponseEntriesBlocksMessageInAppFeedContentBlockTypeText:
		return true
	}
	return false
}

// A set of buttons in a message in an app feed
type UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlock struct {
	Buttons []UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockButton `json:"buttons,required"`
	Name    string                                                                 `json:"name,required"`
	Type    UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockType     `json:"type,required"`
	JSON    userFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockJSON     `json:"-"`
}

// userFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockJSON contains the
// JSON metadata for the struct
// [UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlock]
type userFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockJSON struct {
	Buttons     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockJSON) RawJSON() string {
	return r.raw
}

func (r UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlock) implementsUserFeedGetResponseEntriesBlock() {
}

// A button in a set of buttons
type UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockButton struct {
	Action string                                                                   `json:"action,required"`
	Label  string                                                                   `json:"label,required"`
	Name   string                                                                   `json:"name,required"`
	JSON   userFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockButtonJSON `json:"-"`
}

// userFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockButtonJSON
// contains the JSON metadata for the struct
// [UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockButton]
type userFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockButtonJSON struct {
	Action      apijson.Field
	Label       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockButton) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockButtonJSON) RawJSON() string {
	return r.raw
}

type UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockType string

const (
	UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockTypeButtonSet UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockType = "button_set"
)

func (r UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockType) IsKnown() bool {
	switch r {
	case UserFeedGetResponseEntriesBlocksMessageInAppFeedButtonSetBlockTypeButtonSet:
		return true
	}
	return false
}

type UserFeedGetResponseEntriesBlocksType string

const (
	UserFeedGetResponseEntriesBlocksTypeMarkdown  UserFeedGetResponseEntriesBlocksType = "markdown"
	UserFeedGetResponseEntriesBlocksTypeText      UserFeedGetResponseEntriesBlocksType = "text"
	UserFeedGetResponseEntriesBlocksTypeButtonSet UserFeedGetResponseEntriesBlocksType = "button_set"
)

func (r UserFeedGetResponseEntriesBlocksType) IsKnown() bool {
	switch r {
	case UserFeedGetResponseEntriesBlocksTypeMarkdown, UserFeedGetResponseEntriesBlocksTypeText, UserFeedGetResponseEntriesBlocksTypeButtonSet:
		return true
	}
	return false
}

type UserFeedGetResponseEntriesSource struct {
	Typename   string                               `json:"__typename,required"`
	Categories []string                             `json:"categories,required"`
	Key        string                               `json:"key,required"`
	VersionID  string                               `json:"version_id,required" format:"uuid"`
	JSON       userFeedGetResponseEntriesSourceJSON `json:"-"`
}

// userFeedGetResponseEntriesSourceJSON contains the JSON metadata for the struct
// [UserFeedGetResponseEntriesSource]
type userFeedGetResponseEntriesSourceJSON struct {
	Typename    apijson.Field
	Categories  apijson.Field
	Key         apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetResponseEntriesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseEntriesSourceJSON) RawJSON() string {
	return r.raw
}

type UserFeedGetResponseMeta struct {
	Typename    string                      `json:"__typename,required"`
	TotalCount  int64                       `json:"total_count,required"`
	UnreadCount int64                       `json:"unread_count,required"`
	UnseenCount int64                       `json:"unseen_count,required"`
	JSON        userFeedGetResponseMetaJSON `json:"-"`
}

// userFeedGetResponseMetaJSON contains the JSON metadata for the struct
// [UserFeedGetResponseMeta]
type userFeedGetResponseMetaJSON struct {
	Typename    apijson.Field
	TotalCount  apijson.Field
	UnreadCount apijson.Field
	UnseenCount apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserFeedGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type UserFeedGetResponsePageInfo struct {
	HasNextPage     bool                            `json:"has_next_page,required"`
	HasPreviousPage bool                            `json:"has_previous_page,required"`
	TotalCount      int64                           `json:"total_count,required"`
	JSON            userFeedGetResponsePageInfoJSON `json:"-"`
}

// userFeedGetResponsePageInfoJSON contains the JSON metadata for the struct
// [UserFeedGetResponsePageInfo]
type userFeedGetResponsePageInfoJSON struct {
	HasNextPage     apijson.Field
	HasPreviousPage apijson.Field
	TotalCount      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserFeedGetResponsePageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userFeedGetResponsePageInfoJSON) RawJSON() string {
	return r.raw
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

type UserFeedGetParams struct {
	// The cursor to fetch entries after
	After param.Field[string] `query:"after"`
	// The archived status of the feed items to return
	Archived param.Field[UserFeedGetParamsArchived] `query:"archived"`
	// The cursor to fetch entries before
	Before param.Field[string] `query:"before"`
	// Whether the feed items have a tenant
	HasTenant param.Field[bool] `query:"has_tenant"`
	// The page size to fetch
	PageSize param.Field[int64] `query:"page_size"`
	// The source of the feed items to return
	Source param.Field[string] `query:"source"`
	// The status of the feed items to return
	Status param.Field[UserFeedGetParamsStatus] `query:"status"`
	// The tenant of the feed items to return
	Tenant param.Field[string] `query:"tenant"`
	// The trigger data of the feed items to return (as a JSON string)
	TriggerData param.Field[string] `query:"trigger_data"`
	// The workflow categories of the feed items to return
	WorkflowCategories param.Field[[]string] `query:"workflow_categories"`
}

// URLQuery serializes [UserFeedGetParams]'s query parameters as `url.Values`.
func (r UserFeedGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The archived status of the feed items to return
type UserFeedGetParamsArchived string

const (
	UserFeedGetParamsArchivedExclude UserFeedGetParamsArchived = "exclude"
	UserFeedGetParamsArchivedInclude UserFeedGetParamsArchived = "include"
	UserFeedGetParamsArchivedOnly    UserFeedGetParamsArchived = "only"
)

func (r UserFeedGetParamsArchived) IsKnown() bool {
	switch r {
	case UserFeedGetParamsArchivedExclude, UserFeedGetParamsArchivedInclude, UserFeedGetParamsArchivedOnly:
		return true
	}
	return false
}

// The status of the feed items to return
type UserFeedGetParamsStatus string

const (
	UserFeedGetParamsStatusUnread UserFeedGetParamsStatus = "unread"
	UserFeedGetParamsStatusRead   UserFeedGetParamsStatus = "read"
	UserFeedGetParamsStatusUnseen UserFeedGetParamsStatus = "unseen"
	UserFeedGetParamsStatusSeen   UserFeedGetParamsStatus = "seen"
	UserFeedGetParamsStatusAll    UserFeedGetParamsStatus = "all"
)

func (r UserFeedGetParamsStatus) IsKnown() bool {
	switch r {
	case UserFeedGetParamsStatusUnread, UserFeedGetParamsStatusRead, UserFeedGetParamsStatusUnseen, UserFeedGetParamsStatusSeen, UserFeedGetParamsStatusAll:
		return true
	}
	return false
}

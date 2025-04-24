// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"reflect"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
	"github.com/tidwall/gjson"
)

// RecipientService contains methods and other services that help with interacting
// with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientService] method instead.
type RecipientService struct {
	Options       []option.RequestOption
	Subscriptions *RecipientSubscriptionService
	Preferences   *RecipientPreferenceService
	ChannelData   *RecipientChannelDataService
}

// NewRecipientService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecipientService(opts ...option.RequestOption) (r *RecipientService) {
	r = &RecipientService{}
	r.Options = opts
	r.Subscriptions = NewRecipientSubscriptionService(opts...)
	r.Preferences = NewRecipientPreferenceService(opts...)
	r.ChannelData = NewRecipientChannelDataService(opts...)
	return
}

// A recipient of a notification, which is either a user or an object.
type Recipient struct {
	// The ID for the user that you set when identifying them in Knock.
	ID string `json:"id,required"`
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// The timestamp when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// URL to the user's avatar image.
	Avatar string `json:"avatar,nullable"`
	// The collection this object belongs to.
	Collection string `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// The primary email address for the user.
	Email string `json:"email,nullable"`
	// Display name of the user.
	Name string `json:"name,nullable"`
	// The [E.164](https://www.twilio.com/docs/glossary/what-e164) phone number of the
	// user (required for SMS channels).
	PhoneNumber string `json:"phone_number,nullable"`
	// The timezone of the user. Must be a valid
	// [tz database time zone string](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	// Used for
	// [recurring schedules](/concepts/schedules#scheduling-workflows-with-recurring-schedules-for-recipients).
	Timezone string        `json:"timezone,nullable"`
	JSON     recipientJSON `json:"-"`
	union    RecipientUnion
}

// recipientJSON contains the JSON metadata for the struct [Recipient]
type recipientJSON struct {
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

func (r recipientJSON) RawJSON() string {
	return r.raw
}

func (r *Recipient) UnmarshalJSON(data []byte) (err error) {
	*r = Recipient{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecipientUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [User], [Object].
func (r Recipient) AsUnion() RecipientUnion {
	return r.union
}

// A recipient of a notification, which is either a user or an object.
//
// Union satisfied by [User] or [Object].
type RecipientUnion interface {
	implementsRecipient()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecipientUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(User{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Object{}),
		},
	)
}

// A reference to a recipient, either a user identifier (string) or an object
// reference (ID, collection).
//
// Union satisfied by [shared.UnionString] or [RecipientReferenceObjectReference].
type RecipientReferenceUnion interface {
	ImplementsRecipientReferenceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecipientReferenceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecipientReferenceObjectReference{}),
		},
	)
}

// A reference to a recipient object.
type RecipientReferenceObjectReference struct {
	// An identifier for the recipient object.
	ID string `json:"id"`
	// The collection the recipient object belongs to.
	Collection string                                `json:"collection"`
	JSON       recipientReferenceObjectReferenceJSON `json:"-"`
}

// recipientReferenceObjectReferenceJSON contains the JSON metadata for the struct
// [RecipientReferenceObjectReference]
type recipientReferenceObjectReferenceJSON struct {
	ID          apijson.Field
	Collection  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecipientReferenceObjectReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recipientReferenceObjectReferenceJSON) RawJSON() string {
	return r.raw
}

func (r RecipientReferenceObjectReference) ImplementsRecipientReferenceUnion() {}

// A reference to a recipient, either a user identifier (string) or an object
// reference (ID, collection).
//
// Satisfied by [shared.UnionString], [RecipientReferenceObjectReferenceParam].
type RecipientReferenceUnionParam interface {
	ImplementsRecipientReferenceUnionParam()
}

// A reference to a recipient object.
type RecipientReferenceObjectReferenceParam struct {
	// An identifier for the recipient object.
	ID param.Field[string] `json:"id"`
	// The collection the recipient object belongs to.
	Collection param.Field[string] `json:"collection"`
}

func (r RecipientReferenceObjectReferenceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecipientReferenceObjectReferenceParam) ImplementsRecipientReferenceUnionParam() {}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
type RecipientRequestParam struct {
	// The ID for the user that you set when identifying them in Knock.
	ID param.Field[string] `json:"id,required"`
	// A request to set channel data for a type of channel inline.
	ChannelData param.Field[InlineChannelDataRequestParam] `json:"channel_data"`
	// The collection this object belongs to.
	Collection param.Field[string] `json:"collection"`
	// The creation date of the user from your system.
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Inline set preferences for a recipient.
	Preferences param.Field[InlinePreferenceSetRequestParam] `json:"preferences"`
}

func (r RecipientRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecipientRequestParam) ImplementsRecipientRequestUnionParam() {}

// Specifies a recipient in a request. This can either be a user identifier
// (string), an inline user request (object), or an inline object request, which is
// determined by the presence of a `collection` property.
//
// Satisfied by [shared.UnionString], [InlineIdentifyUserRequestParam],
// [InlineObjectRequestParam], [RecipientRequestParam].
type RecipientRequestUnionParam interface {
	ImplementsRecipientRequestUnionParam()
}

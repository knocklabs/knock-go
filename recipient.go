// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"reflect"
	"time"

	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
	"github.com/stainless-sdks/knock-go/option"
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
	// The email address of the user.
	Email string `json:"email,nullable"`
	// Display name of the user.
	Name string `json:"name,nullable"`
	// Phone number of the user.
	PhoneNumber string `json:"phone_number,nullable"`
	// Timezone of the user.
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
	// Inline set preferences for a recipient, where the key is the preference set name
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

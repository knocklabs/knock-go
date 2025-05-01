// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"time"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/option"
)

// RecipientSubscriptionService contains methods and other services that help with
// interacting with the knock API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientSubscriptionService] method instead.
type RecipientSubscriptionService struct {
	Options []option.RequestOption
}

// NewRecipientSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRecipientSubscriptionService(opts ...option.RequestOption) (r *RecipientSubscriptionService) {
	r = &RecipientSubscriptionService{}
	r.Options = opts
	return
}

// A subscription object.
type Subscription struct {
	// The typename of the schema.
	Typename string `json:"__typename,required"`
	// Timestamp when the resource was created.
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// A custom [Object](/concepts/objects) entity which belongs to a collection.
	Object Object `json:"object,required"`
	// A recipient of a notification, which is either a user or an object.
	Recipient Recipient `json:"recipient,required"`
	// The timestamp when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The custom properties associated with the subscription relationship.
	Properties map[string]interface{} `json:"properties,nullable"`
	JSON       subscriptionJSON       `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	Typename    apijson.Field
	InsertedAt  apijson.Field
	Object      apijson.Field
	Recipient   apijson.Field
	UpdatedAt   apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

package knock

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

func TestUsers_Get(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"User","created_at":null,"id":"user-123","updated_at":"2022-05-15T22:19:44.915Z","email":"jhammond@ingen.net","middle-name":"alfred","name":"John Hammond","phone_number":"+11234567890","welcome":"to jurassic park"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	id := "user-123"
	email := "jhammond@ingen.net"
	name := "John Hammond"
	phone_number := "+11234567890"

	user, err := client.Users.Get(ctx, &GetUserRequest{
		ID: id,
	})

	want := &User{
		Name:        name,
		ID:          id,
		Email:       email,
		PhoneNumber: phone_number,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Date(2022, time.May, 15, 22, 19, 44, 915000000, time.UTC),
		CustomProperties: map[string]interface{}{
			"welcome":     "to jurassic park",
			"middle-name": "alfred",
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(user, qt.DeepEquals, want)
}

func TestUsers_Identify(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"User","created_at":null,"id":"user-123","updated_at":"2022-05-15T22:19:44.915Z","email":"jhammond@ingen.net","middle-name":"alfred","name":"John Hammond","phone_number":"+11234567890","welcome":"to jurassic park"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	id := "user-123"
	email := "jhammond@ingen.net"
	name := "John Hammond"
	phone_number := "+11234567890"

	user, err := client.Users.Identify(ctx, &IdentifyUserRequest{
		ID: id,
		CustomProperties: map[string]interface{}{
			"middle-name": "alfred",
		}})

	want := &User{
		Name:        name,
		ID:          id,
		Email:       email,
		PhoneNumber: phone_number,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Date(2022, time.May, 15, 22, 19, 44, 915000000, time.UTC),
		CustomProperties: map[string]interface{}{
			"welcome":     "to jurassic park",
			"middle-name": "alfred",
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(user, qt.DeepEquals, want)
}

func TestUsers_Merge(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"User","created_at":null,"id":"nice2","updated_at":"2022-05-18T12:03:36.128Z","cool":"cool","email":"nice2@nice2.nice2","name":"nice2","nice":"nice","phone_number":"100"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	have, err := client.Users.Merge(ctx, &MergeUserRequest{
		ID:         "nice2",
		FromUserID: "cool2",
	})

	want := &User{
		ID:          "nice2",
		Name:        "nice2",
		Email:       "nice2@nice2.nice2",
		PhoneNumber: "100",
		CustomProperties: map[string]interface{}{
			"cool": "cool",
			"nice": "nice",
		},
		UpdatedAt: time.Date(2022, time.May, 18, 12, 03, 36, 128000000, time.UTC),
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestUsers_GetMessages(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"items":[{"__cursor":"bigcursor","__typename":"Message","archived_at":null,"channel_id":"5da042d7-02ee-46ed-8b91-9b5717da2028","data":{"middle-name":"alfred","welcome":"to jurassic park"},"id":"29GmBF0R3ZG5ltjw06vyc3H9mDa","inserted_at":"2022-05-17T00:34:18.277163Z","read_at":null,"recipient":"tom","seen_at":null,"source":{"__typename":"NotificationSource","key":"test","version_id":"4dae021a-ba51-473f-9038-77041da8131c"},"status":"delivered","tenant":null,"updated_at":"2022-05-17T00:34:18.318283Z","workflow":"test"}],"page_info":{"__typename":"PageInfo","after":"aftercursor","before":null,"page_size":1}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))

	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient()
	have, err := client.Users.GetMessages(ctx, &GetUserMessagesRequest{
		ID:       "tom",
		PageSize: 1,
	})

	insertedAt, _ := time.Parse(time.RFC3339, "2022-05-17T00:34:18.277163Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2022-05-17T00:34:18.318283Z")
	want := &GetUserMessagesResponse{
		Messages: []*Message{
			{
				Cursor:     "bigcursor",
				ID:         "29GmBF0R3ZG5ltjw06vyc3H9mDa",
				ChannelID:  "5da042d7-02ee-46ed-8b91-9b5717da2028",
				Recipient:  "tom",
				Workflow:   "test",
				Status:     "delivered",
				InsertedAt: insertedAt,
				UpdatedAt:  updatedAt,
				Data: map[string]interface{}{
					"welcome":     "to jurassic park",
					"middle-name": "alfred",
				},
				Source: &NotificationSource{
					Key:       "test",
					VersionID: "4dae021a-ba51-473f-9038-77041da8131c"},
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

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

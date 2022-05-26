package knock

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestObject_Set(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"Object","collection":"test-collection","created_at":null,"id":"cool-object2","properties":{"nice":"cool"},"updated_at":"2022-05-26T13:59:20.701Z"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient() //TODO remove any test client commented code

	have, err := client.Object.Set(ctx, &SetObjectRequest{
		ID:           "cool-object2",
		CollectionID: "test-collection",
		Properties: map[string]interface{}{
			"nice": "cool",
		},
	})

	fmt.Printf("%+v", have)
	want := &SetObjectResponse{
		Object: &Object{
			ObjectID:     "cool-object2",
			CollectionID: "test-collection",
			UpdatedAt:    ParseRFC3339Timestamp("2022-05-26T13:59:20.701Z"),
			Properties: map[string]interface{}{
				"nice": "cool",
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestObject_Get(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"Object","collection":"test-collection","created_at":null,"id":"cool-object2","properties":{"nice":"cool"},"updated_at":"2022-05-26T13:59:20.701Z"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient() //TODO remove any test client commented code

	have, err := client.Object.Get(ctx, &GetObjectRequest{
		ID:           "cool-object2",
		CollectionID: "test-collection",
	})

	fmt.Printf("%+v", have)
	want := &GetObjectResponse{
		Object: &Object{
			ObjectID:     "cool-object2",
			CollectionID: "test-collection",
			UpdatedAt:    ParseRFC3339Timestamp("2022-05-26T13:59:20.701Z"),
			Properties: map[string]interface{}{
				"nice": "cool",
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestObject_GetMessages(t *testing.T) {
	c := qt.New(t)

	// TODO use real api response, this is from docs
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"entries":[{"__typename":"Message","__cursor":"g3QAAAABZAACaWRtAAAAGzFzTXRJc1J2WnRZZjg2YU9ma00yUENwQzZYYw==","id":"1rjI9XBgWQ6EUA3D3Ul3VjUimOD","channel_id":"0bfd9f86-56b0-41f0-ade3-dc5cc6a69bb8","recipient":{"id":"project_1","collection":"projects"},"workflow":"merged-changes","tenant":null,"status":"delivered","read_at":null,"seen_at":null,"archived_at":null,"inserted_at":"2021-03-05T12:00:00Z","updated_at":"2021-03-05T12:00:00Z","source":{"__typename":"WorkflowSource","key":"merged-changes","version_id":"7251cd3f-0028-4d1a-9466-ee79522ba3de"},"data":{"foo":"bar"}}],"page_info":{"__typename":"PageInfo","after":null,"before":null,"page_size":50}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient() //TODO remove any test client commented code

	have, err := client.Object.GetMessages(ctx, &GetObjectMessagesRequest{
		ObjectID:     "cool-object2",
		CollectionID: "test-collection",
	})

	fmt.Printf("%+v", have)
	want := &GetObjectMessagesResponse{
		Items: []*ObjectMessage{
			{
				ID:        "1rjI9XBgWQ6EUA3D3Ul3VjUimOD",
				Cursor:    "g3QAAAABZAACaWRtAAAAGzFzTXRJc1J2WnRZZjg2YU9ma00yUENwQzZYYw==",
				Workflow:  "merged-changes",
				ChannelID: "0bfd9f86-56b0-41f0-ade3-dc5cc6a69bb8",
				Status:    "delivered",
				Recipient: map[string]interface{}{
					"collection": "projects",
					"id":         "project_1",
				},
				Data: map[string]interface{}{
					"foo": "bar",
				},
				Source: &NotificationSource{
					Key:       "merged-changes",
					VersionID: "7251cd3f-0028-4d1a-9466-ee79522ba3de",
				},
				InsertedAt: ParseRFC3339Timestamp("2021-03-05T12:00:00Z"),
				UpdatedAt:  ParseRFC3339Timestamp("2021-03-05T12:00:00Z"),
			},
		},
		PageInfo: PageInfo{
			PageSize: 50,
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestObject_Delete(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	err = client.Object.Delete(ctx, &GetObjectRequest{
		ID:           "cool-object2",
		CollectionID: "test-collection",
	})

	c.Assert(err, qt.IsNil)
}

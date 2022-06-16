package knock

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestObjects_Set(t *testing.T) {
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

	have, err := client.Objects.Set(ctx, &SetObjectRequest{
		ID:           "cool-object2",
		CollectionID: "test-collection",
		Properties: map[string]interface{}{
			"nice": "cool",
		},
	})

	fmt.Printf("%+v", have)
	want := &Object{
		ObjectID:     "cool-object2",
		CollectionID: "test-collection",
		UpdatedAt:    ParseRFC3339Timestamp("2022-05-26T13:59:20.701Z"),
		Properties: map[string]interface{}{
			"nice": "cool",
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestObjects_Get(t *testing.T) {
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

	have, err := client.Objects.Get(ctx, &GetObjectRequest{
		ID:           "cool-object2",
		CollectionID: "test-collection",
	})

	want := &Object{
		ObjectID:     "cool-object2",
		CollectionID: "test-collection",
		UpdatedAt:    ParseRFC3339Timestamp("2022-05-26T13:59:20.701Z"),
		Properties: map[string]interface{}{
			"nice": "cool",
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestObjects_GetMessages(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"entries":[{"__typename":"Message","__cursor":"g3QAAAABZAACaWRtAAAAGzFzTXRJc1J2WnRZZjg2YU9ma00yUENwQzZYYw==","id":"1rjI9XBgWQ6EUA3D3Ul3VjUimOD","channel_id":"0bfd9f86-56b0-41f0-ade3-dc5cc6a69bb8","recipient":{"id":"project_1","collection":"projects"},"workflow":"merged-changes","tenant":null,"status":"delivered","read_at":null,"seen_at":null,"archived_at":null,"inserted_at":"2021-03-05T12:00:00Z","updated_at":"2021-03-05T12:00:00Z","source":{"__typename":"WorkflowSource","key":"merged-changes","version_id":"7251cd3f-0028-4d1a-9466-ee79522ba3de"},"data":{"foo":"bar"}}],"page_info":{"__typename":"PageInfo","after":null,"before":null,"page_size":50}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	haveMessages, _, err := client.Objects.GetMessages(ctx, &GetObjectMessagesRequest{
		ObjectID:     "cool-object2",
		CollectionID: "test-collection",
	})

	want := []*ObjectMessage{
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
	}

	c.Assert(err, qt.IsNil)
	c.Assert(haveMessages, qt.DeepEquals, want)
}

func TestObjects_Delete(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	err = client.Objects.Delete(ctx, &GetObjectRequest{
		ID:           "cool-object2",
		CollectionID: "test-collection",
	})

	c.Assert(err, qt.IsNil)
}

func TestObjects_GetChannelData(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"ChannelData","channel_id":"2b70de44-24e8-40fe-9d54-d16d54e22374","data":{"connections":[{"incoming_webhook":{"url":"cool2"}}]}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	getObjectChannelDataResponse, err := client.Objects.GetChannelData(ctx, &GetObjectChannelDataRequest{
		Collection: "test-collection",
		ObjectID:   "cool-object",
		ChannelID:  "2b70de44-24e8-40fe-9d54-d16d54e22374",
	})

	want := map[string]interface{}{
		"connections": []interface{}{map[string]interface{}{"incoming_webhook": map[string]interface{}{"url": string("cool2")}}},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(getObjectChannelDataResponse, qt.DeepEquals, want)
}

func TestObjects_SetChannelData(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"ChannelData","channel_id":"2b70de44-24e8-40fe-9d54-d16d54e22374","data":{"connections":[{"incoming_webhook":{"url":"cool2"}}]}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient()

	setObjectChannelDataResponse, err := client.Objects.SetChannelData(ctx, &SetObjectChannelDataRequest{
		Collection: "test-collection",
		ObjectID:   "cool-object",
		ChannelID:  "2b70de44-24e8-40fe-9d54-d16d54e22374",
		Data: map[string]interface{}{
			"connections": []interface{}{map[string]interface{}{"incoming_webhook": map[string]interface{}{"url": string("cool2")}}},
		},
	})

	want := map[string]interface{}{
		"connections": []interface{}{map[string]interface{}{"incoming_webhook": map[string]interface{}{"url": string("cool2")}}},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(setObjectChannelDataResponse, qt.DeepEquals, want)
}

func TestObjects_DeleteChannelData(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient()

	err = client.Objects.DeleteChannelData(ctx, &DeleteObjectChannelDataRequest{
		Collection: "test-collection",
		ObjectID:   "cool-object",
		ChannelID:  "2b70de44-24e8-40fe-9d54-d16d54e22374",
	})

	c.Assert(err, qt.IsNil)

}

func TestObjects_GetPreferences(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"PreferenceSet","categories":null,"channel_types":{"email":false,"in_app_feed":true},"id":"default","workflows":{"new-comment":{"channel_types":{"email":false,"in_app_feed":true}}}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient()

	have, err := client.Objects.GetPreferences(ctx, &GetObjectPreferencesRequest{
		PreferenceID: "default",
		ObjectID:     "cool-object",
		Collection:   "test-collection",
	})

	want := &PreferenceSet{
		ID: "default",
		Workflows: map[string]interface{}{
			"new-comment": map[string]interface{}{"channel_types": map[string]interface{}{"email": false, "in_app_feed": true}},
		},
		ChannelTypes: map[string]interface{}{"email": false, "in_app_feed": true},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestObjects_SetPreferences(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"PreferenceSet","categories":null,"channel_types":{"email":true,"in_app_feed":false},"id":"default","workflows":{"new-comment":{"channel_types":{"email":true,"in_app_feed":false}}}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	request := &SetObjectPreferencesRequest{
		PreferenceID: "default",
		ObjectID:     "cool-object",
		Collection:   "test-collection",
	}

	request.AddChannelTypesPreference(map[string]interface{}{
		"email":       true,
		"in_app_feed": false,
	})

	request.AddWorkflowsPreference(map[string]interface{}{
		"new-comment": map[string]interface{}{
			"channel_types": map[string]interface{}{
				"email":       true,
				"in_app_feed": false,
			},
		},
	})

	have, err := client.Objects.SetPreferences(ctx, request)

	want := &PreferenceSet{
		ID: "default",
		Workflows: map[string]interface{}{
			"new-comment": map[string]interface{}{
				"channel_types": map[string]interface{}{
					"email":       true,
					"in_app_feed": false,
				},
			},
		},
		ChannelTypes: map[string]interface{}{"email": true, "in_app_feed": false},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

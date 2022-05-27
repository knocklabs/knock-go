package knock

import (
	"context"
	"fmt"
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

	want := &GetUserResponse{
		User: &User{
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
		User: &User{
			ID: id,
			CustomProperties: map[string]interface{}{
				"middle-name": "alfred",
			},
		}})

	want := &IdentifyUserResponse{
		User: &User{
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

	want := &MergeUserResponse{
		User: &User{
			ID:          "nice2",
			Name:        "nice2",
			Email:       "nice2@nice2.nice2",
			PhoneNumber: "100",
			CustomProperties: map[string]interface{}{
				"cool": "cool",
				"nice": "nice",
			},
			UpdatedAt: time.Date(2022, time.May, 18, 12, 03, 36, 128000000, time.UTC),
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestUsers_GetMessages(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"items":[{"__cursor":"big-cursor","__typename":"Message","archived_at":null,"channel_id":"5da042d7-02ee-46ed-8b91-9b5717da2028","data":{"middle-name":"alfred","welcome":"to jurassic park"},"id":"29GmBF0R3ZG506vyc3H9mDa","inserted_at":"2022-05-17T00:34:18.277163Z","read_at":null,"recipient":"tom","seen_at":null,"source":{"__typename":"NotificationSource","key":"test","version_id":"4dae021a-ba51-473f-9038-77041da8131c"},"status":"delivered","tenant":null,"updated_at":"2022-05-17T00:34:18.318283Z","workflow":"test"}],"page_info":{"__typename":"PageInfo","after":"after-cursor","before":null,"page_size":1}}`
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

	want := &GetUserMessagesResponse{
		Messages: []*Message{
			{
				Cursor:     "big-cursor",
				ID:         "29GmBF0R3ZG506vyc3H9mDa",
				ChannelID:  "5da042d7-02ee-46ed-8b91-9b5717da2028",
				Recipient:  "tom",
				Workflow:   "test",
				Status:     "delivered",
				InsertedAt: ParseRFC3339Timestamp("2022-05-17T00:34:18.277163Z"),
				UpdatedAt:  ParseRFC3339Timestamp("2022-05-17T00:34:18.318283Z"),
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

func TestUsers_BulkIdentify(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"BulkOperation","completed_at":null,"estimated_total_rows":2,"failed_at":null,"id":"b19d032a-fc3b-4a54-9f54-369484527201","inserted_at":"2022-05-27T11:12:08.281201Z","name":"users.identify","processed_rows":0,"progress_path":"/v1/bulk_operations/b19d032a-fc3b-4a54-9f54-369484527201","started_at":null,"status":"queued","updated_at":"2022-05-27T11:12:08.286507Z"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient() //TODO remove any test client commented code

	user, err := client.Users.BulkIdentify(ctx, &BulkIdentifyUserRequest{
		Users: []*User{
			{
				Name:        "John Hammond",
				ID:          "user-123",
				Email:       "jhammond@ingen.net",
				PhoneNumber: "+11234567890",
				CustomProperties: map[string]interface{}{
					"welcome":     "to jurassic park",
					"middle-name": "alfred1",
				},
			},
			{
				Name:        "John Hammond2",
				ID:          "user-1234",
				Email:       "jhammond2@ingen.net",
				PhoneNumber: "+11234567891",
				CustomProperties: map[string]interface{}{
					"welcome":     "to jurassic park",
					"middle-name": "alfred2",
				},
			},
		},
	})

	want := &BulkIdentifyUserResponse{
		BulkOperation: &BulkOperation{
			ID:                 "b19d032a-fc3b-4a54-9f54-369484527201",
			EstimatedTotalRows: 2,
			ProgressPath:       "/v1/bulk_operations/b19d032a-fc3b-4a54-9f54-369484527201",
			Status:             BulkOperationQueued,
			InsertedAt:         ParseRFC3339Timestamp("2022-05-27T11:12:08.281201Z"),
			UpdatedAt:          ParseRFC3339Timestamp("2022-05-27T11:12:08.286507Z"),
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(user, qt.DeepEquals, want)
}

func TestUsers_BulkDelete(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"BulkOperation","completed_at":null,"estimated_total_rows":2,"failed_at":null,"id":"0fb6a596-b579-455e-9a01-32b41fa5613a","inserted_at":"2022-05-27T11:26:50.869070Z","name":"users.delete","processed_rows":0,"progress_path":"/v1/bulk_operations/0fb6a596-b579-455e-9a01-32b41fa5613a","started_at":null,"status":"queued","updated_at":"2022-05-27T11:26:50.879642Z"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	user, err := client.Users.BulkDelete(ctx, &BulkDeleteUserRequest{UserIDs: []string{
		"user-123",
		"user-1234",
	}})

	want := &BulkIdentifyUserResponse{
		BulkOperation: &BulkOperation{
			ID:                 "0fb6a596-b579-455e-9a01-32b41fa5613a",
			EstimatedTotalRows: 2,
			ProgressPath:       "/v1/bulk_operations/0fb6a596-b579-455e-9a01-32b41fa5613a",
			Status:             BulkOperationQueued,
			InsertedAt:         ParseRFC3339Timestamp("2022-05-27T11:26:50.869070Z"),
			UpdatedAt:          ParseRFC3339Timestamp("2022-05-27T11:26:50.879642Z"),
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(user, qt.DeepEquals, want)
}

func TestFeed_Get(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"entries":[{"__typename":"FeedItem","__cursor":"g3QAAAABZAACaWRtAAAAGzFzTXRJc1J2WnRZZjg2YU9ma00yUENwQzZYYw==","activities":[{"__typename":"Activity","actor":{"__typename":"User","id":"c121a5ea-8f2c-4c60-ab40-9966047d5bea","created_at":null,"updated_at":"2021-05-08T20:40:01.340Z","email":"some-user@knock.app","name":"Some User"},"data":{"dest_environment_name":"Production","src_environment_name":"Development","total_merged":1},"id":"activity-id","inserted_at":"2021-05-11T00:50:09.895759Z","recipient":{"__typename":"User","id":"c121a5ea-8f2c-4c60-ab40-9966047d5bea","created_at":null,"updated_at":"2021-05-08T20:40:01.340Z","email":"some-user@knock.app","name":"Some User"},"updated_at":"2021-05-11T00:50:09.895759Z"}],"actors":[{"__typename":"User","id":"c121a5ea-8f2c-4c60-ab40-9966047d5bea","created_at":null,"updated_at":"2021-05-08T20:40:01.340Z","email":"some-user@knock.app","name":"Some User"}],"archived_at":null,"blocks":[{"content":"**{{ actor.name }}** merged {{ total_merged }} {% if total_merged == 1 %} change {% else %} changes {% endif %}\nfrom **{{ src_environment_name }}** into **{{ dest_environment_name }}**.","name":"body","rendered":"<p><strong>The person</strong> merged 1  change \nfrom <strong>Development</strong> into <strong>Production</strong>.</p>","type":"markdown"},{"content":"{{ vars.app_url }}/{{ account_slug }}/commits","name":"action_url","rendered":"https://example.com/thing/commits","type":"text"}],"data":{"dest_environment_name":"Production","src_environment_name":"Development","total_merged":1},"id":"1sMtIsRvZtYf86aOfkM2PCpC6Xc","inserted_at":"2021-05-11T00:50:09.904531Z","read_at":"2021-05-13T02:45:28.559124Z","seen_at":"2021-05-11T00:51:43.617550Z","source":{"__typename":"WorkflowSource","key":"merged-changes","version_id":"7251cd3f-0028-4d1a-9466-ee79522ba3de"},"tenant":null,"total_activities":1,"total_actors":1,"updated_at":"2021-05-13T02:45:28.559863Z"}],"vars":{"app_name":"The app name"},"meta":{"__typename":"FeedMetadata","unread_count":10,"unseen_count":20},"page_info":{"__typename":"PageInfo","after":null,"before":null,"page_size":50}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	have, err := client.Users.GetFeed(ctx, &GetFeedRequest{
		UserID: "test-123",
		FeedID: "5d2377a0-92fb-4616-8315-eee843556566",
	})

	fmt.Printf("%+v", have.Feed)
	want := &GetFeedResponse{
		Feed: &Feed{
			FeedItems: []*FeedItem{
				{
					Activities: []*MessageActivity{
						{
							ID: "activity-id",
							Data: map[string]interface{}{
								"dest_environment_name": "Production",
								"src_environment_name":  "Development",
								"total_merged":          float64(1),
							},
							Actor: &User{
								ID:        "c121a5ea-8f2c-4c60-ab40-9966047d5bea",
								Name:      "Some User",
								Email:     "some-user@knock.app",
								UpdatedAt: ParseRFC3339Timestamp("2021-05-08T20:40:01.340Z"),
							},
							Recipient: &User{
								ID:        "c121a5ea-8f2c-4c60-ab40-9966047d5bea",
								Name:      "Some User",
								Email:     "some-user@knock.app",
								UpdatedAt: ParseRFC3339Timestamp("2021-05-08T20:40:01.34Z"),
							},
						},
					},
					Actors: []*User{
						{
							ID:        "c121a5ea-8f2c-4c60-ab40-9966047d5bea",
							Name:      "Some User",
							Email:     "some-user@knock.app",
							UpdatedAt: ParseRFC3339Timestamp("2021-05-08T20:40:01.340Z"),
						},
					},
					TotalActivities: 1,
					TotalActors:     1,
					Source:          NotificationSource{Key: "merged-changes", VersionID: "7251cd3f-0028-4d1a-9466-ee79522ba3de"},
					ReadAt:          ParseRFC3339Timestamp("2021-05-13T02:45:28.559124Z"),
					InsertedAt:      ParseRFC3339Timestamp("2021-05-11T00:50:09.904531Z"),
					UpdatedAt:       ParseRFC3339Timestamp("2021-05-13T02:45:28.559863Z"),
					SeenAt:          ParseRFC3339Timestamp("2021-05-11T00:51:43.617550Z"),
				},
			},
			FeedMetadata: &FeedMetadata{
				UnreadCount: 10,
				UnseenCount: 20,
			},
			PageInfo: &PageInfo{
				PageSize: 50,
			},
			Vars: map[string]interface{}{
				"app_name": "The app name",
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestChannelData_GetForUser(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"ChannelData","channel_id":"5d2377a0-92fb-4616-8315-eee843556566","data":{"tokens":["a","b"]}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	setUserChannelDataResponse, err := client.Users.GetChannelData(ctx, &GetUserChannelDataRequest{
		UserID:    "test-123",
		ChannelID: "5d2377a0-92fb-4616-8315-eee843556566",
	})

	var responseTokens []interface{}
	responseTokens = append(responseTokens, "a")
	responseTokens = append(responseTokens, "b")

	want := &GetUserChannelDataResponse{
		ChannelData: map[string]interface{}{
			"tokens": responseTokens,
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(setUserChannelDataResponse, qt.DeepEquals, want)
}

func TestChannelData_SetForUser(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"ChannelData","channel_id":"5d2377a0-92fb-4616-8315-eee843556566","data":{"tokens":["a","b"]}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient()

	setUserChannelDataResponse, err := client.Users.SetChannelData(ctx, &SetUserChannelDataRequest{
		UserID:    "test-123",
		ChannelID: "5d2377a0-92fb-4616-8315-eee843556566",
		Data: map[string]interface{}{
			"tokens": []string{"a", "b"},
		},
	})

	var responseTokens []interface{}
	responseTokens = append(responseTokens, "a")
	responseTokens = append(responseTokens, "b")

	want := &GetUserChannelDataResponse{
		ChannelData: map[string]interface{}{
			"tokens": responseTokens,
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(setUserChannelDataResponse, qt.DeepEquals, want)
}

func TestChannelData_DeleteForUser(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"ChannelData","channel_id":"5d2377a0-92fb-4616-8315-eee843556566","data":{"tokens":["a","b"]}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient()

	err = client.Users.DeleteChannelData(ctx, &DeleteUserChannelDataRequest{
		UserID:    "test-123",
		ChannelID: "5d2377a0-92fb-4616-8315-eee843556566",
	})

	c.Assert(err, qt.IsNil)

}

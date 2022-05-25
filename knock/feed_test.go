package knock

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
)

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

	have, err := client.Feed.Get(ctx, &GetFeedRequest{
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

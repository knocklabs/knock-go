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

func TestMessages_List(t *testing.T) {

	// RETURN here to finish this test
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(r)
		w.WriteHeader(200)
		out := `{"items":[{"__cursor":"bigcursor","__typename":"Message","archived_at":null,"channel_id":"5da042d7-02ee-46ed-8b91-9b5717da2028","data":{"middle-name":"alfred","welcome":"to jurassic park"},"id":"message-id","inserted_at":"2022-05-17T00:34:18.277163Z","read_at":null,"recipient":"tom","seen_at":null,"source":{"__typename":"NotificationSource","key":"test","version_id":"4dae021a-ba51-473f-9038-77041da8131c"},"status":"delivered","tenant":null,"updated_at":"2022-05-17T00:34:18.318283Z","workflow":"test"}],"page_info":{"__typename":"PageInfo","after":"bigafter","before":null,"page_size":1}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	have, err := client.Messages.List(ctx, &ListMessagesRequest{
		PageSize: 1,
	})

	if err != nil {
		fmt.Println(err)
	}

	want := &ListMessagesResponse{
		PageInfo: PageInfoMessages{
			PageSize: 1,
			After:    "bigafter",
		},
		Items: []*Message{
			{
				Cursor:     "bigcursor",
				ID:         "message-id",
				ChannelID:  "5da042d7-02ee-46ed-8b91-9b5717da2028",
				Recipient:  "tom",
				Workflow:   "test",
				Tenant:     "",
				Status:     "delivered",
				ReadAt:     time.Time{},
				SeenAt:     time.Time{},
				ArchivedAt: time.Time{},
				InsertedAt: time.Date(2022, time.May, 17, 00, 34, 18, 277163000, time.UTC),
				UpdatedAt:  time.Date(2022, time.May, 17, 00, 34, 18, 318283000, time.UTC),
				Source: &NotificationSource{
					Key:       "test",
					VersionID: "4dae021a-ba51-473f-9038-77041da8131c",
				},
				Data: map[string]interface{}{
					"middle-name": "alfred",
					"welcome":     "to jurassic park",
				},
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

package knock

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
)

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

	setUserChannelDataResponse, err := client.ChannelData.GetForUser(ctx, &GetUserChannelDataRequest{
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

	setUserChannelDataResponse, err := client.ChannelData.SetForUser(ctx, &SetUserChannelDataRequest{
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

	err = client.ChannelData.DeleteForUser(ctx, &DeleteUserChannelDataRequest{
		UserID:    "test-123",
		ChannelID: "5d2377a0-92fb-4616-8315-eee843556566",
	})

	c.Assert(err, qt.IsNil)

}

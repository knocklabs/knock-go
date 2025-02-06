package knock

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestBulkOperation_get(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"BulkOperation","completed_at":null,"estimated_total_rows":42,"failed_at":null,"id":"b4f6f61e-3634-4e80-af0d-9c83e9acc6f3","name":"bulk_op","processed_rows":0,"progress_path":"/v1/bulk_operations/b4f6f61e-3634-4e80-af0d-9c83e9acc6f3","started_at":null,"status":"queued","inserted_at":"2021-03-05T12:00:00Z","updated_at":"2021-03-05T12:00:00Z","error_count":0,"success_count":0,"error_items":[]}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	have, err := client.BulkOperations.Get(ctx, &GetBulkOperationRequest{
		ID: "b4f6f61e-3634-4e80-af0d-9c83e9acc6f3",
	})

	want := &BulkOperation{
		ID:                 "b4f6f61e-3634-4e80-af0d-9c83e9acc6f3",
		ProgressPath:       "/v1/bulk_operations/b4f6f61e-3634-4e80-af0d-9c83e9acc6f3",
		Status:             BulkOperationStatus(Queued),
		InsertedAt:         ParseRFC3339Timestamp("2021-03-05T12:00:00Z"),
		UpdatedAt:          ParseRFC3339Timestamp("2021-03-05T12:00:00Z"),
		EstimatedTotalRows: 42,
		Name:               "bulk_op",
		SuccessCount:       0,
		ErrorCount:         0,
		ErrorItems:         []ErrorItem{},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestBulkOperation_get_with_error(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"BulkOperation","completed_at":null,"estimated_total_rows":5,"failed_at":null,"id":"b4f6f61e-3634-4e80-af0d-9c83e9acc6f3","name":"bulk_op","processed_rows":3,"progress_path":"/v1/bulk_operations/b4f6f61e-3634-4e80-af0d-9c83e9acc6f3","started_at":null,"status":"completed","inserted_at":"2021-03-05T12:00:00Z","updated_at":"2021-03-05T12:00:00Z","error_count":2,"success_count":1,"error_items":[{"id":"1"},{"id":"2"}]}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	have, err := client.BulkOperations.Get(ctx, &GetBulkOperationRequest{
		ID: "b4f6f61e-3634-4e80-af0d-9c83e9acc6f3",
	})

	want := &BulkOperation{
		ID:                 "b4f6f61e-3634-4e80-af0d-9c83e9acc6f3",
		ProgressPath:       "/v1/bulk_operations/b4f6f61e-3634-4e80-af0d-9c83e9acc6f3",
		Status:             BulkOperationStatus(Completed),
		InsertedAt:         ParseRFC3339Timestamp("2021-03-05T12:00:00Z"),
		UpdatedAt:          ParseRFC3339Timestamp("2021-03-05T12:00:00Z"),
		EstimatedTotalRows: 5,
		ProcessedRows:      3,
		Name:               "bulk_op",
		SuccessCount:       1,
		ErrorCount:         2,
		ErrorItems: []ErrorItem{
			map[string]interface{}{
				"id": "1",
			},
			map[string]interface{}{
				"id": "2",
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

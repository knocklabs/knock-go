package knock

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestTenants_Set(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"Tenant","created_at":null,"id":"cool-tenant2","properties":{"name":"cool-tenant-1","settings":{"branding":{"primary_color":"#FFFFFF"}}},"updated_at":"2022-05-26T13:59:20.701Z"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	have, err := client.Tenants.Set(ctx, &SetTenantRequest{
		ID: "cool-tenant2",
		Properties: map[string]interface{}{
			"name": "cool-tenant-1",
			"settings": map[string]interface{}{
				"branding": map[string]interface{}{
					"primary_color": "#FFFFFF",
				},
			},
		},
	})

	want := &Tenant{
		ID:        "cool-tenant2",
		UpdatedAt: ParseRFC3339Timestamp("2022-05-26T13:59:20.701Z"),
		Properties: map[string]interface{}{
			"name": "cool-tenant-1",
			"settings": map[string]interface{}{
				"branding": map[string]interface{}{
					"primary_color": "#FFFFFF",
				},
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestTenants_Get(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"Tenant","created_at":null,"id":"cool-tenant1","properties":{"name":"cool-tenant-1","settings":{"branding":{"primary_color":"#FFFFFF"}}},"updated_at":"2022-05-26T13:59:20.701Z"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	have, err := client.Tenants.Get(ctx, &GetTenantRequest{
		ID: "cool-tenant2",
	})

	want := &Tenant{
		ID:        "cool-tenant1",
		UpdatedAt: ParseRFC3339Timestamp("2022-05-26T13:59:20.701Z"),
		Properties: map[string]interface{}{
			"name": "cool-tenant-1",
			"settings": map[string]interface{}{
				"branding": map[string]interface{}{
					"primary_color": "#FFFFFF",
				},
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestTenants_List(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"items":[{"__typename":"Tenant","properties":{"name":"cool-tenant-1","settings":{"branding":{"primary_color":"#FFFFFF"}}},"id":"tenant-id","updated_at":"2022-05-17T00:34:18.277163Z"}],"page_info":{"__typename":"PageInfo","after":"big-after","before":null,"page_size":1}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	haveTenants, havePageInfo, err := client.Tenants.List(ctx, &ListTenantsRequest{
		PageSize: 1,
	})

	wantTenants :=
		[]*Tenant{
			{
				ID:        "tenant-id",
				UpdatedAt: ParseRFC3339Timestamp("2022-05-17T00:34:18.277163Z"),
				Properties: map[string]interface{}{
					"name": "cool-tenant-1",
					"settings": map[string]interface{}{
						"branding": map[string]interface{}{
							"primary_color": "#FFFFFF",
						},
					},
				},
			},
		}

	wantPageInfo := &PageInfo{
		PageSize: 1,
		After:    "big-after",
	}

	c.Assert(err, qt.IsNil)
	c.Assert(haveTenants, qt.DeepEquals, wantTenants)
	c.Assert(havePageInfo, qt.DeepEquals, wantPageInfo)
}

func TestTenants_Delete(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	err = client.Tenants.Delete(ctx, &GetTenantRequest{
		ID: "cool-tenant2",
	})

	c.Assert(err, qt.IsNil)
}

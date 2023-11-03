package knock

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// TenantsService is an interface for communicating with the Knock
// Tenants API endpoints.
type TenantsService interface {
	List(context.Context, *ListTenantsRequest) ([]*Tenant, *PageInfo, error)
	Get(context.Context, *GetTenantRequest) (*Tenant, error)
	Set(context.Context, *SetTenantRequest) (*Tenant, error)
	Delete(context.Context, *DeleteTenantRequest) error
}
type tenantsService struct {
	client *Client
}

var _ TenantsService = &tenantsService{}

func NewTenantService(client *Client) *tenantsService {
	return &tenantsService{
		client: client,
	}
}

// Context structs
type Tenant struct {
	ID         string                 `json:"id"`
	Properties map[string]interface{} `json:"properties"`
	Settings   map[string]interface{} `json:"settings"`
	UpdatedAt  time.Time              `json:"updated_at"`
	CreatedAt  time.Time              `json:"created_at"`
}

// Client structs
type ListTenantsRequest struct {
	PageSize int    `url:"page_size,omitempty"`
	Before   string `url:"before,omitempty"`
	After    string `url:"after,omitempty"`
	TenantID string `url:"tenant_id,omitempty"`
	Name     string `url:"name,omitempty"`
}

type ListTenantsResponse struct {
	Entries  []*Tenant `json:"entries"`
	PageInfo *PageInfo `json:"page_info"`
}

type GetTenantRequest struct {
	ID string
}
type GetTenantResponse struct {
	Tenant *Tenant
}

type SetTenantRequest struct {
	ID         string                 `json:"-"`
	Properties map[string]interface{} `json:""`
	Settings   map[string]interface{} `json:""`
}

type SetTenantResponse = GetTenantResponse

type DeleteTenantRequest = GetTenantRequest

func tenantAPIPath(tenantID string) string {
	return fmt.Sprintf("v1/tenants/%s", tenantID)
}

const tenantsAPIBasePath = "/v1/tenants"

func (ts *tenantsService) List(ctx context.Context, listReq *ListTenantsRequest) ([]*Tenant, *PageInfo, error) {
	queryString, err := query.Values(listReq)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error parsing request to list tenants")
	}
	path := fmt.Sprintf("%s/?%s", tenantsAPIBasePath, queryString.Encode())

	req, err := ts.client.newRequest(http.MethodGet, path, listReq, nil)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating request to list tenants")
	}
	listRes := &ListTenantsResponse{}

	_, err = ts.client.do(ctx, req, listRes)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error making request to list tenants")
	}

	return listRes.Entries, listRes.PageInfo, nil
}
func (ts *tenantsService) Get(ctx context.Context, getTenantRequest *GetTenantRequest) (*Tenant, error) {
	path := tenantAPIPath(getTenantRequest.ID)

	req, err := ts.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get tenant")
	}

	getTenantResponse := &GetTenantResponse{Tenant: &Tenant{}}
	_, err = ts.client.do(ctx, req, getTenantResponse.Tenant)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get tenant")
	}

	return getTenantResponse.Tenant, nil
}

func (ts *tenantsService) Set(ctx context.Context, setTenantRequest *SetTenantRequest) (*Tenant, error) {
	path := tenantAPIPath(setTenantRequest.ID)

	if len(setTenantRequest.Properties) == 0 || len(setTenantRequest.Settings) == 0 {
		return nil, &Error{msg: "Must set at least one property or settings"}
	}

	// Copy the properties ready to be sent to the API
	properties := make(map[string]interface{})

	for k, v := range setTenantRequest.Properties {
		properties[k] = v
	}

	if len(setTenantRequest.Settings) > 0 {
		// For simplicity we add "settings" under a key in properties, because that's what
		// we want to send to the API anyway.
		properties["settings"] = setTenantRequest.Settings
	}

	req, err := ts.client.newRequest(http.MethodPut, path, properties, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for set tenant")
	}

	setTenantResponse := &SetTenantResponse{Tenant: &Tenant{}}
	_, err = ts.client.do(ctx, req, setTenantResponse.Tenant)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for set tenant")
	}

	return setTenantResponse.Tenant, nil
}

func (ts *tenantsService) Delete(ctx context.Context, deleteTenantRequest *DeleteTenantRequest) error {
	path := tenantAPIPath(deleteTenantRequest.ID)

	req, err := ts.client.newRequest(http.MethodDelete, path, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error creating request for delete tenant")
	}

	_, err = ts.client.do(ctx, req, nil)
	if err != nil {
		return errors.Wrap(err, "error making request for delete tenant")
	}

	return nil
}

package knock

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// BulkOperationsService is an interface for communicating with API endpoints that return BulkOperations,
// which are used by other services including the UsersService and MessagesService
type BulkOperationsService interface {
	Get(context.Context, *GetBulkOperationRequest) (*BulkOperation, error)
}

type bulkOperationsService struct {
	client *Client
}

var _ BulkOperationsService = &bulkOperationsService{}

func NewBulkOperationsService(client *Client) *bulkOperationsService {
	return &bulkOperationsService{
		client: client,
	}
}

type BulkOperationStatus string

const (
	BulkOperationQueued BulkOperationStatus = "queued"
	Processing          BulkOperationStatus = "processing"
	Completed           BulkOperationStatus = "completed"
	Failed              BulkOperationStatus = "failed"
)

type BulkOperation struct {
	ID                 string              `json:"id"`
	CompletedAt        time.Time           `json:"completed_at"`
	FailedAt           time.Time           `json:"failed_at"`
	StartedAt          time.Time           `json:"started_at"`
	InsertedAt         time.Time           `json:"inserted_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	EstimatedTotalRows int                 `json:"estimated_total_rows"`
	ProcessedRows      int                 `json:"processed_rows"`
	ProgressPath       string              `json:"progress_path"`
	Status             BulkOperationStatus `json:"status"`
}

type GetBulkOperationRequest struct {
	ID string
}

type GetBulkOperationResponse struct {
	BulkOperation *BulkOperation
}

func bulkOperationsAPIPath(id string) string {
	return fmt.Sprintf("/v1/bulk_operations/%s", id)
}

func (bos *bulkOperationsService) Get(ctx context.Context, getBulkOperationReq *GetBulkOperationRequest) (*BulkOperation, error) {

	path := bulkOperationsAPIPath(getBulkOperationReq.ID)

	req, err := bos.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request to get bulk operation")
	}

	GetBulkOperationResponse := &GetBulkOperationResponse{BulkOperation: &BulkOperation{}}
	_, err = bos.client.do(ctx, req, GetBulkOperationResponse.BulkOperation)

	if err != nil {
		return nil, errors.Wrap(err, "error making request to get bulk operation")
	}

	return GetBulkOperationResponse.BulkOperation, nil
}

package knock

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type FeedService interface {
	Get(context.Context, *GetFeedRequest) (*GetFeedResponse, error)
}

type feedService struct {
	client *Client
}

var _ FeedService = &feedService{}

func NewFeedService(client *Client) *feedService {
	return &feedService{
		client: client,
	}
}

type Feed struct {
	FeedItems    []*FeedItem            `json:"entries"`
	PageInfo     *PageInfo              `json:"page_info"`
	Vars         map[string]interface{} `json:"vars"`
	FeedMetadata *FeedMetadata          `json:"meta"`
}

type FeedItem struct {
	Activities      []*MessageActivity `json:"activities"`
	Actors          []*User            `json:"actors"`
	TotalActivities int                `json:"total_activities"`
	TotalActors     int                `json:"total_actors"`
	// Blocks TODO blocks
	Data       map[string]interface{} `json:"vars"`
	InsertedAt time.Time              `json:"inserted_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	ReadAt     time.Time              `json:"read_at"`
	SeenAt     time.Time              `json:"seen_at"`
	ArchivedAt time.Time              `json:"archived_at"`
	Source     NotificationSource     `json:"source"`
	Tenant     string                 `json:"tenant"`
}

type FeedMetadata struct {
	UnreadCount int `json:"unread_count"`
	UnseenCount int `json:"unseen_count"`
}

type GetFeedRequest struct {
	UserID string
	FeedID string
}

type GetFeedResponse struct {
	Feed *Feed
}

func feedsAPIPath(userID string, FeedID string) string {
	return fmt.Sprintf("%s/feeds/%s", UsersAPIPath(userID), FeedID)
}

func (fs *feedService) Get(ctx context.Context, getFeedReq *GetFeedRequest) (*GetFeedResponse, error) {
	path := feedsAPIPath(getFeedReq.UserID, getFeedReq.FeedID)

	req, err := fs.client.newRequest(http.MethodGet, path, getFeedReq)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get feed")
	}

	getFeedRes := &GetFeedResponse{Feed: &Feed{}}
	_, err = fs.client.do(ctx, req, getFeedRes.Feed)
	if err != nil {
		return nil, errors.Wrap(err, "error making request for get feed")
	}

	return getFeedRes, nil
}

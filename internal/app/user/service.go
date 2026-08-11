package user

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/marketing-digest/pkg/errorsx"

	"github.com/marketing-digest/auth-service/internal/app/user/models"
)

// Service is the small user/profile API (no OAuth yet).
type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) Ping(_ context.Context, message string) (string, string, error) {
	msg := strings.TrimSpace(message)
	if msg == "" {
		msg = "pong"
	}
	return msg, "auth-service", nil
}

func (s *Service) GetMe(ctx context.Context) (*models.User, error) {
	return s.store.GetOwner(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (*models.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errorsx.ErrInvalidArgument
	}
	return s.store.GetByID(ctx, id)
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	if strings.TrimSpace(username) == "" {
		return nil, errorsx.ErrInvalidArgument
	}
	return s.store.GetByUsername(ctx, username)
}

type TopPost struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Views int64  `json:"views"`
}

type AnalyticsView struct {
	UserID         string
	TotalViews     int64
	TotalPosts     int
	TotalLikes     int64
	Followers      int
	ViewsThisMonth int64
	PostsThisMonth int
	TopPosts       []TopPost
}

func (s *Service) GetAnalytics(ctx context.Context, userID string) (*AnalyticsView, error) {
	if strings.TrimSpace(userID) == "" {
		return nil, errorsx.ErrInvalidArgument
	}
	if _, err := s.store.GetByID(ctx, userID); err != nil {
		return nil, err
	}
	a, err := s.store.GetAnalytics(ctx, userID)
	if err != nil {
		return nil, err
	}
	var tops []TopPost
	_ = json.Unmarshal([]byte(a.TopPostsJSON), &tops)
	return &AnalyticsView{
		UserID: a.UserID, TotalViews: a.TotalViews, TotalPosts: a.TotalPosts,
		TotalLikes: a.TotalLikes, Followers: a.Followers,
		ViewsThisMonth: a.ViewsThisMonth, PostsThisMonth: a.PostsThisMonth,
		TopPosts: tops,
	}, nil
}

func FormatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

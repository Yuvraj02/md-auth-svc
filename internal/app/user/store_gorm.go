package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/marketing-digest/pkg/errorsx"

	"github.com/marketing-digest/auth-service/internal/app/user/models"
)

type row struct {
	ID           string    `gorm:"primaryKey"`
	Email        string    `gorm:"type:varchar(320);uniqueIndex;not null"`
	DisplayName  string    `gorm:"column:display_name;type:varchar(200);not null"`
	Name         string    `gorm:"type:varchar(200);not null;default:''"`
	Username     string    `gorm:"type:varchar(64);not null;default:''"`
	Avatar       string    `gorm:"type:text;not null;default:''"`
	Bio          string    `gorm:"type:text;not null;default:''"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

func (row) TableName() string { return "users" }

type analyticsRow struct {
	UserID         string `gorm:"column:user_id;primaryKey"`
	TotalViews     int64  `gorm:"column:total_views;not null;default:0"`
	TotalPosts     int    `gorm:"column:total_posts;not null;default:0"`
	TotalLikes     int64  `gorm:"column:total_likes;not null;default:0"`
	Followers      int    `gorm:"not null;default:0"`
	ViewsThisMonth int64  `gorm:"column:views_this_month;not null;default:0"`
	PostsThisMonth int    `gorm:"column:posts_this_month;not null;default:0"`
	TopPostsJSON   string `gorm:"column:top_posts_json;not null;default:'[]'"`
}

func (analyticsRow) TableName() string { return "user_analytics" }

type GORMStore struct{ db *gorm.DB }

func NewGORMStore(db *gorm.DB) *GORMStore { return &GORMStore{db: db} }

func (s *GORMStore) GetByID(ctx context.Context, id string) (*models.User, error) {
	var m row
	err := s.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorsx.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return fromRow(&m), nil
}

func (s *GORMStore) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var m row
	err := s.db.WithContext(ctx).First(&m, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorsx.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}
	return fromRow(&m), nil
}

// GetOwner returns the seeded site owner (user-1).
func (s *GORMStore) GetOwner(ctx context.Context) (*models.User, error) {
	return s.GetByID(ctx, "user-1")
}

func (s *GORMStore) GetAnalytics(ctx context.Context, userID string) (*models.Analytics, error) {
	var m analyticsRow
	err := s.db.WithContext(ctx).First(&m, "user_id = ?", userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorsx.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get analytics: %w", err)
	}
	return &models.Analytics{
		UserID: m.UserID, TotalViews: m.TotalViews, TotalPosts: m.TotalPosts,
		TotalLikes: m.TotalLikes, Followers: m.Followers,
		ViewsThisMonth: m.ViewsThisMonth, PostsThisMonth: m.PostsThisMonth,
		TopPostsJSON: m.TopPostsJSON,
	}, nil
}

func fromRow(m *row) *models.User {
	name := m.Name
	if name == "" {
		name = m.DisplayName
	}
	return &models.User{
		ID: m.ID, Email: m.Email, Name: name, DisplayName: m.DisplayName,
		Username: m.Username, Avatar: m.Avatar, Bio: m.Bio,
		PasswordHash: m.PasswordHash, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt,
	}
}

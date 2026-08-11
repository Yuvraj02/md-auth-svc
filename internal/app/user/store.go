package user

import (
	"context"

	"github.com/marketing-digest/auth-service/internal/app/user/models"
)

// Store persists users and analytics.
type Store interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetOwner(ctx context.Context) (*models.User, error)
	GetAnalytics(ctx context.Context, userID string) (*models.Analytics, error)
}

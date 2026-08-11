package transport

import (
	"context"

	authv1 "github.com/Yuvraj02/md-protos/proto/auth/v1"
	"github.com/marketing-digest/pkg/grpcx"

	"github.com/marketing-digest/auth-service/internal/app/user"
	"github.com/marketing-digest/auth-service/internal/app/user/models"
)

// Handler implements AuthService gRPC methods.
type Handler struct {
	authv1.UnimplementedAuthServiceServer
	svc *user.Service
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Ping(ctx context.Context, req *authv1.PingRequest) (*authv1.PingResponse, error) {
	msg, service, err := h.svc.Ping(ctx, req.GetMessage())
	if err != nil {
		return nil, grpcx.ToStatus(err)
	}
	return &authv1.PingResponse{Message: msg, Service: service}, nil
}

func (h *Handler) GetMe(ctx context.Context, _ *authv1.GetMeRequest) (*authv1.GetMeResponse, error) {
	u, err := h.svc.GetMe(ctx)
	if err != nil {
		return nil, grpcx.ToStatus(err)
	}
	return &authv1.GetMeResponse{User: toProtoUser(u)}, nil
}

func (h *Handler) GetUserByUsername(ctx context.Context, req *authv1.GetUserByUsernameRequest) (*authv1.GetUserByUsernameResponse, error) {
	u, err := h.svc.GetByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, grpcx.ToStatus(err)
	}
	return &authv1.GetUserByUsernameResponse{User: toProtoUser(u)}, nil
}

func (h *Handler) GetPublicProfile(ctx context.Context, req *authv1.GetPublicProfileRequest) (*authv1.GetPublicProfileResponse, error) {
	u, err := h.svc.GetByID(ctx, req.GetUserId())
	if err != nil {
		return nil, grpcx.ToStatus(err)
	}
	return &authv1.GetPublicProfileResponse{User: toProtoUser(u)}, nil
}

func (h *Handler) GetUserAnalytics(ctx context.Context, req *authv1.GetUserAnalyticsRequest) (*authv1.GetUserAnalyticsResponse, error) {
	a, err := h.svc.GetAnalytics(ctx, req.GetUserId())
	if err != nil {
		return nil, grpcx.ToStatus(err)
	}
	out := &authv1.UserAnalytics{
		UserId: a.UserID, TotalViews: a.TotalViews, TotalPosts: int32(a.TotalPosts),
		TotalLikes: a.TotalLikes, Followers: int32(a.Followers),
		ViewsThisMonth: a.ViewsThisMonth, PostsThisMonth: int32(a.PostsThisMonth),
	}
	for _, t := range a.TopPosts {
		out.TopPosts = append(out.TopPosts, &authv1.TopPost{Id: t.ID, Title: t.Title, Views: t.Views})
	}
	return &authv1.GetUserAnalyticsResponse{Analytics: out}, nil
}

func toProtoUser(u *models.User) *authv1.User {
	if u == nil {
		return nil
	}
	return &authv1.User{
		Id: u.ID, Name: u.Name, Email: u.Email, Username: u.Username,
		Avatar: u.Avatar, Bio: u.Bio, CreatedAt: user.FormatTime(u.CreatedAt),
	}
}

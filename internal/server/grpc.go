package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	authv1 "github.com/Yuvraj02/md-protos/proto/auth/v1"
	"github.com/marketing-digest/pkg/grpcx"

	"github.com/marketing-digest/auth-service/internal/app/user"
	usertransport "github.com/marketing-digest/auth-service/internal/app/user/transport"
)

// Server owns process-level gRPC lifecycle only.
type Server struct {
	grpc   *grpc.Server
	health *health.Server
	log    *slog.Logger
	addr   string
}

func New(svc *user.Service, log *slog.Logger, port int) *Server {
	s := grpc.NewServer(grpcx.UnaryServerInterceptors(log))
	hs := health.NewServer()

	authv1.RegisterAuthServiceServer(s, usertransport.NewHandler(svc))
	healthpb.RegisterHealthServer(s, hs)
	reflection.Register(s)

	hs.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)
	hs.SetServingStatus("auth.v1.AuthService", healthpb.HealthCheckResponse_NOT_SERVING)

	return &Server{
		grpc: s, health: hs, log: log,
		addr: fmt.Sprintf(":%d", port),
	}
}

func (s *Server) SetReady(ready bool) {
	status := healthpb.HealthCheckResponse_NOT_SERVING
	if ready {
		status = healthpb.HealthCheckResponse_SERVING
	}
	s.health.SetServingStatus("", status)
	s.health.SetServingStatus("auth.v1.AuthService", status)
}

func (s *Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("listen %s: %w", s.addr, err)
	}
	s.log.Info("grpc listening", "addr", s.addr)
	return s.grpc.Serve(ln)
}

func (s *Server) GracefulStop(ctx context.Context) {
	done := make(chan struct{})
	go func() {
		s.grpc.GracefulStop()
		close(done)
	}()
	select {
	case <-done:
	case <-ctx.Done():
		s.grpc.Stop()
	}
}

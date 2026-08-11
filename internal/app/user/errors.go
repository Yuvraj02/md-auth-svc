package user

import "github.com/marketing-digest/pkg/errorsx"

var (
	ErrNotFound         = errorsx.ErrNotFound
	ErrInvalidArgument  = errorsx.ErrInvalidArgument
	ErrUnauthenticated  = errorsx.ErrUnauthenticated
	ErrPermissionDenied = errorsx.ErrPermissionDenied
	ErrConflict         = errorsx.ErrConflict
)

package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AutenticacionService) AutenticacionService

type loggingMiddleware struct {
	logger log.Logger
	next   AutenticacionService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AutenticacionService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AutenticacionService) AutenticacionService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, email string) (e0 error) {
	defer func() {
		l.logger.Log("method", "Create", "email", email, "e0", e0)
	}()
	return l.next.Create(ctx, email)
}

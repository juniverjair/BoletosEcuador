package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(VentasService) VentasService

type loggingMiddleware struct {
	logger log.Logger
	next   VentasService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a VentasService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next VentasService) VentasService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SendEmail(ctx context.Context, email string, content string) (e0 error) {
	defer func() {
		l.logger.Log("method", "SendEmail", "email", email, "content", content, "e0", e0)
	}()
	return l.next.SendEmail(ctx, email, content)
}

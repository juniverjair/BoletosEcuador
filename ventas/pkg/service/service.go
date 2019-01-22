package service

import "context"

// VentasService describes the service.
type VentasService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SendEmail(ctx context.Context, email string, content string) error
}

type basicVentasService struct{}

func (b *basicVentasService) SendEmail(ctx context.Context, email string, content string) (e0 error) {
	// TODO implement the business logic of SendEmail
	return e0
}

// NewBasicVentasService returns a naive, stateless implementation of VentasService.
func NewBasicVentasService() VentasService {
	return &basicVentasService{}
}

// New returns a VentasService with all of the expected middleware wired in.
func New(middleware []Middleware) VentasService {
	var svc VentasService = NewBasicVentasService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

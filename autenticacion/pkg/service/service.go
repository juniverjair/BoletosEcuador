package service

import "context"

// AutenticacionService describes the service.
type AutenticacionService interface {
	// Add your methods here
	Create(ctx context.Context, email string) error
}

type basicAutenticacionService struct{}

func (b *basicAutenticacionService) Create(ctx context.Context, email string) (e0 error) {
	// TODO implement the business logic of Create
	return e0
}

// NewBasicAutenticacionService returns a naive, stateless implementation of AutenticacionService.
func NewBasicAutenticacionService() AutenticacionService {
	return &basicAutenticacionService{}
}

// New returns a AutenticacionService with all of the expected middleware wired in.
func New(middleware []Middleware) AutenticacionService {
	var svc AutenticacionService = NewBasicAutenticacionService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

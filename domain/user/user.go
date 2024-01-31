package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Greeter .
type Greeter struct {
	Hello string
}

// UserRepo .
type UserRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	// ...
}

// UserDomain .
type UserDomain struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserDomain .
func NewUserDomain(repo UserRepo, logger log.Logger) *UserDomain {
	return &UserDomain{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserDomain) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

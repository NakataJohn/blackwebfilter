package domain

import "context"

type DomainFilterRequest struct {
	URL string `form:"url" binding:"required"`
}

type DomainFilterUsecase interface {
	Verify(c context.Context, url string) (string, error)
}

package domain

import "context"

type FilterRequest struct {
	URL string `form:"url" binding:"required"`
}

type FilterUsecase interface {
	Verify(c context.Context, url string) (string, error)
}

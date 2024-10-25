package domain

import "context"

type WordFilterRequest struct {
	URL string `form:"url" binding:"required"`
}

type WordFilterUsecase interface {
	Verify(c context.Context, url string) (string, error)
}

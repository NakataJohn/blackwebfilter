package usecase

import (
	"blackwebfilter/domain"
	"context"
	"time"
)

type pingUsecase struct {
	contextTimeout time.Duration
}

func NewPingUsecase(timeout time.Duration) domain.PingUsecase {
	return &pingUsecase{
		contextTimeout: timeout,
	}
}

func (pu *pingUsecase) Ping(c context.Context) (string, error) {
	_, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return "pong", nil
}

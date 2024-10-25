package domain

import "context"

type PingUsecase interface {
	Ping(c context.Context) (string, error)
}

package route

import (
	"blackwebfilter/api/controller"
	"blackwebfilter/bootstrap"
	"blackwebfilter/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewPingRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	pc := controller.PingController{
		PingUsecase: usecase.NewPingUsecase(timeout),
		Env:         env,
	}
	group.GET("/ping", pc.Ping)
}

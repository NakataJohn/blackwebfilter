package route

import (
	"blackwebfilter/api/controller"
	"blackwebfilter/bootstrap"
	"blackwebfilter/usecase"
	"time"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/gin-gonic/gin"
)

// Black domains filter router.
func NewDomainFilterRouter(env *bootstrap.Env, timeout time.Duration, bloomfilter *bloom.BloomFilter, group *gin.RouterGroup) {
	fc := controller.DomainFilterController{
		DomainFilterUsecase: usecase.NewDomainFilterUsecase(timeout, bloomfilter),
		Env:                 env,
	}
	group.POST("/domains", fc.Verify)
}

package route

import (
	"blackwebfilter/api/controller"
	"blackwebfilter/bootstrap"
	"blackwebfilter/usecase"
	"time"

	"github.com/bits-and-blooms/bloom/v3"

	ahocorasick "github.com/BobuSumisu/aho-corasick"
	"github.com/gin-gonic/gin"
)

// comprehensive evaluation filter router.
func NewFilterRouter(env *bootstrap.Env, timeout time.Duration, bloomfilter *bloom.BloomFilter, trie *ahocorasick.Trie, group *gin.RouterGroup) {
	fc := controller.FilterController{
		Wc: usecase.NewSensitiveWordUsecase(timeout, trie),
		Dc: usecase.NewDomainFilterUsecase(timeout, bloomfilter),
	}
	group.POST("/filter", fc.Verify)
}

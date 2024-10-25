package route

import (
	"blackwebfilter/api/controller"
	"blackwebfilter/bootstrap"
	"blackwebfilter/usecase"
	"time"

	ahocorasick "github.com/BobuSumisu/aho-corasick"
	"github.com/gin-gonic/gin"
)

// Sensitive word filter router.
func NewWordFilterRouter(env *bootstrap.Env, timeout time.Duration, trie *ahocorasick.Trie, group *gin.RouterGroup) {
	fc := controller.WordFilterController{
		WordFilterUsecase: usecase.NewSensitiveWordUsecase(timeout, trie),
		Env:               env,
	}
	group.POST("/sensitive", fc.Verify)
}

package route

import (
	"blackwebfilter/bootstrap"
	"time"

	ahocorasick "github.com/BobuSumisu/aho-corasick"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Setup(env *bootstrap.Env, bloomfilter *bloom.BloomFilter, trie *ahocorasick.Trie, gin *gin.Engine) {

	gin.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	pubilcRouter := gin.Group("/api/v1")
	// pubilcRouter := gin.Group("")
	// All Public APIs
	timeout := time.Duration(env.ContextTimeout) * time.Second

	NewPingRouter(env, timeout, pubilcRouter)
	NewDomainFilterRouter(env, timeout, bloomfilter, pubilcRouter)
	NewWordFilterRouter(env, timeout, trie, pubilcRouter)
	NewFilterRouter(env, timeout, bloomfilter, trie, pubilcRouter)

	// protectedRouter := gin.Group("/api/v1")
	// protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// All Private APIs
	// NewProfileRouter(env, timeout, db, protectedRouter)
	// NewTaskRouter(env, timeout, db, protectedRouter)
	// NewBookRouter(env, timeout, db, protectedRouter)
}

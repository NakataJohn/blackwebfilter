package main

import (
	"blackwebfilter/api/route"
	"blackwebfilter/bootstrap"
	_ "blackwebfilter/docs"
	"blackwebfilter/internal/logger"
	"fmt"
	"io"
	"os"
	"path/filepath"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
)

func init() {
	logo := `
 _   _         _             _   ___ _ _ _           
| |_| |___ ___| |_ _ _ _ ___| |_|  _|_| | |_ ___ ___ 
| . | | .'|  _| '_| | | | -_| . |  _| | |  _| -_|  _|
|___|_|__,|___|_,_|_____|___|___|_| |_|_|_| |___|_|  
	`
	fmt.Println(logo)
}

// @title WebBloom Api
// @version v1.0.0
// @description WebBloom Api
// @termsOfService http://swagger.io/terms/
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey token
// @in header
// @name Authorization
func main() {
	logger.Infof("WebBloom App is starting...")
	app := bootstrap.App()
	env := app.Env

	// 初始化布隆过滤器
	filter, err := bootstrap.NewBloomFilter(env)
	if err != nil {
		panic(err)
	}

	// 初始化AC自动机
	trie, err := bootstrap.NewTrieBuilder(env)
	if err != nil {
		panic(err)
	}

	if env.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	f, _ := os.Create(filepath.Join(env.ZapPath, "gin.log"))
	gin.DefaultWriter = io.MultiWriter(f)
	// 启动服务
	gin := gin.Default()
	// 添加设置最大并发请求数中间件；
	gin.Use(limit.MaxAllowed(env.LimitRequests))
	route.Setup(env, filter, trie, gin)
	gin.Run(env.ServerAddress)
}

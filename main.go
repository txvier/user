package main

import (
	"github.com/gin-gonic/gin"
	"github.com/txvier/base/http"
	"github.com/txvier/user/routers"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	gin.SetMode(gin.DebugMode)

	eg := http.NewGin()
	rg := eg.NewGroup("/txvier/user/api/v1")
	rg.RegisterHttpRouters(routers.HttpRouters...)
	eg.Run(":8888")
}

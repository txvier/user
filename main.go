package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tarzan/base/http"
	"github.com/tarzan/user/routers"
	"runtime"
)


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	gin.SetMode(gin.DebugMode)

	eg := http.NewGin()
	rg := eg.NewGroup("/tarzan/user/api/v1")
	rg.RegisterHttpRouters(routers.HttpRouters...)
	eg.Run(":8888")
}

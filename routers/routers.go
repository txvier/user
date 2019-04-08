package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/txvier/base/http"
)

var (
	HttpRouters = []http.HttpRouter{
		optionsRouter,
		userRouter,
		// add new router here...
	}

	middlewares = []gin.HandlerFunc{
		gin.Logger(),
		//middleware.Recovery(),
		//middleware.ValidateTokenHandlerFunc(),
		CorsMW(),
	}
)

func CorsMW() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"http://localhost:8082"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "HEAD"},
		AllowHeaders:     []string{"authorization", "content-type"},
		AllowCredentials: true,
		// MaxAge: 12 * time.Hour,
	})
}

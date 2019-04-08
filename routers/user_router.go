package routers

import (
	"github.com/txvier/base/http"
	"github.com/txvier/user/api"
)

var (
	userRouter = http.ApiRouter{
		http.Route{
			Name:        "cvs_create_cvm_bsinfo", //新增基本信息
			Methods:     []string{"POST"},
			Pattern:     "/cvs/:cvid/cvms/bsinfo",
			HandlerFunc: http.HttpFuncWrap(api.Loggin),
		},
	}
)

type Student struct {
	Name string
	Age int
}

var arr = []Student{
	{
		"1",
		100,
	},

}
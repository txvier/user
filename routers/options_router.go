package routers

import "github.com/txvier/base/http"

var (
	optionsRouter = http.ApiRouter{
		http.Route{
			Name:        "options_action", 
			Methods:     []string{"OPTIONS"},
			Pattern:     "/*action",
			HandlerFunc: CorsMW(),
		},
	}
)
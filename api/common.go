package api

import (
	"github.com/txvier/base/txlogger"
)

var (
	Logger       = txlogger.Logger
	CURRENT_USER = "CURRENT_USER"
)

const (
// API_KEY   = vars.PROJECT_NAME
)

///////////////////////////////////////////////////////////////////////////////

//PageableRequest
type PageableRequest struct {
	Limit int64 `query:"limit" binding:"required,numeric"`
	Page  int64 `query:"page" binding:"required,numeric"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErr(rte func() error) {
	check(rte())
}

type FuncErr func() error

type ErrIsNilExec struct {
	err error
}

func (eine *ErrIsNilExec) Exec(f FuncErr) *ErrIsNilExec {
	if eine.err == nil {
		eine.err = f()
	}
	return eine
}

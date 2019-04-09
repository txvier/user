package api

import (
	"github.com/txvier/base/common"
	"github.com/txvier/base/http"
	"github.com/txvier/base/txerror"
	"github.com/txvier/base/txjwt"
	"github.com/txvier/user/models"
	"github.com/txvier/user/repositories"
)

var (
	usersRepository = repositories.NewUsersRepository()
)

//用户登录
func Loggin(req *http.ApiRequest) (rsp http.ApiResponse) {
	var u models.User
	if err := req.BindJSON(&u); err != nil {
		return rsp.BindError(txerror.ErrTrace(err))
	}
	if udb, err := usersRepository.GetUserForAuth(u); err != nil {
		return rsp.BindError(txerror.ErrTrace(err))
	} else {
		if udb.UId == "" {
			return rsp.BindError(txerror.ErrForbidden("user name or pwd err."))
		}

		ut := &txjwt.UserToken{
			Name: udb.Name,
			Pwd:  udb.Pwd,
			Id:   udb.UId,
		}
		if token, err := txjwt.CreateToken(ut); err != nil {
			return rsp.BindError(txerror.ErrTrace(err))
		} else {
			return rsp.AddAttribute(txjwt.TOKEN_KEY, token).AddAttribute("auth_user", udb)
		}
	}
}

//新增用户
func CreateUser(req *http.ApiRequest) (rsp http.ApiResponse) {
	var u models.User
	if err := req.BindJSON(&u); err != nil {
		return rsp.BindError(txerror.ErrTrace(err))
	}
	if err := usersRepository.CreateUser(u); err != nil {
		return rsp.BindError(txerror.ErrTrace(err))
	}
	return
}

//用户注册
func RegisterUser(req *http.ApiRequest) (rsp http.ApiResponse) {
	return CreateUser(req)
}

/*
* 返回用户列表
 */
func GetAllUsers(req *http.ApiRequest) (rsp http.ApiResponse) {
	var pr PageableRequest
	if err := req.BindStruct(&pr); err != nil {
		return rsp.BindError(txerror.ErrTrace(err))
	}
	p, err := common.NewPageable(pr.Limit, pr.Page)
	if err != nil {
		return rsp.BindError(txerror.ErrTrace(err))
	}
	p, err = usersRepository.GetAllUsers(p)
	return rsp.BindError(err).AddObject(p)
}

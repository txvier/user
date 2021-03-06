package repositories

import (
	"github.com/txvier/base/common"
	"github.com/txvier/user/models"
	// "github.com/goinggo/mapstructure"
)

type UsersRepository struct{}

var (
	ur = &UsersRepository{}
)

//单例构造函数
func NewUsersRepository() *UsersRepository {
	return ur
}

//返回用户信息
func (this *UsersRepository) GetUserForAuth(u models.User) (udb models.User, err error) {

	/*params := make(map[string]interface{})
	params["name"] = u.Name
	params["pwd"] = u.Pwd

	c := NewCypher("users_repository.get_user_for_auth.cypher").Params(params)

	callback := func(row []interface{}) {
		udb = models.User{
			UId:  common.NilParseString(row[0]),
			Name: common.NilParseString(row[1]),
			Pwd:  common.NilParseString(row[2]),
		}
	}

	return udb, QueryNeo(callback, c)*/
	return
}

//用户是否存在
func (this *UsersRepository) IsExist(u models.User) (b bool, err error) {

	/*params := make(map[string]interface{})
	params["name"] = u.Name

	c := NewCypher("users_repository.is_exist.cypher").Params(params)

	var rowNum int64

	callback := func(row []interface{}) {
		rowNum++
	}

	err := QueryNeo(callback, c)

	if err := txerror.ErrInternalServer(err); err != nil {
		return false, err
	}

	if rowNum == 0 {
		return false, nil
	}

	return true, nil*/
	return
}

//新增用户
func (this *UsersRepository) CreateUser(u models.User) (err error) {

	/*isExist, err := this.IsExist(u)

	if err := txerror.ErrInternalServer(err); err != nil {
		return err
	}

	if isExist {
		return txerror.ErrForbiddenf("user name [%s] is exist..", u.Name)
	}

	m := map[string]interface{}{
		"uid":         common.UUID(),
		"name":        u.Name,
		"pwd":         u.Pwd,
		"create_time": time.Now().UnixNano(),
	}

	c := NewCypher("users_repository.create_user.cypher").Params(m)

	return ExecNeo(c)*/
	return
}

//返回所有用户
func (this *UsersRepository) GetAllUsers(p common.Pageable) (_ common.Pageable, err error) {

	/*c := NewCypher("users_repository.get_all_users.cypher_count")

	var count int64

	callback := func(row []interface{}) {
		count = common.NilParseInt64(row[0])
	}

	err = QueryNeo(callback, c)

	if err := txerror.ErrInternalServer(err); err != nil {
		return p, err
	}

	Logger.Info("row count is:", count)
	if count == 0 {
		return p, nil
	}
	p.SetTotalElements(count)

	params := make(map[string]interface{})
	params["limit"] = p.PageSize
	params["offset"] = p.GetOffSet()

	c = NewCypher("users_repository.get_all_users.cypher_pageable").Params(params)

	callback = func(row []interface{}) {
		m := &models.User{
			UId:        common.NilParseString(row[0]),
			Name:       common.NilParseString(row[1]),
			Pwd:        common.NilParseString(row[2]),
			CreateTime: common.NilParseJSONTime(row[3]),
		}
		p.AddContent(m)
	}

	return p, QueryNeo(callback, c)*/
	return
}

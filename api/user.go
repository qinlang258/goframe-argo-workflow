package api

import "github.com/gogf/gf/v2/frame/g"

type UserAddReq struct {
	g.Meta   `path:"/api/user/add" tags:"user" method:"post" summary:"创建用户"`
	ID       string `json:"user_id"`
	Name     string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Password string `json:"password" description:"密码" v:"required#密码必填"`
	Mobile   string `json:"mobile" description:"密码" v:"required#电话必填"`
}

type UserAddRes struct {
	Id uint `json:"id"`
}

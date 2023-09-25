package user

import (
	"context"
	"goframe-argo-workflow/api"
)

type User struct{}

func init() {

}

func New() *User {
	return &User{}
}

func (c *User) UserCreate(ctx context.Context, req *api.UserAddReq) (err error) {
	user := &api.UserAddReq{}

	user.Name = "qinlang"
	user.Mobile = "18428000826"
	user.Password = "ql2252528"

	return err
}

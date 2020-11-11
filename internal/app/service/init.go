package service

import (
	"jupiter-tutorials/internal/app/model"
	"jupiter-tutorials/internal/app/service/user"
	"jupiter-tutorials/internal/app/service/user/impl"
)

var (
	UserRepository user.Repository
)
//Init instantiate the service
func Init()  {
	UserRepository = impl.NewMysqlImpl(model.MysqlHandler)
}
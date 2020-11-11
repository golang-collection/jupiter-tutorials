package user

import (
	"jupiter-tutorials/internal/app/model/db"
)

type Repository interface {
	Get(id int) (user db.User,err error)
}

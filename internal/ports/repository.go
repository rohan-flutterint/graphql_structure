package ports

import "test/internal/model"

type TestRepo interface {
	GetData(id string) *model.User
	GetDataall() []*model.User
	CreateUser(user *model.User) *model.User
}

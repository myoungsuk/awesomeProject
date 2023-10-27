package repo

import "awesomeProject/internal/user/domain"

type QueryUserData interface {
	FindById(name string) []domain.Account
}

type CommandUserData interface {
	Save(aggregate domain.Account) error
}

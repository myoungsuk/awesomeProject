package article

import "awesomeProject/internal/article/domain"

type QueryService interface {
	SearchUser(name string) (*domain.Article, error)
}

type CommandService interface {
	SaveUser(req SaveUserRequest) error
}

type SaveUserRequest struct {
	Username string
	Password string
	Email    string
	Nickname string
	Memo     string
}

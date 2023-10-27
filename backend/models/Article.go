package models

type Article struct {
	ID         int64
	CreatedAt  string
	CreatedBy  string
	ModifiedAt string
	ModifiedBy string
	Content    string
	Title      string
	UserID     string
}

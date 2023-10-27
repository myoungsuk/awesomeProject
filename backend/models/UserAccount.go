package models

import (
	"gorm.io/gorm"
)

type ArticleComment struct {
	gorm.Model
	ArticleID       uint
	UserAccountID   uint
	ParentCommentID uint
	Content         string
	ChildComments   []*ArticleComment `gorm:"foreignKey:ParentCommentID"`
}

func NewArticleComment(articleID, userAccountID uint, content string) *ArticleComment {
	return &ArticleComment{
		ArticleID:     articleID,
		UserAccountID: userAccountID,
		Content:       content,
	}
}

func (ac *ArticleComment) AddChildComment(child *ArticleComment) {
	child.ParentCommentID = ac.ID
}

func (ac *ArticleComment) Equals(other *ArticleComment) bool {
	return ac.ID == other.ID
}

func (ac *ArticleComment) HashCode() int {
	return int(ac.ID)
}

package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
)

type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`

	UserID     uint64 `gorm:"not null;index"`
	CategoryID uint64 `gorm:"not null;default:4;index"`

	User user.User
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", article.GetStringID())
}

func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}

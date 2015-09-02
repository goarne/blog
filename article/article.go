//Article entity
package article

import (
	"time"
)

//Domain entity for holding articles.
type Article struct {
	Id       int32
	Title    string
	Abstract string
	Text     string
	Date     time.Time
	Author   string
}

//Function creates a new transient instance of an article.
func NewArticle() *Article {
	a := Article{}
	a.Date = time.Now()
	return &a
}

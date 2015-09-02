package article

import (
	"gopkg.in/mgo.v2"
	"testing"
)

var svc ServiceImpl

func TestServiceCRUD(t *testing.T) {
	shouldCreateArticle(t)
	shouldFindArticle(t)
	shoudlUpdateArticle(t)
	shouldDeleteArticle(t)
	shouldFindAllArticles(t)
}

func shouldFindArticle(t *testing.T) {
	executeOnCollection = func(collName string, q func(*mgo.Collection) error) error {
		return nil
	}

	article, err := svc.Find(1)

	if err != nil || article == nil {
		t.Errorf("Could not find article.", err)
	}
}

func shouldCreateArticle(t *testing.T) {
	executeOnCollection = func(collName string, q func(*mgo.Collection) error) error {
		return nil
	}

	newArticle := NewArticle()

	newArticle.Id = 1
	newArticle.Title = "About java"
	newArticle.Text = "Presentation about java."

	err := svc.Create(newArticle)

	if err != nil {
		t.Errorf("Couyld not create new article.", err)
	}
}

func shouldDeleteArticle(t *testing.T) {
	executeOnCollection = func(collName string, q func(*mgo.Collection) error) error {
		return nil
	}

	err := svc.Delete(1)

	if err != nil {
		t.Errorf("Could not delete article.", err)
	}
}

func shoudlUpdateArticle(t *testing.T) {
	executeOnCollection = func(collName string, q func(*mgo.Collection) error) error {
		return nil
	}
	article := NewArticle()
	article.Id = 1
	article.Title = "About java2"
	article.Text = "Advanced about java."

	err := svc.Update(article)

	if err != nil {
		t.Errorf("Could not update article.", err)
	}
}

func shouldFindAllArticles(t *testing.T) {
	var a []*Article

	executeOnCollection = func(collName string, q func(*mgo.Collection) error) error {

		a1 := NewArticle()
		a1.Id = 1
		a1.Title = "About java2"
		a1.Text = "Advanced java."
		a = append(a, a1)

		a2 := NewArticle()
		a2.Id = 1
		a2.Title = "About Golang"
		a2.Text = "Advanced about golang."
		a = append(a, a2)
		return nil
	}

	_, err := svc.FindAll()

	if err != nil {
		t.Errorf("Failed retrieving all articles.")
	}
}

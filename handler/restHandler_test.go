package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nilsen.no/blog/article"
	"nilsen.no/blog/config"

	"testing"
)

func TestHandleRequest(t *testing.T) {

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", config.BLOGG_URL_API, nil)

	controller := createRestHandler()

	controller.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Handlerequest returned %v", resp.Code)
	}
}

func TestGetArticle(t *testing.T) {

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", config.BLOGG_URL_API+"/1", nil)

	params := req.URL.Query()
	params.Add("id", "1")

	req.URL.RawQuery = params.Encode()

	controller := createRestHandler()
	controller.ReceiveGet(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Handlerequest returned %v", resp.Code)
	}

	foundArticle := &article.Article{}
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&foundArticle)

	if foundArticle.Id == 0 {
		t.Errorf("Could not find id.")
	}
}

func TestGetTwoArticles(t *testing.T) {

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", config.BLOGG_URL_API, nil)

	controller := createRestHandler()
	controller.ReceiveGet(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Handlerequest returned %v", resp.Code)
	}

	articles := []article.Article{}
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&articles)

	if len(articles) != 2 {
		t.Errorf("Could not find any articles.")
	}
}

func TestPutArticle(t *testing.T) {
	article := &article.Article{Id: 1, Title: "Testtittel", Text: "Test innhold"}
	articleBytes, _ := json.Marshal(&article)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", config.BLOGG_URL_API+"/1", bytes.NewReader(articleBytes))

	controller := createRestHandler()

	controller.ReceivePut(resp, req) // Want to update existing resource.

	if resp.Code != http.StatusOK {
		t.Errorf("Handlerequest returned %v", resp.Code)
	}
}

func TestPostArticle(t *testing.T) {
	article := &article.Article{Id: 1, Title: "Test", Text: "Test content"}
	articleBytes, _ := json.Marshal(&article)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", config.BLOGG_URL_API, bytes.NewReader(articleBytes))

	controller := createRestHandler()

	controller.ReceivePost(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Handlerequest returned %v", resp.Code)
	}
}

//Helper methods
func createRestHandler() *RestHandler {
	return &RestHandler{WebHandler{MockedServiceImpl{}}}
}

//Mocking service to reduce dependencies.
type MockedServiceImpl struct{}

func (MockedServiceImpl) Create(f *article.Article) error {
	return nil
}

func (MockedServiceImpl) Update(a *article.Article) error {
	return nil
}

func (MockedServiceImpl) Delete(id int32) error {
	return nil
}

func (MockedServiceImpl) Find(id int32) (*article.Article, error) {
	return &article.Article{Id: 1, Title: "test -title", Text: "Test content"}, nil
}

func (MockedServiceImpl) FindAll() ([]article.Article, error) {
	var articles []article.Article
	articles = append(articles, article.Article{Id: 1, Title: "test -title 1", Text: "Test content 1"})
	articles = append(articles, article.Article{Id: 2, Title: "test -title 2", Text: "Test content 2"})
	return articles, nil
}

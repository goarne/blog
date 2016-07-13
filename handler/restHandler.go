package handler

import (
	"encoding/json"
	"net/http"

	"github.com/goarne/blog/article"
)

//The controller handles CRUD requests wiht JSON responses.
type RestHandler struct {
	WebHandler
}

func NewRestHandler() *RestHandler {
	return &RestHandler{WebHandler{article.ServiceImpl{}}}
}

func (c RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.ReceiveGet(w, r)
	case "POST":
		c.ReceivePost(w, r)
	case "PUT":
		c.ReceivePut(w, r)
	case "DELETE":
		c.ReceiveDelete(w, r)
	case "OPTIONS":
		w.Write([]byte("The handler supports GET, POST, PUT and DELETE!"))
	default:
		http.Error(w, "Method not supported.", 405)
	}
}

func (c RestHandler) ReceivePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	newArticle := article.Article{}
	err := decoder.Decode(&newArticle)

	if err != nil {
		http.Error(w, "Invalid content.", 400)
		return
	}

	err = c.Create(&newArticle)

	if err != nil {
		http.Error(w, "Could not update/create article", 404)
		return
	}
}

func (c RestHandler) ReceivePut(w http.ResponseWriter, r *http.Request) {
	id, _ := parseId(r)
	decoder := json.NewDecoder(r.Body)
	newArticle := article.Article{Id: id}
	err := decoder.Decode(&newArticle)

	if err != nil {
		http.Error(w, "Invalid content.", 400)
		return
	}

	article, err := c.Find(id)

	if err != nil {
		http.Error(w, "Could not find article.", 404)
		return
	}

	if article == nil {
		err = c.Create(&newArticle)
	} else {
		err = c.Update(&newArticle)
	}

	if err != nil {
		http.Error(w, "Could not put article.", 404)
		return
	}
}

func (c RestHandler) ReceiveGet(w http.ResponseWriter, r *http.Request) {
	id, _ := parseId(r)

	if id == 0 {
		articles, err := c.FindAll()
		if err != nil {
			http.Error(w, "Could not find article.", 404)
			return
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(&articles)

		return
	}

	article, err := c.Find(id)

	if err != nil {
		http.Error(w, "Could not find article.", 404)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&article)
}

func (c RestHandler) ReceiveDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := parseId(r)

	err := c.Delete(id)

	if err != nil {
		http.Error(w, "Problems deleting article.", 405)
		return
	}

}

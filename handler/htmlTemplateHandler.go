package handler

import (
	"html/template"
	"net/http"
	"nilsen.no/blog/article"
	"strings"
	"time"
)

//The HtmlHandler handles http requests from browser.
//It serves static html files aswell as dynamic html files.
type HtmlTemplateHandler struct {
	WebHandler
}

func NewHtmlTemplateHandler() *HtmlTemplateHandler {
	return &HtmlTemplateHandler{WebHandler{article.ServiceImpl{}}}
}

func (c *HtmlTemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.ReceiveGet(w, r)
	case "POST":
		c.ReceivePost(w, r)
	case "PUT":
		c.ReceivePut(w, r)
	case "DELETE":
		c.ReceiveDelete(w, r)
	default:
		http.Error(w, "Method not supported.", 405)
	}
}

func (c *HtmlTemplateHandler) ReceivePost(w http.ResponseWriter, r *http.Request) {
	var a article.Article

	a.Date = time.Now()
	a.Title = r.FormValue("article_title")
	a.Abstract = r.FormValue("article_abstract")
	a.Text = r.FormValue("article_body")
	a.Author = r.FormValue("article_author")

	err := c.Create(&a)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return

	}
}

func (c *HtmlTemplateHandler) ReceivePut(w http.ResponseWriter, r *http.Request) {

}

func (c *HtmlTemplateHandler) ReceiveGet(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.Path, "list") {
		t, err := template.ParseFiles("./html/list.html")

		if err != nil {
			http.Error(w, "[list]Resource not found", 405)
			return
		}

		alle, _ := c.FindAll()
		t.Execute(w, alle)
		return
	}

	if strings.Contains(r.URL.Path, "ny") {
		t, err := template.ParseFiles("./html/new_article.html")
		if err != nil {
			http.Error(w, "[new]Resource not found.", 404)
			return
		}

		t.Execute(w, nil)
		return
	}

	if strings.Contains(r.URL.Path, "/article/") {
		id, _ := parseId(r)
		a, err := c.Find(id)

		if err != nil {
			http.Error(w, "Could not find article.", 404)
			return
		}

		t, _ := template.ParseFiles("./html/article.html")
		t.Execute(w, &a)

		return
	}

	http.Error(w, "Could not find resource.", 404)
}

func (c *HtmlTemplateHandler) ReceiveDelete(w http.ResponseWriter, r *http.Request) {

}

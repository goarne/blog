//Blog main.go
package main

import (
	//"github.com/davecheney/profile"
	"github.com/nilsengo/web"
	"gopkg.in/mgo.v2"
	"net/http"
	_ "net/http/pprof"
	"nilsen.no/blog/config"
	"nilsen.no/blog/handler"
)

//The application is simple blog webapplication with html and a REST interface.
//It is created with a custom middleware for handing httprequests,
//routing requests to handlers serving html or json content from a database.
func main() {
	//Enables profiler...
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	//Testing the databasesetup.
	initDb()
	router := web.NewWebRouter()

	//The static route handles static html files.
	router.AddRoute(createChainedHtmlRoute())

	//Adds basic html CRUD support for articles
	router.AddRoute(createNewArticleHtmlRoute())
	router.AddRoute(createListArticleHtmlRoute())
	router.AddRoute(createFindArticleHtmlRoute())

	//Adds REST support for CRUD operations on articles.
	router.AddRoute(createJsonRoute())

	http.ListenAndServe(":8888", router)
}

//Creates a html route
func createNewArticleHtmlRoute() *web.Route {
	r := web.NewRoute()
	r.Path("/new_article.html")
	r.Method("POST")
	r.Header("Accept", "html")
	r.Handler(handler.NewHtmlTemplateHandler())
	return r
}

func createListArticleHtmlRoute() *web.Route {
	r := web.NewRoute()
	r.Path("/blog/article/list/")
	r.Method("GET")
	r.Header("Accept", "html")
	r.Handler(handler.NewHtmlTemplateHandler())
	return r
}

func createFindArticleHtmlRoute() *web.Route {
	r := web.NewRoute()
	r.Path("/article/{id}")
	r.Header("Accept", "html")
	r.Method("GET")
	r.Handler(handler.NewHtmlTemplateHandler())
	return r
}

func createChainedHtmlRoute() *web.Route {
	r := web.NewRoute()
	r.Path("/")
	r.PathPrefix("/*.html")
	r.PathPrefix("/css/")
	r.PathPrefix("/ckeditor/")
	r.PathPrefix("/img/")
	r.PathPrefix("/fonts/")
	r.PathPrefix("/js/")
	r.PathPrefix("/less/")
	r.PathPrefix("/mail/")
	r.Method("GET")

	hc := &web.HandlerChain{}

	hc.Add(handler.NewStaticHtmlHandler("index.html"))
	hc.Add(handler.NewFileHandler("./html/"))
	r.Handler(hc)
	return r
}

func createJsonRoute() *web.Route {
	r := web.NewRoute()
	r.Path("/blog/article/{id}")
	r.Method("POST").Method("GET").Method("PUT").Method("DELETE").Method("OPTIONS")
	r.Header("Accept", "json")

	r.Handler(handler.NewRestHandler())
	return r
}

func initDb() {
	session, err := mgo.Dial(config.DOCKER_MONGODB_URL)

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.DB_NAME).C(config.COLL_ARTICLE)

	// Index
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

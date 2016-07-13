package handler

import (
	"github.com/goarne/web"
	"html/template"
	"net/http"
	"github.com/goarne/blog/config"
	"strings"
)

//The StaticHtmlHandler handles http requests for html files.
//It serves static html files by loading the files with template library.
type StaticHtmlHandler struct {
	indexFile string
}

func NewStaticHtmlHandler(index string) *StaticHtmlHandler {
	return &StaticHtmlHandler{index}
}

func (s *StaticHtmlHandler) ServeReq(c *web.ChainedReq) {
	r := c.Req
	w := c.Resp

	if path, isHtml := s.computeHtmlPath(r); isHtml {

		t, err := template.ParseFiles(config.HTML_FOLDER + path)

		if err != nil {
			http.Error(w, "[resource] Resource not found"+path, 404)
			return
		}

		t.Execute(w, nil)
		c.State = web.STOPPED
	}
}

//The function checks if the url path requests loading of static htmlfiles.
func (s *StaticHtmlHandler) computeHtmlPath(r *http.Request) (string, bool) {
	var html string

	if (r.URL.Path == "/") && (s.indexFile != "") {
		html = s.indexFile
	} else {
		html = r.URL.Path
	}

	if strings.Contains(html, config.HTML_FILE_EXT) {
		return html, true
	}

	return "", false
}

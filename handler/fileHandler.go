package handler

import (
	"github.com/goarne/web"
	"net/http"
)

type FileHandler struct {
	http.Handler
}

func NewFileHandler(p string) *FileHandler {
	return &FileHandler{http.FileServer(http.Dir(p))}
}

func (f *FileHandler) ServeReq(c *web.ChainedReq) {

	f.ServeHTTP(c.Resp, c.Req)
	c.State = web.STOPPED
}

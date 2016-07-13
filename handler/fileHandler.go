package handler

import (
	"net/http"

	"github.com/goarne/web"
)

type FileHandler struct {
	http.Handler
}

func NewFileHandler(p string) *FileHandler {
	return &FileHandler{http.FileServer(http.Dir(p))}
}

func (f *FileHandler) ServeReq(c *web.ChainedReq) {
	//http.ServeFile()
	f.ServeHTTP(c.Resp, c.Req)
	c.State = web.STOPPED
}

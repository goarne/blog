package handler

import (
	"net/http"
	"strconv"

	"github.com/goarne/blog/article"
)

//Base controller for handling webrequests.
type WebHandler struct {
	article.Service
}

func parseId(r *http.Request) (int32, error) {
	id := r.URL.Query().Get("id")
	result, err := strconv.ParseInt(id, 10, 32)
	return int32(result), err
}

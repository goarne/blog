package handler

import (
	//	"bytes"
	"github.com/nilsengo/web"
	"net/http"
	"net/http/httptest"
	"nilsen.no/blog/config"
	"testing"
)

func TestStaticHtmlHanderRequestGet(t *testing.T) {

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", config.BLOGG_URL_API, nil)

	controller := StaticHtmlHandler{}
	c := &web.ChainedReq{resp, req, web.ACTIVE}

	controller.ServeReq(c)

	if resp.Code != http.StatusOK {
		t.Errorf("Handlerequest returned %v", resp.Code)
	}
}

func TestStaticHtmlHanderRequest(t *testing.T) {

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", config.BLOGG_URL_API, nil)

	controller := StaticHtmlHandler{}
	c := &web.ChainedReq{resp, req, web.ACTIVE}
	controller.ServeReq(c)

	if resp.Code != http.StatusOK {
		t.Errorf("Handlerequest returned %v", resp.Code)
	}
}

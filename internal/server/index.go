package server

import (
	"net/http"
)

func (s *Server) indexHandler(resp http.ResponseWriter, req *http.Request) {
	bag := gameBag{}

	template := "index.gohtml"
	renderHtml(resp, http.StatusOK, template, bag)
}

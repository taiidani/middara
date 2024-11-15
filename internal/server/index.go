package server

import (
	"net/http"
)

type indexBag struct {
}

func (s *Server) indexHandler(resp http.ResponseWriter, req *http.Request) {
	bag := indexBag{}

	template := "index.gohtml"
	renderHtml(resp, http.StatusOK, template, bag)
}

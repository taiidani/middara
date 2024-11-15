package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

type errorBag struct {
	Message error
}

var (
	errGameIDRequired = fmt.Errorf("game ID is required")
	errGameNotFound   = fmt.Errorf("game not found")
)

func errorResponse(writer http.ResponseWriter, code int, err error) {
	data := errorBag{
		Message: err,
	}

	slog.Error("Displaying error page", "error", err)
	renderHtml(writer, code, "error.gohtml", data)
}

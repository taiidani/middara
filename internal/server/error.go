package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/taiidani/middara/internal/models"
)

var (
	errGameIDRequired   = fmt.Errorf("game ID is required")
	errInvalidGame      = fmt.Errorf("invalid game data")
	errInvalidCharacter = fmt.Errorf("invalid character data")
)

func errorResponse(writer http.ResponseWriter, code int, err error) {
	type errorBag struct {
		Message error
		Game    *models.Game
	}

	data := errorBag{
		Message: err,
	}

	slog.Error("Displaying error page", "error", err)
	renderHtml(writer, code, "error.gohtml", data)
}

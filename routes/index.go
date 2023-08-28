package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", time.Now().String())
}

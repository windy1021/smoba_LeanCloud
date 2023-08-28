package routes

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/leancloud/go-sdk/leancloud"
)

type Todo struct {
	leancloud.Object
	Content string `json:"content"`
}

var Client *leancloud.Client

func init() {
	Client = leancloud.NewEnvClient()
}

func GetTodos(c echo.Context) error {
	todos := []Todo{}
	if err := Client.Class("Todo").NewQuery().Order("createdAt").Find(&todos); err != nil {
		if !strings.Contains(err.Error(), "101") {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	return c.Render(http.StatusOK, "todos", struct {
		Title string
		Todos []Todo
	}{
		Title: "Todo List",
		Todos: todos,
	})
}

func PostTodos(c echo.Context) error {
	content := c.FormValue("content")
	todo := Todo{
		Content: content,
	}

	if _, err := Client.Class("Todo").Create(todo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.Redirect(http.StatusSeeOther, "/todos")
}

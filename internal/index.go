package internal

import "github.com/labstack/echo/v4"

func indexView(c echo.Context) error {
	return c.Render(200, "index", nil)
}

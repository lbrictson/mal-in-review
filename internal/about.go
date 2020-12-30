package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func aboutView(c echo.Context) error {
	return c.Render(200, "about", nil)
}

func aboutForm(c echo.Context) error {
	type Payload struct {
		Email    string
		Question string
	}
	data := Payload{}
	c.Bind(&data)
	if data.Email != "" {
		if data.Question != "" {
			file, _ := json.MarshalIndent(data, "", " ")
			_ = ioutil.WriteFile(fmt.Sprintf("question-%v.json", time.Now().UnixNano()), file, 0644)
		}
	}
	return c.Redirect(http.StatusSeeOther, "/thanks")
}

func thanksView(c echo.Context) error {
	return c.Render(200, "thanks", nil)
}

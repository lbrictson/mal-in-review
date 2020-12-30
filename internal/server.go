package internal

import (
	"fmt"
	"html/template"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Run() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	connectLeaderDB()
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("public/web/templates/*.tmpl")),
	}
	e.Renderer = t
	// Add some better defaults for the http server
	e.Server.IdleTimeout = 30 * time.Second
	e.Server.ReadTimeout = 15 * time.Second
	e.Server.ReadHeaderTimeout = 10 * time.Second
	e.HideBanner = true
	e.Static("/static", "public/web/static")
	e.GET("/", indexView)
	e.GET("/review/:user", singleUserView)
	e.POST("/form/single", singleUserForm)
	e.GET("/battle", battleView)
	e.POST("/form/double", battleForm)
	e.GET("/battle/:user/:user2", battleResults)
	e.GET("/leaderboard", leaderboardView)
	e.GET("/faq", faqView)
	e.GET("/about", aboutView)
	e.POST("/form/question", aboutForm)
	e.GET("/thanks", thanksView)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", conf.Port)))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if data == nil {
		data = map[string]interface{}{"Error": "", "Message": ""}
	}
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		fmt.Printf("error rendering template %v with data payload %v - generated error: %v\n", name, data, err)
	}
	return err
}

func faqView(c echo.Context) error {
	return c.Render(200, "faq", nil)
}

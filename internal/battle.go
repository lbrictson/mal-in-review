package internal

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func battleView(c echo.Context) error {
	return c.Render(http.StatusOK, "battle", map[string]interface{}{})
}

func battleForm(c echo.Context) error {
	type Payload struct {
		Username  string
		OtherUser string
	}
	data := Payload{}
	err := c.Bind(&data)
	if err != nil {
		logrus.Error(err)
		return c.Redirect(http.StatusBadRequest, "/")
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/battle/%v/%v", data.Username, data.OtherUser))
}

func battleResults(c echo.Context) error {
	username := c.Param("user")
	username2 := c.Param("user2")
	if username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	if username2 == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	logrus.Infof("Generating battle report for %v vs %v", username, username2)
	// TODO handle this error gracefully
	data, _ := GetAll(username)
	data2, _ := GetAll(username2)
	stats := generateUserStats(username, data)
	stats2 := generateUserStats(username2, data2)
	return c.Render(http.StatusOK, "double", map[string]interface{}{
		"user":          username,
		"sumMovie":      stats.SumMovie,
		"sumTV":         stats.SumTV,
		"sumOVA":        stats.SumOVA,
		"movieWatched":  stats.MovieWatched,
		"TVWatched":     stats.TVWatched,
		"OVAWatched":    stats.OVAWatched,
		"history":       stats.History,
		"topTV":         stats.TopTV,
		"topOVA":        stats.TopOVA,
		"topMovie":      stats.TopMovie,
		"Rank":          stats.Rank,
		"rawOVA":        stats.RawOVA,
		"rawMovie":      stats.RawMovie,
		"rawTV":         stats.RawTV,
		"user2":         username2,
		"sumMovie2":     stats2.SumMovie,
		"sumTV2":        stats2.SumTV,
		"sumOVA2":       stats2.SumOVA,
		"movieWatched2": stats2.MovieWatched,
		"TVWatched2":    stats2.TVWatched,
		"OVAWatched2":   stats2.OVAWatched,
		"history2":      stats2.History,
		"topTV2":        stats2.TopTV,
		"topOVA2":       stats2.TopOVA,
		"topMovie2":     stats2.TopMovie,
		"Rank2":         stats2.Rank,
		"rawOVA2":       stats2.RawOVA,
		"rawMovie2":     stats2.RawMovie,
		"rawTV2":        stats2.RawTV,
	})
}

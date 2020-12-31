package internal

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

type Historical struct {
	JanMovies       int
	FebMovies       int
	MarchMovies     int
	AprilMovies     int
	MayMovies       int
	JuneMovies      int
	JulyMovies      int
	AugustMovies    int
	SeptemberMovies int
	OctoberMovies   int
	NovemberMovies  int
	DecemberMovies  int
	JanTV           int
	FebTV           int
	MarchTV         int
	AprilTV         int
	MayTV           int
	JuneTV          int
	JulyTV          int
	AugustTV        int
	SeptemberTV     int
	OctoberTV       int
	NovemberTV      int
	DecemberTV      int
	JanOVA          int
	FebOVA          int
	MarchOVA        int
	AprilOVA        int
	MayOVA          int
	JuneOVA         int
	JulyOVA         int
	AugustOVA       int
	SeptemberOVA    int
	OctoberOVA      int
	NovemberOVA     int
	DecemberOVA     int
}

func singleUserForm(c echo.Context) error {
	type Payload struct {
		Username string
	}
	data := Payload{}
	err := c.Bind(&data)
	if err != nil {
		logrus.Error(err)
		return c.Redirect(http.StatusBadRequest, "/")
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/review/%v", data.Username))
}

func singleUserView(c echo.Context) error {
	//movieRank := []RankedItem{}
	username := c.Param("user")
	if username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	logrus.Infof("Generating single report for %v", username)
	// TODO handle this error gracefully
	data, _ := GetAll(username)
	stats := generateUserStats(username, data)
	return c.Render(http.StatusOK, "single", map[string]interface{}{
		"user": username,
		//"current":      data.Current,
		//"finished":     data.Finished,
		"sumMovie":     stats.SumMovie,
		"sumTV":        stats.SumTV,
		"sumOVA":       stats.SumOVA,
		"movieWatched": stats.MovieWatched,
		"TVWatched":    stats.TVWatched,
		"OVAWatched":   stats.OVAWatched,
		"history":      stats.History,
		"topTV":        stats.TopTV,
		"topOVA":       stats.TopOVA,
		"topMovie":     stats.TopMovie,
		"Rank":         stats.Rank,
		"rawOVA":       stats.RawOVA,
		"rawMovie":     stats.RawMovie,
		"rawTV":        stats.RawTV,
	})
}

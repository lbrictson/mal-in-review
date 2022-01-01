package internal

import (
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"strings"
)

type UserStats struct {
	History      Historical
	TopMovie     []FinishedWatchingResponseItem
	TopTV        []FinishedWatchingResponseItem
	TopOVA       []FinishedWatchingResponseItem
	SumOVA       int
	SumMovie     int
	SumTV        int
	TVWatched    int
	OVAWatched   int
	MovieWatched int
	Username     string
	Rank         int
	Minutes      int
	RawOVA       []OVA
	RawTV        []TV
	RawMovie     []Movie
}

type Movie struct {
	Title     string
	ImageLink string
	AnimeID   int
}

type TV struct {
	Title           string
	ImageLink       string
	TotalEpisodes   int
	WatchedEpisodes int
	AnimeID         int
}

type OVA struct {
	Title           string
	ImageLink       string
	TotalEpisodes   int
	WatchedEpisodes int
	AnimeID         int
}

func generateUserStats(username string, data UserData) UserStats {
	u := UserStats{
		Username: username,
	}
	cleanedUpDuplicateData := UserData{}
	for _, x := range data.Finished {
		found := false
		for _, y := range cleanedUpDuplicateData.Finished {
			if x.AnimeID == y.AnimeID {
				found = true
			}
		}
		if found == false {
			cleanedUpDuplicateData.Finished = append(cleanedUpDuplicateData.Finished, x)
		}
	}
	history := Historical{}
	data.Finished = cleanedUpDuplicateData.Finished
	sort.Slice(data.Finished, func(i, j int) bool {
		return data.Finished[i].Score > data.Finished[j].Score
	})
	movies := []FinishedWatchingResponseItem{}
	tvs := []FinishedWatchingResponseItem{}
	ovas := []FinishedWatchingResponseItem{}
	usNumbering := true
	// Determine if non-us numbering
	for _, x := range data.Finished {
		nums := strings.Split(x.FinishDateString, "-")[0]
		number, err := strconv.Atoi(nums)
		if err == nil {
			if number > 12 {
				logrus.Info(number)
				logrus.Infof("user %v is non US numbering", username)
				usNumbering = false
			}
		}
	}
	for _, x := range data.Finished {
		if len(movies) == 5 {
			break
		}
		if x.AnimeMediaTypeString == "Movie" {
			if x.Score != 0 {
				movies = append(movies, x)
			}
		}
	}
	for _, x := range data.Finished {
		if len(tvs) == 5 {
			break
		}
		if x.AnimeMediaTypeString == "TV" {
			if x.Score != 0 {
				tvs = append(tvs, x)
			}
		}
	}
	for _, x := range data.Finished {
		if len(ovas) == 5 {
			break
		}
		if x.AnimeMediaTypeString == "OVA" {
			if x.Score != 0 {
				ovas = append(ovas, x)
			}
		}
		if x.AnimeMediaTypeString == "ONA" {
			if x.Score != 0 {
				ovas = append(ovas, x)
			}
		}
	}
	sumOVA := 0
	sumMovie := 0
	sumTV := 0
	TVWatched := 0
	OVAWatched := 0
	MovieWatched := 0
	rawMovie := []Movie{}
	rawTV := []TV{}
	rawOVA := []OVA{}
	monthSlice := 0
	if !usNumbering {
		monthSlice = 1
	}
	for _, x := range data.Finished {
		switch x.AnimeMediaTypeString {
		case "TV":
			TVWatched = TVWatched + 1
			switch strings.Split(x.FinishDateString, "-")[monthSlice] {
			case "01":
				history.JanTV = history.JanTV + x.NumWatchedEpisodes
			case "02":
				history.FebTV = history.FebTV + x.NumWatchedEpisodes
			case "03":
				history.MarchTV = history.MarchTV + x.NumWatchedEpisodes
			case "04":
				history.AprilTV = history.AprilTV + x.NumWatchedEpisodes
			case "05":
				history.MayTV = history.MayTV + x.NumWatchedEpisodes
			case "06":
				history.JuneTV = history.JuneTV + x.NumWatchedEpisodes
			case "07":
				history.JulyTV = history.JulyTV + x.NumWatchedEpisodes
			case "08":
				history.AugustTV = history.AugustTV + x.NumWatchedEpisodes
			case "09":
				history.SeptemberTV = history.SeptemberTV + x.NumWatchedEpisodes
			case "10":
				history.OctoberTV = history.OctoberTV + x.NumWatchedEpisodes
			case "11":
				history.NovemberTV = history.NovemberTV + x.NumWatchedEpisodes
			case "12":
				history.DecemberTV = history.DecemberTV + x.NumWatchedEpisodes
			}
		case "OVA":
			OVAWatched = OVAWatched + 1
			switch strings.Split(x.FinishDateString, "-")[monthSlice] {
			case "01":
				history.JanOVA = history.JanOVA + x.NumWatchedEpisodes
			case "02":
				history.FebOVA = history.FebOVA + x.NumWatchedEpisodes
			case "03":
				history.MarchOVA = history.MarchOVA + x.NumWatchedEpisodes
			case "04":
				history.AprilOVA = history.AprilOVA + x.NumWatchedEpisodes
			case "05":
				history.MayOVA = history.MayOVA + x.NumWatchedEpisodes
			case "06":
				history.JuneOVA = history.JuneOVA + x.NumWatchedEpisodes
			case "07":
				history.JulyOVA = history.JulyOVA + x.NumWatchedEpisodes
			case "08":
				history.AugustOVA = history.AugustOVA + x.NumWatchedEpisodes
			case "09":
				history.SeptemberOVA = history.SeptemberOVA + x.NumWatchedEpisodes
			case "10":
				history.OctoberOVA = history.OctoberOVA + x.NumWatchedEpisodes
			case "11":
				history.NovemberOVA = history.NovemberOVA + x.NumWatchedEpisodes
			case "12":
				history.DecemberOVA = history.DecemberOVA + x.NumWatchedEpisodes
			}
		case "Movie":
			MovieWatched = MovieWatched + 1
			switch strings.Split(x.FinishDateString, "-")[monthSlice] {
			case "01":
				history.JanMovies = history.JanMovies + 1
			case "02":
				history.FebMovies = history.FebMovies + 1
			case "03":
				history.MarchMovies = history.MarchMovies + 1
			case "04":
				history.AprilMovies = history.AprilMovies + 1
			case "05":
				history.MayMovies = history.MayMovies + 1
			case "06":
				history.JuneMovies = history.JuneMovies + 1
			case "07":
				history.JulyMovies = history.JulyMovies + 1
			case "08":
				history.AugustMovies = history.AugustMovies + 1
			case "09":
				history.SeptemberMovies = history.SeptemberMovies + 1
			case "10":
				history.OctoberMovies = history.OctoberMovies + 1
			case "11":
				history.NovemberMovies = history.NovemberMovies + 1
			case "12":
				history.DecemberMovies = history.DecemberMovies + 1
			}
		case "ONA":
			OVAWatched = OVAWatched + 1
			switch strings.Split(x.FinishDateString, "-")[monthSlice] {
			case "01":
				history.JanOVA = history.JanOVA + x.NumWatchedEpisodes
			case "02":
				history.FebOVA = history.FebOVA + x.NumWatchedEpisodes
			case "03":
				history.MarchOVA = history.MarchOVA + x.NumWatchedEpisodes
			case "04":
				history.AprilOVA = history.AprilOVA + x.NumWatchedEpisodes
			case "05":
				history.MayOVA = history.MayOVA + x.NumWatchedEpisodes
			case "06":
				history.JuneOVA = history.JuneOVA + x.NumWatchedEpisodes
			case "07":
				history.JulyOVA = history.JulyOVA + x.NumWatchedEpisodes
			case "08":
				history.AugustOVA = history.AugustOVA + x.NumWatchedEpisodes
			case "09":
				history.SeptemberOVA = history.SeptemberOVA + x.NumWatchedEpisodes
			case "10":
				history.OctoberOVA = history.OctoberOVA + x.NumWatchedEpisodes
			case "11":
				history.NovemberOVA = history.NovemberOVA + x.NumWatchedEpisodes
			case "12":
				history.DecemberOVA = history.DecemberOVA + x.NumWatchedEpisodes
			}
		}
	}
	for _, x := range data.Current {
		switch x.AnimeMediaTypeString {
		case "TV":
			TVWatched = TVWatched + 1
		case "OVA":
			OVAWatched = OVAWatched + 1
		case "Movie":
			MovieWatched = MovieWatched + 1
		case "ONA":
			OVAWatched = OVAWatched + 1
		}
	}
	for _, x := range data.Finished {
		switch x.AnimeMediaTypeString {
		case "TV":
			totalEps := x.AnimeNumEpisodes
			if totalEps == 0 {
				totalEps = x.NumWatchedEpisodes
			}
			found  := false
			for _, existing := range rawTV {
				if existing.AnimeID == x.AnimeID {
					found = true
				}
			}
			if !found {
				rawTV = append(rawTV, TV{
					Title:           x.AnimeTitle,
					ImageLink:       x.AnimeImagePath,
					TotalEpisodes:   totalEps,
					WatchedEpisodes: x.NumWatchedEpisodes,
					AnimeID:         x.AnimeID,
				})
				sumTV = sumTV + x.NumWatchedEpisodes
			}
		case "OVA":
			totalEps := x.AnimeNumEpisodes
			if totalEps == 0 {
				totalEps = x.NumWatchedEpisodes
			}
			found  := false
			for _, existing := range rawOVA {
				if existing.AnimeID == x.AnimeID {
					found = true
				}
			}
			if !found {
				rawOVA = append(rawOVA, OVA{
					Title:           x.AnimeTitle,
					ImageLink:       x.AnimeImagePath,
					TotalEpisodes:   totalEps,
					WatchedEpisodes: x.NumWatchedEpisodes,
					AnimeID:         x.AnimeID,
				})
				sumOVA = sumOVA + x.NumWatchedEpisodes
			}

		case "Movie":
			found  := false
			for _, existing := range rawMovie {
				if existing.AnimeID == x.AnimeID {
					found = true
				}
			}
			if !found {
				rawMovie = append(rawMovie, Movie{
					Title:     x.AnimeTitle,
					ImageLink: x.AnimeImagePath,
					AnimeID:   x.AnimeID,
				})
				sumMovie = sumMovie + x.NumWatchedEpisodes
			}

		case "ONA":
			totalEps := x.AnimeNumEpisodes
			if totalEps == 0 {
				totalEps = x.NumWatchedEpisodes
			}
			found  := false
			for _, existing := range rawOVA {
				if existing.AnimeID == x.AnimeID {
					found = true
				}
			}
			if !found {
				rawOVA = append(rawOVA, OVA{
					Title:           x.AnimeTitle,
					ImageLink:       x.AnimeImagePath,
					TotalEpisodes:   totalEps,
					WatchedEpisodes: x.NumWatchedEpisodes,
					AnimeID:         x.AnimeID,
				})
				sumOVA = sumOVA + x.NumWatchedEpisodes
			}
		}
	}
	for _, x := range data.Current {
		switch x.AnimeMediaTypeString {
		case "TV":
			totalEps := x.AnimeNumEpisodes
			if totalEps == 0 {
				totalEps = x.NumWatchedEpisodes
			}
			found := false
			for _, existing := range rawTV {
				if existing.AnimeID == x.AnimeID {
					found = true
				}
			}
			if !found {
				rawTV = append(rawTV, TV{
					Title:           x.AnimeTitle,
					ImageLink:       x.AnimeImagePath,
					TotalEpisodes:   totalEps,
					WatchedEpisodes: x.NumWatchedEpisodes,
					AnimeID:         x.AnimeID,
				})
				sumTV = sumTV + x.NumWatchedEpisodes
			}

		case "OVA":
			totalEps := x.AnimeNumEpisodes
			if totalEps == 0 {
				totalEps = x.NumWatchedEpisodes
			}
			found := false
			for _, existing := range rawOVA {
				if existing.AnimeID == x.AnimeID {
					found = true
				}
			}
			if !found {
				rawOVA = append(rawOVA, OVA{
					Title:           x.AnimeTitle,
					ImageLink:       x.AnimeImagePath,
					TotalEpisodes:   totalEps,
					WatchedEpisodes: x.NumWatchedEpisodes,
					AnimeID:         x.AnimeID,
				})
				sumOVA = sumOVA + x.NumWatchedEpisodes
			}

		case "Movie":
			found := false
			for _, existing := range rawMovie {
				if existing.AnimeID == x.AnimeID {
					found = true
				}
			}
			if !found {
				rawMovie = append(rawMovie, Movie{
					Title:     x.AnimeTitle,
					ImageLink: x.AnimeImagePath,
					AnimeID:   x.AnimeID,
				})
				sumMovie = sumMovie + x.NumWatchedEpisodes
			}
		case "ONA":
			found := false
			for _, existing := range rawOVA {
				if existing.AnimeID == x.AnimeID {
					found = true
				}
			}
			totalEps := x.AnimeNumEpisodes
			if totalEps == 0 {
				totalEps = x.NumWatchedEpisodes
			}
			if !found {
				rawOVA = append(rawOVA, OVA{
					Title:           x.AnimeTitle,
					ImageLink:       x.AnimeImagePath,
					TotalEpisodes:   totalEps,
					WatchedEpisodes: x.NumWatchedEpisodes,
					AnimeID:         x.AnimeID,
				})
				sumOVA = sumOVA + x.NumWatchedEpisodes
			}

		}
	}
	u.OVAWatched = OVAWatched
	u.SumOVA = sumOVA
	u.TopOVA = ovas
	u.TVWatched = TVWatched
	u.SumTV = sumTV
	u.TopTV = tvs
	u.MovieWatched = MovieWatched
	u.SumMovie = sumMovie
	u.TopMovie = movies
	u.History = history
	cleanTV := []TV{}
	cleanOVA := []OVA{}
	cleanMovie := []Movie{}
	for _, x := range rawTV {
		found := false
		for _, y := range cleanTV {
			if x.AnimeID == y.AnimeID {
				found = true
			}
		}
		if found == false {
			cleanTV = append(cleanTV, x)
		}
	}
	for _, x := range rawMovie {
		found := false
		for _, y := range cleanMovie {
			if x.AnimeID == y.AnimeID {
				found = true
			}
		}
		if found == false {
			cleanMovie = append(cleanMovie, x)
		}
	}
	for _, x := range rawOVA {
		found := false
		for _, y := range cleanOVA {
			if x.AnimeID == y.AnimeID {
				found = true
			}
		}
		if found == false {
			cleanOVA = append(cleanOVA, x)
		}
	}
	u.RawTV = cleanTV
	u.RawMovie = cleanMovie
	u.RawOVA = cleanOVA
	score := 0
	score = score + (sumMovie * 140)
	score = score + (sumOVA * 24)
	score = score + (sumTV * 24)
	u.Minutes = score
	l := LeaderStat{
		Username: username,
		Score:    score,
	}
	if u.Minutes != 0 {
		l.Store()
		u.Rank = getRank(u.Username)
	}
	return u
}

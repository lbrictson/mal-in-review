package internal

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const animeDateString string = "01-02-06"
const thisYear string = "01-01-21"
const thisYearEuro string = "02-01-06"

type CurrentlyWatchingResponseItem struct {
	Status                int         `json:"status"`
	Score                 int         `json:"score"`
	Tags                  string      `json:"tags"`
	IsRewatching          int         `json:"is_rewatching"`
	NumWatchedEpisodes    int         `json:"num_watched_episodes"`
	AnimeTitle            string      `json:"anime_title"`
	AnimeNumEpisodes      int         `json:"anime_num_episodes"`
	AnimeAiringStatus     int         `json:"anime_airing_status"`
	AnimeID               int         `json:"anime_id"`
	AnimeStudios          interface{} `json:"anime_studios"`
	AnimeLicensors        interface{} `json:"anime_licensors"`
	AnimeSeason           interface{} `json:"anime_season"`
	HasEpisodeVideo       bool        `json:"has_episode_video"`
	HasPromotionVideo     bool        `json:"has_promotion_video"`
	HasVideo              bool        `json:"has_video"`
	VideoURL              string      `json:"video_url"`
	AnimeURL              string      `json:"anime_url"`
	AnimeImagePath        string      `json:"anime_image_path"`
	IsAddedToList         bool        `json:"is_added_to_list"`
	AnimeMediaTypeString  string      `json:"anime_media_type_string"`
	AnimeMpaaRatingString string      `json:"anime_mpaa_rating_string"`
	StartDateString       string      `json:"start_date_string"`
	FinishDateString      interface{} `json:"finish_date_string"`
	AnimeStartDateString  string      `json:"anime_start_date_string"`
	AnimeEndDateString    string      `json:"anime_end_date_string"`
	DaysString            int         `json:"days_string"`
	StorageString         string      `json:"storage_string"`
	PriorityString        string      `json:"priority_string"`
}

type FinishedWatchingResponseItem struct {
	Status                int         `json:"status"`
	Score                 int         `json:"score"`
	Tags                  string      `json:"tags"`
	IsRewatching          int         `json:"is_rewatching"`
	NumWatchedEpisodes    int         `json:"num_watched_episodes"`
	AnimeTitle            string      `json:"anime_title"`
	AnimeNumEpisodes      int         `json:"anime_num_episodes"`
	AnimeAiringStatus     int         `json:"anime_airing_status"`
	AnimeID               int         `json:"anime_id"`
	AnimeStudios          interface{} `json:"anime_studios"`
	AnimeLicensors        interface{} `json:"anime_licensors"`
	AnimeSeason           interface{} `json:"anime_season"`
	HasEpisodeVideo       bool        `json:"has_episode_video"`
	HasPromotionVideo     bool        `json:"has_promotion_video"`
	HasVideo              bool        `json:"has_video"`
	VideoURL              string      `json:"video_url"`
	AnimeURL              string      `json:"anime_url"`
	AnimeImagePath        string      `json:"anime_image_path"`
	IsAddedToList         bool        `json:"is_added_to_list"`
	AnimeMediaTypeString  string      `json:"anime_media_type_string"`
	AnimeMpaaRatingString string      `json:"anime_mpaa_rating_string"`
	StartDateString       string      `json:"start_date_string"`
	FinishDateString      string      `json:"finish_date_string"`
	AnimeStartDateString  string      `json:"anime_start_date_string"`
	AnimeEndDateString    string      `json:"anime_end_date_string"`
	DaysString            int         `json:"days_string"`
	StorageString         string      `json:"storage_string"`
	PriorityString        string      `json:"priority_string"`
	FinishedWatchingDate  time.Time   `json:"-"`
}

type UserData struct {
	Current  []CurrentlyWatchingResponseItem
	Finished []FinishedWatchingResponseItem
}

func GetAll(username string) (UserData, error) {
	thisYearTimed, _ := time.Parse(animeDateString, thisYear)
	current, _ := getCurrentlyWatching(username)
	finished, _ := getFinishedWatching(username)
	finishedWatchingThisYear := []FinishedWatchingResponseItem{}
	usNumbering := true
	logrus.Infof("Trying to determine if %v is us or euro", username)
	// finish_date_string: "18-08-21",
	for _, x := range finished {
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
	logrus.Infof("%v is USnumber %v", username, usNumbering)
	parseString := ""
	if usNumbering == false {
			parseString = thisYearEuro
		} else {
			parseString = animeDateString
	}
	for i, x := range finished {
		t, err := time.Parse(parseString, x.FinishDateString)
		if err == nil {
			finished[i].FinishedWatchingDate = t
			if finished[i].FinishedWatchingDate.UnixNano() >= thisYearTimed.UnixNano() {
				finishedWatchingThisYear = append(finishedWatchingThisYear, finished[i])
			}
		}
	}
	data := UserData{
		Current:  current,
		Finished: finishedWatchingThisYear,
	}
	return data, nil
}

func getCurrentlyWatching(username string) ([]CurrentlyWatchingResponseItem, error) {
	respData := []CurrentlyWatchingResponseItem{}
	url := "https://myanimelist.net/animelist/" + username + "/load.json?status=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respData, err
	}
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return respData, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return respData, err
	}
	err = json.Unmarshal(body, &respData)
	return respData, err
}

func getFinishedWatching(username string) ([]FinishedWatchingResponseItem, error) {
	respData := []FinishedWatchingResponseItem{}
	url := "https://myanimelist.net/animelist/" + username + "/load.json?status=2"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respData, err
	}
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return respData, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return respData, err
	}
	err = json.Unmarshal(body, &respData)
	more := false
	offset := 300
	if len(respData) == 300 {
		more = true
	}
	for more {
		logrus.Infof("looping to get more user data for %v on offset %v", username, offset)
		loopRespData := []FinishedWatchingResponseItem{}
		url := "https://myanimelist.net/animelist/" + username + fmt.Sprintf("/load.json?offset=%v?status=2", offset)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return respData, err
		}
		req.Header.Add("content-type", "application/json")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return respData, err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return respData, err
		}
		err = json.Unmarshal(body, &loopRespData)
		if err != nil {
			return respData, err
		}
		for _, x := range loopRespData {
			respData = append(respData, x)
		}
		if len(loopRespData) == 300 {
			offset = offset + 300
		} else {
			more = false
			break
		}
	}
	logrus.Infof("%v has %v finished items", username, len(respData))
	return respData, err
}

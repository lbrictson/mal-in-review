package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const animeDateString string = "01-02-06"
const thisYear string = "01-01-21"

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
	for i, x := range finished {
		t, err := time.Parse(animeDateString, x.FinishDateString)
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
	return respData, err
}

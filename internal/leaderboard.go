package internal

import (
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var leaderDB *gorm.DB

type LeaderBoardMember struct {
	Username  string
	ScoreRank int
	Score     int
}

func leaderboardView(c echo.Context) error {
	//leaders := []LeaderStat{}
	// db.Order("age desc, name").Find(&users)
	data := []LeaderBoardMember{}
	leaderDB.Raw("SELECT username,ScoreRank,score FROM ( SELECT Username, Score, RANK () OVER ( ORDER BY score DESC ) ScoreRank FROM leader_stats )").Scan(&data)
	//leaderDB.Order("score desc").Find(&leaders)
	return c.Render(200, "leaderboard", map[string]interface{}{"leaders": data})
}

func connectLeaderDB() {
	db, err := gorm.Open(sqlite.Open("leader.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	leaderDB = db
	leaderDB.AutoMigrate(&LeaderStat{})
	return
}

type LeaderStat struct {
	gorm.Model
	Username string `gorm:"unique,index"`
	Score    int
}

func (l LeaderStat) Store() error {
	l.Username = strings.ToLower(l.Username)
	alreadyStored, _ := fetchLeaderStat(l.Username)
	if alreadyStored.Username == l.Username {
		// User already exists, just needs an update
		// db.Model(&product).Update("Price", 200)
		leaderDB.Model(&alreadyStored).Update("Score", l.Score)
	} else {
		leaderDB.Save(&l)
	}
	return nil
}

func fetchLeaderStat(username string) (LeaderStat, error) {
	username = strings.ToLower(username)
	l := LeaderStat{}
	err := leaderDB.First(&l, "username = ?", username).Error
	return l, err
}

func getRank(username string) int {
	username = strings.ToLower(username)
	//SELECT
	//ScoreRank
	//FROM (
	//	SELECT
	//Username,
	//	Score,
	//	RANK () OVER (
	//	ORDER BY score DESC
	//) ScoreRank
	//FROM
	//leader_stats
	//)
	//WHERE
	//username = "Karhu";
	// SELECT ScoreRank FROM ( SELECT Username, Score, RANK () OVER ( ORDER BY score DESC ) ScoreRank FROM leader_stats ) WHERE username = "?"
	result := []Rank{}
	leaderDB.Raw("SELECT ScoreRank FROM ( SELECT Username, Score, RANK () OVER ( ORDER BY score DESC ) ScoreRank FROM leader_stats ) WHERE username = ?", username).Scan(&result)
	if len(result) == 0 {
		return 0
	} else {
		return result[0].ScoreRank
	}
}

type Rank struct {
	ScoreRank int
}

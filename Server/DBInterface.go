package DBInterface

import (
	"strconv"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Structs

type PlayerAchievements struct {
	UserID int
	TotalWins int
	TotalLosses int
	TotalMatches int
	BombsDropped int
}

type MatchHistory struct {
	UserName1 string
	UserName2 string
	WinnerName string
	MatchLength string

}

// Create functions

func CreateUser(db *sql.DB, user string, password string) {
	_, err := db.Exec("INSERT INTO Users (userName, userPassword) VALUES ('" + user + "', '" + password + "');")
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = db.Exec("INSERT INTO PlayerAchievements (userID, totalMatches, totalWins, totalLosses, bombsDropped) " +
					 "VALUES ((SELECT userID FROM Users WHERE userName = '" + user + "'), 0, 0, 0, 0);")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateMatch(db *sql.DB, userID1 string, userID2 string, winnerID string, matchLength string){
	_, err := db.Exec("INSERT INTO MatchHistory (userID1, userID2, winnerID, matchLength) " +
					 "VALUES (" + userID1 + ", " + userID2 + ", " + winnerID + ", '" + matchLength + "');")
	if err != nil {
		log.Fatal(err)
	}
}

// Delete functions

func DeleteUser(db *sql.DB, user string) {
	_, err := db.Exec("DELETE FROM PlayerAchievements WHERE userID = (SELECT userID FROM Users WHERE userName = '" + user + "');")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM Users WHERE userName = '" + user + "';")
	if err != nil {
		log.Fatal(err)
	}
}

// Update functions

func UpdateUserPassword(db *sql.DB, user string, password string){
	_, err := db.Exec("UPDATE Users SET userPassword = '" + password + "' WHERE userName = '" + user + "';")
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePlayerAchievements(db *sql.DB, user string, totalWins string, totalLosses string, totalMatches string, bombsDropped string){
	_, err := db.Exec("UPDATE PlayerAchievements SET totalWins = " + totalWins + ", totalLosses = " + totalLosses +
					   ", totalMatches = " + totalMatches + ", bombsDropped = " + bombsDropped +
					   " WHERE userID = (SELECT userID FROM Users WHERE userName = '" + user + "');")
	if err != nil {
		log.Fatal(err)
	}
}


// Query functions

func GetUserID(db *sql.DB, userName string) int {
	var userID int
	err := db.QueryRow("SELECT userID FROM Users WHERE userName = '" + userName + "';").Scan(&userID)
	
	if err != nil{
		log.Fatal(err)
	}
	return userID;
}

func GetUserName(db *sql.DB, userID int) string {
	var userName string
	err := db.QueryRow("SELECT userName FROM Users WHERE userID = " + strconv.Itoa(userID) + ";").Scan(&userName)
	
	if err != nil{
		log.Fatal(err)
	}
	return userName
}

func GetUserPassword(db *sql.DB, userName string) string {
	var userPassword string
    err := db.QueryRow("SELECT userPassword FROM Users WHERE userName = '" + userName + "';").Scan(&userPassword)
	
	if err != nil {
		log.Fatal(err)
	}
	return userPassword
}

func GetPlayerAchievements(db *sql.DB, userName string) PlayerAchievements{
	userID := GetUserID(db, userName)

	var playerStats PlayerAchievements
	playerStats.UserID = userID
	err := db.QueryRow("SELECT totalWins, totalLosses, totalMatches, bombsDropped FROM PlayerAchievements " +
					   "WHERE userID = '" + strconv.Itoa(userID) + "';").Scan(&playerStats.TotalWins, &playerStats.TotalLosses, &playerStats.TotalMatches, &playerStats.BombsDropped)
	
	if err != nil {
		log.Fatal(err)
	}
		
	return playerStats				
	
}

func GetPlayerMatchCount(db *sql.DB, userID int) int {
	rows1,err1 := db.Query("SELECT * FROM MatchHistory WHERE userID1 = " + strconv.Itoa(userID) + ";")
	if err1 != nil {
		log.Fatal(err1)
	}
	
	rows2,err2 := db.Query("SELECT * FROM MatchHistory WHERE userID2 = " + strconv.Itoa(userID) + ";")
	if err2 != nil {
		log.Fatal(err2)
	}
	
	var count1 int
	var count2 int
	
	for count1 = 0; rows1.Next(); count1++ {	
	}
	
	for count2 = 0; rows2.Next(); count2++ {	
	}
	
	return count1 + count2
}

func GetPlayerMatchHistory(db *sql.DB, userName string) []MatchHistory {
	var userID = GetUserID(db, userName)
	var playerMatchCount = GetPlayerMatchCount(db, userID)
	var playerMatchHistory = make([]MatchHistory, playerMatchCount)
	
	rows, err := db.Query("SELECT userID1, userID2, winnerID, matchLength FROM MatchHistory WHERE userID1 = " + strconv.Itoa(userID) +
						" OR userID2 = " + strconv.Itoa(userID) + ";")
	if err != nil {
		log.Fatal(err)
	}
	
	for i := 0; rows.Next(); i++ {
		var userID1 int
		var userID2 int
		var winnerID int
		err := rows.Scan(&userID1, &userID2, &winnerID, &playerMatchHistory[i].MatchLength)
		if err != nil {
			log.Fatal(err)
		}
		playerMatchHistory[i].UserName1 = GetUserName(db, userID1)
		playerMatchHistory[i].UserName2 = GetUserName(db, userID2)
		playerMatchHistory[i].WinnerName = GetUserName(db, winnerID)
		
	}
	
	return playerMatchHistory
}


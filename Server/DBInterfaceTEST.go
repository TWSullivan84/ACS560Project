package main
import (
	"ACS560/DBInterface"
	"fmt"
	"strconv"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/bmandb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create two Users
	fmt.Println("Creating User 'Tyler' with password 'Hack'\n")
	DBInterface.CreateUser(db, "Tyler", "Hack")
	
	fmt.Println("Creating User 'Eric' with password 'ThePlanet'\n")
	DBInterface.CreateUser(db, "Eric", "ThePlanet")
	
	fmt.Println("Creating a match!\n")
	DBInterface.CreateMatch(db,"1", "2", "2", "135")
	
	fmt.Println("Creating a match!\n")
	DBInterface.CreateMatch(db,"2", "1", "1", "140")
	
	fmt.Println("Getting match history for User 'Tyler'")
	history := DBInterface.GetPlayerMatchHistory(db, "Tyler")
	for i := range history{
		fmt.Printf("%s, %s, %s, %s\n", history[i].UserName1, history[i].UserName2, history[i].WinnerName, history[i].MatchLength)
	}

	fmt.Println("\nGetting userID for User 'Tyler'")
	userID := DBInterface.GetUserID(db, "Tyler")
	fmt.Println(strconv.Itoa(userID))
	
	fmt.Println("\nChanging password to 'Apple' for User 'Tyler'\n")
	DBInterface.UpdateUserPassword(db, "Tyler", "Apple")
	
	fmt.Println("Getting userPassword for User 'Tyler'")
	userPassword := DBInterface.GetUserPassword(db, "Tyler")
	fmt.Println(userPassword)
	
	fmt.Println("\nUpdating Player Achievements to make User 'Tyler' appear legit.\n")
	DBInterface.UpdatePlayerAchievements(db, "Tyler", "5", "2", "7", "37")
	
	fmt.Println("Displaying Player Achievements for User 'Tyler'")
	playerStats := DBInterface.GetPlayerAchievements(db, "Tyler")
	fmt.Printf("UserID: %d, Total Wins: %d  Total Losses: %d Total Matches: %d Bombs Dropped: %d\n", 
				playerStats.UserID, playerStats.TotalWins, playerStats.TotalLosses, playerStats.TotalMatches, playerStats.BombsDropped)
}
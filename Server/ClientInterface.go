package main

import (
	"net"
	"fmt"
	"log"
	"ACS560/DBInterface"
	"strings"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":11337")
	if err != nil{
		log.Fatal(err)
	}
	defer ln.Close()
	
	fmt.Println("Server ready!")
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	fmt.Println("A client has connected!")
	fmt.Println("Reading game results.")
	time.Sleep(2000 * time.Millisecond)
	
	gameData := readGameData(conn)
	splitData := strings.Split(gameData, ",")
	fmt.Println(splitData)
	
	player1Data := splitData[0:3]
	player2Data := splitData[3:6]
	matchData := splitData[6:10]
	
	fmt.Println(player1Data)
	fmt.Println(player2Data)
	fmt.Println(matchData)
	
	updatePlayerAchievements(player1Data)	
	updatePlayerAchievements(player2Data)
	insertMatch(matchData)

	fmt.Println("Done!")
	
	conn.Close()
}

func readGameData(conn net.Conn) string{
	data := make([]byte, 512)
	
	n, err := conn.Read(data)
	if err != nil {
			conn.Close()
	}
	gameData := string(data[:n])
	return gameData
}

func updatePlayerAchievements(playerData []string){
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/bmandb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	userName := playerData[0]
	bombsDropped,_ := strconv.Atoi(playerData[1])
	win,_ := strconv.ParseBool(playerData[2])
	
	PA := DBInterface.GetPlayerAchievements(db, userName)
	PA.BombsDropped += bombsDropped
	if win == true{
		PA.TotalWins += 1
	}else {
		PA.TotalLosses += 1
	}
	PA.TotalMatches += 1

	DBInterface.UpdatePlayerAchievements(db, DBInterface.GetUserName(db,PA.UserID), strconv.Itoa(PA.TotalWins), strconv.Itoa(PA.TotalLosses),
							strconv.Itoa(PA.TotalMatches), strconv.Itoa(PA.BombsDropped))
}

func insertMatch(matchData []string){
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/bmandb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	player1 := matchData[0]
	player2 := matchData[1]
	winner := matchData[2]
	matchLength := matchData[3]
	
	DBInterface.CreateMatch(db, strconv.Itoa(DBInterface.GetUserID(db, player1)), strconv.Itoa(DBInterface.GetUserID(db, player2)),
								strconv.Itoa(DBInterface.GetUserID(db, winner)), matchLength)
}


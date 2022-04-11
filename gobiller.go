package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Biller struct {
	Username     string
	CompleteName string
	Designation  string
}

func getAllBillersName(db *sql.DB) []Biller {
	var result Biller
	allBillers := []Biller{}
	rows, err := db.Query("SELECT * FROM BILLER")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&result)
		allBillers = append(allBillers, result)
	}
	return allBillers
}

func postgresqlConnector(channel chan *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	// !<<<<<<<  HIGHLY SECURE PURPOSE ONLY >>>>>>>>>>
	fmt.Print("Enter the username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter the password: ")
	scanner.Scan()
	password := scanner.Text()

	fmt.Print("Enter the database name: ")
	scanner.Scan()
	database := scanner.Text()

	username = "vivekvijayan"
	password = ""
	database = "houserenew"
	fmt.Println("Connecting to Postgresql data ...")
	connectionString := "postgresql://" + username + ":" + password + "@127.0.0.1:5432/" + database
	db, err := sql.Open("postgres", connectionString)

	if err == nil {
		panic(err)
	}
	channel <- db
	close(channel)
}

func main() {
	fmt.Println("📝 Go Biller Application - Version : 1.0.0")
	databaseChannel := make(chan *sql.DB)
	go postgresqlConnector(databaseChannel)
	fmt.Println("Getting the key ... 🚀")
	databaseKey, err := <-databaseChannel
	if !err {
		fmt.Println("Success")
	}
	fmt.Println("Collected key to the database ✅")
	getAllBillersName(databaseKey)
}

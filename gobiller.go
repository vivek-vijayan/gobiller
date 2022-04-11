package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
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
		log.Fatal(err.Error)
	}
	for rows.Next() {
		rows.Scan(&result)
		allBillers = append(allBillers, result)
	}
	return allBillers
}

func postgresqlConnector(wg sync.WaitGroup) *sql.DB {
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

	fmt.Println("Connecting to Postgresql data ...")
	connectionString := "postgresql://" + username + ":" + password + "@127.0.0.1:5432/" + database
	db, err := sql.Open("postgres", connectionString)
	defer wg.Done()
	if err == nil {
		panic(err)
	}
	return db
}

func main() {
	wg := sync.WaitGroup{}

	fmt.Println("📝 Go Biller Application - Version : 1.0.0")
	wg.Add(1)

	databaseChannel := make(chan *sql.DB)
	go postgresqlConnector(wg)
	fmt.Println("Getting the key ... 🚀")
	databaseKey := <-databaseChannel
	fmt.Println("Collected key to the database ✅")
}

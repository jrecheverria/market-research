package main

import (
	"database/sql"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
	"github.com/jrecheverria/market-research/cmd/newsflow"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	db, err := openDb()
	if err != nil {
		log.Fatalf("Failed to establish database connection pool due to %v\n", err)
	}
	defer db.Close()

	finnhubAPIKey := os.Getenv("FINNHUB_KEY")
	// TODO: Probably need to externalize this so that its available everywhere
	finnhubConfig := finnhub.NewConfiguration()
	finnhubConfig.AddDefaultHeader("X-Finnhub-Token", finnhubAPIKey)
	finnhubClient := finnhub.NewAPIClient(finnhubConfig).DefaultApi

	newsflow.CompanyNews("ASTS", finnhubClient)
}

func openDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "market-research")

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

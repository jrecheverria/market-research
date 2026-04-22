package main

import (
	"database/sql"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jrecheverria/market-research/cmd/news"
	"github.com/jrecheverria/market-research/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type application struct {
	articles *models.ArticleModel
}

func main() {
	db, err := openDb()
	if err != nil {
		log.Fatalf("Failed to establish database connection pool due to %v\n", err)
	}
	defer db.Close()
	m, err := migrate.New("file://migrations", "sqlite3://data.db")

	app := &application{
		articles: &models.ArticleModel{DB: db},
	}

	// Establishing finnhub client
	finnhubAPIKey := os.Getenv("FINNHUB_KEY")
	finnhubConfig := finnhub.NewConfiguration()
	finnhubConfig.AddDefaultHeader("X-Finnhub-Token", finnhubAPIKey)
	finnhubClient := finnhub.NewAPIClient(finnhubConfig).DefaultApi

	news.CompanyNews("ASTS", finnhubClient)
}

func openDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "market-research")

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

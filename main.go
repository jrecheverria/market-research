package main

import (
	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
	"github.com/jrecheverria/market-research/cmd/newsflow"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello to my server"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	finnhubAPIKey := os.Getenv("FINNHUB_KEY")
	// TODO: Probably need to externalize this so that its available everywhere
	finnhubConfig := finnhub.NewConfiguration()
	finnhubConfig.AddDefaultHeader("X-Finnhub-Token", finnhubAPIKey)
	finnhubClient := finnhub.NewAPIClient(finnhubConfig).DefaultApi

	newsflow.CompanyNews("ASTS", finnhubClient)
}

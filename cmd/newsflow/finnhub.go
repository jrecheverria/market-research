package newsflow

import (
	"context"
	"fmt"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
)

// TODO: Need to add proper error handling and retries to these methods
func CompanyNews(ticker string, finnhubClient *finnhub.DefaultApiService) {
	companyNews, _, _ := finnhubClient.CompanyNews(context.Background(), ticker, "2026-04-15", "2026-04-16")
	fmt.Printf("company news data %v\n", companyNews)
}

func Quote()                    {}
func EarningsCalendar()         {}
func StockInsiderTransactions() {}

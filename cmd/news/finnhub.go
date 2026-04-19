package newsflow

import (
	"context"
	"fmt"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
)

func CompanyNews(ticker string, finnhubClient *finnhub.DefaultApiService) {
	companyNews, _, _ := finnhubClient.CompanyNews(context.Background(), ticker, "2026-04-15", "2026-04-16")
	fmt.Printf("company news data %v\n", companyNews)
}

func MarketNews(category string, finnhubClient *finnhub.DefaultApiService) {
	marketNews, _, _ := finnhubClient.GeneralNews(context.Background(), category, nil)
	fmt.Printf("general market news %v\n", marketNews)
}

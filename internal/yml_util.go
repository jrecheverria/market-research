package internal

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type SectorData struct {
	Sectors map[string]Category `yaml:"sectors"`
}

type Category struct {
	DisplayName string   `yaml:"display_name"`
	Description string   `yaml:"description"`
	Tickers     []Ticker `yaml:"tickers"`
}

type Ticker struct {
	Symbol string `yaml:"symbol"`
	Name   string `yaml:"name"`
}

func ReadSector() {
	f, err := os.ReadFile("config/ticker.yaml")
	if err != nil {
		log.Fatal("Unable to read stock data file due to: ", err)
	}

	var sectorData SectorData
	if err := yaml.Unmarshal(f, &sectorData); err != nil {
		log.Fatal("Unable to unmarshall data due to: ", err)
	}

	fmt.Printf("%+v\n", sectorData.Sectors)
}

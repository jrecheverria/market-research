package internal

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Category struct {
	Stocks map[string]string `yaml:"stocks"`
}

func ReadSector() {
	f, err := os.ReadFile("config/ticker.yaml")
	if err != nil {
		log.Fatal("Unable to read stock data file due to: ", err)
	}

	var wrapper struct {
		Sector map[string]Category `yaml:"sector"`
	}
	if err := yaml.Unmarshal(f, &wrapper); err != nil {
		log.Fatal("Unable to unmarshall data due to: ", err)
	}

	fmt.Printf("%+v\n", wrapper.Sector)
}

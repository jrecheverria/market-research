package main

import (
	"fmt"
	"github.com/jrecheverria/market-research/internal"
)

// we'll need a couple of web handlers here to communicate with the front end
func main() {
	fmt.Println("Hello degenerates!")
	internal.ReadSector() 
}

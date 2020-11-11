package main

import (
	"fmt"
	"os"
	"flag"
)

type Input struct {
	baseDate  string
	basePrice int
	numCycle  int
}


func getArgs() {
	t = Input{}
	t.baseDate  := flag.String("date", "dd-mm-yyyy format for date")
	t.basePrice := flag.Int("price", 1000, "An integer which describes the base price")
	t.numCycle  := flag.Int("cycles", 1, "Number of cycles")

	retur
	
}

func main() {
	fmt.Println("This module is to calculate the gann price levels given the proper inputs")
	
}

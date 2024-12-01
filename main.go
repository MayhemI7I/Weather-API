package main

import (
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "User city")
	format := flag.Int("format", 1, "Format output weather")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)

}

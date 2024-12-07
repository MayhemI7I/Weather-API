package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	city := flag.String("city", "", "User city")
	format := flag.Int("format", 1, "Format output weather")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)
	r := strings.NewReader("Maша моя любовь")
	block := make([]byte, 10)
	for {
		_, err := r.Read(block)
		fmt.Printf("%q\n", block)
		if err == io.EOF {
			break
		}

	}

}

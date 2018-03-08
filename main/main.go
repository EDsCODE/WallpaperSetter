package main

import (
	"fmt"
	"github.com/EDsCODE/WallPaperScraper"
	"log"
	"os"
)

func main() {

	if len(os.Args[1:]) == 0 {
		log.Fatal("Subreddit is required")
	}

	_, err := scraper.Get(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s lives in %s.\n", os.Getenv("USER"), os.Getenv("XDG_CURRENT_DESKTOP"))

	// for _, item := range items {
	// 	fmt.Println(item)
	// 	fmt.Println()
	// }
}

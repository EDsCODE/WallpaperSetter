package main

import (
	//	"fmt"
	"log"
	"os"
	"scraper"
)

func main() {

	if len(os.Args[1:]) == 0 {
		log.Fatal("Subreddit is required")
	}

	_, err := reddit.Get(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	// for _, item := range items {
	// 	fmt.Println(item)
	// 	fmt.Println()
	// }
}

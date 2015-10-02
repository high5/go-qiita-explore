package explore_test

import (
	"fmt"
	"github.com/high5/go-qiita-explore"
	"log"
)

func ExampleExplore_GetStocks() {
	explore := explore.NewExplore()
	mainTag := "Ruby"
	articles, err := explore.GetStocks(mainTag)
	if err != nil {
		log.Fatal(err)
	}

	existsTag := true

	for _, article := range articles {
		for _, tag := range article.Tags {
			if tag != mainTag {
				existsTag = false
			} else {
				existsTag = true
				break
			}
		}
	}

	if len(articles) > 0 && existsTag == true {
		fmt.Println("Articles (filtered by Ruby) recieved.")
	} else {
		fmt.Printf("Number of articles recieved: %d (filtered by Ruby %v)", len(articles), existsTag)
	}

	// Output: Articles (filtered by Ruby) recieved.
}

func ExampleExplore_GetItems() {
	explore := explore.NewExplore()
	mainTag := "JavaScript"
	articles, err := explore.GetItems(mainTag, 1)
	if err != nil {
		log.Fatal(err)
	}

	existsTag := true

	for _, article := range articles {
		for _, tag := range article.Tags {
			if tag != mainTag {
				existsTag = false
			} else {
				existsTag = true
				break
			}
		}
	}

	if len(articles) > 0 && existsTag == true {
		fmt.Println("Articles (filtered by JavaScript) recieved.")
	} else {
		fmt.Printf("Number of articles recieved: %d (filtered by JavaScript %v)", len(articles), existsTag)
	}

	// Output: Articles (filtered by JavaScript) recieved.
}

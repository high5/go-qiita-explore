package explore_test

import (
	"fmt"
	"github.com/high5/go-qiita-explore"
	"log"
)

func ExampleExplore_GetStocks() {
	explore := explore.NewExplore()
	articles, err := explore.GetStocks("Ruby")
	if err != nil {
		log.Fatal(err)
	}
	for index, article := range articles {
		if index == 0 {
			fmt.Println("------------------------------------------------------------------------------")
		}
		fmt.Printf("title:    %s \n", article.Title)
		fmt.Printf("url:      %s \n", article.URL.String())
		fmt.Printf("user:     %s \n", article.UserName)
		fmt.Printf("user_url: %s \n", article.UserURL.String())
		fmt.Printf("stock:    %d \n", article.StockCount)
		for index, tag := range article.Tags {
			tagNo := index + 1
			fmt.Printf("tag%d:     %s (http://qiita.com/tags/%s)\n", tagNo, tag, tag)
		}
		fmt.Println("------------------------------------------------------------------------------")
	}

	// Output: Articles (filtered by Ruby) recieved.
}

func ExampleExplore_GetItems() {
	explore := explore.NewExplore()
	articles, err := explore.GetItems("JavaScript", 1)
	if err != nil {
		log.Fatal(err)
	}
	for index, article := range articles {
		if index == 0 {
			fmt.Println("------------------------------------------------------------------------------")
		}
		fmt.Printf("title:    %s \n", article.Title)
		fmt.Printf("url:      %s \n", article.URL.String())
		fmt.Printf("user:     %s \n", article.UserName)
		fmt.Printf("user_url: %s \n", article.UserURL.String())
		fmt.Printf("stock:    %d \n", article.StockCount)
		for index, tag := range article.Tags {
			tagNo := index + 1
			fmt.Printf("tag%d:     %s (http://qiita.com/tags/%s)\n", tagNo, tag, tag)
		}
		fmt.Println("------------------------------------------------------------------------------")
	}

	// Output: Articles (filtered by JavaScript) recieved.
}

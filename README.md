# go-qiita-explore



```

package main

import (
	"fmt"
	"../go-qiita-explore"
	"log"
)

func main() {
  explore := explore.NewExplore()
  //articles, err := explore.GetArticles("Go", "stocks")

	articles, err := explore.GetItems("Go", 1)
	if err != nil {
      log.Fatal(err)
  }

	fmt.Println("------------------------")
	for index, article := range articles {
      no := index + 1
			fmt.Println(no)
      fmt.Println(article.ArticleTitle)
			for i, tag := range article.Tags {
				fmt.Println(i)
			  fmt.Println(tag)
			}
      fmt.Println("------------------------")
  }


	/*
  articles, err := explore.GetArticles("Go", "stocks", 1)
  articles2, err := explore.GetArticles("Go", "stocks", 2)
  articles3, err := explore.GetArticles("Go", "stocks", 3)
  if err != nil {
      log.Fatal(err)
  }
	fmt.Println("------------------------")
  for index, article := range articles {
      no := index + 1
      fmt.Println(no)
      fmt.Println(article.ArticleTitle)
      fmt.Println(article.ArticleURL.String())
      fmt.Println(article.UserName)
      fmt.Println(article.UserURL.String())
      fmt.Println(article.StockCount)
      fmt.Println("------------------------")
  }

	fmt.Println("------------------------")
  for index, article := range articles2 {
      no := index + 1
      fmt.Println(no)
      fmt.Println(article.ArticleTitle)
      fmt.Println(article.ArticleURL.String())
      fmt.Println(article.UserName)
      fmt.Println(article.UserURL.String())
      fmt.Println(article.StockCount)
      fmt.Println("------------------------")
  }

	fmt.Println("------------------------")
  for index, article := range articles3 {
      no := index + 1
      fmt.Println(no)
      fmt.Println(article.ArticleTitle)
      fmt.Println(article.ArticleURL.String())
      fmt.Println(article.UserName)
      fmt.Println(article.UserURL.String())
      fmt.Println(article.StockCount)
      fmt.Println("------------------------")
  }
	*/


}

```

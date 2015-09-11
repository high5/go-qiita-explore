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
  articles, err := explore.GetArticles("go", "stocks")
  if err != nil {
      log.Fatal(err)
  }
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
}

```

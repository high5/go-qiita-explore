# go-qiita-explore

This package were inspired by [andygrunwald/go-trending](https://github.com/andygrunwald/go-trending) (Go)

## Features

* Get newly posted articles
* Get recently stocked articles
* Filtering by tag and page

## Installation

It is go gettable

    $ go get github.com/high5/go-qiita-explore

update

    $ go get -u github.com/high5/go-qiita-explore		

(optional) to run unit / example tests:

    $ cd $GOPATH/src/github.com/high5/go-qiita-explore
    $ go test -v


## Examples

### List recently stocked articles of first page for Ruby

```go
package main

import (
	"fmt"
	"github.com/high5/go-qiita-explore"
	"log"
)

func main() {
	explore := explore.NewExplore()
	articles, err := explore.GetStocks("Ruby") // or explore.GetStocks("Ruby", 1)
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
}
```


### List newly posted articles of second page for JavaScript

```go
package main

import (
	"fmt"
	"github.com/high5/go-qiita-explore"
	"log"
)

func main() {
	explore := explore.NewExplore()
	articles, err := explore.GetItems("JavaScript", 2)
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
}
```


## License

This project is released under the terms of the [MIT license](http://opensource.org/licenses/mit-license.php).

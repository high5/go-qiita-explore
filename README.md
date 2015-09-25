# go-qiita-explore

This package were inspired by [andygrunwald/go-trending](https://github.com/andygrunwald/go-trending) (Go)

## Features

* Get newly posted articles
* Get recently stocked articles
* Filtering by tag and page

## Installation

It is go gettable

    $ go get github.com/high5/go-qiita-explore

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








































# go-trending

[![GoDoc](https://godoc.org/github.com/andygrunwald/go-trending?status.svg)](https://godoc.org/github.com/andygrunwald/go-trending)
[![Build Status](https://travis-ci.org/andygrunwald/go-trending.svg?branch=master)](https://travis-ci.org/andygrunwald/go-trending)
[![Coverage Status](https://coveralls.io/repos/andygrunwald/go-trending/badge.svg?branch=master&service=github)](https://coveralls.io/github/andygrunwald/go-trending?branch=master)

A package to retrieve [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) from Github written in [golang](https://golang.org/).

[![trending package showcase](./img/go-trending-shrinked.png "trending package showcase")](https://raw.githubusercontent.com/andygrunwald/go-trending/master/img/go-trending-shrinked.png)

This package were inspired by [rochefort/git-trend](https://github.com/rochefort/git-trend) (Ruby) and [sheharyarn/github-trending](https://github.com/sheharyarn/github-trending) (Ruby).

## Features

* Get trending repositories
* Get trending developers
* Get trending languages
* Get all programing languages known by GitHub
* Filtering by time and (programing) language

## Installation

It is go gettable

    $ go get github.com/andygrunwald/go-trending

(optional) to run unit / example tests:

    $ cd $GOPATH/src/github.com/andygrunwald/go-trending
    $ go test -v

## API

Please have a look at the [GoDoc documentation](https://godoc.org/github.com/andygrunwald/go-trending) for a detailed API description.

## Examples

Further a few examples how the API can be used.
A few more examples are available in the [GoDoc examples section](https://godoc.org/github.com/andygrunwald/go-trending#pkg-examples).

### List trending repositories of today for all languages

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-trending"
	"log"
)

func main() {
	trend := trending.NewTrending()

	// Show projects of today
	projects, err := trend.GetProjects(trending.TimeToday, "")
	if err != nil {
		log.Fatal(err)
	}
	for index, project := range projects {
		no := index + 1
		if len(project.Language) > 0 {
			fmt.Printf("%d: %s (written in %s with %d ★ )\n", no, project.Name, project.Language, project.Stars)
		} else {
			fmt.Printf("%d: %s (with %d ★ )\n", no, project.Name, project.Stars)
		}
	}
}
```

### List trending repositories of this week for golang

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-trending"
	"log"
)

func main() {
	trend := trending.NewTrending()

	// Show projects of today
	projects, err := trend.GetProjects(trending.TimeWeek, "go")
	if err != nil {
		log.Fatal(err)
	}
	for index, project := range projects {
		no := index + 1
		if len(project.Language) > 0 {
			fmt.Printf("%d: %s (written in %s with %d ★ )\n", no, project.Name, project.Language, project.Stars)
		} else {
			fmt.Printf("%d: %s (with %d ★ )\n", no, project.Name, project.Stars)
		}
	}
}
```

### List trending developers of this month for Swift

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-trending"
	"log"
)

func main() {
	trend := trending.NewTrending()

	developers, err := trend.GetDevelopers(trending.TimeMonth, "swift")
	if err != nil {
		log.Fatal(err)
	}
	for index, developer := range developers {
		no := index + 1
		fmt.Printf("%d: %s (%s)\n", no, developer.DisplayName, developer.FullName)
	}
}
```

### List available languages

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-trending"
	"log"
)

func main() {
	trend := trending.NewTrending()

	// Show languages
	languages, err := trend.GetLanguages()
	if err != nil {
		log.Fatal(err)
	}
	for index, language := range languages {
		no := index + 1
		fmt.Printf("%d: %s (%s)\n", no, language.Name, language.URLName)
	}
}

```

## Implementations

* [sikang99/hub-trend](https://github.com/sikang99/hub-trend/)
* [andygrunwald/TrendingGithub](https://github.com/andygrunwald/TrendingGithub) - [@TrendingGithub](https://twitter.com/TrendingGithub)

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).




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

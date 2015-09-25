package explore

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

// Internal used constants related to qiita`s website / structure.
const (
	baseHost = "http://qiita.com"
	basePath = "/tags/"
)

// Explore reflects the main datastructure of this package.
// It doesn`t provide an exported state, but based on this the methods are called.
// To recieve a new instance just add
//
//		package main
//
//		import (
//			"github.com/high5/go-qiita-explore"
//		)
//
//		func main() {
//			explore := trending.NewExplore()
//			...
//		}
//
type Explore struct {
	document *goquery.Document
}

// Article provides information as printed on the source website http://qiita.com/tags/{tag}/stocks and http://qiita.com/tags/{tag}/items.
// Title is the title of the article. "中規模Web開発のためのMVC分割とレイヤアーキテクチャ"
// Path is the path of the article. "/yuku_t/items/961194a5443b618a4cac"
// URL is the http address of the article reflected as url.URL datastructure like "http://qiita.com/yuku_t/items/961194a5443b618a4cac".
// Tags are a collection of tag.
// UserName is name of the user who posted the article.
// UserPath is path of the user who posted the article.
// User is the http address of the user page of the article reflected as url.URL datastructure like "http://qiita.com/yuku_t".
// StockCount is the number of qiita stock.
type Article struct {
	Title      string
	Path       string
	URL        *url.URL
	Tags       []string
	UserName   string
	UserPath   string
	UserURL    *url.URL
	StockCount int
}

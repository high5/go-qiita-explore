package explore

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
)


const (
	baseHost       = "http://qiita.com"
	basePath       = "/tags/"
)

type Explore struct {
	document *goquery.Document
}


type Article struct {
	ArticleTitle    string
	ArtiClePath     string
	ArticleURL     *url.URL
	Tags           []string
	UserName       string
	UserPath       string
	UserURL        *url.URL
	StockCount     int
}

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
	UserName       string
	UserPath       string
	UserURL        *url.URL
	StockCount     int
}

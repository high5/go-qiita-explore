package explore

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
)

// explore := explore.NewExplore()
// initialize method
func NewExplore() *Explore {
	t := Explore{}
	return &t
}

func (t *Explore) GetStocks(tag string, page ...int) ([]Article, error) {
	var pageQuery int = 1
  if (len(page) > 0) {
		pageQuery = page[0]
	}
	return t.GetArticles(tag, "stocks", pageQuery)
}

func (t *Explore) GetItems(tag string, page ...int) ([]Article, error) {
	var pageQuery int = 1
  if (len(page) > 0) {
		pageQuery = page[0]
	}
	return t.GetArticles(tag, "items", pageQuery)
}

func (t *Explore) GetArticles(tag string, category string, page ...int) ([]Article, error) {
	var articles []Article

	var pageQuery int = 1

  if (len(page) > 0) {
		pageQuery = page[0]
	}

	u, err := t.generateURL(tag, category, pageQuery)
  if err != nil {
		return articles, err
	}

	doc, err := goquery.NewDocument(u.String())
	if err != nil {
		return articles, err
	}

	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		articleTitle := t.trim(s.Find(".publicItem_body a").First().Text())
	  articlePath, articleExists := s.Find(".publicItem_body a").First().Attr("href")
		articleURL := t.appendBaseHostToPath(articlePath, articleExists)

		userName := t.trim(s.Find(".publicItem_status a").Text())
		userPath, userExists := s.Find(".publicItem_status a").First().Attr("href")
		userURL := t.appendBaseHostToPath(userPath, userExists)

		stockCount, err := strconv.Atoi(t.trim(s.Find(".publicItem_stockCount").Text()))
		if err != nil {
			stockCount = 0
		}

		a := Article{
			ArticleTitle: articleTitle,
			ArtiClePath: articlePath,
			ArticleURL:  articleURL,
			UserName:    userName,
			UserPath:    userPath,
			UserURL:     userURL,
			StockCount:  stockCount,
		}
		articles = append(articles, a)

	})

  return articles, nil
}

func (t *Explore) appendBaseHostToPath(address string, exists bool) *url.URL {
	if exists == false {
		return nil
	}

	u, err := url.Parse(baseHost)
	if err != nil {
		return nil
	}

	u.Path = address

	return u
}

func (t *Explore) trim(name string) string {
	trimmedNameParts := []string{}

	nameParts := strings.Split(name, "\n")
	for _, part := range nameParts {
		trimmedNameParts = append(trimmedNameParts, strings.TrimSpace(part))
	}

	return strings.Join(trimmedNameParts, "")
}

func (t *Explore) generateURL(tag string, category string, page int) (*url.URL, error) {
	parseURL := baseHost + basePath + tag + "/" + category + "?page=" + strconv.Itoa(page)

	u, err := url.Parse(parseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()

	u.RawQuery = q.Encode()

	return u, nil
}

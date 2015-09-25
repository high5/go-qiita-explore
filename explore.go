// Package explore provides access to qiita`s tags(items and stocks)
// The data will be collected from qiitas website at http://qiita.com/tags/{tag}/stocks and http://qiita.com/tags/{tag}/items
package explore

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
)

// NewExplore is the main entry point of the explore package.
// It provides access to the API of this package by returning a Article datastructure.
// Usage:
// explore := explore.NewExplore()
// stocks, err := explore.GetItems("swift", 1)
// items, err := explore.GetItems("Go", 1)
// Usage:
//   explore := explore.NewExplore()
//   stocks, err := explore.GetItems("swift", 1)
//   items, err := explore.GetItems("Go", 1)
func NewExplore() *Explore {
	t := Explore{}
	return &t
}

// GetStocks provides a slice of Article(recently stocked) filtered by the given tag and page.
// tag can be filtered by applying a tag by your choice. The input must be a known tag by Qiita.
// page can be filtered by applying by one page of pagers. If an empty int will be applied first page (1) will be the default.
func (t *Explore) GetStocks(tag string, page ...int) ([]Article, error) {
	var pageQuery int = 1
  if (len(page) > 0) {
		pageQuery = page[0]
	}
	return t.getArticles(tag, "stocks", pageQuery)
}

// GetItems provides a slice of Article(articles newly posted) filtered by the given tag and page.
// tag can be filtered by applying a tag by your choice. The input must be a known tag by Qiita.
// page can be filtered by applying by one page of pagers. If an empty int will be applied first page (1) will be the default.
func (t *Explore) GetItems(tag string, page ...int) ([]Article, error) {
	var pageQuery int = 1
  if (len(page) > 0) {
		pageQuery = page[0]
	}
	return t.getArticles(tag, "items", pageQuery)
}


func (t *Explore) getArticles(tag string, category string, page ...int) ([]Article, error) {
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
		title := t.trim(s.Find(".publicItem_body a").First().Text())
	  path, articleExists := s.Find(".publicItem_body a").First().Attr("href")
		URL := t.appendBaseHostToPath(path, articleExists)

		// Collect tag
		var tags []string
		s.Find(".tagList_item a").Each(func(i int, s *goquery.Selection) {
			tags = append(tags, s.Text())
		})

		userName := t.trim(s.Find(".publicItem_status a").Text())
		userPath, userExists := s.Find(".publicItem_status a").First().Attr("href")
		userURL := t.appendBaseHostToPath(userPath, userExists)

		stockCount, err := strconv.Atoi(t.trim(s.Find(".publicItem_stockCount").Text()))
		if err != nil {
			stockCount = 0
		}

		a := Article{
			Title: title,
			Path: path,
			URL:  URL,
			Tags: tags,
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

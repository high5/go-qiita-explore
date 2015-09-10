package explore

import (
	//"github.com/PuerkitoBio/goquery"
	//"net/url"
	//"regexp"
	//"strconv"
	//"strings"
)

// 先頭大文字 public method
// explore := explore.NewExplore()

func NewExplore() *Explore {
	t := Explore{}
	return &t
}

func (t *Explore) GetLastWeekTrendingArticles(tag string) ([]Article, error) {
  return articles, nil
}


func (t *Explore) Test(a int) (int) {
  return 2
}

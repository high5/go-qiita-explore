package explore

import (
	"github.com/PuerkitoBio/goquery"
	//"net/url"
)

type Explore struct {
	document *goquery.Document
}


type Article struct {
	Name           string
	Owner          string
	RepositoryName string
	Description    string
	Language       string
	Stars          int
	//URL            *url.URL
	//ContributerURL *url.URL
	//Contributer    []Developer
}

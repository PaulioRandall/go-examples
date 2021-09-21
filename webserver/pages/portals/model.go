package portals

import (
	"github.com/PaulioRandall/go-examples/webserver/pages"
)

type (
	dataModel struct {
		Title     string
		PageLinks []PageLink
	}

	PageLink struct {
		Name string
		Link string
	}
)

var pageLinks = []PageLink{
	makePageLink("Start", pages.LinkStartPage),
	makePageLink("1. First", pages.LinkFirstPage),
	makePageLink("2. Second", pages.LinkSecondPage),
	makePageLink("3. Third", pages.LinkThirdPage),
}

func makeDataModel() dataModel {
	return dataModel{
		Title:     "Pathfinders Den",
		PageLinks: pageLinks,
	}
}

func makePageLink(name, link string) PageLink {
	return PageLink{
		Name: name,
		Link: link,
	}
}

package server

import (
	"github.com/PaulioRandall/ranch-go/routing/moo"

	"github.com/PaulioRandall/go-examples/server/govuk"
	"github.com/PaulioRandall/go-examples/server/pages"

	"github.com/PaulioRandall/go-examples/server/pages/first"
	"github.com/PaulioRandall/go-examples/server/pages/portals"
	"github.com/PaulioRandall/go-examples/server/pages/second"
	"github.com/PaulioRandall/go-examples/server/pages/start"
	"github.com/PaulioRandall/go-examples/server/pages/third"
)

func CreateHandler() *moo.Handler {

	m := moo.Mux{}

	m.Before(logRequest)
	m.Wrap(timeRequest)

	// Start page
	startPage := moo.Mux{}
	startPage.Before(replacePathsWith(pages.LinkStartPage))
	startPage.GET("/", start.ServePage)
	startPage.GET("/index", start.ServePage)
	startPage.GET("/home", start.ServePage)
	startPage.GET(pages.LinkStartPage, start.ServePage)
	m.Routes(startPage)

	// Form pages
	m.GET("/pages", replacePathWith(pages.LinkPortalsPage, portals.ServePage))
	m.GET(pages.LinkPortalsPage, portals.ServePage)
	m.GET(pages.LinkFirstPage, first.ServePage)
	m.GET(pages.LinkSecondPage, second.ServePage)
	m.GET(pages.LinkThirdPage, third.ServePage)

	// Static content
	m.GET("/favicon.ico", replacePathWith("/govuk/assets/images/favicon.ico", govuk.ServeStatic))
	m.Begins("GET", "/govuk/", govuk.ServeStatic)

	// 404
	m.Route(moo.NotFound())

	return moo.NewHandler(m.Build(""))
}

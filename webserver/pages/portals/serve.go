package portals

import (
	"net/http"

	"github.com/PaulioRandall/go-examples/webserver/shared/components"
	"github.com/PaulioRandall/go-examples/webserver/shared/page"
	"github.com/PaulioRandall/go-examples/webserver/shared/render"
)

var renderer page.PageRenderer = render.NewPageRenderer(
	"webserver/pages/portals/page.html",
	"webserver/pages/portals/page-links.html",
	components.StdHeadHTML,
	components.PageTitleHTML,
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	model := makeDataModel()

	page.Serve(w, renderer, model)
}

package portals

import (
	"net/http"

	"github.com/PaulioRandall/go-examples/server/shared/components"
	"github.com/PaulioRandall/go-examples/server/shared/page"
	"github.com/PaulioRandall/go-examples/server/shared/render"
)

var renderer page.PageRenderer = render.NewPageRenderer(
	"server/pages/portals/page.html",
	"server/pages/portals/page-links.html",
	components.StdHeadHTML,
	components.PageTitleHTML,
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	model := makeDataModel()

	page.Serve(w, renderer, model)
}

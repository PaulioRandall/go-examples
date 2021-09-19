package start

import (
	"net/http"

	"github.com/PaulioRandall/go-examples/server/shared/components"
	"github.com/PaulioRandall/go-examples/server/shared/page"
	"github.com/PaulioRandall/go-examples/server/shared/render"
)

var renderer page.PageRenderer = render.NewPageRenderer(
	"server/pages/start/page.html",
	"server/pages/start/start-now-button.html",
	components.StdHeadHTML,
	components.PageTitleHTML,
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	j := makeJourney()
	model := makeDataModel(j)

	page.Serve(w, renderer, model)
}

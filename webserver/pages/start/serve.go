package start

import (
	"net/http"

	"github.com/PaulioRandall/go-examples/webserver/shared/components"
	"github.com/PaulioRandall/go-examples/webserver/shared/page"
	"github.com/PaulioRandall/go-examples/webserver/shared/render"
)

var renderer page.PageRenderer = render.NewPageRenderer(
	"webserver/pages/start/page.html",
	"webserver/pages/start/start-now-button.html",
	components.StdHeadHTML,
	components.PageTitleHTML,
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	j := makeJourney()
	model := makeDataModel(j)

	page.Serve(w, renderer, model)
}

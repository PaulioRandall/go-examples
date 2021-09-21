package third

import (
	"net/http"

	"github.com/PaulioRandall/go-examples/webserver/pages"
	"github.com/PaulioRandall/go-examples/webserver/shared/components"
	"github.com/PaulioRandall/go-examples/webserver/shared/page"
	"github.com/PaulioRandall/go-examples/webserver/shared/render"
)

var model = struct {
	Title    string
	BackLink string
}{
	Title:    "Third",
	BackLink: pages.LinkSecondPage,
}

var renderer page.PageRenderer = render.NewPageRenderer(
	"webserver/pages/third/page.html",
	components.StdHeadHTML,
	components.PageTitleHTML,
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	page.Serve(w, renderer, model)
}

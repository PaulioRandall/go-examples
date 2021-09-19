package first

import (
	"net/http"

	"github.com/PaulioRandall/go-examples/server/pages"
	"github.com/PaulioRandall/go-examples/server/shared/components"
	"github.com/PaulioRandall/go-examples/server/shared/page"
	"github.com/PaulioRandall/go-examples/server/shared/render"
)

var model = struct {
	Title    string
	BackLink string
	NextLink string
}{
	Title:    "First",
	BackLink: pages.LinkStartPage,
	NextLink: pages.LinkSecondPage,
}

var renderer page.PageRenderer = render.NewPageRenderer(
	"server/pages/first/page.html",
	components.StdHeadHTML,
	components.PageTitleHTML,
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	page.Serve(w, renderer, model)
}

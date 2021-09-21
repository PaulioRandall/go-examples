package start

import (
	"github.com/PaulioRandall/go-examples/webserver/pages"
)

type journey struct {
	nextLink string
}

func makeJourney() journey {
	return journey{
		nextLink: pages.LinkFirstPage,
	}
}

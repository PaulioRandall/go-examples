package webserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	//"github.com/PaulioRandall/go-examples/webserver/govuk"
	"github.com/PaulioRandall/go-examples/webserver/pages"

	"github.com/PaulioRandall/go-examples/webserver/shared/logfmt"

	"github.com/PaulioRandall/go-examples/webserver/pages/first"
	"github.com/PaulioRandall/go-examples/webserver/pages/portals"
	"github.com/PaulioRandall/go-examples/webserver/pages/second"
	"github.com/PaulioRandall/go-examples/webserver/pages/start"
	"github.com/PaulioRandall/go-examples/webserver/pages/third"
)

func ListenAndServe(port string) {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	// Start page
	startPage := logRequest(timeRequest(start.ServePage))
	http.HandleFunc("/", startPage)
	http.HandleFunc("/index", startPage)
	http.HandleFunc("/home", startPage)
	http.HandleFunc(pages.LinkStartPage, startPage)

	// Form pages
	http.HandleFunc("/pages", replacePath(pages.LinkPortalsPage, portals.ServePage))
	http.HandleFunc(pages.LinkPortalsPage, logAndTimeRequest(portals.ServePage))
	http.HandleFunc(pages.LinkFirstPage, logAndTimeRequest(first.ServePage))
	http.HandleFunc(pages.LinkSecondPage, logAndTimeRequest(second.ServePage))
	http.HandleFunc(pages.LinkThirdPage, logAndTimeRequest(third.ServePage))

	// Static content
	fileServer := http.FileServer(http.Dir("./webserver/govuk"))
	http.Handle("/govuk/", http.StripPrefix("/govuk", fileServer))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/govuk/assets/images/favicon.ico"
		fileServer.ServeHTTP(w, r)
	})

	// If we define our own handler we would replace the nil with it
	log.Fatal(http.ListenAndServe(port, nil))
}

func replacePath(newPath string, f http.HandlerFunc) http.HandlerFunc {
	f = replacePathWith(newPath, f)
	return logAndTimeRequest(f)
}

func logAndTimeRequest(f http.HandlerFunc) http.HandlerFunc {
	f = timeRequest(f)
	return logRequest(f)
}

func timeRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now().UnixNano()
		f(w, r)
		stop := time.Now().UnixNano()

		t := time.Duration(stop - start)
		logfmt.LogInfo("Finished in %v microseconds", t.Microseconds())
	}
}

func logRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logfmt.LogInfo(r.URL.String())
		f(w, r)
	}
}

func replacePathWith(newPath string, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = newPath
		f(w, r)
	}
}

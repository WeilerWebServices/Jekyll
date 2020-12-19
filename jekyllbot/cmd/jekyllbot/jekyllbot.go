// jekyllbot is the server which controls @jekyllbot on GitHub.
package main

import (
	_ "expvar"
	"flag"
	"log"
	"net/http"

	"github.com/jekyll/jekyllbot/ctx"
	"github.com/jekyll/jekyllbot/jekyll"
	"github.com/jekyll/jekyllbot/sentry"
)

var context *ctx.Context

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "The port to serve to")
	flag.Parse()
	context = ctx.NewDefaultContext()

	http.HandleFunc("/_ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("ok\n"))
	}))

	jekyllOrgHandler := jekyll.NewJekyllOrgHandler(context)
	http.Handle("/_github/jekyll", sentry.NewHTTPHandler(jekyllOrgHandler, map[string]string{
		"app": "jekyllbot",
	}))

	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

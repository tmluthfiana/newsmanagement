package main

import (
	"github.com/eaciit/knot/knot.v1"
	"net/http"
	_ "newsmanagement/webapp/webext"
)

func main() {
	app := knot.GetApp("newsmanagement")
	if app == nil {
		return
	}

	routes := make(map[string]knot.FnContent, 1)
	routes["/"] = func(r *knot.WebContext) interface{} {
		http.Redirect(r.Writer, r.Request, "/login/default", http.StatusTemporaryRedirect)
		return true
	}
	knot.StartAppWithFn(app, "localhost:8011", routes)
}

package apibgen

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UrlVarExtractor interface {
	Extract(req *http.Request) (vars map[string]string)
}

type GorillaMuxUrlVarExtractor struct {
	Router *mux.Router
}

var _ UrlVarExtractor = &GorillaMuxUrlVarExtractor{}

func (g *GorillaMuxUrlVarExtractor) Extract(req *http.Request) (vars map[string]string) {
	var match mux.RouteMatch
	if g.Router.Match(req, &match) {
		return match.Vars
	}
	return nil
}

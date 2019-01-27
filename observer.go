package apibgen

import (
	"log"
	"net/http"

	"github.com/gaw508/apibgen/apib"
)

type Observer struct {
	UrlVarExtractor UrlVarExtractor
	resources map[string]*apib.Resource
}

func (a *Observer) Observe() func(res *http.Response, req *http.Request) {
	return func(res *http.Response, req *http.Request) {
		// copy request body into Request object
		docReq, err := apib.NewRequest(req)
		if err != nil {
			log.Println("Error:", err.Error())
			return
		}

		// setup resource
		vars := a.UrlVarExtractor.Extract(req)
		u := apib.NewURL(req, vars)
		path := u.ParameterizedPath

		if a.resources[path] == nil {
			a.resources[path] = apib.NewResource(u)
		}

		// store response body in Response object
		docResp := apib.NewResponse(res)

		// find action
		action := a.resources[path].FindAction(req.Method)
		if action == nil {
			// make new action
			action, err = apib.NewAction(req.Method, "TODO: Handler name") // TODO: Handler name
			if err != nil {
				log.Println("Error:", err.Error())
				return
			}

			// add Action to Resource's list of Actions
			a.resources[path].AddAction(action)
		}

		// add request, response to action
		action.AddRequest(docReq, docResp)
	}
}

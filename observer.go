package apibgen

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"

	"github.com/gaw508/apibgen/apib"
	"github.com/steinfletcher/apitest"
)

type Observer struct {
	UrlVarExtractor UrlVarExtractor
	Writer          io.Writer
	resources       map[string]*apib.Resource
	doc             *apib.Doc
}

func NewObserver(extractor UrlVarExtractor, writer io.Writer, resourceGroupName string) *Observer {
	d, err := apib.NewDoc(resourceGroupName, fmt.Sprintf("%s.apib", resourceGroupName))
	if err != nil {
		panic(err)
	}
	return &Observer{
		UrlVarExtractor: extractor,
		Writer:          writer,
		doc:             d,
		resources:       make(map[string]*apib.Resource),
	}
}

func (o *Observer) Observe() apitest.Observe {
	return func(res *http.Response, req *http.Request, _ *apitest.APITest) {
		// copy request body into Request object
		docReq, err := apib.NewRequest(req)
		if err != nil {
			log.Println("Error:", err.Error())
			return
		}

		// setup resource
		vars := o.UrlVarExtractor.Extract(req)
		u := apib.NewURL(req, vars)
		path := u.ParameterizedPath

		if o.resources[path] == nil {
			o.resources[path] = apib.NewResource(u)
		}

		// store response body in Response object
		docResp := apib.NewResponse(res)

		// find action
		action := o.resources[path].FindAction(req.Method)
		if action == nil {
			// make new action
			action, err = apib.NewAction(req.Method, path)
			if err != nil {
				log.Println("Error:", err.Error())
				return
			}

			// add Action to Resource's list of Actions
			o.resources[path].AddAction(action)
		}

		// add request, response to action
		action.AddRequest(docReq, docResp)
	}
}

func (o *Observer) Write() {
	// sort resources by path
	var uris []string
	for k := range o.resources {
		uris = append(uris, k)
	}
	sort.Strings(uris)
	for _, uri := range uris {
		o.doc.AddResource(o.resources[uri])
	}

	err := o.doc.Write(o.Writer)
	if err != nil {
		panic(err)
	}
}

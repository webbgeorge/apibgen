package apibgen

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode"
)

func TestObserver(t *testing.T) {
	testCases := map[string]struct {
		requests []struct {
			req *http.Request
			res *http.Response
		}
		expectedDocPath string
	}{
		// TODO: More test cases
		"simple": {
			requests: []struct {
				req *http.Request
				res *http.Response
			}{
				{
					req: httptest.NewRequest(
						"PUT",
						"/item/123",
						bytes.NewBuffer([]byte(`{"paramOne": "valueOne", "paramTwo": false}`)),
					),
					res: &http.Response{
						StatusCode: http.StatusOK,
						Header: map[string][]string{"Content-Type": {"application/json"}},
						Body: ioutil.NopCloser(bytes.NewBuffer([]byte(`{"paramThree": "valueThree", "paramFour": 10}`))),
					},
				},
			},
			expectedDocPath: "testdata/simple.apib",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			extractor := &GorillaMuxUrlVarExtractor{getTestRouter()}
			var buf bytes.Buffer

			o := NewObserver(extractor, &buf, "Test Resource Group")
			oFn := o.Observe()
			for _, request := range tc.requests {
				oFn(request.res, request.req)
			}
			o.Write()

			expected := trimApiB(string(mustReadFile(tc.expectedDocPath)))
			actual := trimApiB(string(buf.Bytes()))
			assert.Equal(t, expected, actual)
		})
	}
}

func getTestRouter() *mux.Router {
	r := mux.NewRouter()
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	r.Handle("/item", handlerFunc).Methods("GET")
	r.Handle("/item", handlerFunc).Methods("POST")
	r.Handle("/item/{itemId}", handlerFunc).Methods("GET")
	r.Handle("/item/{itemId}", handlerFunc).Methods("PUT")
	r.Handle("/item/{itemId}", handlerFunc).Methods("DELETE")
	return r
}

func mustReadFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func trimApiB(s string) string {
	// TODO: Fix templates to remove trailing whitespace
	lines := strings.Split(strings.TrimSpace(s), "\n")
	newLines := make([]string, 0)
	for _, line := range lines {
		newLines = append(newLines, strings.TrimRightFunc(line, unicode.IsSpace))
	}
	return strings.Join(newLines, "\n")
}

# apibgen

A tool for generating API Blueprint documentation from
[https://github.com/steinfletcher/apitest](https://github.com/steinfletcher/apitest) tests

**NOTE: This library is a WIP, use with caution**

## Example usage

```golang
// This variable holds our observer state between tests
var apibObserver *apibgen.Observer

// The TestMain is used to set up our observer
func TestMain(m *testing.M) {
	// Create an io.Writer, to which the apib doc will be written
	// This example uses an os.File
	f, err := os.Create("docs.apib")
	if err != nil {
		panic(err)
	}

	// In order to parse url params, a UrlVarExtractor must be provided
	// This example uses gorilla mux's
	extractor := &apibgen.GorillaMuxUrlVarExtractor{Router: myRouter}

	// The observer is created with the above dependencies
	apibObserver = apibgen.NewObserver(extractor, f, "Config Service")

	// Execute the tests
	exitCode := m.Run()

	// Write the API Blueprint doc
	apibObserver.Write()

	// Return the exit code from the tests
	os.Exit(exitCode)
}

func TestMyHandler(t *testing.T) {
	apitest.New().
		Observe(apibObserver.Observe()). // This sets the apibgen observer
		Handler(myRouter).
		Get("/items/123").
		Expect(t).
		Status(http.StatusOK).
		Headers(map[string]string{
			"Content-Type": "application/json",
		}).
		Body(string(mustReadFile("testdata/expected-response.json"))).
		End()
}
```

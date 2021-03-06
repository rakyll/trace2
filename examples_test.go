package trace_test

import (
	"context"
	"log"
	"net/http"

	"github.com/rakyll/trace2"
	"github.com/rakyll/trace2/gcp"
)

func ExampleNewTrace() {
	c, err := gcp.NewClient(context.Background(), "project-id")
	if err != nil {
		log.Fatal(err)
	}
	trace.Configure(c)

	span, finish := trace.NewSpan("/foo")
	defer finish()

	span.Annotate("hello", []byte("error happened"))

	req, _ := http.NewRequest("GET", "http://google.com/", nil)
	req, err = span.ToHTTPReq(req)
	if err != nil {
		log.Fatal(err)
	}

	// do the request
}

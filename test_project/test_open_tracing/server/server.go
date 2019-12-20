package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {
	tracer := opentracing.GlobalTracer()

	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		// Extract the context from the headers
		fmt.Println("receive new request")
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		serverSpan := tracer.StartSpan("server", ext.RPCServerOption(spanCtx))
		defer serverSpan.Finish()
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}

package main

import (
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {
	fmt.Println("start")
	tracer := opentracing.GlobalTracer()

	clientSpan := tracer.StartSpan("client")
	defer clientSpan.Finish()

	url := "https://www.baidu.com"
	req, _ := http.NewRequest("GET", url, nil)

	// Set some tags on the clientSpan to annotate that it's the client span. The additional HTTP tags are useful for debugging purposes.
	ext.SpanKindRPCClient.Set(clientSpan)
	ext.HTTPUrl.Set(clientSpan, url)
	ext.HTTPMethod.Set(clientSpan, "GET")

	// Inject the client span context into the headers
	tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	resp, _ := http.DefaultClient.Do(req)
	fmt.Println(resp.Status)
}

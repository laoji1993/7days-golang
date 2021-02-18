package main

import (
	"7days-golang/gee-web/day1-http-base/base3/gee"
	"fmt"
	"net/http"
)

func main() {
	g := gee.New()

	g.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
	})

	g.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "header[%q]=%q\n", k, v)
		}
	})

	g.Run(":1234")
}

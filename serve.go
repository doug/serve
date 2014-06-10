package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.String("port", "8080", "Define what TCP port to bind to")
var root = flag.String("root", ".", "Define the root filesystem path")

func main() {
	flag.Parse()
	addr := ":" + *port
	fmt.Printf("Listening on %s...", addr)
	panic(http.ListenAndServe(addr, http.FileServer(http.Dir(*root))))
}

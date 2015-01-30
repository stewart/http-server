package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	port = flag.Int("port", 3000, "port to serve on")
	host = flag.String("host", "0.0.0.0", "host to serve on")
	path = flag.String("path", "./", "path to serve from")
)

func main() {
	flag.Parse()

	root, err := filepath.Abs(*path)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Static file server running at %s:%d.\n", *host, *port)
	fmt.Println("Ctrl-C to exit.")

	err = http.ListenAndServe(
		*host+":"+strconv.Itoa(*port),
		http.FileServer(http.Dir(root)),
	)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
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
		os.Exit(1)
	}

	fmt.Printf("Static file server running at %s:%d, serving %s\n", *host, *port, *path)
	fmt.Println("Ctrl-C to exit.")

	fileServer := http.FileServer(http.Dir(root))

	err = http.ListenAndServe(*host+":"+strconv.Itoa(*port), fileServer)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

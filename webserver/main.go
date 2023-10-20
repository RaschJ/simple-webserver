package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	port := flag.String("p", "8080", "port on which to serve web interfaces")
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		log.Fatalf("Directory to serve relative to the snap root directory is the only required argument")
	}

	snapdir := os.Getenv("SNAP")
	www := path.Join(snapdir, "index.html")
	
	//panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(www))))
	
 	panic(http.ListenAndServe(":"+*port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello World! ._.</h1>")
	})))
	
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s: [-p port] dir_to_serve\n", os.Args[0])
	flag.PrintDefaults()
}

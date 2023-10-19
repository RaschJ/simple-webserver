package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	port := flag.String("p", "8080", "port on which to serve web interfaces")
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		log.Fatalf("Directory to serve relative to the snap root directory is the only required argument")
	}

	snapdir := os.Getenv("SNAP")
	www := path.Join(snapdir, flag.Arg(0))

	http.HandleFunc("/view/", viewHandler)
	
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(www))))
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s: [-p port] dir_to_serve\n", os.Args[0])
	flag.PrintDefaults()
}

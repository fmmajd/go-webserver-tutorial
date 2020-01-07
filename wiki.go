package main

import (
	"./internal/page"
	"fmt"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := page.LoadPage(title)
	if err != nil {
		fmt.Fprintf(w, "<h1>Not Found</h1><div>%Page Not Found</div>")
	} else {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	}
}

func main () {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}










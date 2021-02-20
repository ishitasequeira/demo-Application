package main

import (
	"html/template"
	"log"
	"net/http"
)

type greet struct {
	msg string
}

func main() {
	http.HandleFunc("/", greetHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	g := greet{msg: r.FormValue("name")}
	tmpl.Execute(w, struct{ Name string }{g.msg})
}

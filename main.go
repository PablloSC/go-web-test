package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r* http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{ true })

		value := r.FormValue("md5") // pega o campo do formulario
		value  = fmt.Sprintf("%x", md5.Sum([]byte(value))) // converte para md5
		fmt.Fprintf(w, "%v", value)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

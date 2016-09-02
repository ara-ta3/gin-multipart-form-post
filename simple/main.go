package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	absTemplate := filepath.Dir(os.Args[0])
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tmpl, e := template.ParseFiles(absTemplate + "/../templates/index.html")
			if e != nil {
				panic(e)
			}
			e = tmpl.Execute(w, nil)
			if e != nil {
				panic(e)
			}
		} else if r.Method == "POST" {
			r.ParseMultipartForm(32 << 20)
			m := map[string]interface{}{
				"message": "ok",
				"forms":   r.MultipartForm.Value,
			}
			j, e := json.Marshal(m)
			if e != nil {
				panic(e)
			}
			fmt.Fprintf(w, string(j))
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

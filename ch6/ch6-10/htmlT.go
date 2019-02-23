package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var DATA []Entry
var tFile string

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tFile))
	myT.ExecuteTemplate(w, tFile, DATA)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Need Database File + Template File!")
		return
	}

	database := args[1]
	tFile := args[2]

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(nil)
	}
}

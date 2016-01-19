package main

import (
	"database/sql"
	"fmt"
	"github.com/smehraein/pgdbGO"
	"github.com/smehraein/sanitizeGO"
	"html/template"
	"io"
	"log"
	"net/http"
)

var db *sql.DB

func hello(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()
	if err != nil {
		io.WriteString(w, "Awwww.")
	} else {
		io.WriteString(w, "GOT DB")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// Display values
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("validEmail:", sanitizeGO.Email(r.Form["email"][0]))
		io.WriteString(w, "Logged in!")
	}
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}
	DBINFO := pgdbGO.PGConnection{"postgres", "", "postgres", "localhost"}
	var connErr error
	db, connErr = pgdbGO.Connect(DBINFO)
	if connErr != nil {
		log.Fatal(connErr)
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/login"] = login
	mux["/"] = hello
	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}

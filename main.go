package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
        "flag"
	"github.com/gorilla/mux"
	expand "github.com/openvenues/gopostal/expand"
	parser "github.com/openvenues/gopostal/parser"
)

type Request struct {
	Query string `json:"query"`
}

func main() {
	
	var port string
    	flag.StringVar(&port, "port", "5000", "default port 5000")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/expand", ExpandHandler).Methods("POST")
	r.HandleFunc("/parser", ParserHandler).Methods("POST")
	fmt.Println("listening on port: " + port)
	http.ListenAndServe(":" + port  , r)
}

func ExpandHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request

	q, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(q, &req)

	expansions := expand.ExpandAddress(req.Query)

	expansionThing, _ := json.Marshal(expansions)
	w.Write(expansionThing)
}

func ParserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request

	q, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(q, &req)

	parsed := parser.ParseAddress(req.Query)
	parseThing, _ := json.Marshal(parsed)
	w.Write(parseThing)
}

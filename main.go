package main

import (
	"encoding/json"
	"net/http"
)

type Quote struct {
	Author   string
	Year     int
	Sentence string
}

func main() {
	http.HandleFunc("/quotes", ListQuotes)
	http.HandleFunc("/quote", GetQuote)
	http.ListenAndServe(":8080", nil)
}

func GetQuote(w http.ResponseWriter, r *http.Request) {
	q := Quote{
		Author:   "Ken Olson",
		Year:     1977,
		Sentence: "I think we dont need five computes in our house",
	}

	jsonStr, _ := json.Marshal(q)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

func ListQuotes(w http.ResponseWriter, r *http.Request) {
	quotes := []Quote{}
	quotes = append(quotes, Quote{
		Author:   "Thomas Watson",
		Year:     1943,
		Sentence: "We need 5 computer",
	})
	quotes = append(quotes, Quote{
		Author:   "Ken Olson",
		Year:     1977,
		Sentence: "I think we dont need five computes in our house",
	})
	jsonStr, _ := json.Marshal(quotes)

	w.Header().Set("content-type", "application/json")
	w.Write(jsonStr)
}

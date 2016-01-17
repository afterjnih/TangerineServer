package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type BookMark struct {
	UserName string `json:"userName"`
	BookName string `json:"bookName"`
	Page     int	`json:"page"`
	UpdateAt time.Time `json:"updated_at"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	bookMark := &BookMark{
		UserName: "koji",
		BookName: "everydayrailsrspec-jp.pdf",
		Page: 75,
		UpdateAt: time.Now(),
	}
	bm, err := json.Marshal(bookMark)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(bm))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}

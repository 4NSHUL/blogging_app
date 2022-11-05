package main

import (
	"encoding/json"
	"net/http"
)

func GetBlogs(w http.ResponseWriter, r *http.Request) {

	output, err := FetchBlogsData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendJsonresponse(w, output)

}

func SendJsonresponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

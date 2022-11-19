package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func SendJsonresponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getBlogs(w http.ResponseWriter, r *http.Request) {

	output, err := FetchBlogsData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendJsonresponse(w, output)

}

func getBlog(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	output, err := FetchBlogData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendJsonresponse(w, output)

}

func createBlog(w http.ResponseWriter, r *http.Request) {
	var blogReqBody Blog
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&blogReqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = CreateBlogData(blogReqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendJsonresponse(w, http.StatusCreated)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := DeleteBlogData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	SendJsonresponse(w, http.StatusOK)
}

func updateBlog(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var blogUpdateReqBody BlogUpdate
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&blogUpdateReqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = UpdateBlogData(id, blogUpdateReqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	SendJsonresponse(w, http.StatusOK)
}

// func createBlog(w http.ResponseWriter, r *http.Request) {

// 	output, err := CreateBlog()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	SendJsonresponse(w, output)

// }

// func SignUp(w http.ResponseWriter, r *http.Request) {
// 	connection := GetDatabase()
// 	defer Closedatabase(connection)

// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		var err Error
// 		err = SetError(err, "Error in reading body")
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	var dbuser User
// 	connection.Where("email = ?", user.Email).First(&dbuser)

// 	//checks if email is already register or not
// 	if dbuser.Email != "" {
// 		var err Error
// 		err = SetError(err, "Email already in use")
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}

// 	user.Password, err = GeneratehashPassword(user.Password)
// 	if err != nil {
// 		log.Fatalln("error in password hash")
// 	}

//insert user details in database
// 	connection.Create(&user)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// }

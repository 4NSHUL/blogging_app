package auth

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/golang-jwt/jwt"
// )

// func SignIn(w http.ResponseWriter, r *http.Request) {
// 	connection := GetDatabase()
// 	defer Closedatabase(connection)

// 	var authdetails Authentication
// 	err := json.NewDecoder(r.Body).Decode(&authdetails)
// 	if err != nil {
// 		var err Error
// 		err = SetError(err, "Error in reading body")
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}

// 	var authuser User
// 	connection.Where("email = ?", authdetails.Email).First(&authuser)
// 	if authuser.Email == "" {
// 		var err Error
// 		err = SetError(err, "Username or Password is incorrect")
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}

// 	check := CheckPasswordHash(authdetails.Password, authuser.Password)

// 	if !check {
// 		var err Error
// 		err = SetError(err, "Username or Password is incorrect")
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}

// 	validToken, err := GenerateJWT(authuser.Email, authuser.Role)
// 	if err != nil {
// 		var err Error
// 		err = SetError(err, "Failed to generate token")
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}

// 	var token Token
// 	token.Email = authuser.Email
// 	token.Role = authuser.Role
// 	token.TokenString = validToken
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(token)
// }

// func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] == nil {
// 			var err Error
// 			err = SetError(err, "No Token Found")
// 			json.NewEncoder(w).Encode(err)
// 			return
// 		}

// 		var mySigningKey = []byte(secretkey)

// 		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("There was an error in parsing")
// 			}
// 			return mySigningKey, nil
// 		})

// 		if err != nil {
// 			var err Error
// 			err = SetError(err, "Your Token has been expired")
// 			json.NewEncoder(w).Encode(err)
// 			return
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			if claims["role"] == "admin" {

// 				r.Header.Set("Role", "admin")
// 				handler.ServeHTTP(w, r)
// 				return

// 			} else if claims["role"] == "user" {

// 				r.Header.Set("Role", "user")
// 				handler.ServeHTTP(w, r)
// 				return
// 			}
// 		}
// 		var reserr Error
// 		reserr = SetError(reserr, "Not Authorized")
// 		json.NewEncoder(w).Encode(err)
// 	}
// }

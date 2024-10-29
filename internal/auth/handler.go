package auth

import "net/http"

type UserHandler struct{
	UserRepo UserRepository
}

func LoginUser(w http.ResponseWriter, r *http.Request){

}

func RegisterUser (w http.ResponseWriter, r *http.Request){

}
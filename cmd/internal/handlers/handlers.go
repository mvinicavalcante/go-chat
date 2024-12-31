package handlers

import "net/http"

//! Lidam com requisições HTTP
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá! Bem-vindo ao nosso site!"))
}
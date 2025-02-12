package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário!"))
}
func SearchAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Todos Usuários!"))
}
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário!"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}

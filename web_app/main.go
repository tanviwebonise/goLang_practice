package main

import "net/http"
import "github.com/controller"
import "github.com/dao"
import "io"

func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong!")
}

func main() {
    dao.InitDB("root:root@tcp(localhost:3306)/test")
	mux:= http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/home", controller.ServeHomePage)
    mux.HandleFunc("/login", controller.Login)
    mux.HandleFunc("/signup", controller.SingupPage)
	http.ListenAndServe(":8080", mux)
}

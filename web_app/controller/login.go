package controller

import "net/http"
import "golang.org/x/crypto/bcrypt"
import "github.com/dao"

var err error
func Login(res http.ResponseWriter, req *http.Request) {
    
    if req.Method != "POST" {
        http.ServeFile(res, req, "template/login.html")
        return
    }   

	username := req.FormValue("username")
    password := req.FormValue("password")

    databaseUsername, databasePassword, err := dao.Login(username)

    if err != nil {
        http.Error(res, "Failed to login", 404)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
    if err != nil {
        http.Error(res, "Password did not match", 401)
        return
    }
    res.Write([]byte("Hello " + databaseUsername))
}
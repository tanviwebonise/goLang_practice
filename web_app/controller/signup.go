package controller

import "net/http"
import "github.com/dao"
import "golang.org/x/crypto/bcrypt"

func SingupPage(res http.ResponseWriter, req *http.Request) {
    
    if req.Method != "POST" {
        http.ServeFile(res, req, "template/signup.html")
        return
    } 

    username := req.FormValue("username")
    address := req.FormValue("address")
    password := req.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(res, "Invalid password", 422)    
    } 
    _, err = dao.Signup(username, address, hashedPassword)
    if err != nil {
        http.Error(res, "Failed to signup user as username already exists", 409)  
    } 
    res.Write([]byte("User created successfully!"))   
}

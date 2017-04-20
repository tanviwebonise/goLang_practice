package controller 

import "net/http"

func ServeHomePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "/template/index.html")	
}	


package dao

type user interface {
    Login(username string, password string) string
    Singup(username string, address string, password string) bool
}


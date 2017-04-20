package dao

var err error
func Login(username string) (string, string, error) {
	var databaseUsername  string
    var databasePassword  string 
    err = db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

    if err != nil {
        return "", "", err
    }
    return databaseUsername, databasePassword, nil
}

func Signup(username string, address string, password []byte) (bool, error) {
    _, err = db.Exec("INSERT INTO users(username, address, password) VALUES(?, ?, ?)", username, address, password)

    if err != nil {
        return false, err
        
    }
    return true, nil
}
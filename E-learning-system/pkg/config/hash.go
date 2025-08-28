package utils

import "golang.org/x/crypto/bcrypt"


func HashePassword(password string) (string, error) {
 HashePassword, err := bcrypt.GenerateFromPassword([]byte (password) , bcrypt.DefaultCost)
 if err != nil {
	return "", err
 }
 return string(HashePassword), nil 

}

// ComparePassword compares the hashed password with the plain text password
func CheckPasswordHash( password, hash string) bool {
 err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
 return  err  == nil
}
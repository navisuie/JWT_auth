# JWT_auth
Created a JWT authentication system using GO, while using a postgres database and Postman for testing

Allows the user to signup with an email and password and then encrypts the password, allows the user to login in while checking the encrypted password to the entered password, and also sends a JWT token as a cookie to the page

Packages:
Gorm (https://github.com/go-gorm/gorm)
Gin (https://github.com/gin-gonic/gin)
Bcrypt (https://github.com/golang/crypto/blob/master/bcrypt/bcrypt.go)
jwt-go (https://github.com/golang-jwt/jwt)

package controllers

import (
	"net/http"
	"os"
	"time"

	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to get body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to create user"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to get body"})
		return
	}
	var user models.User
	initializers.DB.Find(&user, "email = ?", body.Email)

	if user.ID == 0 {

		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return

	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Token is Bad"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AUTH", tokenString, 3600*24*30, "", "", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "logged in"})
}

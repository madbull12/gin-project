package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/madbull12/gin_postgres/initializers"
	"github.com/madbull12/gin_postgres/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	v := validator.New()

	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(v.Struct(&body)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body ",
		})
		return

	}

	//email validation

	if err := v.Var(body.Email, "email"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email address",
		})
		return
	}

	//hash password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	//create the user
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})

}

func Signin(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body ",
		})
		return

	}

	var user models.User
	// initializers.DB.First(&user, "email = ?", body.Email)
	initializers.DB.Where("email = ?", body.Email).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password 3",
		})
		return
	}

	//compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password 2",
		})
		return

	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

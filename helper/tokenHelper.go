package tokenhelper

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type UserDetails struct {
	Email    string
	UserName string
	City     string
	jwt.StandardClaims
}

func ValidateToken(token string) string {
	loaderr := godotenv.Load(".env")

	if loaderr != nil {
		fmt.Println("Unable to load env file")
		log.Fatal(loaderr)
	}

	secretKey := os.Getenv("JWT_SECRET")

	parsedToken, err := jwt.ParseWithClaims(token, &UserDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("Error: Unable to parse token while validating")
		return "Invalid Token"
	}

	claim, valid := parsedToken.Claims.(*UserDetails)

	if !valid {
		fmt.Println("Invalid Token")
		return "Invalid Token"
	}

	if claim.ExpiresAt < time.Now().Local().Unix() {
		fmt.Println("Invalid Token")
		return "Token Expired"
	}

	return "valid"

}
func GenerateToken(email string, username string, city string) string {
	loaderr := godotenv.Load(".env")

	if loaderr != nil {
		fmt.Println("Unable to load env file")
		log.Fatal(loaderr)
	}

	secret := os.Getenv("JWT_SECRET")

	claim := &UserDetails{
		Email:    email,
		UserName: username,
		City:     city,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(120)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))

	if err != nil {
		fmt.Println("ERROR generating token")
		log.Fatal(err)
	}

	return token
}

func ConvertPasswordToHash(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		fmt.Println("Error: Unable to Generate password Hash")
		log.Fatal(err)
	}

	return string(b)
}

func VerifyUserPassword(dbPassword string, userPassword string) bool {
	correct := true

	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(userPassword))

	if err != nil {
		fmt.Println("Password Does'nt match")
		correct = false
	}

	return correct
}

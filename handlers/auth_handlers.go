package handlers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// pull it from and env
var jwtKey = []byte("this is supposed to be super secret")

// Extend this to also generate refresh tokens
// https://dev.to/mdfaizan7/creating-jwt-s-and-signup-route-part-2-3-of-go-authentication-series-5339
type JWTClaims struct {
	Email string
	jwt.StandardClaims
}

func GetJWTKey(user User) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	jwtClaim := &JWTClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)

	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValiadateTokenString(tokenString string) (claim *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*JWTClaims)

	if !ok {
		return nil, errors.New("Could not parse claims")
	}

	if claim.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("The token Expired")
	}

	return claim, nil
}

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email" `
	Password string `json:"password"`
}

var authed_user = []User{
	{
		Id:       "0202",
		Email:    "demo@demo.com",
		Password: "password",
	},
}

func HanldeGetAuthUser(c *fiber.Ctx) error {
	tokenString := string(c.Request().Header.Peek("Authorization"))

	// header format Bearer thisistheactualtokenhere

	tokenArray := strings.Split(tokenString, "Bearer ")

	if len(tokenArray) == 2 {
		tokenString = tokenArray[1]
	}

	claim, err := ValiadateTokenString(tokenString)

	if err != nil {

		return c.JSON(fiber.Map{
			"user":  nil,
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"user":  *claim,
		"error": nil,
	})
}

// creates a jwt token
func HanldeCreateAuthUser(c *fiber.Ctx) error {
	user := User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if user.Email == authed_user[0].Email && user.Password == authed_user[0].Password {
		// create a jwt
		fmt.Println(user)
		tokenString, err := GetJWTKey(user)

		fmt.Println(tokenString, err)

		if err != nil {

			return c.JSON(fiber.Map{
				"access_token": nil,
				"error":        err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"access_token": tokenString,
			"error":        nil,
		})
	}

	return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
		"access_token": nil, "error": "OOPS something wrong with that data",
	})
}

func HandleDeleteAuthUser(c *fiber.Ctx) error {
	return c.JSON("The authenticated user has been deleted.")
}

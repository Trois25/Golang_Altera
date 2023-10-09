package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	JWT_SECRET_KEY string
}

func loadEnv() *AppConfig {
	var res = new(AppConfig)
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : Cannot load config file,", err.Error())
		return nil
	}

	if secretKey, found := os.LookupEnv("JWT_SECRET_KEY"); found {
		res.JWT_SECRET_KEY = secretKey
	}
	return res
}

func JWTMiddleware() echo.MiddlewareFunc {
	key := loadEnv()
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(key.JWT_SECRET_KEY),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId int, name string) (string, error) {
	key := loadEnv()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key.JWT_SECRET_KEY))

}

func ExtractTokenUserId(e echo.Context) (int, string) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		name := claims["name"].(string)
		return int(userId), name
	}
	return 0, ""
}

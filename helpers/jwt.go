package helpers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "rahasia"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	fmt.Println(claims)

	perseToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	fmt.Println(perseToken)
	signedToken, _ := perseToken.SignedString([]byte(secretKey))
	fmt.Println(signedToken)
	return signedToken
}

// func VerifyToken(c *gin.Context) (interface{}, error) {
// 	errReponse := errors.New("sign in to proceed")
// 	headerToken := c.Request.Header.Get("Authorization")
// 	bearer := strings.HasPrefix(headerToken, "Bearer")

// 	if !bearer {
// 		return nil, errReponse
// 	}

// 	stringToken := strings.Split(headerToken, " ")[1]
// 	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errReponse
// 		}
// 		return []byte(secretKey), nil
// 	})

// 	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
// 		return nil, errReponse
// 	}
// 	return token.Claims.(jwt.MapClaims), nil
// }

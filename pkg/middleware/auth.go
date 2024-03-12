package middleware

import (
	"bike-rent-express/model"
	"bike-rent-express/model/dto/json"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	applicationNone  = "incubation-golang"
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignaturKey   = []byte("incubation-golang") // TODO: replace with secure key
)

func GenerateTokenJwt(username string, roles string) (string, error) {
	loginExpDuration := time.Now().Add(24 * time.Hour).Unix()
	claims := model.JWTClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    applicationNone,
			ExpiresAt: loginExpDuration,
		},
		Username: username,
		Roles:    roles,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(jwtSignaturKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func JWTAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			json.NewResponseUnauthorized(c, "Invalid token", "01", "01")
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &model.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtSignaturKey, nil
		})

		if err != nil {
			json.NewResponseError(c, "Invalid token", "01", "01")
			c.Abort()
			return
		}

		if !token.Valid {
			json.NewResponseForbidden(c, "Forbidden", "03", "03")
			c.Abort()
			return
		}

		// validation roles
		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims.Roles {
					validRole = true
					break
				}
			}
		}

		if !validRole {
			json.NewResponseForbidden(c, "Forbidden", "03", "03")
			c.Abort()
			return
		}

		c.Next()
	}
}

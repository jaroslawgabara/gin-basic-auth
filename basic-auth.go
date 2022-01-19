package gin_basic_auth

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strings"
)

func BasicAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(401, "Unauthorized", c)
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		credentials := &Credentials{pair[0], pair[1]}

		c.Set("credentials", credentials)
		c.Next()
	}
}

type Credentials struct {
	Username string
	Password string
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}

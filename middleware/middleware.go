package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/haviz000/API-multi-level-user-golang/helpers"
	"github.com/haviz000/API-multi-level-user-golang/model"
)

func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if bearerIsExist := strings.Contains(auth, "Bearer"); !bearerIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorization",
		})
		return
	}

	token := strings.Split(auth, " ")
	if len(token) < 2 {
		err := errors.New("Must provide Authorization header with format `Bearer {token}`")

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	claims, err := helpers.VerifyAccessToken(token[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	c.Set("user_id", claims.UserID)
	c.Set("role", claims.Role)

	c.Next()
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haviz000/API-multi-level-user-golang/model"
)

func AdminMiddleware(c *gin.Context) {
	role, roleIsExist := c.Get("role")
	if !roleIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorized",
		})
		return
	}

	if role.(bool) == false {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorized",
		})
		return
	}

	c.Next()
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"net/http"
)

const (
	accessTokenHeaderName  = "X-Access-Token"
	refreshTokenHeaderName = "X-Refresh-Token"
	userCtx                = "authorizedUser"
)

func (h *Handler) authorizedUser(c *gin.Context) {
	accessToken := c.GetHeader(accessTokenHeaderName)
	if accessToken == "" {
		newErrorResponse(c, http.StatusUnauthorized, errNotAuthorized)
		return
	}

	u, err := models.ParseAccessToken(accessToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err)
	}

	c.Set(userCtx, u)
}

func (h *Handler) hasRefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader(refreshTokenHeaderName)
	if refreshToken == "" {
		newErrorResponse(c, http.StatusUnauthorized, errNotAuthorized)
		return
	}

	u, err := models.ParseAccessToken(refreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err)
	}

	c.Set(userCtx, u)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

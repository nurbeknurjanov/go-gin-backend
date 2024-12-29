package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/manuals"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"net/http"
)

// @Summary Login
// @Tags auth
// @Description Authorize
// @ID login
// @Accept json
// @Produce json
// @Param inputBody body models.LoginRequest true "login params"
// @Success 200 {object} models.Tokens
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	input := models.LoginRequest{}

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	tokens, err := h.services.Auth.Login(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// @Summary GetAccessToken
// @Tags auth
// @Description Get access token
// @ID getAccessToken
// @Accept json
// @Produce json
// @Param inputBody body models.LoginRequest true "login params"
// @Success 200 {string} string
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/get-access-token [get]
func (h *Handler) getAccessToken(c *gin.Context) {
	uData, _ := c.Get(userCtx)
	token, err := h.services.Auth.GetAccessToken(uData.(*models.User))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, token)
}

func (h *Handler) test(c *gin.Context) {
	var input *manuals.Test

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, input)
}

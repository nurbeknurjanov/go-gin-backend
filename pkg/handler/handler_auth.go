package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"net/http"
)

func (h *Handler) login(c *gin.Context) {
	/*var input struct {
		Email    string
		Password string
	}*/
	input := struct {
		Email    string
		Password string
	}{}

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	tokens, err := h.services.IAuthService.Login(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) getAccessToken(c *gin.Context) {
	uData, _ := c.Get(userCtx)
	token, err := h.services.IAuthService.GetAccessToken(uData.(*models.User))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, token)
}

func (h *Handler) test(c *gin.Context) {
	var input *helpers.Test

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, input)
}

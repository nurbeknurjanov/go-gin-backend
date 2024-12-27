package handlers

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"net/http"
)

func (h *Handler) profile(c *gin.Context) {
	/*h.services.IUsersService.(*service.UsersService).UsersRepo.(*repository.Repositories).IUsersRepository.FindByEmail("sdf")
	h.services.IUsersService.(*service.UsersService).UsersRepo.(*repository.Repositories).IUsersRepository.(*repository.UsersRepository).FindByEmail("sdf")*/
	u, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, u)
}

func (h *Handler) profileUpdate(c *gin.Context) {
	var input *models.UserPartial

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	uData, _ := c.Get(userCtx)
	u := uData.(*models.User)

	if input.Email != nil {
		if existedUser, _ := h.services.Users.FindByEmail(*input.Email); existedUser != nil && existedUser.ID != u.ID {
			newErrorResponse(c, http.StatusBadRequest, validation.Errors{"email": helpers.ErrExistUserEmail})
			return
		}
	}

	if err := h.services.Profile.UpdateProfile(u, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Set(userCtx, u)
	c.JSON(http.StatusOK, u)
}

func (h *Handler) profileChangePassword(c *gin.Context) {
	var input *models.ProfileChangePassword
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	uData, _ := c.Get(userCtx)
	u := uData.(*models.User)
	u, _ = h.services.Users.FindByEmail(*u.Email)
	if err := u.ValidateCurrentPassword(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.Users.ChangeUserPassword(u, input.Password); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

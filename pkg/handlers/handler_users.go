package handlers

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repositories"
	"math"
	"net/http"
	"strconv"
)

// @Summary CreateUser
// @Security AccessTokenHeaderName
// @Tags users
// @Description Create user
// @ID createUser
// @Accept json
// @Produce json
// @Param inputBody body models.UserWithPassword true "create user params"
// @Success 200 {object} models.UserWithPassword
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var input *models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if existedUser, _ := h.services.Users.FindByEmail(*input.Email); existedUser != nil {
		newErrorResponse(c, http.StatusBadRequest, validation.Errors{"email": helpers.ErrExistUserEmail})
		return
	}

	if err := h.services.Users.Create(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if err := h.services.Mailing.SendRegistrationMessage(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, input)
}

// @Summary UpdateUser
// @Security AccessTokenHeaderName
// @Tags users
// @Description Update user
// @ID updateUser
// @Accept json
// @Produce json
// @Param inputBody body models.User true "update user params"
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	var input *models.UserPartial

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	u, err := h.services.Users.Find(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	//input["email"].(string)
	if input.Email != nil {
		if existedUser, _ := h.services.Users.FindByEmail(*input.Email); existedUser != nil && existedUser.ID != u.ID {
			newErrorResponse(c, http.StatusBadRequest, validation.Errors{"Email": helpers.ErrExistUserEmail})
			return
		}
	}

	if err := h.services.Users.Update(u, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	uData, _ := c.Get(userCtx)
	if userAuthorized := uData.(*models.User); userAuthorized.ID == u.ID {
		c.Set(userCtx, u)
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) viewUser(c *gin.Context) {
	u, err := h.services.Users.Find(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

// @Summary ListUsers
// @Security AccessTokenHeaderName
// @Tags users
// @Description List users
// @ID listUsers
// @Accept json
// @Produce json
// @Param pageNumber query integer false "Page number" example(0)
// @Param pageSize query integer false "Page size" example(12)
// @Param sex query int false "Sex(Male=1, Female=2)" Enums(1, 2)
// @Success 200 {array} models.User
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users [get]
func (h *Handler) listUsers(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "12"))
	p := &repositories.PaginationRequest{pageNumber, pageSize}

	sortField := c.DefaultQuery("sortField", "id")
	sortDirection := repositories.SortType(c.DefaultQuery("sortDirection", string(repositories.SORT_ASC)))
	s := &repositories.Sort{sortField, sortDirection}

	f := &models.UserFilter{}
	if Name, ok := c.GetQuery("name"); ok {
		f.Name = &Name
	}
	if Email, ok := c.GetQuery("email"); ok {
		f.Email = &Email
	}
	if UpdatedAtFrom, ok := c.GetQuery("updatedAtFrom"); ok {
		f.UpdatedAtFrom = &UpdatedAtFrom
	}
	if UpdatedAtTo, ok := c.GetQuery("updatedAtTo"); ok {
		f.UpdatedAtTo = &UpdatedAtTo
	}
	if CreatedAtFrom, ok := c.GetQuery("createdAtFrom"); ok {
		f.CreatedAtFrom = &CreatedAtFrom
	}
	if CreatedAtTo, ok := c.GetQuery("createdAtTo"); ok {
		f.CreatedAtTo = &CreatedAtTo
	}

	if StatusValues, ok := c.GetQueryArray("status"); ok && len(StatusValues) > 0 {
		f.Status = &[]models.StatusType{}
		for _, val := range StatusValues {
			*f.Status = append(*f.Status, models.StatusType(val))
		}
	}

	if SexValues, ok := c.GetQueryArray("sex"); ok && len(SexValues) > 0 {
		f.Sex = &[]models.SexType{}
		for _, val := range SexValues {
			intVal, _ := strconv.Atoi(val)
			*f.Sex = append(*f.Sex, models.SexType(intVal))
		}
	}

	if Age, ok := c.GetQuery("age"); ok {
		age, err := strconv.Atoi(Age)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
		f.Age = &age
	}

	list, err := h.services.Users.List(p, s, f)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	count, err := h.services.Users.Count(f)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	pagination := new(repositories.PaginationResponse)
	pagination.PageNumber = pageNumber
	pagination.PageSize = pageSize
	pagination.Total = *count
	//fmt.Println(math.Round(x*100) / 100)
	pagination.PageCount = int(math.Ceil(float64(pagination.Total) / float64(pagination.PageSize)))

	c.JSON(http.StatusOK, repositories.List[*models.User]{&list, pagination})
}

func (h *Handler) deleteUser(c *gin.Context) {
	u, err := h.services.Users.Find(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.Users.Delete(u); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) changeUserPassword(c *gin.Context) {
	var input *models.UserChangePassword

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	u, err := h.services.Users.Find(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := u.ValidateNewPassword(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.Users.ChangeUserPassword(u, input.Password); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

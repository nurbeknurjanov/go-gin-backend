package handler

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repository"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/service"
	"math"
	"net/http"
	"strconv"
)

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

	if existedUser, _ := h.services.IUsersService.(*service.UsersService).UsersRepo.FindByEmail(*input.Email); existedUser != nil {
		newErrorResponse(c, http.StatusBadRequest, validation.Errors{"email": helpers.ErrExistUserEmail})
		return
	}

	if err := h.services.IUsersService.CreateUser(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, input)
}

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

	u, err := h.services.IUsersService.FindUser(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	//input["email"].(string)
	if input.Email != nil {
		if existedUser, _ := h.services.IUsersService.(*service.UsersService).UsersRepo.FindByEmail(*input.Email); existedUser != nil && existedUser.ID != u.ID {
			newErrorResponse(c, http.StatusBadRequest, validation.Errors{"Email": helpers.ErrExistUserEmail})
			return
		}
	}

	if err := h.services.IUsersService.UpdateUser(u, input); err != nil {
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
	u, err := h.services.IUsersService.FindUser(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) listUsers(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "12"))
	p := &repository.PaginationRequest{pageNumber, pageSize}

	sortField := c.DefaultQuery("sortField", "id")
	sortDirection := repository.SortType(c.DefaultQuery("sortDirection", string(repository.SORT_ASC)))
	s := &repository.Sort{sortField, sortDirection}

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

	list, err := h.services.IUsersService.ListUsers(p, s, f)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	count, err := h.services.IUsersService.CountUsers(f)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	pagination := new(repository.PaginationResponse)
	pagination.PageNumber = pageNumber
	pagination.PageSize = pageSize
	pagination.Total = *count
	//fmt.Println(math.Round(x*100) / 100)
	pagination.PageCount = int(math.Ceil(float64(pagination.Total) / float64(pagination.PageSize)))

	c.JSON(http.StatusOK, repository.List[*models.User]{&list, pagination})
}

func (h *Handler) deleteUser(c *gin.Context) {
	u, err := h.services.IUsersService.FindUser(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.IUsersService.DeleteUser(u); err != nil {
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

	u, err := h.services.IUsersService.FindUser(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := u.ValidateNewPassword(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.IUsersService.ChangeUserPassword(u, input.Password); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

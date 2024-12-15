package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"math"
	"net/http"
	"strconv"
)

func (h *Handler) createProduct(c *gin.Context) {
	var input *models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.CreateProduct(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *Handler) updateProduct(c *gin.Context) {
	var input *models.ProductPartial

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	m, err := h.services.FindProduct(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.UpdateProduct(m, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, m)
}

func (h *Handler) viewProduct(c *gin.Context) {
	m, err := h.services.FindProduct(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, m)
}

func (h *Handler) listProducts(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "12"))
	p := &repositories.PaginationRequest{pageNumber, pageSize}

	sortField := c.DefaultQuery("sortField", "id")
	sortDirection := repositories.SortType(c.DefaultQuery("sortDirection", string(repositories.SORT_ASC)))
	s := &repositories.Sort{sortField, sortDirection}

	f := &models.ProductFilter{}
	if Name, ok := c.GetQuery("name"); ok {
		f.Name = &Name
	}
	if Description, ok := c.GetQuery("description"); ok {
		f.Description = &Description
	}

	list, err := h.services.ListProducts(p, s, f)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	count, err := h.services.CountProducts(f)
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

	c.JSON(http.StatusOK, repositories.List[*models.Product]{&list, pagination})
}

func (h *Handler) deleteProduct(c *gin.Context) {
	m, err := h.services.FindProduct(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.DeleteProduct(m); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, m)
}

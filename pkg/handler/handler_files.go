package handler

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) createFile(c *gin.Context) {
	var input models.File

	if err := c.ShouldBind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	/*if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}*/

	//_, fileReader, _ := c.Request.FormFile("fileField")
	fileReader, _ := c.FormFile("fileField")

	input.OriginalFileName = &fileReader.Filename
	parts := strings.Split(*input.OriginalFileName, ".")
	input.Ext = &parts[len(parts)-1]
	uuid := uuid.New().String()
	input.UUID = &uuid

	newFileName := strings.Join([]string{*input.UUID, *input.Ext}, ".")

	if err := c.SaveUploadedFile(fileReader, "public/upload/"+newFileName); err != nil {
		newErrorResponse(c, http.StatusBadRequest, validation.Errors{"email": err})
		return
	}

	if err := h.services.IFilesService.CreateFile(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	//j, _ := json.Marshal(input.Data)
	c.JSON(http.StatusOK, &input)
}

func (h *Handler) listFiles(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "12"))
	p := &repositories.PaginationRequest{pageNumber, pageSize}

	sortField := c.DefaultQuery("sortField", "id")
	sortDirection := repositories.SortType(c.DefaultQuery("sortDirection", string(repositories.SORT_ASC)))
	s := &repositories.Sort{sortField, sortDirection}

	f := &models.FileFilter{}
	if ModelName, ok := c.GetQuery("modelName"); ok {
		f.ModelName = &ModelName
	}
	if ModelId, ok := c.GetQuery("modelId"); ok {
		modelId, _ := strconv.Atoi(ModelId)
		f.ModelId = &modelId
	}

	list, err := h.services.ListFiles(p, s, f)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	count, err := h.services.CountFiles(f)
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

	c.JSON(http.StatusOK, repositories.List[*models.File]{&list, pagination})
}

func (h *Handler) deleteFile(c *gin.Context) {
	m, err := h.services.FindFile(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.DeleteFile(m); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, m)
}

func (h *Handler) viewFile(c *gin.Context) {
	m, err := h.services.FindFile(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &m)
}

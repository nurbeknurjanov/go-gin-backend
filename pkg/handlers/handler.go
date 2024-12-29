package handlers

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/nurbeknurjanov/go-gin-backend/docs"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/services"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	errNotAuthorized = errors.New("Not authorized")
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.StaticFS("/public", gin.Dir("public/upload", false))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	//corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	//corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-Access-Token", "X-Refresh-Token"}
	router.Use(cors.New(corsConfig))

	router.POST("/auth/login", h.login)
	router.GET("/auth/get-access-token", h.hasRefreshToken, h.getAccessToken)
	router.POST("/auth/test", h.test)

	users := router.Group("/users", h.authorizedUser)
	{
		users.GET("", h.listUsers)
		users.POST("", h.createUser)
		users.GET("/:id", h.viewUser)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)

		users.PUT("/:id/change-password", h.changeUserPassword)
	}

	profile := router.Group("/profile", h.authorizedUser)
	{
		profile.GET("", h.profile)
		profile.POST("", h.profileUpdate)
		profile.PUT("/change-password", h.profileChangePassword)
	}

	products := router.Group("/products")
	{
		products.GET("", h.listProducts)
		products.POST("", h.authorizedUser, h.createProduct)
		products.GET("/:id", h.viewProduct)
		products.PUT("/:id", h.authorizedUser, h.updateProduct)
		products.DELETE("/:id", h.authorizedUser, h.deleteProduct)
	}

	files := router.Group("/files", h.authorizedUser)
	{
		files.POST("/upload", h.createFile)
		files.GET("", h.listFiles)
		files.DELETE("/:id", h.deleteFile)
		files.GET("/:id", h.viewFile)
	}

	return router
}

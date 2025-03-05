package main

import (
	"ms-categories/cmd/api/controllers"
	"ms-categories/internal/repositories"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.GET("/", func(ctx *gin.Context) {
		controllers.ListCategories(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.PUT("/:id", func(ctx *gin.Context) {
		controllers.UpdateCategory(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.DELETE("/:id", func(ctx *gin.Context) {
		controllers.DeleteCategory(ctx, inMemoryCategoryRepository)
	})
}

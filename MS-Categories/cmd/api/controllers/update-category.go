package controllers

import (
	"ms-categories/internal/repositories"
	use_cases "ms-categories/internal/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	useCase := use_cases.NewUpdateCategoryUseCase(repository)

	err := useCase.Execute(id, body.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

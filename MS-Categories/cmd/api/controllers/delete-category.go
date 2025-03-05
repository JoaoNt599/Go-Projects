package controllers

import (
	"log"
	"ms-categories/internal/repositories"
	use_cases "ms-categories/internal/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	id := ctx.Param("id")

	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": "ID is required"})
		return
	}

	useCase := use_cases.NewDeleteCategoryUseCase(repository)

	err := useCase.Execute(id)
	if err != nil {
		log.Println("Delete Error:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

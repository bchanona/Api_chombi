package controllers

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/gin-gonic/gin"
)

type GetPdfController struct {
	useCase *usecases.GetUrlPdfUseCase
}

func NewGetPdfController(useCase *usecases.GetUrlPdfUseCase) *GetPdfController {
	return &GetPdfController{useCase: useCase}
}

func (controller *GetPdfController) Execute(ctx *gin.Context) {
	urls, err := controller.useCase.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"pdf_urls": urls})
}

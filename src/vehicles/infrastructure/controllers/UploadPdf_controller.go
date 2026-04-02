package controllers

import (
	"time"

	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/gin-gonic/gin"
)

type UploadPdfController struct {
	useCase *usecases.UploadPdfUseCase
}

func NewUploadPdfController(useCase *usecases.UploadPdfUseCase) *UploadPdfController {
	return &UploadPdfController{useCase: useCase}
}

func (controller *UploadPdfController) Execute(ctx * gin.Context){
	//Realizar la carga del archivo PDF
	file, header, err := ctx.Request.FormFile("pdf")
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "PDF file is required",
		})
		return
	}
	defer file.Close()
	
	err = controller.useCase.Execute(file, header.Filename)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to upload PDF",
		})
		return
	}
	time.Sleep(5 * time.Second) // Simular tiempo de procesamiento
	
	ctx.JSON(200, gin.H{
		"message": "PDF uploaded successfully",
	})
}
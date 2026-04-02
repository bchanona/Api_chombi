package usecases

import (
	"mime/multipart"

	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/ports"
)

type UploadPdfUseCase struct {
	repo ports.ImageStorage
}

func NewUploadPdfUseCase(repo ports.ImageStorage) *UploadPdfUseCase {
	return &UploadPdfUseCase{repo: repo}
}

func (useCase *UploadPdfUseCase) Execute(file multipart.File, fileName string) error{
	_, err := useCase.repo.UploadPdf(file, fileName)
	if err != nil {
		return err
	}
	return nil
}
package usecases

import "github.com/bchanona/Api_chombi.git/src/vehicles/domain/ports"

type GetUrlPdfUseCase struct {
	imageStorage ports.ImageStorage
}

func NewGetUrlPdfUseCase(imageStorage ports.ImageStorage) *GetUrlPdfUseCase {
	return &GetUrlPdfUseCase{imageStorage: imageStorage}
}

func (useCase *GetUrlPdfUseCase) Execute() ([]string, error){
	urls, err := useCase.imageStorage.GetPdfUrl()
	return urls, err
}
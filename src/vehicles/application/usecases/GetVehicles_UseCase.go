package usecases

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/output"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/repositories"
)

type GetVehiclesUseCase struct {
	db repositories.IVehicleRepository
}

func NewGetVehiclesUseCase(db repositories.IVehicleRepository) *GetVehiclesUseCase {
	return &GetVehiclesUseCase{db: db}
}

func (useCase *GetVehiclesUseCase) Execute(userId string) ([]output.VehicleResponse, error) {
	result, err := useCase.db.GetVehicles(userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
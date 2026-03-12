package usecases

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/output"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/repositories"
)

type GetVehicleHistoryUseCase struct {
	db repositories.IVehicleRepository
}

func NewGetVehicleHistoryUseCase(db repositories.IVehicleRepository) *GetVehicleHistoryUseCase {
	return &GetVehicleHistoryUseCase{db: db}
}

func (useCase *GetVehicleHistoryUseCase) Execute(userId string) ([]output.VehicleShiftHistory, error) {
	result, err := useCase.db.GetVehiclesByUserId(userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
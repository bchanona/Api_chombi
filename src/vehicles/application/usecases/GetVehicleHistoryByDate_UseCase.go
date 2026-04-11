package usecases

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/output"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/repositories"
)

type GetVehicleHistoryByDateUseCase struct {
	VehicleRepository repositories.IVehicleRepository
}

func NewGetVehicleHistoryByDateUseCase(vehicleRepository repositories.IVehicleRepository) *GetVehicleHistoryByDateUseCase {
	return &GetVehicleHistoryByDateUseCase{VehicleRepository: vehicleRepository}
}

func (uc *GetVehicleHistoryByDateUseCase) Execute(userId string, date string) ([]output.VehicleShiftHistory, error) {
	history, err := uc.VehicleRepository.GetVehicleHistoryByDate(userId, date)
	if err != nil {
		return nil, err
	}
	return history, nil
}

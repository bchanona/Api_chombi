package usecases

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/output"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/repositories"
)

type GetVehicleByUnitNumberUseCase struct {
	VehicleRepository repositories.IVehicleRepository
}

func NewGetVehicleByUnitNumberUseCase(vehicleRepository repositories.IVehicleRepository) *GetVehicleByUnitNumberUseCase {
	return &GetVehicleByUnitNumberUseCase{VehicleRepository: vehicleRepository}
}

func (uc *GetVehicleByUnitNumberUseCase) Execute(userId string, unitNumber int) (output.VehicleResponse, error) {
	vehicle, err := uc.VehicleRepository.GetVehicleByUnitNumber(userId, unitNumber)
	if err != nil {
		return output.VehicleResponse{}, err
	}
	return vehicle, nil
}

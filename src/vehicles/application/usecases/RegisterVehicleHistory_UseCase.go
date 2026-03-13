package usecases

import "github.com/bchanona/Api_chombi.git/src/vehicles/domain/repositories"

type RegisterVehicleHistoryUseCase struct {
	VehicleRepository repositories.IVehicleRepository
}

func NewRegisterVehicleHistoryUseCase(vehicleRepository repositories.IVehicleRepository) *RegisterVehicleHistoryUseCase {
	return &RegisterVehicleHistoryUseCase{VehicleRepository: vehicleRepository}
}

func (uc *RegisterVehicleHistoryUseCase) Execute() error{
	err := uc.VehicleRepository.RegisterVehicleShiftHistory()
	if err != nil {
		return err
	}
	return nil
}
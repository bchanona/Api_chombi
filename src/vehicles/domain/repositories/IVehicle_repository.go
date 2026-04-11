package repositories

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/input"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/output"
)

type IVehicleRepository interface {
	RegisterVehicle(input.VehicleRequest) error
	RegisterVehicleShiftHistory() error
	GetVehiclesByUserId(userId string) ([]output.VehicleShiftHistory, error)
	GetVehicles(userId string) ([]output.VehicleResponse, error)
	UpdateVehicle(userId string, vehicleId string, vehicleData input.UpdateVehicleRequest) error
	DeleteVehicle(userId string, vehicleId string) error
	GetVehicleByUnitNumber(userId string, unitNumber int) (output.VehicleResponse, error)
	GetVehicleHistoryByDate(userId string, date string) ([]output.VehicleShiftHistory, error)
}

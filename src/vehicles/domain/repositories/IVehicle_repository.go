package repositories

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/input"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/output"
)


type IVehicleRepository interface {
	RegisterVehicle(input.VehicleRequest) error
	GetVehiclesByUserId(userId string)([]output.VehicleShiftHistory, error)
	GetVehicles(userId string)([]output.VehicleResponse, error)
}
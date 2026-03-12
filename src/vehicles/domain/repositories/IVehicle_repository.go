package repositories

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/input"
)


type IVehicleRepository interface {
	RegisterVehicle(input.VehicleRequest) error

}
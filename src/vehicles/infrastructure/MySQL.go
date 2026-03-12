package infrastructure

import (
	"database/sql"

	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/input"
	"github.com/bchanona/Api_chombi.git/src/vehicles/infrastructure/queries"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) RegisterVehicle(vehicle input.VehicleRequest) error {
	_, err := sql.db.Exec(queries.RegisterVehicleQuery,
		vehicle.Id, vehicle.LicensePlate,
		vehicle.UnitNumber, vehicle.Shift,
		vehicle.IsWorking, vehicle.ImageURL,
		vehicle.NumberOfPassengers,
		vehicle.Model,
		vehicle.UserId,
		vehicle.DriverName)

	if err != nil {
		return err
	}
	return nil
}

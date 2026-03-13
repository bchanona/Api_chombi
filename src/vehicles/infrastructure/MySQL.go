package infrastructure

import (
	"database/sql"

	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/input"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/output"
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

func (sql *MySQL) GetVehiclesByUserId(userId string) ([]output.VehicleShiftHistory, error) {
	rows, err := sql.db.Query(queries.GetAllVehicleShiftHistory, userId)

	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var vehicleShiftHistory []output.VehicleShiftHistory

	for rows.Next() {
		var vehicle output.VehicleShiftHistory
		err := rows.Scan(&vehicle.Id, &vehicle.DirverName, &vehicle.LicensePlate, &vehicle.ShiftOrder, &vehicle.Date)
		if err != nil {
			return nil, err
		}
		vehicleShiftHistory = append(vehicleShiftHistory, vehicle)
	}

	return vehicleShiftHistory, nil

}

func (sql *MySQL) GetVehicles(userId string) ([]output.VehicleResponse, error) {
	rows, err := sql.db.Query(queries.GetAllVehiclesQuery, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []output.VehicleResponse

	for rows.Next(){
		var vehicle output.VehicleResponse
		err := rows.Scan(
			&vehicle.Id,
			&vehicle.LicensePlate,
			&vehicle.UnitNumber,
			&vehicle.Shift,
			&vehicle.IsWorking,
			&vehicle.ImageURL,
			&vehicle.Model,
			&vehicle.DriverName,
		)
		if err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

func (sql *MySQL) RegisterVehicleShiftHistory() error {
	_, err := sql.db.Exec(queries.RegisterVehicleShiftHistory)
	
	if err != nil {
		return err
	}
	return nil

}

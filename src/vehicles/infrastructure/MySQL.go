package infrastructure

import (
	"database/sql"
	"fmt"

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

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicleShiftHistory []output.VehicleShiftHistory

	for rows.Next() {
		var vehicle output.VehicleShiftHistory
		err := rows.Scan(&vehicle.Id, &vehicle.DriverName, &vehicle.LicensePlate, &vehicle.ShiftOrder, &vehicle.Date)
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

	for rows.Next() {
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

func (sql *MySQL) UpdateVehicle(userId string, vehicleId string, vehicleData input.UpdateVehicleRequest) error {
	result, err := sql.db.Exec(queries.UpdateVehicleQuery,
		vehicleData.DriverName,
		vehicleData.LicensePlate,
		vehicleData.Model,
		vehicleData.UnitNumber,
		vehicleData.ImageURL,
		vehicleId,
		userId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("vehicle not found or access denied")
	}

	return nil
}

func (sql *MySQL) DeleteVehicle(userId string, vehicleId string) error {
	result, err := sql.db.Exec(queries.DeleteVehicleQuery,
		vehicleId,
		userId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("vehicle not found or access denied")
	}

	return nil
}

func (m *MySQL) GetVehicleByUnitNumber(userId string, unitNumber int) (output.VehicleResponse, error) {
	var vehicle output.VehicleResponse

	err := m.db.QueryRow(queries.GetVehicleByUnitNumberQuery, unitNumber, userId).Scan(
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
		if err == sql.ErrNoRows {
			return output.VehicleResponse{}, fmt.Errorf("vehicle not found")
		}
		return output.VehicleResponse{}, err
	}

	return vehicle, nil
}

func (sql *MySQL) GetVehicleHistoryByDate(userId string, date string) ([]output.VehicleShiftHistory, error) {
	rows, err := sql.db.Query(queries.GetHistoryByDateQuery, userId, date)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicleHistory []output.VehicleShiftHistory

	for rows.Next() {
		var vehicle output.VehicleShiftHistory
		err := rows.Scan(&vehicle.Id, &vehicle.DriverName, &vehicle.LicensePlate, &vehicle.ShiftOrder, &vehicle.Date)
		if err != nil {
			return nil, err
		}
		vehicleHistory = append(vehicleHistory, vehicle)
	}

	if len(vehicleHistory) == 0 {
		return nil, fmt.Errorf("no history found for the specified date")
	}

	return vehicleHistory, nil
}

package queries

const (
	RegisterVehicleQuery = `
		INSERT INTO Vehicles (
			id,
			licensePlate,
			unitNumber,
			shift,
			isWorking,
			image_url,
			number_of_passenggers,
			model,
			user_id,
			driver_name
		)
		VALUES (
			UUID_TO_BIN(?),
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			UUID_TO_BIN(?),
			?
		)
	`

	GetAllVehicleShiftHistory = `
		SELECT
			BIN_TO_UUID(v.id) AS vehicle_id,
			v.driver_name,
			v.licensePlate,
			h.shift_order,
			h.snapshot_date
		FROM vehicle_shift_history h
		JOIN Vehicles v 
			ON v.id = h.vehicle_id
		WHERE v.user_id = UUID_TO_BIN(?)
		ORDER BY 
			h.snapshot_date DESC,
			h.shift_order ASC
	`
)
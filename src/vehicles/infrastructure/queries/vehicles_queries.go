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

	GetAllVehiclesQuery = ` 
	SELECT 
		BIN_TO_UUID(v.id) AS id,
		v.licensePlate,
		v.unitNumber,
		v.shift,
		v.isWorking,
		v.image_url,
		v.model,
		v.driver_name
		FROM Vehicles v
		WHERE v.user_id = UUID_TO_BIN(?);
	`
	RegisterVehicleShiftHistory = `
	INSERT INTO vehicle_shift_history (
    	id,
    	vehicle_id,
    	shift_order,
    	snapshot_date
	)
	SELECT
	UUID_TO_BIN(UUID()),
	id,
	shift,
	CURDATE()
	FROM Vehicles;

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

	UpdateVehicleQuery = `
	UPDATE Vehicles
	SET 
		driver_name = ?,
		licensePlate = ?,
		model = ?,
		unitNumber = ?,
		image_url = COALESCE(?, image_url)
	WHERE 
	id = UUID_TO_BIN(?) AND user_id = UUID_TO_BIN(?)
	`

	DeleteVehicleQuery = `
	DELETE FROM Vehicles
	WHERE id = UUID_TO_BIN(?) AND user_id = UUID_TO_BIN(?)
	`

	GetVehicleByUnitNumberQuery = `
	SELECT 
		BIN_TO_UUID(v.id) AS id,
		v.licensePlate,
		v.unitNumber,
		v.shift,
		v.isWorking,
		v.image_url,
		v.model,
		v.driver_name
	FROM Vehicles v
	WHERE v.unitNumber = ? AND v.user_id = UUID_TO_BIN(?)
	`

	GetHistoryByDateQuery = `
	SELECT
		BIN_TO_UUID(v.id) AS vehicle_id,
		v.driver_name,
		v.licensePlate,
		h.shift_order,
		h.snapshot_date
	FROM vehicle_shift_history h
	JOIN Vehicles v 
		ON v.id = h.vehicle_id
	WHERE v.user_id = UUID_TO_BIN(?) AND DATE(h.snapshot_date) = ?
	ORDER BY 
		h.shift_order ASC
	`
)

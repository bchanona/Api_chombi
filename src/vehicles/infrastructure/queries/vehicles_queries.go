package queries

const (
	RegisterVehicleQuery = "INSERT INTO Vehicles (id,licensePlate,unitNumber,shift,isWorking,image_url,number_of_passenggers, model,user_id, driver_name) VALUES (UUID_TO_BIN(?),?,?,?,?,?,?,?,UUID_TO_BIN(?),?)"
)


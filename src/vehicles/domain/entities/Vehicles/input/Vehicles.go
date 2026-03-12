package input

type VehicleRequest struct {
	Id                 string `json:"id" form:"id"`
	DriverName         string `json:"driver_name" form:"driver_name" binding:"required"`
	LicensePlate       string `json:"license_plate" form:"license_plate" binding:"required"`
	Shift              int    `json:"shift" form:"shift"`
	Model              string `json:"model" form:"model" binding:"required"`
	UnitNumber         int    `json:"unit_number" form:"unit_number" binding:"required"`
	IsWorking          bool   `json:"is_working" form:"is_working"`
	NumberOfPassengers int    `json:"number_of_passengers" form:"number_of_passengers"`
	ImageURL           string `json:"image_url" form:"image_url"`
	UserId             string `json:"user_id"`
}
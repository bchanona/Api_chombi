package output

type VehicleShiftHistory struct {
	Id string `json:"id"`
	DriverName string `json:"driver_name"`
	LicensePlate string `json:"license_plate"`
	ShiftOrder int `json:"shift_order"`
	Date string `json:"date"`	
}

type VehicleResponse struct {
	Id				 string `json:"id"`
	LicensePlate	 string `json:"license_plate"`
	UnitNumber		 int    `json:"unit_number"`
	Shift 			 int    `json:"shift"`
	IsWorking		 bool   `json:"is_working"`
	ImageURL		 string `json:"image_url"`
	Model			 string `json:"model"`
	DriverName		 string `json:"driver_name"`
	
}
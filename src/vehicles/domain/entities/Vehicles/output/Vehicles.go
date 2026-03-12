package output

type VehicleShiftHistory struct {
	Id string `json:"id"`
	DirverName string `json:"driver_name"`
	LicensePlate string `json:"license_plate"`
	ShiftOrder int `json:"shift_order"`
	Date string `json:"date"`	
}
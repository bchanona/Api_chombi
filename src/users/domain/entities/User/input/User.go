package input


type UserRequest struct {
	ID string `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required, min=8"`
	RolID string `json:"rol_id"`
}
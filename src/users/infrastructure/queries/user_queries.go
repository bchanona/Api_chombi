package queries

const (
	RegisterUserQuery = "INSERT INTO Users (id, firstName, lastName, email, password, role_id) values (UUID_TO_BIN(?),?,?,?,?,UUID_TO_BIN(?))"
	FindByEmailQuery  = "SELECT BIN_TO_UUID(u.id), u.firstName, u.lastName,u.email,u.password,r.name AS role_name FROM Users u JOIN Roles r ON u.role_id = r.id WHERE u.email = ?"
)

package queries

const(
	RegisterUserQuery = "INSERT INTO users (id, firstName, lastName, email, password, role_id) values (UUID_TO_BIN(?),?,?,?,?,UUID_TO_BIN(?))"
)
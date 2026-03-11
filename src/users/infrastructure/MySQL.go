package infrastructure

import (
	"database/sql"

	"github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/input"
	"github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/output"
	"github.com/bchanona/Api_chombi.git/src/users/infrastructure/queries"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) RegisterUser(user input.UserRequest) error {

	_, err := sql.db.Exec(queries.RegisterUserQuery, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.RolID)
	return err

}

func (sql *MySQL) FindByEmail(email string) (output output.UserLoginDBResponse, err error) {
	row := sql.db.QueryRow(queries.FindByEmailQuery, email)
	err = row.Scan(&output.Id, &output.FirstName, &output.LastName, &output.Email, &output.Password, &output.RoleName)
	return
}

// services/user_service.go
package services

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"smartschoolsystem.go/utils"
)

func AuthenticateUser(email, password string) (string, error) {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return "", err
	}
	defer db.Close()

	var storedPwd, role string
	err = db.QueryRow("SELECT password, role FROM userlogin WHERE email = ?", email).Scan(&storedPwd, &role)
	if err != nil || storedPwd != password {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateJWT(email, role)
}

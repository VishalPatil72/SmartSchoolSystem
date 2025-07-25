package services

import (
	"smartschoolsystem.go/database"
	"smartschoolsystem.go/models"
)

func GetAllDivisions() ([]models.Division, error) {
	db := database.InitMySQL()
	rows, err := db.Query("SELECT divisionId, divisionName FROM divisionmaster")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var divisions []models.Division
	for rows.Next() {
		var d models.Division
		if err := rows.Scan(&d.DivisionID, &d.DivisionName); err != nil {
			return nil, err
		}
		divisions = append(divisions, d)
	}
	return divisions, nil
}

func GetDivisionById(divisionId uint) (models.Division, error) {
	db := database.InitMySQL()
	var d models.Division

	err := db.QueryRow("SELECT divisionId, divisionName FROM divisionmaster WHERE divisionId = ?", divisionId).Scan(&d.DivisionID, &d.DivisionName)
	if err != nil {
		return models.Division{}, err
	}
	return d, nil
}
func CreateDivision(division models.Division) (uint, error) {
	db := database.InitMySQL()
	qry := `INSERT INTO divisionmaster
				( 
				divisionName
				)
				values(?)`
	result, err := db.Exec(qry, division.DivisionName)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(lastInsertId), nil
}
func UpdateDivision(division models.Division) error {

	db := database.InitMySQL()
	qry := `UPDATE divisionmaster
				SET divisionName = ?		
				WHERE divisionId = ?`
	_, err := db.Exec(qry, division.DivisionName, division.DivisionID)
	return err
}
func DeleteDivision(divisionId uint) error {
	db := database.InitMySQL()
	qry := `DELETE FROM divisionmaster WHERE divisionId = ?`
	_, err := db.Exec(qry, divisionId)
	if err != nil {
		return err
	}
	return nil
}

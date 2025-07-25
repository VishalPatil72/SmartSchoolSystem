package services

import (
	"smartschoolsystem.go/database"
	"smartschoolsystem.go/models"
)

func GetAllClasses() ([]models.Class, error) {

	db := database.InitMySQL()
	rows, err := db.Query("SELECT classId, className FROM classmaster")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []models.Class
	for rows.Next() {
		var c models.Class
		if err := rows.Scan(&c.ClassID, &c.ClassName); err != nil {
			return nil, err
		}
		classes = append(classes, c)
	}
	return classes, nil
}

func GetClassById(classId uint) (models.Class, error) {
	db := database.InitMySQL()
	var c models.Class

	err := db.QueryRow("SELECT classId, className FROM classmaster WHERE classId = ?", classId).Scan(&c.ClassID, &c.ClassName)
	if err != nil {
		return models.Class{}, err
	}
	return c, nil
}
func CreateClass(class models.Class) (uint, error) {
	db := database.InitMySQL()
	qry := `INSERT INTO classmaster
				( 
				className
				)
				values(?)`
	result, err := db.Exec(qry, class.ClassName)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(lastInsertId), nil
}

func UpdateClass(class models.Class) error {
	db := database.InitMySQL()
	qry := `UPDATE classmaster
				SET className = ?		
				WHERE classId = ?`
	_, err := db.Exec(qry, class.ClassName, class.ClassID)
	if err != nil {
		return err
	}
	return nil
}
func DeleteClass(classId uint) error {
	db := database.InitMySQL()
	qry := `DELETE FROM classmaster WHERE classId = ?`
	_, err := db.Exec(qry, classId)
	if err != nil {
		return err
	}
	return nil
}

package services

import (
	"smartschoolsystem.go/database"
	"smartschoolsystem.go/models"
)

func GetAllSubjects() ([]models.Subject, error) {
	db := database.InitMySQL()
	rows, err := db.Query("SELECT subjectId, subjectName FROM subjectmaster")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []models.Subject
	for rows.Next() {
		var s models.Subject
		if err := rows.Scan(&s.SubjectID, &s.SubjectName); err != nil {
			return nil, err
		}
		subjects = append(subjects, s)
	}
	return subjects, nil
}
func GetSubjectById(subjectId uint) (models.Subject, error) {
	db := database.InitMySQL()
	var s models.Subject

	err := db.QueryRow("SELECT subjectId, subjectName FROM subjectmaster WHERE subjectId = ?", subjectId).Scan(&s.SubjectID, &s.SubjectName)
	if err != nil {
		return models.Subject{}, err
	}
	return s, nil
}
func CreateSubject(subject models.Subject) (uint, error) {
	db := database.InitMySQL()
	qry := `INSERT INTO subjectmaster
				( 
				subjectName
				)
				values(?)`
	result, err := db.Exec(qry, subject.SubjectName)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(lastInsertId), nil
}
func UpdateSubject(subject models.Subject) error {
	db := database.InitMySQL()
	qry := `UPDATE subjectmaster
				SET subjectName = ?		
				WHERE subjectId = ?`
	_, err := db.Exec(qry, subject.SubjectName, subject.SubjectID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSubject(subjectId uint) error {
	db := database.InitMySQL()
	qry := `DELETE FROM subjectmaster WHERE subjectId = ?`
	_, err := db.Exec(qry, subjectId)
	if err != nil {
		return err
	}
	return nil
}

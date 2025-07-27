package services

import (
	"smartschoolsystem.go/database"
	"smartschoolsystem.go/models"
)

func GetAllStudents() ([]models.Student, error) {
	db := database.InitMySQL()
	rows, err := db.Query(
		`SELECT student_id, last_name, first_name, middle_name, mother_name,gender, date_of_birth, address,
	 			pin_code, mobile, email, photo_path, religion_id, caste_id, sub_caste_id, class_id, 
				created_at, updated_at 
				FROM student`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []models.Student
	for rows.Next() {
		var s models.Student
		if err := rows.Scan(&s.StudentID, &s.LastName, &s.FirstName, &s.MiddleName, &s.MotherName, &s.Gender,
			&s.DateOfBirth, &s.Address, &s.PinCode, &s.Mobile, &s.Email, &s.PhotoPath, &s.ReligionID, &s.CasteID,
			&s.SubCasteID, &s.ClassID, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}
func GetStudentById(studentId int) (models.Student, error) {
	db := database.InitMySQL()
	var s models.Student

	err := db.QueryRow(
		`SELECT student_id, last_name, first_name, middle_name, mother_name,gender, date_of_birth, address,
	 			pin_code, mobile, email, photo_path, religion_id, caste_id, sub_caste_id, class_id, 
				created_at, updated_at 
				FROM student WHERE student_id = ?`, studentId).Scan(
		&s.StudentID, &s.LastName, &s.FirstName, &s.MiddleName, &s.MotherName, &s.Gender,
		&s.DateOfBirth, &s.Address, &s.PinCode, &s.Mobile, &s.Email, &s.PhotoPath, &s.ReligionID, &s.CasteID,
		&s.SubCasteID, &s.ClassID, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return models.Student{}, err
	}
	return s, nil
}

func CreateStudent(student models.Student) (int, error) {
	db := database.InitMySQL()
	qry := `INSERT INTO student
				( 
				last_name, first_name, middle_name, mother_name,gender, date_of_birth, address,
	 			pin_code, mobile, email, photo_path, religion_id, caste_id, sub_caste_id, class_id, 
				created_at, updated_at)
				values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`
	result, err := db.Exec(qry, student.LastName, student.FirstName, student.MiddleName, student.MotherName,
		student.Gender,
		student.DateOfBirth, student.Address, student.PinCode, student.Mobile, student.Email, student.PhotoPath, student.ReligionID, student.CasteID,
		student.SubCasteID, student.ClassID, student.CreatedAt, student.UpdatedAt)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastInsertId), nil
}
func UpdateStudent(student models.Student) error {
	db := database.InitMySQL()
	qry := `UPDATE students SET 
		last_name = ?, first_name = ?, middle_name = ?, mother_name = ?, gender = ?, 
		date_of_birth = ?, address = ?, pin_code = ?, mobile = ?, email = ?, 
		photo_path = ?, religion_id = ?, caste_id = ?, sub_caste_id = ?, class_id = ?, 
		updated_at = ?
		WHERE student_id = ?`
	_, err := db.Exec(qry, student.LastName, student.FirstName, student.MiddleName, student.MotherName,
		student.Gender, student.DateOfBirth, student.Address, student.PinCode, student.Mobile, student.Email,
		student.PhotoPath, student.ReligionID, student.CasteID, student.SubCasteID, student.ClassID,
		student.UpdatedAt, student.StudentID)
	return err
}
func DeleteStudent(studentId int) error {
	db := database.InitMySQL()
	qry := `DELETE FROM student WHERE student_id = ?`
	_, err := db.Exec(qry, studentId)
	if err != nil {
		return err
	}
	return nil
}

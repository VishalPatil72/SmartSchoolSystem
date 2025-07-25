package services

import (
	"smartschoolsystem.go/database"
	"smartschoolsystem.go/models"
)

func GetAllCategories() ([]models.Category, error) {
	db := database.InitMySQL()
	rows, err := db.Query("SELECT categoryId, categoryName FROM categorymaster")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.CategoryID, &c.CategoryName); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}
func GetCategoryById(categoryId uint) (models.Category, error) {
	db := database.InitMySQL()
	var c models.Category

	err := db.QueryRow(("SELECT categoryId, categoryName FROM categorymaster WHERE categoryId = ?"), categoryId).Scan(&c.CategoryID, &c.CategoryName)
	if err != nil {
		return models.Category{}, err
	}
	return c, nil
}
func CreateCategory(category models.Category) (uint, error) {
	db := database.InitMySQL()
	qry := `INSERT INTO categorymaster
				( 
				categoryName
				)
				values(?)`
	result, err := db.Exec(qry, category.CategoryName)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(lastInsertId), nil
}
func UpdateCategory(category models.Category) error {
	db := database.InitMySQL()
	qry := `UPDATE categorymaster
				SET categoryName = ?		
				WHERE categoryId = ?`
	_, err := db.Exec(qry, category.CategoryName, category.CategoryID)
	if err != nil {
		return err
	}
	return nil
}
func DeleteCategory(categoryId uint) error {
	db := database.InitMySQL()
	qry := `DELETE FROM categorymaster WHERE categoryId = ?`
	_, err := db.Exec(qry, categoryId)
	if err != nil {
		return err
	}
	return nil
}

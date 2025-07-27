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
func GetSubCategories() ([]models.SubCategory, error) {
	db := database.InitMySQL()
	rows, err := db.Query("SELECT subCategoryId, subCategoryName, categoryId FROM subcategorymaster")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subCategories []models.SubCategory
	for rows.Next() {
		var sc models.SubCategory
		if err := rows.Scan(&sc.SubCategoryID, &sc.SubCategoryName, &sc.CategoryID); err != nil {
			return nil, err
		}
		subCategories = append(subCategories, sc)
	}
	return subCategories, nil

}
func GetAllReligions() ([]models.Religion, error) {
	db := database.InitMySQL()
	rows, err := db.Query("SELECT religionId, religionName FROM religionmaster")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var religions []models.Religion
	for rows.Next() {
		var r models.Religion
		if err := rows.Scan(&r.ReligionID, &r.ReligionName); err != nil {
			return nil, err
		}
		religions = append(religions, r)
	}
	return religions, nil
}

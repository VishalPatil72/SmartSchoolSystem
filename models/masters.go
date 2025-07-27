package models

type Category struct {
	CategoryID   uint   `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}
type SubCategory struct {
	SubCategoryID   uint   `json:"subCategoryId"`
	SubCategoryName string `json:"subCategoryName"`
	CategoryID      uint   `json:"categoryId"`
}
type Class struct {
	ClassID   uint   `json:"classId"`
	ClassName string `json:"className"`
}

type Division struct {
	DivisionID   uint   `json:"divisionId"`
	DivisionName string `json:"divisionName"`
}
type Subject struct {
	SubjectID   uint   `json:"subjectId"`
	SubjectName string `json:"subjectName"`
}
type Religion struct {
	ReligionID   uint   `json:"religionId"`
	ReligionName string `json:"religionName"`
}

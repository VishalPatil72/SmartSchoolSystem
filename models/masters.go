package models

type Category struct {
	CategoryID   uint   `json:"categoryId"`
	CategoryName string `json:"categoryName"`
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

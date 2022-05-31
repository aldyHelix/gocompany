package models

type Company struct {
	ID                 uint   `json:"id" gorm:"primary_key"`
	CompanyName        string `json:"company_name"`
	CompanyDescription string `json:"company_description"`
}

type InsertCompany struct {
	CompanyName        string `json:"company_name" binding:"required"`
	CompanyDescription string `json:"company_description" binding:"required"`
}

type EditCompany struct {
	CompanyName        string `json:"company_name"`
	CompanyDescription string `json:"company_description"`
}

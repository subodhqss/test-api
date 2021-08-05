package models

import (
	"crypto/rand"
	"math/big"
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmployeeNumber string    `gorm:"column:employeeNumber" json:"employee_number"`
	FirstName      string    `gorm:"column:firstName" json:"first_name"`
	LastName       string    `gorm:"column:lastName" json:"last_name"`
	Email          string    `json:"email"`
	ReportsTo      *Employee `gorm:"foreignKey:ReportsToId;references:EmployeeNumber" json:"reports_to.omitempty"`
	ReportsToId    int       `gorm:"column:reportsTo" json:"reports_to_id"`
	OfficeCode     int       `gorm:"column:officeCode" json:"office_code"`
	Extension      string    `json:"extension"`
	JobTitle       string    `gorm:"column:jobTitle" json:"job_title"`
	Office         *Offices  `gorm:"foreignKey:OfficeCode;references:OfficeCode"`
	Password       string    `gorm:"column:password" json:"password"`
}

type Offices struct {
	OfficeCode   int        `gorm:"column:officeCode" json:"office_code"`
	City         string     `json:"city"`
	Phone        string     `json:"phone"`
	AddressLine1 string     `gorm:"addressLine1" json:"address_line1"`
	Employee     []Employee `gorm:"foreignKey:OfficeCode;references:OfficeCode" json:"employees,omitempty"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	// uuidStr := uuid.New().String()
	a, _ := rand.Int(rand.Reader, new(big.Int).SetInt64(time.Now().UnixNano()))
	e.EmployeeNumber = a.String()[13:]

	return nil
}

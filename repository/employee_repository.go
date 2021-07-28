package repository

import (
	"fmt"
	"log"

	"github.com/subodhqss/test-api/models"
)

type EmployeeRepository interface {
	GetEmployees() ([]*models.Employee, error)
	GetOneEmployeeById(empNo string) (*models.Employee, error)
	AddEmployee(emp *models.Employee) (*models.Employee, error)
	GetEmployeeInOffice(officeCode string) (*models.Offices, error)
	UpsertTwoFA(secret string) (int, error)
}

type employeeRepo struct {
}

const (
	TwoFAUpsertQuery = "REPLACE INTO two_fa(secret) VALUES(?)"
)

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepo{}
}

func (er *employeeRepo) GetEmployees() ([]*models.Employee, error) {
	employees := []*models.Employee{}
	db := GormDB()
	if err := db.Limit(10).Preload("ReportsTo").Find(&employees).Error; err != nil {
		fmt.Println("err in db ", err)
		return nil, err
	}

	return employees, nil
}

func (er *employeeRepo) GetOneEmployeeById(empNo string) (*models.Employee, error) {
	emp := &models.Employee{}
	res := gormDB.Preload("ReportsTo").Preload("Office").Where("employeeNumber", empNo).First(&emp)

	if err := res.Error; err != nil {
		log.Print("Error in getting employee by employeeNumber", err)
		return nil, err
	}

	// office := []*models.Offices{}
	// gormDB.Preload("Employee").Where("officeCode", emp.OfficeCode).Find(&office)

	// b, _ := json.MarshalIndent(emp, " ", " ")
	// fmt.Println(string(b), "office Code ", emp.OfficeCode)
	return emp, nil
}

func (er *employeeRepo) AddEmployee(emp *models.Employee) (*models.Employee, error) {
	res := gormDB.Create(emp)
	if err := res.Error; err != nil {
		log.Print("Error in adding the employee in DB", err)
		return nil, err
	}
	return emp, nil
}

func (er *employeeRepo) UpsertTwoFA(secret string) (int, error) {
	stmt, err := sqlDB.Prepare(TwoFAUpsertQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, resErr := stmt.Exec(secret)

	if resErr != nil {
		return 0, resErr
	}
	insertedId, IdErr := res.LastInsertId()
	return int(insertedId), IdErr
}

func (er *employeeRepo) GetEmployeeInOffice(officeCode string) (*models.Offices, error) {
	office := &models.Offices{}
	res := gormDB.Preload("Employee").Preload("Employee.ReportsTo").Where("officeCode", officeCode).Find(&office)

	if err := res.Error; err != nil {
		log.Print("Error in getting office by code", err)
		return nil, err
	}

	return office, nil
}

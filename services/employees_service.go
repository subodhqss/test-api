package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/subodhqss/test-api/models"
	"github.com/subodhqss/test-api/repository"
)

type EmployeeService interface {
	GetEmployees() ([]*models.Employee, error)
	GetEmployeeById(empNo string) (*models.Employee, error)
	AddEmployee(emp *models.Employee) (*models.Employee, error)
	GetEmployeeInOffice(officeCode string) (*models.Offices, error)
	GetQRCodeImage() (*models.QRCode, error)
	GetTOTPKey() (*otp.Key, error)
}

type employeeSrv struct {
	employeeRepo repository.EmployeeRepository
}

func NewUserService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &employeeSrv{employeeRepo: employeeRepo}
}

func (es *employeeSrv) GetEmployees() ([]*models.Employee, error) {
	return es.employeeRepo.GetEmployees()

	// return []*models.Employee{{ID: "08203", Name: "subodh"}}, nil
}

func (es *employeeSrv) GetEmployeeById(empNo string) (*models.Employee, error) {
	return es.employeeRepo.GetOneEmployeeById(empNo)
}

func (es *employeeSrv) AddEmployee(emp *models.Employee) (*models.Employee, error) {
	return es.employeeRepo.AddEmployee(emp)
}

func (es *employeeSrv) GetEmployeeInOffice(officeCode string) (*models.Offices, error) {
	return es.employeeRepo.GetEmployeeInOffice(officeCode)
}

func (es *employeeSrv) GetQRCodeImage() (*models.QRCode, error) {
	key, err := es.GetTOTPKey()
	if err != nil {
		return nil, err
	}
	qrCode, err := GenerateQRCode(key)
	if err != nil {
		fmt.Println("error in generating the QR code", err)
		return nil, err
	}

	id, idErr := es.employeeRepo.UpsertTwoFA(qrCode.Secret)
	if idErr != nil {
		return nil, idErr
	}
	qrCode.ID = id
	return qrCode, nil
}

func GenerateQRCode(key *otp.Key) (*models.QRCode, error) {

	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		return nil, err
	}
	png.Encode(&buf, img)
	base64String := base64.StdEncoding.EncodeToString(buf.Bytes())
	qrCode := &models.QRCode{Image: base64String, Secret: key.Secret()}
	return qrCode, nil
}

func (es *employeeSrv) GetTOTPKey() (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "wndyr.com",
		AccountName: "subodh@wndyr.com",
	})
	return key, err
}

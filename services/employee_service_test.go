package services

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/test-api/mocks"
	"github.com/subodhqss/test-api/models"
)

func TestLoginEmployee(t *testing.T) {
	mockRepo := new(mocks.EmployeeRepository)
	svc := NewUserService(mockRepo)

	loginData := models.Login{Email: "sk@gmail.com", Password: "test@123"}

	// hasPass, err := bcrypt.GenerateFromPassword([]byte(loginData.Password), 5)
	// if err != nil {
	// 	log.Print("Hass pass :", string(hasPass))
	// 	assert.NotNil(t, err)
	// }
	// log.Print("Hass pass :", string(hasPass))
	data := svc.LoginEmployee(loginData)

	log.Printf("data returned %#v", data)
	assert.NotNil(t, data)
}

func TestGeneratePass(t *testing.T) {
	pass := "test@123"

	err := generatePass(&pass)
	log.Println("Generated Pass : ", pass)
	assert.Nil(t, err)
}

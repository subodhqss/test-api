package services

import (
	"testing"

	"github.com/pquerna/otp"
	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/test-api/mocks"
)

func TestGenerateQRCode(t *testing.T) {
	mockRepo := new(mocks.EmployeeRepository)
	mockEmpSrv := new(mocks.EmployeeService)
	userSvc := NewUserService(mockRepo)

	// qrCode := &models.QRCode{Secret: "JJODHSPXQETK7KVWFNW5DFLCYMXMSKJR"}
	id := 1
	key := &otp.Key{}
	mockEmpSrv.On("GetTOTPKey").Return(key, nil)
	mockRepo.On("UpsertTwoFA", key.Secret()).Return(id, nil)
	code, err := userSvc.GetQRCodeImage()

	assert.Nil(t, err)
	assert.NotNil(t, code)
}

package repository

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	repo := NewEmployeeRepository()

	res, err := repo.GetEmployees()

	b, _ := json.MarshalIndent(res, " ", " ")

	fmt.Println(string(b))

	// assert.Nil(t, res)
	assert.Nil(t, err)
}

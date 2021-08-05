package repository

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitDBConnection()
	os.Exit(m.Run())
}

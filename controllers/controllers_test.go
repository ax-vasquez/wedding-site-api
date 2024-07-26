package controllers

import (
	"log"
	"os"
	"testing"

	"github.com/ax-vasquez/wedding-site-api/models"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panic("Error loading .env file: ", err.Error())
	}
	os.Setenv("TEST_ENV", "true")
	models.Setup()
	models.SeedTestData()

	// This will be 0 if passing, 1 if failing
	exitCode := m.Run()
	// Switch back to "prod" DB so we can drop the test DB
	models.SwitchConnectedDB(os.Getenv("PGSQL_DBNAME"))
	models.DropTestDB()

	// Must return status code - if you don't all tests will be marked as "passing" by returning 0 for all tests
	os.Exit(exitCode)
}

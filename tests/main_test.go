package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set up your testing environment here
	// LoadEnvVariablesTest()
	setupTestingEnvironment()

	// Run all the tests
	exitCode := m.Run()

	// Tear down your testing environment here
	teardownTestingEnvironment()

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}

// Define functions for setting up and tearing down the testing environment
func setupTestingEnvironment() {
	// Set up your test database, environment variables, etc.
	LoadEnvVariablesTest()
	ConnectToPostgresqlTest()
	SyncPostgresqlTest()
}

func teardownTestingEnvironment() {
	// Clean up resources created during testing
}

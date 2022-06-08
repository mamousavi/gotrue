package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/netlify/gotrue/conf"
	"github.com/netlify/gotrue/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MFATestSuite struct {
	suite.Suite
	API    *API
	Config *conf.Configuration

	instanceID uuid.UUID
}

func TestMFA(t *testing.T) {
	api, config, instanceID, err := setupAPIForTestForInstance()
	require.NoError(t, err)

	ts := &MFATestSuite{
		API:        api,
		Config:     config,
		instanceID: instanceID,
	}
	defer api.db.Close()

	suite.Run(t, ts)
}

func (ts *MFATestSuite) SetupTest() {
	models.TruncateAll(ts.API.db)
}

func (ts *MFATestSuite) TestMFAEnable() {
	// Enable MFA for a single user
}

func (ts *MFATestSuite) TestMFADisable() {
	// Disable MFA for a single user
}

func (ts *MFATestSuite) TestMFAEnrollDevice() {
	// Make call to enroll
	// Device should be in the DB
	// the user id stored should match that of the user enrolled
	require.Equal(ts.T(), "hello", "hello", "enroll device not implemented")
}

func (ts *MFATestSuite) TestMFAUnenrollDevice() {
	// Create a user with enrolled device
	// Make call to unenroll endpoint
	// Device should be set to false in the db
	// Corresponding user should have no associtaed devices
	require.Equal(ts.T(), "hello", "hello", "Unenroll device not implemented")
}

func (ts *MFATestSuite) TestMFABackupCodeGeneration() {
	// Setup request and pass in user
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/mfa/generate_backup_codes", nil)
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ts.token))
	ts.API.handler.ServeHTTP(w, req)
	require.Equal(ts.T(), http.StatusOK, w.Code)
	// Make a call to the generate backup codes function
	// eight backup codes should be generated
	// The timestamp should be correct
	// All eight backup codes should be unique and unused
}

func (ts *MFATestSuite) TestMFAListDevices() {
	// Generate four devices
	// List four devices and expect that the corresponding primary keys match when calling
	// listDevices
	require.Equal(ts.T(), "hello", "hello", "Backup code generation not implemented")
}

func (ts *MFATestSuite) TestMFACreateChallenge() {
	// Make a call to create challenge, ensure that it is set in DB

	require.Equal(ts.T(), "hello", "hello", "Backup code generation not implemented")
}

func (ts *MFATestSuite) TestMFAVerifyChallenge() {
	// Verify against a challenge with fixed code using a mocked response
	// Verify with incorrect code should return false
	// Verify with correct code shoulde return a successful response
	require.Equal(ts.T(), "hello", "hello", "Backup code generation not implemented")
}

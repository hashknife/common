package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashknife/api/config"
	"github.com/stretchr/testify/suite"
)

// FrontendTestSuite
type FrontendTestSuite struct {
	suite.Suite
	conf     *config.Config
	endpoint Frontender
}

// SetupSuite runs code needed for the test suite
func (f *FrontendTestSuite) SetupSuite() {
	path := "/frontend"
	f.conf = &config.Config{
		FrontendPath: &path,
	}
	f.endpoint = NewFrontend(f.conf)
}

// TestFrontendTestSuite
func TestFrontendTestSuite(t *testing.T) {
	suite.Run(t, &FrontendTestSuite{})
}

// TestFrontend_Success
func (f *FrontendTestSuite) TestFrontend_Success() {
	server := httptest.NewServer(f.endpoint)
	defer server.Close()
	resp, err := http.Get(server.URL)
	f.Require().NoError(err)
	f.Require().Equal(resp.StatusCode, http.StatusOK)
}

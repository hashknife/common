package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/briandowns/hashknife/hashknife-api/config"
	"github.com/stretchr/testify/suite"
)

// Tile38TestSuite
type Tile38TestSuite struct {
	suite.Suite
	conf    *config.Config
	service Tile38er
}

// SetupSuite runs code needed for the test suite
func (p *Tile38TestSuite) SetupSuite() {
	p.conf = &config.Config{}
	p.endpoint = NewTile38(p.conf)
}

// TestFrontendTestSuite
func TestFTile38TestSuite(t *testing.T) {
	suite.Run(t, &Tile38TestSuite{})
}

// TestFrontend_Success
func (p *Tile38TestSuite) TestFrontend_Success() {
	server := httptest.NewServer(f.endpoint)
	defer server.Close()
	resp, err := http.Get(server.URL)
	f.Require().NoError(err)
	f.Require().Equal(resp.StatusCode, http.StatusOK)
}

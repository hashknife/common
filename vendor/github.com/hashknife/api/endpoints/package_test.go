package endpoints

import (
	"testing"

	"github.com/hashknife/api/config"
	"github.com/stretchr/testify/suite"
)

// PackageServiceTestSuite
type PackageServiceTestSuite struct {
	suite.Suite
	conf     *config.Config
	endpoint PackageServicer
}

// SetupSuite runs code needed for the test suite
func (p *PackageServiceTestSuite) SetupSuite() {
	p.conf = &config.Config{}
	p.endpoint = NewPackageService(p.conf)
}

// TestPackageServiceTestSuite
func TestPackageServiceTestSuite(t *testing.T) {
	suite.Run(t, &PackageServiceTestSuite{})
}

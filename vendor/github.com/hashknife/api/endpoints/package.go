package endpoints

import (
	"context"
	"errors"

	"github.com/hashknife/api/config"
)

// PackageServicer
type PackageServicer interface {
	Status(ctx context.Context, i interface{}) (interface{}, error)
	Deliver(ctx context.Context, i interface{}) (interface{}, error)
	Accept(ctx context.Context, i interface{}) (interface{}, error)
	Request(ctx context.Context, i interface{}) (interface{}, error)
}

// PackageService
type PackageService struct {
	conf *config.Config
}

// compile time validation
var _ PackageServicer = (*PackageService)(nil)

// NewPackageService
func NewPackageService(c *config.Config) PackageServicer {
	return &PackageService{c}
}

// PackageStatusRequest
type PackageStatusRequest struct {
	PackageID string `json:"package_id"`
}

// PackageStatusResponse
type PackageStatusResponse struct {
	PackageID string `json:"package_id"`
}

// PackageDeliverRequest
type PackageDeliverRequest struct {
	PackageID string `json:"package_id"`
}

// PackageDeliverResponse
type PackageDeliverResponse struct {
	PackageID string `json:"package_id"`
}

// PackageAcceptRequest
type PackageAcceptRequest struct {
	PackageID string `json:"package_id"`
}

// PackageAcceptResponse
type PackageAcceptResponse struct {
	PackageID string `json:"package_id"`
}

// PackageReqRequest
type PackageReqRequest struct {
	PackageID string `json:"package_id"`
}

// PackageReqResponse
type PackageReqResponse struct {
	PackageID string `json:"package_id"`
}

// Status
func (c *PackageService) Status(ctx context.Context, i interface{}) (interface{}, error) {
	req, ok := i.(*PackageStatusRequest)
	if !ok {
		return nil, errors.New("unable to convert request to PackageServiceRetrieveRequest type")
	}
	// look in the database
	return &PackageStatusResponse{req.PackageID}, nil
}

// Deliver
func (c *PackageService) Deliver(ctx context.Context, i interface{}) (interface{}, error) {
	req, ok := i.(*PackageDeliverRequest)
	if !ok {
		return nil, errors.New("unable to convert request to PackageServiceRetrieveRequest type")
	}
	// look in the database
	return &PackageDeliverResponse{req.PackageID}, nil
}

// Accept
func (c *PackageService) Accept(ctx context.Context, i interface{}) (interface{}, error) {
	req, ok := i.(*PackageAcceptRequest)
	if !ok {
		return nil, errors.New("unable to convert request to PackageServiceRetrieveRequest type")
	}
	// look in the database
	return &PackageAcceptResponse{req.PackageID}, nil
}

// Request
func (c *PackageService) Request(ctx context.Context, i interface{}) (interface{}, error) {
	req, ok := i.(*PackageReqRequest)
	if !ok {
		return nil, errors.New("unable to convert request to PackageServiceRetrieveRequest type")
	}
	// look in the database
	return &PackageReqResponse{req.PackageID}, nil
}

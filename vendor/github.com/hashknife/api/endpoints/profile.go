package endpoints

import (
	"context"
	"errors"

	"github.com/hashknife/api/config"
)

// ProfileServicer
type ProfileServicer interface {
	Retrieve(ctx context.Context, i interface{}) (interface{}, error)
	Update(ctx context.Context, i interface{}) (interface{}, error)
}

// ProfileService
type ProfileService struct {
	conf *config.Config
}

// compile time validation
var _ ProfileServicer = (*ProfileService)(nil)

// NewProfileService
func NewProfileService(c *config.Config) ProfileServicer {
	return &ProfileService{c}
}

// ProfileRetrieveRequest
type ProfileRetrieveRequest struct {
	ProfileID string `json:"Profile_id"`
}

// ProfileRetrieveResponse
type ProfileRetrieveResponse struct {
	ProfileID string `json:"Profile_id"`
}

// ProfileUpdateRequest
type ProfileUpdateRequest struct {
	ProfileID string `json:"Profile_id"`
}

// ProfileUpdateResponse
type ProfileUpdateResponse struct {
	ProfileID string `json:"Profile_id"`
}

// Retrieve
func (c *ProfileService) Retrieve(ctx context.Context, i interface{}) (interface{}, error) {
	req, ok := i.(*ProfileRetrieveRequest)
	if !ok {
		return nil, errors.New("unable to convert request to ProfileServiceRetrieveRequest type")
	}
	// look in the database
	return &ProfileRetrieveResponse{req.ProfileID}, nil
}

// Update
func (c *ProfileService) Update(ctx context.Context, i interface{}) (interface{}, error) {
	req, ok := i.(*ProfileUpdateRequest)
	if !ok {
		return nil, errors.New("unable to convert request to ProfileServiceRetrieveRequest type")
	}
	// look in the database
	return &ProfileUpdateResponse{req.ProfileID}, nil
}

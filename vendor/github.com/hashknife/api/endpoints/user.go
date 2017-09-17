package endpoints

import (
	"context"
	"errors"

	"github.com/hashknife/api/config"
)

// UserServicer
type UserServicer interface {
	Retrieve(ctx context.Context, i interface{}) (interface{}, error)
	Update(ctx context.Context, i interface{}) (interface{}, error)
	Disable(ctx context.Context, i interface{}) (interface{}, error)
}

// UserService
type UserService struct {
	conf *config.Config
}

// compile time validation
var _ UserServicer = (*UserService)(nil)

// NewUserService
func NewUserService(c *config.Config) UserServicer {
	return &UserService{c}
}

// UserServiceRetrieveRequest
type UserServiceRetrieveRequest struct {
	UserID string `json:"user_id"`
}

// UserServiceRetrieveResponse
type UserServiceRetrieveResponse struct {
	UserID string `json:"user_id"`
}

// UserServiceUpdateRequest
type UserServiceUpdateRequest struct {
	UserID string `json:"user_id"`
}

// UserServiceUpdateResponse
type UserServiceUpdateResponse struct {
	UserID string `json:"user_id"`
}

// UserServiceDisableRequest
type UserServiceDisableRequest struct {
	UserID string `json:"user_id"`
}

// UserServiceDisableResponse
type UserServiceDisableResponse struct {
	UserID string `json:"user_id"`
}

// Retrieve
func (c *UserService) Retrieve(ctx context.Context, i interface{}) (interface{}, error) {
	req, ok := i.(*UserServiceRetrieveRequest)
	if !ok {
		return nil, errors.New("unable to convert request to UserServiceRetrieveRequest type")
	}
	// look in the database
	return &UserServiceRetrieveResponse{req.UserID}, nil
}

// Update
func (c *UserService) Update(ctx context.Context, i interface{}) (interface{}, error) {
	req, ok := i.(*UserServiceUpdateRequest)
	if !ok {
		return nil, errors.New("unable to convert request to UserServiceUpdateRequest type")
	}
	// look in the database
	return &UserServiceUpdateResponse{req.UserID}, nil
}

// Disable
func (c *UserService) Disable(ctx context.Context, i interface{}) (interface{}, error) {
	req, ok := i.(*UserServiceDisableRequest)
	if !ok {
		return nil, errors.New("unable to convert request to UserServiceDisableRequest type")
	}
	// look in the database
	return &UserServiceDisableResponse{req.UserID}, nil
}

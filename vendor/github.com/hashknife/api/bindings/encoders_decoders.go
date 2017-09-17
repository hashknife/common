package bindings

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashknife/api/endpoints"
)

// encodeHealthCheckHTTPResponse encodes the response given by the healthCheck Endpoint
func encodeHealthCheckHTTPResponse(ctx context.Context, w http.ResponseWriter, i interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(i.(*endpoints.HealthCheckResponse))
}

// encodeResponse is used by the application routes to return their data
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	addPoweredHeaders(w)
	return json.NewEncoder(w).Encode(response)
}

// noOpDecodeRequest is of type GoKit http.DecodeRequestFunc and is used in place of writing out
// an empty function (func(context.Context, *http.Request) (interface{}, error) { return struct{}{}, nil },)
// directly in the NewServer call.
func noOpDecodeRequest(context.Context, *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

// decodePackageStatusHTTPRequest
func decodePackageStatusHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var psr endpoints.PackageStatusRequest
	params := mux.Vars(r)
	if params["package_id"] == "" {
		return nil, errors.New("no package id")
	}
	psr.PackageID = params["package_id"]
	return &endpoints.PackageStatusRequest{PackageID: psr.PackageID}, nil
}

// decodePackageDeliverHTTPRequest
func decodePackageDeliverHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var psr endpoints.PackageStatusRequest
	params := mux.Vars(r)
	if params["package_id"] == "" {
		return nil, errors.New("no package id")
	}
	psr.PackageID = params["package_id"]
	return &endpoints.PackageStatusRequest{PackageID: psr.PackageID}, nil
}

// decodePackageAcceptHTTPRequest
func decodePackageAcceptHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var psr endpoints.PackageStatusRequest
	params := mux.Vars(r)
	if params["package_id"] == "" {
		return nil, errors.New("no package id")
	}
	psr.PackageID = params["package_id"]
	return &endpoints.PackageStatusRequest{PackageID: psr.PackageID}, nil
}

// decodeRetrieveUserHTTPRequest
func decodeRetrieveUserHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var psr endpoints.PackageStatusRequest
	params := mux.Vars(r)
	if params["package_id"] == "" {
		return nil, errors.New("no package id")
	}
	psr.PackageID = params["package_id"]
	return &endpoints.PackageStatusRequest{PackageID: psr.PackageID}, nil
}

// decodeUpdateUserHTTPRequest
func decodeUpdateUserHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var psr endpoints.PackageStatusRequest
	params := mux.Vars(r)
	if params["package_id"] == "" {
		return nil, errors.New("no package id")
	}
	psr.PackageID = params["package_id"]
	return &endpoints.PackageStatusRequest{PackageID: psr.PackageID}, nil
}

// decodeCreateDisableHTTPRequest
func decodeCreateDisableHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var psr endpoints.PackageStatusRequest
	params := mux.Vars(r)
	if params["package_id"] == "" {
		return nil, errors.New("no package id")
	}
	psr.PackageID = params["package_id"]
	return &endpoints.PackageStatusRequest{PackageID: psr.PackageID}, nil
}

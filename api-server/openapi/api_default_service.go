/*
 * MassBank Tool API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// SpectraAccessionGet -
func (s *DefaultApiService) SpectraAccessionGet(ctx context.Context, accession string) (ImplResponse, error) {
	// TODO - update SpectraAccessionGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Spectrum{}) or use other options such as http.Ok ...
	//return Response(200, Spectrum{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("SpectraAccessionGet method not implemented")
}

// SpectraGet -
func (s *DefaultApiService) SpectraGet(ctx context.Context, limit int64, offset int64, page int64) (ImplResponse, error) {
	// TODO - update SpectraGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []SpectrumListItem{}) or use other options such as http.Ok ...
	//return Response(200, []SpectrumListItem{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SpectraGet method not implemented")
}

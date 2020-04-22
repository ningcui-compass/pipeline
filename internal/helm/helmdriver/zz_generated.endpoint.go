// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package helmdriver

import (
	"context"
	"errors"
	"github.com/banzaicloud/pipeline/internal/helm"
	"github.com/go-kit/kit/endpoint"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
)

// endpointError identifies an error that should be returned as an endpoint error.
type endpointError interface {
	EndpointError() bool
}

// serviceError identifies an error that should be returned as a service error.
type serviceError interface {
	ServiceError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	AddRepository       endpoint.Endpoint
	CheckRelease        endpoint.Endpoint
	DeleteRelease       endpoint.Endpoint
	DeleteRepository    endpoint.Endpoint
	GetChart            endpoint.Endpoint
	GetRelease          endpoint.Endpoint
	GetReleaseResources endpoint.Endpoint
	InstallRelease      endpoint.Endpoint
	ListCharts          endpoint.Endpoint
	ListReleases        endpoint.Endpoint
	ListRepositories    endpoint.Endpoint
	ModifyRepository    endpoint.Endpoint
	UpdateRepository    endpoint.Endpoint
	UpgradeRelease      endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service helm.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{
		AddRepository:       kitxendpoint.OperationNameMiddleware("helm.AddRepository")(mw(MakeAddRepositoryEndpoint(service))),
		CheckRelease:        kitxendpoint.OperationNameMiddleware("helm.CheckRelease")(mw(MakeCheckReleaseEndpoint(service))),
		DeleteRelease:       kitxendpoint.OperationNameMiddleware("helm.DeleteRelease")(mw(MakeDeleteReleaseEndpoint(service))),
		DeleteRepository:    kitxendpoint.OperationNameMiddleware("helm.DeleteRepository")(mw(MakeDeleteRepositoryEndpoint(service))),
		GetChart:            kitxendpoint.OperationNameMiddleware("helm.GetChart")(mw(MakeGetChartEndpoint(service))),
		GetRelease:          kitxendpoint.OperationNameMiddleware("helm.GetRelease")(mw(MakeGetReleaseEndpoint(service))),
		GetReleaseResources: kitxendpoint.OperationNameMiddleware("helm.GetReleaseResources")(mw(MakeGetReleaseResourcesEndpoint(service))),
		InstallRelease:      kitxendpoint.OperationNameMiddleware("helm.InstallRelease")(mw(MakeInstallReleaseEndpoint(service))),
		ListCharts:          kitxendpoint.OperationNameMiddleware("helm.ListCharts")(mw(MakeListChartsEndpoint(service))),
		ListReleases:        kitxendpoint.OperationNameMiddleware("helm.ListReleases")(mw(MakeListReleasesEndpoint(service))),
		ListRepositories:    kitxendpoint.OperationNameMiddleware("helm.ListRepositories")(mw(MakeListRepositoriesEndpoint(service))),
		ModifyRepository:    kitxendpoint.OperationNameMiddleware("helm.ModifyRepository")(mw(MakeModifyRepositoryEndpoint(service))),
		UpdateRepository:    kitxendpoint.OperationNameMiddleware("helm.UpdateRepository")(mw(MakeUpdateRepositoryEndpoint(service))),
		UpgradeRelease:      kitxendpoint.OperationNameMiddleware("helm.UpgradeRelease")(mw(MakeUpgradeReleaseEndpoint(service))),
	}
}

// AddRepositoryRequest is a request struct for AddRepository endpoint.
type AddRepositoryRequest struct {
	OrganizationID uint
	Repository     helm.Repository
}

// AddRepositoryResponse is a response struct for AddRepository endpoint.
type AddRepositoryResponse struct {
	Err error
}

func (r AddRepositoryResponse) Failed() error {
	return r.Err
}

// MakeAddRepositoryEndpoint returns an endpoint for the matching method of the underlying service.
func MakeAddRepositoryEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRepositoryRequest)

		err := service.AddRepository(ctx, req.OrganizationID, req.Repository)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return AddRepositoryResponse{Err: err}, nil
			}

			return AddRepositoryResponse{Err: err}, err
		}

		return AddRepositoryResponse{}, nil
	}
}

// CheckReleaseRequest is a request struct for CheckRelease endpoint.
type CheckReleaseRequest struct {
	OrganizationID uint
	ClusterID      uint
	ReleaseName    string
	Options        helm.Options
}

// CheckReleaseResponse is a response struct for CheckRelease endpoint.
type CheckReleaseResponse struct {
	R0  string
	Err error
}

func (r CheckReleaseResponse) Failed() error {
	return r.Err
}

// MakeCheckReleaseEndpoint returns an endpoint for the matching method of the underlying service.
func MakeCheckReleaseEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CheckReleaseRequest)

		r0, err := service.CheckRelease(ctx, req.OrganizationID, req.ClusterID, req.ReleaseName, req.Options)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return CheckReleaseResponse{
					Err: err,
					R0:  r0,
				}, nil
			}

			return CheckReleaseResponse{
				Err: err,
				R0:  r0,
			}, err
		}

		return CheckReleaseResponse{R0: r0}, nil
	}
}

// DeleteReleaseRequest is a request struct for DeleteRelease endpoint.
type DeleteReleaseRequest struct {
	OrganizationID uint
	ClusterID      uint
	ReleaseName    string
	Options        helm.Options
}

// DeleteReleaseResponse is a response struct for DeleteRelease endpoint.
type DeleteReleaseResponse struct {
	Err error
}

func (r DeleteReleaseResponse) Failed() error {
	return r.Err
}

// MakeDeleteReleaseEndpoint returns an endpoint for the matching method of the underlying service.
func MakeDeleteReleaseEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteReleaseRequest)

		err := service.DeleteRelease(ctx, req.OrganizationID, req.ClusterID, req.ReleaseName, req.Options)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return DeleteReleaseResponse{Err: err}, nil
			}

			return DeleteReleaseResponse{Err: err}, err
		}

		return DeleteReleaseResponse{}, nil
	}
}

// DeleteRepositoryRequest is a request struct for DeleteRepository endpoint.
type DeleteRepositoryRequest struct {
	OrganizationID uint
	RepoName       string
}

// DeleteRepositoryResponse is a response struct for DeleteRepository endpoint.
type DeleteRepositoryResponse struct {
	Err error
}

func (r DeleteRepositoryResponse) Failed() error {
	return r.Err
}

// MakeDeleteRepositoryEndpoint returns an endpoint for the matching method of the underlying service.
func MakeDeleteRepositoryEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRepositoryRequest)

		err := service.DeleteRepository(ctx, req.OrganizationID, req.RepoName)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return DeleteRepositoryResponse{Err: err}, nil
			}

			return DeleteRepositoryResponse{Err: err}, err
		}

		return DeleteRepositoryResponse{}, nil
	}
}

// GetChartRequest is a request struct for GetChart endpoint.
type GetChartRequest struct {
	OrganizationID uint
	ChartFilter    helm.ChartFilter
	Options        helm.Options
}

// GetChartResponse is a response struct for GetChart endpoint.
type GetChartResponse struct {
	ChartDetails map[string]interface{}
	Err          error
}

func (r GetChartResponse) Failed() error {
	return r.Err
}

// MakeGetChartEndpoint returns an endpoint for the matching method of the underlying service.
func MakeGetChartEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetChartRequest)

		chartDetails, err := service.GetChart(ctx, req.OrganizationID, req.ChartFilter, req.Options)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return GetChartResponse{
					ChartDetails: chartDetails,
					Err:          err,
				}, nil
			}

			return GetChartResponse{
				ChartDetails: chartDetails,
				Err:          err,
			}, err
		}

		return GetChartResponse{ChartDetails: chartDetails}, nil
	}
}

// GetReleaseRequest is a request struct for GetRelease endpoint.
type GetReleaseRequest struct {
	OrganizationID uint
	ClusterID      uint
	ReleaseName    string
	Options        helm.Options
}

// GetReleaseResponse is a response struct for GetRelease endpoint.
type GetReleaseResponse struct {
	R0  helm.Release
	Err error
}

func (r GetReleaseResponse) Failed() error {
	return r.Err
}

// MakeGetReleaseEndpoint returns an endpoint for the matching method of the underlying service.
func MakeGetReleaseEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetReleaseRequest)

		r0, err := service.GetRelease(ctx, req.OrganizationID, req.ClusterID, req.ReleaseName, req.Options)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return GetReleaseResponse{
					Err: err,
					R0:  r0,
				}, nil
			}

			return GetReleaseResponse{
				Err: err,
				R0:  r0,
			}, err
		}

		return GetReleaseResponse{R0: r0}, nil
	}
}

// GetReleaseResourcesRequest is a request struct for GetReleaseResources endpoint.
type GetReleaseResourcesRequest struct {
	OrganizationID uint
	ClusterID      uint
	Release        helm.Release
	Options        helm.Options
}

// GetReleaseResourcesResponse is a response struct for GetReleaseResources endpoint.
type GetReleaseResourcesResponse struct {
	R0  []helm.ReleaseResource
	Err error
}

func (r GetReleaseResourcesResponse) Failed() error {
	return r.Err
}

// MakeGetReleaseResourcesEndpoint returns an endpoint for the matching method of the underlying service.
func MakeGetReleaseResourcesEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetReleaseResourcesRequest)

		r0, err := service.GetReleaseResources(ctx, req.OrganizationID, req.ClusterID, req.Release, req.Options)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return GetReleaseResourcesResponse{
					Err: err,
					R0:  r0,
				}, nil
			}

			return GetReleaseResourcesResponse{
				Err: err,
				R0:  r0,
			}, err
		}

		return GetReleaseResourcesResponse{R0: r0}, nil
	}
}

// InstallReleaseRequest is a request struct for InstallRelease endpoint.
type InstallReleaseRequest struct {
	OrganizationID uint
	ClusterID      uint
	Release        helm.Release
	Options        helm.Options
}

// InstallReleaseResponse is a response struct for InstallRelease endpoint.
type InstallReleaseResponse struct {
	Err error
}

func (r InstallReleaseResponse) Failed() error {
	return r.Err
}

// MakeInstallReleaseEndpoint returns an endpoint for the matching method of the underlying service.
func MakeInstallReleaseEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InstallReleaseRequest)

		err := service.InstallRelease(ctx, req.OrganizationID, req.ClusterID, req.Release, req.Options)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return InstallReleaseResponse{Err: err}, nil
			}

			return InstallReleaseResponse{Err: err}, err
		}

		return InstallReleaseResponse{}, nil
	}
}

// ListChartsRequest is a request struct for ListCharts endpoint.
type ListChartsRequest struct {
	OrganizationID uint
	Filter         helm.ChartFilter
	Options        helm.Options
}

// ListChartsResponse is a response struct for ListCharts endpoint.
type ListChartsResponse struct {
	Charts []interface{}
	Err    error
}

func (r ListChartsResponse) Failed() error {
	return r.Err
}

// MakeListChartsEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListChartsEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListChartsRequest)

		charts, err := service.ListCharts(ctx, req.OrganizationID, req.Filter, req.Options)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListChartsResponse{
					Charts: charts,
					Err:    err,
				}, nil
			}

			return ListChartsResponse{
				Charts: charts,
				Err:    err,
			}, err
		}

		return ListChartsResponse{Charts: charts}, nil
	}
}

// ListReleasesRequest is a request struct for ListReleases endpoint.
type ListReleasesRequest struct {
	OrganizationID uint
	ClusterID      uint
	Filters        interface{}
	Options        helm.Options
}

// ListReleasesResponse is a response struct for ListReleases endpoint.
type ListReleasesResponse struct {
	R0  []helm.Release
	Err error
}

func (r ListReleasesResponse) Failed() error {
	return r.Err
}

// MakeListReleasesEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListReleasesEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListReleasesRequest)

		r0, err := service.ListReleases(ctx, req.OrganizationID, req.ClusterID, req.Filters, req.Options)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListReleasesResponse{
					Err: err,
					R0:  r0,
				}, nil
			}

			return ListReleasesResponse{
				Err: err,
				R0:  r0,
			}, err
		}

		return ListReleasesResponse{R0: r0}, nil
	}
}

// ListRepositoriesRequest is a request struct for ListRepositories endpoint.
type ListRepositoriesRequest struct {
	OrganizationID uint
}

// ListRepositoriesResponse is a response struct for ListRepositories endpoint.
type ListRepositoriesResponse struct {
	Repos []helm.Repository
	Err   error
}

func (r ListRepositoriesResponse) Failed() error {
	return r.Err
}

// MakeListRepositoriesEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListRepositoriesEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRepositoriesRequest)

		repos, err := service.ListRepositories(ctx, req.OrganizationID)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListRepositoriesResponse{
					Err:   err,
					Repos: repos,
				}, nil
			}

			return ListRepositoriesResponse{
				Err:   err,
				Repos: repos,
			}, err
		}

		return ListRepositoriesResponse{Repos: repos}, nil
	}
}

// ModifyRepositoryRequest is a request struct for ModifyRepository endpoint.
type ModifyRepositoryRequest struct {
	OrganizationID uint
	Repository     helm.Repository
}

// ModifyRepositoryResponse is a response struct for ModifyRepository endpoint.
type ModifyRepositoryResponse struct {
	Err error
}

func (r ModifyRepositoryResponse) Failed() error {
	return r.Err
}

// MakeModifyRepositoryEndpoint returns an endpoint for the matching method of the underlying service.
func MakeModifyRepositoryEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ModifyRepositoryRequest)

		err := service.ModifyRepository(ctx, req.OrganizationID, req.Repository)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ModifyRepositoryResponse{Err: err}, nil
			}

			return ModifyRepositoryResponse{Err: err}, err
		}

		return ModifyRepositoryResponse{}, nil
	}
}

// UpdateRepositoryRequest is a request struct for UpdateRepository endpoint.
type UpdateRepositoryRequest struct {
	OrganizationID uint
	Repository     helm.Repository
}

// UpdateRepositoryResponse is a response struct for UpdateRepository endpoint.
type UpdateRepositoryResponse struct {
	Err error
}

func (r UpdateRepositoryResponse) Failed() error {
	return r.Err
}

// MakeUpdateRepositoryEndpoint returns an endpoint for the matching method of the underlying service.
func MakeUpdateRepositoryEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRepositoryRequest)

		err := service.UpdateRepository(ctx, req.OrganizationID, req.Repository)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return UpdateRepositoryResponse{Err: err}, nil
			}

			return UpdateRepositoryResponse{Err: err}, err
		}

		return UpdateRepositoryResponse{}, nil
	}
}

// UpgradeReleaseRequest is a request struct for UpgradeRelease endpoint.
type UpgradeReleaseRequest struct {
	OrganizationID uint
	ClusterID      uint
	Release        helm.Release
	Options        helm.Options
}

// UpgradeReleaseResponse is a response struct for UpgradeRelease endpoint.
type UpgradeReleaseResponse struct {
	Err error
}

func (r UpgradeReleaseResponse) Failed() error {
	return r.Err
}

// MakeUpgradeReleaseEndpoint returns an endpoint for the matching method of the underlying service.
func MakeUpgradeReleaseEndpoint(service helm.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpgradeReleaseRequest)

		err := service.UpgradeRelease(ctx, req.OrganizationID, req.ClusterID, req.Release, req.Options)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return UpgradeReleaseResponse{Err: err}, nil
			}

			return UpgradeReleaseResponse{Err: err}, err
		}

		return UpgradeReleaseResponse{}, nil
	}
}

/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.3.0
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

type ListDeploymentsResponseInner struct {
	ReleaseName  string `json:"releaseName,omitempty"`
	Chart        string `json:"chart,omitempty"`
	ChartName    string `json:"chartName,omitempty"`
	ChartVersion string `json:"chartVersion,omitempty"`
	Version      int32  `json:"version,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
	Status       string `json:"status,omitempty"`
	Namespace    string `json:"namespace,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	Supported    bool   `json:"supported,omitempty"`
}

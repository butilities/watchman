/*
 * Watchman API
 *
 * Moov Watchman is an HTTP API and Go library to download, parse and offer search functions over numerous trade sanction lists from the United States, European Union governments, agencies, and non profits for complying with regional laws. Also included is a web UI and async webhook notification service to initiate processes on remote systems.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// BisEntities Bureau of Industry and Security Entity List
type BisEntities struct {
	// The name of the entity
	Name string `json:"name,omitempty"`
	// Addresses associated with the entity
	Addresses []string `json:"addresses,omitempty"`
	// Known aliases associated with the entity
	AlternateNames []string `json:"alternateNames,omitempty"`
	// Date when the restriction came into effect
	StartDate string `json:"startDate,omitempty"`
	// Specifies the license requirement imposed on the named entity
	LicenseRequirement string `json:"licenseRequirement,omitempty"`
	// Identifies the policy BIS uses to review the licenseRequirements
	LicensePolicy string `json:"licensePolicy,omitempty"`
	// Identifies the corresponding Notice in the Federal Register
	FrNotice string `json:"frNotice,omitempty"`
	// The link to the official SSI list
	SourceListURL string `json:"sourceListURL,omitempty"`
	// The link for information regarding the source
	SourceInfoURL string `json:"sourceInfoURL,omitempty"`
}

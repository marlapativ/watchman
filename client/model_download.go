/*
 * Watchman API
 *
 * Moov Watchman is an HTTP API and Go library to download, parse and offer search functions over numerous trade sanction lists from the United States, European Union governments, agencies, and non profits for complying with regional laws. Also included is a web UI and async webhook notification service to initiate processes on remote systems.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// Download Metadata and stats about downloaded OFAC data
type Download struct {
	SDNs      int32     `json:"SDNs,omitempty"`
	AltNames  int32     `json:"altNames,omitempty"`
	Addresses int32     `json:"addresses,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

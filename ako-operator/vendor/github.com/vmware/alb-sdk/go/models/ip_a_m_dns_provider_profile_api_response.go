// Copyright © 2025 Broadcom Inc. and/or its subsidiaries. All Rights Reserved.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// IPAMDNSProviderProfileAPIResponse ipam Dns provider profile Api response
// swagger:model IpamDnsProviderProfileApiResponse
type IPAMDNSProviderProfileAPIResponse struct {

	// count
	// Required: true
	Count *int32 `json:"count"`

	// next
	Next *string `json:"next,omitempty"`

	// results
	// Required: true
	Results []*IPAMDNSProviderProfile `json:"results,omitempty"`
}

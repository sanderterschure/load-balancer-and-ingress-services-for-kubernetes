// Copyright © 2025 Broadcom Inc. and/or its subsidiaries. All Rights Reserved.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// HostHdrMatch host hdr match
// swagger:model HostHdrMatch
type HostHdrMatch struct {

	// Case sensitivity to use for the match. Enum options - SENSITIVE, INSENSITIVE. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	MatchCase *string `json:"match_case,omitempty"`

	// Criterion to use for the host header value match. Enum options - HDR_EXISTS, HDR_DOES_NOT_EXIST, HDR_BEGINS_WITH, HDR_DOES_NOT_BEGIN_WITH, HDR_CONTAINS, HDR_DOES_NOT_CONTAIN, HDR_ENDS_WITH, HDR_DOES_NOT_END_WITH, HDR_EQUALS, HDR_DOES_NOT_EQUAL. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	// Required: true
	MatchCriteria *string `json:"match_criteria"`

	// String value(s) in the host header. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	Value []string `json:"value,omitempty"`
}

// Copyright © 2025 Broadcom Inc. and/or its subsidiaries. All Rights Reserved.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// OCSPConfig o c s p config
// swagger:model OCSPConfig
type OCSPConfig struct {

	// Describes the Time Interval after which the next OCSP job needs to be scheduled in case of the OCSP job failures. Allowed values are 60-86400. Field introduced in 20.1.1. Unit is SEC. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	FailedOcspJobsRetryInterval *uint32 `json:"failed_ocsp_jobs_retry_interval,omitempty"`

	// Maximum number of times the failed OCSP jobs can be scheduled. Field introduced in 20.1.1. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	MaxTries *uint32 `json:"max_tries,omitempty"`

	// Interval between the OCSP queries. Allowed values are 60-31536000. Field introduced in 20.1.1. Unit is SEC. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	OcspReqInterval *uint32 `json:"ocsp_req_interval,omitempty"`

	// Time in seconds that the system waits for a reply from the OCSP responder before dropping the connection. Field introduced in 20.1.1. Unit is SEC. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	OcspRespTimeout *uint32 `json:"ocsp_resp_timeout,omitempty"`

	// List of Responder URLs configured by user to do failover/override the AIA extension contained in the OCSP responder's SSL/TLS certificate. Field introduced in 20.1.1. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	ResponderURLLists []string `json:"responder_url_lists,omitempty"`

	// Describes the type of action to take with the Responder URLs. Enum options - OCSP_RESPONDER_URL_FAILOVER, OCSP_RESPONDER_URL_OVERRIDE. Field introduced in 20.1.1. Allowed in Enterprise edition with any value, Essentials, Basic, Enterprise with Cloud Services edition.
	URLAction *string `json:"url_action,omitempty"`
}

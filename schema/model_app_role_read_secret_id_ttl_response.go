// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleReadSecretIdTtlResponse struct for AppRoleReadSecretIdTtlResponse
type AppRoleReadSecretIdTtlResponse struct {
	// Duration in seconds after which the issued secret ID expires.
	SecretIdTtl int32 `json:"secret_id_ttl,omitempty"`
}

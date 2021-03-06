/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package upgrade

// An error id and message pair
type ComputeInstanceErrorMessage struct {

	// Error message string to indicate, if it is NSX or cloud operation generated error.
	DetailedMessage string `json:"detailed_message,omitempty"`

	// an error id contract obtained from PCM
	ErrorId int64 `json:"error_id,omitempty"`
}

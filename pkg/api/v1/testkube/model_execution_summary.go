/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

import (
	"time"
)

// execution summary
type ExecutionSummary struct {
	// execution id
	Id string `json:"id"`
	// execution name
	Name string `json:"name"`
	// name of the script
	ScriptName string `json:"testName"`
	// the type of script for this execution
	ScriptType string           `json:"scriptType"`
	Status     *ExecutionStatus `json:"status"`
	// test execution start time
	StartTime time.Time `json:"startTime,omitempty"`
	// test execution end time
	EndTime time.Time `json:"endTime,omitempty"`
}

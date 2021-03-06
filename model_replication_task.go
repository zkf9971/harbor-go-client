/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The replication task
type ReplicationTask struct {
	// The ID
	Id int32 `json:"id,omitempty"`
	// The execution ID
	ExecutionId int32 `json:"execution_id,omitempty"`
	// The resource type
	ResourceType string `json:"resource_type,omitempty"`
	// The source resource
	SrcResource string `json:"src_resource,omitempty"`
	// The destination resource
	DstResource string `json:"dst_resource,omitempty"`
	// The job ID
	JobId string `json:"job_id,omitempty"`
	// The status
	Status string `json:"status,omitempty"`
	// The start time
	StartTime string `json:"start_time,omitempty"`
	// The end time
	EndTime string `json:"end_time,omitempty"`
}

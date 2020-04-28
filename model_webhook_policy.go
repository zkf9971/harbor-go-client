/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The webhook policy object
type WebhookPolicy struct {
	// The webhook policy ID.
	Id int64 `json:"id,omitempty"`
	// The name of webhook policy.
	Name string `json:"name,omitempty"`
	// The description of webhook policy.
	Description string `json:"description,omitempty"`
	// The project ID of webhook policy.
	ProjectId int32 `json:"project_id,omitempty"`
	Targets []WebhookTargetObject `json:"targets,omitempty"`
	EventTypes []string `json:"event_types,omitempty"`
	// The creator of the webhook policy.
	Creator string `json:"creator,omitempty"`
	// The create time of the webhook policy.
	CreationTime string `json:"creation_time,omitempty"`
	// The update time of the webhook policy.
	UpdateTime string `json:"update_time,omitempty"`
	// Whether the webhook policy is enabled or not.
	Enabled bool `json:"enabled,omitempty"`
}

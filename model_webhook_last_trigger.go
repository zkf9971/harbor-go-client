/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The webhook policy and last trigger time group by event type.
type WebhookLastTrigger struct {
	// The webhook event type.
	EventType string `json:"event_type,omitempty"`
	// Whether or not the webhook policy enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The creation time of webhook policy.
	CreationTime string `json:"creation_time,omitempty"`
	// The last trigger time of webhook policy.
	LastTriggerTime string `json:"last_trigger_time,omitempty"`
}

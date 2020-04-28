/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Label struct {
	// The ID of label.
	Id int32 `json:"id,omitempty"`
	// The name of label.
	Name string `json:"name,omitempty"`
	// The description of label.
	Description string `json:"description,omitempty"`
	// The color of label.
	Color string `json:"color,omitempty"`
	// The scope of label, g for global labels and p for project labels.
	Scope string `json:"scope,omitempty"`
	// The project ID if the label is a project label.
	ProjectId int32 `json:"project_id,omitempty"`
	// The creation time of label.
	CreationTime string `json:"creation_time,omitempty"`
	// The update time of label.
	UpdateTime string `json:"update_time,omitempty"`
	// The label is deleted or not.
	Deleted bool `json:"deleted,omitempty"`
}

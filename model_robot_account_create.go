/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RobotAccountCreate struct {
	// The name of robot account
	Name string `json:"name,omitempty"`
	// The description of robot account
	Description string `json:"description,omitempty"`
	// The permission of robot account
	Access []RobotAccountAccess `json:"access,omitempty"`
}
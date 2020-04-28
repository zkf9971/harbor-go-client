/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UserEntity struct {
	// The ID of the user.
	UserId int32 `json:"user_id,omitempty"`
	// The name of the user.
	Username string `json:"username,omitempty"`
}

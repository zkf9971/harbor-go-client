/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RetentionSelector struct {
	Kind string `json:"kind,omitempty"`
	Decoration string `json:"decoration,omitempty"`
	Pattern string `json:"pattern,omitempty"`
}

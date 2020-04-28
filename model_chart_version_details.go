/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The detailed information of the chart entry
type ChartVersionDetails struct {
	Metadata *ChartVersion `json:"metadata,omitempty"`
	Security *SecurityReport `json:"security,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
	Values map[string]interface{} `json:"values,omitempty"`
	Files map[string]string `json:"files,omitempty"`
	Labels *Labels `json:"labels,omitempty"`
}

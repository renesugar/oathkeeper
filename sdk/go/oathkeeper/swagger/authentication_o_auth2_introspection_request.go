/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type AuthenticationOAuth2IntrospectionRequest struct {

	// Scopes is an array of scopes that are required.
	Scope []string `json:"scope,omitempty"`

	// Token is the token to introspect.
	Token string `json:"token,omitempty"`
}

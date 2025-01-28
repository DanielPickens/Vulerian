/*
 * vulerian dev
 *
 * API interface for 'vulerian dev'
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DevstateResourceResourceNamePatchRequest struct {
	Inlined string `json:"inlined,omitempty"`

	Uri string `json:"uri,omitempty"`

	DeployByDefault string `json:"deployByDefault,omitempty"`
}

// AssertDevstateResourceResourceNamePatchRequestRequired checks if the required fields are not zero-ed
func AssertDevstateResourceResourceNamePatchRequestRequired(obj DevstateResourceResourceNamePatchRequest) error {
	return nil
}

// AssertRecurseDevstateResourceResourceNamePatchRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DevstateResourceResourceNamePatchRequest (e.g. [][]DevstateResourceResourceNamePatchRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDevstateResourceResourceNamePatchRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDevstateResourceResourceNamePatchRequest, ok := obj.(DevstateResourceResourceNamePatchRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDevstateResourceResourceNamePatchRequestRequired(aDevstateResourceResourceNamePatchRequest)
	})
}

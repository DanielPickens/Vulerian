/*
 * vulerian dev
 *
 * API interface for 'vulerian dev'
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DevstateCommandCommandNameMovePostRequest struct {

	// Initial group of the command
	FromGroup string `json:"fromGroup,omitempty"`

	// Initial index of the command in the group
	FromIndex int32 `json:"fromIndex,omitempty"`

	// Target group of the command
	ToGroup string `json:"toGroup,omitempty"`

	// Target index of the command in the group
	ToIndex int32 `json:"toIndex,omitempty"`
}

// AssertDevstateCommandCommandNameMovePostRequestRequired checks if the required fields are not zero-ed
func AssertDevstateCommandCommandNameMovePostRequestRequired(obj DevstateCommandCommandNameMovePostRequest) error {
	return nil
}

// AssertRecurseDevstateCommandCommandNameMovePostRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DevstateCommandCommandNameMovePostRequest (e.g. [][]DevstateCommandCommandNameMovePostRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDevstateCommandCommandNameMovePostRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDevstateCommandCommandNameMovePostRequest, ok := obj.(DevstateCommandCommandNameMovePostRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDevstateCommandCommandNameMovePostRequestRequired(aDevstateCommandCommandNameMovePostRequest)
	})
}

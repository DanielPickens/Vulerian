/*
 * vulerian dev
 *
 * API interface for 'vulerian dev'
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DevstateApplyCommandPostRequest struct {

	// Name of the command
	Name string `json:"name,omitempty"`

	Component string `json:"component,omitempty"`
}

// AssertDevstateApplyCommandPostRequestRequired checks if the required fields are not zero-ed
func AssertDevstateApplyCommandPostRequestRequired(obj DevstateApplyCommandPostRequest) error {
	return nil
}

// AssertRecurseDevstateApplyCommandPostRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DevstateApplyCommandPostRequest (e.g. [][]DevstateApplyCommandPostRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDevstateApplyCommandPostRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDevstateApplyCommandPostRequest, ok := obj.(DevstateApplyCommandPostRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDevstateApplyCommandPostRequestRequired(aDevstateApplyCommandPostRequest)
	})
}

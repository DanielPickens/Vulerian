/*
 * vulerian dev
 *
 * API interface for 'vulerian dev'
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Events struct {
	PreStart []string `json:"preStart,omitempty"`

	PostStart []string `json:"postStart,omitempty"`

	PreStop []string `json:"preStop,omitempty"`

	PostStop []string `json:"postStop,omitempty"`
}

// AssertEventsRequired checks if the required fields are not zero-ed
func AssertEventsRequired(obj Events) error {
	return nil
}

// AssertRecurseEventsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Events (e.g. [][]Events), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEventsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEvents, ok := obj.(Events)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEventsRequired(aEvents)
	})
}

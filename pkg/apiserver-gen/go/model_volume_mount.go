/*
 * vulerian dev
 *
 * API interface for 'vulerian dev'
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VolumeMount struct {
	Name string `json:"name"`

	Path string `json:"path"`
}

// AssertVolumeMountRequired checks if the required fields are not zero-ed
func AssertVolumeMountRequired(obj VolumeMount) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"path": obj.Path,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseVolumeMountRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of VolumeMount (e.g. [][]VolumeMount), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseVolumeMountRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aVolumeMount, ok := obj.(VolumeMount)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertVolumeMountRequired(aVolumeMount)
	})
}

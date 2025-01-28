/*
 * vulerian dev
 *
 * API interface for 'vulerian dev'
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DevstateContainerPostRequest struct {

	// Name of the container
	Name string `json:"name"`

	// Container image
	Image string `json:"image"`

	// Entrypoint of the container
	Command []string `json:"command,omitempty"`

	// Args passed to the Container entrypoint
	Args []string `json:"args,omitempty"`

	// Environment variables to define
	Env []Env `json:"env,omitempty"`

	// Requested memory for the deployed container
	MemReq string `json:"memReq,omitempty"`

	// Memory limit for the deployed container
	MemLimit string `json:"memLimit,omitempty"`

	// Requested CPU for the deployed container
	CpuReq string `json:"cpuReq,omitempty"`

	// CPU limit for the deployed container
	CpuLimit string `json:"cpuLimit,omitempty"`

	// Volume to mount into the container filesystem
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`

	// If false, mountSources and sourceMapping values are not considered
	ConfigureSources bool `json:"configureSources,omitempty"`

	// If true, sources are mounted into container's filesystem
	MountSources bool `json:"mountSources,omitempty"`

	// Specific directory on which to mount sources
	SourceMapping string `json:"sourceMapping,omitempty"`

	Annotation Annotation `json:"annotation,omitempty"`

	// Endpoints exposed by the container
	Endpoints []Endpoint `json:"endpoints,omitempty"`
}

// AssertDevstateContainerPostRequestRequired checks if the required fields are not zero-ed
func AssertDevstateContainerPostRequestRequired(obj DevstateContainerPostRequest) error {
	elements := map[string]interface{}{
		"name":  obj.Name,
		"image": obj.Image,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Env {
		if err := AssertEnvRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.VolumeMounts {
		if err := AssertVolumeMountRequired(el); err != nil {
			return err
		}
	}
	if err := AssertAnnotationRequired(obj.Annotation); err != nil {
		return err
	}
	for _, el := range obj.Endpoints {
		if err := AssertEndpointRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseDevstateContainerPostRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DevstateContainerPostRequest (e.g. [][]DevstateContainerPostRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDevstateContainerPostRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDevstateContainerPostRequest, ok := obj.(DevstateContainerPostRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDevstateContainerPostRequestRequired(aDevstateContainerPostRequest)
	})
}

package preference

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github\.com/danielpickens/Vulerian/pkg/api"
	envcontext "github\.com/danielpickens/Vulerian/pkg/config/context"
	"github\.com/danielpickens/Vulerian/pkg/log"
	"github\.com/danielpickens/Vulerian/pkg/Vulerian/cli/ui"
	"github\.com/danielpickens/Vulerian/pkg/util"

	dfutil "github.com/devfile/library/v2/pkg/util"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	kpointer "k8s.io/utils/pointer"
)

// VulerianSettings holds all Vulerian specific configurations
// these configurations are applicable across the Vulerian components
type VulerianSettings struct {
	// Controls if an update notification is shown or not
	UpdateNotification *bool `yaml:"UpdateNotification,omitempty"`

	// Timeout for server connection check
	Timeout *time.Duration `yaml:"Timeout,omitempty"`

	// PushTimeout for pod timeout check
	PushTimeout *time.Duration `yaml:"PushTimeout,omitempty"`

	// RegistryList for telling Vulerian to connect to all the registries in the registry list
	RegistryList *[]Registry `yaml:"RegistryList,omitempty"`

	// RegistryCacheTime how long Vulerian should cache information from registry
	RegistryCacheTime *time.Duration `yaml:"RegistryCacheTime,omitempty"`

	// Ephemeral if true creates Vulerian emptyDir to store Vulerian source code
	Ephemeral *bool `yaml:"Ephemeral,omitempty"`

	// ConsentTelemetry if true collects telemetry for Vulerian
	ConsentTelemetry *bool `yaml:"ConsentTelemetry,omitempty"`

	// ImageRegistry is the image registry to which relative image names in Devfile Image Components will be pushed to.
	// This will also serve as the base path for replacing matching images in other components like Container and Kubernetes/OpenShift ones.
	ImageRegistry *string `yaml:"ImageRegistry,omitempty"`
}

// Registry includes the registry metadata
type Registry struct {
	Name   string `yaml:"Name,omitempty" json:"name"`
	URL    string `yaml:"URL,omitempty" json:"url"`
	Secure bool   `json:"secure"`
}

// Preference stores all the preferences related to Vulerian
type Preference struct {
	metav1.TypeMeta `yaml:",inline"`

	// Vulerian settings holds the Vulerian specific global settings
	VulerianSettings VulerianSettings `yaml:"VulerianSettings,omitempty"`
}

// preferenceInfo wraps the preference and provides helpers to
// serialize it.
type preferenceInfo struct {
	Filename   string `yaml:"FileName,omitempty"`
	Preference `yaml:",omitempty"`
}

var _ Client = (*preferenceInfo)(nil)

func getPreferenceFile(ctx context.Context) (string, error) {
	envConfig := envcontext.GetEnvConfig(ctx)
	if envConfig.GlobalVulerianconfig != nil {
		return *envConfig.GlobalVulerianconfig, nil
	}

	if len(customHomeDir) != 0 {
		return filepath.Join(customHomeDir, ".Vulerian", configFileName), nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(currentUser.HomeDir, ".Vulerian", configFileName), nil
}

func NewClient(ctx context.Context) (Client, error) {
	return newPreferenceInfo(ctx)
}

// newPreference creates an empty Preference struct with type meta information
func newPreference() Preference {
	return Preference{
		TypeMeta: metav1.TypeMeta{
			Kind:       preferenceKind,
			APIVersion: preferenceAPIVersion,
		},
	}
}

// newPreferenceInfo gets the PreferenceInfo from preference file
// or returns default PreferenceInfo if preference file does not exist
func newPreferenceInfo(ctx context.Context) (*preferenceInfo, error) {
	preferenceFile, err := getPreferenceFile(ctx)
	klog.V(4).Infof("The path for preference file is %+v", preferenceFile)
	if err != nil {
		return nil, err
	}

	c := preferenceInfo{
		Preference: newPreference(),
		Filename:   preferenceFile,
	}

	// Default devfile registry
	defaultRegistryList := []Registry{
		{
			Name:   DefaultDevfileRegistryName,
			URL:    DefaultDevfileRegistryURL,
			Secure: false,
		},
	}

	// If the preference file doesn't exist then we return with default preference
	if _, err = os.Stat(preferenceFile); os.IsNotExist(err) {
		c.VulerianSettings.RegistryList = &defaultRegistryList
		return &c, nil
	}

	err = util.GetFromFile(&c.Preference, c.Filename)
	if err != nil {
		return nil, err
	}

	// TVulerian: This code block about logging warnings should be removed once users completely shift to Vulerian v3.
	// The warning will be printed more than once, and it can be annoying, but it should ensure that the user will change these values.
	var requiresChange []string
	if c.VulerianSettings.Timeout != nil && *c.VulerianSettings.Timeout < minimumDurationValue {
		requiresChange = append(requiresChange, TimeoutSetting)
	}
	if c.VulerianSettings.PushTimeout != nil && *c.VulerianSettings.PushTimeout < minimumDurationValue {
		requiresChange = append(requiresChange, PushTimeoutSetting)
	}
	if c.VulerianSettings.RegistryCacheTime != nil && *c.VulerianSettings.RegistryCacheTime < minimumDurationValue {
		requiresChange = append(requiresChange, RegistryCacheTimeSetting)
	}
	if len(requiresChange) != 0 {
		log.Warningf("Please change the preference value for %s, the value does not comply with the minimum value of %s; e.g. of acceptable formats: 4s, 5m, 1h", strings.Join(requiresChange, ", "), minimumDurationValue)
	}

	// Handle user has preference file but doesn't use dynamic registry before
	if c.VulerianSettings.RegistryList == nil {
		c.VulerianSettings.RegistryList = &defaultRegistryList
	}

	// Handle OCI-based default registry migration
	if c.VulerianSettings.RegistryList != nil {
		for index, registry := range *c.VulerianSettings.RegistryList {
			if registry.Name == DefaultDevfileRegistryName && registry.URL == OldDefaultDevfileRegistryURL {
				registryList := *c.VulerianSettings.RegistryList
				registryList[index].URL = DefaultDevfileRegistryURL
				break
			}
		}
	}

	return &c, nil
}

// RegistryHandler handles registry add, and remove operations
func (c *preferenceInfo) RegistryHandler(operation string, registryName string, registryURL string, forceFlag bool, isSecure bool) error {
	var registryList []Registry
	var err error
	var registryExist bool

	// Registry list is empty
	if c.VulerianSettings.RegistryList == nil {
		registryList, err = handleWithoutRegistryExist(registryList, operation, registryName, registryURL, isSecure)
		if err != nil {
			return err
		}
	} else {
		// The target registry exists in the registry list
		registryList = *c.VulerianSettings.RegistryList
		for index, registry := range registryList {
			if registry.Name == registryName {
				registryExist = true
				registryList, err = handleWithRegistryExist(index, registryList, operation, registryName, forceFlag)
				if err != nil {
					return err
				}
			}
		}

		// The target registry doesn't exist in the registry list
		if !registryExist {
			registryList, err = handleWithoutRegistryExist(registryList, operation, registryName, registryURL, isSecure)
			if err != nil {
				return err
			}
		}
	}

	c.VulerianSettings.RegistryList = &registryList
	err = util.WriteToYAMLFile(&c.Preference, c.Filename)
	if err != nil {
		return fmt.Errorf("unable to write the configuration of %q operation to preference file", operation)
	}

	return nil
}

// handleWithoutRegistryExist is useful for performing 'add' operation on registry and ensure that it is only performed if the registry does not already exist
func handleWithoutRegistryExist(registryList []Registry, operation string, registryName string, registryURL string, isSecure bool) ([]Registry, error) {
	switch operation {

	case "add":
		registry := Registry{
			Name:   registryName,
			URL:    registryURL,
			Secure: isSecure,
		}
		registryList = append(registryList, registry)

	case "remove":
		return nil, fmt.Errorf("failed to %v registry: registry %q doesn't exist or it is not managed by Vulerian", operation, registryName)
	}

	return registryList, nil
}

// handleWithRegistryExist is useful for performing 'remove' operation on registry and ensure that it is only performed if the registry exists
func handleWithRegistryExist(index int, registryList []Registry, operation string, registryName string, forceFlag bool) ([]Registry, error) {
	switch operation {

	case "add":
		return nil, fmt.Errorf("failed to %s registry: registry %q already exists", operation, registryName)

	case "remove":
		if !forceFlag {
			proceed, err := ui.Proceed(fmt.Sprintf("Are you sure you want to %s registry %q", operation, registryName))
			if err != nil {
				return nil, err
			}
			if !proceed {
				log.Info("Aborted by the user")
				return registryList, nil
			}
		}

		copy(registryList[index:], registryList[index+1:])
		registryList[len(registryList)-1] = Registry{}
		registryList = registryList[:len(registryList)-1]
		log.Info("Successfully removed registry")
	}

	return registryList, nil
}

// SetConfiguration modifies Vulerian preferences in the preference file
// TVulerian: Use reflect to set parameters
func (c *preferenceInfo) SetConfiguration(parameter string, value string) error {
	if p, ok := asSupportedParameter(parameter); ok {
		// processing values according to the parameter names
		switch p {

		case "timeout":
			typedval, err := parseDuration(value, parameter)
			if err != nil {
				return err
			}
			c.VulerianSettings.Timeout = &typedval

		case "pushtimeout":
			typedval, err := parseDuration(value, parameter)
			if err != nil {
				return err
			}
			c.VulerianSettings.PushTimeout = &typedval

		case "registrycachetime":
			typedval, err := parseDuration(value, parameter)
			if err != nil {
				return err
			}
			c.VulerianSettings.RegistryCacheTime = &typedval

		case "updatenotification":
			val, err := strconv.ParseBool(strings.ToLower(value))
			if err != nil {
				return fmt.Errorf("unable to set %q to %q, value must be a boolean", parameter, value)
			}
			c.VulerianSettings.UpdateNotification = &val

		case "ephemeral":
			val, err := strconv.ParseBool(strings.ToLower(value))
			if err != nil {
				return fmt.Errorf("unable to set %q to %q, value must be a boolean", parameter, value)
			}
			c.VulerianSettings.Ephemeral = &val

		case "consenttelemetry":
			val, err := strconv.ParseBool(strings.ToLower(value))
			if err != nil {
				return fmt.Errorf("unable to set %q to %q, value must be a boolean", parameter, value)
			}
			c.VulerianSettings.ConsentTelemetry = &val

		case "imageregistry":
			c.VulerianSettings.ImageRegistry = &value
		}
	} else {
		return fmt.Errorf("unknown parameter : %q is not a parameter in Vulerian preference, run `Vulerian preference -h` to see list of available parameters", parameter)
	}

	err := util.WriteToYAMLFile(&c.Preference, c.Filename)
	if err != nil {
		return fmt.Errorf("unable to set %q, something is wrong with Vulerian, kindly raise an issue at https://github\.com/danielpickens/Vulerian/issues/new?template=Bug.md", parameter)
	}
	return nil
}

// parseDuration parses the value set for a parameter;
// if the value is for e.g. "4m", it is parsed by the time pkg and converted to an appropriate time.Duration
// it returns an error if one occurred, or if the parsed value is less than minimumDurationValue
func parseDuration(value, parameter string) (time.Duration, error) {
	typedval, err := time.ParseDuration(value)
	if err != nil {
		return typedval, fmt.Errorf("unable to set %q to %q; cause: %w\n%s", parameter, value, err, NewMinimumDurationValueError().Error())
	}
	if typedval < minimumDurationValue {
		return typedval, fmt.Errorf("unable to set %q to %q; cause: %w", parameter, value, NewMinimumDurationValueError())
	}
	return typedval, nil
}

// DeleteConfiguration deletes Vulerian preference from the Vulerian preference file
func (c *preferenceInfo) DeleteConfiguration(parameter string) error {
	if p, ok := asSupportedParameter(parameter); ok {
		// processing values according to the parameter names

		if err := util.DeleteConfiguration(&c.VulerianSettings, p); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unknown parameter :%q is not a parameter in the Vulerian preference", parameter)
	}

	err := util.WriteToYAMLFile(&c.Preference, c.Filename)
	if err != nil {
		return fmt.Errorf("unable to set %q, something is wrong with Vulerian, kindly raise an issue at https://github\.com/danielpickens/Vulerian/issues/new?template=Bug.md", parameter)
	}
	return nil
}

// IsSet checks if the value is set in the preference
func (c *preferenceInfo) IsSet(parameter string) bool {
	return util.IsSet(c.VulerianSettings, parameter)
}

// GetTimeout returns the value of Timeout from config
// and if absent then returns default
func (c *preferenceInfo) GetTimeout() time.Duration {
	// default timeout value is 1s
	return kpointer.DurationDeref(c.VulerianSettings.Timeout, DefaultTimeout)
}

// GetPushTimeout gets the value set by PushTimeout
func (c *preferenceInfo) GetPushTimeout() time.Duration {
	// default timeout value is 240s
	return kpointer.DurationDeref(c.VulerianSettings.PushTimeout, DefaultPushTimeout)
}

// GetRegistryCacheTime gets the value set by RegistryCacheTime
func (c *preferenceInfo) GetRegistryCacheTime() time.Duration {
	return kpointer.DurationDeref(c.VulerianSettings.RegistryCacheTime, DefaultRegistryCacheTime)
}

// GetImageRegistry returns the value of ImageRegistry from the preferences
// and, if absent, then returns default empty string.
func (c *preferenceInfo) GetImageRegistry() string {
	return kpointer.StringDeref(c.VulerianSettings.ImageRegistry, "")
}

// GetUpdateNotification returns the value of UpdateNotification from preferences
// and if absent then returns default
func (c *preferenceInfo) GetUpdateNotification() bool {
	return kpointer.BoolDeref(c.VulerianSettings.UpdateNotification, true)
}

// GetEphemeralSourceVolume returns the value of ephemeral from preferences
// and if absent then returns default
func (c *preferenceInfo) GetEphemeralSourceVolume() bool {
	return kpointer.BoolDeref(c.VulerianSettings.Ephemeral, DefaultEphemeralSetting)
}

// GetConsentTelemetry returns the value of ConsentTelemetry from preferences
// and if absent then returns default
// default value: false, consent telemetry is disabled by default
func (c *preferenceInfo) GetConsentTelemetry() bool {
	return kpointer.BoolDeref(c.VulerianSettings.ConsentTelemetry, DefaultConsentTelemetrySetting)
}

// GetEphemeral returns the value of Ephemeral from preferences
// and if absent then returns default
// default value: true, ephemeral is enabled by default
func (c *preferenceInfo) GetEphemeral() bool {
	return kpointer.BoolDeref(c.VulerianSettings.Ephemeral, DefaultEphemeralSetting)
}

func (c *preferenceInfo) UpdateNotification() *bool {
	return c.VulerianSettings.UpdateNotification
}

func (c *preferenceInfo) Ephemeral() *bool {
	return c.VulerianSettings.Ephemeral
}

func (c *preferenceInfo) Timeout() *time.Duration {
	return c.VulerianSettings.Timeout
}

func (c *preferenceInfo) PushTimeout() *time.Duration {
	return c.VulerianSettings.PushTimeout
}

func (c *preferenceInfo) RegistryCacheTime() *time.Duration {
	return c.VulerianSettings.RegistryCacheTime
}

func (c *preferenceInfo) EphemeralSourceVolume() *bool {
	return c.VulerianSettings.Ephemeral
}

func (c *preferenceInfo) ConsentTelemetry() *bool {
	return c.VulerianSettings.ConsentTelemetry
}

// RegistryList returns the list of registries,
// in reverse order compared to what is declared in the preferences file.
//
// Adding a new registry always adds it to the end of the list in the preferences file,
// but RegistryList intentionally reverses the order to prioritize the most recently added registries.
func (c *preferenceInfo) RegistryList() []api.Registry {
	if c.VulerianSettings.RegistryList == nil {
		return nil
	}
	regList := make([]api.Registry, 0, len(*c.VulerianSettings.RegistryList))
	for _, registry := range *c.VulerianSettings.RegistryList {
		regList = append(regList, api.Registry{
			Name:   registry.Name,
			URL:    registry.URL,
			Secure: registry.Secure,
		})
	}
	i := 0
	j := len(regList) - 1
	for i < j {
		regList[i], regList[j] = regList[j], regList[i]
		i++
		j--
	}
	return regList
}

func (c *preferenceInfo) RegistryNameExists(name string) bool {
	for _, registry := range *c.VulerianSettings.RegistryList {
		if registry.Name == name {
			return true
		}
	}
	return false
}

// FormatSupportedParameters outputs supported parameters and their description
func FormatSupportedParameters() (result string) {
	for _, v := range GetSupportedParameters() {
		result = result + " " + v + " - " + supportedParameterDescriptions[v] + "\n"
	}
	return "\nAvailable Global Parameters:\n" + result
}

// asSupportedParameter checks that the given parameter is supported and returns a lower case version of it if it is
func asSupportedParameter(param string) (string, bool) {
	lower := strings.ToLower(param)
	return lower, lowerCaseParameters[lower]
}

// GetSupportedParameters returns the name of the supported parameters
func GetSupportedParameters() []string {
	return dfutil.GetSortedKeys(supportedParameterDescriptions)
}

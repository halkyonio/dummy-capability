package plugin

import (
	v1beta12 "halkyon.io/api/capability/v1beta1"
	"halkyon.io/api/v1beta1"
	framework "halkyon.io/operator-framework"
	v1 "k8s.io/api/core/v1"
	"strings"
)

func OwnerAsCapability(res framework.DependentResource) *v1beta12.Capability {
	return res.Owner().(*v1beta12.Capability)
}

// Convert Array of parameters to a Map
func ParametersAsMap(parameters []v1beta1.NameValuePair) map[string]string {
	result := make(map[string]string)
	for _, parameter := range parameters {
		result[parameter.Name] = parameter.Value
	}
	return result
}

// TODO : TO BE REVIEWED
func GetSecretOrDefault(needsSecret NeedsSecret, parameters map[string]string) *v1.SecretVolumeSource {
	if secretName, ok := parameters[DummyConfigName]; ok {
		return &v1.SecretVolumeSource{SecretName: secretName}
	} else {
		// generate default secret name
		return &v1.SecretVolumeSource{SecretName: needsSecret.GetSecretName()}
	}
}

// TODO : TO BE REVIEWED
func DefaultSecretNameFor(secretOwner NeedsSecret) string {
	c := secretOwner.Owner().(*v1beta12.Capability)
	paramsMap := ParametersAsMap(c.Spec.Parameters)
	return SetDefaultSecretNameIfEmpty(c.Name, paramsMap[DummyConfigName])
}

// TODO : TO BE REVIEWED
func SetDefaultSecretNameIfEmpty(capabilityName, paramSecretName string) string {
	if paramSecretName == "" {
		return strings.ToLower(capabilityName) + "-config"
	} else {
		return paramSecretName
	}
}

//getAppLabels returns an string map with the labels which wil be associated to the kubernetes/ocp resource which will be created and managed by this operator
func GetAppLabels(name string) map[string]string {
	return map[string]string{
		"app": name,
	}
}

func ReplicaNumber(num int) *int32 {
	q := int32(num)
	return &q
}

package plugin

import (
	v1beta12 "halkyon.io/api/capability/v1beta1"
	"halkyon.io/api/v1beta1"
	framework "halkyon.io/operator-framework"
)

func OwnerAsCapability(res framework.DependentResource) *v1beta12.Capability {
	return res.Owner().(*v1beta12.Capability)
}

// Convert Array of parameters to a Map
func ParametersAsMap(parameters []v1beta1.NameValuePair) map[string][]byte {
	result := make(map[string][]byte)
	for _, parameter := range parameters {
		result[parameter.Name] = []byte(parameter.Value)
	}
	return result
}

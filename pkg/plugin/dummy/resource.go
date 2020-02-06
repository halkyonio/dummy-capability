package dummy

import (
	"github.com/hashicorp/go-hclog"
	v1capability "halkyon.io/api/capability/v1beta1"
	"halkyon.io/api/v1beta1"
	"halkyon.io/operator-framework"
	"halkyon.io/operator-framework/plugins/capability"
)

var _ capability.PluginResource = &DummyPluginResource{}
var versionsMapping = make(map[string]string, 11)

func NewPluginResource() capability.PluginResource {
	return &DummyPluginResource{capability.NewQueryingSimplePluginResourceStem(v1capability.LoggingCategory, resolver)}
}

// TODO: Check how to get the namespace
func resolver(logger hclog.Logger) capability.TypeInfo {
	info := capability.TypeInfo{
		Type:     v1capability.CapabilityType(dummyType),
		Versions: []string{dummyVersion},
	}
	return info
}

type DummyPluginResource struct {
	capability.QueryingSimplePluginResourceStem
}

func (p *DummyPluginResource) GetDependentResourcesWith(owner v1beta1.HalkyonResource) []framework.DependentResource {
	/*dummy := NewDummy(owner)
	return []framework.DependentResource{
		framework.NewOwnedRole(dummy),
		plugin.NewRoleBinding(dummy),
		plugin.NewSecret(dummy),
		dummy,
	}*/
	return nil
}

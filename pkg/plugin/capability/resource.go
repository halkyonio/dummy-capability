package capability

import (
	"github.com/hashicorp/go-hclog"
	v1capability "halkyon.io/api/capability/v1beta1"
	"halkyon.io/operator-framework"
	"halkyon.io/operator-framework/plugins/capability"
)

var _ capability.PluginResource = &PluginResource{}

// The PluginResource
func NewPluginResource() capability.PluginResource {
	return &PluginResource{SimplePluginResourceStem: capability.NewSimplePluginResourceStem(
		"example",
		capability.TypeInfo{
			Type:     "foo",
			Versions: []string{"1.0"},
		},
	)}
}

func (p *PluginResource) SetLogger(logger hclog.Logger) {
	p.logger = logger
}

type PluginResource struct {
	capability.SimplePluginResourceStem
	logger hclog.Logger
}

func (p *PluginResource) GetSupportedCategory() v1capability.CapabilityCategory {
	p.logger.Info("calling GetSupportedCategory")
	return p.SimplePluginResourceStem.GetSupportedCategory()
}

func (p *PluginResource) GetSupportedTypes() []capability.TypeInfo {
	p.logger.Info("calling GetSupportedTypes")
	return p.SimplePluginResourceStem.GetSupportedTypes()
}

func (p *PluginResource) GetDependentResourcesWith(owner framework.SerializableResource) []framework.DependentResource {
	p.logger.Info("calling GetDependentResourcesWith")
	ownerResource := NewOwnerResource(owner)
	return []framework.DependentResource{
		ownerResource,
		framework.NewSecret(ownerResource),
	}
}

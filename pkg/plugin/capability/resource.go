package capability

import (
	"github.com/hashicorp/go-hclog"
	v1capability "halkyon.io/api/capability/v1beta1"
	"halkyon.io/api/v1beta1"
	"halkyon.io/example-capability/pkg/plugin"
	"halkyon.io/operator-framework"
	"halkyon.io/operator-framework/plugins/capability"
)

var _ capability.PluginResource = &PluginResource{}

// The PluginResource
func NewPluginResource() capability.PluginResource {
	return &PluginResource{SimplePluginResourceStem: capability.NewSimplePluginResourceStem(
		v1capability.LoggingCategory,
		capability.TypeInfo{
		   Type:     "logrus",
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

func (p *PluginResource) GetDependentResourcesWith(owner v1beta1.HalkyonResource) []framework.DependentResource {
	p.logger.Info("calling GetDependentResourcesWith")
	ownerResource := NewOwnerResource(owner)
	return []framework.DependentResource{
		ownerResource,
		plugin.NewSecret(ownerResource),
	}
}

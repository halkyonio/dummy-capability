package dummy

import (
	"github.com/hashicorp/go-hclog"
	v1capability "halkyon.io/api/capability/v1beta1"
	"halkyon.io/api/v1beta1"
	"halkyon.io/operator-framework"
	"halkyon.io/operator-framework/plugins/capability"
)

var _ capability.PluginResource = &DummyPluginResource{}

func NewPluginResource() capability.PluginResource {
	return &DummyPluginResource{SimplePluginResourceStem: capability.NewSimplePluginResourceStem(v1capability.LoggingCategory, capability.TypeInfo{
		Type:     "logrus",
		Versions: []string{"1.0"},
	})}
}

func (p *DummyPluginResource) SetLogger(logger hclog.Logger) {
	p.logger = logger
}

type DummyPluginResource struct {
	capability.SimplePluginResourceStem
	logger hclog.Logger
}

func (p *DummyPluginResource) GetSupportedCategory() v1capability.CapabilityCategory {
	p.logger.Info("calling GetSupportedCategory")
	return p.SimplePluginResourceStem.GetSupportedCategory()
}

func (p *DummyPluginResource) GetSupportedTypes() []capability.TypeInfo {
	p.logger.Info("calling GetSupportedTypes")
	return p.SimplePluginResourceStem.GetSupportedTypes()
}

func (p *DummyPluginResource) GetDependentResourcesWith(owner v1beta1.HalkyonResource) []framework.DependentResource {
	p.logger.Info("calling GetDependentResourcesWith")
	/*dummy := NewDummy(owner)
	return []framework.DependentResource{
		framework.NewOwnedRole(dummy),
		plugin.NewRoleBinding(dummy),
		plugin.NewSecret(dummy),
		dummy,
	}*/
	return nil
}

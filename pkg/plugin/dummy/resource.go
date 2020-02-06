package dummy

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"halkyon.io/api/capability/v1beta1"
	v1beta12 "halkyon.io/api/v1beta1"
	"halkyon.io/dummy-capability/pkg/plugin"
	"halkyon.io/operator-framework"
	"halkyon.io/operator-framework/plugins/capability"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ capability.PluginResource = &DummyPluginResource{}
var versionsMapping = make(map[string]string, 11)

func NewPluginResource() capability.PluginResource {
	return &DummyPluginResource{capability.NewQueryingSimplePluginResourceStem(v1beta1.LoggingCategory, resolver)}
}

// TODO: Check how to get the namespace
func resolver(logger hclog.Logger) capability.TypeInfo {
	list, err := plugin.Client.Pods("").List(v1.ListOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("error retrieving versions: %v", err))
	}
	versions := make([]string, 0, len(list.Items))
	info := capability.TypeInfo{
		Type:     "DummyKind",
		Versions: versions,
	}
	return info
}

type DummyPluginResource struct {
	capability.QueryingSimplePluginResourceStem
}

func (p *DummyPluginResource) GetDependentResourcesWith(owner v1beta12.HalkyonResource) []framework.DependentResource {
	dummy := NewDummy(owner)
	return []framework.DependentResource{
		framework.NewOwnedRole(dummy),
		plugin.NewRoleBinding(dummy),
		plugin.NewSecret(dummy),
		dummy,
	}
}

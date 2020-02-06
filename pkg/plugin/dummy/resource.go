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

var _ capability.PluginResource = &PostgresPluginResource{}
var versionsMapping = make(map[string]string, 11)

func NewPluginResource() capability.PluginResource {
	return &PostgresPluginResource{capability.NewQueryingSimplePluginResourceStem(v1beta1.DatabaseCategory, resolver)}
}

func resolver(logger hclog.Logger) capability.TypeInfo {
	list, err := plugin.Client.Deployments("").List(v1.ListOptions{})
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

type PostgresPluginResource struct {
	capability.QueryingSimplePluginResourceStem
}

func (p *PostgresPluginResource) GetDependentResourcesWith(owner v1beta12.HalkyonResource) []framework.DependentResource {
	postgres := NewDummy(owner)
	return []framework.DependentResource{
		framework.NewOwnedRole(postgres),
		plugin.NewRoleBinding(postgres),
		plugin.NewSecret(postgres),
		postgres,
	}
}

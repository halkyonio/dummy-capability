package dummy

import (
	"fmt"
	"halkyon.io/api/v1beta1"
	"halkyon.io/dummy-capability/pkg/plugin"
	framework "halkyon.io/operator-framework"
	apps "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	// Dummy const
	key1VarName     = "KEY1"
	key2VarName     = "KEY2"
)

var (
	// dummyGVK = kubedbv1.SchemeGroupVersion.WithKind(kubedbv1.ResourceKindPostgres)
)

type dummy struct {
	*framework.BaseDependentResource
}

func (res dummy) Fetch() (runtime.Object, error) {
	panic("should never be called")
}

var _ framework.DependentResource = &dummy{}

func (res dummy) Update(_ runtime.Object) (bool, error) {
	return false, nil
}

func NewPostgres(owner v1beta1.HalkyonResource) *dummy {
	config := framework.NewConfig(dummyGVK)
	config.CheckedForReadiness = true
	config.OwnerStatusField = "PodName" // todo: find a way to compute this as above instead of hardcoding it
	p := &dummy{framework.NewConfiguredBaseDependentResource(owner, config)}
	return p
}

func (res dummy) Name() string {
	return framework.DefaultDependentResourceNameFor(res.Owner())
}

//buildSecret returns the dummy resource
func (res dummy) Build(empty bool) (runtime.Object, error) {
	dummy := &kubedbv1.Postgres{}
	if !empty {
		c := plugin.OwnerAsCapability(res)
		ls := plugin.GetAppLabels(c.Name)
		dummy.ObjectMeta = metav1.ObjectMeta{
			Name:      res.Name(),
			Namespace: c.Namespace,
			Labels:    ls,
		}
		dummy.Spec = kubedbv1.PostgresSpec{
			Version:  plugin.GetVersionFrom(c, versionsMapping),
			Replicas: plugin.ReplicaNumber(1),
			UpdateStrategy: apps.StatefulSetUpdateStrategy{
				Type: apps.RollingUpdateStatefulSetStrategyType,
			},
			StorageType:       kubedbv1.StorageTypeEphemeral,
			TerminationPolicy: kubedbv1.TerminationPolicyDelete,
		}

		paramsMap := plugin.ParametersAsMap(c.Spec.Parameters)
		if secret := plugin.GetSecretOrDefault(res, paramsMap); secret != nil {
			dummy.Spec.DatabaseSecret = secret
		}
		if dbNameConfig := plugin.GetDatabaseNameConfigOrNil(dbNameVarName, paramsMap); dbNameConfig != nil {
			dummy.Spec.PodTemplate = *dbNameConfig
		}
	}
	return dummy, nil
}

func (res dummy) IsReady(underlying runtime.Object) (ready bool, message string) {
	psql := underlying.(*kubedbv1.Postgres)
	ready = psql.Status.Phase == kubedbv1.DatabasePhaseRunning
	if !ready {
		msg := ""
		reason := psql.Status.Reason
		if len(reason) > 0 {
			msg = ": " + reason
		}
		message = fmt.Sprintf("%s is not ready%s", psql.Name, msg)
	}
	return
}

func (res dummy) NameFrom(underlying runtime.Object) string {
	return underlying.(*kubedbv1.Postgres).Name
}

func (res dummy) GetRoleBindingName() string {
	return "use-scc-privileged"
}

func (res dummy) GetAssociatedRoleName() string {
	return res.GetRoleName()
}

func (res dummy) GetServiceAccountName() string {
	return res.Name()
}

func (res dummy) GetRoleName() string {
	return "scc-privileged-role"
}

func (res dummy) GetDataMap() map[string][]byte {
	c := plugin.OwnerAsCapability(res)
	paramsMap := plugin.ParametersAsMap(c.Spec.Parameters)
	return map[string][]byte{
		dbUserVarName:     []byte(paramsMap[plugin.DbUser]),
		dbPasswordVarName: []byte(paramsMap[plugin.DbPassword]),
		dbNameVarName:     []byte(plugin.SetDefaultDatabaseName(paramsMap[plugin.DbName])),
		// TODO : To be reviewed according to the discussion started with issue #75
		// as we will create another secret when a link will be issued
		plugin.DbHost:     []byte(plugin.SetDefaultDatabaseHost(c.Name, paramsMap[plugin.DbHost])),
		plugin.DbPort:     []byte(plugin.SetDefaultDatabasePort(paramsMap[plugin.DbPort])),
		plugin.DbName:     []byte(plugin.SetDefaultDatabaseName(paramsMap[plugin.DbName])),
		plugin.DbUser:     []byte((paramsMap[plugin.DbUser])),
		plugin.DbPassword: []byte(paramsMap[plugin.DbPassword]),
	}
}

func (res dummy) GetSecretName() string {
	return plugin.DefaultSecretNameFor(res)
}

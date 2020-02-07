package dummy

import (
	"fmt"
	"halkyon.io/api/v1beta1"
	"halkyon.io/dummy-capability/pkg/plugin"
	framework "halkyon.io/operator-framework"
	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	// Dummy const
	key1VarName = "KEY1"
	key2VarName = "KEY2"
)

var podGVK = v1.SchemeGroupVersion.WithKind("Pod")

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

func NewDummy(owner v1beta1.HalkyonResource) *dummy {
	config := framework.NewConfig(podGVK)
	config.CheckedForReadiness = true
	config.OwnerStatusField = "PodName"
	p := &dummy{framework.NewConfiguredBaseDependentResource(owner, config)}
	return p
}

func (res dummy) Name() string {
	return framework.DefaultDependentResourceNameFor(res.Owner())
}

//Build returns a Pod resource
func (res dummy) Build(empty bool) (runtime.Object, error) {
	pod := &v1.Pod{}
	if !empty {
		c := plugin.OwnerAsCapability(res)
		ls := plugin.GetAppLabels(c.Name)
		pod.ObjectMeta = metav1.ObjectMeta{
			Name:      res.Name(),
			Namespace: c.Namespace,
			Labels:    ls,
		}
		pod.Spec = v1.PodSpec{
			Containers: []v1.Container{
				v1.Container{
					Name:  "dummy",
					Image: "busybox",
				},
			},
		}
	}
	return pod, nil
}

// Check if the status of the Deployment is ready
func (res dummy) IsReady(underlying runtime.Object) (ready bool, message string) {
	deploy := underlying.(*v1.Pod)
	ready = deploy.Status.Conditions[0].Status == v1.ConditionTrue
	if !ready {
		msg := ""
		reason := deploy.Status.Conditions[0].Reason
		if len(reason) > 0 {
			msg = ": " + reason
		}
		message = fmt.Sprintf("%s is not ready%s", dummy.Name, msg)
	}
	return
}

// Return the name of the Kubernetes Deployment Resources
func (res dummy) NameFrom(underlying runtime.Object) string {
	return underlying.(*apps.Deployment).Name
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
	_ = plugin.ParametersAsMap(c.Spec.Parameters)
	return map[string][]byte{
		// TODO
	}
}

func (res dummy) GetSecretName() string {
	return plugin.DefaultSecretNameFor(res)
}

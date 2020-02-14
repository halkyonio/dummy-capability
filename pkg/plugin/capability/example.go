package capability

import (
	"fmt"
	"halkyon.io/api/v1beta1"
	"halkyon.io/example-capability/pkg/plugin"
	framework "halkyon.io/operator-framework"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var podGVK = v1.SchemeGroupVersion.WithKind("Pod")

type example struct {
	*framework.BaseDependentResource
}

func (res example) Fetch() (runtime.Object, error) {
	panic("should never be called")
}

var _ framework.DependentResource = &example{}

func (res example) Update(_ runtime.Object) (bool, error) {
	return false, nil
}

func NewOwnerResource(owner v1beta1.HalkyonResource) *example {
	config := framework.NewConfig(podGVK)
	config.CheckedForReadiness = true
	config.OwnerStatusField = "PodName"
	p := &example{framework.NewConfiguredBaseDependentResource(owner, config)}
	return p
}

func (res example) Name() string {
	return framework.DefaultDependentResourceNameFor(res.Owner())
}

//Build returns a Pod resource
func (res example) Build(empty bool) (runtime.Object, error) {
	pod := &v1.Pod{}
	if !empty {
		c := plugin.OwnerAsCapability(res)
		pod.ObjectMeta = metav1.ObjectMeta{
			Name:      res.Name(),
			Namespace: c.Namespace,
		}
		pod.Spec = v1.PodSpec{
			Containers: []v1.Container{
				v1.Container{
					Name:  "example",
					Image: "busybox",
				},
			},
		}
	}
	return pod, nil
}

// Check if the status of the Deployment is ready
func (res example) IsReady(underlying runtime.Object) (ready bool, message string) {
	deploy := underlying.(*v1.Pod)
	ready = deploy.Status.Conditions[0].Status == v1.ConditionTrue
	if !ready {
		msg := ""
		reason := deploy.Status.Conditions[0].Reason
		if len(reason) > 0 {
			msg = ": " + reason
		}
		message = fmt.Sprintf("%s is not ready%s", example.Name, msg)
	}
	return
}

// Return the name of the Kubernetes Deployment Resources
func (res example) NameFrom(underlying runtime.Object) string {
	return underlying.(*v1.Pod).Name
}

func (res example) GetDataMap() map[string][]byte {
	c := plugin.OwnerAsCapability(res)
	return plugin.ParametersAsMap(c.Spec.Parameters)
}

func (res example) GetSecretName() string {
	return res.Owner().GetName() + "-secret"
}

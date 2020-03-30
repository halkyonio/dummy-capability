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

// Records the GVK for the underlying type we're interested in working with (here, a Pod)
var podGVK = v1.SchemeGroupVersion.WithKind("Pod")

// example is a simple, example implementation of DependentResource
type example struct {
	*framework.BaseDependentResource
}

// blank assignment to make sure that our struct properly implements the DependentResource interface
var _ framework.DependentResource = &example{}

// NewOwnerResource creates a new example instance given the specified owner Resource as a SerializableResource
func NewOwnerResource(owner framework.SerializableResource) *example {
	// Create a new, default config with the specified GVK
	config := framework.NewConfig(podGVK)
	// Override some of the default configuration if needed, here we want to check this dependent for its
	// readiness when computing the owner's status
	config.CheckedForReadiness = true
	// Create an instance of the struct, properly initializing the embedded BaseDependentResource
	p := &example{framework.NewConfiguredBaseDependentResource(owner, config)}
	return p
}

func (res example) Update(toUpdate runtime.Object) (bool, runtime.Object, error) {
	return false, toUpdate, nil
}

func (res example) GetCondition(underlying runtime.Object, err error) *v1beta1.DependentCondition {
	return framework.DefaultCustomizedGetConditionFor(res, err, underlying, func(underlying runtime.Object, cond *v1beta1.DependentCondition) {
		pod := underlying.(*v1.Pod)
		for _, c := range pod.Status.Conditions {
			if c.Type == v1.PodReady {
				cond.Type = v1beta1.DependentReady
				if c.Status != v1.ConditionTrue {
					cond.Type = v1beta1.DependentPending
				}
				cond.Message = c.Message
				cond.Reason = c.Reason
			}
		}
		return
	})
}

func (res example) Fetch() (runtime.Object, error) {
	panic("should never be called")
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
				{
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

func (res example) GetDataMap() map[string][]byte {
	c := plugin.OwnerAsCapability(res)
	return plugin.ParametersAsMap(c.Spec.Parameters)
}

func (res example) GetSecretName() string {
	return res.Owner().GetName() + "-secret"
}

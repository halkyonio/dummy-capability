package plugin

import (
	"halkyon.io/api/v1beta1"
	framework "halkyon.io/operator-framework"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var secretGVK = v1.SchemeGroupVersion.WithKind("Secret")

type NeedsSecret interface {
	GetDataMap() map[string][]byte
	GetSecretName() string
	Owner() v1beta1.HalkyonResource
}

type secret struct {
	*framework.BaseDependentResource
	Delegate NeedsSecret
}

func (res secret) NameFrom(underlying runtime.Object) string {
	return framework.DefaultNameFrom(res, underlying)
}

func (res secret) Fetch() (runtime.Object, error) {
	return framework.DefaultFetcher(res)
}

func (res secret) IsReady(underlying runtime.Object) (ready bool, message string) {
	return framework.DefaultIsReady(underlying)
}

var _ framework.DependentResource = &secret{}

func (res secret) Update(_ runtime.Object) (bool, error) {
	return false, nil
}

func NewSecret(owner NeedsSecret) secret {
	config := framework.NewConfig(secretGVK)
	config.Watched = false
	return secret{BaseDependentResource: framework.NewConfiguredBaseDependentResource(owner.Owner(), config), Delegate: owner}
}

//buildSecret returns the secret resource
func (res secret) Build(empty bool) (runtime.Object, error) {
	secret := &v1.Secret{}
	if !empty {
		c := OwnerAsCapability(res)
		secret.ObjectMeta = metav1.ObjectMeta{
			Name:      res.Name(),
			Namespace: c.Namespace,
		}
		secret.Data = res.Delegate.GetDataMap()
	}

	return secret, nil
}

func (res secret) Name() string {
	return res.Delegate.GetSecretName()
}

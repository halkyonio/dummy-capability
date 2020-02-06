module halkyon.io/dummy-capability

go 1.13

require (
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	halkyon.io/api v1.0.0-rc.3
	halkyon.io/operator-framework v1.0.0-beta.3.0.20200123132453-20459cfad8a3
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v11.0.1-0.20190805182715-88a2adca7e76+incompatible
	sigs.k8s.io/controller-runtime v0.3.0
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190805182251-6c9aa3caf3d6 // kubernetes-1.14.5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20190811223248-5a95b2df4348
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190805182715-88a2adca7e76+incompatible
)

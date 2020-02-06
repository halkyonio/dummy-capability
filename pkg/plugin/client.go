package plugin

import (
	client "k8s.io/client-go/kubernetes"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

var (
	// Return a Client accessing the Resource tht you will like to handle with your Dummy Capability
	// Eg. Kubernetes Pod, ...
	Client = client.NewForConfigOrDie(controllerruntime.GetConfigOrDie()).CoreV1()
)

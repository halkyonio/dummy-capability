package plugin

import (
	client "k8s.io/client-go/kubernetes"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

var (
	// Return a Client accessing the Kubernetes Core V1
	Client = client.NewForConfigOrDie(controllerruntime.GetConfigOrDie()).AppsV1()
)

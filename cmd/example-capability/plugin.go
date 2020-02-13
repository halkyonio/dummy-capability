package main

import (
	"halkyon.io/example-capability/pkg/plugin/capability"
	plugins "halkyon.io/operator-framework/plugins/capability"
)

func main() {
	plugins.StartPluginServerFor(capability.NewPluginResource())
}

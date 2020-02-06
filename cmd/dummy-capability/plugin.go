package main

import (
	"halkyon.io/dummy-capability/pkg/plugin/dummy"
	plugins "halkyon.io/operator-framework/plugins/capability"
)

func main() {
	plugins.StartPluginServerFor(dummy.NewPluginResource())
}

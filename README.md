# Instructions to create a new plugin

The process to create a new `plugin` able to manage a new capability for the `halkyon` operator is pretty forward 
and consists in a few steps :

- Git clone the project `https://github.com/halkyonio/plugin-example-capability` and rename it 
- Rename the `module name` as defined within the `go.mod` file to use your package name
- Find and replace the `example` word with the name of your `plugin`, resource, ...
- Build and deploy it under the local `plugins` folder of the `halkyon` operator to test it

## Build the plugin rpc server

To build the plugin, execute the following command where
```bash
go build -o example-plugin ./cmd/example-capability/plugin.go
```
If you plan to debug, simply  add the following compilation parameters
```bash
go build -gcflags="all=-N -l" -o example-plugin ./cmd/example-capability/plugin.go
```
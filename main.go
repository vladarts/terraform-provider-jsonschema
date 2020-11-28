package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/xxxbobrxxx/terraform-provider-jsonschema/jsonschema"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: jsonschema.Provider})
}

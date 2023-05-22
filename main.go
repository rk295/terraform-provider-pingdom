package main

import (
	// // document generation
	// _ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/rk295/terraform-provider-pingdom/pingdom"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pingdom.Provider,
	})

}

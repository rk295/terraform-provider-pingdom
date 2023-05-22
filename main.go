package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/rk295/terraform-provider-pingdom/pingdom"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pingdom.Provider,
	})
}

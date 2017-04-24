package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/gridscale/terraform-provider-gridscale/gridscale"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gridscale.Provider,
	})
}

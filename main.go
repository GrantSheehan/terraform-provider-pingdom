package main

import (
	"github.com/GrantSheehan/terraform-provider-pingdom/pingdom"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pingdom.Provider,
	})
}

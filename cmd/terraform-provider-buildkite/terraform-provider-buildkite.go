package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"

	"github.com/saymedia/terraform-buildkite/buildkite/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}

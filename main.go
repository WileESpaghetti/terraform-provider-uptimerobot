package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/WileESpaghetti/terraform-provider-uptimerobot/uptimerobot"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: uptimerobot.Provider})
}

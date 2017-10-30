package uptimerobot

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"uptimerobot_monitor": resourceMonitor(),
		},
	}
}

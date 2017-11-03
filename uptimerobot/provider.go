package uptimerobot

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,
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

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ApiKey: d.Get("api_key").(string),
	}

	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	return client, nil
}

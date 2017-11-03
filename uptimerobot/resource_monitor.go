package uptimerobot

import (
	"github.com/WileESpaghetti/go-uptimemonitor-v2/v2"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitorCreate,
		Read:   resourceMonitorRead,
		Update: resourceMonitorUpdate,
		Delete: resourceMonitorDelete,

		Schema: map[string]*schema.Schema{
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(uptimerobot.Client)
	monitor := uptimerobot.Monitor{
		FriendlyName: d.Get("friendly_name").(string),
		Url: d.Get("url").(string),
		Type: 1,
	}

	_, err := client.NewMonitor(&monitor)
	if err != nil {
		// TODO update state on already_exists errors
		return err
	}

	d.SetId(monitor.Id.String())

	return resourceMonitorCreate(d, meta)
}

func resourceMonitorRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMonitorDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

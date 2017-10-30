package uptimerobot

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitorCreate,
		Read:   resourceMonitorRead,
		Update: resourceMonitorUpdate,
		Delete: resourceMonitorDelete,

		Schema: map[string]*schema.Schema{},
	}
}

func resourceMonitorCreate(d *schema.ResourceData, m interface{}) error {
	return nil
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

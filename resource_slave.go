package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSlave() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlaveCreate,
		Read:   resourceSlaveRead,
		Update: resourceSlaveUpdate,
		Delete: resourceSlaveDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSlaveCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return nil
}

func resourceSlaveRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSlaveUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSlaveDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceEvent() *schema.Resource {
	return &schema.Resource{
		Create: resourceWcProductCreate,
		Read:   resourceWcProductRead,
		Update: resourceWcProductUpdate,
		Delete: resourceWcProductDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"regular_price": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"images": &schema.Schema{
				Type:     schema.TypeList,
				Required: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"categories": &schema.Schema{
				Type:     schema.TypeList,
				Required: false,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}
func resourceWcProductCreate(d *schema.ResourceData, meta interface{}) error {
	// TODO
	return nil
}

func resourceWcProductRead(d *schema.ResourceData, meta interface{}) error {
	// TODO
	return nil
}

func resourceWcProductUpdate(d *schema.ResourceData, meta interface{}) error {
	// TODO
	return nil
}

func resourceWcProductDelete(d *schema.ResourceData, meta interface{}) error {
	// TODO
	return nil
}

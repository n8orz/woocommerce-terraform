package woocommerce

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"woocommerce_key": {
				Type:     schema.TypeString,
				Optional: false,
			},
			"woocommerce_secret": {
				Type:     schema.TypeString,
				Optional: false,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Optional: false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"woocommerce_product": resourceEvent(),
		},
		ConfigureFunc: providerConfigure,
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey: d.Get("woocommerce_key").(string),
		SecretKey: d.Get("woocommerce_secret").(string),
		Endpoint:  d.Get("endpoint").(string),
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", d.Get("endpoint").(string)+"/wp-json/wc/v3/orders", nil)
	if err != nil {
		return err, nil
	}
	req.SetBasicAuth(d.Get("woocommerce_key").(string), d.Get("woocommerce_secret").(string))
	resp, err := client.Do(req)
	println(resp)
	if err != nil {
		return err, nil
	}
	return &config, nil
}

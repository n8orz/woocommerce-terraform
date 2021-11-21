package woocommerce

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"wc_key": {
				Type:     schema.TypeString,
				Optional: false,
			},
			"wc_secret": {
				Type:     schema.TypeString,
				Optional: false,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Optional: false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"wc_product": resourceEvent(),
		},
		ConfigureFunc: providerConfigure,
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey: d.Get("wc_key").(string),
		SecretKey: d.Get("wc_secret").(string),
		Endpoint:  d.Get("endpoint").(string),
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", d.Get("endpoint").(string)+"/wp-json/wc/v3/orders", nil)
	if err != nil {
		return err, nil
	}
	req.SetBasicAuth(d.Get("wc_key").(string), d.Get("wc_secret").(string))
	resp, err := client.Do(req)
	println(resp)
	if err != nil {
		return err, nil
	}
	return &config, nil
}

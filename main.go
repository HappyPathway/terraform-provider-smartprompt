package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

type providerConfig struct {
	apiURL  string
	timeout int
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   false,
				Description: "The URL of the Smart Prompt API",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30,
				Description: "Timeout in seconds for API requests",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"smartprompt_refined": dataSourceRefinedPrompt(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := providerConfig{
		apiURL:  d.Get("api_url").(string),
		timeout: d.Get("timeout").(int),
	}
	return &config, nil
}
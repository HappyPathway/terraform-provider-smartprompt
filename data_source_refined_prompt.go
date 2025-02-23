package main

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/yourusername/smartprompt-client/client"
)

func dataSourceRefinedPrompt() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRefinedPromptRead,

		Schema: map[string]*schema.Schema{
			"lazy_prompt": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The simple prompt to be refined",
			},
				"domain": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(client.DomainArchitecture),
					string(client.DomainDevelopment),
					string(client.DomainInfrastructure),
					string(client.DomainSecurity),
					string(client.DomainGeneral),
				}, false),
				Description: "Technical domain for the prompt (architecture, development, infrastructure, security, general)",
			},
			"expertise_level": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(client.ExpertiseBeginner),
					string(client.ExpertiseIntermediate),
					string(client.ExpertiseExpert),
				}, false),
				Description: "Target expertise level (beginner, intermediate, expert)",
			},
			"output_format": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(client.OutputFormatSimple),
					string(client.OutputFormatDetailed),
					string(client.OutputFormatTutorial),
					string(client.OutputFormatChecklist),
				}, false),
				Description: "Desired output format (simple, detailed, tutorial, checklist)",
			},
			"include_best_practices": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Include industry best practices in the response",
			},
			"include_examples": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Include examples in the response",
			},
			"refined_prompt": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resulting refined prompt",
			},
			"detected_topics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of technical topics detected in the prompt",
			},
			"recommended_references": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of recommended technical references and documentation",
			},
		},
	}
}

func dataSourceRefinedPromptRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*providerConfig)
	
	c := client.NewClient(config.apiURL, time.Duration(config.timeout)*time.Second)
	
	// Convert schema values to client types
	var domain *client.DomainType
	if v, ok := d.GetOk("domain"); ok {
		d := client.DomainType(v.(string))
		domain = &d
	}

	var expertiseLevel *client.ExpertiseLevel
	if v, ok := d.GetOk("expertise_level"); ok {
		e := client.ExpertiseLevel(v.(string))
		expertiseLevel = &e
	}

	var outputFormat *client.OutputFormat
	if v, ok := d.GetOk("output_format"); ok {
		f := client.OutputFormat(v.(string))
		outputFormat = &f
	}

	var includeBestPractices *bool
	if v, ok := d.GetOk("include_best_practices"); ok {
		b := v.(bool)
		includeBestPractices = &b
	}

	var includeExamples *bool
	if v, ok := d.GetOk("include_examples"); ok {
		e := v.(bool)
		includeExamples = &e
	}

	response, err := c.RefinePromptWithOptions(
		d.Get("lazy_prompt").(string),
		domain,
		expertiseLevel,
		outputFormat,
		includeBestPractices,
		includeExamples,
	)
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("refined_prompt", response.RefinedPrompt)
	d.Set("detected_topics", response.DetectedTopics)
	if response.RecommendedReferences != nil {
		d.Set("recommended_references", response.RecommendedReferences)
	}

	return nil
}
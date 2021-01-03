package guacamole

import (
	"context"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_URL", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_PASSWORD", nil),
			},
			"disable_tls_verification": {
				Type:      schema.TypeBool,
				Optional:  true,
				Sensitive: true,
				Default:   false,
			},
			"disable_cookies": {
				Type:      schema.TypeBool,
				Optional:  true,
				Sensitive: true,
				Default:   false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"guacamole_user": guacamoleUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"guacamole_user": dataSourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := d.Get("url").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	disableTLS := d.Get("disable_tls_verification").(bool)
	disableCookies := d.Get("disable_cookies").(bool)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	config := guac.Config{
		URL:                    url,
		Username:               username,
		Password:               password,
		DisableTLSVerification: disableTLS,
		DisableCookies:         disableCookies,
	}

	err := validate(config)
	if merr, ok := err.(*multierror.Error); ok {
		for e := range merr.Errors {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Errors in guacamole provider configuration",
				Detail:   merr.Errors[e].Error(),
			})
		}
		return nil, diags
	}

	c := guac.New(config)

	err = c.Connect()

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create guacamole client",
			Detail:   "Unable to authenticate user for guacamole client",
		})

		return nil, diags
	}

	return &c, diags
}

// validate validates the config needed to initialize a guacamole client,
// returning a single error with all validation errors, or nil if no error.
func validate(config guac.Config) error {
	var err *multierror.Error
	if config.URL == "" {
		err = multierror.Append(err, fmt.Errorf("URL must be configured for the guacamole provider"))
	}
	if config.Username == "" {
		err = multierror.Append(err, fmt.Errorf("Username must be configured for the guacamole provider"))
	}
	if config.Password == "" {
		err = multierror.Append(err, fmt.Errorf("Password must be configured for the guacamole provider"))
	}
	return err.ErrorOrNil()
}

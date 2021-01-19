package guacamole

import (
	"context"

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
				Type:        schema.TypeBool,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_DISABLE_TLS", false),
			},
			"disable_cookies": {
				Type:        schema.TypeBool,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_DISABLE_COOKIES", false),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"guacamole_user":                  guacamoleUser(),
			"guacamole_user_group":            guacamoleUserGroup(),
			"guacamole_connection_ssh":        guacamoleConnectionSSH(),
			"guacamole_connection_telnet":     guacamoleConnectionTelnet(),
			"guacamole_connection_rdp":        guacamoleConnectionRDP(),
			"guacamole_connection_vnc":        guacamoleConnectionVNC(),
			"guacamole_connection_kubernetes": guacamoleConnectionKubernetes(),
			"guacamole_connection_group":      guacamoleConnectionGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"guacamole_user":                  dataSourceUser(),
			"guacamole_user_group":            dataSourceUserGroup(),
			"guacamole_connection_ssh":        dataSourceConnectionSSH(),
			"guacamole_connection_telnet":     dataSourceConnectionTelnet(),
			"guacamole_connection_rdp":        dataSourceConnectionRDP(),
			"guacamole_connection_vnc":        dataSourceConnectionVNC(),
			"guacamole_connection_kubernetes": dataSourceConnectionKubernetes(),
			"guacamole_connection_group":      dataSourceConnectionGroup(),
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

	// Check for required provider parameters
	check := validate(config)

	if check.HasError() {
		return nil, check
	}

	client := guac.New(config)

	err := client.Connect()

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create guacamole client",
			Detail:   "Unable to authenticate user for guacamole client",
		})

		return nil, diags
	}

	return &client, diags
}

// validate validates the config needed to initialize a guacamole client,
// returning a single error with all validation errors, or nil if no error.
func validate(config guac.Config) diag.Diagnostics {
	var diags diag.Diagnostics

	if config.URL == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing provider parameter",
			Detail:   "URL must be configured for the guacamole provider",
		})
	}
	if config.Username == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing provider parameter",
			Detail:   "Username must be configured for the guacamole provider",
		})
	}
	if config.Password == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing provider parameter",
			Detail:   "Password must be configured for the guacamole provider",
		})
	}
	return diags
}

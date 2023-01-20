package guacamole

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"password"},
				DefaultFunc:  schema.EnvDefaultFunc("GUACAMOLE_USERNAME", nil),
			},
			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"username"},
				AtLeastOneOf: []string{"password", "token"},
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("GUACAMOLE_PASSWORD", nil),
			},
			"token": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"password", "token"},
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("GUACAMOLE_TOKEN", nil),
			},
			"data_source": {
				Type:             schema.TypeString,
				Optional:         true,
				RequiredWith:     []string{"token"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"postgresql", "mysql"}, true)),
				DefaultFunc:      schema.EnvDefaultFunc("GUACAMOLE_DATA_SOURCE", nil),
			},
			"cookies": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
	url := strings.TrimRight(d.Get("url").(string), "/")
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	token := d.Get("token").(string)
	data_source := d.Get("data_source").(string)
	disableTLS := d.Get("disable_tls_verification").(bool)
	disableCookies := d.Get("disable_cookies").(bool)

	cookies := make(map[string]string)
	cookieMap := d.Get("cookies").(map[string]interface{})
	if len(cookieMap) > 0 {
		for k, v := range cookieMap {
			cookies[k] = v.(string)
		}
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	config := guac.Config{
		URL:                    url,
		Username:               username,
		Password:               password,
		Token:                  token,
		DataSource:             data_source,
		Cookies:                cookies,
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
			Detail:   err.Error(),
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
	if config.Password == "" && config.Token == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing provider parameter",
			Detail:   "Either username/password or token/data_source must be configured for the guacamole provider",
		})
	}
	return diags
}

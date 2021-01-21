package guacamole

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGuacamoleConnectionKubernetesBasic(t *testing.T) {
	testProviderConnectionKubernetes := map[string]interface{}{
		"name": "testProviderConnectionKubernetes",
		"attributes": map[string]interface{}{
			"guacd_hostname":           "example.example.com",
			"guacd_port":               "8443",
			"guacd_encryption":         "ssl",
			"failover_only":            true,
			"weight":                   "10",
			"max_connections":          "4",
			"max_connections_per_user": "2",
		},
		"parameters": map[string]interface{}{
			"hostname":                    "hostname.example.com",
			"port":                        "22",
			"ignore_cert":                 true,
			"ca_cert":                     "ca cert",
			"pod":                         "pod",
			"namespace":                   "namespace",
			"container":                   "container",
			"client_cert":                 "client cert",
			"client_key":                  "client key",
			"color_scheme":                "green-black",
			"font_name":                   "Helvetica, sans-serif",
			"font_size":                   "12",
			"max_scrollback_size":         "200",
			"readonly":                    true,
			"backspace":                   "127",
			"typescript_path":             "typescript path",
			"typescript_name":             "typescript name",
			"typescript_auto_create_path": true,
			"recording_path":              "recording path",
			"recording_name":              "recording name",
			"recording_exclude_output":    true,
			"recording_exclude_mouse":     true,
			"recording_include_keys":      true,
			"recording_auto_create_path":  true,
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGuacamoleConnectionKubernetesConfigBasic(toHclString(testProviderConnectionKubernetes, true)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuacamoleConnectionKubernetesExists("guacamole_connection_kubernetes.new"),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "name", testProviderConnectionKubernetes["name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "attributes.0.guacd_hostname", testProviderConnectionKubernetes["attributes"].(map[string]interface{})["guacd_hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "attributes.0.guacd_port", testProviderConnectionKubernetes["attributes"].(map[string]interface{})["guacd_port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "attributes.0.guacd_encryption", testProviderConnectionKubernetes["attributes"].(map[string]interface{})["guacd_encryption"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "attributes.0.failover_only", boolToString(testProviderConnectionKubernetes["attributes"].(map[string]interface{})["failover_only"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "attributes.0.weight", testProviderConnectionKubernetes["attributes"].(map[string]interface{})["weight"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "attributes.0.max_connections", testProviderConnectionKubernetes["attributes"].(map[string]interface{})["max_connections"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "attributes.0.max_connections_per_user", testProviderConnectionKubernetes["attributes"].(map[string]interface{})["max_connections_per_user"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.hostname", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.port", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.ignore_cert", boolToString(testProviderConnectionKubernetes["parameters"].(map[string]interface{})["ignore_cert"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.ca_cert", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["ca_cert"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.pod", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["pod"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.namespace", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["namespace"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.container", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["container"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.client_cert", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["client_cert"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.client_key", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["client_key"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.color_scheme", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["color_scheme"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.font_name", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["font_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.font_size", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["font_size"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.max_scrollback_size", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["max_scrollback_size"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.readonly", boolToString(testProviderConnectionKubernetes["parameters"].(map[string]interface{})["readonly"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.backspace", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["backspace"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.typescript_path", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["typescript_path"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.typescript_name", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["typescript_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.typescript_auto_create_path", boolToString(testProviderConnectionKubernetes["parameters"].(map[string]interface{})["typescript_auto_create_path"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.recording_path", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["recording_path"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.recording_name", testProviderConnectionKubernetes["parameters"].(map[string]interface{})["recording_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.recording_exclude_output", boolToString(testProviderConnectionKubernetes["parameters"].(map[string]interface{})["recording_exclude_output"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.recording_exclude_mouse", boolToString(testProviderConnectionKubernetes["parameters"].(map[string]interface{})["recording_exclude_mouse"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.recording_include_keys", boolToString(testProviderConnectionKubernetes["parameters"].(map[string]interface{})["recording_include_keys"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_kubernetes.new", "parameters.0.recording_auto_create_path", boolToString(testProviderConnectionKubernetes["parameters"].(map[string]interface{})["recording_auto_create_path"].(bool))),
				),
			},
		},
	})
}

func testAccCheckGuacamoleConnectionKubernetesConfigBasic(connection string) string {
	return fmt.Sprintf(`
	resource "guacamole_connection_kubernetes" "new" %s
	`, connection)
}

func testAccCheckGuacamoleConnectionKubernetesExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No connection id set")
		}

		return nil
	}
}

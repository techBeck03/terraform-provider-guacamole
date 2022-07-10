---
page_title: "Connection Kubernetes Resource - terraform-provider-guacamole"
subcategory: ""
description: |-
  The connection_kubernetes resource allows you to configure a guacamole kubernetes connection
---

# Resource `guacamole_connection_kubernetes`

The connection_kubernetes resource allows you to configure a guacamole kubernetes connection

## Example Usage

```terraform
resource "guacamole_connection_kubernetes" "kubernetes" {
  name = "Test K8s Connection"
  parent_identifier = "ROOT"
  attributes {
    guacd_hostname = "guac.test.com"
    guacd_encryption = "ssl"
  }
  parameters {
    hostname = "https://kube.example.com"
    port = 6443
    use_ssl = true
    ignore_cert = true
    namespace = "default"
    pod = "testPod"
    container = "user-container"
    client_cert = <<-EOT
    PLACE CLIENT CERT CONTENTS HERE
    EOT
    client_key = <<-EOT
    PLACE CLIENT Key CONTENTS HERE
    EOT
    color_scheme = "green-black"
    font_size = 48
  }
}
```


## Argument Reference

### Base

- `name` -  (string, Required) Name of the connection
- `parent_identifier` -  (string, Required) Numeric identifier of the parent connection

### Attributes

- `max_connections` - (string) max allowed connections
- `max_connections_per_user` - (string) max allowed connections per user
- `weigth` - (string) connectivity weight
- `failover_only` - (bool) used for failover only
- `guacd_hostname` - (string) guacamole proxy hostname
- `guacd_port` - (string) guacamole proxy port
- `guacd_encryption` - (string) guacamole proxy encryption type:  Value should be on of:
  - `none`
  - `ssl`

### Parameters

#### *Network*
- `hostname` - (string) hostname
- `port` - (string) port
- `public_host_key` - (string) public host key
#### Container
- `namespace` - (string)
- `pod` - (string)
- `container` - (string)
#### *Authentication*
- `client_certificate` - (string) client certificate
- `client_key` - (string) client key
#### *Display*
- `color_scheme` - (string) color scheme: Value should be on of:
  - `black-white`
  - `gray-black`
  - `green-black`
  - `white-black`
- `font_name` - (string) font family name
- `font_size` - (string) font size. Value should be on of:
  - `8`
  - `9`
  - `10`
  - `11`
  - `12`
  - `14`
  - `18`
  - `24`
  - `30`
  - `36`
  - `48`
  - `60`
  - `72`
  - `96`
- `max_scrollback_size` - (string) max scrollback size
- `readonly` - (bool) display is read-only
#### *Terminal Behavior*
- `backspace` - (string) backspace key sends.  Value should be on of:
  - `127`
  - `8`
#### *Typescript (Text Session Recording)*
- `typescript_path` - (string) typescript path
- `typescript_name` - (string) typescript name
- `typescript_auto_create_path` - (bool) automatically create recording path
#### *Screen Recording*
- `recording_path` - (string) recording path
- `recording_name` - (string) recording name
- `recording_exclude_output` - (bool) exclude graphics/streams
- `recording_exclude_mouse` - (bool) exclude mouse
- `recording_include_keys` - (bool) include key events
- `recording_auto_create_path` - (bool) automatically create recording path

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

#### Base
- `identifier` -  (string) Numeric identifier of the kubernetes connection
- `protocol` -  (string) protocol of the connection (`kubernetes`)
- `active_connections` - (sting) Number of active connections for the group

## Import

Kubernetes connection can be imported using the `resource id`, e.g.

```shell
terraform import guacamole_connection_kubernetes.kubernetes 1
```

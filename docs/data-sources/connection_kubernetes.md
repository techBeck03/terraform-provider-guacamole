---
page_title: "Connection Kubernetes Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The connection_kubernetes data source allows you to retrieve a guacamole kubernetes connection details by identifier or path
---

# Data Source `guacamole_connection_kubernetes`

The connection_kubernetes data source allows you to retrieve a guacamole kubernetes connection details by identifier or path

## Example Usage

```terraform
data "guacamole_connection_kubernetes" "kubernetes" {
  identifier = 1234
}
```

```terraform
data "guacamole_connection_kubernetes" "kubernetes" {
  path = "parentGroupName/connectionName"
}
```

## Attributes Reference

The following attributes are exported.

### Base

- `name` -  (string) Name of the connection
- `path` -  (string) Used in place of identifier to find a path by "ParentName/TargetName" when the identifier is unknown
- `identifier` -  (string) Numeric identifier of the kubernetes connection
- `parent_identifier` -  (string) Numeric identifier of the parent connection
- `protocol` -  (string) protocol of the connection (`kubernetes`).
- `active_connections` - (sting) Number of active connections for the group


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

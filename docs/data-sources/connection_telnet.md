---
page_title: "Connection Telnet Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The connection_telnet data source allows you to retrieve a guacamole telnet connection details by identifier or path
---

# Data Source `guacamole_connection_telnet`

The user group data source allows you to retrieve a guacamole user group by identifier

## Example Usage

```terraform
data "guacamole_connection_telnet" "telnet" {
  identifier = 1234
}
```

```terraform
data "guacamole_connection_telnet" "telnet" {
  path = "parentGroupName/connectionName"
}
```

## Attributes Reference

The following attributes are exported.

### Base

- `name` -  (string) Name of the connection
- `path` -  (string) Used in place of identifier to find a path by "ParentName/TargetName" when the identifier is unknown
- `identifier` -  (string) Numeric identifier of the telnet connection
- `parent_identifier` -  (string) Numeric identifier of the parent connection
- `protocol` -  (string) protocol of the connection (`telnet`).
- `active_connections` - (sting) Number of active connections for the group


### Attributes

- `max_connections` - (string) max allowed connections
- `max_connections_per_user` - (string) max allowed connections per user
- `weigth` - (string) connectivity weight
- `failover_only` - (bool) used for failover only
- `guacd_hostnmae` - (string) guacamole proxy hostname
- `guacd_port` - (string) guacamole proxy port
- `guacd_encryption` - (string) guacamole proxy encryption type:  Value should be on of:
  - `none`
  - `ssl`

### Parameters

#### *Network*
- `hostname` - (string) hostname
- `port` - (string) port
#### *Authentication*
- `username` - (string) username
- `username_regex` - (string) username regular expression
- `password_regex` - (string) password regular expression
- `login_success_regex` - (string) login success regular expression
- `login_failure_regex` - (string) login failure regular expression
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
#### *Clipboard*
- `disable_copy` - (bool) disable copying from the terminal
- `disable_paste` - (bool) disable pastiong from client
#### *Terminal Behavior*
- `backspace` - (string) backspace key sends.  Value should be on of:
  - `127`
  - `8`
- `terminal_type` - (string) terminal type. Value should be one of:
  - `ansi`
  - `linux`
  - `vt100`
  - `vt220`
  - `xterm`
  - `xterm-25color`
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
#### *Wake-on-LAN (WoL)*
- `wol_send_packet` - (bool) send WoL packet
- `wol_mac_address` - (string) MAC address of the remote host
- `wol_broadcast_address` - (string) broadcast address for WoL packet
- `wol_boot_wait_time` - (string) host boot wait time

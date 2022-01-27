---
page_title: "Connection VNC Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The connection_vnc data source allows you to retrieve a guacamole vnc connection details by identifier or path
---

# Data Source `guacamole_connection_vnc`

The connection_vnc data source allows you to retrieve a guacamole vnc connection details by identifier or path

## Example Usage

```terraform
data "guacamole_connection_vnc" "vnc" {
  identifier = 1234
}
```

```terraform
data "guacamole_connection_vnc" "vnc" {
  path = "parentGroupName/connectionName"
}
```

## Attributes Reference

The following attributes are exported.

### Base

- `name` -  (string) Name of the connection
- `path` -  (string) Used in place of identifier to find a path by "ParentName/TargetName" when the identifier is unknown
- `identifier` -  (string) Numeric identifier of the vnc connection
- `parent_identifier` -  (string) Numeric identifier of the parent connection
- `protocol` -  (string) protocol of the connection (`vnc`).
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
#### *Authentication*
- `username` - (string) username
#### *Display*
- `readonly` - (bool) display is read-only
- `swap_red_blue` - (bool) swap red/blue components
- `cursor` - (bool) cursor.  Value should be on of:
  - `local`
  - `remote`
- `color_depth` - (string) color depth.  Value should be on of:
  - `8`
  - `16`
  - `24`
  - `32`
#### *Clipboard*
- `disable_copy` - (bool) disable copying from the terminal
- `disable_paste` - (bool) disable pastiong from client
#### VNC Repeater
- `destination_host` - (string) destination host
- `destination_port` - (string) destination port
#### *Screen Recording*
- `recording_path` - (string) recording path
- `recording_name` - (string) recording name
- `recording_exclude_output` - (bool) exclude graphics/streams
- `recording_exclude_mouse` - (bool) exclude mouse
- `recording_include_keys` - (bool) include key events
- `recording_auto_create_path` - (bool) automatically create recording path
#### *SFTP*
- `sftp_enable` - (bool) enable SFTP
- `sftp_hostname` - (string) hostname
- `sftp_port` - (string) port
- `sftp_host_key` - (string) public host key (Base64)
- `sftp_username` - (string) username
- `sftp_password` - (string) password
- `sftp_private_key` - (string) private key
- `sftp_passphrase` - (string) passphrase
- `sftp_root_directory` - (string) file browser root directory
- `sftp_upload_directory` - (string) default upload directory
- `sftp_keepalive_interval` - (string) SFTP keepalive interval
- `sftp_dsiable_file_download` - (bool) disable file download
- `sftp_disable_file_upload` - (bool) disable file upload
#### Audio
- `enable_audio` - (bool) enable audio
- `audio_server_name` - (string) audio server name
#### *Wake-on-LAN (WoL)*
- `wol_send_packet` - (bool) send WoL packet
- `wol_mac_address` - (string) MAC address of the remote host
- `wol_broadcast_address` - (string) broadcast address for WoL packet
- `wol_boot_wait_time` - (string) host boot wait time

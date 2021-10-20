---
page_title: "Connection RDP Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The connection_rdp data source allows you to retrieve a guacamole rdp connection details by identifier or path
---

# Data Source `guacamole_connection_rdp`

The connection_rdp data source allows you to retrieve a guacamole rdp connection details by identifier or path

## Example Usage

```terraform
data "guacamole_connection_rdp" "rdp" {
  identifier = 1234
}
```

```terraform
data "guacamole_connection_rdp" "rdp" {
  path = "parentGroupName/connectionName"
}
```

## Attributes Reference

The following attributes are exported.

### Base

- `name` -  (string) Name of the connection
- `path` -  (string) Used in place of identifier to find a path by "ParentName/TargetName" when the identifier is unknown
- `identifier` -  (string) Numeric identifier of the rdp connection
- `parent_identifier` -  (string) Numeric identifier of the parent connection
- `protocol` -  (string) protocol of the connection (`rdp`).
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
- `domain` - (string) active directory domain name
- `security_mode` - (string) security mode.  Value should be on of:
  - `any`
  - `nla`
  - `rdp`
  - `tls`
  - `vmconnect`
- `disable_authentication` - (bool) disable authentication
- `ignore_cert` - (bool) ignore server certificate
#### *Remote Desktop Gateway*
- `gateway_hostname` - (string) remote desktop gateway hostname
- `gateway_port` - (string) remote desktop gateway port
- `gateway_username` - (string) remote desktop gateway username
- `gateway_password` - (string) remote desktop gateway password
- `gateway_domain` - (string) remote desktop gateway domain name
- `initial_program` - (string) initial program
- `client_name` - (string) client name
- `keyboard_layout` - (string) keyboard layout.  Value should be on of:
  - `da-dk-qwerty`
  - `de-ch-qwertz`
  - `de-de-qwertz`
  - `en-gb-qwerty`
  - `en-us-qwerty`
  - `es-es-qwerty`
  - `es-latam-qwerty`
  - `failsafe`
  - `fr-be-azerty`
  - `fr-ch-qwertz`
  - `fr-fr-azerty`
  - `hu-hu-qwertz`
  - `it-it-qwerty`
  - `ja-jp-qwerty`
  - `pt-br-qwerty`
  - `sv-se-qwerty`
  - `tr-tr-qwerty`
- `timezone` - (string) timezone string. Example `America/Chicago`
- `administrator_console` - (bool) administrator console
#### *Display*
- `width` - (string) display width
- `height` - (string) display height
- `dpi` - (string) resolution (DPI)
- `color_depth` - (string) color depth.  Value should be on of:
  - `8`
  - `16`
  - `24`
  - `32`
- `resize_method` - (string) display resize method.  Value should be on of:
  - `display-update`
  - `reconnect`
- `readonly` - (bool) display is read-only
#### *Clipboard*
- `disable_copy` - (bool) disable copying from the terminal
- `disable_paste` - (bool) disable pastiong from client
#### *Device Redirection*
- `console_audio` - (bool) support audio in console
- `disable_audio` - (bool) disable audio
- `enable_audio_input` - (bool) enable audio input (microphone)
- `enable_printing` - (bool) enable printing
- `printer_name` - (string) redirected printer name
- `enable_drive` - (bool) enable drive
- `drive_name` - (string) drive name
- `disable_file_download` - (bool) disable file download
- `diable_file_upload` - (bool) disable file upload
- `drive_path` - (string) drive path
- `create_drive_path` - (bool) automatically create drive
- `static_channels` - (string) static channel names
#### *Performance*
- `enable_wallpaper` - (bool) enable wallpaper
- `enable_theming` - (bool) enable theming
- `enable_font_smoothing` - (bool) enable font smoothing (ClearType)
- `enable_full_window_drag` - (bool) enable full-window drag
- `enable_desktop_composition` - (bool) enable desktop composition (Aero)
- `enable_menu_animations` - (bool) enable menu animations
- `disable_bitmap_caching` - (bool) disable bitmap caching
- `disable_offscreen_caching` - (bool) diable off-screen caching
- `disable_glyph_caching` - (bool) disable glyph caching
#### *RemoteApp*
- `remote_app` - (string) program
- `remote_app_working_directory` - (string) working directory
- `remote_app_parameters` - (string) parameters
#### *Preconnection PDU/Hyper-V*
- `preconnection_id` - (string) RDP source ID
- `preconnection_blob` - (string) Preconnection BLOB (VM ID)
#### *Load Balancing*
- `load_balance_info` - (string) load balance info/cookie
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
#### *Wake-on-LAN (WOL)*
- `wol_send_packet` - (bool) send WoL packet
- `wol_mac_address` - (string) MAC address of the remote host
- `wol_broadcast_address` - (string) broadcast address for WoL packet
- `wol_boot_wait_time` - (string) host boot wait time
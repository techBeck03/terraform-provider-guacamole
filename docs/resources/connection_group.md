---
page_title: "Connection Group Resource - terraform-provider-guacamole"
subcategory: ""
description: |-
  The connection group resource allows you to configure a connection group
---

# Resource `guacamole_connection_group`

The user group data source allows you to retrieve a guacamole user group by identifier

## Example Usage

```terraform
resource "guacamole_connection_group" "group" {
  parent_identifier = "ROOT"
  name = "Testing Group"
  type = "organizational"
  attributes {
    max_connections_per_user = 4
  }
}
```

## Argument Reference

### Base

- `name` -  (string, Required) name of the connection group
- `parent_identifier` -  (string, Required) numeric identifier of the parent connection group
- `type` -  (string) type of connection group.  Value should be on of:
  - `ORGANIZATIONAL`
  - `BALANCING`

### Attributes

- `enable_session_affinity` - (bool) whether session affinity is enabled
- `max_connections` - (string) max allowed connections for the group
- `max_connections_per_user` - (string) max allowed connections per user for the group

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

#### Base
- `identifier` -  (string) numeric identifier of the connection group
- `active_connections` - (sting) number of active connections for the group
- `member_connections` - (List) list of connection identifiers whose parent is this connection group
- `member_connection_groups` - (List) list of connection group identifiers whose parent is this user group

## Import

Connection group can be imported using the `resource id`, e.g.

```shell
terraform import guacamole_connection_group.group 1
```
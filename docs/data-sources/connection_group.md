---
page_title: "Connection Group Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The connection group data source allows you to retrieve connection group details by identifier or path
---

# Data Source `connection_group`

The user group data source allows you to retrieve a guacamole user group by identifier

## Example Usage

```terraform
data "guacamole_connection_group" "group" {
  identifier = 1234
}
```

```terraform
data "guacamole_connection_group" "group" {
  path = "parentGroupName/targetGroupName"
}
```

## Attributes Reference

The following attributes are exported.

### User

- `name` -  (string) Name of the 
- `path` -  (string) Used in place of identifier to find a path by "ParentName/TargetName" when the identifier is unknown
- `identifier` -  (string) Numeric identifier of the connection group
- `parent_identifier` -  (string) Numeric identifier of the parent connection group
- `type` -  (string) type of connection group.  Valid choices are:
  - `ORGANIZATIONAL`
  - `BALANCING`
- `active_connections` - (sting) Number of active connections for the group
- `member_connections` - (List) List of connection identifiers whose parent is this connection group
- `member_connection_groups` - (List) List of connection group identifiers whose parent is this user group

### Attributes

- `enable_session_affinity` - (bool) 
- `max_connections` - (string) max allowed connections for the group
- `max_connections_per_user` - (string) whether session affinity is enabled


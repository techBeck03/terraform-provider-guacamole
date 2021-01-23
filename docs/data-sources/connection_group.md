---
page_title: "Connection Group Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The connection group data source allows you to retrieve connection group details by identifier or path
---

# Data Source `guacamole_connection_group`

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

- `name` -  (string) name of the connection group
- `path` -  (string) used in place of identifier to find a path by "ParentName/TargetName" when the identifier is unknown
- `identifier` -  (string) numeric identifier of the connection group
- `parent_identifier` -  (string) numeric identifier of the parent connection group
- `type` -  (string) type of connection group.  Value should be on of:
  - `ORGANIZATIONAL`
  - `BALANCING`
- `active_connections` - (sting) number of active connections for the group
- `member_connections` - (List) list of connection identifiers whose parent is this connection group
- `member_connection_groups` - (List) list of connection group identifiers whose parent is this user group

### Attributes

- `enable_session_affinity` - (bool) whether session affinity is enabled
- `max_connections` - (string) max allowed connections for the group
- `max_connections_per_user` - (string) max allowed connections per user for the group

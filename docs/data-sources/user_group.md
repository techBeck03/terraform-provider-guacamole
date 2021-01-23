---
page_title: "User Group Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The user group data source allows you to retrieve a guacamole user group by identifier
---

# Data Source `guacamole_user_group`

The user group data source allows you to retrieve a guacamole user group by identifier

## Example Usage

```terraform
data "guacamole_user_group" "group" {
  identifier = "testGuacamoleUserGroup"
}

```

## Attributes Reference

The following attributes are exported.

### User

- `identifier` -  (string) the guacamole user group identifier
- `group_membership` - (List) list of user group identifiers that this group is a member of
- `system_permissions` - (List) list of system permissions assigned to the user
- `member_users` - (List) user identifiers that are members of this group
- `member_groups` - (List) user group identifiers that are members of this group
- `connections` - list of connection identifiers assigned to the user group.  This list currently does not include connection identifiers from parent user groups.
- `connection_groups` - (List) list of connection group identifiers assigned to the user group.  This list currently does not include connection group identifiers from parent user groups.

### Attributes

- `disabled` - (bool) whether the user group is disabled


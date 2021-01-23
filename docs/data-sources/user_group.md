---
page_title: "User Group Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The user group data source allows you to retrieve a guacamole user group by identifier
---

# Data Source `user_group`

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

- `identifier` -  (string) The guacamole user group identifier
- `group_membership` - (List) List of user group identifiers that this group is a member of
- `system_permissions` - (List) List of system permissions assigned to the user
- `member_users` - (List) User identifiers that are members of this group
- `member_groups` - (List) User group identifiers that are members of this group
- `connections` - List of connection identifiers assigned to the user group.  This list currently does not include connection identifiers from parent user groups.
- `connection_groups` - (List) List of connection group identifiers assigned to the user group.  This list currently does not include connection group identifiers from parent user groups.

### Attributes

- `disabled` - (bool) whether the user group is disabled


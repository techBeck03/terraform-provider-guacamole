---
page_title: "User Group Resource - terraform-provider-guacamole"
subcategory: ""
description: |-
  The user group resource allows you to  a guacamole user group by identifier
---

# Resource `guacamole_user_group`

The user group data source allows you to configure a guacamole user group

## Example Usage

```terraform
resource "guacamole_user_group" "group" {
  identifier = "testGuacamoleUserGroup"
  system_permissions = ["ADMINISTER", "CREATE_USER"]
  group_membership = ["Parent Group"]
  connections = [
    "12345"
  ]
  connection_groups = [
    "678910"
  ]
  attributes {
    disabled = true
  }
}

```

## Argument Reference

### Base

- `identifier` -  (string, Required) the guacamole user group identifier
- `group_membership` - (List) list of user group identifiers that this group is a member of
- `system_permissions` - (List) list of system permissions assigned to the user
- `member_groups` - (List) user group identifiers that are members of this group
- `connections` - list of connection identifiers assigned to the user group.  This list currently does not include connection identifiers from parent user groups.
- `connection_groups` - (List) list of connection group identifiers assigned to the user group.  This list currently does not include connection group identifiers from parent user groups.

### Attributes

- `disabled` - (bool) whether the user group is disabled

## Import

User group can be imported using the `resource id`, e.g.

```shell
terraform import guacamole_user_group.group group_name
```

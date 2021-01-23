---
page_title: "User Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The user data source allows you to retrieve a guacamole user by username
---

# Data Source `guacamole_user`

The user data source allows you to retrieve a guacamole user by username

## Example Usage

```terraform
data "guacamole_user" "user" {
  username = "testGuacamoleUser"
}

```

## Attributes Reference

The following attributes are exported.

### User

- `username` -  (string) the guacamole user username
- `last_active` - (string) timestamp of last activity
- `group_membership` - (List) list of user group identifiers
- `system_permissions` - (List) list of system permissions assigned to the user
- `connections` - (List) list of connection identifiers assigned to the user.  This list currently does not include connection identifiers from parent user groups.
- `connection_groups` - (List) list of connection group identifiers assigned to the user.  This list currently does not include connection group identifiers from parent user groups.

### Attributes

- `organizational_role` - (string) assigned organizational role
- `full_name` - (string) full name
- `email` - (string) email address
- `expired` - (bool) whether the account is expired
- `timezone` - (string) the timezone string ("America/Chicago")
- `access_window_start` - (string) access window start time
- `access_window_end` - (string) access window end time
- `disabled` - (bool) whether the account is disabled
- `valid_from` - (string) account valid start date
- `valid_until` - (string) account valid end date

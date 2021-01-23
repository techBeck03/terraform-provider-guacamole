---
page_title: "User Data Source - terraform-provider-guacamole"
subcategory: ""
description: |-
  The user data source allows you to retrieve a guacamole user by username
---

# Data Source `user`

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

- `username` -  (string) The guacamole user username
- `last_active` - (string) Timestamp of last activity
- `group_membership` - (List) List of user group identifiers
- `system_permissions` - (List) List of system permissions assigned to the user
- `connections` - List of connection identifiers assigned to the user.  This list currently does not include connection identifiers from parent user groups.
- `connection_groups` - (List) List of connection group identifiers assigned to the user.  This list currently does not include connection group identifiers from parent user groups.

### Attributes

- `organizational_role` - (string) Assigned organizational role
- `full_name` - (string) Full name
- `email` - (string) Email address
- `expired` - (bool) whether the account is expired
- `timezone` - (string) The timezone string ("America/Chicago")
- `access_window_start` - (string) Access window start time
- `access_window_end` - (string) Access window end time
- `disabled` - (bool) whether the account is disabled
- `valid_from` - (string) Account valid start date
- `valid_until` - (string) Account valid end date

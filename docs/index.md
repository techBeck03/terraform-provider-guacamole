---
page_title: "Provider: Guacamole"
subcategory: ""
description: |-
  Terraform provider for interacting with Apache Guacamole web API
---

# Guacamole Provider


The guacamole provider is used to interact with the [Apache Guacamole](https://guacamole.apache.org/) web API.  The provider includes the following capabilities

1. Data sources to read users, user groups, connections for ssh, telnet, vnc, rdp, and kubernetes, and connection groups.
2. Resources to perform CRUD operations on users, user groups, connections for ssh, telnet, vnc, rdp, and kubernetes, and connection groups.

Provider versioning will align with Guacamole's release versioning starting with release `1.2`

## Example Usage

To use the provider you will need the Guacamole web server url, username, and password as shown in the example below

```terraform
provider "guacamole" {
  url      = "https://guacamole.example.com"
  username = "guacadmin"
  password = "guacadmin"
  disable_tls_verification = true
}
```

## Schema

- **url** (String) URL of guacamole web server (defaults to env `GUACAMOLE_URL`)
- **username** (String) Username to authenticate to guacamole (defaults to `GUACAMOLE_USERNAME`)
- **password** (String, Optional) Password to authenticate to guacamole (defaults to `GUACAMOLE_PASSWORD`)
- **disable_tls_verification** (Bool, Optional) Whether to disable tls verification for ssl connections (defaults to `false`)
- **disable_cookies** (Bool, Optional) Whether to disable cookie collection in session (defaults to `false`)

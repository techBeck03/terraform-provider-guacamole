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

- **url** (String) URL of guacamole web server (defaults to environment variable `GUACAMOLE_URL`)
- **username** (String) Username to authenticate to guacamole (defaults to environment variable `GUACAMOLE_USERNAME`)
- **password** (String) Password to authenticate to guacamole (defaults to environment variable `GUACAMOLE_PASSWORD`)
- **disable_tls_verification** (Bool, Optional) Whether to disable tls verification for ssl connections (defaults to `false`)
- **disable_cookies** (Bool, Optional) Whether to disable cookie collection in session (defaults to `false`)

## Using Guacamole Parameter Tokens

Apache Guacamole allows users to use system generated [parmater tokens](https://guacamole.apache.org/doc/gug/configuring-guacamole.html#parameter-tokens) within connection definitions.  The parameter token syntax is the same syntax used for HCL string interpolation of variables and must therefore be escaped.

Below is an example of how to use the `${GUAC_USERNAME}` and `${GUAC_PASSWORD}` parameter tokens

```terraform
resource "guacamole_connection_ssh" "ssh" {
  name = "sshConnection"
  parent_identifier = "ROOT"
  parameters {
    hostname = "testing.example.com"
    username = "$${GUAC_USERNAME}"
    password ="$${GUAC_PASSWORD}"
  }
}
```
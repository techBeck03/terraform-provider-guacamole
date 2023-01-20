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

Alternatively, if you are using dual factor authentication, you will need the Guacamole web server url, token, and data_source as shown in the example below.  One way to retrieve the token is to login via the UI using DFA then open the developer console of your browser.  Inspect any Fetch/XHR operation and you'll find the token in the request headers under `guacamole-token` (this assumes you are using guacamole `1.4` or later)

```terraform
provider "guacamole" {
  url         = "https://guacamole.example.com"
  token       = "8675309"
  data_source = "mysql"
  disable_tls_verification = true
}
```

One more example including optional cookies

```terraform
provider "guacamole" {
  url         = "https://guacamole.example.com"
  token       = "8675309"
  data_source = "mysql"
  cookies     = {
    SERVERID = "field-guacp1-3"
  }
  disable_tls_verification = true
}
```

## Schema

- **url** (String) URL of guacamole web server (defaults to environment variable `GUACAMOLE_URL`)
- **username** (String) Username to authenticate to guacamole (defaults to environment variable `GUACAMOLE_USERNAME`)
- **password** (String) Password to authenticate to guacamole (defaults to environment variable `GUACAMOLE_PASSWORD`). This parameter is mutually exclusive to `token`
- **token** (String) Token to authenticate to guacamole (defaults to environment variable `GUACAMOLE_TOKEN`).  This parameter is mutually exclusive to `password`
- **data_source** (String) Datasource for guacamole configuration data (defaults to environment variable `GUACAMOLE_DATA_SOURCE`).  This parameter is required for token based authentication.  Values must be one of:
  - `mysql`
  - `postgresql`
- **cookies** (Map[string], Optional) Map of cookies to be included with requests if using `token` based authentication.  This parameter helps support cookie based load balancing use cases coupled with dual factor authentication.
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
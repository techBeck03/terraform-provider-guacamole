# Go Apache Guacamole API client

This go-based client is intended to be used as an sdk for the [Apache Guacamole](https://guacamole.apache.org/) web API.  To my knowledge this API is undocumented and could therefore change from release to release.

The intent of this sdk is to be used by a Terraform provider that is still a work in progress.  Once completed I will link the community provider here.

Be sure to checkout the release that matches the version of Guacamole you are running.  The developement of this sdk started with the Guacamole 1.2.0 release so no prior releases are currently supported.

# Usage

## Connect/Discconect


```go
package main

import (
	"fmt"

	guac "github.com/techBeck03/guacamole-api-client"
	"github.com/techBeck03/guacamole-api-client/types"
)

func main() {
	client := guac.New(guac.Config{
		URL:                    "https://guac.example.com",
		Username:               "guacadmin",
		Password:               "guacadmin",
		DisableTLSVerification: true,
	})

    err := client.Connect()

    if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection successful")
    }
    
    err = client.Disconnect()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Disconnect successful")
    }
}
```

## Testing

1) Copy/Rename the example env file and change values to match your envrionment <br>
   ```bash
    cp .env_example .env
    ```
2) Run tests
    ```bash
    make test
    ```

## Examples
Check the examples folder for more detailed usage examples
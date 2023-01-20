package guacamole

import (
	"encoding/json"
	"log"
)

func prettyPrint(object interface{}) {
	output, _ := json.MarshalIndent(object, "", "    ")
	log.Printf("%s", string(output))
}

package guacamole

import "strconv"

func stringToBool(v string) bool {
	if v == "" {
		return false
	}
	b, err := strconv.ParseBool(v)

	if err != nil {
		return false
	}
	return b
}

func boolToString(b bool) string {
	if b == true {
		return "true"
	}
	return ""
}

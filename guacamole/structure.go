package guacamole

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

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

func sliceDiff(slice1 []string, slice2 []string, bidirectional bool) []string {
	var diff []string

	var loopCount int
	if bidirectional {
		loopCount = 2
	} else {
		loopCount = 1
	}

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < loopCount; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func stringInSlice(valid []string, test []string) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, t := range test {
		matchFlag := false
		for _, v := range valid {
			if v == t {
				matchFlag = true
				break
			}
		}
		if !matchFlag {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Invalid value entered"),
				Detail:   fmt.Sprintf("%s is not one of supported values: %s", t, strings.Join(valid[:], ", ")),
			})
		}
	}
	return diags
}

func checkForDuplicates(slice1 []string) diag.Diagnostics {
	var diags diag.Diagnostics
	var check []string
	var duplicates []string

	for _, v1 := range slice1 {
		matchFlag := false
		for _, v2 := range check {
			if v2 == v1 {
				matchFlag = true
				duplicates = append(duplicates, v2)
			}
		}
		if !matchFlag {
			check = append(check, v1)
		}
	}
	if len(duplicates) > 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Duplicate entries found in array"),
			Detail:   fmt.Sprintf("Found the duplicate entries: %s", strings.Join(duplicates[:], ", ")),
		})
		return diags
	}
	return diags
}

// sorts slice 2 by slice 1
func sortSliceBySlice(slice1 []string, slice2 []string) []string {
	var sorted []string
	for _, v1 := range slice1 {
		for _, v2 := range slice2 {
			if v1 == v2 {
				sorted = append(sorted, v1)
				break
			}
		}
	}
	for _, v1 := range slice2 {
		matchFlag := false
		for _, v2 := range sorted {
			if v1 == v2 {
				matchFlag = true
				break
			}
		}
		if !matchFlag {
			sorted = append(sorted, v1)
		}
	}
	return sorted
}

func toHclString(value interface{}, isNested bool) string {
	// Ideally, we'd use a type switch here to identify slices and maps, but we can't do that, because Go doesn't
	// support generics, and the type switch only matches concrete types. So we could match []interface{}, but if
	// a user passes in []string{}, that would NOT match (the same logic applies to maps). Therefore, we have to
	// use reflection and manually convert into []interface{} and map[string]interface{}.

	if slice, isSlice := tryToConvertToGenericSlice(value); isSlice {
		return sliceToHclString(slice)
	} else if m, isMap := tryToConvertToGenericMap(value); isMap {
		return mapToHclString(m)
	} else {
		return primitiveToHclString(value, isNested)
	}
}

// Try to convert the given value to a generic slice. Return the slice and true if the underlying value itself was a
// slice and an empty slice and false if it wasn't. This is necessary because Go is a shitty language that doesn't
// have generics, nor useful utility methods built-in. For more info, see: http://stackoverflow.com/a/12754757/483528
func tryToConvertToGenericSlice(value interface{}) ([]interface{}, bool) {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() != reflect.Slice {
		return []interface{}{}, false
	}

	genericSlice := make([]interface{}, reflectValue.Len())

	for i := 0; i < reflectValue.Len(); i++ {
		genericSlice[i] = reflectValue.Index(i).Interface()
	}

	return genericSlice, true
}

// Try to convert the given value to a generic map. Return the map and true if the underlying value itself was a
// map and an empty map and false if it wasn't. This is necessary because Go is a shitty language that doesn't
// have generics, nor useful utility methods built-in. For more info, see: http://stackoverflow.com/a/12754757/483528
func tryToConvertToGenericMap(value interface{}) (map[string]interface{}, bool) {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() != reflect.Map {
		return map[string]interface{}{}, false
	}

	reflectType := reflect.TypeOf(value)
	if reflectType.Key().Kind() != reflect.String {
		return map[string]interface{}{}, false
	}

	genericMap := make(map[string]interface{}, reflectValue.Len())

	mapKeys := reflectValue.MapKeys()
	for _, key := range mapKeys {
		genericMap[key.String()] = reflectValue.MapIndex(key).Interface()
	}

	return genericMap, true
}

// Convert a slice to an HCL string. See ToHclString for details.
func sliceToHclString(slice []interface{}) string {
	hclValues := []string{}

	for _, value := range slice {
		hclValue := toHclString(value, true)
		hclValues = append(hclValues, hclValue)
	}

	return fmt.Sprintf("[%s]", strings.Join(hclValues, ", "))
}

// Convert a map to an HCL string. See ToHclString for details.
func mapToHclString(m map[string]interface{}) string {
	keyValuePairs := []string{}

	for key, value := range m {
		var keyValuePair string
		if _, isMap := tryToConvertToGenericMap(value); isMap {
			keyValuePair = fmt.Sprintf(`%s %s`, key, toHclString(value, true))
		} else {
			keyValuePair = fmt.Sprintf(`%s = %s`, key, toHclString(value, true))
		}
		keyValuePairs = append(keyValuePairs, keyValuePair)
	}

	return fmt.Sprintf("{\n%s\n}", strings.Join(keyValuePairs, "\n"))
}

// Convert a primitive, such as a bool, int, or string, to an HCL string. If this isn't a primitive, force its value
// using Sprintf. See ToHclString for details.
func primitiveToHclString(value interface{}, isNested bool) string {
	if value == nil {
		return "null"
	}

	switch v := value.(type) {

	case bool:
		return strconv.FormatBool(v)

	case string:
		// If string is nested in a larger data structure (e.g. list of string, map of string), ensure value is quoted
		if isNested {
			return fmt.Sprintf("\"%v\"", v)
		}

		return fmt.Sprintf("%v", v)

	default:
		return fmt.Sprintf("%v", v)
	}
}

func validateTimestring(timeString string, name string) diag.Diagnostics {
	var diags diag.Diagnostics
	regex := `^\d{4}[-]\d{2}[-]\d{2}$`
	matched, err := regexp.MatchString(regex, timeString)
	if err != nil || !matched {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Invalid timestring format for: %s", name),
			Detail:   "Date string must be in the form of YYYY-DD-MM",
		})
	}
	return diags
}

func testAccCheckTestSliceVals(resourceName string, key string, expected []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		v, ok := rs.Primary.Attributes[fmt.Sprintf("%s.#", key)]
		if !ok {
			return fmt.Errorf("%s: Attribute '%s.#' not found", resourceName, key)
		}
		testCount, _ := strconv.Atoi(v)
		if testCount == 0 {
			return fmt.Errorf("No entries found in state for key: %s", key)
		}

		var sv []string
		for i := 0; i < testCount; i++ {
			sv = append(sv, rs.Primary.Attributes[fmt.Sprintf("%s.%d", key, i)])
		}

		diff := sliceDiff(expected, sv, true)
		if len(diff) > 0 {
			return fmt.Errorf("Set values: %s do not match expected value: %s", sv, expected)
		}
		return nil
	}
}

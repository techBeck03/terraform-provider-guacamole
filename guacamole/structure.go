package guacamole

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func complexStringInSlice(valid []string, test []string) diag.Diagnostics {
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

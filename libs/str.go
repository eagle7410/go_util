package lib

import "strings"

// Example use v := "Fist" => "fist"
func FirstToUpper (val *string) string {
	str := *val
	return strings.ToUpper(str[:1]) + str[1:]
}

// Example use v := "fist" => "Fist"
func FirstToLower (val *string) string {
	str := *val
	return strings.ToLower(str[:1]) + str[1:]
}

func IsHasStringInArray(find *string, arr *[]string) bool {
	for _, iter := range *arr {
		if *find == iter {
			return true
		}
	}

	return false
}

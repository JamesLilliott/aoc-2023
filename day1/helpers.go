package main

func contains(array []string, search string) bool {
	for _, val := range array {
		if val == search {
			return true
		}
	}

	return false
}

func reverseString(str string) string {
	result := ""
	for _, val := range str {
		result = string(val) + result
	}

	return result
}

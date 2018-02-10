package filter

// Less compares a and b and returns true if a is less than b.
func Less(a interface{}, b interface{}) bool {
	switch v := a.(type) {
	case float32:
		return v < b.(float32)
	case float64:
		return v < b.(float64)
	case int:
		return v < b.(int)
	case string:
		return v < b.(string)
	}
	return false
}

// Equals compares a and b and returns true if a equals b.
func Equals(a interface{}, b interface{}) bool {
	switch v := a.(type) {
	case float32:
		return v == b.(float32)
	case float64:
		return v == b.(float64)
	case int:
		return v == b.(int)
	case string:
		return v == b.(string)
	}
	return false
}

// Greater compares a and b and returns true if a is greater than b.
func Greater(a interface{}, b interface{}) bool {
	switch v := a.(type) {
	case float32:
		return v > b.(float32)
	case float64:
		return v > b.(float64)
	case int:
		return v > b.(int)
	case string:
		return v > b.(string)
	}
	return false
}

package ternary

// Ternary 三目运算
func Ternary(b bool, v1, v2 interface{}) interface{} {
	if b {
		return v1
	}
	return v2
}

func String(b bool, v1, v2 string) string {
	if b {
		return v1
	}
	return v2
}

func StringEmpty(v1, v2 string) string {
	if v1 != "" {
		return v1
	}
	return v2
}

func Int(b bool, v1, v2 int) int {
	if b {
		return v1
	}
	return v2
}

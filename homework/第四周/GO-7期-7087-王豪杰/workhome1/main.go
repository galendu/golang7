package workhome1

import "fmt"

func Square(num interface{}) interface{} {
	switch result := num.(type) {
	case int:
		return result * result
	case byte:
		return result * result
	case float32:
		return fmt.Sprintf("%.2f", result*result)
	case float64:
		return fmt.Sprintf("%.2f", result*result)
	default:
		return "类型不正确..."
	}
}

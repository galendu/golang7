package work4

import "fmt"

func Square(num interface{}) interface{} {
	switch v := num.(type) {
	case float32:
		return v * v
	case float64:
		return v * v
	case int:
		return v * v
	case byte:
		return v * v
	default:
		fmt.Printf("unsurport data type %T\n", num)
		return nil
	}
}

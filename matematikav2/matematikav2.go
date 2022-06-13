package matematikav2

func Sayarray(num ...int) []string {
	result := []string{}
	for _, i := range num {
		a := i % 2
		if a == 0 {
			result = append(result, "GENAP")
		} else {
			result = append(result, "GANJIL")
		}
	}
	return result
}

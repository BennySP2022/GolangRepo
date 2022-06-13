package matematikav1

func sayganjilgenap(a int) string {
	b := a % 2
	if b == 0 {
		return "Genap"
	} else {
		return "Ganjil"
	}
}

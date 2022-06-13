package main

import (
	"fmt"
	"ganjil-genap/matematikav1"
	"ganjil-genap/matematikav2"
)

func main() {
	xhasil := matematikav1.Sayganjilgenap(5)
	fmt.Println(xhasil)

	number := []int{2, 7, 4, 9}
	xhasil2 := matematikav2.Sayarray(number...)
	fmt.Println(xhasil2)
}

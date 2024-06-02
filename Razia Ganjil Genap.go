package main

import (
	"fmt"
	"strings"
)

type Kendaraan struct {
	Name string
	Plat string
	Type string
	Rute []string
}

// Function untuk menentukan apakah nomor plat genap atau ganjil
func isPlatGenap(plat string) bool {
	lastDigit := plat[len(plat)-1]
	return (lastDigit-'0')%2 == 0
}

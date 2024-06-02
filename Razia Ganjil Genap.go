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

// Function untuk mengecek pelanggaran razia ganjil genap
func kenaRazia(date int, data []Kendaraan) []map[string]interface{} {
	lokasiGanjilGenap := map[string]bool{
		"Gajah Mada":       true,
		"Hayam Wuruk":      true,
		"Sisingamangaraja": true,
		"Panglima Polim":   true,
		"Fatmawati":        true,
		"Tomang Raya":      true,
	}

	pelanggaran := []map[string]interface{}{}
	for _, kendaraan := range data {
		if strings.Contains(kendaraan.Type, "Mobil") {
			for _, rute := range kendaraan.Rute {
				if _, ok := lokasiGanjilGenap[rute]; ok {
					platGenap := isPlatGenap(kendaraan.Plat)
					if (date%2 == 0 && !platGenap) || (date%2 != 0 && platGenap) {
						pelanggaran = append(pelanggaran, map[string]interface{}{"name": kendaraan.Name, "tilang": 1})
						break
					}
				}
			}
		}
	}

	return pelanggaran
}

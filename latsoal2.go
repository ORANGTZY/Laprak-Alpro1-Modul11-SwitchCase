package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	var tarif int
	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}

	switch kendaraan {
	case "motor":
		tarif = durasi * 2000

	case "mobil":
		tarif = durasi * 5000

	case "truk":
		tarif = durasi * 8000

	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}

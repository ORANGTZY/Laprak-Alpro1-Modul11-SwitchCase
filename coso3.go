package main

import "fmt"

func main() {
	var tipe_kendaran string
	var durasi, tarif int

	fmt.Print("Masukkan jenis kendaran (Motor/Mobil/Truk): ")
	fmt.Scan(&tipe_kendaran)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch {
	case tipe_kendaran == "Motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case tipe_kendaran == "Motor" && durasi > 2:
		tarif = 9000
	case tipe_kendaran == "Mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case tipe_kendaran == "Mobil" && durasi > 2:
		tarif = 20000
	case tipe_kendaran == "Truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case tipe_kendaran == "Truk" && durasi > 2:
		tarif = 35000
	default:
		fmt.Println("Jenis kendaran atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif Parkir : Rp %d", tarif)
}

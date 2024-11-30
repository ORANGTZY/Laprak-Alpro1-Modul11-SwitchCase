package main

import "fmt"

func main() {
	var ph float64
	fmt.Scan(&ph)

	switch {
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air Layak Minum")
	case (ph > 0 && ph < 6.5) || (ph <= 14 && ph > 8.6):
		fmt.Println("Air Tidak Layak Minum")
	default:
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14")
	}
}

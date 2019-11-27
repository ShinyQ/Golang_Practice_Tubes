package main

import "fmt"

const N = 3

var arrBerat [N]float64

func main() {

	var berat float64

	for i := 0; i < N; i++ {
		fmt.Print("Masukkan Berat Bayi Ke-", i+1, ": ")
		fmt.Scanln(&berat)
		arrBerat[i] = berat
	}

	maxmin()
	rerata()

}

func maxmin() {

	max := arrBerat[0]
	min := arrBerat[0]
	for i := 0; i < N; i++ {
		if max < arrBerat[i] {
			max = arrBerat[i]
		}

		if min > arrBerat[i] {
			min = arrBerat[i]
		}
	}

	fmt.Println("Nilai Maksimum : ", max)
	fmt.Println("Nilai Minimum : ", min)
}

func rerata() {
	rerata := 0.0
	for i := 0; i < N; i++ {
		rerata += arrBerat[i]
	}

	fmt.Println("Rata-Rata Berat Bayi : ", rerata)
}

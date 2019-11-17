package main

import "fmt"

var (
	Menu     int
	dataBuku []Buku
)

type Buku struct {
	ID, Judul, Penerbit, Tahun string
	Stok                       int
}

func tambahBuku() {
	var (
		Selesai                               bool
		inputLagi, ID, Judul, Penerbit, Tahun string
		Stok                                  int
	)

	jumlah := len(dataBuku)

	for i := jumlah; Selesai != true; i++ {
		fmt.Println("\nMasukkan Data Buku")

		fmt.Print("Masukkan ID Buku : ")
		fmt.Scanln(&ID)

		fmt.Print("Masukkan Judul Buku : ")
		fmt.Scanln(&Judul)

		fmt.Print("Masukkan Penerbit Buku : ")
		fmt.Scanln(&Penerbit)

		fmt.Print("Masukkan Tahun Buku : ")
		fmt.Scanln(&Tahun)

		fmt.Print("Masukkan Stok Buku : ")
		fmt.Scanln(&Stok)

		buku := Buku{
			ID, Judul, Penerbit, Tahun, Stok,
		}

		dataBuku = append(dataBuku, buku)

		fmt.Println("SUKSES MENAMBAH DATA BUKU")

		fmt.Print("\nInput Lagi ? (Y/N) ")
		fmt.Scanln(&inputLagi)

		if inputLagi == "N" || inputLagi == "No" {
			Selesai = true
		}
	}
	main()
}

func main() {

	fmt.Println(dataBuku)

	fmt.Println("1. Tambah Data Buku")
	fmt.Println("2. Cari Data Buku")
	fmt.Println("3. Tambah Stok Buku")
	fmt.Println("4. Cari Buku Berdasarkan Jumlah")
	fmt.Println("5. Ubah Penerbit Buku")

	fmt.Print("Pilih Menu : ")
	fmt.Scanln(&Menu)

	if Menu == 1 {
		tambahBuku()
	} else if Menu == 2 {

	} else if Menu == 3 {

	} else if Menu == 4 {

	} else if Menu == 5 {

	} else {
		fmt.Println("Menu Tersebut Tidak Ada !")
		main()
	}
}

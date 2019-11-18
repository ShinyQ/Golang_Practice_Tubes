package main

import "fmt"

var (
	Menu     int
	dataBuku []Buku
)

type Buku struct {
	ID, Judul, Penerbit string
	Stok, Tahun         int
}

func tambahBuku() {
	var (
		Selesai                        bool
		inputLagi, ID, Judul, Penerbit string
		Stok, Tahun                    int
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

func prosesCariBuku(key string) int {
	var index int
	index = -1
	for i := 0; i < len(dataBuku); i++ {
		if dataBuku[i].ID == key {
			index = i
		}
	}
	return index
}

func cariBuku() {
	var (
		ID string
	)

	fmt.Print("Masukkan ID Buku : ")
	fmt.Scanln(&ID)

	index := prosesCariBuku(ID)
	if index != -1 {
		fmt.Println("\nData Ditemukan :)")
		fmt.Println("ID :", dataBuku[index].ID)
		fmt.Println("Nama :", dataBuku[index].Judul)
		fmt.Println("Penerbit :", dataBuku[index].Penerbit)
		fmt.Println("Stok :", dataBuku[index].Stok)
		fmt.Println("Tahun :", dataBuku[index].Tahun)
	} else {
		fmt.Println("\nData Tidak Ditemukan :(")
	}

	main()
}

func tambahStok() {
	var ID, InputBuku string
	var maxTambah, Stok int
	var validStok bool

	fmt.Print("Masukkan Kode Buku : ")
	fmt.Scanln(&ID)

	index := prosesCariBuku(ID)

	if index != -1 {

		if dataBuku[index].Tahun > 2018 {

			if dataBuku[index].Stok >= 3 || dataBuku[index].Stok <= 6 {
				maxTambah = 5
			} else if dataBuku[index].Stok >= 7 || dataBuku[index].Stok <= 9 {
				maxTambah = 1
			} else if dataBuku[index].Stok == 10 {
				maxTambah = 0
			}

		} else if dataBuku[index].Tahun >= 2010 || dataBuku[index].Tahun <= 2018 {

			if dataBuku[index].Stok >= 3 || dataBuku[index].Stok <= 6 {
				maxTambah = 7
			} else if dataBuku[index].Stok >= 7 || dataBuku[index].Stok <= 9 {
				maxTambah = 2
			} else if dataBuku[index].Stok == 15 {
				maxTambah = 0
			}

		} else if dataBuku[index].Tahun < 2010 {
			maxTambah = -1
		}

		if maxTambah == 0 {
			fmt.Println("\nMaaf Stok Buku Sudah Tidak Dapat Ditambah Lagi")
		} else {
			for !validStok {
				fmt.Print("Masukkan Jumlah Penambahan Stok : ")
				fmt.Scanln(&Stok)

				if Stok > maxTambah {
					fmt.Println("Hanya Dapat Menambahkan", maxTambah, "Stok Buku")
				} else {
					dataBuku[index].Stok += Stok
					validStok = true
					fmt.Println("Stok Buku Berhasil Ditambahkan !")
				}
			}
		}

	} else {
		fmt.Println("ID Buku Tersebut Tidak Ada")
		fmt.Print("Ingin Menambahkan Data Buku Baru ? (Y/N) ")
		fmt.Scanln(&InputBuku)

		if InputBuku == "Y" || InputBuku == "y" {
			tambahBuku()
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
		cariBuku()
	} else if Menu == 3 {
		tambahStok()
	} else if Menu == 4 {

	} else if Menu == 5 {

	} else {
		fmt.Println("Menu Tersebut Tidak Ada !")
		main()
	}
}

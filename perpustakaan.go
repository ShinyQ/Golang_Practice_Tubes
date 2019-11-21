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

				if Stok > maxTambah || maxTambah != -1 {
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

func sortBukuJumlah() {
	var dataBukuSort []Buku

	j := 0
	for i := 0; i < len(dataBuku); i++ {
		if dataBuku[i].Stok >= 7 || dataBuku[i].Stok <= 9 {
			dataBukuSort = append(dataBukuSort, dataBuku[i])
			j++
		}
	}

	sorted := false
	jumlahBuku := len(dataBukuSort)

	if jumlahBuku != 0 {

		for !sorted {
			swapped := false
			for i := 0; i < jumlahBuku-1; i++ {
				if dataBukuSort[i].Stok < dataBukuSort[i+1].Stok {
					dataBukuSort[i+1], dataBukuSort[i] = dataBukuSort[i], dataBukuSort[i+1]
					swapped = true
				}
			}
			if !swapped {
				sorted = true
			}
			jumlahBuku--
		}

		for i := 0; i < jumlahBuku; i++ {
			fmt.Println("")
			fmt.Println("ID Buku :", dataBukuSort[i].ID)
			fmt.Println("Judul Buku :", dataBukuSort[i].Judul)
			fmt.Println("Penerbit Buku :", dataBukuSort[i].Penerbit)
			fmt.Println("Stok Buku :", dataBukuSort[i].Stok)
			fmt.Println("Tahun Buku :", dataBukuSort[i].Tahun)
			fmt.Println("")
		}
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}

	main()
}

func updatePenerbitBuku() {
	var penerbit string
	var dataPenerbit []Buku

	fmt.Print("Masukkan Nama Penerbit : ")
	fmt.Scanln(&penerbit)

	jumlahBuku := len(dataBuku)
	j := 0

	for i := 0; i < jumlahBuku; i++ {
		if dataBuku[i].Penerbit == penerbit {
			dataPenerbit[j] = dataBuku[i]
			j++
		}
	}

	jumlahBukuPenerbit := len(dataPenerbit)

	fmt.Println("Ditemukan", jumlahBuku, "Buku Dengan Nama Penerbit", penerbit)

	for i := 0; i < jumlahBukuPenerbit; i++ {
		fmt.Println("ID Buku :", dataPenerbit[i].ID)
		fmt.Println("Judul Buku :", dataPenerbit[i].Judul)
		fmt.Println("Penerbit Buku :", dataPenerbit[i].Penerbit)
		fmt.Println("Stok Buku :", dataPenerbit[i].Stok)
		fmt.Println("Tahun Buku :", dataPenerbit[i].Tahun)
		fmt.Println("")
	}

	fmt.Print("Masukkan Nama Penerbit Yang Akan Diubah : ")
	fmt.Scanln(&penerbit)
	fmt.Println("")

	for i := 0; i < jumlahBukuPenerbit; i++ {
		dataPenerbit[i].Penerbit = penerbit
	}

	fmt.Println("Data Nama Penerbit Sukses Diubah")

	for i := 0; i < jumlahBukuPenerbit; i++ {
		fmt.Println("ID Buku :", dataPenerbit[i].ID)
		fmt.Println("Judul Buku :", dataPenerbit[i].Judul)
		fmt.Println("Penerbit Buku :", dataPenerbit[i].Penerbit)
		fmt.Println("Stok Buku :", dataPenerbit[i].Stok)
		fmt.Println("Tahun Buku :", dataPenerbit[i].Tahun)
		fmt.Println("")
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
	fmt.Println("6. Keluar Program")

	fmt.Print("Pilih Menu : ")
	fmt.Scanln(&Menu)

	if Menu == 1 {
		tambahBuku()
	} else if Menu == 2 {
		cariBuku()
	} else if Menu == 3 {
		tambahStok()
	} else if Menu == 4 {
		sortBukuJumlah()
	} else if Menu == 5 {
		updatePenerbitBuku()
	} else if Menu == 6 {
		defer fmt.Println("Sukses Keluar Program")
	} else {
		fmt.Println("Menu Tersebut Tidak Ada !")
		main()
	}
}

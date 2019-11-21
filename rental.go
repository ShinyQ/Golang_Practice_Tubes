package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
)

var (
	itemAnggota []Anggota
	itemRental  []Rental

	Menu    int
	scanner = bufio.NewScanner(os.Stdin)
)

type Anggota struct {
	ID, nama, alamat  string
	umur, saldo, poin int
}

type Rental struct {
	ID, kota, mobil string
	harga           int
}

func inputDataAnggota() {
	var (
		Selesai           bool
		ID, nama, alamat  string
		umur, saldo, poin int
		inputLagi         string
	)

	for i := 0; !Selesai; i++ {

		fmt.Println("\nMasukkan Data-Data Anggota Baru")

		fmt.Print("ID Anggota : ")
		fmt.Scanln(&ID)

		fmt.Print("Nama : ")
		scanner.Scan()
		nama = scanner.Text()

		fmt.Print("Alamat : ")
		scanner.Scan()
		alamat = scanner.Text()

		fmt.Print("Umur : ")
		fmt.Scanln(&umur)

		fmt.Print("Saldo : ")
		fmt.Scanln(&saldo)

		poin = 0

		anggota := Anggota{ID, nama, alamat, umur, saldo, poin}
		itemAnggota = append(itemAnggota, anggota)

		fmt.Println("\n Sukses Menambah Data Anggota")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID Anggota", "Nama", "Alamat", "Umur", "Saldo", "Poin"})
		table.Append(
			[]string{
				itemAnggota[i].ID,
				itemAnggota[i].nama,
				strconv.Itoa(itemAnggota[i].umur),
				itemAnggota[i].alamat,
				"Rp" + humanize.Comma(int64(itemAnggota[i].saldo)) + ",00",
				strconv.Itoa(itemAnggota[i].poin),
			},
		)
		table.Render()

		fmt.Print("\nInput Lagi (Y/N) ? ")
		fmt.Scanln(&inputLagi)

		if inputLagi == "n" || inputLagi == "N" {
			Selesai = true
		}
	}
	main()
}

func cariProsesPemesan(ID string) int {
	index := -1
	for i := 0; i < len(itemAnggota); i++ {
		if itemAnggota[i].ID == ID {
			index = i
		}
	}

	return index
}

func cariPemesan() {
	var ID string

	fmt.Println("\nCari Data Pemesan")
	fmt.Print("Masukkan Kode Anggota : ")
	fmt.Scanln(&ID)

	index := cariProsesPemesan(ID)

	if index != -1 {
		fmt.Println("\nData Ditemukan :")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID Anggota", "Nama", "Alamat", "Umur", "Saldo", "Poin"})
		table.Append(
			[]string{
				itemAnggota[index].ID,
				itemAnggota[index].nama,
				strconv.Itoa(itemAnggota[index].umur),
				itemAnggota[index].alamat,
				"Rp" + humanize.Comma(int64(itemAnggota[index].saldo)) + ",00",
				strconv.Itoa(itemAnggota[index].poin),
			},
		)
		table.Render()
	} else {
		fmt.Println("ID Pemesan Tidak Valid, Data Tidak Ditemukan")
	}
}

func inputTujuan() {
	var (
		Selesai                                         bool
		ID, kota, mobil                                 string
		inputLagi                                       string
		validKota, validMobil, validInputan, validSaldo bool
		validasiPesanan                                 string
	)

	for i := 0; !Selesai; i++ {
		fmt.Println("\nMasukkan Data-Data Perentalan")

		fmt.Print("ID Anggota : ")
		fmt.Scanln(&ID)

		fmt.Print("Kota : ")
		fmt.Scanln(&kota)
		for !validKota {
			if kota == "A" || kota == "B" {
				validKota = true
			} else {
				fmt.Println("Kota Tidak Valid")
				fmt.Print("\nKota : ")
				fmt.Scanln(&kota)
			}
		}

		fmt.Print("Jenis Mobil: ")
		fmt.Scanln(&mobil)

		for !validMobil {
			if kota == "A" {
				if mobil == "Inova" || mobil == "Avanza" || mobil == "Alya" {
					validMobil = true
				} else {
					fmt.Println("Mobil Tidak Valid")
					fmt.Print("\nJenis Mobil")
					fmt.Scanln(&mobil)
				}
			} else {
				if mobil == "Mobilio" || mobil == "Brio" {
					validMobil = true
				} else {
					fmt.Println("Mobil Tidak Valid")
					fmt.Print("\nJenis Mobil")
					fmt.Scanln(&mobil)
				}
			}
		}

		harga := 0
		poin := 0

		if kota == "A" {
			if mobil == "Inova" {
				harga = 8000 * 50
				poin = 10
			} else if mobil == "Avanza" {
				harga = 5000 * 50
				poin = 5
			} else {
				harga = 4000 * 50
			}
		} else {
			if mobil == "Mobilio" {
				harga = 5000 * 80
			} else {
				harga = 4000 * 80
				poin = 5
			}
		}

		index := cariProsesPemesan(ID)
		if harga > itemAnggota[index].saldo {
			fmt.Println("\nSaldo Anda Tidak Mencukupi")
			fmt.Println("Saldo Anda Saat Ini Adalah : Rp" + humanize.Comma(int64(itemAnggota[index].saldo)) + ",00")
			fmt.Println("Jumlah Saldo Yang Dibutuhkan : RP" + humanize.Comma(int64(harga)) + ",00")
			validSaldo = false
		} else {
			validInputan = true
			validSaldo = true
		}

		if validSaldo == true {
			fmt.Println("\nData Rental Anda  ")
			fmt.Println("Kota Perjalanan : Kota", kota)
			fmt.Println("Jenis Mobil : Mobil", mobil)
			fmt.Println("Total Harga : RP" + humanize.Comma(int64(harga)) + ",00")
			fmt.Print("\nApakah Data Ini Sudah Benar (Ok, Cancel) ? ")
			fmt.Scanln(&validasiPesanan)

			if validasiPesanan == "OK" || validasiPesanan == "Ok" {
				validInputan = true
			} else {
				validInputan = false
				fmt.Println("\n Pesanan DIbatalkan")
			}
		}

		if validInputan == true {
			rental := Rental{ID, kota, mobil, harga}
			itemRental = append(itemRental, rental)

			fmt.Println("\n Sukses Menambah Data Rental")
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID Anggota", "Kota", "Mobil", "Total Harga"})
			table.Append(
				[]string{
					itemRental[i].ID,
					itemRental[i].kota,
					itemRental[i].mobil,
					"Rp" + humanize.Comma(int64(itemRental[i].harga)) + ",00",
				},
			)
			table.Render()

			if poin != 0 {
				poinSekarang := itemAnggota[index].poin + poin
				itemAnggota[index].poin = poinSekarang
				fmt.Println("Selamat Anda Mendapatkan", poin, "Poin")
			}

			fmt.Println("Saldo Anda Sebelumnya : Rp" + humanize.Comma(int64(itemAnggota[index].saldo)) + ",00")
			saldoSekarang := itemAnggota[index].saldo - harga

			itemAnggota[index].saldo = saldoSekarang
			fmt.Println("Saldo Anda Saat Ini : Rp" + humanize.Comma(int64(saldoSekarang)) + ",00")
		}

		fmt.Print("\nInput Lagi (Y/N) ? ")
		fmt.Scanln(&inputLagi)

		if inputLagi == "n" || inputLagi == "N" {
			Selesai = true
		}
	}
	main()
}

func sortDataAnggota() {
	sorted := false
	n := len(itemAnggota)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if itemAnggota[i].nama > itemAnggota[i+1].nama {
				itemAnggota[i+1], itemAnggota[i] = itemAnggota[i], itemAnggota[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n--
	}

	for i := 0; i < len(itemAnggota); i++ {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID Anggota", "Nama", "Alamat", "Umur", "Saldo", "Poin"})
		table.Append(
			[]string{
				itemAnggota[i].ID,
				itemAnggota[i].nama,
				strconv.Itoa(itemAnggota[i].umur),
				itemAnggota[i].alamat,
				"Rp" + humanize.Comma(int64(itemAnggota[i].saldo)) + ",00",
				strconv.Itoa(itemAnggota[i].poin),
			},
		)
		table.Render()
	}
}

func sortDataBySaldo() {
	var sortAnggota []Anggota
	found := false
	for i := 0; i < len(itemAnggota); i++ {
		if itemAnggota[i].saldo > 50000 {
			anggota := Anggota{
				nama:   itemAnggota[i].nama,
				alamat: itemAnggota[i].alamat,
				ID:     itemAnggota[i].ID,
				saldo:  itemAnggota[i].saldo,
				umur:   itemAnggota[i].umur,
				poin:   itemAnggota[i].poin,
			}

			sortAnggota = append(sortAnggota, anggota)
			found = true
		}
	}

	if found == true {
		sorted := false
		n := len(sortAnggota)
		for !sorted {
			swapped := false
			for i := 0; i < n-1; i++ {
				if sortAnggota[i].saldo > sortAnggota[i+1].saldo {
					sortAnggota[i+1], sortAnggota[i] = sortAnggota[i], sortAnggota[i+1]
					swapped = true
				}
			}
			if !swapped {
				sorted = true
			}
			n--
		}

		for i := 0; i < len(sortAnggota); i++ {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID Anggota", "Nama", "Alamat", "Umur", "Saldo", "Poin"})
			table.Append(
				[]string{
					sortAnggota[i].ID,
					sortAnggota[i].nama,
					strconv.Itoa(sortAnggota[i].umur),
					sortAnggota[i].alamat,
					"Rp" + humanize.Comma(int64(sortAnggota[i].saldo)) + ",00",
					strconv.Itoa(sortAnggota[i].poin),
				},
			)
			table.Render()
		}
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}

func tampilTransaksi() {
	if len(itemRental) != 0 {
		for i := 0; i < len(itemRental); i++ {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID Anggota", "Kota", "Total Harga"})
			table.Append(
				[]string{
					itemRental[i].ID,
					itemRental[i].kota,
					itemRental[i].mobil,
					"Rp" + humanize.Comma(int64(itemRental[i].harga)) + ",00",
				},
			)
			table.Render()
		}
	} else {
		fmt.Println("Belum Ada Transaksi ")
	}
}

func main() {

	fmt.Println("Silhakan Pilih Menu : ")
	fmt.Println("1. Input Data Anggota")
	fmt.Println("2. Input Tujuan")
	fmt.Println("3. Cari Data Pemesan")
	fmt.Println("4. Menampilkan Data Anggota")
	fmt.Println("5. Menampilkan Data Anggota Berdasarkan Jumlah Saldo")
	fmt.Println("6. Tampil Transaksi Rental")
	fmt.Println("7. Keluar Program")

	fmt.Print("Silahkan Pilih Menu : ")
	fmt.Scanln(&Menu)

	if Menu == 1 {
		inputDataAnggota()
	} else if Menu == 2 {
		inputTujuan()
	} else if Menu == 3 {
		cariPemesan()
	} else if Menu == 4 {
		sortDataAnggota()
	} else if Menu == 5 {
		sortDataBySaldo()
	} else if Menu == 6 {
		tampilTransaksi()
	} else if Menu == 7 {
		defer fmt.Println("Sukses Keluar Program")
	} else {
		fmt.Println("Menu Tersebut Tidak Ada !")
		main()
	}
}

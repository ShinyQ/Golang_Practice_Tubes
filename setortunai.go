package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type TNasabah struct {
	nik, nama, norek, jenis          string
	pin, saldo, nominalTerakhir      int
	tanggalDaftar, transaksiTerakhir string
}

var TArrNasabah []TNasabah
var pilih int
var waktu = time.Now().Format("2006-01-02")

func cariRekening(jenisRek string) int {
	count := 0
	if len(TArrNasabah) != 0 {
		for i := 0; i < len(TArrNasabah); i++ {
			if TArrNasabah[i].jenis == jenisRek {
				count++
			}
		}
		count++
	} else {
		count = 1
	}

	return count
}

func cari_nik(nik string) int {
	validnik := -1
	for i := 0; i < len(TArrNasabah); i++ {
		if TArrNasabah[i].nik == nik {
			validnik = 0
		}
	}
	return validnik
}

func registrasi() {
	var (
		nik, nama, norek, jenisRek                  string
		pin, jenis, saldo, nominalTerakhir          int
		tanggalDaftar, transaksiTerakhir            string
		validnik, validJenis, validSetoran, Selesai bool
		inputLagi                                   string
	)

	for !Selesai {
		fmt.Println("======================================")
		fmt.Println(">>> REGISTRASI NASABAH BARU <<<")
		fmt.Println("======================================")

		for !validnik {
			fmt.Print("Masukkan NIK Nasabah : ")
			fmt.Scanln(&nik)
			if len(nik) != 16 {
				fmt.Println("\nNIK Harus 16 Digit\n")
			} else if cari_nik(nik) == -1 {
				validnik = true
			} else {
				fmt.Println("\nNIK Tersebut Sudah Ada\n")
			}
		}

		fmt.Print("Masukkan nama Nasabah    : ")
		fmt.Scanln(&nama)

		fmt.Print("Masukkan jenis Nasabah   : ")
		fmt.Scanln(&jenis)

		for !validSetoran {
			fmt.Print("Masukkan setoran Nasabah : ")
			fmt.Scanln(&saldo)

			if saldo < 1000 && saldo%100 != 0 {
				fmt.Println("\nJumlah Setoran Tidak Valid\n")
			} else {
				validSetoran = true
			}
		}

		if jenis == 1 {
			jenisRek = "Silver"
			norek = "XYZ-S" + fmt.Sprintf("%03d", cariRekening(jenisRek))
		} else if jenis == 2 {
			jenisRek = "Gold"
			norek = "XYZ-G" + fmt.Sprintf("%03d", cariRekening(jenisRek))
		} else {
			jenisRek = "Platinum"
			norek = "XYZ-P" + fmt.Sprintf("%03d", cariRekening(jenisRek))
		}

		tanggalDaftar = waktu
		transaksiTerakhir = waktu
		rand.Seed(time.Now().UnixNano())
		pin = 100000 + rand.Intn(999999-100000+1)
		index := len(TArrNasabah)

		nasabah := TNasabah{
			nik, nama, norek, jenisRek, pin,
			saldo, nominalTerakhir,
			tanggalDaftar, transaksiTerakhir,
		}
		TArrNasabah = append(TArrNasabah, nasabah)

		fmt.Println("\nData Nasabah Sukses Di Inputkan :")
		fmt.Println("NIK :", TArrNasabah[index].nik)
		fmt.Println("Nama :", TArrNasabah[index].nama)
		fmt.Println("PIN :", TArrNasabah[index].pin)
		fmt.Println("Nomor Rekening :", TArrNasabah[index].norek)
		fmt.Println("Jenis ATM :", TArrNasabah[index].jenis)
		fmt.Println("Setoran Awal :", TArrNasabah[index].saldo)

		fmt.Print("\nInput Data Nasabah Lagi (Y/N)? ")
		fmt.Scanln(&inputLagi)

		if inputLagi == "N" || inputLagi == "n" {
			Selesai = true
		}
	}
	login_cs()
}

func login_cs() {
	var username, password string
	var validcs bool

	fmt.Println("\n======================================")
	fmt.Println("1). LOGIN CUSTOMER SERVICE")
	fmt.Println("======================================")

	for !validcs {
		fmt.Print("Masukkan username Customer Service : ")
		fmt.Scanln(&username)
		fmt.Print("Masukkan password Customer Service : ")
		fmt.Scanln(&password)
		if username != "admin" && password != "root" {
			fmt.Println("\nUsername Atau Password Salah\n")
		} else {
			validcs = true
		}
	}

	fmt.Println("\n======================================")
	fmt.Println(">>> MENU CUSTOMER SERVICE <<<")
	fmt.Println("======================================")
	fmt.Println("1). Registrasi nasabah baru")
	fmt.Println("2). Menu Utama")
	fmt.Print("Masukkan nomor menu pilihan anda : ")
	fmt.Scanln(&pilih)
	if pilih == 1 {
		registrasi()
	} else if pilih == 2 {
		main()
	} else {
		fmt.Println("ERROR, Masukkan nomor pilihan yang tersedia")
	}
}

func login_nasabah_validasi(rekening string, pin int) int {
	index := -1
	for i := 0; i < len(TArrNasabah); i++ {
		if TArrNasabah[i].norek == rekening {
			index = i
		}
	}

	if TArrNasabah[index].pin != pin {
		index = -1
	}

	return index
}

func nasabah_setoran(index int) {
	var (
		setoran, max          int
		validSetoran, Selesai bool
		inputLagi             string
	)

	if TArrNasabah[index].jenis == "Silver" {
		max = 10000
	} else if TArrNasabah[index].jenis == "Gold" {
		max = 20000
	} else {
		max = 50000
	}

	for !Selesai {
		for !validSetoran {
			fmt.Println("Masukkan Jumlah Setoran (Max :", max, ") : ")
			fmt.Scanln(&setoran)
			if setoran < 0 {
				fmt.Println("Jumlah Setoran Tidak Valid")
			} else if setoran%100 != 0 {
				fmt.Println("Setoran Harus Kelipatan 100")
			} else if setoran > max {
				fmt.Println("Anda Sudah Melebihi Maksimal Setoran")
			} else {
				validSetoran = true
			}
		}

		TArrNasabah[index].saldo += setoran
		TArrNasabah[index].transaksiTerakhir = waktu
		TArrNasabah[index].nominalTerakhir = setoran

		fmt.Println("Sukses Menambahkan Setoran Sebesar", setoran)

		fmt.Print("\nInput Lagi ? ")
		fmt.Scanln(&inputLagi)
		if inputLagi == "N" || inputLagi == "n" {
			Selesai = true
		}
	}

	login_nasabah()
}

func login_nasabah() {
	var rekening string
	var pin int

	fmt.Println("======================================")
	fmt.Println("2). LOGIN NASABAH")
	fmt.Println("======================================")
	fmt.Println("1). Tambah Setoran")
	fmt.Println("2). Menu Utama")

	fmt.Print("Silahkan Pilih Menu : ")
	fmt.Scanln(&pilih)

	if pilih == 1 {
		fmt.Print("Masukkan No. Rekening Nasabah : ")
		fmt.Scanln(&rekening)
		fmt.Print("Masukkan PIN Nasabah : ")
		fmt.Scanln(&pin)

		index := login_nasabah_validasi(rekening, pin)
		if index != -1 {
			nasabah_setoran(index)
		} else {
			fmt.Println("\nNomor Rekening Atau PIN Salah\n")
			login_nasabah()
		}
	} else if pilih == 2 {
		main()
	} else {
		fmt.Println("Menu Tersebut Tidak Ada")
	}
}

func proses_cari_nasabah(norek string) int {
	index := -1
	for i := 0; i < len(TArrNasabah); i++ {
		if TArrNasabah[i].norek == norek {
			index = i
		}
	}
	return index
}

func cari_nasabah() {
	var norek string

	fmt.Print("Masukkan Nomor Rekening Nasabah : ")
	fmt.Scanln(&norek)

	index := proses_cari_nasabah(norek)

	if index != -1 {
		fmt.Println("Data Ditemukan :")
		fmt.Println("NIK :", TArrNasabah[index].nik)
		fmt.Println("Nama :", TArrNasabah[index].nama)
		fmt.Println("PIN :", TArrNasabah[index].pin)
		fmt.Println("Nomor Rekening :", TArrNasabah[index].norek)
		fmt.Println("Jenis ATM :", TArrNasabah[index].jenis)
		fmt.Println("Setoran :", TArrNasabah[index].saldo)
		fmt.Println("Transaksi Terahir :", TArrNasabah[index].transaksiTerakhir)
		fmt.Println("Nominal Transaksi Terahir :", TArrNasabah[index].nominalTerakhir)
	} else {
		fmt.Println("Maaf, Data Nasabah Tidak Ditemukan")
	}
	login_manager()
}

func tampil_nasabah() {
	var sortNasabah []TNasabah
	var sorted bool

	jumlah := len(TArrNasabah)
	if jumlah != 0 {

		for i := 0; i < len(TArrNasabah); i++ {
			nasabah := TNasabah{
				nama:              TArrNasabah[i].nama,
				jenis:             TArrNasabah[i].jenis,
				nik:               TArrNasabah[i].nik,
				nominalTerakhir:   TArrNasabah[i].nominalTerakhir,
				norek:             TArrNasabah[i].norek,
				pin:               TArrNasabah[i].pin,
				saldo:             TArrNasabah[i].saldo,
				tanggalDaftar:     TArrNasabah[i].tanggalDaftar,
				transaksiTerakhir: TArrNasabah[i].transaksiTerakhir,
			}
			sortNasabah = append(sortNasabah, nasabah)
		}

		for !sorted {
			swapped := false
			for i := 0; i < jumlah-1; i++ {
				if sortNasabah[i].norek > sortNasabah[i+1].norek {
					sortNasabah[i+1], sortNasabah[i] = sortNasabah[i], sortNasabah[i+1]
					swapped = true
				}
			}
			if !swapped {
				sorted = true
			}
			jumlah--
		}

		fmt.Println("Ditemukan : ", len(sortNasabah), "Data Nasabah")
		for i := 0; i < len(sortNasabah); i++ {
			fmt.Println("")
			fmt.Println("NIK :", sortNasabah[i].nik)
			fmt.Println("Nama :", sortNasabah[i].nama)
			fmt.Println("PIN :", sortNasabah[i].pin)
			fmt.Println("Nomor Rekening :", sortNasabah[i].norek)
			fmt.Println("Jenis ATM :", sortNasabah[i].jenis)
			fmt.Println("Setoran :", sortNasabah[i].saldo)
			fmt.Println("Transaksi Terahir :", sortNasabah[i].transaksiTerakhir)
			fmt.Println("Nominal Transaksi Terahir :", sortNasabah[i].nominalTerakhir)
		}
		fmt.Println("")
	} else {
		fmt.Println("Tidak Ada Data Nasabah")
	}

	login_manager()
}

func nasabah_jenis_rekening() {
	var jenis string
	var validJenis, sorted bool
	var sortNasabah []TNasabah

	for !validJenis {
		fmt.Print("Masukkan Jenis Rekening : ")
		fmt.Scanln(&jenis)

		jenis = strings.ToUpper(jenis)
		if jenis == "SILVER" || jenis == "GOLD" || jenis == "PLATINUM" {
			validJenis = true
		} else {
			fmt.Println("Jenis Rekening Tersebut Tidak Ada\n")
		}
	}

	if len(TArrNasabah) != 0 {
		for i := 0; i < len(TArrNasabah); i++ {
			if strings.ToUpper(TArrNasabah[i].jenis) == strings.ToUpper(jenis) {
				nasabah := TNasabah{
					nama:              TArrNasabah[i].nama,
					jenis:             TArrNasabah[i].jenis,
					nik:               TArrNasabah[i].nik,
					nominalTerakhir:   TArrNasabah[i].nominalTerakhir,
					norek:             TArrNasabah[i].norek,
					pin:               TArrNasabah[i].pin,
					saldo:             TArrNasabah[i].saldo,
					tanggalDaftar:     TArrNasabah[i].tanggalDaftar,
					transaksiTerakhir: TArrNasabah[i].transaksiTerakhir,
				}
				sortNasabah = append(sortNasabah, nasabah)
			}
		}
		if len(sortNasabah) != 0 {
			jumlah := len(sortNasabah)
			for !sorted {
				swapped := false
				for i := 0; i < jumlah-1; i++ {
					if sortNasabah[i].nama > sortNasabah[i+1].nama {
						sortNasabah[i+1], sortNasabah[i] = sortNasabah[i], sortNasabah[i+1]
						swapped = true
					}
				}
				if !swapped {
					sorted = true
				}
				jumlah--
			}

			fmt.Println("Ditemukan : ", len(sortNasabah), "Data Nasabah Dengan Jenis Rekening", sortNasabah[0].jenis)
			for i := 0; i < len(sortNasabah); i++ {
				fmt.Println("")
				fmt.Println("NIK :", sortNasabah[i].nik)
				fmt.Println("Nama :", sortNasabah[i].nama)
				fmt.Println("PIN :", sortNasabah[i].pin)
				fmt.Println("Nomor Rekening :", sortNasabah[i].norek)
				fmt.Println("Jenis ATM :", sortNasabah[i].jenis)
				fmt.Println("Setoran :", sortNasabah[i].saldo)
				fmt.Println("Transaksi Terahir :", sortNasabah[i].transaksiTerakhir)
				fmt.Println("Nominal Transaksi Terahir :", sortNasabah[i].nominalTerakhir)
			}
		} else {
			fmt.Println("Data Nasabah Dengan Jenis Rekening", jenis, "Tidak Ditemukan")
		}

		fmt.Println("")
	} else {
		fmt.Println("Belum Ada Data Nasabah")
	}

	login_manager()
}

func login_manager() {
	var username, password string
	var pilih int

	fmt.Println("======================================")
	fmt.Println("3). LOGIN MANAGER")
	fmt.Println("======================================")
	fmt.Print("Masukkan username Manager : ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password Manager : ")
	fmt.Scanln(&password)
	for username != "manager" && password != "root" {
		fmt.Print("Masukkan username Manager : ")
		fmt.Scanln(&username)
		fmt.Print("Masukkan password Manager : ")
		fmt.Scanln(&password)
	}
	fmt.Println("======================================")
	fmt.Println(">>> MENU MANAGER <<<")
	fmt.Println("======================================")
	fmt.Println("1). Tampilkan satu nasabah")
	fmt.Println("2). Tampilkan seluruh nasabah")
	fmt.Println("3). Tampilkan nasabah yang tidak melakukan transaksi")
	fmt.Println("4). Tampilkan nasabah per jenis rekening")
	fmt.Println("5). Menu Utama")
	fmt.Print("Masukkan nomor menu pilihan anda : ")
	fmt.Scanln(&pilih)
	if pilih == 1 {
		cari_nasabah()
	} else if pilih == 2 {
		tampil_nasabah()
	} else if pilih == 3 {
		fmt.Print("DUMMY TIDAK MELAKUKAN TRANSAKSI")
	} else if pilih == 4 {
		nasabah_jenis_rekening()
	} else if pilih == 5 {
		main()
	} else {
		fmt.Println("ERROR, Masukkan nomor pilihan yang tersedia\n")
	}
}

func main() {

	TArrNasabah = append(TArrNasabah,
		TNasabah{"1234567890123456", "Dony", "XYZ-S002", "Silver", 123456, 500000, 500000, "2019-11-25", "2019-11-25"},
		TNasabah{"1234567890123456", "Aini", "XYZ-S001", "Silver", 123456, 500000, 500000, "2019-11-25", "2019-11-25"},
		TNasabah{"1234567890123456", "Angela", "XYZ-G001", "Gold", 123456, 500000, 500000, "2019-11-25", "2019-11-25"},
	)

	fmt.Println(TArrNasabah)

	fmt.Println("======================================")
	fmt.Println("MENU UTAMA SETOR TUNAI ATM")
	fmt.Println("======================================")
	fmt.Println("1). LOGIN CUSTOMER SERVICE")
	fmt.Println("2). LOGIN NASABAH")
	fmt.Println("3). LOGIN MANAGER")
	fmt.Println("======================================")
	fmt.Print("Masukkan nomor menu pilihan anda : ")
	fmt.Scanln(&pilih)

	if pilih == 1 {
		login_cs()
	} else if pilih == 2 {
		login_nasabah()
	} else if pilih == 3 {
		login_manager()
	} else {
		fmt.Println("ERROR, Masukkan nomor pilihan yang tersedia")
	}

}

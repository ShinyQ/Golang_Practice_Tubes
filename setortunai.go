package main

import (
	"fmt"
	"math/rand"
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

func registrasi() {
	var nik, nama, norek, jenisRek string
	var pin, jenis, saldo, nominalTerakhir int
	var tanggalDaftar, transaksiTerakhir string

	for len(nik) != 16 || nama == "" || jenis < 1 && jenis > 3 || saldo < 1000 {
		fmt.Println("======================================")
		fmt.Println(">>> REGISTRASI NASABAH BARU <<<")
		fmt.Println("======================================")
		fmt.Print("Masukkan NIK Nasabah     : ")
		fmt.Scanln(&nik)

		fmt.Print("Masukkan nama Nasabah    : ")
		fmt.Scanln(&nama)

		fmt.Print("Masukkan jenis Nasabah   : ")
		fmt.Scanln(&jenis)

		fmt.Print("Masukkan setoran Nasabah : ")
		fmt.Scanln(&saldo)
	}

	if jenis == 1 {
		jenisRek = "Silver"
		norek = "XYZ-S" + fmt.Sprintf("%03d", cariRekening(jenisRek))
	} else if jenis == 2 {
		jenisRek = "Gold"
		norek = "XYZ-G" + fmt.Sprintf("%04d", cariRekening(jenisRek))
	} else {
		jenisRek = "Platinum"
		norek = "XYZ-P" + fmt.Sprintf("%04d", cariRekening(jenisRek))
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

	login_cs()
}

func login_cs() {
	var username, password string

	fmt.Println("\n======================================")
	fmt.Println("1). LOGIN CUSTOMER SERVICE")
	fmt.Println("======================================")
	fmt.Print("Masukkan username Customer Service : ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password Customer Service : ")
	fmt.Scanln(&password)
	for username != "admin" && password != "root" {
		fmt.Print("Masukkan username Customer Service : ")
		fmt.Scanln(&username)
		fmt.Print("Masukkan password Customer Service : ")
		fmt.Scanln(&password)
	}
	fmt.Println("======================================")
	fmt.Println(">>> MENU CUSTOMER SERVICE <<<")
	fmt.Println("======================================")
	fmt.Println("1). Registrasi nasabah baru")
	fmt.Print("Masukkan nomor menu pilihan anda : ")
	fmt.Scanln(&pilih)
	if pilih == 1 {
		registrasi()
	} else {
		fmt.Println("ERROR, Masukkan nomor pilihan yang tersedia")
	}
}

func login_ns(rekening string, pin int) int {
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
	var setoran int
	var max int
	var validSetoran bool

	if TArrNasabah[index].jenis == "Silver" {
		max = 10000
	} else if TArrNasabah[index].jenis == "Gold" {
		max = 20000
	} else {
		max = 50000
	}

	for !validSetoran {
		fmt.Println("Masukkan Jumlah Setoran (Max :", max, ") : ")
		fmt.Scanln(&setoran)

		if setoran%100 != 0 {
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

	login_nasabah()
}

func login_nasabah() {
	var rekening string
	var pin int

	fmt.Println("======================================")
	fmt.Println("2). LOGIN NASABAH")
	fmt.Println("======================================")
	fmt.Print("Masukkan No. Rekening Nasabah : ")
	fmt.Scanln(&rekening)
	fmt.Print("Masukkan PIN Nasabah : ")
	fmt.Scanln(&pin)

	index := login_ns(rekening, pin)
	if index != -1 {
		nasabah_setoran(index)
	} else {
		fmt.Println("\nNomor Rekening Atau PIN Salah\n")
		login_nasabah()
	}
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
	fmt.Print("Masukkan nomor menu pilihan anda : ")
	fmt.Scanln(&pilih)
	if pilih == 1 {
		fmt.Print("DUMMY TAMPIL SATU NASABAH")
	} else if pilih == 2 {
		fmt.Print("DUMMY TAMPIL SELURUH NASABAH")
	} else if pilih == 3 {
		fmt.Print("DUMMY TIDAK MELAKUKAN TRANSAKSI")
	} else if pilih == 4 {
		fmt.Print("DUMMY PER JENIS REKENING")
	} else {
		fmt.Println("ERROR, Masukkan nomor pilihan yang tersedia")
	}
}

func main() {

	TArrNasabah = append(TArrNasabah,
		TNasabah{"1234567890123456", "Dony", "XYZ-S0001", "Silver", 123456, 500000, 500000, "2019-11-25", "2019-11-25"},
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

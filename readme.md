## Case 1 : Aplikasi Kelola Stok Buku

1. Menerima input ID buku, kemudian mencari ID buku dalam suatu array buku. Jika ada,
di layar akan tampil “ok” dan fungsi/prosedur tambah stok buku dapat diproses. Jika
tidak, di layar akan tampil “ID tidak dikenali” dan fungsi/prosedur tambah buku baru
dapat diproses.

2. Menambahkan buku baru dengan identitas: ID, judul, penerbit, tahun, stok.

3. Menampilkan identitas suatu buku dengan ID tertentu.

4. Menambah stok buku dengan ketentuan:
• Untuk buku dengan tahun terbit > 2010:
jika stok buku saat ini berjumlah antara 3-6, maka penambahan buku maksimal 5
buah. Jika stok buku saat ini berjumlah 7-9 penambahan buku hanya boleh 1 buah
buku. Jika stok buku saat ini berjumlah 10, maka tidak boleh ada penambahan buku.• Untuk buku dengan tahun terbit antara 2010-2018:
jika stok buku saat ini berjumlah antara 3-6, maka penambahan buku maksimal 7
buah. Jika stok buku saat ini berjumlah 7-9 penambahan buku hanya boleh 2 buah
buku. Jika stok buku saat ini berjumlah 15, maka tidak boleh ada penambahan buku.
• Untuk buku dengan tahun terbit antara < 2010:
Penambahan buku boleh berapa-pun.


## Case 2 : Aplikasi Rental Mobil

1. Menerima input data kartu anggota, disimpan dalam record pemesan.

2. Mencari data pemesan dengan inputan berupa ID anggota. Tampilkan data anggota
tersebut (ID, nama, umur, alamat, saldo, poin)

3. Menerima input tujuan. Menampilkan mobil yang dapat dipinjam dengan ketentuan:
1) Saldo harus diatas Rp50000.
2) Mobil yang tersedia ke kota A: Inova (Rp8000/km), Avanza (Rp5000/km), Alya
(Rp4000/km). Mobil yang tersedia ke kota B: Mobilio (Rp5000/km) dan Brio
(Rp4000/km).3) Jarak ke Kota A=50km, jarak ke Kota B=80km.
4) Mobil yang bisa dipinjam yaitu yang bisa ke kota tujuan dan saldo anggota cukup
untukmembayar perjalanan tersebut.

4. Menampilkan pilihan “ok” atau “cancel” setelah tampil pilihan mobil. Jika “ok”, ubah
saldo anggota dan tampilkan kembali.

5. Menghitung poin dengan ketentuan: 1) Jika mobil yang terpilih adalah Inova, maka
poin yang didapat sebesar 10 poin, 2) Avanza 5 poin, 3) Brio 5 poin. Tampilkan jumlah
poin setelah “ok”.

6. Menampilkan seluruh data anggota dan diiurutkan secara descending berdasarkan
nama anggota.

7. Menampilkan seluruh data anggota yang memiliki saldo > Rp 50000 dan diiurutkan
secara descending berdasarkan nama anggota.

## Case 3 : Setor Tunai ATM

Disimpan data nasabah di Bank XYZ yang terdiri dari nomor rekening, pin, NIK, nama, jenis rekening,
saldo, tanggal pendaftaran, tanggal transaksi terakhir, dan nominal transaksi terakhir. Seorang
nasabah di Bank XYZ hanya boleh memiliki satu rekening. Buat program untuk mengelola transaksi
setor tunai di Bank XYZ, dengan spesifikasi sebagai berikut:

1. Login Customer Service
Asumsi hanya ada satu akun untuk login Customer Service.
a. Pendaftaran nasabah baru
i. Input NIK, nama, jenis rekening, dan setoran awal.
ii. NIK, nama, jenis rekening, dan setoran awal tidak boleh kosong.
iii. NIK valid jika berupa 16 digit angka. NIK tidak boleh duplikat.
iv. Ada 3 jenis rekening, yaitu Silver, Gold, dan Platinum.
v. Setoran awal harus positif, minimal adalah 1000.
vi. Nomor rekening di-generate secara otomatis dengan format: XYZ-[S/G/P
sesuai jenis nasabah][nomor urut sesuai jenis rekening dalam 3 digit],
contoh: XYZ-S004 (artinya nasabah rekening Silver urutan ke-4).
vii. Pin di-generate secara otomatis berup 6 digit angka.
viii. Saldo untuk rekening tersebut diinisialisai dengan setoran awal.
ix. Tanggal pendaftaran dan tanggal transaksi terakhir diambil otomatis dari
tanggal sistem saat ini.
x. Nominal transaksi terakhir = setoran awal.
xi. Disimpan data nasabah berupa nomor rekening, pin, NIK, nama, jenis
rekening, saldo, tanggal pendaftaran, tanggal transaksi terakhir, dan nominal
transaksi terakhir.

2. Login Nasabah
Nasabah login dengan menggunakan nomor rekening dan pin.
a. Setor tunai
i. Input jumlah setoran.
ii. Jumlah setoran harus positif, kelipatan 100, maksimum per transaksi 10.000
untuk Silver, 20.000 untuk Gold, dan 50.000 untuk Platinum.
iii. Saldo untuk rekening tersebut bertambah sesuai jumlah setoran.
iv. Tanggal transaksi terakhir berubah, diambil otomatis dari tanggal sistem
saat ini.
v. Nominal transaksi terakhir = jumlah setoran.

3. Login Manager
Asumsi hanya ada satu akun untuk login Manager.
a. Menampilkan satu data nasabah tertentu berdasarkan nomor rekening.
b. Menampilkan jumlah dan data seluruh nasabah secara terurut berdasarkan nomor
rekening.
c. Menampilkan jumlah dan data semua nasabah yang tidak melakukan transaksi lebih
dari 8 bulan, secara terurut berdasarkan tanggal transaksi terakhir.
d. Menampilkan jumlah dan data nasabah per jenis rekening secara terurut
berdasarkan nomor rekening

5. Menampilkan buku yang jumlah stok-nya 7-9 buah dan diurutkan mulai dari stok
tertinggi ke terendah.

6. Mengubah semua buku yang memiliki nama penerbit yang sama. Contoh: semula
nama penerbitnya “X”, diubah menjadi “Y”. Tampilkan di layar bahwa nama penerbit
setelah diubah.

# BLU CLI Banking System

CLI banking sederhana berbasis file teks.

## Cara pakai

```
go run main.go create_account <name> <amount>
go run main.go transfer <from> <to> <amount>
go run main.go add_deposit <name> <amount>
go run main.go accrue_interest
```

Contoh:

```
go run main.go create_account Alice 1000
go run main.go create_account Rudolph 2000
go run main.go add_deposit Alice 500
go run main.go transfer Rudolph Alice 500
# tunggu 10 menit
go run main.go accrue_interest
```

## Struktur

```
main.go                     -> baca argumen, panggil services
services/                   -> aturan bisnis (validasi, cek saldo, dll)
repository/account/         -> baca/tulis file akun
repository/deposit/         -> baca/tulis file deposit + hitung bunga
filebase/<name>.txt         -> saldo akun
filebase/deposits/<name>.txt-> satu baris per deposit: <amount>|<timestamp>
```

Format file akun:

```
name:Alice
balance:1500.00
```

## Asumsi

- Nama akun hanya boleh huruf. Ini juga mencegah path traversal (misal `../`).
- Saldo awal boleh 0, tidak boleh negatif.
- Akun dengan nama yang sama tidak bisa dibuat dua kali.
- Transfer: jumlah harus > 0, tidak bisa ke akun sendiri, saldo pengirim harus cukup,
  dan kedua akun harus ada.
- Deposit **tidak** memotong saldo utama (sesuai aturan soal). Karena itu, setelah contoh
  di atas saldo Alice = 1000 + 500 (transfer) = 1500. Angka 2000 di bagian "Expected" soal
  hanya masuk akal kalau deposit juga menambah saldo utama, yang bertentangan dengan aturan
  "Deposits are NOT deducted from the main account balance".
- Bunga 1% per menit, compound: `amount × 1.01^menit`. Hanya menit penuh yang dihitung.
  Setelah dihitung, timestamp maju sebanyak menit tersebut, jadi menjalankan
  `accrue_interest` berkali-kali tidak menghitung bunga ganda, dan sisa detik tetap
  terhitung di run berikutnya.
- Contoh: deposit 500 setelah 10 menit menjadi 552.31.
- Nilai uang disimpan dengan `float64` dan dibulatkan 2 desimal saat ditulis. Untuk sistem
  produksi lebih aman memakai integer satuan terkecil (sen) atau tipe decimal.

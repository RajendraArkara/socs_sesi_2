package services

import (
	"time"

	"github.com/RajendraArkara/socs_sesi_2/repository/deposit"
)

// AccrueInterest menghitung bunga 1% per menit (compound) untuk semua deposit.
// Saldo utama tidak berubah.
func AccrueInterest() (int, error) {
	return deposit.AccrueAll(time.Now())
}

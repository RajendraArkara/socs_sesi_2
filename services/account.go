package services

import (
	"errors"
	"unicode"

	"github.com/RajendraArkara/socs_sesi_2/repository/account"
)

func CreateAccount(name string, amount float64) error {
	if len(name) == 0 {
		return errors.New("nama harus di isi")
	}

	for _, r := range name {
		if !unicode.IsLetter(r) {
			return errors.New("username tidak memenuhi syarat")
		}
	}

	if amount < 0 {
		return errors.New("tidak boleh kurang dari 0")
	}

	if account.FindAccount(name) {
		return errors.New("akun sudah ada")
	}

	return account.Create(name, amount)
}

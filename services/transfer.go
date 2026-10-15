package services

import (
	"errors"
	"strings"

	"github.com/RajendraArkara/socs_sesi_2/repository/account"
)

func Transfer(from string, to string, amount float64) error {
	from = strings.TrimSpace(from)
	to = strings.TrimSpace(to)

	if from == "" || to == "" {
		return errors.New("pengirim dan penerima tidak boleh kosong")
	}

	if amount <= 0 {
		return errors.New("transfer tidak boleh kosong atau kurang dari 0")
	}

	if from == to {
		return errors.New("tidak bisa transfer ke akun sendiri")
	}

	fromBalance, err := account.FindByName(from)
	if err != nil {
		return err
	}

	toBalance, err := account.FindByName(to)
	if err != nil {
		return err
	}

	if fromBalance < amount {
		return errors.New("saldo tidak cukup")
	}

	if err := account.Save(from, fromBalance-amount); err != nil {
		return err
	}

	return account.Save(to, toBalance+amount)
}

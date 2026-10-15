package services

import (
	"errors"
	"strings"
	"time"

	"github.com/RajendraArkara/socs_sesi_2/repository/account"
	"github.com/RajendraArkara/socs_sesi_2/repository/deposit"
)

func AddDeposit(name string, amount float64) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return errors.New("nama harus di isi")
	}

	if amount <= 0 {
		return errors.New("deposit harus lebih dari 0")
	}

	if !account.FindAccount(name) {
		return errors.New("akun tidak di temukan")
	}

	return deposit.Add(name, amount, time.Now())
}

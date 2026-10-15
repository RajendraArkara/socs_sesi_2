package account

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const accountsDir = "filebase"

func FindAccount(account string) bool {
	filePath := filepath.Join(accountsDir, account+".txt")
	_, err := os.Stat(filePath)
	return err == nil
}

func Create(name string, balance float64) error {
	err := os.MkdirAll(accountsDir, 0755)
	if err != nil {
		return err
	}

	return Save(name, balance)
}

func FindByName(name string) (float64, error) {
	path := filepath.Join(accountsDir, name+".txt")

	data, err := os.ReadFile(path)
	if err != nil {
		return 0, errors.New("akun tidak di temukan")
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "balance:") {
			angkaString := strings.TrimPrefix(line, "balance:")
			angkaString = strings.TrimSpace(angkaString)

			balance, err := strconv.ParseFloat(angkaString, 64)
			if err != nil {
				return 0, errors.New("format saldo tidak valid")
			}

			return balance, nil
		}
	}
	return 0, errors.New("saldo tidak di temukan di file akun")
}

func Save(name string, balance float64) error {
	path := filepath.Join(accountsDir, name+".txt")

	content := fmt.Sprintf("name:%s\nbalance:%.2f\n", name, balance)

	return os.WriteFile(path, []byte(content), 0644)
}

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RajendraArkara/socs_sesi_2/services"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go <command> [args]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create_account":
		if len(os.Args) != 4 {
			fmt.Println("usage: create_account <name> <amount>")
			os.Exit(1)
		}

		amount, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("amount harus angka")
			os.Exit(1)
		}

		if err := services.CreateAccount(os.Args[2], amount); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		fmt.Printf("akun %s berhasil dibuat dengan saldo %.2f\n", os.Args[2], amount)

	case "transfer":
		if len(os.Args) != 5 {
			fmt.Println("usage: transfer <from> <to> <amount>")
			os.Exit(1)
		}

		amount, err := strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			fmt.Println("amount harus angka")
			os.Exit(1)
		}

		if err := services.Transfer(os.Args[2], os.Args[3], amount); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		fmt.Printf("transfer %.2f dari %s ke %s berhasil\n", amount, os.Args[2], os.Args[3])

	case "add_deposit":
		if len(os.Args) != 4 {
			fmt.Println("usage: add_deposit <name> <amount>")
			os.Exit(1)
		}

		amount, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("amount harus angka")
			os.Exit(1)
		}

		if err := services.AddDeposit(os.Args[2], amount); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		fmt.Printf("deposit %.2f untuk %s berhasil ditambahkan\n", amount, os.Args[2])

	case "accrue_interest":
		if len(os.Args) != 2 {
			fmt.Println("usage: accrue_interest")
			os.Exit(1)
		}

		updated, err := services.AccrueInterest()
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		fmt.Printf("bunga dihitung untuk %d deposit\n", updated)

	default:
		fmt.Println("command tidak dikenal:", os.Args[1])
		os.Exit(1)
	}
}

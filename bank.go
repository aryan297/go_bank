package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeFile(Value float64) float64 {

	balanceText := fmt.Sprint(Value)

	err := os.WriteFile("balance.txt", []byte(balanceText), 0644)

	if err != nil {
		return 0
	}
	return Value
}

func readFile() (float64, error) {
	file, err := os.ReadFile("balance.txt")
	if err != nil {
		fmt.Println("Error in reading file")
		return 0, err
	}
	balance, err := strconv.ParseFloat(string(file), 64)
	//fmt.Println("Your balance is", string(file))
	return balance, err

}

func main() {
	accountBalance, err := readFile()
	if err != nil {
		fmt.Println("Error in reading file, please restart APP")
		return
	}
	for i := 0; i < 2; i++ {
		fmt.Println("Welcome to go bank ")
		fmt.Println("Please select the option")
		fmt.Println("1.Check balance")
		fmt.Println("2.Deposit")
		fmt.Println("3.Withdraw")
		fmt.Println("4.Exit")
		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			accountBalance, err := readFile()
			if err != nil {
				fmt.Println("Error in reading file")
				fmt.Println("sorry for the issue")
				return
			}
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Println("Enter the amount to deposit")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				return
			}
			accountBalance += depositAmount
			accountBalanceValue := writeFile(accountBalance)
			fmt.Println("Your balance is", accountBalanceValue)

		case 3:
			fmt.Println("Enter the amount to withdraw")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance")
			} else {
				accountBalance -= withdrawAmount
				accountBalanceValue := writeFile(accountBalance)
				fmt.Println("Your balance is", accountBalanceValue)
			}
		default:
			fmt.Println("Thank you for using go bank")

			return

		}

		/* 		if option == 1 {
		   			fmt.Println("Your balance is", accountBalance)
		   		} else if option == 2 {
		   			fmt.Println("Enter the amount to deposit")
		   			var depositAmount float64
		   			fmt.Scan(&depositAmount)
		   			if depositAmount <= 0 {
		   				fmt.Println("Invalid amount")
		   				return
		   			}
		   			accountBalance += depositAmount
		   			fmt.Println("Your balance is", accountBalance)
		   		} else if option == 3 {
		   			fmt.Println("Enter the amount to withdraw")
		   			var withdrawAmount float64
		   			fmt.Scan(&withdrawAmount)
		   			if withdrawAmount > accountBalance {
		   				fmt.Println("Insufficient balance")
		   			} else {
		   				accountBalance -= withdrawAmount
		   				fmt.Println("Your balance is", accountBalance)
		   			}
		   		} else {
		   			fmt.Println("Thank you for using go bank")
		   		} */
	}

}

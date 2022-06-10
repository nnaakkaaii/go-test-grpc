package main

import (
	"bufio"
	"fmt"
	"os"
	"test-grpc/service"
)

func main() {
	fmt.Println("start Rock-paper-scissors game.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("0: exit")
		fmt.Println("1: play game")
		fmt.Println("2: show match results")
		fmt.Println("3: notify messages")
		fmt.Println("4: sum values")
		fmt.Println("5: chat messages")
		fmt.Print("please enter > ")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "0":
			fmt.Println("bye.")
			goto M
		case "1":
			fmt.Println("Please enter Rock, Paper, or Scissors.")
			fmt.Println("1: Rock")
			fmt.Println("2: Paper")
			fmt.Println("3: Scissors")
			fmt.Print("please enter > ")

			service.PlayGame(scanner)
			continue
		case "2":
			fmt.Println("Here are your match results.")
			service.ReportMatchResults()
			continue
		case "3":
			fmt.Println("Please enter the number of notifications.")
			fmt.Print("please enter > ")
			service.NotifyMessages(scanner)
			continue
		case "4":
			fmt.Println("Please enter numbers (enter 0 if finished).")
			fmt.Println("please enter ↓")
			service.SumValues(scanner)
			continue
		case "5":
			fmt.Println("Please enter messages (no message if finished).")
			fmt.Println("please enter ↓")
			service.ChatMessages(scanner)
			continue
		default:
			fmt.Println("Invalid command.")
			continue
		}
	}
M:
}

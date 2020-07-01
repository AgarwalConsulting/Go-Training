package main

import "fmt"

func printBoard(board [3][3]string) {
	for i := 0; i < cap(board); i++ {
		row := board[i]
		for j := 0; j < cap(row); j++ {
			fmt.Print(" | ", row[j], " | ")
		}
		fmt.Println("")
	}
}

func main() {
	board := [3][3]string{
		{"-", "X", "-"},
		{"O", "-", "X"},
		{"-", "-", "-"},
	}

	fmt.Println("Center Left: ", board[1][0])

	printBoard(board)
}

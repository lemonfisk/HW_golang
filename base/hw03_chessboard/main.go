package main

import (
	"fmt"
	"strings"
)

func main() {
	var boardSize int
	var board strings.Builder

	fmt.Printf("Please enter the size of the board: ")
	fmt.Scanf("%d", &boardSize)
	if boardSize < 0 {
		fmt.Println("The size of the board must be greater than zero")
		return
	}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if (i+j)%2 == 0 {
				board.WriteString("0")
			} else {
				board.WriteString("#")
			}
		}
		if i < boardSize-1 {
			board.WriteString("\n")
		}
	}
	fmt.Printf(board.String())

}

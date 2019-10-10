package main

import (
	"fmt"
	"strings"
)

func PrintBoard(board [][]string) {
	for i:=0; i<len(board); i++ {
         fmt.Println(strings.Join(board[i]," "))
     }
}

func main() {

	board := [][]string{
			[]string{"_","_","_"},
			[]string{"_","_","_"},
			[]string{"_","_","_"},
		}
	PrintBoard(board)
}

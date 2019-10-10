package main

import (
	"fmt"
	"strings"
)

func PrintSlice(s []int){
	fmt.Printf("len=%d,cap=%d,s=%v\n",len(s),cap(s),s)
}

func PrintBoard(board [][]string){
    for i:=0;i<len(board);i++ {
  		fmt.Println(strings.Join(board[i]," "))
	}
}

func  main() {
	s := make([]int,0);
	PrintSlice(s);

	s = append(s,0);
	PrintSlice(s);

	s = append(s,1)
	PrintSlice(s);

	s = append(s,2,3,4,5)
	PrintSlice(s)

	board := make([][]string,0)
	PrintBoard(board)

	board = append(board,[]string{"_","_","_"})	
//	PrintBoard(board)

	board = append(board,[]string{"_","_","_"}) 
 //	PrintBoard(board)

	board = append(board,[]string{"_","_","_"})     
	PrintBoard(board)

}

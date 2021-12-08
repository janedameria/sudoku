package main

import "fmt"

const N int = 9

var board = [][]int{
	{7, 8, 0, 4, 0, 0, 1, 2, 0},
	{6, 0, 0, 0, 7, 5, 0, 0, 9},
	{0, 0, 0, 6, 0, 1, 0, 7, 8},
	{0, 0, 7, 0, 4, 0, 2, 6, 0},
	{0, 0, 1, 0, 5, 0, 9, 3, 0},
	{9, 0, 4, 0, 6, 0, 0, 0, 5},
	{0, 7, 0, 3, 0, 0, 0, 1, 2},
	{1, 2, 0, 0, 0, 7, 4, 0, 0},
	{0, 4, 9, 2, 0, 6, 0, 0, 7},
}

/* 
func printBoard(board [][]int) {
	for i := 0; i < len(board); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("----------------------")
		}

		for j := 0; j < len(board[0]); j++ {
			if j != len(board) - 1 {
				if j % 3 == 0 && j != 0 {
					fmt.Printf("| ")
				}

				fmt.Printf("%v ", board[i][j])

			} else { 
				fmt.Printf("%v\n", board[i][j])
			}
		}
	}
}
*/

func printBoard(board [][] int){
	for row_i, row_v := range board {
		if row_i % 3 == 0 && row_i != 0 {
			fmt.Println("----------------------")
		}
		for col_i, col_v := range row_v {
			if col_i % 3 == 0 && col_i != 0 {
				fmt.Print("| ")
			} 
			fmt.Printf("%v ", col_v)
			if col_i == len(row_v) - 1{
				fmt.Print("\n")
			}
		}
	}

}

func findEmpty(board [][] int) []int{
	for i, row_v := range board {
		for j, col_v := range row_v {
			if col_v == 0 {
				return []int{i, j}
			}
		} 
	}

	return nil 
}

func solve(board [][] int) bool {
	value := findEmpty(board)

	if value == nil {
		fmt.Println("Sudoku has been solved")
		fmt.Println("Printing Result:")
		printBoard(board)
		return true;
	}

	//1. Check if there's empty value
	//2. IF theres a empty value then input a number ranging from 1-9
		//2.1 Check if its valid
		//2.2 If valid then input the number -> continue to solve itu
		//2.3 If not backtrack

		for num := 1; num <=9; num++ {
			isValid := valid(board, value, num)

			if isValid {
				board[value[0]][value[1]] = num

				if solve(board) {
					return true
				}

				board[value[0]][value[1]] = 0

			}
		}
		return false
}

func valid(board [][]int, pos []int, num int) bool{
	//checking on row
	for _, row_v := range board[pos[0]] {
		if row_v == num {
			return false
		}
	}

	//checking on col
	for _, row_v := range board{
		if row_v[pos[1]] == num {
			return false
		}
	}
	

	r := pos[0] / 3
	c := pos[1] / 3

	for i:=r*3; i< (r*3)+3; i++ {
		for j :=c*3; j <(c*3)+3; j++ {
				if board[i][j] == num {
					return false
				}
		}
	}

	return true

}

func printIndex(board [][]int){
	for row_i, row_val := range board {
				if row_i % 3 == 0 && row_i != 0 {
				fmt.Println("")
			}
		for col_i := range row_val {
			if col_i % 3 == 0 && col_i != 0 {
				fmt.Print("    |    ")
			}
			fmt.Printf("[%v, %v] = %v ", row_i, col_i, row_i + col_i)
		}
		fmt.Println("")
	}
}


func main() {
	fmt.Println("Board before : ")
	printBoard(board)
	fmt.Println()
	solve(board)
}

package main

import (
	."fmt"
)

func main() {

	var row int
	var col int
	Println("Enter Rows")
	_, _ = Scan(&row)
	Println("Enter Column")
	_, _ = Scan(&col)
	Println(row)
	Matrix := IntiMatrix(row,col)
	printMatrix(*Matrix, row,col)

}

func MakeMatrix(row ,col int)[][]int{

	Matrix := make([][]int, row)
	for i := range Matrix {
		Matrix[i] = make([]int, col)
	}
	return Matrix
}
func IntiMatrix(row ,col int) *[][]int {
	 Matrix :=MakeMatrix(row, col)
	for i :=0; i <row; i++{

		for j :=0; j < col; j++{
			Matrix[i][j] = (i+1 * j*2)/ 2
		}
	}
	return &Matrix
}
func printMatrix(Matrix [][] int, row int,col int){

	for i:=0; i < row; i++{
		Print("[")
		for j:=0; j < col;j++{
			Print(Matrix[i][j]," ")
		}
		Print("]")
		Print("\n")
	}
}

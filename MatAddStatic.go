package main

import "fmt"


//Allcoates Memory and returns matrix

func allocate_mat(size int, ch chan [][]int) {
	// Allocate Memory
mat := make([][]int, size)
	     for i := range mat {
		     mat[i] = make([]int, size)
	     }
     fmt.Println("mat[0][0] =", mat[0][0])
	     ch <- mat
}


func main() {
	var size int
		
		chA := make(chan [][]int)					//Creating channel for matrix A
		//	var b [][]int
		// 	var c [][]int

		//ch := make(chan bool)					//Channel to wait for allocation
		fmt.Print("enter size: ")
		fmt.Scan(&size)


		//Allcoating Memory of Matrices concurrently
		go allocate_mat( size, chA)
		//go allocate_mat(b, size, ch)
		//go allocate_mat(c, size, ch)

		//for i := 0 ; i < 3 ; i++{
		a := <-chA
		//}

		// array elements initialized to 0
		fmt.Println("a[0][0] =", a[0][0])

		// assign
		a[size-1][size-1] = 7

		// retrieve
		fmt.Printf("a[%d][%d] = %d\n", size-1, size-1, a[size-1][size-1])

		// remove only reference
		a = nil
		// memory allocated earlier with make can now be garbage collected.
}

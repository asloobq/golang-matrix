package main

import (
        "fmt"
        "math/rand"
       )

//Allocates Memory and returns matrix

func allocateMat(size int, ch chan [][]int) {
	// Allocate Memory
mat := make([][]int, size)
	     for i := range mat {
		     mat[i] = make([]int, size)
	     }
	     ch <- mat
}


func initializeMat(mat [][]int, size int, ch chan [][]int){
    //var seed int64
    //seed = 10
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
           // seed =  rand.Seed(seed);
            mat[i][j] = rand.Intn(10);
        }
    }

    ch <- mat
}

func print(mat [][]int, size int) {
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            fmt.Printf("ary[%d][%d] = %d  ", i, j, mat[i][j]);
        }
        fmt.Println("");
    }
}


func main() {
	var size int
		
		chA := make(chan [][]int)			//Creating channel for matrix A
		chB := make(chan [][]int)			//Creating channel for matrix B
		chC := make(chan [][]int)			//Creating channel for matrix C

        //Accept matrix size as input
		fmt.Print("enter size: ")
		fmt.Scan(&size)


		//Allcoating Memory of Matrices concurrently
		go allocateMat(size, chA)
		go allocateMat(size, chB)
		go allocateMat(size, chC)

		//Wait for allocation
		matA := <-chA
		matB := <-chB
        matC := <-chC

        // Initialize matrix A
        go initializeMat(matA, size, chA)
        go initializeMat(matB, size, chB)
        go initializeMat(matC, size, chC)

        //Wait for initialization
        matA = <-chA
    	matB = <-chB
        matC = <-chC
    
        //Print matrix A
        fmt.Println("Matrix A loaded with random numbers:")
        print(matA, size)
        
        fmt.Println("Matrix B loaded with random numbers:")
        print(matB, size)
        
        fmt.Println("Matrix C loaded with random numbers:")
		print(matC, size)
		
        
        //fmt.Println("a[0][0] =", a[0][0])

		// assign
		//a[size-1][size-1] = 7

		// retrieve
		//fmt.Printf("a[%d][%d] = %d\n", size-1, size-1, a[size-1][size-1])

		// remove only reference
		//a = nil
		// memory allocated earlier with make can now be garbage collected.
}

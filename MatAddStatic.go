package main

import (
        "fmt"
        "math/rand"
       )

//Allocates Memory and returns matrix

/*
 Calculates addition of the matrices A and B for row 
 numbers [start, end) 
 */
func addMatBlock(matA [][]int, matB [][]int, matC [][]int, 
        size int, start int, end int, ch chan bool){
    
    for i := start; i < end; i++ {
        for j := 0; j < size; j++ {
            matC[i][j] = matA[i][j] + matB[i][j]
        }
    }

    ch <- true
}

/*
  Breaks down matrix A and B into 'size / split' chunks.
   Calls 'addMatBlock' split number of times.

 */
func addMatStatic(matA [][]int, matB [][]int, matC [][]int, 
        size int, splits int, ch chan bool){
    
    //chBool := make(chan bool)
    //chunkSize := size / splits
    for i:= 0; i < splits; i++ {
        //addMatBlock(matA, matB, matC, size, i*chunkSize, 
        //        i*chunkSize + chunkSize, chBool)

    }

    ch <- true
}

func allocateMat(size int, ch chan [][]int) {
	// Allocate Memory
mat := make([][]int, size)
	     for i := range mat {
		     mat[i] = make([]int, size)
	     }
	     ch <- mat
}


func initializeMat(mat [][]int, size int, ch chan bool){
    
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            mat[i][j] = rand.Intn(10);
        }
    }

    ch <- true
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


        chABool := make(chan bool) 
        chBBool := make(chan bool)
       // chCBool := make(chan bool)

        // Initialize matrix A
        go initializeMat(matA, size, chABool)
        go initializeMat(matB, size, chBBool)
        //go initializeMat(matC, size, chCBool)

        //Wait for initialization
        <- chABool
        <- chBBool
        //<- chCBool

        //Print matrix A
        fmt.Println("Matrix A loaded with random numbers:")
        print(matA, size)
        
        fmt.Println("Matrix B loaded with random numbers:")
        print(matB, size)
        
        fmt.Println("Matrix C loaded with -1s:")
		print(matC, size)
    /*
        //Do addition
        go addMatBlock(matA, matB, matC, size, 0, size, chC)
        matC = <- chC        
        
        fmt.Println("Matrix C holds result after A[] + B[]:")
		print(matC, size)
    */
}

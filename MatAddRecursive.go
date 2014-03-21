/****************************************************************************/
//          MATRIX ADDITION WITH RECURSIVE DIVIDE AND CONQUER
/****************************************************************************/
package main

import (
        "fmt"
        "math/rand"
        "os"
        "strconv"
        "time"
       )

func addMatRecursive(matA [][]int, matB [][]int, matC [][]int, baseSize int,
        ch chan bool) {
    //Get current size
size := len(matA)
    if(size <= baseSize) {
        //Do serial addition
        for i := 0; i < size; i++ {
              for j := 0; j < size; j++ {
                  matC[i][j] = matA[i][j] + matB[i][j]
              }
          }
        ch <- true
    } else {

        //Recursive divide and conquer

        //End is mid point
        end := size / 2

        ch1 := make(chan bool)
        ch2 := make(chan bool)

        go addMatRecursive(matA[0:end], matB[0:end], matC[0:end], baseSize,
                ch1)
        go addMatRecursive(matA[end:size], matB[end:size], matC[end:size],
               baseSize, ch2)
        fmt.Printf("addMatRecursive size = %d | split from 0 to %d\n", 
                size, end) 
        //Wait for multiplications to complete
        <- ch1
        <- ch2 

        ch <-true
    }
}

func addMatRecursiveBegin(matA [][]int, matB [][]int, matC [][]int, splits int,
        ch chan bool){
size := len(matA)
          baseSize := size /splits
          chStart := make(chan bool)
          go addMatRecursive(matA, matB, matC, baseSize, chStart)
          <- chStart

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


func initializeMat(mat [][]int, ch chan bool){
size := len(mat)
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

t1 := time.Now()
chA := make(chan [][]int)			//Creating channel for matrix A
         chB := make(chan [][]int)			//Creating channel for matrix B
         chC := make(chan [][]int)			//Creating channel for matrix C

         //Accept matrix size as input

         //Commnad Line Arguments for splits and size of matrix
         if len(os.Args) == 3{

             size, _ := strconv.Atoi(os.Args[2])
                 splits, _ := strconv.Atoi(os.Args[1])
                 //Allcoating Memory of Matrices concurrently
                 go allocateMat(size, chA)
                 go allocateMat(size, chB)
                 go allocateMat(size, chC)

                 //Wait for allocation
                 matA := <-chA
                 matB := <-chB
                 matC := <-chC


                 chBool := make(chan bool) 
                 // chBBool := make(chan bool)
                 // chCBool := make(chan bool)

                 // Initialize matrix A
                 go initializeMat(matA, chBool)
                 go initializeMat(matB, chBool)
                 //go initializeMat(matC, chCBool)

                 //Wait for initialization
                 <- chBool
                 <- chBool
                 //<- chCBool

                 //Print matrix A
                 fmt.Println("Matrix A loaded with random numbers:")
                 print(matA, size)

                 fmt.Println("Matrix B loaded with random numbers:")
                 print(matB, size)

                 fmt.Println("Matrix C loaded with -1s:")
                 print(matC, size)

                 /********  ADDITION  ***********/

                 //Do addition
                 go addMatRecursive(matA, matB, matC, splits, chBool)
                 <- chBool

                 fmt.Println("Matrix C holds result after A[] + B[]:")
                 print(matC, size)

                 duration := time.Since(t1)
                 seconds := duration.Seconds()
                 fmt.Println("Running time = ", seconds, " s")
         } else {
             fmt.Println("usage is ./executable <splits> <size>")
         }
}

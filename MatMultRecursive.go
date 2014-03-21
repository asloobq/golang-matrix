/****************************************************************************/
//MATRIX MULTIPLICATION WITH RECURSIVE WORK DIVISION
/****************************************************************************/
package main

import (
        "fmt"
        "math/rand"
        "os"
        "strconv"
       )


func multMatRecursive(matA [][]int, matB [][]int, matC [][]int, baseSize int,
        startRow int, rowCount int, ch chan bool) {
    //Get current size
size := len(matA)
          if(rowCount <= baseSize) {
              //Do serial multiplication
endRow := startRow + rowCount

            for i := startRow; i < endRow; i++ {      
                for j := 0; j < size; j++ {

                    matC[i][j] = 0
                        for k:=0; k < size; k++ {
                            matC[i][j] += ( matA[i][k] * matB[k][j] )
                        }
                }
            }

        ch <- true
          } else {

              //Recursive divide and conquer

ch1 := make(chan bool)
         ch2 := make(chan bool)

         go multMatRecursive(matA, matB, matC, baseSize, startRow,
                 rowCount/2, ch1)
         go multMatRecursive(matA, matB, matC, baseSize, startRow + (rowCount/2),
                 rowCount - (rowCount /2), ch2)
         fmt.Printf("multMatRecursive size = %d | split from %d to %d\n", 
                 size, startRow, rowCount/2) 
         //Wait for multiplications to complete
         <- ch1
         <- ch2 

         ch <-true
          }
}

func multMatRecursiveBegin(matA [][]int, matB [][]int, matC [][]int, 
        splits int, ch chan bool) {
size := len(matA)
          baseSize := size /splits
          chStart := make(chan bool)
          go multMatRecursive(matA, matB, matC, baseSize, 0, size, chStart)
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

                 /********  MULTIPLICATION  ***********/

                 //Do multiplication
                 go multMatRecursiveBegin(matA, matB, matC, splits, chBool)
                 <- chBool

                 fmt.Println("Matrix C holds result after A[] * B[]:")
                 print(matC, size)
         } else {
             fmt.Println("usage is ./executable <splits> <size>")
         }
}

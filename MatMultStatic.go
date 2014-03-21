/****************************************************************************/
//MATRIX MULTIPLICATION WITH STATIC WORK DIVISION
/****************************************************************************/
package main

import (
        "fmt"
        "math/rand"
        "os"
        "strconv"
       )

//Allocates Memory and returns matrix

/*
   Calculates addition of the matrices A and B for row 
   numbers [start, end) 
 */
func multMatBlock(matA [][]int, matB [][]int, matC [][]int, startRow int, 
        endRow int, ch chan bool) {
cols := len(matA)
          for i := startRow; i < endRow; i++ {      
              for j := 0; j < cols; j++ {

                  matC[i][j] = 0
                      for k:=0; k < cols; k++ {
                          matC[i][j] += ( matA[i][k] * matB[k][j] )
                      }
              }
          }

      ch <- true
}

/*
   Breaks down matrix A and B into 'size / split' chunks.
   Calls 'multMatBlock' split number of times.

 */
func multMatStatic(matA [][]int, matB [][]int, matC [][]int, splits int,
        ch chan bool){

size := len(matA)   //number of rows and columns
          chBool := make(chan bool)
          chunkSize := size / splits  //rows to compute by 1 go routine

          fmt.Printf("multMatStatic chunkSize = %d \n", chunkSize)

          i := 0
          for ; i <= (splits-2); i++ {

start := i*chunkSize
           end := start + chunkSize
           go multMatBlock(matA, matB, matC, start, end, chBool)
           fmt.Printf("multMatStatic i = %d \n", i)
          }

      //handle last split separately
start := i*chunkSize
           go multMatBlock(matA, matB, matC, start, size, chBool)
           fmt.Printf("multMatStatic i = %d \n", i)

           for i:= 0; i < splits; i++ {
               <- chBool
                   fmt.Printf("multMatStatic sync i = %d\n", i)
           }

       ch <- true;

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
                 go multMatStatic(matA, matB, matC, splits, chBool)
                 <- chBool

                 fmt.Println("Matrix C holds result after A[] * B[]:")
                 print(matC, size)
         } else {
             fmt.Println("usage is ./executable <splits> <size>")
         }
}

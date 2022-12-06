/*
Created By Anwar Sanusi
*/

package main

import (
   "fmt";
   "math";
)

var (
   x  float64
)

//============SUDUT COSINUS, SINUS, DAN TANGEN=========//
func SdtCoSinTan(){
   fmt.Print("Masukkan nilai sudut = ")
   fmt.Scanf("%f", &x)
   sin := math.Sin(x)
   cos := math.Cos(x)
   tan := math.Tan(x)

   fmt.Println("Hasil hitung sudut Sin = ", sin)
   fmt.Println("Hasil hitung sudut Cos = ", cos)
   fmt.Println("Hasil hitung sudut Tan = ", tan)
}

func main() {
   fmt.Println("=================WELCOME================")
   SdtCoSinTan()
}
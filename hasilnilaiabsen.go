package main

import "fmt"
func main(){
	var nilai int
fmt.Print("Masukkan Nilai: ")
fmt.Scan(&nilai)

var absen int
fmt.Print("Masukkan Jumlah Absen: ")
fmt.Scan(&absen)

if nilai >70 && absen <5 {
	fmt.Println("Lulus")
}else{
	fmt.Println("Tidak Lulus")
}
}
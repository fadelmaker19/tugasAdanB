package main

import "fmt"

type karyawan struct {
	nama       string
	nik        int
	departemen string
}

func main() {
	var mhs karyawan
	fmt.Println("MEMASUKAN INPUTAN DATA KARYAWAN")

	for i := 0; i < 6; i++ {

		fmt.Println("Masukan Nama :")
		fmt.Scan(&mhs.nama)
		fmt.Println("Masukan Nik :")
		fmt.Scan(&mhs.nik)
		fmt.Println("Masukan Departemen :")
		fmt.Scan(&mhs.departemen)
	}

}

package main

import "fmt"

type sistemoperasi struct {
	jenis      string
	seri       string
	tahunrilis int
}

func main() {
	var kumpulan []sistemoperasi

	kumpulan = []sistemoperasi{
		sistemoperasi{
			jenis:      "MS-DOS",
			seri:       "81",
			tahunrilis: 1981,
		},
		sistemoperasi{
			jenis:      "Windows",
			seri:       "Windows 1.0",
			tahunrilis: 1985,
		},
		sistemoperasi{
			jenis:      "Windows",
			seri:       "Windows 3.1",
			tahunrilis: 1990,
		},
		sistemoperasi{
			jenis:      "Windows",
			seri:       "Windows 95",
			tahunrilis: 1995,
		},
		sistemoperasi{
			jenis:      "Windows",
			seri:       "Windows 1.0",
			tahunrilis: 1985,
		},
		sistemoperasi{
			jenis:      "Windows",
			seri:       "Windows 98",
			tahunrilis: 1998,
		},
		sistemoperasi{
			jenis:      "Windows",
			seri:       "Windows XP",
			tahunrilis: 2001,
		},
		sistemoperasi{
			jenis:      "Windows",
			seri:       "Windows ME",
			tahunrilis: 2000,
		},
		sistemoperasi{
			jenis:      "Windows",
			seri:       "Windows NT 31",
			tahunrilis: 1993,
		},
		sistemoperasi{
			jenis:      "Windows",
			seri:       "Windows 2000",
			tahunrilis: 2000,
		},
	}

	fmt.Println("-------Data Sistem Operasi--------")
	fmt.Println(kumpulan[0])
	fmt.Println(kumpulan[1])
	fmt.Println(kumpulan[2])
	fmt.Println(kumpulan[3])
	fmt.Println(kumpulan[4])
	fmt.Println(kumpulan[5])
	fmt.Println(kumpulan[6])
	fmt.Println(kumpulan[7])
	fmt.Println(kumpulan[8])
	fmt.Println(kumpulan[9])

}

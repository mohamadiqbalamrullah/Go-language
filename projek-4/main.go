package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Siswa struct {
	nomor     int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

type SiswaSlice struct {
	siswa []Siswa
}

func (s SiswaSlice) find(nomor int) (Siswa, bool) {
	for _, v := range s.siswa {
		if v.nomor == nomor {
			return v, true
		}
	}
	return Siswa{}, false
}

func main() {
	args := os.Args[1:]
	paraSiswa := SiswaSlice{
		siswa: []Siswa{
			{1, "Aku", "wakanda", "alien", "menguasai dunia dengan golang"},
			{2, "Dia", "bumi", "manusia", "Mau belajar golang ngalahin alien"},
			{3, "Kamu", "mars", "groot", "Mau belajar ngalahin"},
		},
	}
	for _, arg := range args {
		if a, err := strconv.ParseInt(arg, 10, 0); err == nil {
			if siswa, ok := paraSiswa.find(int(a)); ok {
				fmt.Println(strings.Repeat("=", 25))

				reflectValue := reflect.ValueOf(siswa)

				for i := 0; i < reflectValue.NumField(); i++ {
					fmt.Println(reflectValue.Type().Field(i).Name, ":", reflectValue.Field(i))
				}

			} else {
				fmt.Printf("siswa dengan nomor absen %d tidak ditemukan", a)
			}

		} else {
			fmt.Println("args harus berupa angka")
		}
	}
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Peserta struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	dataPeserta := [5]Peserta{}
	dataPeserta[0].nama = "Naily Khairiya"
	dataPeserta[0].alamat = "Surabaya"
	dataPeserta[0].pekerjaan = "Mahasiswa"
	dataPeserta[0].alasan = "Menambah Pengetahuan"

	dataPeserta[1].nama = "Muna"
	dataPeserta[1].alamat = "Nganjuk"
	dataPeserta[1].pekerjaan = "Mahasiswa"
	dataPeserta[1].alasan = "Menambah Ilmu"

	dataPeserta[2].nama = "Rani"
	dataPeserta[2].alasan = "Kediri"
	dataPeserta[2].pekerjaan = "Mahasiswa"
	dataPeserta[2].alamat = "Menambah Relasi"

	var argsRaw = os.Args
	var args, _ = strconv.Atoi(argsRaw[1])

	fmt.Printf("\nData Peserta absen %v\nNama : %s\n", args, dataPeserta[args-1].nama)
	fmt.Printf("Alamat : %s\n", dataPeserta[args-1].alamat)
	fmt.Printf("Pekerjaan : %s\n", dataPeserta[args-1].pekerjaan)
	fmt.Printf("Alamat : %s\n", dataPeserta[args-1].alamat)
}

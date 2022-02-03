package interfaces

import (
	"fmt"
)

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 1001 * asal.Angka
}

func main() {
	persegi := Persegi{Sisi: 5}
	luas := SeberapaLuas(persegi)
	fmt.Println("Luas Persegi:", luas)

	persegiPanjang := PersegiPanjang{Panjang: 10, Lebar: 8}
	luas = SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang:", luas)

	asal := Asal{Angka: 9}
	luas = SeberapaLuas(asal)
	fmt.Println("Luas Asal:", luas)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}

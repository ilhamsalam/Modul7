package main

import (
        "fmt"
        "math"
)

// Struktur untuk merepresentasikan titik
type Titik struct {
        x int
        y int
}

// Struktur untuk merepresentasikan lingkaran
type Lingkaran struct {
        pusat Titik
        radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func hitungJarak(a, b Titik) float64 {
        return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

// Fungsi untuk menentukan posisi titik terhadap dua lingkaran
func tentukanPosisi(titik Titik, lingkaran1, lingkaran2 Lingkaran) string {
        jarak1 := hitungJarak(titik, lingkaran1.pusat)
        jarak2 := hitungJarak(titik, lingkaran2.pusat)

        if jarak1 <= float64(lingkaran1.radius) && jarak2 <= float64(lingkaran2.radius) {
                return "Titik di dalam lingkaran 1 dan 2"
        } else if jarak1 <= float64(lingkaran1.radius) {
                return "Titik di dalam lingkaran 1"
        } else if jarak2 <= float64(lingkaran2.radius) {
                return "Titik di dalam lingkaran 2"
        } else {
                return "Titik di luar lingkaran 1 dan 2"
        }
}

func main() {
        var lingkaran1, lingkaran2 Lingkaran
        var titik Titik

        // Input koordinat (asumsi input sudah valid)
        fmt.Scan(&lingkaran1.pusat.x, &lingkaran1.pusat.y, &lingkaran1.radius)
        fmt.Scan(&lingkaran2.pusat.x, &lingkaran2.pusat.y, &lingkaran2.radius)
        fmt.Scan(&titik.x, &titik.y)

        // Panggil fungsi untuk menentukan posisi
        hasil := tentukanPosisi(titik, lingkaran1, lingkaran2)
        fmt.Println(hasil)
}
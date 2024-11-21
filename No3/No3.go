package main

import "fmt"

func main() {
        var klubA, klubB string
        var skorA, skorB int

        fmt.Print("Nama klub A: ")
        fmt.Scan(&klubA)
        fmt.Print("Nama klub B: ")
        fmt.Scan(&klubB)

        var pemenang []string

        for {
                fmt.Printf("Pertandingan: ")
                fmt.Scan(&skorA, &skorB)

                if skorA < 0 || skorB < 0 {
                        break
                }

                if skorA > skorB {
                        pemenang = append(pemenang, klubA)
                } else if skorB > skorA {
                        pemenang = append(pemenang, klubB)
                } else {
                        pemenang = append(pemenang, "Draw")
                }
        }

        fmt.Println("Hasil Pertandingan:")
        for i, hasil := range pemenang {
                fmt.Printf("Hasil %d: %s\n", i+1, hasil)
        }
}
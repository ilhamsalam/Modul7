# Latihan modul 7
Penjelasan singkat pada tiap latihan soal modul 7

## Setiap soal
Pada setiap soal saya menggunakan beberapa komponen :
- `Package main` agar program dapat dieksekusi
- `import fmt` agar dapat menggunakan beberapa operasi dasar bahasa program `go`
- `import math` agar dapat memasukkan fungsi matematika dalam program `go`
- `main()` sebagai titik awal program dan dieksekusi ketika program dijalankan

## Soal 1
```go
type Titik struct {
        x int
        y int
}
```
Operasi diatas merupakan operasi menyimpan nilai koordinat titik dari x dan y bertipe integer

```go
type Lingkaran struct {
        pusat Titik
        radius int
}
```
Operasi diatas merupakan operasi menyimpan nilai dari pusat yang terdiri dari titik x dan y bertipe integer serta radius atau jari-jari lingkaran yang bertipe integer

```go
func hitungJarak(a, b Titik) float64 {
        return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}
```
Operasi diatas merupakan operasi menghitung jarak antara kedua titik

```go
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
```
Operasi diatas merupakan operasi menghitung  jarak antara titik yang diberikan dengan pusat masing-masing lingkaran dan membandingkan jarak tersebut dengan jari-jari masing-masing lingkaran hasilnya menunjukkan posisi titik relatif terhadap kedua lingkaran

## Soal 2
```go
func rataRata(arr []int) float64 {
    sum := 0
    for _, val := range arr {
        sum += val
    }
    return float64(sum) / float64(len(arr))
}
```
Operasi diatas merupakan operasi menjumlahkan semua elemen dalam array kemudian membaginya dengan banyaknya elemen pada array

```go
func standarDeviasi(arr []int) float64 {
    mean := rataRata(arr)
    var sumSquares float64
    for _, val := range arr {
        sumSquares += math.Pow(float64(val)-mean, 2)
    }
    return math.Sqrt(sumSquares / float64(len(arr)))
}
```
Operasi diatas merupakan operasi menghitung selisih kuadrat antara setiap elemen dengan rata-rata kemudian jumlahkan semua selisih kuadrat terakhir hitung akar kuadrat dari rata-rata selisih kuadrat untuk mendapatkan standar deviasi.

```go
func frekuensi(arr []int, elemen int) int {
    count := 0
    for _, val := range arr {
        if val == elemen {
            count++
        }
    }
    return count
}
```
Operasi diatas merupakan operasi menghitung jumlah kemunculan suatu elemen tertentu dalam array dengan cara iterasi dan perbandingan.

## Soal 3
```go
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

```
Operasi diatas merupakan operasi menginput skor tiap pertandingan klub a melawan klub b dengan ketentuan jika salah satu klub memiliki skor lebih tinggi dari klub satunya maka dia dinyatakan sebagai pemenang pertandingan tersebut dan jika seri maka akan dianggap sebagai seri. Hasil pertandingan akan disimpan pada array pemenang dan proses penginputan akan berakhir jika salah satu skor tidak valid

## Soal 4
```go
func isiArray(t *tabel, n *int) {
    var char rune
    fmt.Println("Masukkan karakter (akhiri dengan '.'): ")
    for *n = 0; *n < NMAX; *n++ {
        fmt.Scanf("%c", &char)
        if char == '.' {
            break
        }
        t[*n] = char
    }
}
```
Operasi diatas merupakan operasi mengisi array dengan karakter yang diinputkan oleh pengguna. Fungsi ini akan berhenti ketika memasukkan karakter titik ('.').

```go
func cetakArray(t tabel, n int) {
    for i := 0; i < n; i++ {
        fmt.Printf("%c", t[i])
    }
    fmt.Println()
}
```
Operasi diatas merupakan operasi mencetak isi array ke layar.

```go
func balikanArray(t *tabel, n int) {
    for i := 0; i < n/2; i++ {
        t[i], t[n-i-1] = t[n-i-1], t[i]
    }
}
```
Operasi diatas merupakan operasi membalikkan urutan elemen dalam array. Operasi ini bekerja dengan menukar elemen pada indeks i dengan elemen pada indeks n-i-1, di mana n adalah panjang array.

```go
func palindrom(t tabel, n int) bool {
    for i := 0; i < n/2; i++ {
        if t[i] != t[n-i-1] {
            return false
        }
    }
    return true
}
```
Operasi diatas merupakan operasi  memeriksa apakah array merupakan palindrom. Operasi ini membandingkan karakter pada indeks i dengan karakter pada indeks n-i-1 untuk setiap indeks i hingga mencapai tengah array. Jika semua pasangan karakter sama, maka array tersebut adalah palindrom.

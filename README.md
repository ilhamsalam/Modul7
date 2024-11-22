# Latihan modul 7
Penjelasan singkat pada tiap latihan soal modul 7

## Setiap soal
Pada setiap soal saya menggunakan beberapa komponen :
- `Package main` agar program dapat dieksekusi
- `import fmt` agar dapat menggunakan beberapa operasi dasar bahasa program `go`
- `main()` sebagai titik awal program dan dieksekusi ketika program dijalankan

## Soal 1
```go
func fibonacci(n int) int {
        if n <= 1 {
                return n
        }
        return fibonacci(n-1) + fibonacci(n-2)
}
```
Operasi diatas merupakan operasi menghitung bilangan fibonacci yakni bilangan selanjutnya merupakan hasil penjumlahan dari dua angka sebelumnya

```go
 for i := 0; i <= n; i++ {
            fmt.Printf("%d ", fibonacci(i))
        }
        fmt.Println()
}
```
Operasi diatas merupakan operasi menghitung dan mencetak hasil deret fibonacci dari ke 0 sampai ke n

## Soal 2
```go
func printStars(n int) {
        if n == 1 {
                fmt.Println("*")
                return
        }
        printStars(n - 1)
        for i := 0; i < n; i++ {
                fmt.Print("*")
        }
        fmt.Println()
}


```
Operasi diatas merupakan operasi mencetak bintang dengan tinggi n dengan cara fungsi memanggil dirinya sendiri dengan n-1 sebagai argumen dan mencetak segitiga ke n sampai n=1

## Soal 3
```go
func findFactors(number, divisor int) {
        if divisor > number {
                return
        }
        if number%divisor == 0 {
                fmt.Print(divisor, " ")
        }
        findFactors(number, divisor+1)
}
```
Operasi diatas merupakan operasi mencari faktor bilangan dari suatu angka dengan cara jika suatu angka habis dibagi oleh pembagi atau hasilnya sama dengan nol maka pembagi adalah faktor angka tersebut. Hal ini dilakukan sampai pembagi lebih besar sampai angka yang dibagi

## Soal 4
```go
func printSequence(n int) {
        if n == 1 {
                fmt.Print(n, " ")
                return
        }
        fmt.Print(n, " ")
        printSequence(n - 1)
        if n > 1 {
                fmt.Print(n, " ")
        }
}
```
Operasi diatas merupakan operasi mencetak bilangan n ke 1 kemudian ke n lagi

## Soal 5
```go
func printOddNumbers(n int) {
        if n < 1 {
                return
        }
        printOddNumbers(n - 2)
        fmt.Print(n, " ")
}
```
Operasi diatas merupakan operasi mencetak bilangan ganjil dengan cara pertama melihat nilai n adalah genap maka dicari dulu nilai ganjil terdekat kemudian dijalankan fungsi mencari nilai ganjil dengan n-2 sebagai argumen dan dicetak sampai n kurang dari 1

## Soal 6
```go
func power(x, y int) int {
        if y == 0 {
                return 1
        } else if y > 0 {
                return x * power(x, y-1)
        } else {
                return 1 / power(x, -y)
        }
}
```
Operasi diatas merupakan operasi menghitung pangkat suatu bilangan dengan cara jika bilangan pangkat positif maka fungsi akan memanggil dirinya sendiri dengan x yang sama dan y-1 sebagai pangkat. Hasilnya kemudian dikalikan dengan x. Jika pangkat negatif maka maka fungsi menghitung pangkat positif dari x dengan pangkat -y dan kemudian membagi 1 dengan hasil tersebut sampai y = 0 dengan hasil 1

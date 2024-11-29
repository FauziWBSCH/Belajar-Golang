package main

import "fmt"

func tipe() {
	fmt.Println("TIPE DATA")
	angka := 10
	teks := "Halo world"
	desimal := 3.14
	boolean := true
	fmt.Printf("tipe data: %T\n", angka)
	fmt.Printf("tipe data: %T\n", teks)
	fmt.Printf("tipe data: %T\n", desimal)
	fmt.Printf("tipe data: %t\n", boolean)
}

func tambah(a int, b int) int {
	fmt.Println("OPERASI PERTAMBAHAN")
	return a + b
}

func kurang(a int, b int) int {
	fmt.Println("OPERASI PENGURANGAN")
	return a - b
}

func kali(a int, b int) int {
	fmt.Println("OPERASI PERKALIAN")
    return a * b
}

func bagi(a int, b int) int {
	fmt.Println("OPERASI PEMBAGIAN")
    return a / b
}


func loop() {
	fmt.Println("PERULANGAN")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func cabang() {
	fmt.Println("PERCABANGAN MENENTUKAN ANGKA POSITIF DAN NEGATIF")
	var angka int
	fmt.Print("masukkan angka: ")
	fmt.Scan(&angka)
	if angka > 0 {
		fmt.Println(angka,"Angka positif")
	} else if angka < 0 {
		fmt.Println(angka,"Angka negatif")
	} else {
		fmt.Println("Angka Nol")
	}
}


func swcase() {
	fmt.Println("PERCABANGAN SWITCH CASE")
	buah := "apple"
	switch buah {
		case "apple":
            fmt.Println("Buah ini adalah Apel")
        case "banana":
            fmt.Println("Buah ini adalah Pepaya")
        default:
            fmt.Println("Buah yang anda masukkan tidak terdaftar")
}
}



func main() {
	tipe()
	fmt.Printf("================================================\n")
	var angka1 int
	var angka2 int
	fmt.Print("masukkan angka 1 dan angka 2 (dipisahkan dengan spasi): ")
    fmt.Scan(&angka1, &angka2)
	hasil1 := tambah(angka1, angka2)
	fmt.Println(angka1, "+", angka2,"=",hasil1)

	hasil2 := kurang(angka1, angka2)
	fmt.Println(angka1, "-", angka2,"=",hasil2)

	hasil3 := kali(angka1, angka2)
	fmt.Println(angka1, "x", angka2,"=",hasil3)

	hasil4 := bagi(angka1, angka2)
	fmt.Println(angka1, ":", angka2,"=",hasil4)

	fmt.Printf("================================================\n")
	loop()
	fmt.Printf("================================================\n")
    cabang()
	fmt.Printf("================================================\n")
	swcase()
}